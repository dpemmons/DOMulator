HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 2.6 Common DOM interfaces](common-dom-interfaces.html) --- [Table of
Contents](./) --- [3 Semantics, structure, and APIs of HTML documents
→](dom.html)

1.  ::: {#toc-infrastructure}
    1.  [[2.7]{.secno} Safe passing of structured
        data](structured-data.html#safe-passing-of-structured-data)
        1.  [[2.7.1]{.secno} Serializable
            objects](structured-data.html#serializable-objects)
        2.  [[2.7.2]{.secno} Transferable
            objects](structured-data.html#transferable-objects)
        3.  [[2.7.3]{.secno} StructuredSerializeInternal (
            `value`{.variable}, `forStorage`{.variable} \[ ,
            `memory`{.variable} \]
            )](structured-data.html#structuredserializeinternal)
        4.  [[2.7.4]{.secno} StructuredSerialize ( `value`{.variable}
            )](structured-data.html#structuredserialize)
        5.  [[2.7.5]{.secno} StructuredSerializeForStorage (
            `value`{.variable}
            )](structured-data.html#structuredserializeforstorage)
        6.  [[2.7.6]{.secno} StructuredDeserialize (
            `serialized`{.variable}, `targetRealm`{.variable} \[ ,
            `memory`{.variable} \]
            )](structured-data.html#structureddeserialize)
        7.  [[2.7.7]{.secno} StructuredSerializeWithTransfer (
            `value`{.variable}, `transferList`{.variable}
            )](structured-data.html#structuredserializewithtransfer)
        8.  [[2.7.8]{.secno} StructuredDeserializeWithTransfer (
            `serializeWithTransferResult`{.variable},
            `targetRealm`{.variable}
            )](structured-data.html#structureddeserializewithtransfer)
        9.  [[2.7.9]{.secno} Performing serialization and transferring
            from other
            specifications](structured-data.html#performing-structured-clones-from-other-specifications)
        10. [[2.7.10]{.secno} Structured cloning
            API](structured-data.html#structured-cloning)
    :::

### [2.7]{.secno} Safe passing of structured data[](#safe-passing-of-structured-data){.self-link}

[]{#structured-clone}To support passing JavaScript objects, including
[platform
objects](https://webidl.spec.whatwg.org/#dfn-platform-object){#safe-passing-of-structured-data:platform-object
x-internal="platform-object"}, across
[realm](https://tc39.es/ecma262/#sec-code-realms){#safe-passing-of-structured-data:realm
x-internal="realm"} boundaries, this specification defines the following
infrastructure for serializing and deserializing objects, including in
some cases transferring the underlying data instead of copying it.
Collectively this serialization/deserialization process is known as
\"structured cloning\", although most APIs perform separate
serialization and deserialization steps. (With the notable exception
being the
[`structuredClone()`{#safe-passing-of-structured-data:dom-structuredclone}](#dom-structuredclone)
method.)

This section uses the terminology and typographic conventions from the
JavaScript specification.
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

#### [2.7.1]{.secno} [Serializable objects]{.dfn}[](#serializable-objects){.self-link} {#serializable-objects lt="serializable object" export=""}

::::::::::::::::: {.mdn-anno .wrapped}
MDN

:::: feature
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
[/developer.mozilla.org/en-US/docs/Glossary/Serializable_object](https://developer.mozilla.org/en-US/docs/Web/https://developer.mozilla.org/en-US/docs/Glossary/Serializable_object "Serializable objects are objects that can be serialized and later deserialized in any JavaScript environment ("realm"). This allows them to, for example, be stored on disk and later restored, or cloned with structuredClone(), or shared between workers using DedicatedWorkerGlobalScope.postMessage().")

::: support
[Firefox103+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome77+]{.chrome
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
:::::::::::::::::

[Serializable
objects](#serializable-objects){#serializable-objects:serializable-objects}
support being serialized, and later deserialized, in a way that is
independent of any given
[realm](https://tc39.es/ecma262/#sec-code-realms){#serializable-objects:realm
x-internal="realm"}. This allows them to be stored on disk and later
restored, or cloned across
[agent](https://tc39.es/ecma262/#sec-agents){#serializable-objects:agent
x-internal="agent"} and even [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#serializable-objects:agent-cluster
x-internal="agent-cluster"} boundaries.

Not all objects are [serializable
objects](#serializable-objects){#serializable-objects:serializable-objects-2},
and not all aspects of objects that are [serializable
objects](#serializable-objects){#serializable-objects:serializable-objects-3}
are necessarily preserved when they are serialized.

[Platform
objects](https://webidl.spec.whatwg.org/#dfn-platform-object){#serializable-objects:platform-object
x-internal="platform-object"} can be [serializable
objects](#serializable-objects){#serializable-objects:serializable-objects-4}
if their [primary
interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#serializable-objects:primary-interface
x-internal="primary-interface"} is decorated with the
[`[Serializable]`]{#serializable .dfn dfn-type="extended-attribute"
lt="Serializable"} IDL [extended
attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute){#serializable-objects:extended-attribute
x-internal="extended-attribute"}. Such interfaces must also define the
following algorithms:

[serialization steps]{#serialization-steps .dfn export=""}, taking a [platform object](https://webidl.spec.whatwg.org/#dfn-platform-object){#serializable-objects:platform-object-2 x-internal="platform-object"} `value`{.variable}, a [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#serializable-objects:record x-internal="record"} `serialized`{.variable}, and a boolean `forStorage`{.variable}

:   A set of steps that serializes the data in `value`{.variable} into
    fields of `serialized`{.variable}. The resulting data serialized
    into `serialized`{.variable} must be independent of any
    [realm](https://tc39.es/ecma262/#sec-code-realms){#serializable-objects:realm-2
    x-internal="realm"}.

    These steps may throw an exception if serialization is not possible.

    These steps may perform a
    [sub-serialization](#sub-serialization){#serializable-objects:sub-serialization}
    to serialize nested data structures. They should not call
    [StructuredSerialize](#structuredserialize){#serializable-objects:structuredserialize}
    directly, as doing so will omit the important `memory`{.variable}
    argument.

    The introduction of these steps should omit mention of the
    `forStorage`{.variable} argument if it is not relevant to the
    algorithm.

[deserialization steps]{#deserialization-steps .dfn export=""}, taking a [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#serializable-objects:record-2 x-internal="record"} `serialized`{.variable}, a [platform object](https://webidl.spec.whatwg.org/#dfn-platform-object){#serializable-objects:platform-object-3 x-internal="platform-object"} `value`{.variable}, and a [realm](https://tc39.es/ecma262/#sec-code-realms){#serializable-objects:realm-3 x-internal="realm"} `targetRealm`{.variable}

:   A set of steps that deserializes the data in
    `serialized`{.variable}, using it to set up `value`{.variable} as
    appropriate. `value`{.variable} will be a newly-created instance of
    the [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#serializable-objects:platform-object-4
    x-internal="platform-object"} type in question, with none of its
    internal data set up; setting that up is the job of these steps.

    These steps may throw an exception if deserialization is not
    possible.

    These steps may perform a
    [sub-deserialization](#sub-deserialization){#serializable-objects:sub-deserialization}
    to deserialize nested data structures. They should not call
    [StructuredDeserialize](#structureddeserialize){#serializable-objects:structureddeserialize}
    directly, as doing so will omit the important
    `targetRealm`{.variable} and `memory`{.variable} arguments.

It is up to the definition of individual platform objects to determine
what data is serialized and deserialized by these steps. Typically the
steps are very symmetric.

The
[`[Serializable]`{#serializable-objects:serializable}](#serializable)
extended attribute must take no arguments, and must only appear on an
interface. It must not appear more than once on an interface.

For a given [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#serializable-objects:platform-object-5
x-internal="platform-object"}, only the object\'s [primary
interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#serializable-objects:primary-interface-2
x-internal="primary-interface"} is considered during the
(de)serialization process. Thus, if inheritance is involved in defining
the interface, each
[`[Serializable]`{#serializable-objects:serializable-2}](#serializable)-annotated
interface in the inheritance chain needs to define standalone
[serialization
steps](#serialization-steps){#serializable-objects:serialization-steps}
and [deserialization
steps](#deserialization-steps){#serializable-objects:deserialization-steps},
including taking into account any important data that might come from
inherited interfaces.

::: example
Let\'s say we were defining a platform object `Person`, which had
associated with it two pieces of associated data:

- a name value, which is a string; and

- a best friend value, which is either another `Person` instance or
  null.

We could then define `Person` instances to be [serializable
objects](#serializable-objects){#serializable-objects:serializable-objects-5}
by annotating the `Person` interface with the
[`[Serializable]`{#serializable-objects:serializable-3}](#serializable)
[extended
attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute){#serializable-objects:extended-attribute-2
x-internal="extended-attribute"}, and defining the following
accompanying algorithms:

[serialization steps](#serialization-steps){#serializable-objects:serialization-steps-2}

:   1.  Set `serialized`{.variable}.\[\[Name\]\] to
        `value`{.variable}\'s associated name value.

    2.  Let `serializedBestFriend`{.variable} be the
        [sub-serialization](#sub-serialization){#serializable-objects:sub-serialization-2}
        of `value`{.variable}\'s associated best friend value.

    3.  Set `serialized`{.variable}.\[\[BestFriend\]\] to
        `serializedBestFriend`{.variable}.

[deserialization steps](#deserialization-steps){#serializable-objects:deserialization-steps-2}

:   1.  Set `value`{.variable}\'s associated name value to
        `serialized`{.variable}.\[\[Name\]\].

    2.  Let `deserializedBestFriend`{.variable} be the
        [sub-deserialization](#sub-deserialization){#serializable-objects:sub-deserialization-2}
        of `serialized`{.variable}.\[\[BestFriend\]\].

    3.  Set `value`{.variable}\'s associated best friend value to
        `deserializedBestFriend`{.variable}.
:::

Objects defined in the JavaScript specification are handled by the
[StructuredSerialize](#structuredserialize){#serializable-objects:structuredserialize-2}
abstract operation directly.

Originally, this specification defined the concept of \"cloneable
objects\", which could be cloned from one
[realm](https://tc39.es/ecma262/#sec-code-realms){#serializable-objects:realm-4
x-internal="realm"} to another. However, to better specify the behavior
of certain more complex situations, the model was updated to make the
serialization and deserialization explicit.

#### [2.7.2]{.secno} [Transferable objects]{.dfn}[](#transferable-objects){.self-link} {#transferable-objects lt="transferable object" export=""}

[Transferable
objects](#transferable-objects){#transferable-objects:transferable-objects}
support being transferred across
[agents](https://tc39.es/ecma262/#sec-agents){#transferable-objects:agent
x-internal="agent"}. Transferring is effectively recreating the object
while sharing a reference to the underlying data and then detaching the
object being transferred. This is useful to transfer ownership of
expensive resources. Not all objects are [transferable
objects](#transferable-objects){#transferable-objects:transferable-objects-2}
and not all aspects of objects that are [transferable
objects](#transferable-objects){#transferable-objects:transferable-objects-3}
are necessarily preserved when transferred.

Transferring is an irreversible and non-idempotent operation. Once an
object has been transferred, it cannot be transferred, or indeed used,
again.

[Platform
objects](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object
x-internal="platform-object"} can be [transferable
objects](#transferable-objects){#transferable-objects:transferable-objects-4}
if their [primary
interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#transferable-objects:primary-interface
x-internal="primary-interface"} is decorated with the
[`[Transferable]`]{#transferable .dfn dfn-type="extended-attribute"
lt="Transferable"} IDL [extended
attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute){#transferable-objects:extended-attribute
x-internal="extended-attribute"}. Such interfaces must also define the
following algorithms:

[transfer steps]{#transfer-steps .dfn export=""}, taking a [platform object](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object-2 x-internal="platform-object"} `value`{.variable} and a [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#transferable-objects:record x-internal="record"} `dataHolder`{.variable}

:   A set of steps that transfers the data in `value`{.variable} into
    fields of `dataHolder`{.variable}. The resulting data held in
    `dataHolder`{.variable} must be independent of any
    [realm](https://tc39.es/ecma262/#sec-code-realms){#transferable-objects:realm
    x-internal="realm"}.

    These steps may throw an exception if transferral is not possible.

[transfer-receiving steps]{#transfer-receiving-steps .dfn export=""}, taking a [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#transferable-objects:record-2 x-internal="record"} `dataHolder`{.variable} and a [platform object](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object-3 x-internal="platform-object"} `value`{.variable}

:   A set of steps that receives the data in `dataHolder`{.variable},
    using it to set up `value`{.variable} as appropriate.
    `value`{.variable} will be a newly-created instance of the [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object-4
    x-internal="platform-object"} type in question, with none of its
    internal data set up; setting that up is the job of these steps.

    These steps may throw an exception if it is not possible to receive
    the transfer.

It is up to the definition of individual platform objects to determine
what data is transferred by these steps. Typically the steps are very
symmetric.

The
[`[Transferable]`{#transferable-objects:transferable}](#transferable)
extended attribute must take no arguments, and must only appear on an
interface. It must not appear more than once on an interface.

For a given [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object-5
x-internal="platform-object"}, only the object\'s [primary
interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#transferable-objects:primary-interface-2
x-internal="primary-interface"} is considered during the transferring
process. Thus, if inheritance is involved in defining the interface,
each
[`[Transferable]`{#transferable-objects:transferable-2}](#transferable)-annotated
interface in the inheritance chain needs to define standalone [transfer
steps](#transfer-steps){#transferable-objects:transfer-steps} and
[transfer-receiving
steps](#transfer-receiving-steps){#transferable-objects:transfer-receiving-steps},
including taking into account any important data that might come from
inherited interfaces.

[Platform
objects](https://webidl.spec.whatwg.org/#dfn-platform-object){#transferable-objects:platform-object-6
x-internal="platform-object"} that are [transferable
objects](#transferable-objects){#transferable-objects:transferable-objects-5}
have a [\[\[Detached\]\]]{#detached .dfn dfn-for="platform object"
dfn-type="attribute"} internal slot. This is used to ensure that once a
platform object has been transferred, it cannot be transferred again.

Objects defined in the JavaScript specification are handled by the
[StructuredSerializeWithTransfer](#structuredserializewithtransfer){#transferable-objects:structuredserializewithtransfer}
abstract operation directly.

#### [2.7.3]{.secno} [StructuredSerializeInternal]{.dfn} ( `value`{.variable}, `forStorage`{.variable} \[ , `memory`{.variable} \] )[](#structuredserializeinternal){.self-link} {#structuredserializeinternal dfn-type="abstract-op" noexport="" lt="StructuredSerializeInternal"}

The
[StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal}
abstract operation takes as input a JavaScript value `value`{.variable}
and serializes it to a
[realm](https://tc39.es/ecma262/#sec-code-realms){#structuredserializeinternal:realm
x-internal="realm"}-independent form, represented here as a
[Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:record
x-internal="record"}. This serialized form has all the information
necessary to later deserialize into a new JavaScript value in a
different realm.

This process can throw an exception, for example when trying to
serialize un-serializable objects.

1.  If `memory`{.variable} was not supplied, let `memory`{.variable} be
    an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#structuredserializeinternal:ordered-map
    x-internal="ordered-map"}.

    The purpose of the `memory`{.variable} map is to avoid serializing
    objects twice. This ends up preserving cycles and the identity of
    duplicate objects in graphs.

2.  If `memory`{.variable}\[`value`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#structuredserializeinternal:map-exists
    x-internal="map-exists"}, then return
    `memory`{.variable}\[`value`{.variable}\].

3.  Let `deep`{.variable} be false.

4.  If `value`{.variable} is undefined, null, [a
    Boolean](https://tc39.es/ecma262/#sec-ecmascript-language-types-boolean-type){#structuredserializeinternal:js-boolean
    x-internal="js-boolean"}, [a
    Number](https://tc39.es/ecma262/#sec-ecmascript-language-types-number-type){#structuredserializeinternal:js-number
    x-internal="js-number"}, [a
    BigInt](https://tc39.es/ecma262/#sec-ecmascript-language-types-bigint-type){#structuredserializeinternal:js-bigint
    x-internal="js-bigint"}, or [a
    String](https://tc39.es/ecma262/#sec-ecmascript-language-types-string-type){#structuredserializeinternal:js-string
    x-internal="js-string"}, then return { \[\[Type\]\]: \"primitive\",
    \[\[Value\]\]: `value`{.variable} }.

5.  If `value`{.variable} [is a
    Symbol](https://tc39.es/ecma262/#sec-ecmascript-language-types-symbol-type){#structuredserializeinternal:js-symbol
    x-internal="js-symbol"}, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror
    x-internal="datacloneerror"}
    [`DOMException`{#structuredserializeinternal:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  Let `serialized`{.variable} be an uninitialized value.

7.  If `value`{.variable} has a \[\[BooleanData\]\] internal slot, then
    set `serialized`{.variable} to { \[\[Type\]\]: \"Boolean\",
    \[\[BooleanData\]\]: `value`{.variable}.\[\[BooleanData\]\] }.

8.  Otherwise, if `value`{.variable} has a \[\[NumberData\]\] internal
    slot, then set `serialized`{.variable} to { \[\[Type\]\]:
    \"Number\", \[\[NumberData\]\]:
    `value`{.variable}.\[\[NumberData\]\] }.

9.  Otherwise, if `value`{.variable} has a \[\[BigIntData\]\] internal
    slot, then set `serialized`{.variable} to { \[\[Type\]\]:
    \"BigInt\", \[\[BigIntData\]\]:
    `value`{.variable}.\[\[BigIntData\]\] }.

10. Otherwise, if `value`{.variable} has a \[\[StringData\]\] internal
    slot, then set `serialized`{.variable} to { \[\[Type\]\]:
    \"String\", \[\[StringData\]\]:
    `value`{.variable}.\[\[StringData\]\] }.

11. Otherwise, if `value`{.variable} has a \[\[DateValue\]\] internal
    slot, then set `serialized`{.variable} to { \[\[Type\]\]: \"Date\",
    \[\[DateValue\]\]: `value`{.variable}.\[\[DateValue\]\] }.

12. Otherwise, if `value`{.variable} has a \[\[RegExpMatcher\]\]
    internal slot, then set `serialized`{.variable} to { \[\[Type\]\]:
    \"RegExp\", \[\[RegExpMatcher\]\]:
    `value`{.variable}.\[\[RegExpMatcher\]\], \[\[OriginalSource\]\]:
    `value`{.variable}.\[\[OriginalSource\]\], \[\[OriginalFlags\]\]:
    `value`{.variable}.\[\[OriginalFlags\]\] }.

13. Otherwise, if `value`{.variable} has an \[\[ArrayBufferData\]\]
    internal slot, then:

    1.  If
        [IsSharedArrayBuffer](https://tc39.es/ecma262/#sec-issharedarraybuffer){#structuredserializeinternal:issharedarraybuffer
        x-internal="issharedarraybuffer"}(`value`{.variable}) is true,
        then:

        1.  If the [current settings
            object](webappapis.html#current-settings-object){#structuredserializeinternal:current-settings-object}\'s
            [cross-origin isolated
            capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#structuredserializeinternal:concept-settings-object-cross-origin-isolated-capability}
            is false, then throw a
            [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-2
            x-internal="datacloneerror"}
            [`DOMException`{#structuredserializeinternal:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

            This check is only needed when serializing (and not when
            deserializing) as the [cross-origin isolated
            capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#structuredserializeinternal:concept-settings-object-cross-origin-isolated-capability-2}
            cannot change over time and a
            [`SharedArrayBuffer`{#structuredserializeinternal:sharedarraybuffer}](https://tc39.es/ecma262/#sec-sharedarraybuffer-objects){x-internal="sharedarraybuffer"}
            cannot leave an [agent
            cluster](https://tc39.es/ecma262/#sec-agent-clusters){#structuredserializeinternal:agent-cluster
            x-internal="agent-cluster"}.

        2.  If `forStorage`{.variable} is true, then throw a
            [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-3
            x-internal="datacloneerror"}
            [`DOMException`{#structuredserializeinternal:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        3.  If `value`{.variable} has an
            \[\[ArrayBufferMaxByteLength\]\] internal slot, then set
            `serialized`{.variable} to { \[\[Type\]\]:
            \"GrowableSharedArrayBuffer\", \[\[ArrayBufferData\]\]:
            `value`{.variable}.\[\[ArrayBufferData\]\],
            \[\[ArrayBufferByteLengthData\]\]:
            `value`{.variable}.\[\[ArrayBufferByteLengthData\]\],
            \[\[ArrayBufferMaxByteLength\]\]:
            `value`{.variable}.\[\[ArrayBufferMaxByteLength\]\],
            \[\[AgentCluster\]\]: the [surrounding
            agent](https://tc39.es/ecma262/#surrounding-agent){#structuredserializeinternal:surrounding-agent
            x-internal="surrounding-agent"}\'s [agent
            cluster](https://tc39.es/ecma262/#sec-agent-clusters){#structuredserializeinternal:agent-cluster-2
            x-internal="agent-cluster"} }.

        4.  Otherwise, set `serialized`{.variable} to { \[\[Type\]\]:
            \"SharedArrayBuffer\", \[\[ArrayBufferData\]\]:
            `value`{.variable}.\[\[ArrayBufferData\]\],
            \[\[ArrayBufferByteLength\]\]:
            `value`{.variable}.\[\[ArrayBufferByteLength\]\],
            \[\[AgentCluster\]\]: the [surrounding
            agent](https://tc39.es/ecma262/#surrounding-agent){#structuredserializeinternal:surrounding-agent-2
            x-internal="surrounding-agent"}\'s [agent
            cluster](https://tc39.es/ecma262/#sec-agent-clusters){#structuredserializeinternal:agent-cluster-3
            x-internal="agent-cluster"} }.

    2.  Otherwise:

        1.  If
            [IsDetachedBuffer](https://tc39.es/ecma262/#sec-isdetachedbuffer){#structuredserializeinternal:isdetachedbuffer
            x-internal="isdetachedbuffer"}(`value`{.variable}) is true,
            then throw a
            [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-4
            x-internal="datacloneerror"}
            [`DOMException`{#structuredserializeinternal:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Let `size`{.variable} be
            `value`{.variable}.\[\[ArrayBufferByteLength\]\].

        3.  Let `dataCopy`{.variable} be ?
            [CreateByteDataBlock](https://tc39.es/ecma262/#sec-createbytedatablock){#structuredserializeinternal:createbytedatablock
            x-internal="createbytedatablock"}(`size`{.variable}).

            This can throw a
            [`RangeError`{#structuredserializeinternal:js-rangeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}
            exception upon allocation failure.

        4.  Perform
            [CopyDataBlockBytes](https://tc39.es/ecma262/#sec-copydatablockbytes){#structuredserializeinternal:copydatablockbytes
            x-internal="copydatablockbytes"}(`dataCopy`{.variable}, 0,
            `value`{.variable}.\[\[ArrayBufferData\]\], 0,
            `size`{.variable}).

        5.  If `value`{.variable} has an
            \[\[ArrayBufferMaxByteLength\]\] internal slot, then set
            `serialized`{.variable} to { \[\[Type\]\]:
            \"ResizableArrayBuffer\", \[\[ArrayBufferData\]\]:
            `dataCopy`{.variable}, \[\[ArrayBufferByteLength\]\]:
            `size`{.variable}, \[\[ArrayBufferMaxByteLength\]\]:
            `value`{.variable}.\[\[ArrayBufferMaxByteLength\]\] }.

        6.  Otherwise, set `serialized`{.variable} to { \[\[Type\]\]:
            \"ArrayBuffer\", \[\[ArrayBufferData\]\]:
            `dataCopy`{.variable}, \[\[ArrayBufferByteLength\]\]:
            `size`{.variable} }.

14. Otherwise, if `value`{.variable} has a \[\[ViewedArrayBuffer\]\]
    internal slot, then:

    1.  If
        [IsArrayBufferViewOutOfBounds](https://tc39.es/ecma262/#sec-isarraybufferviewoutofbounds){#structuredserializeinternal:isarraybufferviewoutofbounds
        x-internal="isarraybufferviewoutofbounds"}(`value`{.variable})
        is true, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-5
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializeinternal:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Let `buffer`{.variable} be the value of `value`{.variable}\'s
        \[\[ViewedArrayBuffer\]\] internal slot.

    3.  Let `bufferSerialized`{.variable} be ?
        [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-2}(`buffer`{.variable},
        `forStorage`{.variable}, `memory`{.variable}).

    4.  [Assert](https://infra.spec.whatwg.org/#assert){#structuredserializeinternal:assert
        x-internal="assert"}: `bufferSerialized`{.variable}.\[\[Type\]\]
        is \"ArrayBuffer\", \"ResizableArrayBuffer\",
        \"SharedArrayBuffer\", or \"GrowableSharedArrayBuffer\".

    5.  If `value`{.variable} has a \[\[DataView\]\] internal slot, then
        set `serialized`{.variable} to { \[\[Type\]\]:
        \"ArrayBufferView\", \[\[Constructor\]\]: \"DataView\",
        \[\[ArrayBufferSerialized\]\]: `bufferSerialized`{.variable},
        \[\[ByteLength\]\]: `value`{.variable}.\[\[ByteLength\]\],
        \[\[ByteOffset\]\]: `value`{.variable}.\[\[ByteOffset\]\] }.

    6.  Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#structuredserializeinternal:assert-2
            x-internal="assert"}: `value`{.variable} has a
            \[\[TypedArrayName\]\] internal slot.

        2.  Set `serialized`{.variable} to { \[\[Type\]\]:
            \"ArrayBufferView\", \[\[Constructor\]\]:
            `value`{.variable}.\[\[TypedArrayName\]\],
            \[\[ArrayBufferSerialized\]\]:
            `bufferSerialized`{.variable}, \[\[ByteLength\]\]:
            `value`{.variable}.\[\[ByteLength\]\], \[\[ByteOffset\]\]:
            `value`{.variable}.\[\[ByteOffset\]\], \[\[ArrayLength\]\]:
            `value`{.variable}.\[\[ArrayLength\]\] }.

15. Otherwise, if `value`{.variable} has \[\[MapData\]\] internal slot,
    then:

    1.  Set `serialized`{.variable} to { \[\[Type\]\]: \"Map\",
        \[\[MapData\]\]: a new empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list
        x-internal="js-list"} }.

    2.  Set `deep`{.variable} to true.

16. Otherwise, if `value`{.variable} has \[\[SetData\]\] internal slot,
    then:

    1.  Set `serialized`{.variable} to { \[\[Type\]\]: \"Set\",
        \[\[SetData\]\]: a new empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list-2
        x-internal="js-list"} }.

    2.  Set `deep`{.variable} to true.

17. Otherwise, if `value`{.variable} has an \[\[ErrorData\]\] internal
    slot and `value`{.variable} is not a [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#structuredserializeinternal:platform-object
    x-internal="platform-object"}, then:

    1.  Let `name`{.variable} be ?
        [Get](https://tc39.es/ecma262/#sec-get-o-p){#structuredserializeinternal:js-get
        x-internal="js-get"}(`value`{.variable}, \"name\").

    2.  If `name`{.variable} is not one of \"Error\", \"EvalError\",
        \"RangeError\", \"ReferenceError\", \"SyntaxError\",
        \"TypeError\", or \"URIError\", then set `name`{.variable} to
        \"Error\".

    3.  Let `valueMessageDesc`{.variable} be ?
        `value`{.variable}.\[\[GetOwnProperty\]\](\"`message`\").

    4.  Let `message`{.variable} be undefined if
        [IsDataDescriptor](https://tc39.es/ecma262/#sec-isdatadescriptor){#structuredserializeinternal:isdatadescriptor
        x-internal="isdatadescriptor"}(`valueMessageDesc`{.variable}) is
        false, and ?
        [ToString](https://tc39.es/ecma262/#sec-tostring){#structuredserializeinternal:tostring
        x-internal="tostring"}(`valueMessageDesc`{.variable}.\[\[Value\]\])
        otherwise.

    5.  Set `serialized`{.variable} to { \[\[Type\]\]: \"Error\",
        \[\[Name\]\]: `name`{.variable}, \[\[Message\]\]:
        `message`{.variable} }.

    6.  User agents should attach a serialized representation of any
        interesting accompanying data which are not yet specified,
        notably the `stack` property, to `serialized`{.variable}.

        See the Error Stacks proposal for in-progress work on specifying
        this data.
        [\[JSERRORSTACKS\]](references.html#refsJSERRORSTACKS)

18. Otherwise, if `value`{.variable} is an Array exotic object, then:

    1.  Let `valueLenDescriptor`{.variable} be ?
        [OrdinaryGetOwnProperty](https://tc39.es/ecma262/#sec-ordinarygetownproperty){#structuredserializeinternal:ordinarygetownproperty
        x-internal="ordinarygetownproperty"}(`value`{.variable},
        \"`length`\").

    2.  Let `valueLen`{.variable} be
        `valueLenDescriptor`{.variable}.\[\[Value\]\].

    3.  Set `serialized`{.variable} to { \[\[Type\]\]: \"Array\",
        \[\[Length\]\]: `valueLen`{.variable}, \[\[Properties\]\]: a new
        empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list-3
        x-internal="js-list"} }.

    4.  Set `deep`{.variable} to true.

19. Otherwise, if `value`{.variable} is a [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#structuredserializeinternal:platform-object-2
    x-internal="platform-object"} that is a [serializable
    object](#serializable-objects){#structuredserializeinternal:serializable-objects}:

    1.  If `value`{.variable} has a
        [\[\[Detached\]\]](#detached){#structuredserializeinternal:detached}
        internal slot whose value is true, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-6
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializeinternal:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Let `typeString`{.variable} be the identifier of the [primary
        interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#structuredserializeinternal:primary-interface
        x-internal="primary-interface"} of `value`{.variable}.

    3.  Set `serialized`{.variable} to { \[\[Type\]\]:
        `typeString`{.variable} }.

    4.  Set `deep`{.variable} to true.

20. Otherwise, if `value`{.variable} is a [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#structuredserializeinternal:platform-object-3
    x-internal="platform-object"}, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-7
    x-internal="datacloneerror"}
    [`DOMException`{#structuredserializeinternal:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

21. Otherwise, if
    [IsCallable](https://tc39.es/ecma262/#sec-iscallable){#structuredserializeinternal:iscallable
    x-internal="iscallable"}(`value`{.variable}) is true, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-8
    x-internal="datacloneerror"}
    [`DOMException`{#structuredserializeinternal:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

22. Otherwise, if `value`{.variable} has any internal slot other than
    \[\[Prototype\]\], \[\[Extensible\]\], or \[\[PrivateElements\]\],
    then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-9
    x-internal="datacloneerror"}
    [`DOMException`{#structuredserializeinternal:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    For instance, a \[\[PromiseState\]\] or \[\[WeakMapData\]\] internal
    slot.

23. Otherwise, if `value`{.variable} is an exotic object and
    `value`{.variable} is not the
    [%Object.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-object-prototype-object){#structuredserializeinternal:object.prototype
    x-internal="object.prototype"} intrinsic object associated with any
    [realm](https://tc39.es/ecma262/#sec-code-realms){#structuredserializeinternal:realm-2
    x-internal="realm"}, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializeinternal:datacloneerror-10
    x-internal="datacloneerror"}
    [`DOMException`{#structuredserializeinternal:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    For instance, a proxy object.

24. Otherwise:

    1.  Set `serialized`{.variable} to { \[\[Type\]\]: \"Object\",
        \[\[Properties\]\]: a new empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list-4
        x-internal="js-list"} }.

    2.  Set `deep`{.variable} to true.

    [%Object.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-object-prototype-object){#structuredserializeinternal:object.prototype-2
    x-internal="object.prototype"} will end up being handled via this
    step and subsequent steps. The end result is that its exoticness is
    ignored, and after deserialization the result will be an empty
    object (not an [immutable prototype exotic
    object](https://tc39.es/ecma262/#immutable-prototype-exotic-object){#structuredserializeinternal:immutable-prototype-exotic-object
    x-internal="immutable-prototype-exotic-object"}).

25. [Set](https://infra.spec.whatwg.org/#map-set){#structuredserializeinternal:map-set
    x-internal="map-set"} `memory`{.variable}\[`value`{.variable}\] to
    `serialized`{.variable}.

26. If `deep`{.variable} is true, then:

    1.  If `value`{.variable} has a \[\[MapData\]\] internal slot, then:

        1.  Let `copiedList`{.variable} be a new empty
            [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list-5
            x-internal="js-list"}.

        2.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializeinternal:list-iterate
            x-internal="list-iterate"}
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:record-2
            x-internal="record"} { \[\[Key\]\], \[\[Value\]\] }
            `entry`{.variable} of `value`{.variable}.\[\[MapData\]\]:

            1.  Let `copiedEntry`{.variable} be a new
                [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:record-3
                x-internal="record"} { \[\[Key\]\]:
                `entry`{.variable}.\[\[Key\]\], \[\[Value\]\]:
                `entry`{.variable}.\[\[Value\]\] }.

            2.  If `copiedEntry`{.variable}.\[\[Key\]\] is not the
                special value *empty*,
                [append](https://infra.spec.whatwg.org/#list-append){#structuredserializeinternal:list-append
                x-internal="list-append"} `copiedEntry`{.variable} to
                `copiedList`{.variable}.

        3.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializeinternal:list-iterate-2
            x-internal="list-iterate"}
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:record-4
            x-internal="record"} { \[\[Key\]\], \[\[Value\]\] }
            `entry`{.variable} of `copiedList`{.variable}:

            1.  Let `serializedKey`{.variable} be ?
                [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-3}(`entry`{.variable}.\[\[Key\]\],
                `forStorage`{.variable}, `memory`{.variable}).

            2.  Let `serializedValue`{.variable} be ?
                [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-4}(`entry`{.variable}.\[\[Value\]\],
                `forStorage`{.variable}, `memory`{.variable}).

            3.  [Append](https://infra.spec.whatwg.org/#list-append){#structuredserializeinternal:list-append-2
                x-internal="list-append"} { \[\[Key\]\]:
                `serializedKey`{.variable}, \[\[Value\]\]:
                `serializedValue`{.variable} } to
                `serialized`{.variable}.\[\[MapData\]\].

    2.  Otherwise, if `value`{.variable} has a \[\[SetData\]\] internal
        slot, then:

        1.  Let `copiedList`{.variable} be a new empty
            [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:js-list-6
            x-internal="js-list"}.

        2.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializeinternal:list-iterate-3
            x-internal="list-iterate"} `entry`{.variable} of
            `value`{.variable}.\[\[SetData\]\]:

            1.  If `entry`{.variable} is not the special value *empty*,
                [append](https://infra.spec.whatwg.org/#list-append){#structuredserializeinternal:list-append-3
                x-internal="list-append"} `entry`{.variable} to
                `copiedList`{.variable}.

        3.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializeinternal:list-iterate-4
            x-internal="list-iterate"} `entry`{.variable} of
            `copiedList`{.variable}:

            1.  Let `serializedEntry`{.variable} be ?
                [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-5}(`entry`{.variable},
                `forStorage`{.variable}, `memory`{.variable}).

            2.  [Append](https://infra.spec.whatwg.org/#list-append){#structuredserializeinternal:list-append-4
                x-internal="list-append"} `serializedEntry`{.variable}
                to `serialized`{.variable}.\[\[SetData\]\].

    3.  Otherwise, if `value`{.variable} is a [platform
        object](https://webidl.spec.whatwg.org/#dfn-platform-object){#structuredserializeinternal:platform-object-4
        x-internal="platform-object"} that is a [serializable
        object](#serializable-objects){#structuredserializeinternal:serializable-objects-2},
        then perform the [serialization
        steps](#serialization-steps){#structuredserializeinternal:serialization-steps}
        for `value`{.variable}\'s [primary
        interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#structuredserializeinternal:primary-interface-2
        x-internal="primary-interface"}, given `value`{.variable},
        `serialized`{.variable}, and `forStorage`{.variable}.

        The [serialization
        steps](#serialization-steps){#structuredserializeinternal:serialization-steps-2}
        may need to perform a [sub-serialization]{#sub-serialization
        .dfn export=""}. This is an operation which takes as input a
        value `subValue`{.variable}, and returns
        [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-6}(`subValue`{.variable},
        `forStorage`{.variable}, `memory`{.variable}). (In other words,
        a
        [sub-serialization](#sub-serialization){#structuredserializeinternal:sub-serialization}
        is a specialization of
        [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-7}
        to be consistent within this invocation.)

    4.  Otherwise, for each `key`{.variable} in !
        [EnumerableOwnProperties](https://tc39.es/ecma262/#sec-enumerableownproperties){#structuredserializeinternal:enumerableownproperties
        x-internal="enumerableownproperties"}(`value`{.variable}, key):

        1.  If !
            [HasOwnProperty](https://tc39.es/ecma262/#sec-hasownproperty){#structuredserializeinternal:hasownproperty
            x-internal="hasownproperty"}(`value`{.variable},
            `key`{.variable}) is true, then:

            1.  Let `inputValue`{.variable} be ?
                `value`{.variable}.\[\[Get\]\](`key`{.variable},
                `value`{.variable}).

            2.  Let `outputValue`{.variable} be ?
                [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-8}(`inputValue`{.variable},
                `forStorage`{.variable}, `memory`{.variable}).

            3.  [Append](https://infra.spec.whatwg.org/#list-append){#structuredserializeinternal:list-append-5
                x-internal="list-append"} { \[\[Key\]\]:
                `key`{.variable}, \[\[Value\]\]:
                `outputValue`{.variable} } to
                `serialized`{.variable}.\[\[Properties\]\].

27. Return `serialized`{.variable}.

::: example
It\'s important to realize that the
[Records](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializeinternal:record-5
x-internal="record"} produced by
[StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-9}
might contain \"pointers\" to other records that create circular
references. For example, when we pass the following JavaScript object
into
[StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeinternal:structuredserializeinternal-10}:

``` js
const o = {};
o.myself = o;
```

it produces the following result:

    {
      [[Type]]: "Object",
      [[Properties]]: «
        {
          [[Key]]: "myself",
          [[Value]]: <a pointer to this whole structure>
        }
      »
    }
:::

#### [2.7.4]{.secno} [StructuredSerialize]{.dfn} ( `value`{.variable} )[](#structuredserialize){.self-link} {#structuredserialize dfn-type="abstract-op" lt="StructuredSerialize"}

1.  Return ?
    [StructuredSerializeInternal](#structuredserializeinternal){#structuredserialize:structuredserializeinternal}(`value`{.variable},
    false).

#### [2.7.5]{.secno} [StructuredSerializeForStorage]{.dfn} ( `value`{.variable} )[](#structuredserializeforstorage){.self-link} {#structuredserializeforstorage dfn-type="abstract-op" lt="StructuredSerializeForStorage"}

1.  Return ?
    [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializeforstorage:structuredserializeinternal}(`value`{.variable},
    true).

#### [2.7.6]{.secno} [StructuredDeserialize]{.dfn} ( `serialized`{.variable}, `targetRealm`{.variable} \[ , `memory`{.variable} \] )[](#structureddeserialize){.self-link} {#structureddeserialize dfn-type="abstract-op" lt="StructuredDeserialize"}

The
[StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize}
abstract operation takes as input a
[Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:record
x-internal="record"} `serialized`{.variable}, which was previously
produced by
[StructuredSerialize](#structuredserialize){#structureddeserialize:structuredserialize}
or
[StructuredSerializeForStorage](#structuredserializeforstorage){#structureddeserialize:structuredserializeforstorage},
and deserializes it into a new JavaScript value, created in
`targetRealm`{.variable}.

This process can throw an exception, for example when trying to allocate
memory for the new objects (especially `ArrayBuffer` objects).

1.  If `memory`{.variable} was not supplied, let `memory`{.variable} be
    an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#structureddeserialize:ordered-map
    x-internal="ordered-map"}.

    The purpose of the `memory`{.variable} map is to avoid deserializing
    objects twice. This ends up preserving cycles and the identity of
    duplicate objects in graphs.

2.  If `memory`{.variable}\[`serialized`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#structureddeserialize:map-exists
    x-internal="map-exists"}, then return
    `memory`{.variable}\[`serialized`{.variable}\].

3.  Let `deep`{.variable} be false.

4.  Let `value`{.variable} be an uninitialized value.

5.  If `serialized`{.variable}.\[\[Type\]\] is \"primitive\", then set
    `value`{.variable} to `serialized`{.variable}.\[\[Value\]\].

6.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Boolean\",
    then set `value`{.variable} to a new Boolean object in
    `targetRealm`{.variable} whose \[\[BooleanData\]\] internal slot
    value is `serialized`{.variable}.\[\[BooleanData\]\].

7.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Number\",
    then set `value`{.variable} to a new Number object in
    `targetRealm`{.variable} whose \[\[NumberData\]\] internal slot
    value is `serialized`{.variable}.\[\[NumberData\]\].

8.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"BigInt\",
    then set `value`{.variable} to a new BigInt object in
    `targetRealm`{.variable} whose \[\[BigIntData\]\] internal slot
    value is `serialized`{.variable}.\[\[BigIntData\]\].

9.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"String\",
    then set `value`{.variable} to a new String object in
    `targetRealm`{.variable} whose \[\[StringData\]\] internal slot
    value is `serialized`{.variable}.\[\[StringData\]\].

10. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Date\", then
    set `value`{.variable} to a new Date object in
    `targetRealm`{.variable} whose \[\[DateValue\]\] internal slot value
    is `serialized`{.variable}.\[\[DateValue\]\].

11. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"RegExp\",
    then set `value`{.variable} to a new RegExp object in
    `targetRealm`{.variable} whose \[\[RegExpMatcher\]\] internal slot
    value is `serialized`{.variable}.\[\[RegExpMatcher\]\], whose
    \[\[OriginalSource\]\] internal slot value is
    `serialized`{.variable}.\[\[OriginalSource\]\], and whose
    \[\[OriginalFlags\]\] internal slot value is
    `serialized`{.variable}.\[\[OriginalFlags\]\].

12. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is
    \"SharedArrayBuffer\", then:

    1.  If `targetRealm`{.variable}\'s corresponding [agent
        cluster](https://tc39.es/ecma262/#sec-agent-clusters){#structureddeserialize:agent-cluster
        x-internal="agent-cluster"} is not
        `serialized`{.variable}.\[\[AgentCluster\]\], then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserialize:datacloneerror
        x-internal="datacloneerror"}
        [`DOMException`{#structureddeserialize:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Otherwise, set `value`{.variable} to a new SharedArrayBuffer
        object in `targetRealm`{.variable} whose \[\[ArrayBufferData\]\]
        internal slot value is
        `serialized`{.variable}.\[\[ArrayBufferData\]\] and whose
        \[\[ArrayBufferByteLength\]\] internal slot value is
        `serialized`{.variable}.\[\[ArrayBufferByteLength\]\].

13. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is
    \"GrowableSharedArrayBuffer\", then:

    1.  If `targetRealm`{.variable}\'s corresponding [agent
        cluster](https://tc39.es/ecma262/#sec-agent-clusters){#structureddeserialize:agent-cluster-2
        x-internal="agent-cluster"} is not
        `serialized`{.variable}.\[\[AgentCluster\]\], then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserialize:datacloneerror-2
        x-internal="datacloneerror"}
        [`DOMException`{#structureddeserialize:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Otherwise, set `value`{.variable} to a new SharedArrayBuffer
        object in `targetRealm`{.variable} whose \[\[ArrayBufferData\]\]
        internal slot value is
        `serialized`{.variable}.\[\[ArrayBufferData\]\], whose
        \[\[ArrayBufferByteLengthData\]\] internal slot value is
        `serialized`{.variable}.\[\[ArrayBufferByteLengthData\]\], and
        whose \[\[ArrayBufferMaxByteLength\]\] internal slot value is
        `serialized`{.variable}.\[\[ArrayBufferMaxByteLength\]\].

14. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is
    \"ArrayBuffer\", then set `value`{.variable} to a new ArrayBuffer
    object in `targetRealm`{.variable} whose \[\[ArrayBufferData\]\]
    internal slot value is
    `serialized`{.variable}.\[\[ArrayBufferData\]\], and whose
    \[\[ArrayBufferByteLength\]\] internal slot value is
    `serialized`{.variable}.\[\[ArrayBufferByteLength\]\].

    If this throws an exception, catch it, and then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserialize:datacloneerror-3
    x-internal="datacloneerror"}
    [`DOMException`{#structureddeserialize:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This step might throw an exception if there is not enough memory
    available to create such an ArrayBuffer object.

15. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is
    \"ResizableArrayBuffer\", then set `value`{.variable} to a new
    ArrayBuffer object in `targetRealm`{.variable} whose
    \[\[ArrayBufferData\]\] internal slot value is
    `serialized`{.variable}.\[\[ArrayBufferData\]\], whose
    \[\[ArrayBufferByteLength\]\] internal slot value is
    `serialized`{.variable}.\[\[ArrayBufferByteLength\]\], and whose
    \[\[ArrayBufferMaxByteLength\]\] internal slot value is
    `serialized`{.variable}.\[\[ArrayBufferMaxByteLength\]\].

    If this throws an exception, catch it, and then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserialize:datacloneerror-4
    x-internal="datacloneerror"}
    [`DOMException`{#structureddeserialize:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This step might throw an exception if there is not enough memory
    available to create such an ArrayBuffer object.

16. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is
    \"ArrayBufferView\", then:

    1.  Let `deserializedArrayBuffer`{.variable} be ?
        [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-2}(`serialized`{.variable}.\[\[ArrayBufferSerialized\]\],
        `targetRealm`{.variable}, `memory`{.variable}).

    2.  If `serialized`{.variable}.\[\[Constructor\]\] is \"DataView\",
        then set `value`{.variable} to a new DataView object in
        `targetRealm`{.variable} whose \[\[ViewedArrayBuffer\]\]
        internal slot value is `deserializedArrayBuffer`{.variable},
        whose \[\[ByteLength\]\] internal slot value is
        `serialized`{.variable}.\[\[ByteLength\]\], and whose
        \[\[ByteOffset\]\] internal slot value is
        `serialized`{.variable}.\[\[ByteOffset\]\].

    3.  Otherwise, set `value`{.variable} to a new typed array object in
        `targetRealm`{.variable}, using the constructor given by
        `serialized`{.variable}.\[\[Constructor\]\], whose
        \[\[ViewedArrayBuffer\]\] internal slot value is
        `deserializedArrayBuffer`{.variable}, whose
        \[\[TypedArrayName\]\] internal slot value is
        `serialized`{.variable}.\[\[Constructor\]\], whose
        \[\[ByteLength\]\] internal slot value is
        `serialized`{.variable}.\[\[ByteLength\]\], whose
        \[\[ByteOffset\]\] internal slot value is
        `serialized`{.variable}.\[\[ByteOffset\]\], and whose
        \[\[ArrayLength\]\] internal slot value is
        `serialized`{.variable}.\[\[ArrayLength\]\].

17. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Map\", then:

    1.  Set `value`{.variable} to a new Map object in
        `targetRealm`{.variable} whose \[\[MapData\]\] internal slot
        value is a new empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:js-list
        x-internal="js-list"}.

    2.  Set `deep`{.variable} to true.

18. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Set\", then:

    1.  Set `value`{.variable} to a new Set object in
        `targetRealm`{.variable} whose \[\[SetData\]\] internal slot
        value is a new empty
        [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:js-list-2
        x-internal="js-list"}.

    2.  Set `deep`{.variable} to true.

19. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Array\",
    then:

    1.  Let `outputProto`{.variable} be
        `targetRealm`{.variable}.\[\[Intrinsics\]\].\[\[[%Array.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-array-prototype-object){#structureddeserialize:array.prototype
        x-internal="array.prototype"}\]\].

    2.  Set `value`{.variable} to !
        [ArrayCreate](https://tc39.es/ecma262/#sec-arraycreate){#structureddeserialize:arraycreate
        x-internal="arraycreate"}(`serialized`{.variable}.\[\[Length\]\],
        `outputProto`{.variable}).

    3.  Set `deep`{.variable} to true.

20. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Object\",
    then:

    1.  Set `value`{.variable} to a new Object in
        `targetRealm`{.variable}.

    2.  Set `deep`{.variable} to true.

21. Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Error\",
    then:

    1.  Let `prototype`{.variable} be
        [%Error.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-error-prototype-object){#structureddeserialize:error.prototype
        x-internal="error.prototype"}.

    2.  If `serialized`{.variable}.\[\[Name\]\] is \"EvalError\", then
        set `prototype`{.variable} to
        [%EvalError.prototype%](infrastructure.html#evalerror.prototype){#structureddeserialize:evalerror.prototype}.

    3.  If `serialized`{.variable}.\[\[Name\]\] is \"RangeError\", then
        set `prototype`{.variable} to
        [%RangeError.prototype%](infrastructure.html#rangeerror.prototype){#structureddeserialize:rangeerror.prototype}.

    4.  If `serialized`{.variable}.\[\[Name\]\] is \"ReferenceError\",
        then set `prototype`{.variable} to
        [%ReferenceError.prototype%](infrastructure.html#referenceerror.prototype){#structureddeserialize:referenceerror.prototype}.

    5.  If `serialized`{.variable}.\[\[Name\]\] is \"SyntaxError\", then
        set `prototype`{.variable} to
        [%SyntaxError.prototype%](infrastructure.html#syntaxerror.prototype){#structureddeserialize:syntaxerror.prototype}.

    6.  If `serialized`{.variable}.\[\[Name\]\] is \"TypeError\", then
        set `prototype`{.variable} to
        [%TypeError.prototype%](infrastructure.html#typeerror.prototype){#structureddeserialize:typeerror.prototype}.

    7.  If `serialized`{.variable}.\[\[Name\]\] is \"URIError\", then
        set `prototype`{.variable} to
        [%URIError.prototype%](infrastructure.html#urierror.prototype){#structureddeserialize:urierror.prototype}.

    8.  Let `message`{.variable} be
        `serialized`{.variable}.\[\[Message\]\].

    9.  Set `value`{.variable} to
        [OrdinaryObjectCreate](https://tc39.es/ecma262/#sec-objectcreate){#structureddeserialize:ordinaryobjectcreate
        x-internal="ordinaryobjectcreate"}(`prototype`{.variable}, «
        \[\[ErrorData\]\] »).

    10. Let `messageDesc`{.variable} be
        [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#structureddeserialize:propertydescriptor
        x-internal="propertydescriptor"}{ \[\[Value\]\]:
        `message`{.variable}, \[\[Writable\]\]: true,
        \[\[Enumerable\]\]: false, \[\[Configurable\]\]: true }.

    11. If `message`{.variable} is not undefined, then perform !
        [OrdinaryDefineOwnProperty](https://tc39.es/ecma262/#sec-ordinarydefineownproperty){#structureddeserialize:ordinarydefineownproperty
        x-internal="ordinarydefineownproperty"}(`value`{.variable},
        \"`message`\", `messageDesc`{.variable}).

    12. Any interesting accompanying data attached to
        `serialized`{.variable} should be deserialized and attached to
        `value`{.variable}.

22. Otherwise:

    1.  Let `interfaceName`{.variable} be
        `serialized`{.variable}.\[\[Type\]\].

    2.  If the interface identified by `interfaceName`{.variable} is not
        [exposed](https://webidl.spec.whatwg.org/#dfn-exposed){#structureddeserialize:idl-exposed
        x-internal="idl-exposed"} in `targetRealm`{.variable}, then
        throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserialize:datacloneerror-5
        x-internal="datacloneerror"}
        [`DOMException`{#structureddeserialize:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Set `value`{.variable} to a new instance of the interface
        identified by `interfaceName`{.variable}, created in
        `targetRealm`{.variable}.

    4.  Set `deep`{.variable} to true.

23. [Set](https://infra.spec.whatwg.org/#map-set){#structureddeserialize:map-set
    x-internal="map-set"} `memory`{.variable}\[`serialized`{.variable}\]
    to `value`{.variable}.

24. If `deep`{.variable} is true, then:

    1.  If `serialized`{.variable}.\[\[Type\]\] is \"Map\", then:

        1.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structureddeserialize:list-iterate
            x-internal="list-iterate"}
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:record-2
            x-internal="record"} { \[\[Key\]\], \[\[Value\]\] }
            `entry`{.variable} of
            `serialized`{.variable}.\[\[MapData\]\]:

            1.  Let `deserializedKey`{.variable} be ?
                [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-3}(`entry`{.variable}.\[\[Key\]\],
                `targetRealm`{.variable}, `memory`{.variable}).

            2.  Let `deserializedValue`{.variable} be ?
                [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-4}(`entry`{.variable}.\[\[Value\]\],
                `targetRealm`{.variable}, `memory`{.variable}).

            3.  [Append](https://infra.spec.whatwg.org/#list-append){#structureddeserialize:list-append
                x-internal="list-append"} { \[\[Key\]\]:
                `deserializedKey`{.variable}, \[\[Value\]\]:
                `deserializedValue`{.variable} } to
                `value`{.variable}.\[\[MapData\]\].

    2.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Set\",
        then:

        1.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structureddeserialize:list-iterate-2
            x-internal="list-iterate"} `entry`{.variable} of
            `serialized`{.variable}.\[\[SetData\]\]:

            1.  Let `deserializedEntry`{.variable} be ?
                [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-5}(`entry`{.variable},
                `targetRealm`{.variable}, `memory`{.variable}).

            2.  [Append](https://infra.spec.whatwg.org/#list-append){#structureddeserialize:list-append-2
                x-internal="list-append"} `deserializedEntry`{.variable}
                to `value`{.variable}.\[\[SetData\]\].

    3.  Otherwise, if `serialized`{.variable}.\[\[Type\]\] is \"Array\"
        or \"Object\", then:

        1.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#structureddeserialize:list-iterate-3
            x-internal="list-iterate"}
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:record-3
            x-internal="record"} { \[\[Key\]\], \[\[Value\]\] }
            `entry`{.variable} of
            `serialized`{.variable}.\[\[Properties\]\]:

            1.  Let `deserializedValue`{.variable} be ?
                [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-6}(`entry`{.variable}.\[\[Value\]\],
                `targetRealm`{.variable}, `memory`{.variable}).

            2.  Let `result`{.variable} be !
                [CreateDataProperty](https://tc39.es/ecma262/#sec-createdataproperty){#structureddeserialize:createdataproperty
                x-internal="createdataproperty"}(`value`{.variable},
                `entry`{.variable}.\[\[Key\]\],
                `deserializedValue`{.variable}).

            3.  [Assert](https://infra.spec.whatwg.org/#assert){#structureddeserialize:assert
                x-internal="assert"}: `result`{.variable} is true.

    4.  Otherwise:

        1.  Perform the appropriate [deserialization
            steps](#deserialization-steps){#structureddeserialize:deserialization-steps}
            for the interface identified by
            `serialized`{.variable}.\[\[Type\]\], given
            `serialized`{.variable}, `value`{.variable}, and
            `targetRealm`{.variable}.

            The [deserialization
            steps](#deserialization-steps){#structureddeserialize:deserialization-steps-2}
            may need to perform a
            [sub-deserialization]{#sub-deserialization .dfn export=""}.
            This is an operation which takes as input a
            previously-serialized
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserialize:record-4
            x-internal="record"} `subSerialized`{.variable}, and returns
            [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-7}(`subSerialized`{.variable},
            `targetRealm`{.variable}, `memory`{.variable}). (In other
            words, a
            [sub-deserialization](#sub-deserialization){#structureddeserialize:sub-deserialization}
            is a specialization of
            [StructuredDeserialize](#structureddeserialize){#structureddeserialize:structureddeserialize-8}
            to be consistent within this invocation.)

25. Return `value`{.variable}.

#### [2.7.7]{.secno} [StructuredSerializeWithTransfer]{.dfn} ( `value`{.variable}, `transferList`{.variable} )[](#structuredserializewithtransfer){.self-link} {#structuredserializewithtransfer dfn-type="abstract-op" lt="StructuredSerializeWithTransfer"}

1.  Let `memory`{.variable} be an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#structuredserializewithtransfer:ordered-map
    x-internal="ordered-map"}.

    In addition to how it is used normally by
    [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializewithtransfer:structuredserializeinternal},
    in this algorithm `memory`{.variable} is also used to ensure that
    [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializewithtransfer:structuredserializeinternal-2}
    ignores items in `transferList`{.variable}, and let us do our own
    handling instead.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializewithtransfer:list-iterate
    x-internal="list-iterate"} `transferable`{.variable} of
    `transferList`{.variable}:

    1.  If `transferable`{.variable} has neither an
        \[\[ArrayBufferData\]\] internal slot nor a
        [\[\[Detached\]\]](#detached){#structuredserializewithtransfer:detached}
        internal slot, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializewithtransfer:datacloneerror
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializewithtransfer:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  If `transferable`{.variable} has an \[\[ArrayBufferData\]\]
        internal slot and
        [IsSharedArrayBuffer](https://tc39.es/ecma262/#sec-issharedarraybuffer){#structuredserializewithtransfer:issharedarraybuffer
        x-internal="issharedarraybuffer"}(`transferable`{.variable}) is
        true, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializewithtransfer:datacloneerror-2
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializewithtransfer:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  If `memory`{.variable}\[`transferable`{.variable}\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#structuredserializewithtransfer:map-exists
        x-internal="map-exists"}, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializewithtransfer:datacloneerror-3
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializewithtransfer:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    4.  [Set](https://infra.spec.whatwg.org/#map-set){#structuredserializewithtransfer:map-set
        x-internal="map-set"}
        `memory`{.variable}\[`transferable`{.variable}\] to {
        \[\[Type\]\]: an uninitialized value }.

        `transferable`{.variable} is not transferred yet as transferring
        has side effects and
        [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializewithtransfer:structuredserializeinternal-3}
        needs to be able to throw first.

3.  Let `serialized`{.variable} be ?
    [StructuredSerializeInternal](#structuredserializeinternal){#structuredserializewithtransfer:structuredserializeinternal-4}(`value`{.variable},
    false, `memory`{.variable}).

4.  Let `transferDataHolders`{.variable} be a new empty
    [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structuredserializewithtransfer:js-list
    x-internal="js-list"}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#structuredserializewithtransfer:list-iterate-2
    x-internal="list-iterate"} `transferable`{.variable} of
    `transferList`{.variable}:

    1.  If `transferable`{.variable} has an \[\[ArrayBufferData\]\]
        internal slot and
        [IsDetachedBuffer](https://tc39.es/ecma262/#sec-isdetachedbuffer){#structuredserializewithtransfer:isdetachedbuffer
        x-internal="isdetachedbuffer"}(`transferable`{.variable}) is
        true, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializewithtransfer:datacloneerror-4
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializewithtransfer:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  If `transferable`{.variable} has a
        [\[\[Detached\]\]](#detached){#structuredserializewithtransfer:detached-2}
        internal slot and
        `transferable`{.variable}.[\[\[Detached\]\]](#detached){#structuredserializewithtransfer:detached-3}
        is true, then throw a
        [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structuredserializewithtransfer:datacloneerror-5
        x-internal="datacloneerror"}
        [`DOMException`{#structuredserializewithtransfer:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Let `dataHolder`{.variable} be
        `memory`{.variable}\[`transferable`{.variable}\].

    4.  If `transferable`{.variable} has an \[\[ArrayBufferData\]\]
        internal slot, then:

        1.  If `transferable`{.variable} has an
            \[\[ArrayBufferMaxByteLength\]\] internal slot, then:

            1.  Set `dataHolder`{.variable}.\[\[Type\]\] to
                \"ResizableArrayBuffer\".

            2.  Set `dataHolder`{.variable}.\[\[ArrayBufferData\]\] to
                `transferable`{.variable}.\[\[ArrayBufferData\]\].

            3.  Set
                `dataHolder`{.variable}.\[\[ArrayBufferByteLength\]\] to
                `transferable`{.variable}.\[\[ArrayBufferByteLength\]\].

            4.  Set
                `dataHolder`{.variable}.\[\[ArrayBufferMaxByteLength\]\]
                to
                `transferable`{.variable}.\[\[ArrayBufferMaxByteLength\]\].

        2.  Otherwise:

            1.  Set `dataHolder`{.variable}.\[\[Type\]\] to
                \"ArrayBuffer\".

            2.  Set `dataHolder`{.variable}.\[\[ArrayBufferData\]\] to
                `transferable`{.variable}.\[\[ArrayBufferData\]\].

            3.  Set
                `dataHolder`{.variable}.\[\[ArrayBufferByteLength\]\] to
                `transferable`{.variable}.\[\[ArrayBufferByteLength\]\].

        3.  Perform ?
            [DetachArrayBuffer](https://tc39.es/ecma262/#sec-detacharraybuffer){#structuredserializewithtransfer:detacharraybuffer
            x-internal="detacharraybuffer"}(`transferable`{.variable}).

            Specifications can use the \[\[ArrayBufferDetachKey\]\]
            internal slot to prevent
            [`ArrayBuffer`{#structuredserializewithtransfer:idl-arraybuffer}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}s
            from being detached. This is used in WebAssembly JavaScript
            Interface, for example.
            [\[WASMJS\]](references.html#refsWASMJS)

    5.  Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#structuredserializewithtransfer:assert
            x-internal="assert"}: `transferable`{.variable} is a
            [platform
            object](https://webidl.spec.whatwg.org/#dfn-platform-object){#structuredserializewithtransfer:platform-object
            x-internal="platform-object"} that is a [transferable
            object](#transferable-objects){#structuredserializewithtransfer:transferable-objects}.

        2.  Let `interfaceName`{.variable} be the identifier of the
            [primary
            interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#structuredserializewithtransfer:primary-interface
            x-internal="primary-interface"} of
            `transferable`{.variable}.

        3.  Set `dataHolder`{.variable}.\[\[Type\]\] to
            `interfaceName`{.variable}.

        4.  Perform the appropriate [transfer
            steps](#transfer-steps){#structuredserializewithtransfer:transfer-steps}
            for the interface identified by `interfaceName`{.variable},
            given `transferable`{.variable} and `dataHolder`{.variable}.

        5.  Set
            `transferable`{.variable}.[\[\[Detached\]\]](#detached){#structuredserializewithtransfer:detached-4}
            to true.

    6.  [Append](https://infra.spec.whatwg.org/#list-append){#structuredserializewithtransfer:list-append
        x-internal="list-append"} `dataHolder`{.variable} to
        `transferDataHolders`{.variable}.

6.  Return { \[\[Serialized\]\]: `serialized`{.variable},
    \[\[TransferDataHolders\]\]: `transferDataHolders`{.variable} }.

#### [2.7.8]{.secno} [StructuredDeserializeWithTransfer]{.dfn} ( `serializeWithTransferResult`{.variable}, `targetRealm`{.variable} )[](#structureddeserializewithtransfer){.self-link} {#structureddeserializewithtransfer dfn-type="abstract-op" lt="StructuredDeserializeWithTransfer"}

1.  Let `memory`{.variable} be an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#structureddeserializewithtransfer:ordered-map
    x-internal="ordered-map"}.

    Analogous to
    [StructuredSerializeWithTransfer](#structuredserializewithtransfer){#structureddeserializewithtransfer:structuredserializewithtransfer},
    in addition to how it is used normally by
    [StructuredDeserialize](#structureddeserialize){#structureddeserializewithtransfer:structureddeserialize},
    in this algorithm `memory`{.variable} is also used to ensure that
    [StructuredDeserialize](#structureddeserialize){#structureddeserializewithtransfer:structureddeserialize-2}
    ignores items in
    `serializeWithTransferResult`{.variable}.\[\[TransferDataHolders\]\],
    and let us do our own handling instead.

2.  Let `transferredValues`{.variable} be a new empty
    [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#structureddeserializewithtransfer:js-list
    x-internal="js-list"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#structureddeserializewithtransfer:list-iterate
    x-internal="list-iterate"} `transferDataHolder`{.variable} of
    `serializeWithTransferResult`{.variable}.\[\[TransferDataHolders\]\]:

    1.  Let `value`{.variable} be an uninitialized value.

    2.  If `transferDataHolder`{.variable}.\[\[Type\]\] is
        \"ArrayBuffer\", then set `value`{.variable} to a new
        ArrayBuffer object in `targetRealm`{.variable} whose
        \[\[ArrayBufferData\]\] internal slot value is
        `transferDataHolder`{.variable}.\[\[ArrayBufferData\]\], and
        whose \[\[ArrayBufferByteLength\]\] internal slot value is
        `transferDataHolder`{.variable}.\[\[ArrayBufferByteLength\]\].

        In cases where the original memory occupied by
        \[\[ArrayBufferData\]\] is accessible during the
        deserialization, this step is unlikely to throw an exception, as
        no new memory needs to be allocated: the memory occupied by
        \[\[ArrayBufferData\]\] is instead just getting transferred into
        the new ArrayBuffer. This could be true, for example, when both
        the source and target realms are in the same process.

    3.  Otherwise, if `transferDataHolder`{.variable}.\[\[Type\]\] is
        \"ResizableArrayBuffer\", then set `value`{.variable} to a new
        ArrayBuffer object in `targetRealm`{.variable} whose
        \[\[ArrayBufferData\]\] internal slot value is
        `transferDataHolder`{.variable}.\[\[ArrayBufferData\]\], whose
        \[\[ArrayBufferByteLength\]\] internal slot value is
        `transferDataHolder`{.variable}.\[\[ArrayBufferByteLength\]\],
        and whose \[\[ArrayBufferMaxByteLength\]\] internal slot value
        is
        `transferDataHolder`{.variable}.\[\[ArrayBufferMaxByteLength\]\].

        For the same reason as the previous step, this step is also
        unlikely to throw an exception.

    4.  Otherwise:

        1.  Let `interfaceName`{.variable} be
            `transferDataHolder`{.variable}.\[\[Type\]\].

        2.  If the interface identified by `interfaceName`{.variable} is
            not exposed in `targetRealm`{.variable}, then throw a
            [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structureddeserializewithtransfer:datacloneerror
            x-internal="datacloneerror"}
            [`DOMException`{#structureddeserializewithtransfer:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        3.  Set `value`{.variable} to a new instance of the interface
            identified by `interfaceName`{.variable}, created in
            `targetRealm`{.variable}.

        4.  Perform the appropriate [transfer-receiving
            steps](#transfer-receiving-steps){#structureddeserializewithtransfer:transfer-receiving-steps}
            for the interface identified by `interfaceName`{.variable}
            given `transferDataHolder`{.variable} and
            `value`{.variable}.

    5.  [Set](https://infra.spec.whatwg.org/#map-set){#structureddeserializewithtransfer:map-set
        x-internal="map-set"}
        `memory`{.variable}\[`transferDataHolder`{.variable}\] to
        `value`{.variable}.

    6.  [Append](https://infra.spec.whatwg.org/#list-append){#structureddeserializewithtransfer:list-append
        x-internal="list-append"} `value`{.variable} to
        `transferredValues`{.variable}.

4.  Let `deserialized`{.variable} be ?
    [StructuredDeserialize](#structureddeserialize){#structureddeserializewithtransfer:structureddeserialize-3}(`serializeWithTransferResult`{.variable}.\[\[Serialized\]\],
    `targetRealm`{.variable}, `memory`{.variable}).

5.  Return { \[\[Deserialized\]\]: `deserialized`{.variable},
    \[\[TransferredValues\]\]: `transferredValues`{.variable} }.

#### [2.7.9]{.secno} Performing serialization and transferring from other specifications[](#performing-structured-clones-from-other-specifications){.self-link} {#performing-structured-clones-from-other-specifications}

Other specifications may use the abstract operations defined here. The
following provides some guidance on when each abstract operation is
typically useful, with examples.

[StructuredSerializeWithTransfer](#structuredserializewithtransfer){#performing-structured-clones-from-other-specifications:structuredserializewithtransfer}\
[StructuredDeserializeWithTransfer](#structureddeserializewithtransfer){#performing-structured-clones-from-other-specifications:structureddeserializewithtransfer}

:   Cloning a value to another
    [realm](https://tc39.es/ecma262/#sec-code-realms){#performing-structured-clones-from-other-specifications:realm
    x-internal="realm"}, with a transfer list, but where the target
    realm is not known ahead of time. In this case the serialization
    step can be performed immediately, with the deserialization step
    delayed until the target realm becomes known.

    [`messagePort.postMessage()`{#performing-structured-clones-from-other-specifications:dom-messageport-postmessage}](web-messaging.html#dom-messageport-postmessage)
    uses this pair of abstract operations, as the destination realm is
    not known until the
    [`MessagePort`{#performing-structured-clones-from-other-specifications:messageport}](web-messaging.html#messageport)
    [has been
    shipped](web-messaging.html#has-been-shipped){#performing-structured-clones-from-other-specifications:has-been-shipped}.

[StructuredSerialize](#structuredserialize){#performing-structured-clones-from-other-specifications:structuredserialize}\
[StructuredSerializeForStorage](#structuredserializeforstorage){#performing-structured-clones-from-other-specifications:structuredserializeforstorage}\
[StructuredDeserialize](#structureddeserialize){#performing-structured-clones-from-other-specifications:structureddeserialize}

:   Creating a
    [realm](https://tc39.es/ecma262/#sec-code-realms){#performing-structured-clones-from-other-specifications:realm-2
    x-internal="realm"}-independent snapshot of a given value which can
    be saved for an indefinite amount of time, and then reified back
    into a JavaScript value later, possibly multiple times.

    [StructuredSerializeForStorage](#structuredserializeforstorage){#performing-structured-clones-from-other-specifications:structuredserializeforstorage-2}
    can be used for situations where the serialization is anticipated to
    be stored in a persistent manner, instead of passed between realms.
    It throws when attempting to serialize
    [`SharedArrayBuffer`{#performing-structured-clones-from-other-specifications:sharedarraybuffer}](https://tc39.es/ecma262/#sec-sharedarraybuffer-objects){x-internal="sharedarraybuffer"}
    objects, since storing shared memory does not make sense. Similarly,
    it can throw or possibly have different behavior when given a
    [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#performing-structured-clones-from-other-specifications:platform-object
    x-internal="platform-object"} with custom [serialization
    steps](#serialization-steps){#performing-structured-clones-from-other-specifications:serialization-steps}
    when the `forStorage`{.variable} argument is true.

    [`history.pushState()`{#performing-structured-clones-from-other-specifications:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate)
    and
    [`history.replaceState()`{#performing-structured-clones-from-other-specifications:dom-history-replacestate}](nav-history-apis.html#dom-history-replacestate)
    use
    [StructuredSerializeForStorage](#structuredserializeforstorage){#performing-structured-clones-from-other-specifications:structuredserializeforstorage-3}
    on author-supplied state objects, storing them as [serialized
    state](browsing-the-web.html#serialized-state){#performing-structured-clones-from-other-specifications:serialized-state}
    in the appropriate [session history
    entry](browsing-the-web.html#session-history-entry){#performing-structured-clones-from-other-specifications:session-history-entry}.
    Then,
    [StructuredDeserialize](#structureddeserialize){#performing-structured-clones-from-other-specifications:structureddeserialize-2}
    is used so that the
    [`history.state`{#performing-structured-clones-from-other-specifications:dom-history-state}](nav-history-apis.html#dom-history-state)
    property can return a clone of the originally-supplied state object.

    [`broadcastChannel.postMessage()`{#performing-structured-clones-from-other-specifications:dom-broadcastchannel-postmessage}](web-messaging.html#dom-broadcastchannel-postmessage)
    uses
    [StructuredSerialize](#structuredserialize){#performing-structured-clones-from-other-specifications:structuredserialize-2}
    on its input, then uses
    [StructuredDeserialize](#structureddeserialize){#performing-structured-clones-from-other-specifications:structureddeserialize-3}
    multiple times on the result to produce a fresh clone for each
    destination being broadcast to. Note that transferring does not make
    sense in multi-destination situations.

    Any API for persisting JavaScript values to the filesystem would
    also use
    [StructuredSerializeForStorage](#structuredserializeforstorage){#performing-structured-clones-from-other-specifications:structuredserializeforstorage-4}
    on its input and
    [StructuredDeserialize](#structureddeserialize){#performing-structured-clones-from-other-specifications:structureddeserialize-4}
    on its output.

In general, call sites may pass in Web IDL values instead of JavaScript
values; this is to be understood to perform an implicit
[conversion](https://webidl.spec.whatwg.org/#es-type-mapping){#performing-structured-clones-from-other-specifications:concept-idl-convert
x-internal="concept-idl-convert"} to the JavaScript value before
invoking these algorithms.

------------------------------------------------------------------------

Call sites that are not invoked as a result of author code synchronously
calling into a user agent method must take care to properly [prepare to
run
script](webappapis.html#prepare-to-run-script){#performing-structured-clones-from-other-specifications:prepare-to-run-script}
and [prepare to run a
callback](webappapis.html#prepare-to-run-a-callback){#performing-structured-clones-from-other-specifications:prepare-to-run-a-callback}
before invoking
[StructuredSerialize](#structuredserialize){#performing-structured-clones-from-other-specifications:structuredserialize-3},
[StructuredSerializeForStorage](#structuredserializeforstorage){#performing-structured-clones-from-other-specifications:structuredserializeforstorage-5},
or
[StructuredSerializeWithTransfer](#structuredserializewithtransfer){#performing-structured-clones-from-other-specifications:structuredserializewithtransfer-2}
abstract operations, if they are being performed on arbitrary objects.
This is necessary because the serialization process can invoke
author-defined accessors as part of its final deep-serialization steps,
and these accessors could call into operations that rely on the
[entry](webappapis.html#concept-entry-everything){#performing-structured-clones-from-other-specifications:concept-entry-everything}
and
[incumbent](webappapis.html#concept-incumbent-everything){#performing-structured-clones-from-other-specifications:concept-incumbent-everything}
concepts being properly set up.

[`window.postMessage()`{#performing-structured-clones-from-other-specifications:dom-window-postmessage}](web-messaging.html#dom-window-postmessage)
performs
[StructuredSerializeWithTransfer](#structuredserializewithtransfer){#performing-structured-clones-from-other-specifications:structuredserializewithtransfer-3}
on its arguments, but is careful to do so immediately, inside the
synchronous portion of its algorithm. Thus it is able to use the
algorithms without needing to [prepare to run
script](webappapis.html#prepare-to-run-script){#performing-structured-clones-from-other-specifications:prepare-to-run-script-2}
and [prepare to run a
callback](webappapis.html#prepare-to-run-a-callback){#performing-structured-clones-from-other-specifications:prepare-to-run-a-callback-2}.

In contrast, a hypothetical API that used
[StructuredSerialize](#structuredserialize){#performing-structured-clones-from-other-specifications:structuredserialize-4}
to serialize some author-supplied object periodically, directly from a
[task](webappapis.html#concept-task){#performing-structured-clones-from-other-specifications:concept-task}
on the [event
loop](webappapis.html#event-loop){#performing-structured-clones-from-other-specifications:event-loop},
would need to ensure it performs the appropriate preparations
beforehand. As of this time, we know of no such APIs on the platform;
usually it is simpler to perform the serialization ahead of time, as a
synchronous consequence of author code.

#### [2.7.10]{.secno} Structured cloning API[](#structured-cloning){.self-link} {#structured-cloning}

`result`{.variable}` = self.`[`structuredClone`](#dom-structuredclone){#dom-structuredclone-dev}`(``value`{.variable}`[, { `[`transfer`](web-messaging.html#dom-structuredserializeoptions-transfer){#structured-cloning:dom-structuredserializeoptions-transfer}` }])`

:   Takes the input value and returns a deep copy by performing the
    structured clone algorithm. [Transferable
    objects](#transferable-objects){#structured-cloning:transferable-objects}
    listed in the
    [`transfer`{#structured-cloning:dom-structuredserializeoptions-transfer-2}](web-messaging.html#dom-structuredserializeoptions-transfer)
    array are transferred, not just cloned, meaning that they are no
    longer usable in the input value.

    Throws a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#structured-cloning:datacloneerror
    x-internal="datacloneerror"}
    [`DOMException`{#structured-cloning:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if any part of the input value is not
    [serializable](#serializable-objects){#structured-cloning:serializable-objects}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[structuredClone](https://developer.mozilla.org/en-US/docs/Web/API/structuredClone "The global structuredClone() method creates a deep clone of a given value using the structured clone algorithm.")

Support in all current engines.

::: support
[Firefox94+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome98+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge98+]{.edge_blink .yes}

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
[`structuredClone(``value`{.variable}`, ``options`{.variable}`)`]{#dom-structuredclone
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps
are:

1.  Let `serialized`{.variable} be ?
    [StructuredSerializeWithTransfer](#structuredserializewithtransfer){#structured-cloning:structuredserializewithtransfer}(`value`{.variable},
    `options`{.variable}\[\"[`transfer`{#structured-cloning:dom-structuredserializeoptions-transfer-3}](web-messaging.html#dom-structuredserializeoptions-transfer)\"\]).

2.  Let `deserializeRecord`{.variable} be ?
    [StructuredDeserializeWithTransfer](#structureddeserializewithtransfer){#structured-cloning:structureddeserializewithtransfer}(`serialized`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#structured-cloning:this
    x-internal="this"}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#structured-cloning:concept-relevant-realm}).

3.  Return `deserializeRecord`{.variable}.\[\[Deserialized\]\].

[← 2.6 Common DOM interfaces](common-dom-interfaces.html) --- [Table of
Contents](./) --- [3 Semantics, structure, and APIs of HTML documents
→](dom.html)
