HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 2.4 URLs](urls-and-fetching.html) --- [Table of Contents](./) ---
[2.7 Safe passing of structured data →](structured-data.html)

1.  ::: {#toc-infrastructure}
    1.  [[2.6]{.secno} Common DOM
        interfaces](common-dom-interfaces.html#common-dom-interfaces)
        1.  [[2.6.1]{.secno} Reflecting content attributes in IDL
            attributes](common-dom-interfaces.html#reflecting-content-attributes-in-idl-attributes)
        2.  [[2.6.2]{.secno} Using reflect in
            specifications](common-dom-interfaces.html#using-reflect-in-specifications)
        3.  [[2.6.3]{.secno}
            Collections](common-dom-interfaces.html#collections)
            1.  [[2.6.3.1]{.secno} The `HTMLAllCollection`
                interface](common-dom-interfaces.html#the-htmlallcollection-interface)
                1.  [[2.6.3.1.1]{.secno} \[\[Call\]\] (
                    `thisArgument`{.variable},
                    `argumentsList`{.variable}
                    )](common-dom-interfaces.html#HTMLAllCollection-call)
            2.  [[2.6.3.2]{.secno} The `HTMLFormControlsCollection`
                interface](common-dom-interfaces.html#the-htmlformcontrolscollection-interface)
            3.  [[2.6.3.3]{.secno} The `HTMLOptionsCollection`
                interface](common-dom-interfaces.html#the-htmloptionscollection-interface)
        4.  [[2.6.4]{.secno} The `DOMStringList`
            interface](common-dom-interfaces.html#the-domstringlist-interface)
    :::

### [2.6]{.secno} Common DOM interfaces[](#common-dom-interfaces){.self-link}

#### [2.6.1]{.secno} Reflecting content attributes in IDL attributes[](#reflecting-content-attributes-in-idl-attributes){.self-link}

The building blocks for reflecting are as follows:

- A [reflected target]{#reflected-target .dfn} is an element or
  [`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals}](custom-elements.html#elementinternals)
  object. It is typically clear from context and typically identical to
  the interface of the [reflected IDL
  attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute}.
  It is always identical to that interface when it is an
  [`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals-2}](custom-elements.html#elementinternals)
  object.

- A [reflected IDL attribute]{#reflected-idl-attribute .dfn} is an
  attribute interface member.

- A [reflected content attribute name]{#reflected-content-attribute-name
  .dfn} is a string. When the [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target}
  is an element, it represents the local name of a content attribute
  whose namespace is null. When the [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-2}
  is an
  [`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals-3}](custom-elements.html#elementinternals)
  object, it represents a key of the [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-3}\'s
  [target
  element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target}\'s
  [internal content attribute
  map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map}.

A [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-2}
can be defined to [reflect]{#reflect .dfn export=""} a [reflected
content attribute
name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name}
of a [reflected
target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-4}.
In general this means that the IDL attribute getter returns the current
value of the content attribute, and the setter changes the value of the
content attribute to the given value.

If the [reflected
target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-5}
is an element, then the [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-3}
can additionally declare to [support
`ElementInternals`]{#support-elementinternals .dfn}. This means that the
[`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals-4}](custom-elements.html#elementinternals)
interface also has a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-4},
with the same identifier, and that [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-5}
[reflects](#reflect){#reflecting-content-attributes-in-idl-attributes:reflect}
the same [reflected content attribute
name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-2}.

The `fooBar`{.variable} IDL attribute must
[reflect](#reflect){#reflecting-content-attributes-in-idl-attributes:reflect-2}
the `foobar`{.variable} content attribute and [support
`ElementInternals`](#support-elementinternals){#reflecting-content-attributes-in-idl-attributes:support-elementinternals}.

[Reflected
targets](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-6}
have these associated algorithms:

- [get the element]{#get-the-element .dfn}: takes no arguments; returns
  an element.
- [get the content attribute]{#get-the-content-attribute .dfn}: takes no
  arguments; returns null or a string.
- [set the content attribute]{#set-the-content-attribute .dfn}: takes a
  string `value`{.variable}; returns nothing.
- [delete the content attribute]{#delete-the-content-attribute .dfn}:
  takes no arguments; returns nothing.

For a [reflected
target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-7}
that is an element `element`{.variable}, these are defined as follows:

[get the element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element}

:   1.  Return `element`{.variable}.

[get the content attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute}

:   1.  Let `attribute`{.variable} be the result of running [get an
        attribute by namespace and local
        name](https://dom.spec.whatwg.org/#concept-element-attributes-get-by-namespace){#reflecting-content-attributes-in-idl-attributes:concept-element-attributes-get-by-namespace
        x-internal="concept-element-attributes-get-by-namespace"} given
        null, the [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-3},
        and `element`{.variable}.

    2.  If `attribute`{.variable} is null, then return null.

    3.  Return `attribute`{.variable}\'s
        [value](https://dom.spec.whatwg.org/#concept-attribute-value){#reflecting-content-attributes-in-idl-attributes:concept-attribute-value
        x-internal="concept-attribute-value"}.

[set the content attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute} with a string `value`{.variable}

:   1.  [Set an attribute
        value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#reflecting-content-attributes-in-idl-attributes:concept-element-attributes-set-value
        x-internal="concept-element-attributes-set-value"} given
        `element`{.variable}, the [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-4},
        and `value`{.variable}.

[delete the content attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute}

:   1.  [Remove an attribute by namespace and local
        name](https://dom.spec.whatwg.org/#concept-element-attributes-remove-by-namespace){#reflecting-content-attributes-in-idl-attributes:concept-element-attributes-remove-by-namespace
        x-internal="concept-element-attributes-remove-by-namespace"}
        given null, the [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-5},
        and `element`{.variable}.

For a [reflected
target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-8}
that is an
[`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals-5}](custom-elements.html#elementinternals)
object `elementInternals`{.variable}, they are defined as follows:

[get the element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-2}

:   1.  Return `elementInternals`{.variable}\'s [target
        element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-2}.

[get the content attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-2}

:   1.  If `elementInternals`{.variable}\'s [target
        element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-3}\'s
        [internal content attribute
        map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map-2}\[the
        [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-6}\]
        [does not
        exist](https://infra.spec.whatwg.org/#map-exists){#reflecting-content-attributes-in-idl-attributes:map-exists
        x-internal="map-exists"}, then return null.

    2.  Return `elementInternals`{.variable}\'s [target
        element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-4}\'s
        [internal content attribute
        map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map-3}\[the
        [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-7}\].

[set the content attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-2} with a string `value`{.variable}

:   1.  [Set](https://infra.spec.whatwg.org/#map-set){#reflecting-content-attributes-in-idl-attributes:map-set
        x-internal="map-set"} `elementInternals`{.variable}\'s [target
        element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-5}\'s
        [internal content attribute
        map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map-4}\[the
        [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-8}\]
        to `value`{.variable}.

[delete the content attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute-2}

:   1.  [Remove](https://infra.spec.whatwg.org/#map-remove){#reflecting-content-attributes-in-idl-attributes:map-remove
        x-internal="map-remove"} `elementInternals`{.variable}\'s
        [target
        element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-6}\'s
        [internal content attribute
        map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map-5}\[the
        [reflected content attribute
        name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-9}\].

This results in somewhat redundant data structures for
[`ElementInternals`{#reflecting-content-attributes-in-idl-attributes:elementinternals-6}](custom-elements.html#elementinternals)
objects as their [target
element](custom-elements.html#internals-target){#reflecting-content-attributes-in-idl-attributes:internals-target-7}\'s
[internal content attribute
map](custom-elements.html#internal-content-attribute-map){#reflecting-content-attributes-in-idl-attributes:internal-content-attribute-map-6}
cannot be directly manipulated and as such reflection is only happening
in a single direction. This approach was nevertheless chosen to make it
less error-prone to define IDL attributes that are shared between
[reflected
targets](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-9}
and benefit from common API semantics.

------------------------------------------------------------------------

IDL attributes of type
[`DOMString`{#reflecting-content-attributes-in-idl-attributes:idl-domstring}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"}
or
[`DOMString`](https://webidl.spec.whatwg.org/#idl-DOMString){#reflecting-content-attributes-in-idl-attributes:idl-domstring-2
x-internal="idl-domstring"}`?` that
[reflect](#reflect){#reflecting-content-attributes-in-idl-attributes:reflect-3}
[enumerated](common-microsyntaxes.html#enumerated-attribute){#reflecting-content-attributes-in-idl-attributes:enumerated-attribute}
content attributes can be [limited to only known
values]{#limited-to-only-known-values .dfn}. Per the processing models
below, those will cause the getters for such IDL attributes to only
return keywords for those enumerated attributes, or the empty string or
null.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-6}
has the type
[`DOMString`{#reflecting-content-attributes-in-idl-attributes:idl-domstring-3}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"}:

- The getter steps are:

  1.  Let `element`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this
      x-internal="this"}\'s [get the
      element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-3}.

  2.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-2
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-3}.

  3.  Let `attributeDefinition`{.variable} be the attribute definition
      of `element`{.variable}\'s content attribute whose namespace is
      null and local name is the [reflected content attribute
      name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-10}.

  4.  If `attributeDefinition`{.variable} indicates it is an [enumerated
      attribute](common-microsyntaxes.html#enumerated-attribute){#reflecting-content-attributes-in-idl-attributes:enumerated-attribute-2}
      and the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-7}
      is defined to be [limited to only known
      values](#limited-to-only-known-values){#reflecting-content-attributes-in-idl-attributes:limited-to-only-known-values}:

      1.  If `contentAttributeValue`{.variable} does not correspond to
          any state of `attributeDefinition`{.variable} (e.g., it is
          null and there is no *[missing value
          default](common-microsyntaxes.html#missing-value-default)*),
          or if it is in a state of `attributeDefinition`{.variable}
          with no associated keyword value, then return the empty
          string.

      2.  Return the [canonical
          keyword](common-microsyntaxes.html#canonical-keyword){#reflecting-content-attributes-in-idl-attributes:canonical-keyword}
          for the state of `attributeDefinition`{.variable} that
          `contentAttributeValue`{.variable} corresponds to.

  5.  If `contentAttributeValue`{.variable} is null, then return the
      empty string.

  6.  Return `contentAttributeValue`{.variable}.

- The setter steps are to run
  [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-3
  x-internal="this"}\'s [set the content
  attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-3}
  with the given value.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-8}
has the type
[`DOMString`](https://webidl.spec.whatwg.org/#idl-DOMString){#reflecting-content-attributes-in-idl-attributes:idl-domstring-4
x-internal="idl-domstring"}`?`:

- The getter steps are:

  1.  Let `element`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-4
      x-internal="this"}\'s [get the
      element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-4}.

  2.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-5
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-4}.

  3.  Let `attributeDefinition`{.variable} be the attribute definition
      of `element`{.variable}\'s content attribute whose namespace is
      null and local name is the [reflected content attribute
      name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-11}.

  4.  If `attributeDefinition`{.variable} indicates it is an [enumerated
      attribute](common-microsyntaxes.html#enumerated-attribute){#reflecting-content-attributes-in-idl-attributes:enumerated-attribute-3}:

      1.  [Assert](https://infra.spec.whatwg.org/#assert){#reflecting-content-attributes-in-idl-attributes:assert
          x-internal="assert"}: the [reflected IDL
          attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-9}
          is [limited to only known
          values](#limited-to-only-known-values){#reflecting-content-attributes-in-idl-attributes:limited-to-only-known-values-2}.

      2.  [Assert](https://infra.spec.whatwg.org/#assert){#reflecting-content-attributes-in-idl-attributes:assert-2
          x-internal="assert"}: `contentAttributeValue`{.variable}
          corresponds to a state of `attributeDefinition`{.variable}.

      3.  If `contentAttributeValue`{.variable} corresponds to a state
          of `attributeDefinition`{.variable} with no associated keyword
          value, then return null.

      4.  Return the [canonical
          keyword](common-microsyntaxes.html#canonical-keyword){#reflecting-content-attributes-in-idl-attributes:canonical-keyword-2}
          for the state of `attributeDefinition`{.variable} that
          `contentAttributeValue`{.variable} corresponds to.

  5.  Return `contentAttributeValue`{.variable}.

- The setter steps are:

  1.  If the given value is null, then run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-6
      x-internal="this"}\'s [delete the content
      attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute-3}.

  2.  Otherwise, run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-7
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-4}
      with the given value.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-10}
has the type
[`USVString`{#reflecting-content-attributes-in-idl-attributes:idl-usvstring}](https://webidl.spec.whatwg.org/#idl-USVString){x-internal="idl-usvstring"}:

- The getter steps are:

  1.  Let `element`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-8
      x-internal="this"}\'s [get the
      element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-5}.

  2.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-9
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-5}.

  3.  Let `attributeDefinition`{.variable} be the attribute definition
      of `element`{.variable}\'s content attribute whose namespace is
      null and local name is the [reflected content attribute
      name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-12}.

  4.  If `attributeDefinition`{.variable} indicates it contains a
      [URL](https://url.spec.whatwg.org/#concept-url){#reflecting-content-attributes-in-idl-attributes:url
      x-internal="url"}:

      1.  If `contentAttributeValue`{.variable} is null, then return the
          empty string.

      2.  Let `urlString`{.variable} be the result of
          [encoding-parsing-and-serializing a
          URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#reflecting-content-attributes-in-idl-attributes:encoding-parsing-and-serializing-a-url}
          given `contentAttributeValue`{.variable}, relative to
          `element`{.variable}\'s [node
          document](https://dom.spec.whatwg.org/#concept-node-document){#reflecting-content-attributes-in-idl-attributes:node-document
          x-internal="node-document"}.

      3.  If `urlString`{.variable} is not failure, then return
          `urlString`{.variable}.

  5.  Return `contentAttributeValue`{.variable}, [converted to a scalar
      value
      string](https://infra.spec.whatwg.org/#javascript-string-convert){#reflecting-content-attributes-in-idl-attributes:convert
      x-internal="convert"}.

- The setter steps are to run
  [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-10
  x-internal="this"}\'s [set the content
  attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-5}
  with the given value.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-11}
has the type
[`boolean`{#reflecting-content-attributes-in-idl-attributes:idl-boolean}](https://webidl.spec.whatwg.org/#idl-boolean){x-internal="idl-boolean"}:

- The getter steps are:

  1.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-11
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-6}.

  2.  If `contentAttributeValue`{.variable} is null, then return false.

  3.  Return true.

- The setter steps are:

  1.  If the given value is false, then run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-12
      x-internal="this"}\'s [delete the content
      attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute-4}.

  2.  If the given value is true, then run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-13
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-6}
      with the empty string.

This corresponds to the rules for [boolean content
attributes](common-microsyntaxes.html#boolean-attribute){#reflecting-content-attributes-in-idl-attributes:boolean-attribute}.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-12}
has the type
[`long`{#reflecting-content-attributes-in-idl-attributes:idl-long}](https://webidl.spec.whatwg.org/#idl-long){x-internal="idl-long"},
optionally [limited to only non-negative
numbers]{#limited-to-only-non-negative-numbers .dfn} and optionally with
a [default value]{#default-value .dfn} `defaultValue`{.variable}:

- The getter steps are:

  1.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-14
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-7}.

  2.  If `contentAttributeValue`{.variable} is not null:

      1.  Let `parsedValue`{.variable} be the result of [integer
          parsing](common-microsyntaxes.html#rules-for-parsing-integers){#reflecting-content-attributes-in-idl-attributes:rules-for-parsing-integers}
          `contentAttributeValue`{.variable} if the [reflected IDL
          attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-13}
          is not [limited to only non-negative
          numbers](#limited-to-only-non-negative-numbers){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers};
          otherwise the result of [non-negative integer
          parsing](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#reflecting-content-attributes-in-idl-attributes:rules-for-parsing-non-negative-integers}
          `contentAttributeValue`{.variable}.

      2.  If `parsedValue`{.variable} is not an error and is within the
          [`long`{#reflecting-content-attributes-in-idl-attributes:idl-long-2}](https://webidl.spec.whatwg.org/#idl-long){x-internal="idl-long"}
          range, then return `parsedValue`{.variable}.

  3.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-14}
      has a [default
      value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value},
      then return `defaultValue`{.variable}.

  4.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-15}
      is [limited to only non-negative
      numbers](#limited-to-only-non-negative-numbers){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-2},
      then return −1.

  5.  Return 0.

- The setter steps are:

  1.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-16}
      is [limited to only non-negative
      numbers](#limited-to-only-non-negative-numbers){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-3}
      and the given value is negative, then throw an
      [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#reflecting-content-attributes-in-idl-attributes:indexsizeerror
      x-internal="indexsizeerror"}
      [`DOMException`{#reflecting-content-attributes-in-idl-attributes:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

  2.  Run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-15
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-7}
      with the given value converted to the shortest possible string
      representing the number as a [valid
      integer](common-microsyntaxes.html#valid-integer){#reflecting-content-attributes-in-idl-attributes:valid-integer}.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-17}
has the type
[`unsigned long`{#reflecting-content-attributes-in-idl-attributes:idl-unsigned-long}](https://webidl.spec.whatwg.org/#idl-unsigned-long){x-internal="idl-unsigned-long"},
optionally [limited to only positive
numbers]{#limited-to-only-non-negative-numbers-greater-than-zero .dfn},
[limited to only positive numbers with
fallback]{#limited-to-only-non-negative-numbers-greater-than-zero-with-fallback
.dfn}, or [clamped to the range]{#clamped-to-the-range .dfn}
\[`clampedMin`{.variable}, `clampedMax`{.variable}\], and optionally
with a [default
value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value-2}
`defaultValue`{.variable}:

- The getter steps are:

  1.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-16
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-8}.

  2.  Let `minimum`{.variable} be 0.

  3.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-18}
      is [limited to only positive
      numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero}
      or [limited to only positive numbers with
      fallback](#limited-to-only-non-negative-numbers-greater-than-zero-with-fallback){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-with-fallback},
      then set `minimum`{.variable} to 1.

  4.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-19}
      is [clamped to the
      range](#clamped-to-the-range){#reflecting-content-attributes-in-idl-attributes:clamped-to-the-range},
      then set `minimum`{.variable} to `clampedMin`{.variable}.

  5.  Let `maximum`{.variable} be 2147483647 if the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-20}
      is not [clamped to the
      range](#clamped-to-the-range){#reflecting-content-attributes-in-idl-attributes:clamped-to-the-range-2};
      otherwise `clampedMax`{.variable}.

  6.  If `contentAttributeValue`{.variable} is not null:

      1.  Let `parsedValue`{.variable} be the result of [non-negative
          integer
          parsing](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#reflecting-content-attributes-in-idl-attributes:rules-for-parsing-non-negative-integers-2}
          `contentAttributeValue`{.variable}.

      2.  If `parsedValue`{.variable} is not an error and is in the
          range `minimum`{.variable} to `maximum`{.variable}, inclusive,
          then return `parsedValue`{.variable}.

      3.  If `parsedValue`{.variable} is not an error and the [reflected
          IDL
          attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-21}
          is [clamped to the
          range](#clamped-to-the-range){#reflecting-content-attributes-in-idl-attributes:clamped-to-the-range-3}:

          1.  If `parsedValue`{.variable} is less than
              `minimum`{.variable}, then return `minimum`{.variable}.

          2.  Return `maximum`{.variable}.

  7.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-22}
      has a [default
      value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value-3},
      then return `defaultValue`{.variable}.

  8.  Return `minimum`{.variable}.

- The setter steps are:

  1.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-23}
      is [limited to only positive
      numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-2}
      and the given value is 0, then throw an
      [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#reflecting-content-attributes-in-idl-attributes:indexsizeerror-2
      x-internal="indexsizeerror"}
      [`DOMException`{#reflecting-content-attributes-in-idl-attributes:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

  2.  Let `minimum`{.variable} be 0.

  3.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-24}
      is [limited to only positive
      numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-3}
      or [limited to only positive numbers with
      fallback](#limited-to-only-non-negative-numbers-greater-than-zero-with-fallback){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-with-fallback-2},
      then set `minimum`{.variable} to 1.

  4.  Let `newValue`{.variable} be `minimum`{.variable}.

  5.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-25}
      has a [default
      value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value-4},
      then set `newValue`{.variable} to `defaultValue`{.variable}.

  6.  If the given value is in the range `minimum`{.variable} to
      2147483647, inclusive, then set `newValue`{.variable} to it.

  7.  Run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-17
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-8}
      with `newValue`{.variable} converted to the shortest possible
      string representing the number as a [valid non-negative
      integer](common-microsyntaxes.html#valid-non-negative-integer){#reflecting-content-attributes-in-idl-attributes:valid-non-negative-integer}.

  [Clamped to the
  range](#clamped-to-the-range){#reflecting-content-attributes-in-idl-attributes:clamped-to-the-range-4}
  has no effect on the setter steps.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-26}
has the type
[`double`{#reflecting-content-attributes-in-idl-attributes:idl-double}](https://webidl.spec.whatwg.org/#idl-double){x-internal="idl-double"},
optionally [limited to only positive
numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#limited-to-numbers-greater-than-zero}
and optionally with a [default
value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value-5}
`defaultValue`{.variable}:

- The getter steps are:

  1.  Let `contentAttributeValue`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-18
      x-internal="this"}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-9}.

  2.  If `contentAttributeValue`{.variable} is not null:

      1.  Let `parsedValue`{.variable} be the result of [floating-point
          number
          parsing](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#reflecting-content-attributes-in-idl-attributes:rules-for-parsing-floating-point-number-values}
          `contentAttributeValue`{.variable}.

      2.  If `parsedValue`{.variable} is not an error and is greater
          than 0, then return `parsedValue`{.variable}.

      3.  If `parsedValue`{.variable} is not an error and the [reflected
          IDL
          attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-27}
          is not [limited to only positive
          numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-4},
          then return `parsedValue`{.variable}.

  3.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-28}
      has a [default
      value](#default-value){#reflecting-content-attributes-in-idl-attributes:default-value-6},
      then return `defaultValue`{.variable}.

  4.  Return 0.

- The setter steps are:

  1.  If the [reflected IDL
      attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-29}
      is [limited to only positive
      numbers](#limited-to-only-non-negative-numbers-greater-than-zero){#reflecting-content-attributes-in-idl-attributes:limited-to-only-non-negative-numbers-greater-than-zero-5}
      and the given value is not greater than 0, then return.

  2.  Run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-19
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-9}
      with the given value, converted to the [best representation of the
      number as a floating-point
      number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#reflecting-content-attributes-in-idl-attributes:best-representation-of-the-number-as-a-floating-point-number}.

The values Infinity and Not-a-Number (NaN) values throw an exception on
setting, as defined in Web IDL. [\[WEBIDL\]](references.html#refsWEBIDL)

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-30}
has the type
[`DOMTokenList`{#reflecting-content-attributes-in-idl-attributes:domtokenlist}](https://dom.spec.whatwg.org/#interface-domtokenlist){x-internal="domtokenlist"},
then its getter steps are to return a
[`DOMTokenList`{#reflecting-content-attributes-in-idl-attributes:domtokenlist-2}](https://dom.spec.whatwg.org/#interface-domtokenlist){x-internal="domtokenlist"}
object whose associated element is
[this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-20
x-internal="this"} and associated attribute\'s local name is the
[reflected content attribute
name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-13}.
Specification authors cannot use [support
`ElementInternals`](#support-elementinternals){#reflecting-content-attributes-in-idl-attributes:support-elementinternals-2}
for IDL attributes of this type.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-31}
has the type `T`{.variable}`?`, where `T`{.variable} is either
[`Element`{#reflecting-content-attributes-in-idl-attributes:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or an interface that inherits from
[`Element`{#reflecting-content-attributes-in-idl-attributes:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
then with `attr`{.variable} being the [reflected content attribute
name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-14}:

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-10}
  has an [explicitly set
  `attr`{.variable}-element]{#explicitly-set-attr-element .dfn}, which
  is a weak reference to an element or null. It is initially null.

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-11}
  `reflectedTarget`{.variable} has a [get the
  `attr`{.variable}-associated element]{#attr-associated-element .dfn
  dfn-for="Element,ElementInternals" export=""} algorithm, that runs
  these steps:

  1.  Let `element`{.variable} be the result of running
      `reflectedTarget`{.variable}\'s [get the
      element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-6}.

  2.  Let `contentAttributeValue`{.variable} be the result of running
      `reflectedTarget`{.variable}\'s [get the content
      attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-10}.

  3.  If `reflectedTarget`{.variable}\'s [explicitly set
      `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element}
      is not null:

      1.  If `reflectedTarget`{.variable}\'s [explicitly set
          `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element-2}
          is a
          [descendant](https://dom.spec.whatwg.org/#concept-tree-descendant){#reflecting-content-attributes-in-idl-attributes:descendant
          x-internal="descendant"} of any of `element`{.variable}\'s
          [shadow-including
          ancestors](https://dom.spec.whatwg.org/#concept-shadow-including-ancestor){#reflecting-content-attributes-in-idl-attributes:concept-shadow-including-ancestor
          x-internal="concept-shadow-including-ancestor"}, then return
          `reflectedTarget`{.variable}\'s [explicitly set
          `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element-3}.

      2.  Return null.

  4.  Otherwise, if `contentAttributeValue`{.variable} is not null,
      return the first element `candidate`{.variable}, in [tree
      order](https://dom.spec.whatwg.org/#concept-tree-order){#reflecting-content-attributes-in-idl-attributes:tree-order
      x-internal="tree-order"}, that meets the following criteria:

      - `candidate`{.variable}\'s
        [root](https://dom.spec.whatwg.org/#concept-tree-root){#reflecting-content-attributes-in-idl-attributes:root
        x-internal="root"} is the same as `element`{.variable}\'s
        [root](https://dom.spec.whatwg.org/#concept-tree-root){#reflecting-content-attributes-in-idl-attributes:root-2
        x-internal="root"};

      - `candidate`{.variable}\'s
        [ID](https://dom.spec.whatwg.org/#concept-id){#reflecting-content-attributes-in-idl-attributes:concept-id
        x-internal="concept-id"} is `contentAttributeValue`{.variable};
        and

      - `candidate`{.variable}
        [implements](https://webidl.spec.whatwg.org/#implements){#reflecting-content-attributes-in-idl-attributes:implements
        x-internal="implements"} `T`{.variable}.

      If no such element exists, then return null.

  5.  Return null.

- The getter steps are to return the result of running
  [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-21
  x-internal="this"}\'s [get the `attr`{.variable}-associated
  element](#attr-associated-element){#reflecting-content-attributes-in-idl-attributes:attr-associated-element}.

- The setter steps are:

  1.  If the given value is null, then:

      1.  Set
          [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-22
          x-internal="this"}\'s [explicitly set
          `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element-4}
          to null.

      2.  Run
          [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-23
          x-internal="this"}\'s [delete the content
          attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute-5}.

      3.  Return.

  2.  Run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-24
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-10}
      with the empty string.

  3.  Set
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-25
      x-internal="this"}\'s [explicitly set
      `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element-5}
      to a weak reference to the given value.

- For element [reflected
  targets](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-12}
  only: the following [attribute change
  steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#reflecting-content-attributes-in-idl-attributes:concept-element-attributes-change-ext
  x-internal="concept-element-attributes-change-ext"}, given
  `element`{.variable}, `localName`{.variable}, `oldValue`{.variable},
  `value`{.variable}, and `namespace`{.variable}, are used to
  synchronize between the content attribute and the IDL attribute:

  1.  If `localName`{.variable} is not `attr`{.variable} or
      `namespace`{.variable} is not null, then return.

  2.  Set `element`{.variable}\'s [explicitly set
      `attr`{.variable}-element](#explicitly-set-attr-element){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-element-6}
      to null.

[Reflected IDL
attributes](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-32}
of this type are strongly encouraged to have their identifier end in
\"`Element`\" for consistency.

If a [reflected IDL
attribute](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-33}
has the type `FrozenArray<``T`{.variable}`>?`, where `T`{.variable} is
either
[`Element`{#reflecting-content-attributes-in-idl-attributes:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or an interface that inherits from
[`Element`{#reflecting-content-attributes-in-idl-attributes:element-4}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
then with `attr`{.variable} being the [reflected content attribute
name](#reflected-content-attribute-name){#reflecting-content-attributes-in-idl-attributes:reflected-content-attribute-name-15}:

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-13}
  has an [explicitly set
  `attr`{.variable}-elements]{#explicitly-set-attr-elements .dfn}, which
  is either a
  [list](https://infra.spec.whatwg.org/#list){#reflecting-content-attributes-in-idl-attributes:list
  x-internal="list"} of weak references to elements or null. It is
  initially null.

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-14}
  has a [cached `attr`{.variable}-associated
  elements]{#cached-attr-associated-elements .dfn}, which is a
  [list](https://infra.spec.whatwg.org/#list){#reflecting-content-attributes-in-idl-attributes:list-2
  x-internal="list"} of elements. It is initially « ».

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-15}
  has a [cached `attr`{.variable}-associated elements
  object]{#cached-attr-associated-elements-object .dfn}, which is a
  `FrozenArray<``T`{.variable}`>?`. It is initially null.

- Its [reflected
  target](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-16}
  `reflectedTarget`{.variable} has a [get the
  `attr`{.variable}-associated elements]{#attr-associated-elements .dfn
  dfn-for="Element,ElementInternals" export=""} algorithm, which runs
  these steps:

  1.  Let `elements`{.variable} be an empty
      [list](https://infra.spec.whatwg.org/#list){#reflecting-content-attributes-in-idl-attributes:list-3
      x-internal="list"}.

  2.  Let `element`{.variable} be the result of running
      `reflectedTarget`{.variable}\'s [get the
      element](#get-the-element){#reflecting-content-attributes-in-idl-attributes:get-the-element-7}.

  3.  If `reflectedTarget`{.variable}\'s [explicitly set
      `attr`{.variable}-elements](#explicitly-set-attr-elements){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-elements}
      is not null:

      1.  [For
          each](https://infra.spec.whatwg.org/#list-iterate){#reflecting-content-attributes-in-idl-attributes:list-iterate
          x-internal="list-iterate"} `attrElement`{.variable} in
          `reflectedTarget`{.variable}\'s [explicitly set
          `attr`{.variable}-elements](#explicitly-set-attr-elements){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-elements-2}:

          1.  If `attrElement`{.variable} is not a
              [descendant](https://dom.spec.whatwg.org/#concept-tree-descendant){#reflecting-content-attributes-in-idl-attributes:descendant-2
              x-internal="descendant"} of any of `element`{.variable}\'s
              [shadow-including
              ancestors](https://dom.spec.whatwg.org/#concept-shadow-including-ancestor){#reflecting-content-attributes-in-idl-attributes:concept-shadow-including-ancestor-2
              x-internal="concept-shadow-including-ancestor"}, then
              [continue](https://infra.spec.whatwg.org/#iteration-continue){#reflecting-content-attributes-in-idl-attributes:continue
              x-internal="continue"}.

          2.  [Append](https://infra.spec.whatwg.org/#list-append){#reflecting-content-attributes-in-idl-attributes:list-append
              x-internal="list-append"} `attrElement`{.variable} to
              `elements`{.variable}.

  4.  Otherwise:

      1.  Let `contentAttributeValue`{.variable} be the result of
          running `reflectedTarget`{.variable}\'s [get the content
          attribute](#get-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:get-the-content-attribute-11}.

      2.  If `contentAttributeValue`{.variable} is null, then return
          null.

      3.  Let `tokens`{.variable} be `contentAttributeValue`{.variable},
          [split on ASCII
          whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#reflecting-content-attributes-in-idl-attributes:split-a-string-on-spaces
          x-internal="split-a-string-on-spaces"}.

      4.  [For
          each](https://infra.spec.whatwg.org/#list-iterate){#reflecting-content-attributes-in-idl-attributes:list-iterate-2
          x-internal="list-iterate"} `id`{.variable} of
          `tokens`{.variable}:

          1.  Let `candidate`{.variable} be the first element, in [tree
              order](https://dom.spec.whatwg.org/#concept-tree-order){#reflecting-content-attributes-in-idl-attributes:tree-order-2
              x-internal="tree-order"}, that meets the following
              criteria:

              - `candidate`{.variable}\'s
                [root](https://dom.spec.whatwg.org/#concept-tree-root){#reflecting-content-attributes-in-idl-attributes:root-3
                x-internal="root"} is the same as
                `element`{.variable}\'s
                [root](https://dom.spec.whatwg.org/#concept-tree-root){#reflecting-content-attributes-in-idl-attributes:root-4
                x-internal="root"};

              - `candidate`{.variable}\'s
                [ID](https://dom.spec.whatwg.org/#concept-id){#reflecting-content-attributes-in-idl-attributes:concept-id-2
                x-internal="concept-id"} is `id`{.variable}; and

              - `candidate`{.variable}
                [implements](https://webidl.spec.whatwg.org/#implements){#reflecting-content-attributes-in-idl-attributes:implements-2
                x-internal="implements"} `T`{.variable}.

              If no such element exists, then
              [continue](https://infra.spec.whatwg.org/#iteration-continue){#reflecting-content-attributes-in-idl-attributes:continue-2
              x-internal="continue"}.

          2.  [Append](https://infra.spec.whatwg.org/#list-append){#reflecting-content-attributes-in-idl-attributes:list-append-2
              x-internal="list-append"} `candidate`{.variable} to
              `elements`{.variable}.

  5.  Return `elements`{.variable}.

- The getter steps are:

  1.  Let `elements`{.variable} be the result of running
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-26
      x-internal="this"}\'s [get the `attr`{.variable}-associated
      elements](#attr-associated-elements){#reflecting-content-attributes-in-idl-attributes:attr-associated-elements}.

  2.  If the contents of `elements`{.variable} is equal to the contents
      of
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-27
      x-internal="this"}\'s [cached `attr`{.variable}-associated
      elements](#cached-attr-associated-elements){#reflecting-content-attributes-in-idl-attributes:cached-attr-associated-elements},
      then return
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-28
      x-internal="this"}\'s [cached `attr`{.variable}-associated
      elements
      object](#cached-attr-associated-elements-object){#reflecting-content-attributes-in-idl-attributes:cached-attr-associated-elements-object}.

  3.  Let `elementsAsFrozenArray`{.variable} be `elements`{.variable},
      [converted](https://webidl.spec.whatwg.org/#es-type-mapping){#reflecting-content-attributes-in-idl-attributes:concept-idl-convert
      x-internal="concept-idl-convert"} to a
      `FrozenArray<``T`{.variable}`>?`.

  4.  Set
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-29
      x-internal="this"}\'s [cached `attr`{.variable}-associated
      elements](#cached-attr-associated-elements){#reflecting-content-attributes-in-idl-attributes:cached-attr-associated-elements-2}
      to `elements`{.variable}.

  5.  Set
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-30
      x-internal="this"}\'s [cached `attr`{.variable}-associated
      elements
      object](#cached-attr-associated-elements-object){#reflecting-content-attributes-in-idl-attributes:cached-attr-associated-elements-object-2}
      to `elementsAsFrozenArray`{.variable}.

  6.  Return `elementsAsFrozenArray`{.variable}.

  This extra caching layer is necessary to preserve the invariant that
  `element.reflectedElements === element.reflectedElements`.

- The setter steps are:

  1.  If the given value is null:

      1.  Set
          [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-31
          x-internal="this"}\'s [explicitly set
          `attr`{.variable}-elements](#explicitly-set-attr-elements){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-elements-3}
          to null.

      2.  Run
          [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-32
          x-internal="this"}\'s [delete the content
          attribute](#delete-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:delete-the-content-attribute-6}.

      3.  Return.

  2.  Run
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-33
      x-internal="this"}\'s [set the content
      attribute](#set-the-content-attribute){#reflecting-content-attributes-in-idl-attributes:set-the-content-attribute-11}
      with the empty string.

  3.  Let `elements`{.variable} be an empty
      [list](https://infra.spec.whatwg.org/#list){#reflecting-content-attributes-in-idl-attributes:list-4
      x-internal="list"}.

  4.  [For
      each](https://infra.spec.whatwg.org/#list-iterate){#reflecting-content-attributes-in-idl-attributes:list-iterate-3
      x-internal="list-iterate"} `element`{.variable} in the given
      value:

      1.  [Append](https://infra.spec.whatwg.org/#list-append){#reflecting-content-attributes-in-idl-attributes:list-append-3
          x-internal="list-append"} a weak reference to
          `element`{.variable} to `elements`{.variable}.

  5.  Set
      [this](https://webidl.spec.whatwg.org/#this){#reflecting-content-attributes-in-idl-attributes:this-34
      x-internal="this"}\'s [explicitly set
      `attr`{.variable}-elements](#explicitly-set-attr-elements){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-elements-4}
      to `elements`{.variable}.

- For element [reflected
  targets](#reflected-target){#reflecting-content-attributes-in-idl-attributes:reflected-target-17}
  only: the following [attribute change
  steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#reflecting-content-attributes-in-idl-attributes:concept-element-attributes-change-ext-2
  x-internal="concept-element-attributes-change-ext"}, given
  `element`{.variable}, `localName`{.variable}, `oldValue`{.variable},
  `value`{.variable}, and `namespace`{.variable}, are used to
  synchronize between the content attribute and the IDL attribute:

  1.  If `localName`{.variable} is not `attr`{.variable} or
      `namespace`{.variable} is not null, then return.

  2.  Set `element`{.variable}\'s [explicitly set
      `attr`{.variable}-elements](#explicitly-set-attr-elements){#reflecting-content-attributes-in-idl-attributes:explicitly-set-attr-elements-5}
      to null.

[Reflected IDL
attributes](#reflected-idl-attribute){#reflecting-content-attributes-in-idl-attributes:reflected-idl-attribute-34}
of this type are strongly encouraged to have their identifier end in
\"`Elements`\" for consistency.

#### [2.6.2]{.secno} Using reflect in specifications[](#using-reflect-in-specifications){.self-link}

[Reflection](#reflect){#using-reflect-in-specifications:reflect} is
primarily about improving web developer ergonomics by giving them typed
access to content attributes through [reflected IDL
attributes](#reflected-idl-attribute){#using-reflect-in-specifications:reflected-idl-attribute}.
The ultimate source of truth, which the web platform builds upon, is the
content attributes themselves. That is, specification authors must not
use the [reflected IDL
attribute](#reflected-idl-attribute){#using-reflect-in-specifications:reflected-idl-attribute-2}
getter or setter steps, but instead must use the content attribute
presence and value. (Or an abstraction on top, such as the state of an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#using-reflect-in-specifications:enumerated-attribute}.)

Two important exceptions to this are [reflected IDL
attributes](#reflected-idl-attribute){#using-reflect-in-specifications:reflected-idl-attribute-3}
whose type is one of the following:

- `T`{.variable}`?`, where `T`{.variable} is either
  [`Element`{#using-reflect-in-specifications:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
  or an interface that inherits from
  [`Element`{#using-reflect-in-specifications:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}

- `FrozenArray<``T`{.variable}`>?`, where `T`{.variable} is either
  [`Element`{#using-reflect-in-specifications:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
  or an interface that inherits from
  [`Element`{#using-reflect-in-specifications:element-4}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}

For those, specification authors must use the [reflected
target](#reflected-target){#using-reflect-in-specifications:reflected-target}\'s
[get the `attr`{.variable}-associated
element](#attr-associated-element){#using-reflect-in-specifications:attr-associated-element}
and [get the `attr`{.variable}-associated
elements](#attr-associated-elements){#using-reflect-in-specifications:attr-associated-elements},
respectively. The content attribute presence and value must not be used
as they cannot be fully synchronized with the [reflected IDL
attribute](#reflected-idl-attribute){#using-reflect-in-specifications:reflected-idl-attribute-4}.

A [reflected
target](#reflected-target){#using-reflect-in-specifications:reflected-target-2}\'s
[explicitly set
`attr`{.variable}-element](#explicitly-set-attr-element){#using-reflect-in-specifications:explicitly-set-attr-element},
[explicitly set
`attr`{.variable}-elements](#explicitly-set-attr-elements){#using-reflect-in-specifications:explicitly-set-attr-elements},
[cached `attr`{.variable}-associated
elements](#cached-attr-associated-elements){#using-reflect-in-specifications:cached-attr-associated-elements},
and [cached `attr`{.variable}-associated elements
object](#cached-attr-associated-elements-object){#using-reflect-in-specifications:cached-attr-associated-elements-object}
are to be treated as internal implementation details and not to be built
upon.

#### [2.6.3]{.secno} Collections[](#collections){.self-link}

The
[`HTMLFormControlsCollection`{#collections:htmlformcontrolscollection}](#htmlformcontrolscollection)
and
[`HTMLOptionsCollection`{#collections:htmloptionscollection}](#htmloptionscollection)
interfaces are
[collections](https://dom.spec.whatwg.org/#concept-collection){#collections:concept-collection
x-internal="concept-collection"} derived from the
[`HTMLCollection`{#collections:htmlcollection}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
interface. The
[`HTMLAllCollection`{#collections:htmlallcollection}](#htmlallcollection)
interface is a
[collection](https://dom.spec.whatwg.org/#concept-collection){#collections:concept-collection-2
x-internal="concept-collection"}, but is not so derived.

##### [2.6.3.1]{.secno} The [`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection}](#htmlallcollection) interface[](#the-htmlallcollection-interface){.self-link}

The
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-2}](#htmlallcollection)
interface is used for the legacy
[`document.all`{#the-htmlallcollection-interface:dom-document-all}](obsolete.html#dom-document-all)
attribute. It operates similarly to
[`HTMLCollection`{#the-htmlallcollection-interface:htmlcollection}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"};
the main differences are that it allows a staggering variety of
different (ab)uses of its methods to all end up returning something, and
that it can be called as a function as an alternative to property
access.

All
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-3}](#htmlallcollection)
objects are rooted at a
[`Document`{#the-htmlallcollection-interface:document}](dom.html#document)
and have a filter that matches all elements, so the elements
[represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmlallcollection-interface:represented-by-the-collection
x-internal="represented-by-the-collection"} of an
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-4}](#htmlallcollection)
object consist of all the descendant elements of the root
[`Document`{#the-htmlallcollection-interface:document-2}](dom.html#document).

Objects that implement the
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-5}](#htmlallcollection)
interface are [legacy platform
objects](https://webidl.spec.whatwg.org/#dfn-legacy-platform-object){#the-htmlallcollection-interface:legacy-platform-object
x-internal="legacy-platform-object"} with an additional \[\[Call\]\]
internal method described in the [section
below](#HTMLAllCollection-call). They also have an
[\[\[IsHTMLDDA\]\]](https://tc39.es/ecma262/#sec-IsHTMLDDA-internal-slot){#the-htmlallcollection-interface:ishtmldda
x-internal="ishtmldda"} internal slot.

::: note
Objects that implement the
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-6}](#htmlallcollection)
interface have several unusual behaviors, due of the fact that they have
an
[\[\[IsHTMLDDA\]\]](https://tc39.es/ecma262/#sec-IsHTMLDDA-internal-slot){#the-htmlallcollection-interface:ishtmldda-2
x-internal="ishtmldda"} internal slot:

- The
  [ToBoolean](https://tc39.es/ecma262/#sec-toboolean){#the-htmlallcollection-interface:js-toboolean
  x-internal="js-toboolean"} abstract operation in JavaScript returns
  false when given objects implementing the
  [`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-7}](#htmlallcollection)
  interface.

- The
  [IsLooselyEqual](https://tc39.es/ecma262/#sec-islooselyequal){#the-htmlallcollection-interface:js-abstract-equality
  x-internal="js-abstract-equality"} abstract operation, when given
  objects implementing the
  [`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-8}](#htmlallcollection)
  interface, returns true when compared to the `undefined` and `null`
  values. (Comparisons using the
  [IsStrictlyEqual](https://tc39.es/ecma262/#sec-isstrictlyequal){#the-htmlallcollection-interface:js-strict-equality
  x-internal="js-strict-equality"} abstract operation, and
  IsLooselyEqual comparisons to other values such as strings or objects,
  are unaffected.)

- The
  [`typeof`{#the-htmlallcollection-interface:js-typeof}](https://tc39.es/ecma262/#sec-typeof-operator){x-internal="js-typeof"}
  operator in JavaScript returns the string `"undefined"` when applied
  to objects implementing the
  [`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-9}](#htmlallcollection)
  interface.

These special behaviors are motivated by a desire for compatibility with
two classes of legacy content: one that uses the presence of
[`document.all`{#the-htmlallcollection-interface:dom-document-all-2}](obsolete.html#dom-document-all)
as a way to detect legacy user agents, and one that only supports those
legacy user agents and uses the
[`document.all`{#the-htmlallcollection-interface:dom-document-all-3}](obsolete.html#dom-document-all)
object without testing for its presence first.
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)
:::

``` idl
[Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface HTMLAllCollection {
  readonly attribute unsigned long length;
  getter Element (unsigned long index);
  getter (HTMLCollection or Element)? namedItem(DOMString name);
  (HTMLCollection or Element)? item(optional DOMString nameOrIndex);

  // Note: HTMLAllCollection objects have a custom [[Call]] internal method and an [[IsHTMLDDA]] internal slot.
};
```

The object\'s [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-htmlallcollection-interface:supported-property-indices
x-internal="supported-property-indices"} are as defined for
[`HTMLCollection`{#the-htmlallcollection-interface:htmlcollection-4}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
objects.

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#the-htmlallcollection-interface:supported-property-names
x-internal="supported-property-names"} consist of the non-empty values
of all the
[`id`{#the-htmlallcollection-interface:the-id-attribute}](dom.html#the-id-attribute)
attributes of all the elements [represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmlallcollection-interface:represented-by-the-collection-2
x-internal="represented-by-the-collection"}, and the non-empty values of
all the `name` attributes of all the [\"all\"-named
elements](#all-named-elements){#the-htmlallcollection-interface:all-named-elements}
[represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmlallcollection-interface:represented-by-the-collection-3
x-internal="represented-by-the-collection"}, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlallcollection-interface:tree-order
x-internal="tree-order"}, ignoring later duplicates, with the
[`id`{#the-htmlallcollection-interface:the-id-attribute-2}](dom.html#the-id-attribute)
of an element preceding its `name` if it contributes both, they differ
from each other, and neither is the duplicate of an earlier entry.

The [`length`]{#dom-htmlallcollection-length .dfn
dfn-for="HTMLAllCollection" dfn-type="attribute"} getter steps are to
return the number of nodes [represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmlallcollection-interface:represented-by-the-collection-4
x-internal="represented-by-the-collection"}.

The indexed property getter must return the result of [getting the
\"all\"-indexed
element](#concept-get-all-indexed){#the-htmlallcollection-interface:concept-get-all-indexed}
from
[this](https://webidl.spec.whatwg.org/#this){#the-htmlallcollection-interface:this
x-internal="this"} given the passed index.

The [`namedItem(``name`{.variable}`)`]{#dom-htmlallcollection-nameditem
.dfn dfn-for="HTMLAllCollection" dfn-type="method"} method steps are to
return the result of [getting the \"all\"-named
element(s)](#concept-get-all-named){#the-htmlallcollection-interface:concept-get-all-named}
from
[this](https://webidl.spec.whatwg.org/#this){#the-htmlallcollection-interface:this-2
x-internal="this"} given `name`{.variable}.

The [`item(``nameOrIndex`{.variable}`)`]{#dom-htmlallcollection-item
.dfn dfn-for="HTMLAllCollection" dfn-type="method"} method steps are:

1.  If `nameOrIndex`{.variable} was not provided, return null.

2.  Return the result of [getting the \"all\"-indexed or named
    element(s)](#concept-get-all-indexed-or-named){#the-htmlallcollection-interface:concept-get-all-indexed-or-named}
    from
    [this](https://webidl.spec.whatwg.org/#this){#the-htmlallcollection-interface:this-3
    x-internal="this"}, given `nameOrIndex`{.variable}.

------------------------------------------------------------------------

The following elements are [\"all\"-named elements]{#all-named-elements
.dfn}:
[`a`{#the-htmlallcollection-interface:the-a-element}](text-level-semantics.html#the-a-element),
[`button`{#the-htmlallcollection-interface:the-button-element}](form-elements.html#the-button-element),
[`embed`{#the-htmlallcollection-interface:the-embed-element}](iframe-embed-object.html#the-embed-element),
[`form`{#the-htmlallcollection-interface:the-form-element}](forms.html#the-form-element),
[`frame`{#the-htmlallcollection-interface:frame}](obsolete.html#frame),
[`frameset`{#the-htmlallcollection-interface:frameset}](obsolete.html#frameset),
[`iframe`{#the-htmlallcollection-interface:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`img`{#the-htmlallcollection-interface:the-img-element}](embedded-content.html#the-img-element),
[`input`{#the-htmlallcollection-interface:the-input-element}](input.html#the-input-element),
[`map`{#the-htmlallcollection-interface:the-map-element}](image-maps.html#the-map-element),
[`meta`{#the-htmlallcollection-interface:the-meta-element}](semantics.html#the-meta-element),
[`object`{#the-htmlallcollection-interface:the-object-element}](iframe-embed-object.html#the-object-element),
[`select`{#the-htmlallcollection-interface:the-select-element}](form-elements.html#the-select-element),
and
[`textarea`{#the-htmlallcollection-interface:the-textarea-element}](form-elements.html#the-textarea-element)

To [get the \"all\"-indexed element]{#concept-get-all-indexed .dfn} from
an
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-10}](#htmlallcollection)
`collection`{.variable} given an index `index`{.variable}, return the
`index`{.variable}^th^ element in `collection`{.variable}, or null if
there is no such `index`{.variable}^th^ element.

To [get the \"all\"-named element(s)]{#concept-get-all-named .dfn} from
an
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-11}](#htmlallcollection)
`collection`{.variable} given a name `name`{.variable}, perform the
following steps:

1.  If `name`{.variable} is the empty string, return null.

2.  Let `subCollection`{.variable} be an
    [`HTMLCollection`{#the-htmlallcollection-interface:htmlcollection-5}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    object rooted at the same
    [`Document`{#the-htmlallcollection-interface:document-3}](dom.html#document)
    as `collection`{.variable}, whose filter matches only elements that
    are either:

    - [\"all\"-named
      elements](#all-named-elements){#the-htmlallcollection-interface:all-named-elements-2}
      with a `name` attribute equal to `name`{.variable}, or,

    - elements with an
      [ID](https://dom.spec.whatwg.org/#concept-id){#the-htmlallcollection-interface:concept-id
      x-internal="concept-id"} equal to `name`{.variable}.

3.  If there is exactly one element in `subCollection`{.variable}, then
    return that element.

4.  Otherwise, if `subCollection`{.variable} is empty, return null.

5.  Otherwise, return `subCollection`{.variable}.

To [get the \"all\"-indexed or named
element(s)]{#concept-get-all-indexed-or-named .dfn} from an
[`HTMLAllCollection`{#the-htmlallcollection-interface:htmlallcollection-12}](#htmlallcollection)
`collection`{.variable} given `nameOrIndex`{.variable}:

1.  If `nameOrIndex`{.variable},
    [converted](https://webidl.spec.whatwg.org/#es-type-mapping){#the-htmlallcollection-interface:concept-idl-convert
    x-internal="concept-idl-convert"} to a JavaScript String value, is
    an [array index property
    name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#the-htmlallcollection-interface:array-index-property-name
    x-internal="array-index-property-name"}, return the result of
    [getting the \"all\"-indexed
    element](#concept-get-all-indexed){#the-htmlallcollection-interface:concept-get-all-indexed-2}
    from `collection`{.variable} given the number represented by
    `nameOrIndex`{.variable}.

2.  Return the result of [getting the \"all\"-named
    element(s)](#concept-get-all-named){#the-htmlallcollection-interface:concept-get-all-named-2}
    from `collection`{.variable} given `nameOrIndex`{.variable}.

###### [2.6.3.1.1]{.secno} \[\[Call\]\] ( `thisArgument`{.variable}, `argumentsList`{.variable} )[](#HTMLAllCollection-call){.self-link} {#HTMLAllCollection-call}

1.  If `argumentsList`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#HTMLAllCollection-call:list-size
    x-internal="list-size"} is zero, or if
    `argumentsList`{.variable}\[0\] is undefined, return null.

2.  Let `nameOrIndex`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#HTMLAllCollection-call:concept-idl-convert
    x-internal="concept-idl-convert"} `argumentsList`{.variable}\[0\] to
    a
    [`DOMString`{#HTMLAllCollection-call:idl-domstring}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"}.

3.  Let `result`{.variable} be the result of [getting the
    \"all\"-indexed or named
    element(s)](#concept-get-all-indexed-or-named){#HTMLAllCollection-call:concept-get-all-indexed-or-named}
    from this
    [`HTMLAllCollection`{#HTMLAllCollection-call:htmlallcollection}](#htmlallcollection)
    given `nameOrIndex`{.variable}.

4.  Return the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#HTMLAllCollection-call:concept-idl-convert-2
    x-internal="concept-idl-convert"} `result`{.variable} to an
    ECMAScript value.

The `thisArgument`{.variable} is ignored, and thus code such as
`Function.prototype.call.call(document.all, null, "x")` will still
search for elements. (`document.all.call` does not exist, since
`document.all` does not inherit from `Function.prototype`.)

##### [2.6.3.2]{.secno} The [`HTMLFormControlsCollection`{#the-htmlformcontrolscollection-interface:htmlformcontrolscollection}](#htmlformcontrolscollection) interface[](#the-htmlformcontrolscollection-interface){.self-link}

The
[`HTMLFormControlsCollection`{#the-htmlformcontrolscollection-interface:htmlformcontrolscollection-2}](#htmlformcontrolscollection)
interface is used for
[collections](https://dom.spec.whatwg.org/#concept-collection){#the-htmlformcontrolscollection-interface:concept-collection
x-internal="concept-collection"} of [listed
elements](forms.html#category-listed){#the-htmlformcontrolscollection-interface:category-listed}
in
[`form`{#the-htmlformcontrolscollection-interface:the-form-element}](forms.html#the-form-element)
elements.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormControlsCollection](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormControlsCollection "The HTMLFormControlsCollection interface represents a collection of HTML form control elements, returned by the HTMLFormElement interface's elements property.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
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

:::: feature
[RadioNodeList](https://developer.mozilla.org/en-US/docs/Web/API/RadioNodeList "The RadioNodeList interface represents a collection of radio elements in a <form> or a <fieldset> element.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome21+]{.chrome
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
:::::::

``` idl
[Exposed=Window]
interface HTMLFormControlsCollection : HTMLCollection {
  // inherits length and item()
  getter (RadioNodeList or Element)? namedItem(DOMString name); // shadows inherited namedItem()
};

[Exposed=Window]
interface RadioNodeList : NodeList {
  attribute DOMString value;
};
```

`collection`{.variable}`.`[`length`](https://dom.spec.whatwg.org/#dom-htmlcollection-length){#dom-htmlcollection-length-dev
x-internal="dom-htmlcollection-length"}

Returns the number of elements in `collection`{.variable}.

`element`{.variable}` = ``collection`{.variable}`.`[`item`](https://dom.spec.whatwg.org/#dom-htmlcollection-item){#dom-htmlcollection-item-dev
x-internal="dom-htmlcollection-item"}`(``index`{.variable}`)`

`element`{.variable}` = ``collection`{.variable}`[``index`{.variable}`]`

Returns the item at index `index`{.variable} in `collection`{.variable}.
The items are sorted in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order
x-internal="tree-order"}.

`element`{.variable}` = ``collection`{.variable}`.`[`namedItem`](#dom-htmlformcontrolscollection-nameditem){#dom-htmlformcontrolscollection-nameditem-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormControlsCollection/namedItem](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormControlsCollection/namedItem "The HTMLFormControlsCollection.namedItem() method returns the RadioNodeList or the Element in the collection whose name or id match the specified name, or null if no node matches.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
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

`radioNodeList`{.variable}` = ``collection`{.variable}`.`[`namedItem`](#dom-htmlformcontrolscollection-nameditem){#the-htmlformcontrolscollection-interface:dom-htmlformcontrolscollection-nameditem-2}`(``name`{.variable}`)`

`element`{.variable}` = ``collection`{.variable}`[``name`{.variable}`]`

`radioNodeList`{.variable}` = ``collection`{.variable}`[``name`{.variable}`]`

Returns the item with
[ID](https://dom.spec.whatwg.org/#concept-id){#the-htmlformcontrolscollection-interface:concept-id
x-internal="concept-id"} or
[`name`{#the-htmlformcontrolscollection-interface:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
`name`{.variable} from `collection`{.variable}.

If there are multiple matching items, then a
[`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-2}](#radionodelist)
object containing all those elements is returned.

`radioNodeList`{.variable}`.`[`value`](#dom-radionodelist-value){#dom-radionodelist-value-dev}

Returns the value of the first checked radio button represented by
`radioNodeList`{.variable}.

`radioNodeList`{.variable}`.`[`value`](#dom-radionodelist-value){#the-htmlformcontrolscollection-interface:dom-radionodelist-value-2}` = ``value`{.variable}

Checks the first radio button represented by `radioNodeList`{.variable}
that has value `value`{.variable}.

The object\'s [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-htmlformcontrolscollection-interface:supported-property-indices
x-internal="supported-property-indices"} are as defined for
[`HTMLCollection`{#the-htmlformcontrolscollection-interface:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
objects.

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#the-htmlformcontrolscollection-interface:supported-property-names
x-internal="supported-property-names"} consist of the non-empty values
of all the
[`id`{#the-htmlformcontrolscollection-interface:the-id-attribute}](dom.html#the-id-attribute)
and
[`name`{#the-htmlformcontrolscollection-interface:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attributes of all the elements [represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmlformcontrolscollection-interface:represented-by-the-collection
x-internal="represented-by-the-collection"}, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order-2
x-internal="tree-order"}, ignoring later duplicates, with the
[`id`{#the-htmlformcontrolscollection-interface:the-id-attribute-2}](dom.html#the-id-attribute)
of an element preceding its
[`name`{#the-htmlformcontrolscollection-interface:attr-fe-name-3}](form-control-infrastructure.html#attr-fe-name)
if it contributes both, they differ from each other, and neither is the
duplicate of an earlier entry.

The
[`namedItem(``name`{.variable}`)`]{#dom-htmlformcontrolscollection-nameditem
.dfn dfn-for="HTMLFormControlsCollection" dfn-type="method"} method must
act according to the following algorithm:

1.  If `name`{.variable} is the empty string, return null and stop the
    algorithm.
2.  If, at the time the method is called, there is exactly one node in
    the collection that has either an
    [`id`{#the-htmlformcontrolscollection-interface:the-id-attribute-3}](dom.html#the-id-attribute)
    attribute or a
    [`name`{#the-htmlformcontrolscollection-interface:attr-fe-name-4}](form-control-infrastructure.html#attr-fe-name)
    attribute equal to `name`{.variable}, then return that node and stop
    the algorithm.
3.  Otherwise, if there are no nodes in the collection that have either
    an
    [`id`{#the-htmlformcontrolscollection-interface:the-id-attribute-4}](dom.html#the-id-attribute)
    attribute or a
    [`name`{#the-htmlformcontrolscollection-interface:attr-fe-name-5}](form-control-infrastructure.html#attr-fe-name)
    attribute equal to `name`{.variable}, then return null and stop the
    algorithm.
4.  Otherwise, create a new
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-3}](#radionodelist)
    object representing a
    [live](infrastructure.html#live){#the-htmlformcontrolscollection-interface:live}
    view of the
    [`HTMLFormControlsCollection`{#the-htmlformcontrolscollection-interface:htmlformcontrolscollection-3}](#htmlformcontrolscollection)
    object, further filtered so that the only nodes in the
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-4}](#radionodelist)
    object are those that have either an
    [`id`{#the-htmlformcontrolscollection-interface:the-id-attribute-5}](dom.html#the-id-attribute)
    attribute or a
    [`name`{#the-htmlformcontrolscollection-interface:attr-fe-name-6}](form-control-infrastructure.html#attr-fe-name)
    attribute equal to `name`{.variable}. The nodes in the
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-5}](#radionodelist)
    object must be sorted in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order-3
    x-internal="tree-order"}.
5.  Return that
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-6}](#radionodelist)
    object.

------------------------------------------------------------------------

Members of the
[`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-7}](#radionodelist)
interface inherited from the
[`NodeList`{#the-htmlformcontrolscollection-interface:nodelist-2}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
interface must behave as they would on a
[`NodeList`{#the-htmlformcontrolscollection-interface:nodelist-3}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
object.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[RadioNodeList/value](https://developer.mozilla.org/en-US/docs/Web/API/RadioNodeList/value "If the underlying element collection contains radio buttons, the RadioNodeList.value property represents the checked radio button. On retrieving the value property, the value of the currently checked radio button is returned as a string. If the collection does not contain any radio buttons or none of the radio buttons in the collection is in checked state, the empty string is returned. On setting the value property, the first radio button input element whose value property is equal to the new value will be set to checked.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome21+]{.chrome
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

The [`value`]{#dom-radionodelist-value .dfn dfn-for="RadioNodeList"
dfn-type="attribute"} IDL attribute on the
[`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-8}](#radionodelist)
object, on getting, must return the value returned by running the
following steps:

1.  Let `element`{.variable} be the first element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order-4
    x-internal="tree-order"} represented by the
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-9}](#radionodelist)
    object that is an
    [`input`{#the-htmlformcontrolscollection-interface:the-input-element}](input.html#the-input-element)
    element whose
    [`type`{#the-htmlformcontrolscollection-interface:attr-input-type}](input.html#attr-input-type)
    attribute is in the [Radio
    Button](input.html#radio-button-state-(type=radio)){#the-htmlformcontrolscollection-interface:radio-button-state-(type=radio)}
    state and whose
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-htmlformcontrolscollection-interface:concept-fe-checked}
    is true. Otherwise, let it be null.

2.  If `element`{.variable} is null, return the empty string.

3.  If `element`{.variable} is an element with no
    [`value`{#the-htmlformcontrolscollection-interface:attr-input-value}](input.html#attr-input-value)
    attribute, return the string \"`on`\".

4.  Otherwise, return the value of `element`{.variable}\'s
    [`value`{#the-htmlformcontrolscollection-interface:attr-input-value-2}](input.html#attr-input-value)
    attribute.

On setting, the
[`value`{#the-htmlformcontrolscollection-interface:dom-radionodelist-value-3}](#dom-radionodelist-value)
IDL attribute must run the following steps:

1.  If the new value is the string \"`on`\": let `element`{.variable} be
    the first element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order-5
    x-internal="tree-order"} represented by the
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-10}](#radionodelist)
    object that is an
    [`input`{#the-htmlformcontrolscollection-interface:the-input-element-2}](input.html#the-input-element)
    element whose
    [`type`{#the-htmlformcontrolscollection-interface:attr-input-type-2}](input.html#attr-input-type)
    attribute is in the [Radio
    Button](input.html#radio-button-state-(type=radio)){#the-htmlformcontrolscollection-interface:radio-button-state-(type=radio)-2}
    state and whose
    [`value`{#the-htmlformcontrolscollection-interface:attr-input-value-3}](input.html#attr-input-value)
    content attribute is either absent, or present and equal to the new
    value, if any. If no such element exists, then instead let
    `element`{.variable} be null.

    Otherwise: let `element`{.variable} be the first element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmlformcontrolscollection-interface:tree-order-6
    x-internal="tree-order"} represented by the
    [`RadioNodeList`{#the-htmlformcontrolscollection-interface:radionodelist-11}](#radionodelist)
    object that is an
    [`input`{#the-htmlformcontrolscollection-interface:the-input-element-3}](input.html#the-input-element)
    element whose
    [`type`{#the-htmlformcontrolscollection-interface:attr-input-type-3}](input.html#attr-input-type)
    attribute is in the [Radio
    Button](input.html#radio-button-state-(type=radio)){#the-htmlformcontrolscollection-interface:radio-button-state-(type=radio)-3}
    state and whose
    [`value`{#the-htmlformcontrolscollection-interface:attr-input-value-4}](input.html#attr-input-value)
    content attribute is present and equal to the new value, if any. If
    no such element exists, then instead let `element`{.variable} be
    null.

2.  If `element`{.variable} is not null, then set its
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-htmlformcontrolscollection-interface:concept-fe-checked-2}
    to true.

##### [2.6.3.3]{.secno} The [`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection}](#htmloptionscollection) interface[](#the-htmloptionscollection-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTMLOptionsCollection](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOptionsCollection "The HTMLOptionsCollection interface represents a collection of <option> HTML elements (in document order) and offers methods and properties for selecting from the list as well as optionally altering its items. This object is returned only by the options property of select.")

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
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The
[`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-2}](#htmloptionscollection)
interface is used for
[collections](https://dom.spec.whatwg.org/#concept-collection){#the-htmloptionscollection-interface:concept-collection
x-internal="concept-collection"} of
[`option`{#the-htmloptionscollection-interface:the-option-element}](form-elements.html#the-option-element)
elements. It is always rooted on a
[`select`{#the-htmloptionscollection-interface:the-select-element}](form-elements.html#the-select-element)
element and has attributes and methods that manipulate that element\'s
descendants.

``` idl
[Exposed=Window]
interface HTMLOptionsCollection : HTMLCollection {
  // inherits item(), namedItem()
  [CEReactions] attribute unsigned long length; // shadows inherited length
  [CEReactions] setter undefined (unsigned long index, HTMLOptionElement? option);
  [CEReactions] undefined add((HTMLOptionElement or HTMLOptGroupElement) element, optional (HTMLElement or long)? before = null);
  [CEReactions] undefined remove(long index);
  attribute long selectedIndex;
};
```

`collection`{.variable}`.`[`length`](#dom-htmloptionscollection-length){#dom-htmloptionscollection-length-dev}
:   Returns the number of elements in `collection`{.variable}.

`collection`{.variable}`.`[`length`](#dom-htmloptionscollection-length){#the-htmloptionscollection-interface:dom-htmloptionscollection-length-2}` = ``value`{.variable}

:   When set to a smaller number than the existing length, truncates the
    number of
    [`option`{#the-htmloptionscollection-interface:the-option-element-2}](form-elements.html#the-option-element)
    elements in the container corresponding to `collection`{.variable}.

    When set to a greater number than the existing length, if that
    number is less than or equal to 100000, adds new blank
    [`option`{#the-htmloptionscollection-interface:the-option-element-3}](form-elements.html#the-option-element)
    elements to the container corresponding to `collection`{.variable}.

`element`{.variable}` = ``collection`{.variable}`.`[`item`](https://dom.spec.whatwg.org/#dom-htmlcollection-item){#the-htmloptionscollection-interface:dom-htmlcollection-item-2 x-internal="dom-htmlcollection-item"}`(``index`{.variable}`)`\
`element`{.variable}` = ``collection`{.variable}`[``index`{.variable}`]`
:   Returns the item at index `index`{.variable} in
    `collection`{.variable}. The items are sorted in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmloptionscollection-interface:tree-order
    x-internal="tree-order"}.

`collection`{.variable}`[``index`{.variable}`] = ``element`{.variable}

:   When `index`{.variable} is a greater number than the number of items
    in `collection`{.variable}, adds new blank
    [`option`{#the-htmloptionscollection-interface:the-option-element-4}](form-elements.html#the-option-element)
    elements in the corresponding container.

    When set to null, removes the item at index `index`{.variable} from
    `collection`{.variable}.

    When set to an
    [`option`{#the-htmloptionscollection-interface:the-option-element-5}](form-elements.html#the-option-element)
    element, adds or replaces it at index `index`{.variable} in
    `collection`{.variable}.

`element`{.variable}` = ``collection`{.variable}`.`[`namedItem`](https://dom.spec.whatwg.org/#dom-htmlcollection-nameditem){#dom-htmlcollection-nameditem-dev x-internal="dom-htmlcollection-nameditem"}`(``name`{.variable}`)`\
`element`{.variable}` = ``collection`{.variable}`[``name`{.variable}`]`

:   Returns the item with
    [ID](https://dom.spec.whatwg.org/#concept-id){#the-htmloptionscollection-interface:concept-id
    x-internal="concept-id"} or
    [`name`{#the-htmloptionscollection-interface:attr-option-name}](obsolete.html#attr-option-name)
    `name`{.variable} from `collection`{.variable}.

    If there are multiple matching items, then the first is returned.

`collection`{.variable}`.`[`add`](#dom-htmloptionscollection-add){#dom-htmloptionscollection-add-dev}`(``element`{.variable}`[, ``before`{.variable}`])`

:   Inserts `element`{.variable} before the node given by
    `before`{.variable}.

    The `before`{.variable} argument can be a number, in which case
    `element`{.variable} is inserted before the item with that number,
    or an element from `collection`{.variable}, in which case
    `element`{.variable} is inserted before that element.

    If `before`{.variable} is omitted, null, or a number out of range,
    then `element`{.variable} will be added at the end of the list.

    Throws a
    [\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-htmloptionscollection-interface:hierarchyrequesterror
    x-internal="hierarchyrequesterror"}
    [`DOMException`{#the-htmloptionscollection-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if `element`{.variable} is an ancestor of the element into which it
    is to be inserted.

`collection`{.variable}`.`[`remove`](#dom-htmloptionscollection-remove){#dom-htmloptionscollection-remove-dev}`(``index`{.variable}`)`
:   Removes the item with index `index`{.variable} from
    `collection`{.variable}.

`collection`{.variable}`.`[`selectedIndex`](#dom-htmloptionscollection-selectedindex){#the-htmloptionscollection-interface:dom-htmloptionscollection-selectedindex-2}
:   Returns the index of the first selected item, if any, or −1 if there
    is no selected item.

`collection`{.variable}`.`[`selectedIndex`](#dom-htmloptionscollection-selectedindex){#dom-htmloptionscollection-selectedindex-dev}` = ``index`{.variable}

:   Changes the selection to the
    [`option`{#the-htmloptionscollection-interface:the-option-element-6}](form-elements.html#the-option-element)
    element at index `index`{.variable} in `collection`{.variable}.

The object\'s [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-htmloptionscollection-interface:supported-property-indices
x-internal="supported-property-indices"} are as defined for
[`HTMLCollection`{#the-htmloptionscollection-interface:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
objects.

The [`length`]{#dom-htmloptionscollection-length .dfn
dfn-for="HTMLOptionsCollection" dfn-type="attribute"} getter steps are
to return the number of nodes [represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection
x-internal="represented-by-the-collection"}.

The
[`length`{#the-htmloptionscollection-interface:dom-htmloptionscollection-length-3}](#dom-htmloptionscollection-length)
setter steps are:

1.  Let `current`{.variable} be the number of nodes [represented by the
    collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection-2
    x-internal="represented-by-the-collection"}.

2.  If the given value is greater than `current`{.variable}, then:

    1.  If the given value is greater than 100,000, then return.

    2.  Let `n`{.variable} be `value`{.variable} − `current`{.variable}.

    3.  Append `n`{.variable} new
        [`option`{#the-htmloptionscollection-interface:the-option-element-7}](form-elements.html#the-option-element)
        elements with no attributes and no child nodes to the
        [`select`{#the-htmloptionscollection-interface:the-select-element-2}](form-elements.html#the-select-element)
        element on which
        [this](https://webidl.spec.whatwg.org/#this){#the-htmloptionscollection-interface:this
        x-internal="this"} is rooted.

3.  If the given value is less than `current`{.variable}, then:

    1.  Let `n`{.variable} be `current`{.variable} − `value`{.variable}.

    2.  Remove the last `n`{.variable} nodes in the collection from
        their parent nodes.

Setting
[`length`{#the-htmloptionscollection-interface:dom-htmloptionscollection-length-4}](#dom-htmloptionscollection-length)
never removes or adds any
[`optgroup`{#the-htmloptionscollection-interface:the-optgroup-element}](form-elements.html#the-optgroup-element)
elements, and never adds new children to existing
[`optgroup`{#the-htmloptionscollection-interface:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
elements (though it can remove children from them).

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#the-htmloptionscollection-interface:supported-property-names
x-internal="supported-property-names"} consist of the non-empty values
of all the
[`id`{#the-htmloptionscollection-interface:the-id-attribute}](dom.html#the-id-attribute)
and
[`name`{#the-htmloptionscollection-interface:attr-option-name-2}](obsolete.html#attr-option-name)
attributes of all the elements [represented by the
collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection-3
x-internal="represented-by-the-collection"}, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-htmloptionscollection-interface:tree-order-2
x-internal="tree-order"}, ignoring later duplicates, with the
[`id`{#the-htmloptionscollection-interface:the-id-attribute-2}](dom.html#the-id-attribute)
of an element preceding its
[`name`{#the-htmloptionscollection-interface:attr-option-name-3}](obsolete.html#attr-option-name)
if it contributes both, they differ from each other, and neither is the
duplicate of an earlier entry.

When the user agent is to [set the value of a new indexed
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-a-new-indexed-property){#the-htmloptionscollection-interface:set-the-value-of-a-new-indexed-property
x-internal="set-the-value-of-a-new-indexed-property"} or [set the value
of an existing indexed
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-an-existing-indexed-property){#the-htmloptionscollection-interface:set-the-value-of-an-existing-indexed-property
x-internal="set-the-value-of-an-existing-indexed-property"} for a given
property index `index`{.variable} to a new value `value`{.variable}, it
must run the following algorithm:

1.  If `value`{.variable} is null, invoke the steps for the
    [`remove`{#the-htmloptionscollection-interface:dom-htmloptionscollection-remove-2}](#dom-htmloptionscollection-remove)
    method with `index`{.variable} as the argument, and return.

2.  Let `length`{.variable} be the number of nodes [represented by the
    collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection-4
    x-internal="represented-by-the-collection"}.

3.  Let `n`{.variable} be `index`{.variable} minus `length`{.variable}.

4.  If `n`{.variable} is greater than zero, then
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-htmloptionscollection-interface:concept-node-append
    x-internal="concept-node-append"} a
    [`DocumentFragment`{#the-htmloptionscollection-interface:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    consisting of `n`{.variable}-1 new
    [`option`{#the-htmloptionscollection-interface:the-option-element-8}](form-elements.html#the-option-element)
    elements with no attributes and no child nodes to the
    [`select`{#the-htmloptionscollection-interface:the-select-element-3}](form-elements.html#the-select-element)
    element on which the
    [`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-3}](#htmloptionscollection)
    is rooted.

5.  If `n`{.variable} is greater than or equal to zero,
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-htmloptionscollection-interface:concept-node-append-2
    x-internal="concept-node-append"} `value`{.variable} to the
    [`select`{#the-htmloptionscollection-interface:the-select-element-4}](form-elements.html#the-select-element)
    element. Otherwise,
    [replace](https://dom.spec.whatwg.org/#concept-node-replace){#the-htmloptionscollection-interface:concept-node-replace
    x-internal="concept-node-replace"} the `index`{.variable}th element
    in the collection by `value`{.variable}.

The
[`add(``element`{.variable}`, ``before`{.variable}`)`]{#dom-htmloptionscollection-add
.dfn dfn-for="HTMLOptionsCollection" dfn-type="method"} method must act
according to the following algorithm:

1.  If `element`{.variable} is an ancestor of the
    [`select`{#the-htmloptionscollection-interface:the-select-element-5}](form-elements.html#the-select-element)
    element on which the
    [`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-4}](#htmloptionscollection)
    is rooted, then throw a
    [\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-htmloptionscollection-interface:hierarchyrequesterror-2
    x-internal="hierarchyrequesterror"}
    [`DOMException`{#the-htmloptionscollection-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `before`{.variable} is an element, but that element isn\'t a
    descendant of the
    [`select`{#the-htmloptionscollection-interface:the-select-element-6}](form-elements.html#the-select-element)
    element on which the
    [`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-5}](#htmloptionscollection)
    is rooted, then throw a
    [\"`NotFoundError`\"](https://webidl.spec.whatwg.org/#notfounderror){#the-htmloptionscollection-interface:notfounderror
    x-internal="notfounderror"}
    [`DOMException`{#the-htmloptionscollection-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `element`{.variable} and `before`{.variable} are the same
    element, then return.

4.  If `before`{.variable} is a node, then let `reference`{.variable} be
    that node. Otherwise, if `before`{.variable} is an integer, and
    there is a `before`{.variable}th node in the collection, let
    `reference`{.variable} be that node. Otherwise, let
    `reference`{.variable} be null.

5.  If `reference`{.variable} is not null, let `parent`{.variable} be
    the parent node of `reference`{.variable}. Otherwise, let
    `parent`{.variable} be the
    [`select`{#the-htmloptionscollection-interface:the-select-element-7}](form-elements.html#the-select-element)
    element on which the
    [`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-6}](#htmloptionscollection)
    is rooted.

6.  [Pre-insert](https://dom.spec.whatwg.org/#concept-node-pre-insert){#the-htmloptionscollection-interface:pre-insert
    x-internal="pre-insert"} `element`{.variable} into
    `parent`{.variable} node before `reference`{.variable}.

The [`remove(``index`{.variable}`)`]{#dom-htmloptionscollection-remove
.dfn dfn-for="HTMLOptionsCollection" dfn-type="method"} method must act
according to the following algorithm:

1.  If the number of nodes [represented by the
    collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection-5
    x-internal="represented-by-the-collection"} is zero, return.

2.  If `index`{.variable} is not a number greater than or equal to 0 and
    less than the number of nodes [represented by the
    collection](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-htmloptionscollection-interface:represented-by-the-collection-6
    x-internal="represented-by-the-collection"}, return.

3.  Let `element`{.variable} be the `index`{.variable}th element in the
    collection.

4.  Remove `element`{.variable} from its parent node.

The [`selectedIndex`]{#dom-htmloptionscollection-selectedindex .dfn
dfn-for="HTMLOptionsCollection" dfn-type="attribute"} IDL attribute must
act like the identically named attribute on the
[`select`{#the-htmloptionscollection-interface:the-select-element-8}](form-elements.html#the-select-element)
element on which the
[`HTMLOptionsCollection`{#the-htmloptionscollection-interface:htmloptionscollection-7}](#htmloptionscollection)
is rooted

#### [2.6.4]{.secno} The [`DOMStringList`{#the-domstringlist-interface:domstringlist}](#domstringlist) interface[](#the-domstringlist-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DOMStringList](https://developer.mozilla.org/en-US/docs/Web/API/DOMStringList "The DOMString interface is a legacy type returned by some APIs and represents a non-modifiable list of strings (DOMString). Modern APIs use Array objects (in WebIDL: sequence<DOMString>) instead.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome1+]{.chrome
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

The
[`DOMStringList`{#the-domstringlist-interface:domstringlist-2}](#domstringlist)
interface is a non-fashionable retro way of representing a list of
strings.

``` idl
[Exposed=(Window,Worker)]
interface DOMStringList {
  readonly attribute unsigned long length;
  getter DOMString? item(unsigned long index);
  boolean contains(DOMString string);
};
```

New APIs must use `sequence<DOMString>` or equivalent rather than
[`DOMStringList`{#the-domstringlist-interface:domstringlist-3}](#domstringlist).

`strings`{.variable}`.`[`length`](#dom-domstringlist-length){#dom-domstringlist-length-dev}
:   Returns the number of strings in `strings`{.variable}.

`strings`{.variable}`[``index`{.variable}`]`\
`strings`{.variable}`.`[`item`](#dom-domstringlist-item){#dom-domstringlist-item-dev}`(``index`{.variable}`)`
:   Returns the string with index `index`{.variable} from
    `strings`{.variable}.

`strings`{.variable}`.`[`contains`](#dom-domstringlist-contains){#dom-domstringlist-contains-dev}`(``string`{.variable}`)`

:   Returns true if `strings`{.variable} contains `string`{.variable},
    and false otherwise.

Each
[`DOMStringList`{#the-domstringlist-interface:domstringlist-4}](#domstringlist)
object has an associated
[list](https://infra.spec.whatwg.org/#list){#the-domstringlist-interface:list
x-internal="list"}.

The
[`DOMStringList`{#the-domstringlist-interface:domstringlist-5}](#domstringlist)
interface [supports indexed
properties](https://webidl.spec.whatwg.org/#dfn-support-indexed-properties){#the-domstringlist-interface:supports-indexed-properties
x-internal="supports-indexed-properties"}. The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-domstringlist-interface:supported-property-indices
x-internal="supported-property-indices"} are the
[indices](https://infra.spec.whatwg.org/#list-get-the-indices){#the-domstringlist-interface:indices
x-internal="indices"} of
[this](https://webidl.spec.whatwg.org/#this){#the-domstringlist-interface:this
x-internal="this"}\'s associated list.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DOMStringList/length](https://developer.mozilla.org/en-US/docs/Web/API/DOMStringList/length "The read-only length property indicates the number of strings in the DOMStringList.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome1+]{.chrome
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

The [`length`]{#dom-domstringlist-length .dfn dfn-for="DOMStringList"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-domstringlist-interface:this-2
x-internal="this"}\'s associated list\'s
[size](https://infra.spec.whatwg.org/#list-size){#the-domstringlist-interface:list-size
x-internal="list-size"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DOMStringList/item](https://developer.mozilla.org/en-US/docs/Web/API/DOMStringList/item "The item() method returns a string from a DOMStringList by index.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome1+]{.chrome
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

The [`item(``index`{.variable}`)`]{#dom-domstringlist-item .dfn
dfn-for="DOMStringList" dfn-type="method"} method steps are to return
the `index`{.variable}th item in
[this](https://webidl.spec.whatwg.org/#this){#the-domstringlist-interface:this-3
x-internal="this"}\'s associated list, or null if `index`{.variable}
plus one is greater than
[this](https://webidl.spec.whatwg.org/#this){#the-domstringlist-interface:this-4
x-internal="this"}\'s associated list\'s
[size](https://infra.spec.whatwg.org/#list-size){#the-domstringlist-interface:list-size-2
x-internal="list-size"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DOMStringList/contains](https://developer.mozilla.org/en-US/docs/Web/API/DOMStringList/contains "The contains() method returns a boolean indicating whether the given string is in the list.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome1+]{.chrome
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

The [`contains(``string`{.variable}`)`]{#dom-domstringlist-contains .dfn
dfn-for="DOMStringList" dfn-type="method"} method steps are to return
true if
[this](https://webidl.spec.whatwg.org/#this){#the-domstringlist-interface:this-5
x-internal="this"}\'s associated list
[contains](https://infra.spec.whatwg.org/#list-contain){#the-domstringlist-interface:list-contains
x-internal="list-contains"} `string`{.variable}, and false otherwise.

[← 2.4 URLs](urls-and-fetching.html) --- [Table of Contents](./) ---
[2.7 Safe passing of structured data →](structured-data.html)
