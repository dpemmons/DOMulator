## [7]{.secno} Loading web pages[](#browsers){.self-link} {#browsers}

This section describes features that apply most directly to web
browsers. Having said that, except where specified otherwise, the
requirements defined in this section *do* apply to all user agents,
whether they are web browsers or not.

### [7.1]{.secno} Supporting concepts[](#loading-web-pages-supporting-concepts){.self-link} {#loading-web-pages-supporting-concepts}

#### [7.1.1]{.secno} Origins[](#origin){.self-link} {#origin}

Origins are the fundamental currency of the web\'s security model. Two
actors in the web platform that share an origin are assumed to trust
each other and to have the same authority. Actors with differing origins
are considered potentially hostile versus each other, and are isolated
from each other to varying degrees.

For example, if Example Bank\'s web site, hosted at `bank.example.com`,
tries to examine the DOM of Example Charity\'s web site, hosted at
`charity.example.org`, a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#origin:securityerror
x-internal="securityerror"}
[`DOMException`{#origin:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
will be raised.

------------------------------------------------------------------------

An [origin]{#concept-origin .dfn export=""} is one of the following:

An [opaque origin]{#concept-origin-opaque .dfn export=""}
:   An internal value, with no serialization it can be recreated from
    (it is serialized as \"`null`\" per [serialization of an
    origin](#ascii-serialisation-of-an-origin){#origin:ascii-serialisation-of-an-origin}),
    for which the only meaningful operation is testing for equality.

A [tuple origin]{#concept-origin-tuple .dfn export=""}

:   A [tuple](https://infra.spec.whatwg.org/#tuple){#origin:tuple
    x-internal="tuple"} consisting of:

    - A [scheme]{#concept-origin-scheme .dfn dfn-for="origin" export=""}
      (an [ASCII
      string](https://infra.spec.whatwg.org/#ascii-string){#origin:ascii-string
      x-internal="ascii-string"}).
    - A [host]{#concept-origin-host .dfn dfn-for="origin" export=""} (a
      [host](https://url.spec.whatwg.org/#concept-host){#origin:concept-host
      x-internal="concept-host"}).
    - A [port]{#concept-origin-port .dfn dfn-for="origin" export=""}
      (null or a 16-bit unsigned integer).
    - A [domain]{#concept-origin-domain .dfn dfn-for="origin" export=""}
      (null or a
      [domain](https://url.spec.whatwg.org/#concept-domain){#origin:concept-domain
      x-internal="concept-domain"}). Null unless stated otherwise.

[Origins](#concept-origin){#origin:concept-origin} can be shared, e.g.,
among multiple [`Document`{#origin:document}](dom.html#document)
objects. Furthermore,
[origins](#concept-origin){#origin:concept-origin-2} are generally
immutable. Only the
[domain](#concept-origin-domain){#origin:concept-origin-domain} of a
[tuple origin](#concept-origin-tuple){#origin:concept-origin-tuple} can
be changed, and only through the
[`document.domain`{#origin:dom-document-domain}](#dom-document-domain)
API.

The [effective domain]{#concept-origin-effective-domain .dfn export=""}
of an [origin](#concept-origin){#origin:concept-origin-3}
`origin`{.variable} is computed as follows:

1.  If `origin`{.variable} is an [opaque
    origin](#concept-origin-opaque){#origin:concept-origin-opaque}, then
    return null.

2.  If `origin`{.variable}\'s
    [domain](#concept-origin-domain){#origin:concept-origin-domain-2} is
    non-null, then return `origin`{.variable}\'s
    [domain](#concept-origin-domain){#origin:concept-origin-domain-3}.

3.  Return `origin`{.variable}\'s
    [host](#concept-origin-host){#origin:concept-origin-host}.

The [serialization of an origin]{#ascii-serialisation-of-an-origin .dfn
lt="serialization of an
  origin|ASCII serialization of an origin" export=""} is the string
obtained by applying the following algorithm to the given
[origin](#concept-origin){#origin:concept-origin-4} `origin`{.variable}:

1.  If `origin`{.variable} is an [opaque
    origin](#concept-origin-opaque){#origin:concept-origin-opaque-2},
    then return \"`null`\".

2.  Otherwise, let `result`{.variable} be `origin`{.variable}\'s
    [scheme](#concept-origin-scheme){#origin:concept-origin-scheme}.

3.  Append \"`://`\" to `result`{.variable}.

4.  Append `origin`{.variable}\'s
    [host](#concept-origin-host){#origin:concept-origin-host-2},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#origin:host-serializer
    x-internal="host-serializer"}, to `result`{.variable}.

5.  If `origin`{.variable}\'s
    [port](#concept-origin-port){#origin:concept-origin-port} is
    non-null, append a U+003A COLON character (:), and
    `origin`{.variable}\'s
    [port](#concept-origin-port){#origin:concept-origin-port-2},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#origin:serialize-an-integer
    x-internal="serialize-an-integer"}, to `result`{.variable}.

6.  Return `result`{.variable}.

The
[serialization](#ascii-serialisation-of-an-origin){#origin:ascii-serialisation-of-an-origin-2}
of (\"`https`\", \"`xn--maraa-rta.example`\", null, null) is
\"`https://xn--maraa-rta.example`\".

There used to also be a *Unicode serialization of an origin*. However,
it was never widely adopted.

------------------------------------------------------------------------

Two [origins](#concept-origin){#origin:concept-origin-5}, `A`{.variable}
and `B`{.variable}, are said to be [same origin]{#same-origin .dfn
export=""} if the following algorithm returns true:

1.  If `A`{.variable} and `B`{.variable} are the same [opaque
    origin](#concept-origin-opaque){#origin:concept-origin-opaque-3},
    then return true.

2.  If `A`{.variable} and `B`{.variable} are both [tuple
    origins](#concept-origin-tuple){#origin:concept-origin-tuple-2} and
    their
    [schemes](#concept-origin-scheme){#origin:concept-origin-scheme-2},
    [hosts](#concept-origin-host){#origin:concept-origin-host-3}, and
    [port](#concept-origin-port){#origin:concept-origin-port-3} are
    identical, then return true.

3.  Return false.

Two [origins](#concept-origin){#origin:concept-origin-6}, `A`{.variable}
and `B`{.variable}, are said to be [same
origin-domain]{#same-origin-domain .dfn export=""} if the following
algorithm returns true:

1.  If `A`{.variable} and `B`{.variable} are the same [opaque
    origin](#concept-origin-opaque){#origin:concept-origin-opaque-4},
    then return true.

2.  If `A`{.variable} and `B`{.variable} are both [tuple
    origins](#concept-origin-tuple){#origin:concept-origin-tuple-3}:

    1.  If `A`{.variable} and `B`{.variable}\'s
        [schemes](#concept-origin-scheme){#origin:concept-origin-scheme-3}
        are identical, and their
        [domains](#concept-origin-domain){#origin:concept-origin-domain-4}
        are identical and non-null, then return true.

    2.  Otherwise, if `A`{.variable} and `B`{.variable} are [same
        origin](#same-origin){#origin:same-origin} and their
        [domains](#concept-origin-domain){#origin:concept-origin-domain-5}
        are both null, return true.

3.  Return false.

::: example
`A`{.variable}

`B`{.variable}

[same origin](#same-origin){#origin:same-origin-2}

[same origin-domain](#same-origin-domain){#origin:same-origin-domain}

(\"`https`\", \"`example.org`\", null, null)

(\"`https`\", \"`example.org`\", null, null)

✅

✅

(\"`https`\", \"`example.org`\", 314, null)

(\"`https`\", \"`example.org`\", 420, null)

❌

❌

(\"`https`\", \"`example.org`\", 314, \"`example.org`\")

(\"`https`\", \"`example.org`\", 420, \"`example.org`\")

❌

✅

(\"`https`\", \"`example.org`\", null, null)

(\"`https`\", \"`example.org`\", null, \"`example.org`\")

✅

❌

(\"`https`\", \"`example.org`\", null, \"`example.org`\")

(\"`http`\", \"`example.org`\", null, \"`example.org`\")

❌

❌
:::

##### [7.1.1.1]{.secno} Sites[](#sites){.self-link}

A [scheme-and-host]{#scheme-and-host .dfn} is a
[tuple](https://infra.spec.whatwg.org/#tuple){#sites:tuple
x-internal="tuple"} of a [scheme]{#concept-scheme-and-host-scheme .dfn}
(an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#sites:ascii-string
x-internal="ascii-string"}) and a [host]{#concept-scheme-and-host-host
.dfn} (a
[host](https://url.spec.whatwg.org/#concept-host){#sites:concept-host
x-internal="concept-host"}).

A [site]{#site .dfn export=""} is an [opaque
origin](#concept-origin-opaque){#sites:concept-origin-opaque} or a
[scheme-and-host](#scheme-and-host){#sites:scheme-and-host}.

To [obtain a site]{#obtain-a-site .dfn export=""}, given an origin
`origin`{.variable}, run these steps:

1.  If `origin`{.variable} is an [opaque
    origin](#concept-origin-opaque){#sites:concept-origin-opaque-2},
    then return `origin`{.variable}.

2.  If `origin`{.variable}\'s
    [host](#concept-origin-host){#sites:concept-origin-host}\'s
    [registrable
    domain](https://url.spec.whatwg.org/#host-registrable-domain){#sites:registrable-domain
    x-internal="registrable-domain"} is null, then return
    (`origin`{.variable}\'s
    [scheme](#concept-origin-scheme){#sites:concept-origin-scheme},
    `origin`{.variable}\'s
    [host](#concept-origin-host){#sites:concept-origin-host-2}).

3.  Return (`origin`{.variable}\'s
    [scheme](#concept-origin-scheme){#sites:concept-origin-scheme-2},
    `origin`{.variable}\'s
    [host](#concept-origin-host){#sites:concept-origin-host-3}\'s
    [registrable
    domain](https://url.spec.whatwg.org/#host-registrable-domain){#sites:registrable-domain-2
    x-internal="registrable-domain"}).

Two [sites](#site){#sites:site}, `A`{.variable} and `B`{.variable}, are
said to be [same site]{#concept-site-same-site .dfn dfn-for="site"
export=""} if the following algorithm returns true:

1.  If `A`{.variable} and `B`{.variable} are the same [opaque
    origin](#concept-origin-opaque){#sites:concept-origin-opaque-3},
    then return true.

2.  If `A`{.variable} or `B`{.variable} is an [opaque
    origin](#concept-origin-opaque){#sites:concept-origin-opaque-4},
    then return false.

3.  If `A`{.variable}\'s and `B`{.variable}\'s
    [scheme](#concept-scheme-and-host-scheme){#sites:concept-scheme-and-host-scheme}
    values are different, then return false.

4.  If `A`{.variable}\'s and `B`{.variable}\'s
    [host](#concept-scheme-and-host-host){#sites:concept-scheme-and-host-host}
    values are not
    [equal](https://url.spec.whatwg.org/#concept-host-equals){#sites:host-equals
    x-internal="host-equals"}, then return false.

5.  Return true.

The [serialization of a site]{#serialization-of-a-site .dfn export=""}
is the string obtained by applying the following algorithm to the given
[site](#site){#sites:site-2} `site`{.variable}:

1.  If `site`{.variable} is an [opaque
    origin](#concept-origin-opaque){#sites:concept-origin-opaque-5},
    then return \"`null`\".

2.  Let `result`{.variable} be `site`{.variable}\[0\].

3.  Append \"`://`\" to `result`{.variable}.

4.  Append `site`{.variable}\[1\],
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#sites:host-serializer
    x-internal="host-serializer"}, to `result`{.variable}.

5.  Return `result`{.variable}.

It needs to be clear from context that the serialized value is a site,
not an origin, as there is not necessarily a syntactic difference
between the two. For example, the origin (\"`https`\",
\"`shop.example`\", null, null) and the site (\"`https`\",
\"`shop.example`\") have the same serialization:
\"`https://shop.example`\".

Two [origins](#concept-origin){#sites:concept-origin}, `A`{.variable}
and `B`{.variable}, are said to be [schemelessly same
site]{#schemelessly-same-site .dfn export=""} if the following algorithm
returns true:

1.  If `A`{.variable} and `B`{.variable} are the same [opaque
    origin](#concept-origin-opaque){#sites:concept-origin-opaque-6},
    then return true.

2.  If `A`{.variable} and `B`{.variable} are both [tuple
    origins](#concept-origin-tuple){#sites:concept-origin-tuple}, then:

    1.  Let `hostA`{.variable} be `A`{.variable}\'s
        [host](#concept-origin-host){#sites:concept-origin-host-4}, and
        let `hostB`{.variable} be `B`{.variable}\'s
        [host](#concept-origin-host){#sites:concept-origin-host-5}.

    2.  If `hostA`{.variable}
        [equals](https://url.spec.whatwg.org/#concept-host-equals){#sites:host-equals-2
        x-internal="host-equals"} `hostB`{.variable} and
        `hostA`{.variable}\'s [registrable
        domain](https://url.spec.whatwg.org/#host-registrable-domain){#sites:registrable-domain-3
        x-internal="registrable-domain"} is null, then return true.

    3.  If `hostA`{.variable}\'s [registrable
        domain](https://url.spec.whatwg.org/#host-registrable-domain){#sites:registrable-domain-4
        x-internal="registrable-domain"}
        [equals](https://url.spec.whatwg.org/#concept-host-equals){#sites:host-equals-3
        x-internal="host-equals"} `hostB`{.variable}\'s [registrable
        domain](https://url.spec.whatwg.org/#host-registrable-domain){#sites:registrable-domain-5
        x-internal="registrable-domain"} and is non-null, then return
        true.

3.  Return false.

Two [origins](#concept-origin){#sites:concept-origin-2}, `A`{.variable}
and `B`{.variable}, are said to be [same site]{#same-site .dfn
export=""} if the following algorithm returns true:

1.  Let `siteA`{.variable} be the result of [obtaining a
    site](#obtain-a-site){#sites:obtain-a-site} given `A`{.variable}.

2.  Let `siteB`{.variable} be the result of [obtaining a
    site](#obtain-a-site){#sites:obtain-a-site-2} given `B`{.variable}.

3.  If `siteA`{.variable} is [same
    site](#concept-site-same-site){#sites:concept-site-same-site} with
    `siteB`{.variable}, then return true.

4.  Return false.

Unlike the [same origin](#same-origin){#sites:same-origin} and [same
origin-domain](#same-origin-domain){#sites:same-origin-domain} concepts,
for [schemelessly same
site](#schemelessly-same-site){#sites:schemelessly-same-site} and [same
site](#same-site){#sites:same-site}, the
[port](#concept-origin-port){#sites:concept-origin-port} and
[domain](#concept-origin-domain){#sites:concept-origin-domain}
components are ignored.

For the reasons [explained in
URL](https://url.spec.whatwg.org/#warning-avoid-psl), the [same
site](#same-site){#sites:same-site-2} and [schemelessly same
site](#schemelessly-same-site){#sites:schemelessly-same-site-2} concepts
should be avoided when possible, in favor of [same
origin](#same-origin){#sites:same-origin-2} checks.

::: {#example-same-site .example}
Given that `wildlife.museum`, `museum`, and `com` are [public
suffixes](https://url.spec.whatwg.org/#host-public-suffix){#sites:public-suffix
x-internal="public-suffix"} and that `example.com` is not:

`A`{.variable}

`B`{.variable}

[schemelessly same
site](#schemelessly-same-site){#sites:schemelessly-same-site-3}

[same site](#same-site){#sites:same-site-3}

(\"`https`\", \"`example.com`\")

(\"`https`\", \"`sub.example.com`\")

✅

✅

(\"`https`\", \"`example.com`\")

(\"`https`\", \"`sub.other.example.com`\")

✅

✅

(\"`https`\", \"`example.com`\")

(\"`http`\", \"`non-secure.example.com`\")

✅

❌

(\"`https`\", \"`r.wildlife.museum`\")

(\"`https`\", \"`sub.r.wildlife.museum`\")

✅

✅

(\"`https`\", \"`r.wildlife.museum`\")

(\"`https`\", \"`sub.other.r.wildlife.museum`\")

✅

✅

(\"`https`\", \"`r.wildlife.museum`\")

(\"`https`\", \"`other.wildlife.museum`\")

❌

❌

(\"`https`\", \"`r.wildlife.museum`\")

(\"`https`\", \"`wildlife.museum`\")

❌

❌

(\"`https`\", \"`wildlife.museum`\")

(\"`https`\", \"`wildlife.museum`\")

✅

✅

(\"`https`\", \"`example.com`\")

(\"`https`\", \"`example.com.`\")

❌

❌

(Here we have omitted the
[port](#concept-origin-port){#sites:concept-origin-port-2} and
[domain](#concept-origin-domain){#sites:concept-origin-domain-2}
components since they are not considered.)
:::

##### [7.1.1.2]{.secno} Relaxing the same-origin restriction[](#relaxing-the-same-origin-restriction){.self-link}

`document`{.variable}`.`[`domain`](#dom-document-domain){#dom-document-domain-dev}` [ = ``domain`{.variable}` ]`

:   Returns the current domain used for security checks.

    Can be set to a value that removes subdomains, to change the
    [origin](#concept-origin){#relaxing-the-same-origin-restriction:concept-origin}\'s
    [domain](#concept-origin-domain){#relaxing-the-same-origin-restriction:concept-origin-domain}
    to allow pages on other subdomains of the same domain (if they do
    the same thing) to access each other. This enables pages on
    different hosts of a domain to synchronously access each other\'s
    DOMs.

    In sandboxed
    [`iframe`{#relaxing-the-same-origin-restriction:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s,
    [`Document`{#relaxing-the-same-origin-restriction:document}](dom.html#document)s
    with [opaque
    origins](#concept-origin-opaque){#relaxing-the-same-origin-restriction:concept-origin-opaque},
    and
    [`Document`{#relaxing-the-same-origin-restriction:document-2}](dom.html#document)s
    without a [browsing
    context](document-sequences.html#concept-document-bc){#relaxing-the-same-origin-restriction:concept-document-bc},
    the setter will throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#relaxing-the-same-origin-restriction:securityerror
    x-internal="securityerror"} exception. In cases where
    [`crossOriginIsolated`{#relaxing-the-same-origin-restriction:dom-crossoriginisolated}](webappapis.html#dom-crossoriginisolated)
    or
    [`originAgentCluster`{#relaxing-the-same-origin-restriction:dom-originagentcluster}](#dom-originagentcluster)
    return true, the setter will do nothing.

::: critical
Avoid using the
[`document.domain`{#relaxing-the-same-origin-restriction:dom-document-domain}](#dom-document-domain)
setter. It undermines the security protections provided by the
same-origin policy. This is especially acute when using shared hosting;
for example, if an untrusted third party is able to host an HTTP server
at the same IP address but on a different port, then the same-origin
protection that normally protects two different sites on the same host
will fail, as the ports are ignored when comparing origins after the
[`document.domain`{#relaxing-the-same-origin-restriction:dom-document-domain-2}](#dom-document-domain)
setter has been used.

Because of these security pitfalls, this feature is in the process of
being removed from the web platform. (This is a long process that takes
many years.)

Instead, use
[`postMessage()`{#relaxing-the-same-origin-restriction:dom-window-postmessage}](web-messaging.html#dom-window-postmessage)
or
[`MessageChannel`{#relaxing-the-same-origin-restriction:messagechannel}](web-messaging.html#messagechannel)
objects to communicate across origins in a safe manner.
:::

The [`domain`]{#dom-document-domain .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are:

1.  Let `effectiveDomain`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#relaxing-the-same-origin-restriction:this
    x-internal="this"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#relaxing-the-same-origin-restriction:concept-document-origin
    x-internal="concept-document-origin"}\'s [effective
    domain](#concept-origin-effective-domain){#relaxing-the-same-origin-restriction:concept-origin-effective-domain}.

2.  If `effectiveDomain`{.variable} is null, then return the empty
    string.

3.  Return `effectiveDomain`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#relaxing-the-same-origin-restriction:host-serializer
    x-internal="host-serializer"}.

The
[`domain`{#relaxing-the-same-origin-restriction:dom-document-domain-3}](#dom-document-domain)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#relaxing-the-same-origin-restriction:this-2
    x-internal="this"}\'s [browsing
    context](document-sequences.html#concept-document-bc){#relaxing-the-same-origin-restriction:concept-document-bc-2}
    is null, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#relaxing-the-same-origin-restriction:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#relaxing-the-same-origin-restriction:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#relaxing-the-same-origin-restriction:this-3
    x-internal="this"}\'s [active sandboxing flag
    set](#active-sandboxing-flag-set){#relaxing-the-same-origin-restriction:active-sandboxing-flag-set}
    has its [sandboxed `document.domain` browsing context
    flag](#sandboxed-document.domain-browsing-context-flag){#relaxing-the-same-origin-restriction:sandboxed-document.domain-browsing-context-flag}
    set, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#relaxing-the-same-origin-restriction:securityerror-3
    x-internal="securityerror"}
    [`DOMException`{#relaxing-the-same-origin-restriction:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `effectiveDomain`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#relaxing-the-same-origin-restriction:this-4
    x-internal="this"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#relaxing-the-same-origin-restriction:concept-document-origin-2
    x-internal="concept-document-origin"}\'s [effective
    domain](#concept-origin-effective-domain){#relaxing-the-same-origin-restriction:concept-origin-effective-domain-2}.

4.  If `effectiveDomain`{.variable} is null, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#relaxing-the-same-origin-restriction:securityerror-4
    x-internal="securityerror"}
    [`DOMException`{#relaxing-the-same-origin-restriction:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If the given value [is not a registrable domain suffix of and is not
    equal
    to](#is-a-registrable-domain-suffix-of-or-is-equal-to){#relaxing-the-same-origin-restriction:is-a-registrable-domain-suffix-of-or-is-equal-to}
    `effectiveDomain`{.variable}, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#relaxing-the-same-origin-restriction:securityerror-5
    x-internal="securityerror"}
    [`DOMException`{#relaxing-the-same-origin-restriction:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  If the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#relaxing-the-same-origin-restriction:surrounding-agent
    x-internal="surrounding-agent"}\'s [agent
    cluster](https://tc39.es/ecma262/#sec-agent-clusters){#relaxing-the-same-origin-restriction:agent-cluster
    x-internal="agent-cluster"}\'s [is
    origin-keyed](webappapis.html#is-origin-keyed){#relaxing-the-same-origin-restriction:is-origin-keyed}
    is true, then return.

7.  Set
    [this](https://webidl.spec.whatwg.org/#this){#relaxing-the-same-origin-restriction:this-5
    x-internal="this"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#relaxing-the-same-origin-restriction:concept-document-origin-3
    x-internal="concept-document-origin"}\'s
    [domain](#concept-origin-domain){#relaxing-the-same-origin-restriction:concept-origin-domain-2}
    to the result of
    [parsing](https://url.spec.whatwg.org/#concept-host-parser){#relaxing-the-same-origin-restriction:host-parser
    x-internal="host-parser"} the given value.

To determine if a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#relaxing-the-same-origin-restriction:scalar-value-string
x-internal="scalar-value-string"} `hostSuffixString`{.variable} [is a
registrable domain suffix of or is equal
to]{#is-a-registrable-domain-suffix-of-or-is-equal-to .dfn
lt="is a registrable domain suffix of or is equal to|is not a registrable domain suffix of
  and is not equal to" export=""} a
[host](https://url.spec.whatwg.org/#concept-host){#relaxing-the-same-origin-restriction:concept-host
x-internal="concept-host"} `originalHost`{.variable}:

1.  If `hostSuffixString`{.variable} is the empty string, then return
    false.

2.  Let `hostSuffix`{.variable} be the result of
    [parsing](https://url.spec.whatwg.org/#concept-host-parser){#relaxing-the-same-origin-restriction:host-parser-2
    x-internal="host-parser"} `hostSuffixString`{.variable}.

3.  If `hostSuffix`{.variable} is failure, then return false.

4.  If `hostSuffix`{.variable} does not
    [equal](https://url.spec.whatwg.org/#concept-host-equals){#relaxing-the-same-origin-restriction:host-equals
    x-internal="host-equals"} `originalHost`{.variable}, then:

    1.  If `hostSuffix`{.variable} or `originalHost`{.variable} is not a
        [domain](https://url.spec.whatwg.org/#concept-domain){#relaxing-the-same-origin-restriction:concept-domain
        x-internal="concept-domain"}, then return false.

        This excludes
        [hosts](https://url.spec.whatwg.org/#concept-host){#relaxing-the-same-origin-restriction:concept-host-2
        x-internal="concept-host"} that are [IP
        addresses](https://url.spec.whatwg.org/#ip-address){#relaxing-the-same-origin-restriction:ip-address
        x-internal="ip-address"}.

    2.  If `hostSuffix`{.variable}, prefixed by U+002E (.), does not
        match the end of `originalHost`{.variable}, then return false.

    3.  If any of the following are true:

        - `hostSuffix`{.variable}
          [equals](https://url.spec.whatwg.org/#concept-host-equals){#relaxing-the-same-origin-restriction:host-equals-2
          x-internal="host-equals"} `hostSuffix`{.variable}\'s [public
          suffix](https://url.spec.whatwg.org/#host-public-suffix){#relaxing-the-same-origin-restriction:public-suffix
          x-internal="public-suffix"}; or

        - `hostSuffix`{.variable}, prefixed by U+002E (.), matches the
          end of `originalHost`{.variable}\'s [public
          suffix](https://url.spec.whatwg.org/#host-public-suffix){#relaxing-the-same-origin-restriction:public-suffix-2
          x-internal="public-suffix"},

        then return false. [\[URL\]](references.html#refsURL)

    4.  [Assert](https://infra.spec.whatwg.org/#assert){#relaxing-the-same-origin-restriction:assert
        x-internal="assert"}: `originalHost`{.variable}\'s [public
        suffix](https://url.spec.whatwg.org/#host-public-suffix){#relaxing-the-same-origin-restriction:public-suffix-3
        x-internal="public-suffix"}, prefixed by U+002E (.), matches the
        end of `hostSuffix`{.variable}.

5.  Return true.

::: {#example-registrable-domain-suffix .example}
`hostSuffixString`{.variable}

`originalHost`{.variable}

Outcome of [is a registrable domain suffix of or is equal
to](#is-a-registrable-domain-suffix-of-or-is-equal-to){#relaxing-the-same-origin-restriction:is-a-registrable-domain-suffix-of-or-is-equal-to-2}

Notes

\"`0.0.0.0`\"

`0.0.0.0`

✅

\"`0x10203`\"

`0.1.2.3`

✅

\"`[0::1]`\"

`::1`

✅

\"`example.com`\"

`example.com`

✅

\"`example.com`\"

`example.com.`

❌

Trailing dot is significant.

\"`example.com.`\"

`example.com`

❌

\"`example.com`\"

`www.example.com`

✅

\"`com`\"

`example.com`

❌

At the time of writing, `com` is a public suffix.

\"`example`\"

`example`

✅

\"`compute.amazonaws.com`\"

`example.compute.amazonaws.com`

❌

At the time of writing, `*`{.variable}`.compute.amazonaws.com` is a
public suffix.

\"`example.compute.amazonaws.com`\"

`www.example.compute.amazonaws.com`

❌

\"`amazonaws.com`\"

`www.example.compute.amazonaws.com`

❌

\"`amazonaws.com`\"

`test.amazonaws.com`

✅

At the time of writing, `amazonaws.com` is a registrable domain.
:::

#### [7.1.2]{.secno} []{#origin-isolation}Origin-keyed agent clusters[](#origin-keyed-agent-clusters){.self-link}

`window.`[`originAgentCluster`](#dom-originagentcluster){#dom-originagentcluster-dev}

:   Returns true if this
    [`Window`{#origin-keyed-agent-clusters:window}](nav-history-apis.html#window)
    belongs to an [agent
    cluster](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster
    x-internal="agent-cluster"} which is
    [origin](#concept-origin){#origin-keyed-agent-clusters:concept-origin}-[keyed](webappapis.html#agent-cluster-key){#origin-keyed-agent-clusters:agent-cluster-key},
    in the manner described in this section.

A [`Document`{#origin-keyed-agent-clusters:document}](dom.html#document)
delivered over a [secure
context](webappapis.html#secure-context){#origin-keyed-agent-clusters:secure-context}
can request that it be placed in an
[origin](#concept-origin){#origin-keyed-agent-clusters:concept-origin-2}-[keyed](webappapis.html#agent-cluster-key){#origin-keyed-agent-clusters:agent-cluster-key-2}
[agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster-2
x-internal="agent-cluster"}, by using the
\`[`Origin-Agent-Cluster`]{#origin-agent-cluster .dfn
dfn-type="http-header"}\` HTTP response header. This header is a
[structured
header](https://httpwg.org/specs/rfc8941.html){#origin-keyed-agent-clusters:http-structured-header
x-internal="http-structured-header"} whose value must be a
[boolean](https://httpwg.org/specs/rfc8941.html#boolean){#origin-keyed-agent-clusters:http-structured-header-boolean
x-internal="http-structured-header-boolean"}.
[\[STRUCTURED-FIELDS\]](references.html#refsSTRUCTURED-FIELDS)

Per the processing model in the [create and initialize a new `Document`
object](document-lifecycle.html#initialise-the-document-object){#origin-keyed-agent-clusters:initialise-the-document-object},
values that are not the [structured header
boolean](https://httpwg.org/specs/rfc8941.html#boolean){#origin-keyed-agent-clusters:http-structured-header-boolean-2
x-internal="http-structured-header-boolean"} true value (i.e., \``?1`\`)
will be ignored.

The consequences of using this header are that the resulting
[`Document`{#origin-keyed-agent-clusters:document-2}](dom.html#document)\'s
[agent cluster
key](webappapis.html#agent-cluster-key){#origin-keyed-agent-clusters:agent-cluster-key-3}
is its
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#origin-keyed-agent-clusters:concept-document-origin
x-internal="concept-document-origin"}, instead of the [corresponding
site](#obtain-a-site){#origin-keyed-agent-clusters:obtain-a-site}. In
terms of observable effects, this means that attempting to [relax the
same-origin restriction](#relaxing-the-same-origin-restriction) using
[`document.domain`{#origin-keyed-agent-clusters:dom-document-domain}](#dom-document-domain)
will instead do nothing, and it will not be possible to send
[`WebAssembly.Module`{#origin-keyed-agent-clusters:webassembly.module}](https://webassembly.github.io/spec/js-api/#module){x-internal="webassembly.module"}
objects to cross-origin
[`Document`{#origin-keyed-agent-clusters:document-3}](dom.html#document)s
(even if they are [same
site](#same-site){#origin-keyed-agent-clusters:same-site}). Behind the
scenes, this isolation can allow user agents to allocate
implementation-specific resources corresponding to [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster-3
x-internal="agent-cluster"}, such as processes or threads, more
efficiently.

Note that within a [browsing context
group](document-sequences.html#browsing-context-group){#origin-keyed-agent-clusters:browsing-context-group},
the
\`[`Origin-Agent-Cluster`{#origin-keyed-agent-clusters:origin-agent-cluster}](#origin-agent-cluster)\`
header can never cause same-origin
[`Document`{#origin-keyed-agent-clusters:document-4}](dom.html#document)
objects to end up in different [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster-4
x-internal="agent-cluster"}, even if one sends the header and the other
doesn\'t. This is prevented by means of the [historical agent cluster
key
map](document-sequences.html#historical-agent-cluster-key-map){#origin-keyed-agent-clusters:historical-agent-cluster-key-map}.

This means that the
[`originAgentCluster`{#origin-keyed-agent-clusters:dom-originagentcluster}](#dom-originagentcluster)
getter can return false, even if the header is set, if the header was
omitted on a previously-loaded same-origin page in the same [browsing
context
group](document-sequences.html#browsing-context-group){#origin-keyed-agent-clusters:browsing-context-group-2}.
Similarly, it can return true even when the header is not set.

The [`originAgentCluster`]{#dom-originagentcluster .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are to return the [surrounding
agent](https://tc39.es/ecma262/#surrounding-agent){#origin-keyed-agent-clusters:surrounding-agent
x-internal="surrounding-agent"}\'s [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster-5
x-internal="agent-cluster"}\'s [is
origin-keyed](webappapis.html#is-origin-keyed){#origin-keyed-agent-clusters:is-origin-keyed}.

[`Document`{#origin-keyed-agent-clusters:document-5}](dom.html#document)s
with an [opaque
origin](#concept-origin-opaque){#origin-keyed-agent-clusters:concept-origin-opaque}
can be considered unconditionally origin-keyed; for them the header has
no effect, and the
[`originAgentCluster`{#origin-keyed-agent-clusters:dom-originagentcluster-2}](#dom-originagentcluster)
getter will always return true.

Similarly,
[`Document`{#origin-keyed-agent-clusters:document-6}](dom.html#document)s
whose [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#origin-keyed-agent-clusters:agent-cluster-6
x-internal="agent-cluster"}\'s [cross-origin isolation
mode](webappapis.html#agent-cluster-cross-origin-isolation){#origin-keyed-agent-clusters:agent-cluster-cross-origin-isolation}
is not
\"[`none`{#origin-keyed-agent-clusters:cross-origin-isolation-none}](document-sequences.html#cross-origin-isolation-none)\"
are automatically origin-keyed. The
\`[`Origin-Agent-Cluster`{#origin-keyed-agent-clusters:origin-agent-cluster-2}](#origin-agent-cluster)\`
header might be useful as an additional hint to implementations about
resource allocation, since the
\`[`Cross-Origin-Opener-Policy`{#origin-keyed-agent-clusters:cross-origin-opener-policy-2}](#cross-origin-opener-policy-2)\`
and
\`[`Cross-Origin-Embedder-Policy`{#origin-keyed-agent-clusters:cross-origin-embedder-policy}](#cross-origin-embedder-policy)\`
headers used to achieve cross-origin isolation are more about ensuring
that everything in the same address space opts in to being there. But
adding it would have no additional observable effects on author code.

#### [7.1.3]{.secno} Cross-origin opener policies[](#cross-origin-opener-policies){.self-link}

An [opener policy value]{#cross-origin-opener-policy-value .dfn} allows
a document which is navigated to in a [top-level browsing
context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context}
to force the creation of a new [top-level browsing
context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context-2},
and a corresponding
[group](document-sequences.html#tlbc-group){#cross-origin-opener-policies:tlbc-group}.
The possible values are:

\"[`unsafe-none`]{#coop-unsafe-none .dfn}\"
:   This is the (current) default and means that the document will
    occupy the same [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context-3}
    as its predecessor, unless that document specified a different
    [opener
    policy](#cross-origin-opener-policy){#cross-origin-opener-policies:cross-origin-opener-policy}.

\"[`same-origin-allow-popups`]{#coop-same-origin-allow-popups .dfn}\"
:   This forces the creation of a new [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context-4}
    for the document, unless its predecessor specified the same [opener
    policy](#cross-origin-opener-policy){#cross-origin-opener-policies:cross-origin-opener-policy-2}
    and they are [same
    origin](#same-origin){#cross-origin-opener-policies:same-origin}.

\"[`same-origin`]{#coop-same-origin .dfn}\"
:   This behaves the same as
    \"[`same-origin-allow-popups`{#cross-origin-opener-policies:coop-same-origin-allow-popups}](#coop-same-origin-allow-popups)\",
    with the addition that any [auxiliary browsing
    context](document-sequences.html#auxiliary-browsing-context){#cross-origin-opener-policies:auxiliary-browsing-context}
    created needs to contain [same
    origin](#same-origin){#cross-origin-opener-policies:same-origin-2}
    documents that also have the same [opener
    policy](#cross-origin-opener-policy){#cross-origin-opener-policies:cross-origin-opener-policy-3}
    or it will appear closed to the opener.

\"[`same-origin-plus-COEP`]{#coop-same-origin-plus-coep .dfn}\"

:   This behaves the same as
    \"[`same-origin`{#cross-origin-opener-policies:coop-same-origin}](#coop-same-origin)\",
    with the addition that it sets the (new) [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context-5}\'s
    [group](document-sequences.html#tlbc-group){#cross-origin-opener-policies:tlbc-group-2}\'s
    [cross-origin isolation
    mode](document-sequences.html#bcg-cross-origin-isolation){#cross-origin-opener-policies:bcg-cross-origin-isolation}
    to one of
    \"[`logical`{#cross-origin-opener-policies:cross-origin-isolation-logical}](document-sequences.html#cross-origin-isolation-logical)\"
    or
    \"[`concrete`{#cross-origin-opener-policies:cross-origin-isolation-concrete}](document-sequences.html#cross-origin-isolation-concrete)\".

    \"[`same-origin-plus-COEP`{#cross-origin-opener-policies:coop-same-origin-plus-coep}](#coop-same-origin-plus-coep)\"
    cannot be directly set via the
    \`[`Cross-Origin-Opener-Policy`{#cross-origin-opener-policies:cross-origin-opener-policy-2-2}](#cross-origin-opener-policy-2)\`
    header, but results from a combination of setting both
    \`[`Cross-Origin-Opener-Policy`](#cross-origin-opener-policy-2){#cross-origin-opener-policies:cross-origin-opener-policy-2-3}`: `[`same-origin`](#coop-same-origin){#cross-origin-opener-policies:coop-same-origin-2}\`
    and a
    \`[`Cross-Origin-Embedder-Policy`{#cross-origin-opener-policies:cross-origin-embedder-policy}](#cross-origin-embedder-policy)\`
    header whose value is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#cross-origin-opener-policies:compatible-with-cross-origin-isolation}
    together.

\"[`noopener-allow-popups`]{#coop-noopener-allow-popups .dfn}\"

:   This forces the creation of a new [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#cross-origin-opener-policies:top-level-browsing-context-6}
    for the document, regardless of its predecessor.

    ::: note
    While including a
    [`noopener-allow-popups`{#cross-origin-opener-policies:coop-noopener-allow-popups}](#coop-noopener-allow-popups)
    value severs the opener relationship between the document on which
    it is applied and its opener, it does not create a robust security
    boundary between those same-origin documents.

    Other risks from same-origin applications include:

    - Same-origin requests fetching the document\'s content --- could be
      mitigated through Fetch Metadata filtering.
      [\[FETCHMETADATA\]](references.html#refsFETCHMETADATA)

    - Same-origin framing - could be mitigated through
      [`X-Frame-Options`{#cross-origin-opener-policies:x-frame-options}](document-lifecycle.html#x-frame-options)
      or CSP
      [`frame-ancestors`{#cross-origin-opener-policies:frame-ancestors-directive}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"}.

    - JavaScript accessible cookies - can be mitigated by ensuring all
      cookies are `httponly`.

    - [`localStorage`{#cross-origin-opener-policies:dom-localstorage}](webstorage.html#dom-localstorage)
      access to sensitive data.

    - Service worker installation.

    - [Cache API](https://w3c.github.io/ServiceWorker/#cache)
      manipulation or access to sensitive data.
      [\[SW\]](references.html#refsSW)

    - `postMessage` or
      [`BroadcastChannel`{#cross-origin-opener-policies:broadcastchannel}](web-messaging.html#broadcastchannel)
      messaging that exposes sensitive information.

    - Autofill which may not require user interaction for same-origin
      documents.

    Developers using
    [`noopener-allow-popups`{#cross-origin-opener-policies:coop-noopener-allow-popups-2}](#coop-noopener-allow-popups)
    need to make sure that their sensitive applications don\'t rely on
    client-side features accessible to other same-origin documents,
    e.g.,
    [`localStorage`{#cross-origin-opener-policies:dom-localstorage-2}](webstorage.html#dom-localstorage)
    and other client-side storage APIs,
    [`BroadcastChannel`{#cross-origin-opener-policies:broadcastchannel-2}](web-messaging.html#broadcastchannel)
    and related same-origin communication mechanisms. They also need to
    make sure that their server-side endpoints don\'t return sensitive
    data to non-navigation requests, whose response content is
    accessible to same-origin documents.
    :::

An [opener policy]{#cross-origin-opener-policy .dfn} consists of:

- A [value]{#coop-struct-value .dfn}, which is an [opener policy
  value](#cross-origin-opener-policy-value){#cross-origin-opener-policies:cross-origin-opener-policy-value},
  initially
  \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none}](#coop-unsafe-none)\".

- A [reporting endpoint]{#coop-struct-report-endpoint .dfn}, which is
  string or null, initially null.

- A [report-only value]{#coop-struct-report-only-value .dfn}, which is
  an [opener policy
  value](#cross-origin-opener-policy-value){#cross-origin-opener-policies:cross-origin-opener-policy-value-2},
  initially
  \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none-2}](#coop-unsafe-none)\".

- A [report-only reporting endpoint]{#coop-struct-report-only-endpoint
  .dfn}, which is a string or null, initially null.

To [match opener policy values]{#matching-coop .dfn}, given an [opener
policy
value](#cross-origin-opener-policy-value){#cross-origin-opener-policies:cross-origin-opener-policy-value-3}
`documentCOOP`{.variable}, an
[origin](#concept-origin){#cross-origin-opener-policies:concept-origin}
`documentOrigin`{.variable}, an [opener policy
value](#cross-origin-opener-policy-value){#cross-origin-opener-policies:cross-origin-opener-policy-value-4}
`responseCOOP`{.variable}, and an
[origin](#concept-origin){#cross-origin-opener-policies:concept-origin-2}
`responseOrigin`{.variable}:

1.  If `documentCOOP`{.variable} is
    \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none-3}](#coop-unsafe-none)\"
    and `responseCOOP`{.variable} is
    \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none-4}](#coop-unsafe-none)\",
    then return true.

2.  If `documentCOOP`{.variable} is
    \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none-5}](#coop-unsafe-none)\"
    or `responseCOOP`{.variable} is
    \"[`unsafe-none`{#cross-origin-opener-policies:coop-unsafe-none-6}](#coop-unsafe-none)\",
    then return false.

3.  If `documentCOOP`{.variable} is `responseCOOP`{.variable} and
    `documentOrigin`{.variable} is [same
    origin](#same-origin){#cross-origin-opener-policies:same-origin-3}
    with `responseOrigin`{.variable}, then return true.

4.  Return false.

##### [7.1.3.1]{.secno} The headers[](#the-coop-headers){.self-link} {#the-coop-headers}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Headers/Cross-Origin-Opener-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Opener-Policy "The HTTP Cross-Origin-Opener-Policy (COOP) response header allows you to ensure a top-level document does not share a browsing context group with cross-origin documents.")

Support in all current engines.

::: support
[Firefox79+]{.firefox .yes}[Safari15.2+]{.safari
.yes}[Chrome83+]{.chrome .yes}

------------------------------------------------------------------------

[OperaNo]{.opera .no}[Edge83+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
AndroidNo]{.webview_android .no}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
AndroidNo]{.opera_android .no}
:::
::::
:::::

A [`Document`{#the-coop-headers:document}](dom.html#document)\'s
[cross-origin opener
policy](dom.html#concept-document-coop){#the-coop-headers:concept-document-coop}
is derived from the
\`[`Cross-Origin-Opener-Policy`]{#cross-origin-opener-policy-2 .dfn
dfn-type="http-header"}\` and
\`[`Cross-Origin-Opener-Policy-Report-Only`]{#cross-origin-opener-policy-report-only
.dfn dfn-type="http-header"}\` HTTP response headers. These headers are
[structured
headers](https://httpwg.org/specs/rfc8941.html){#the-coop-headers:http-structured-header
x-internal="http-structured-header"} whose value must be a
[token](https://httpwg.org/specs/rfc8941.html#token){#the-coop-headers:http-structured-header-token
x-internal="http-structured-header-token"}.
[\[STRUCTURED-FIELDS\]](references.html#refsSTRUCTURED-FIELDS)

The valid
[token](https://httpwg.org/specs/rfc8941.html#token){#the-coop-headers:http-structured-header-token-2
x-internal="http-structured-header-token"} values are the [opener policy
values](#cross-origin-opener-policy-value){#the-coop-headers:cross-origin-opener-policy-value}.
The token may also have attached
[parameters](https://httpwg.org/specs/rfc8941.html#param){#the-coop-headers:http-structured-header-parameters
x-internal="http-structured-header-parameters"}; of these, the
\"[`report-to`]{#coop-report-to .dfn}\" parameter can have a [valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#the-coop-headers:valid-url-string
x-internal="valid-url-string"} identifying an appropriate reporting
endpoint. [\[REPORTING\]](references.html#refsREPORTING)

Per the processing model described below, user agents will ignore this
header if it contains an invalid value. Likewise, user agents will
ignore this header if the value cannot be parsed as a
[token](https://httpwg.org/specs/rfc8941.html#token){#the-coop-headers:http-structured-header-token-3
x-internal="http-structured-header-token"}.

------------------------------------------------------------------------

To [obtain an opener policy]{#obtain-coop .dfn} given a
[response](https://fetch.spec.whatwg.org/#concept-response){#the-coop-headers:concept-response
x-internal="concept-response"} `response`{.variable} and an
[environment](webappapis.html#environment){#the-coop-headers:environment}
`reservedEnvironment`{.variable}:

1.  Let `policy`{.variable} be a new [opener
    policy](#cross-origin-opener-policy){#the-coop-headers:cross-origin-opener-policy}.

2.  If `reservedEnvironment`{.variable} is a [non-secure
    context](webappapis.html#non-secure-context){#the-coop-headers:non-secure-context},
    then return `policy`{.variable}.

3.  Let `parsedItem`{.variable} be the result of [getting a structured
    field
    value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header){#the-coop-headers:getting-a-structured-field-value
    x-internal="getting-a-structured-field-value"} given
    \`[`Cross-Origin-Opener-Policy`{#the-coop-headers:cross-origin-opener-policy-2}](#cross-origin-opener-policy-2)\`
    and \"`item`\" from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-coop-headers:concept-response-header-list
    x-internal="concept-response-header-list"}.

4.  If `parsedItem`{.variable} is not null, then:

    1.  If `parsedItem`{.variable}\[0\] is
        \"[`same-origin`{#the-coop-headers:coop-same-origin}](#coop-same-origin)\",
        then:

        1.  Let `coep`{.variable} be the result of [obtaining a
            cross-origin embedder
            policy](#obtain-an-embedder-policy){#the-coop-headers:obtain-an-embedder-policy}
            from `response`{.variable} and
            `reservedEnvironment`{.variable}.

        2.  If `coep`{.variable}\'s
            [value](#embedder-policy-value-2){#the-coop-headers:embedder-policy-value-2}
            is [compatible with cross-origin
            isolation](#compatible-with-cross-origin-isolation){#the-coop-headers:compatible-with-cross-origin-isolation},
            then set `policy`{.variable}\'s
            [value](#coop-struct-value){#the-coop-headers:coop-struct-value}
            to
            \"[`same-origin-plus-COEP`{#the-coop-headers:coop-same-origin-plus-coep}](#coop-same-origin-plus-coep)\".

        3.  Otherwise, set `policy`{.variable}\'s
            [value](#coop-struct-value){#the-coop-headers:coop-struct-value-2}
            to
            \"[`same-origin`{#the-coop-headers:coop-same-origin-2}](#coop-same-origin)\".

    2.  If `parsedItem`{.variable}\[0\] is
        \"[`same-origin-allow-popups`{#the-coop-headers:coop-same-origin-allow-popups}](#coop-same-origin-allow-popups)\",
        then set `policy`{.variable}\'s
        [value](#coop-struct-value){#the-coop-headers:coop-struct-value-3}
        to
        \"[`same-origin-allow-popups`{#the-coop-headers:coop-same-origin-allow-popups-2}](#coop-same-origin-allow-popups)\".

    3.  If `parsedItem`{.variable}\[0\] is
        \"[`noopener-allow-popups`{#the-coop-headers:coop-noopener-allow-popups}](#coop-noopener-allow-popups)\",
        then set `policy`{.variable}\'s
        [value](#coop-struct-value){#the-coop-headers:coop-struct-value-4}
        to
        \"[`noopener-allow-popups`{#the-coop-headers:coop-noopener-allow-popups-2}](#coop-noopener-allow-popups)\".

    4.  If
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coop-headers:coop-report-to}](#coop-report-to)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#the-coop-headers:map-exists
        x-internal="map-exists"} and it is a string, then set
        `policy`{.variable}\'s [reporting
        endpoint](#coop-struct-report-endpoint){#the-coop-headers:coop-struct-report-endpoint}
        to
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coop-headers:coop-report-to-2}](#coop-report-to)\"\].

5.  Set `parsedItem`{.variable} to the result of [getting a structured
    field
    value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header){#the-coop-headers:getting-a-structured-field-value-2
    x-internal="getting-a-structured-field-value"} given
    \`[`Cross-Origin-Opener-Policy-Report-Only`{#the-coop-headers:cross-origin-opener-policy-report-only}](#cross-origin-opener-policy-report-only)\`
    and \"`item`\" from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-coop-headers:concept-response-header-list-2
    x-internal="concept-response-header-list"}.

6.  If `parsedItem`{.variable} is not null, then:

    1.  If `parsedItem`{.variable}\[0\] is
        \"[`same-origin`{#the-coop-headers:coop-same-origin-3}](#coop-same-origin)\",
        then:

        1.  Let `coep`{.variable} be the result of [obtaining a
            cross-origin embedder
            policy](#obtain-an-embedder-policy){#the-coop-headers:obtain-an-embedder-policy-2}
            from `response`{.variable} and
            `reservedEnvironment`{.variable}.

        2.  If `coep`{.variable}\'s
            [value](#embedder-policy-value-2){#the-coop-headers:embedder-policy-value-2-2}
            is [compatible with cross-origin
            isolation](#compatible-with-cross-origin-isolation){#the-coop-headers:compatible-with-cross-origin-isolation-2}
            or `coep`{.variable}\'s [report-only
            value](#embedder-policy-report-only-value){#the-coop-headers:embedder-policy-report-only-value}
            is [compatible with cross-origin
            isolation](#compatible-with-cross-origin-isolation){#the-coop-headers:compatible-with-cross-origin-isolation-3},
            then set `policy`{.variable}\'s [report-only
            value](#coop-struct-report-only-value){#the-coop-headers:coop-struct-report-only-value}
            to
            \"[`same-origin-plus-COEP`{#the-coop-headers:coop-same-origin-plus-coep-2}](#coop-same-origin-plus-coep)\".

            Report only COOP also considers report-only COEP to assign
            the special
            \"[`same-origin-plus-COEP`{#the-coop-headers:coop-same-origin-plus-coep-3}](#coop-same-origin-plus-coep)\"
            value. This allows developers more freedom in the order of
            deployment of COOP and COEP.

        3.  Otherwise, set `policy`{.variable}\'s [report-only
            value](#coop-struct-report-only-value){#the-coop-headers:coop-struct-report-only-value-2}
            to
            \"[`same-origin`{#the-coop-headers:coop-same-origin-4}](#coop-same-origin)\".

    2.  If `parsedItem`{.variable}\[0\] is
        \"[`same-origin-allow-popups`{#the-coop-headers:coop-same-origin-allow-popups-3}](#coop-same-origin-allow-popups)\",
        then set `policy`{.variable}\'s [report-only
        value](#coop-struct-report-only-value){#the-coop-headers:coop-struct-report-only-value-3}
        to
        \"[`same-origin-allow-popups`{#the-coop-headers:coop-same-origin-allow-popups-4}](#coop-same-origin-allow-popups)\".

    3.  If
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coop-headers:coop-report-to-3}](#coop-report-to)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#the-coop-headers:map-exists-2
        x-internal="map-exists"} and it is a string, then set
        `policy`{.variable}\'s [report-only reporting
        endpoint](#coop-struct-report-only-endpoint){#the-coop-headers:coop-struct-report-only-endpoint}
        to
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coop-headers:coop-report-to-4}](#coop-report-to)\"\].

7.  Return `policy`{.variable}.

##### [7.1.3.2]{.secno} Browsing context group switches due to opener policy[](#browsing-context-group-switches-due-to-cross-origin-opener-policy){.self-link} {#browsing-context-group-switches-due-to-cross-origin-opener-policy}

To [check if popup COOP values require a browsing context group
switch]{#check-browsing-context-group-switch-coop-value-popup .dfn},
given two
[origins](#concept-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin}
`responseOrigin`{.variable} and
`activeDocumentNavigationOrigin`{.variable}, and two [opener policy
values](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value}
`responseCOOPValue`{.variable} and `activeDocumentCOOPValue`{.variable}:

1.  If `responseCOOPValue`{.variable} is
    \"[`noopener-allow-popups`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-noopener-allow-popups}](#coop-noopener-allow-popups)\",
    then return true.

2.  If all of the following are true:

    - `activeDocumentCOOPValue`{.variable}\'s
      [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-2}
      is
      \"[`same-origin-allow-popups`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-same-origin-allow-popups}](#coop-same-origin-allow-popups)\"
      or
      \"[`noopener-allow-popups`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-noopener-allow-popups-2}](#coop-noopener-allow-popups)\";
      and

    - `responseCOOPValue`{.variable} is
      \"[`unsafe-none`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-unsafe-none}](#coop-unsafe-none)\",

    then return false.

3.  If the result of
    [matching](#matching-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:matching-coop}
    `activeDocumentCOOPValue`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOPValue`{.variable}, and `responseOrigin`{.variable} is
    true, then return false.

4.  Return true.

To [check if COOP values require a browsing context group
switch]{#check-browsing-context-group-switch-coop-value .dfn}, given a
boolean `isInitialAboutBlank`{.variable}, two
[origins](#concept-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin-2}
`responseOrigin`{.variable} and
`activeDocumentNavigationOrigin`{.variable}, and two [opener policy
values](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-3}
`responseCOOPValue`{.variable} and `activeDocumentCOOPValue`{.variable}:

1.  If `isInitialAboutBlank`{.variable} is true, then return the result
    of [checking if popup COOP values requires a browsing context group
    switch](#check-browsing-context-group-switch-coop-value-popup){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-browsing-context-group-switch-coop-value-popup}
    with `responseOrigin`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOPValue`{.variable}, and
    `activeDocumentCOOPValue`{.variable}.

2.  Here we are dealing with a non-popup navigation.

    If the result of
    [matching](#matching-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:matching-coop-2}
    `activeDocumentCOOPValue`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOPValue`{.variable}, and `responseOrigin`{.variable} is
    true, then return false.

3.  Return true.

To [check if enforcing report-only COOP would require a browsing context
group switch]{#check-bcg-switch-navigation-report-only .dfn}, given a
boolean `isInitialAboutBlank`{.variable}, two
[origins](#concept-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin-3}
`responseOrigin`{.variable},
`activeDocumentNavigationOrigin`{.variable}, and two [opener
policies](#cross-origin-opener-policy){#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-opener-policy}
`responseCOOP`{.variable} and `activeDocumentCOOP`{.variable}:

1.  If the result of [checking if COOP values require a browsing context
    group
    switch](#check-browsing-context-group-switch-coop-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-browsing-context-group-switch-coop-value}
    given `isInitialAboutBlank`{.variable}, `responseOrigin`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOP`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-report-only-value},
    and `activeDocumentCOOPReportOnly`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-report-only-value-2}
    is false, then return false.

    Matching report-only policies allows a website to specify the same
    report-only opener policy on all its pages and not receive violation
    reports for navigations between these pages.

2.  If the result of [checking if COOP values require a browsing context
    group
    switch](#check-browsing-context-group-switch-coop-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-browsing-context-group-switch-coop-value-2}
    given `isInitialAboutBlank`{.variable}, `responseOrigin`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOP`{.variable}\'s
    [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-4},
    and `activeDocumentCOOPReportOnly`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-report-only-value-3}
    is true, then return true.

3.  If the result of [checking if COOP values require a browsing context
    group
    switch](#check-browsing-context-group-switch-coop-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-browsing-context-group-switch-coop-value-3}
    given `isInitialAboutBlank`{.variable}, `responseOrigin`{.variable},
    `activeDocumentNavigationOrigin`{.variable},
    `responseCOOP`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-report-only-value-4},
    and `activeDocumentCOOPReportOnly`{.variable}\'s
    [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-5}
    is true, then return true.

4.  Return false.

An [opener policy enforcement result]{#coop-enforcement-result .dfn} is
a
[struct](https://infra.spec.whatwg.org/#struct){#browsing-context-group-switches-due-to-cross-origin-opener-policy:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#browsing-context-group-switches-due-to-cross-origin-opener-policy:struct-item
x-internal="struct-item"}:

- A boolean [needs a browsing context group
  switch]{#coop-enforcement-bcg-switch .dfn}, initially false.

- A boolean [would need a browsing context group switch due to
  report-only]{#coop-enforcement-bcg-switch-report-only .dfn}, initially
  false.

- A
  [URL](https://url.spec.whatwg.org/#concept-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:url
  x-internal="url"} [url]{#coop-enforcement-url .dfn}.

- An
  [origin](#concept-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin-4}
  [origin]{#coop-enforcement-origin .dfn}.

- An [opener
  policy](#cross-origin-opener-policy){#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-opener-policy-2}
  [opener policy]{#coop-enforcement-coop .dfn}.

- A boolean [current context is navigation
  source]{#coop-enforcement-source .dfn}, initially false.

To [enforce a response\'s opener policy]{#coop-enforce .dfn}, given a
[browsing
context](document-sequences.html#browsing-context){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context}
`browsingContext`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:url-2
x-internal="url"} `responseURL`{.variable}, an
[origin](#concept-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin-5}
`responseOrigin`{.variable}, an [opener
policy](#cross-origin-opener-policy){#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-opener-policy-3}
`responseCOOP`{.variable}, an [opener policy enforcement
result](#coop-enforcement-result){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-result}
`currentCOOPEnforcementResult`{.variable}, and a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-request-referrer
x-internal="concept-request-referrer"} `referrer`{.variable}:

1.  Let `newCOOPEnforcementResult`{.variable} be a new [opener policy
    enforcement
    result](#coop-enforcement-result){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-result-2}
    with

    [needs a browsing context group switch](#coop-enforcement-bcg-switch){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch}
    :   `currentCOOPEnforcementResult`{.variable}\'s [needs a browsing
        context group
        switch](#coop-enforcement-bcg-switch){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-2}

    [would need a browsing context group switch due to report-only](#coop-enforcement-bcg-switch-report-only){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-report-only}
    :   `currentCOOPEnforcementResult`{.variable}\'s [would need a
        browsing context group switch due to
        report-only](#coop-enforcement-bcg-switch-report-only){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-report-only-2}

    [url](#coop-enforcement-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-url}
    :   `responseURL`{.variable}

    [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin}
    :   `responseOrigin`{.variable}

    [opener policy](#coop-enforcement-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-coop}
    :   `responseCOOP`{.variable}

    [current context is navigation source](#coop-enforcement-source){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-source}
    :   true

2.  Let `isInitialAboutBlank`{.variable} be
    `browsingContext`{.variable}\'s [active
    document](document-sequences.html#active-document){#browsing-context-group-switches-due-to-cross-origin-opener-policy:active-document}\'s
    [is initial
    `about:blank`](dom.html#is-initial-about:blank){#browsing-context-group-switches-due-to-cross-origin-opener-policy:is-initial-about:blank}.

3.  If `isInitialAboutBlank`{.variable} is true and
    `browsingContext`{.variable}\'s [initial
    URL](document-sequences.html#browsing-context-initial-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context-initial-url}
    is null, set `browsingContext`{.variable}\'s [initial
    URL](document-sequences.html#browsing-context-initial-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context-initial-url-2}
    to `responseURL`{.variable}.

4.  If the result of [checking if COOP values require a browsing context
    group
    switch](#check-browsing-context-group-switch-coop-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-browsing-context-group-switch-coop-value-4}
    given `isInitialAboutBlank`{.variable},
    `currentCOOPEnforcementResult`{.variable}\'s [opener
    policy](#coop-enforcement-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-coop-2}\'s
    [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-6},
    `currentCOOPEnforcementResult`{.variable}\'s
    [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-2},
    `responseCOOP`{.variable}\'s
    [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-7},
    and `responseOrigin`{.variable} is true, then:

    1.  Set `newCOOPEnforcementResult`{.variable}\'s [needs a browsing
        context group
        switch](#coop-enforcement-bcg-switch){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-3}
        to true.

    2.  If `browsingContext`{.variable}\'s
        [group](document-sequences.html#tlbc-group){#browsing-context-group-switches-due-to-cross-origin-opener-policy:tlbc-group}\'s
        [browsing context
        set](document-sequences.html#browsing-context-set){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context-set}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#browsing-context-group-switches-due-to-cross-origin-opener-policy:list-size
        x-internal="list-size"} is greater than 1, then:

        1.  [Queue a violation report for browsing context group switch
            when navigating to a COOP
            response](#coop-violation-navigation-to){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-violation-navigation-to}
            with `responseCOOP`{.variable}, \"`enforce`\",
            `responseURL`{.variable},
            `currentCOOPEnforcementResult`{.variable}\'s
            [url](#coop-enforcement-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-url-2},
            `currentCOOPEnforcementResult`{.variable}\'s
            [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-3},
            `responseOrigin`{.variable}, and `referrer`{.variable}.

        2.  [Queue a violation report for browsing context group switch
            when navigating away from a COOP
            response](#coop-violation-navigation-from){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-violation-navigation-from}
            with `currentCOOPEnforcementResult`{.variable}\'s [opener
            policy](#coop-enforcement-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-coop-3},
            \"`enforce`\", `currentCOOPEnforcementResult`{.variable}\'s
            [url](#coop-enforcement-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-url-3},
            `responseURL`{.variable},
            `currentCOOPEnforcementResult`{.variable}\'s
            [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-4},
            `responseOrigin`{.variable}, and
            `currentCOOPEnforcementResult`{.variable}\'s [current
            context is navigation
            source](#coop-enforcement-source){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-source-2}.

5.  If the result of [checking if enforcing report-only COOP would
    require a browsing context group
    switch](#check-bcg-switch-navigation-report-only){#browsing-context-group-switches-due-to-cross-origin-opener-policy:check-bcg-switch-navigation-report-only}
    given `isInitialAboutBlank`{.variable}, `responseOrigin`{.variable},
    `currentCOOPEnforcementResult`{.variable}\'s
    [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-5},
    `responseCOOP`{.variable}, and
    `currentCOOPEnforcementResult`{.variable}\'s [opener
    policy](#coop-enforcement-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-coop-4},
    is true, then:

    1.  Set `newCOOPEnforcementResult`{.variable}\'s [would need a
        browsing context group switch due to
        report-only](#coop-enforcement-bcg-switch-report-only){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-report-only-3}
        to true.

    2.  If `browsingContext`{.variable}\'s
        [group](document-sequences.html#tlbc-group){#browsing-context-group-switches-due-to-cross-origin-opener-policy:tlbc-group-2}\'s
        [browsing context
        set](document-sequences.html#browsing-context-set){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context-set-2}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#browsing-context-group-switches-due-to-cross-origin-opener-policy:list-size-2
        x-internal="list-size"} is greater than 1, then:

        1.  [Queue a violation report for browsing context group switch
            when navigating to a COOP
            response](#coop-violation-navigation-to){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-violation-navigation-to-2}
            with `responseCOOP`{.variable}, \"`reporting`\",
            `responseURL`{.variable},
            `currentCOOPEnforcementResult`{.variable}\'s
            [url](#coop-enforcement-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-url-4},
            `currentCOOPEnforcementResult`{.variable}\'s
            [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-6},
            `responseOrigin`{.variable}, and `referrer`{.variable}.

        2.  [Queue a violation report for browsing context group switch
            when navigating away from a COOP
            response](#coop-violation-navigation-from){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-violation-navigation-from-2}
            with `currentCOOPEnforcementResult`{.variable}\'s [opener
            policy](#coop-enforcement-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-coop-5},
            \"`reporting`\",
            `currentCOOPEnforcementResult`{.variable}\'s
            [url](#coop-enforcement-url){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-url-5},
            `responseURL`{.variable},
            `currentCOOPEnforcementResult`{.variable}\'s
            [origin](#coop-enforcement-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-origin-7},
            `responseOrigin`{.variable}, and
            `currentCOOPEnforcementResult`{.variable}\'s [current
            context is navigation
            source](#coop-enforcement-source){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-source-3}.

6.  Return `newCOOPEnforcementResult`{.variable}.

To [obtain a browsing context to use for a navigation
response]{#obtain-browsing-context-navigation .dfn}, given [navigation
params](browsing-the-web.html#navigation-params){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params}
`navigationParams`{.variable}:

1.  Let `browsingContext`{.variable} be `navigationParams`{.variable}\'s
    [navigable](browsing-the-web.html#navigation-params-navigable){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-navigable}\'s
    [active browsing
    context](document-sequences.html#nav-bc){#browsing-context-group-switches-due-to-cross-origin-opener-policy:nav-bc}.

2.  If `browsingContext`{.variable} is not a [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#browsing-context-group-switches-due-to-cross-origin-opener-policy:top-level-browsing-context},
    then return `browsingContext`{.variable}.

3.  Let `coopEnforcementResult`{.variable} be
    `navigationParams`{.variable}\'s [COOP enforcement
    result](browsing-the-web.html#navigation-params-coop-enforcement-result){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-coop-enforcement-result}.

4.  Let `swapGroup`{.variable} be `coopEnforcementResult`{.variable}\'s
    [needs a browsing context group
    switch](#coop-enforcement-bcg-switch){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-4}.

5.  Let `sourceOrigin`{.variable} be `browsingContext`{.variable}\'s
    [active
    document](document-sequences.html#active-document){#browsing-context-group-switches-due-to-cross-origin-opener-policy:active-document-2}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-document-origin
    x-internal="concept-document-origin"}.

6.  Let `destinationOrigin`{.variable} be
    `navigationParams`{.variable}\'s
    [origin](browsing-the-web.html#navigation-params-origin){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-origin}.

7.  If `sourceOrigin`{.variable} is not [same
    site](#same-site){#browsing-context-group-switches-due-to-cross-origin-opener-policy:same-site}
    with `destinationOrigin`{.variable}:

    1.  If either of `sourceOrigin`{.variable} or
        `destinationOrigin`{.variable} have a
        [scheme](#concept-origin-scheme){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-origin-scheme}
        that is not an [HTTP(S)
        scheme](https://fetch.spec.whatwg.org/#http-scheme){#browsing-context-group-switches-due-to-cross-origin-opener-policy:http(s)-scheme
        x-internal="http(s)-scheme"} and the user agent considers it
        necessary for `sourceOrigin`{.variable} and
        `destinationOrigin`{.variable} to be isolated from each other
        (for
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#browsing-context-group-switches-due-to-cross-origin-opener-policy:implementation-defined
        x-internal="implementation-defined"} reasons), optionally set
        `swapGroup`{.variable} to true.

        For example, if a user navigates from `about:settings` to
        `https://example.com`, the user agent could force a swap.

        [Issue #10842](https://github.com/whatwg/html/issues/10842)
        tracks settling on an interoperable behavior here, instead of
        letting this be optional.

    2.  If `navigationParams`{.variable}\'s [user
        involvement](browsing-the-web.html#navigation-params-user-involvement){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-user-involvement}
        is
        \"[`browser UI`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\",
        optionally set `swapGroup`{.variable} to true.

        [Issue #6356](https://github.com/whatwg/html/issues/6356) tracks
        settling on an interoperable behavior here, instead of letting
        this be optional.

8.  If `browsingContext`{.variable}\'s
    [group](document-sequences.html#tlbc-group){#browsing-context-group-switches-due-to-cross-origin-opener-policy:tlbc-group-3}\'s
    [browsing context
    set](document-sequences.html#browsing-context-set){#browsing-context-group-switches-due-to-cross-origin-opener-policy:browsing-context-set-3}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#browsing-context-group-switches-due-to-cross-origin-opener-policy:list-size-3
    x-internal="list-size"} is 1, optionally set `swapGroup`{.variable}
    to true.

    Some implementations swap browsing context groups here for
    performance reasons.

    The check for other contexts that could script this one is not
    sufficient to prevent differences in behavior that could affect a
    web page. Even if there are currently no other contexts, the
    destination page could open a window, then if the user navigates
    back, the previous page could expect to be able to script the opened
    window. Doing a swap here would break that use case.

9.  If `swapGroup`{.variable} is false, then:

    1.  If `coopEnforcementResult`{.variable}\'s [would need a browsing
        context group switch due to
        report-only](#coop-enforcement-bcg-switch-report-only){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-enforcement-bcg-switch-report-only-4}
        is true, set `browsingContext`{.variable}\'s [virtual browsing
        context group
        ID](document-sequences.html#virtual-browsing-context-group-id){#browsing-context-group-switches-due-to-cross-origin-opener-policy:virtual-browsing-context-group-id}
        to a new unique identifier.

    2.  Return `browsingContext`{.variable}.

10. Let `newBrowsingContext`{.variable} be the first return value of
    [creating a new top-level browsing context and
    document](document-sequences.html#creating-a-new-top-level-browsing-context){#browsing-context-group-switches-due-to-cross-origin-opener-policy:creating-a-new-top-level-browsing-context}.

    In this case we are going to perform a browsing context group swap.
    `browsingContext`{.variable} will not be used by the new
    [`Document`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:document}](dom.html#document)
    that we are about to
    [create](document-lifecycle.html#initialise-the-document-object){#browsing-context-group-switches-due-to-cross-origin-opener-policy:initialise-the-document-object}.
    If it is not used by other
    [`Document`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:document-2}](dom.html#document)s
    either (such as ones in the back/forward cache), then the user agent
    might [destroy
    it](document-sequences.html#a-browsing-context-is-discarded) at this
    point.

11. Let `navigationCOOP`{.variable} be `navigationParams`{.variable}\'s
    [cross-origin opener
    policy](browsing-the-web.html#navigation-params-coop){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-coop}.

12. If `navigationCOOP`{.variable}\'s
    [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-8}
    is
    \"[`same-origin-plus-COEP`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-same-origin-plus-coep}](#coop-same-origin-plus-coep)\",
    then set `newBrowsingContext`{.variable}\'s
    [group](document-sequences.html#tlbc-group){#browsing-context-group-switches-due-to-cross-origin-opener-policy:tlbc-group-4}\'s
    [cross-origin isolation
    mode](document-sequences.html#bcg-cross-origin-isolation){#browsing-context-group-switches-due-to-cross-origin-opener-policy:bcg-cross-origin-isolation}
    to either
    \"[`logical`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-isolation-logical}](document-sequences.html#cross-origin-isolation-logical)\"
    or
    \"[`concrete`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-isolation-concrete}](document-sequences.html#cross-origin-isolation-concrete)\".
    The choice of which is
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#browsing-context-group-switches-due-to-cross-origin-opener-policy:implementation-defined-2
    x-internal="implementation-defined"}.

    It is difficult on some platforms to provide the security properties
    required by the [cross-origin isolated
    capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#browsing-context-group-switches-due-to-cross-origin-opener-policy:concept-settings-object-cross-origin-isolated-capability}.
    \"[`concrete`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-isolation-concrete-2}](document-sequences.html#cross-origin-isolation-concrete)\"
    grants access to it and
    \"[`logical`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:cross-origin-isolation-logical-2}](document-sequences.html#cross-origin-isolation-logical)\"
    does not.

13. Let `sandboxFlags`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#browsing-context-group-switches-due-to-cross-origin-opener-policy:list-clone
    x-internal="list-clone"} of `navigationParams`{.variable}\'s [final
    sandboxing flag
    set](browsing-the-web.html#navigation-params-sandboxing){#browsing-context-group-switches-due-to-cross-origin-opener-policy:navigation-params-sandboxing}.

14. If `sandboxFlags`{.variable} is not empty, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#browsing-context-group-switches-due-to-cross-origin-opener-policy:assert
        x-internal="assert"}: `navigationCOOP`{.variable}\'s
        [value](#coop-struct-value){#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-struct-value-9}
        is
        \"[`unsafe-none`{#browsing-context-group-switches-due-to-cross-origin-opener-policy:coop-unsafe-none-2}](#coop-unsafe-none)\".

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#browsing-context-group-switches-due-to-cross-origin-opener-policy:assert-2
        x-internal="assert"}: `newBrowsingContext`{.variable}\'s [popup
        sandboxing flag
        set](#popup-sandboxing-flag-set){#browsing-context-group-switches-due-to-cross-origin-opener-policy:popup-sandboxing-flag-set}
        [is
        empty](https://infra.spec.whatwg.org/#list-is-empty){#browsing-context-group-switches-due-to-cross-origin-opener-policy:list-is-empty
        x-internal="list-is-empty"}.

    3.  Set `newBrowsingContext`{.variable}\'s [popup sandboxing flag
        set](#popup-sandboxing-flag-set){#browsing-context-group-switches-due-to-cross-origin-opener-policy:popup-sandboxing-flag-set-2}
        to `sandboxFlags`{.variable}.

15. Return `newBrowsingContext`{.variable}.

##### [7.1.3.3]{.secno} []{#reporting}Reporting[](#coop-reporting){.self-link} {#coop-reporting}

An [accessor-accessed relationship]{#accessor-accessed-relationship
.dfn} is an enum that describes the relationship between two [browsing
contexts](document-sequences.html#browsing-context){#coop-reporting:browsing-context}
between which an access happened. It can take the following values:

[accessor is opener]{#accessor-accessed-opener .dfn}
:   The accessor [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-2}
    or one of its
    [ancestors](document-sequences.html#ancestor-browsing-context){#coop-reporting:ancestor-browsing-context}
    is the [opener browsing
    context](document-sequences.html#opener-browsing-context){#coop-reporting:opener-browsing-context}
    of the accessed [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-3}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc}.

[accessor is openee]{#accessor-accessed-openee .dfn}
:   The accessed [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-4}
    or one of its
    [ancestors](document-sequences.html#ancestor-browsing-context){#coop-reporting:ancestor-browsing-context-2}
    is the [opener browsing
    context](document-sequences.html#opener-browsing-context){#coop-reporting:opener-browsing-context-2}
    of the accessor [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-5}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-2}.

[none]{#accessor-accessed-none .dfn}

:   There is no opener relationship between the accessor [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-6},
    the accessor [browsing
    context](document-sequences.html#browsing-context){#coop-reporting:browsing-context-7},
    or any of their
    [ancestors](document-sequences.html#ancestor-browsing-context){#coop-reporting:ancestor-browsing-context-3}.

To [check if an access between two browsing contexts should be
reported]{#coop-check-access-report .dfn}, given two [browsing
contexts](document-sequences.html#browsing-context){#coop-reporting:browsing-context-8}
`accessor`{.variable} and `accessed`{.variable}, a JavaScript property
name `P`{.variable}, and an [environment settings
object](webappapis.html#environment-settings-object){#coop-reporting:environment-settings-object}
`environment`{.variable}:

1.  If `P`{.variable} is not a [cross-origin accessible window property
    name](nav-history-apis.html#cross-origin-accessible-window-property-name){#coop-reporting:cross-origin-accessible-window-property-name},
    then return.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#coop-reporting:assert
    x-internal="assert"}: `accessor`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document}
    and `accessed`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-2}
    are both [fully
    active](document-sequences.html#fully-active){#coop-reporting:fully-active}.

3.  Let `accessorTopDocument`{.variable} be `accessor`{.variable}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-3}\'s
    [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-3}.

4.  Let `accessorInclusiveAncestorOrigins`{.variable} be the list
    obtained by taking the
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin
    x-internal="concept-document-origin"} of the [active
    document](document-sequences.html#nav-document){#coop-reporting:nav-document}
    of each of `accessor`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-4}\'s
    [inclusive ancestor
    navigables](document-sequences.html#inclusive-ancestor-navigables){#coop-reporting:inclusive-ancestor-navigables}.

5.  Let `accessedTopDocument`{.variable} be `accessed`{.variable}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-4}\'s
    [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-5}.

6.  Let `accessedInclusiveAncestorOrigins`{.variable} be the list
    obtained by taking the
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin-2
    x-internal="concept-document-origin"} of the [active
    document](document-sequences.html#nav-document){#coop-reporting:nav-document-2}
    of each of `accessed`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-6}\'s
    [inclusive ancestor
    navigables](document-sequences.html#inclusive-ancestor-navigables){#coop-reporting:inclusive-ancestor-navigables-2}.

7.  If any of `accessorInclusiveAncestorOrigins`{.variable} are not
    [same origin](#same-origin){#coop-reporting:same-origin} with
    `accessorTopDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin-3
    x-internal="concept-document-origin"}, or if any of
    `accessedInclusiveAncestorOrigins`{.variable} are not [same
    origin](#same-origin){#coop-reporting:same-origin-2} with
    `accessedTopDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin-4
    x-internal="concept-document-origin"}, then return.

    This avoids leaking information about cross-origin iframes to a top
    level frame with opener policy reporting.

8.  If `accessor`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-5}\'s
    [virtual browsing context group
    ID](document-sequences.html#virtual-browsing-context-group-id){#coop-reporting:virtual-browsing-context-group-id}
    is `accessed`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-6}\'s
    [virtual browsing context group
    ID](document-sequences.html#virtual-browsing-context-group-id){#coop-reporting:virtual-browsing-context-group-id-2},
    then return.

9.  Let `accessorAccessedRelationship`{.variable} be a new
    [accessor-accessed
    relationship](#accessor-accessed-relationship){#coop-reporting:accessor-accessed-relationship}
    with value
    [none](#accessor-accessed-none){#coop-reporting:accessor-accessed-none}.

10. If `accessed`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-7}\'s
    [opener browsing
    context](document-sequences.html#opener-browsing-context){#coop-reporting:opener-browsing-context-3}
    is `accessor`{.variable} or is an
    [ancestor](document-sequences.html#ancestor-browsing-context){#coop-reporting:ancestor-browsing-context-4}
    of `accessor`{.variable}, then set
    `accessorAccessedRelationship`{.variable} to [accessor is
    opener](#accessor-accessed-opener){#coop-reporting:accessor-accessed-opener}.

11. If `accessor`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-8}\'s
    [opener browsing
    context](document-sequences.html#opener-browsing-context){#coop-reporting:opener-browsing-context-4}
    is `accessed`{.variable} or is an
    [ancestor](document-sequences.html#ancestor-browsing-context){#coop-reporting:ancestor-browsing-context-5}
    of `accessed`{.variable}, then set
    `accessorAccessedRelationship`{.variable} to [accessor is
    openee](#accessor-accessed-openee){#coop-reporting:accessor-accessed-openee}.

12. [Queue violation reports for
    accesses](#coop-violation-access){#coop-reporting:coop-violation-access},
    given `accessorAccessedRelationship`{.variable},
    `accessorTopDocument`{.variable}\'s [opener
    policy](dom.html#concept-document-coop){#coop-reporting:concept-document-coop},
    `accessedTopDocument`{.variable}\'s [opener
    policy](dom.html#concept-document-coop){#coop-reporting:concept-document-coop-2},
    `accessor`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-7}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#coop-reporting:the-document's-address
    x-internal="the-document's-address"}, `accessed`{.variable}\'s
    [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-8}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#coop-reporting:the-document's-address-2
    x-internal="the-document's-address"}, `accessor`{.variable}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-9}\'s
    [initial
    URL](document-sequences.html#browsing-context-initial-url){#coop-reporting:browsing-context-initial-url},
    `accessed`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-10}\'s
    [initial
    URL](document-sequences.html#browsing-context-initial-url){#coop-reporting:browsing-context-initial-url-2},
    `accessor`{.variable}\'s [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-9}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin-5
    x-internal="concept-document-origin"}, `accessed`{.variable}\'s
    [active
    document](document-sequences.html#active-document){#coop-reporting:active-document-10}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#coop-reporting:concept-document-origin-6
    x-internal="concept-document-origin"}, `accessor`{.variable}\'s
    [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-11}\'s
    [opener origin at
    creation](document-sequences.html#opener-origin-at-creation){#coop-reporting:opener-origin-at-creation},
    `accessed`{.variable}\'s [top-level browsing
    context](document-sequences.html#bc-tlbc){#coop-reporting:bc-tlbc-12}\'s
    [opener origin at
    creation](document-sequences.html#opener-origin-at-creation){#coop-reporting:opener-origin-at-creation-2},
    `accessorTopDocument`{.variable}\'s
    [referrer](dom.html#dom-document-referrer){#coop-reporting:dom-document-referrer},
    `accessedTopDocument`{.variable}\'s
    [referrer](dom.html#dom-document-referrer){#coop-reporting:dom-document-referrer-2},
    `P`{.variable}, and `environment`{.variable}.

To [sanitize a URL to send in a report]{#sanitize-url-report .dfn} given
a [URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url
x-internal="url"} `url`{.variable}:

1.  Let `sanitizedURL`{.variable} be a copy of `url`{.variable}.

2.  [Set the
    username](https://url.spec.whatwg.org/#set-the-username){#coop-reporting:set-the-username
    x-internal="set-the-username"} given `sanitizedURL`{.variable} and
    the empty string.

3.  [Set the
    password](https://url.spec.whatwg.org/#set-the-password){#coop-reporting:set-the-password
    x-internal="set-the-password"} given `sanitizedURL`{.variable} and
    the empty string.

4.  Return the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#coop-reporting:concept-url-serializer
    x-internal="concept-url-serializer"} of `sanitizedURL`{.variable}
    with [*exclude
    fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#coop-reporting:url-serializer-exclude-fragment
    x-internal="url-serializer-exclude-fragment"} set to true.

To [queue a violation report for browsing context group switch when
navigating to a COOP response]{#coop-violation-navigation-to .dfn} given
an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy}
`coop`{.variable}, a string `disposition`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-2
x-internal="url"} `coopURL`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-3
x-internal="url"} `previousResponseURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin}
`coopOrigin`{.variable} and `previousResponseOrigin`{.variable}, and a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#coop-reporting:concept-request-referrer
x-internal="concept-request-referrer"} `referrer`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint}
    is null, return.

2.  Let `coopValue`{.variable} be `coop`{.variable}\'s
    [value](#coop-struct-value){#coop-reporting:coop-struct-value}.

3.  If `disposition`{.variable} is \"`reporting`\", then set
    `coopValue`{.variable} to `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value}.

4.  Let `serializedReferrer`{.variable} be an empty string.

5.  If `referrer`{.variable} is a
    [URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-4
    x-internal="url"}, set `serializedReferrer`{.variable} to the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#coop-reporting:concept-url-serializer-2
    x-internal="concept-url-serializer"} of `referrer`{.variable}.

6.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    `disposition`{.variable}

    effectivePolicy

    `coopValue`{.variable}

    previousResponseURL

    If `coopOrigin`{.variable} and `previousResponseOrigin`{.variable}
    are [same origin](#same-origin){#coop-reporting:same-origin-3} this
    is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report}
    of `previousResponseURL`{.variable}, null otherwise.

    referrer

    `serializedReferrer`{.variable}

    type

    \"`navigation-to-response`\"

7.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-2}
    with `coopURL`{.variable}.

To [queue a violation report for browsing context group switch when
navigating away from a COOP response]{#coop-violation-navigation-from
.dfn} given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-2}
`coop`{.variable}, a string `disposition`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-5
x-internal="url"} `coopURL`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-6
x-internal="url"} `nextResponseURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin-2}
`coopOrigin`{.variable} and `nextResponseOrigin`{.variable}, and a
boolean `isCOOPResponseNavigationSource`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-3}
    is null, return.

2.  Let `coopValue`{.variable} be `coop`{.variable}\'s
    [value](#coop-struct-value){#coop-reporting:coop-struct-value-2}.

3.  If `disposition`{.variable} is \"`reporting`\", then set
    `coopValue`{.variable} to `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-2}.

4.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    `disposition`{.variable}

    effectivePolicy

    `coopValue`{.variable}

    nextResponseURL

    If `coopOrigin`{.variable} and `nextResponseOrigin`{.variable} are
    [same origin](#same-origin){#coop-reporting:same-origin-4} or
    `isCOOPResponseNavigationSource`{.variable} is true, this is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-2}
    of `nextResponseURL`{.variable}, null otherwise.

    type

    \"`navigation-from-response`\"

5.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-2
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-4}
    with `coopURL`{.variable}.

To [queue violation reports for accesses]{#coop-violation-access .dfn},
given an [accessor-accessed
relationship](#accessor-accessed-relationship){#coop-reporting:accessor-accessed-relationship-2}
`accessorAccessedRelationship`{.variable}, two [opener
policies](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-3}
`accessorCOOP`{.variable} and `accessedCOOP`{.variable}, four
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-7
x-internal="url"} `accessorURL`{.variable}, `accessedURL`{.variable},
`accessorInitialURL`{.variable}, `accessedInitialURL`{.variable}, four
[origins](#concept-origin){#coop-reporting:concept-origin-3}
`accessorOrigin`{.variable}, `accessedOrigin`{.variable},
`accessorCreatorOrigin`{.variable} and
`accessedCreatorOrigin`{.variable}, two
[referrers](dom.html#dom-document-referrer){#coop-reporting:dom-document-referrer-3}
`accessorReferrer`{.variable} and `accessedReferrer`{.variable}, a
string `propertyName`{.variable}, and an [environment settings
object](webappapis.html#environment-settings-object){#coop-reporting:environment-settings-object-2}
`environment`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-5}
    is null, return.

2.  Let `coopValue`{.variable} be `coop`{.variable}\'s
    [value](#coop-struct-value){#coop-reporting:coop-struct-value-3}.

3.  If `disposition`{.variable} is \"`reporting`\", then set
    `coopValue`{.variable} to `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-3}.

4.  If `accessorAccessedRelationship`{.variable} is [accessor is
    opener](#accessor-accessed-opener){#coop-reporting:accessor-accessed-opener-2}:

    1.  [Queue a violation report for access to an opened
        window](#coop-violation-access-to-opened){#coop-reporting:coop-violation-access-to-opened},
        given `accessorCOOP`{.variable}, `accessorURL`{.variable},
        `accessedURL`{.variable}, `accessedInitialURL`{.variable},
        `accessorOrigin`{.variable}, `accessedOrigin`{.variable},
        `accessedCreatorOrigin`{.variable}, `propertyName`{.variable},
        and `environment`{.variable}.

    2.  [Queue a violation report for access from the
        opener](#coop-violation-access-from-opener){#coop-reporting:coop-violation-access-from-opener},
        given `accessedCOOP`{.variable}, `accessedURL`{.variable},
        `accessorURL`{.variable}, `accessedOrigin`{.variable},
        `accessorOrigin`{.variable}, `propertyName`{.variable}, and
        `accessedReferrer`{.variable}.

5.  Otherwise, if `accessorAccessedRelationship`{.variable} is [accessor
    is
    openee](#accessor-accessed-openee){#coop-reporting:accessor-accessed-openee-2}:

    1.  [Queue a violation report for access to the
        opener](#coop-violation-access-to-opener){#coop-reporting:coop-violation-access-to-opener},
        given `accessorCOOP`{.variable}, `accessorURL`{.variable},
        `accessedURL`{.variable}, `accessorOrigin`{.variable},
        `accessedOrigin`{.variable}, `propertyName`{.variable},
        `accessorReferrer`{.variable}, and `environment`{.variable}.

    2.  [Queue a violation report for access from an opened
        window](#coop-violation-access-from-opened){#coop-reporting:coop-violation-access-from-opened},
        given `accessedCOOP`{.variable}, `accessedURL`{.variable},
        `accessorURL`{.variable}, `accessorInitialURL`{.variable},
        `accessedOrigin`{.variable}, `accessorOrigin`{.variable},
        `accessorCreatorOrigin`{.variable}, and
        `propertyName`{.variable}.

6.  Otherwise:

    1.  [Queue a violation report for access to another
        window](#coop-violation-access-to-opened){#coop-reporting:coop-violation-access-to-opened-2},
        given `accessorCOOP`{.variable}, `accessorURL`{.variable},
        `accessedURL`{.variable}, `accessorOrigin`{.variable},
        `accessedOrigin`{.variable}, `propertyName`{.variable}, and
        `environment`{.variable}.

    2.  [Queue a violation report for access from another
        window](#coop-violation-access-from-other){#coop-reporting:coop-violation-access-from-other},
        given `accessedCOOP`{.variable}, `accessedURL`{.variable},
        `accessorURL`{.variable}, `accessedOrigin`{.variable},
        `accessorOrigin`{.variable}, and `propertyName`{.variable}.

To [queue a violation report for access to the
opener]{#coop-violation-access-to-opener .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-4}
`coop`{.variable}, two
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-8
x-internal="url"} `coopURL`{.variable} and `openerURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin-4}
`coopOrigin`{.variable} and `openerOrigin`{.variable}, a string
`propertyName`{.variable}, a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#coop-reporting:concept-request-referrer-2
x-internal="concept-request-referrer"} `referrer`{.variable}, and an
[environment settings
object](webappapis.html#environment-settings-object){#coop-reporting:environment-settings-object-3}
`environment`{.variable}:

1.  Let `sourceFile`{.variable}, `lineNumber`{.variable}, and
    `columnNumber`{.variable} be the relevant script URL and problematic
    position which triggered this report.

2.  Let `serializedReferrer`{.variable} be an empty string.

3.  If `referrer`{.variable} is a
    [URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-9
    x-internal="url"}, set `serializedReferrer`{.variable} to the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#coop-reporting:concept-url-serializer-3
    x-internal="concept-url-serializer"} of `referrer`{.variable}.

4.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-4}

    property

    `propertyName`{.variable}

    openerURL

    If `coopOrigin`{.variable} and `openerOrigin`{.variable} are [same
    origin](#same-origin){#coop-reporting:same-origin-5}, this is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-3}
    of `openerURL`{.variable}, null otherwise.

    referrer

    `serializedReferrer`{.variable}

    sourceFile

    `sourceFile`{.variable}

    lineNumber

    `lineNumber`{.variable}

    columnNumber

    `columnNumber`{.variable}

    type

    \"`access-to-opener`\"

5.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-3
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-6}
    with `coopURL`{.variable} and `environment`{.variable}.

To [queue a violation report for access to an opened
window]{#coop-violation-access-to-opened .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-5}
`coop`{.variable}, three
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-10
x-internal="url"} `coopURL`{.variable}, `openedWindowURL`{.variable} and
`initialWindowURL`{.variable}, three
[origins](#concept-origin){#coop-reporting:concept-origin-5}
`coopOrigin`{.variable}, `openedWindowOrigin`{.variable}, and
`openerInitialOrigin`{.variable}, a string `propertyName`{.variable},
and an [environment settings
object](webappapis.html#environment-settings-object){#coop-reporting:environment-settings-object-4}
`environment`{.variable}:

1.  Let `sourceFile`{.variable}, `lineNumber`{.variable}, and
    `columnNumber`{.variable} be the relevant script URL and problematic
    position which triggered this report.

2.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-5}

    property

    `propertyName`{.variable}

    openedWindowURL

    If `coopOrigin`{.variable} and `openedWindowOrigin`{.variable} are
    [same origin](#same-origin){#coop-reporting:same-origin-6}, this is
    the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-4}
    of `openedWindowURL`{.variable}, null otherwise.

    openedWindowInitialURL

    If `coopOrigin`{.variable} and `openerInitialOrigin`{.variable} are
    [same origin](#same-origin){#coop-reporting:same-origin-7}, this is
    the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-5}
    of `initialWindowURL`{.variable}, null otherwise.

    sourceFile

    `sourceFile`{.variable}

    lineNumber

    `lineNumber`{.variable}

    columnNumber

    `columnNumber`{.variable}

    type

    \"`access-to-opener`\"

3.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-4
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-7}
    with `coopURL`{.variable} and `environment`{.variable}.

To [queue a violation report for access to another
window]{#coop-violation-access-to-other .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-6}
`coop`{.variable}, two
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-11
x-internal="url"} `coopURL`{.variable} and `otherURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin-6}
`coopOrigin`{.variable} and `otherOrigin`{.variable}, a string
`propertyName`{.variable}, and an [environment settings
object](webappapis.html#environment-settings-object){#coop-reporting:environment-settings-object-5}
`environment`{.variable}:

1.  Let `sourceFile`{.variable}, `lineNumber`{.variable}, and
    `columnNumber`{.variable} be the relevant script URL and problematic
    position which triggered this report.

2.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-6}

    property

    `propertyName`{.variable}

    otherURL

    If `coopOrigin`{.variable} and `otherOrigin`{.variable} are [same
    origin](#same-origin){#coop-reporting:same-origin-8}, this is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-6}
    of `otherURL`{.variable}, null otherwise.

    sourceFile

    `sourceFile`{.variable}

    lineNumber

    `lineNumber`{.variable}

    columnNumber

    `columnNumber`{.variable}

    type

    \"`access-to-opener`\"

3.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-5
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-8}
    with `coopURL`{.variable} and `environment`{.variable}.

To [queue a violation report for access from the
opener]{#coop-violation-access-from-opener .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-7}
`coop`{.variable}, two
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-12
x-internal="url"} `coopURL`{.variable} and `openerURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin-7}
`coopOrigin`{.variable} and `openerOrigin`{.variable}, a string
`propertyName`{.variable}, and a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#coop-reporting:concept-request-referrer-3
x-internal="concept-request-referrer"} `referrer`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-9}
    is null, return.

2.  Let `serializedReferrer`{.variable} be an empty string.

3.  If `referrer`{.variable} is a
    [URL](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-13
    x-internal="url"}, set `serializedReferrer`{.variable} to the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#coop-reporting:concept-url-serializer-4
    x-internal="concept-url-serializer"} of `referrer`{.variable}.

4.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-7}

    property

    `propertyName`{.variable}

    openerURL

    If `coopOrigin`{.variable} and `openerOrigin`{.variable} are [same
    origin](#same-origin){#coop-reporting:same-origin-9}, this is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-7}
    of `openerURL`{.variable}, null otherwise.

    referrer

    `serializedReferrer`{.variable}

    type

    \"`access-to-opener`\"

5.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-6
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-10}
    with `coopURL`{.variable}.

To [queue a violation report for access from an opened
window]{#coop-violation-access-from-opened .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-8}
`coop`{.variable}, three
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-14
x-internal="url"} `coopURL`{.variable}, `openedWindowURL`{.variable} and
`initialWindowURL`{.variable}, three
[origins](#concept-origin){#coop-reporting:concept-origin-8}
`coopOrigin`{.variable}, `openedWindowOrigin`{.variable}, and
`openerInitialOrigin`{.variable}, and a string
`propertyName`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-11}
    is null, return.

2.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coopValue`{.variable}

    property

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-8}

    openedWindowURL

    If `coopOrigin`{.variable} and `openedWindowOrigin`{.variable} are
    [same origin](#same-origin){#coop-reporting:same-origin-10}, this is
    the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-8}
    of `openedWindowURL`{.variable}, null otherwise.

    openedWindowInitialURL

    If `coopOrigin`{.variable} and `openerInitialOrigin`{.variable} are
    [same origin](#same-origin){#coop-reporting:same-origin-11}, this is
    the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-9}
    of `initialWindowURL`{.variable}, null otherwise.

    type

    \"`access-to-opener`\"

3.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-7
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-12}
    with `coopURL`{.variable}.

To [queue a violation report for access from another
window]{#coop-violation-access-from-other .dfn}, given an [opener
policy](#cross-origin-opener-policy){#coop-reporting:cross-origin-opener-policy-9}
`coop`{.variable}, two
[URLs](https://url.spec.whatwg.org/#concept-url){#coop-reporting:url-15
x-internal="url"} `coopURL`{.variable} and `otherURL`{.variable}, two
[origins](#concept-origin){#coop-reporting:concept-origin-9}
`coopOrigin`{.variable} and `otherOrigin`{.variable}, and a string
`propertyName`{.variable}:

1.  If `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-13}
    is null, return.

2.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    disposition

    \"`reporting`\"

    effectivePolicy

    `coop`{.variable}\'s [report-only
    value](#coop-struct-report-only-value){#coop-reporting:coop-struct-report-only-value-9}

    property

    `propertyName`{.variable}

    otherURL

    If `coopOrigin`{.variable} and `otherOrigin`{.variable} are [same
    origin](#same-origin){#coop-reporting:same-origin-12}, this is the
    [sanitization](#sanitize-url-report){#coop-reporting:sanitize-url-report-10}
    of `otherURL`{.variable}, null otherwise.

    type

    `access-to-opener`

3.  [Queue](https://w3c.github.io/reporting/#queue-report){#coop-reporting:queue-a-report-8
    x-internal="queue-a-report"} `body`{.variable} as \"`coop`\" for
    `coop`{.variable}\'s [reporting
    endpoint](#coop-struct-report-endpoint){#coop-reporting:coop-struct-report-endpoint-14}
    with `coopURL`{.variable}.

#### [7.1.4]{.secno} Cross-origin embedder policies[](#coep){.self-link} {#coep}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Headers/Cross-Origin-Embedder-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Embedder-Policy "The HTTP Cross-Origin-Embedder-Policy (COEP) response header configures embedding cross-origin resources into the document.")

Support in all current engines.

::: support
[Firefox79+]{.firefox .yes}[Safari15.2+]{.safari
.yes}[Chrome83+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge83+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android86+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

An [embedder policy value]{#embedder-policy-value .dfn export=""} is one
of three strings that controls the fetching of cross-origin resources
without explicit permission from resource owners.

\"[`unsafe-none`]{#coep-unsafe-none .dfn dfn-for="embedder policy value" export=""}\"
:   This is the default value. When this value is used, cross-origin
    resources can be fetched without giving explicit permission through
    the [CORS
    protocol](https://fetch.spec.whatwg.org/#http-cors-protocol){#coep:cors-protocol
    x-internal="cors-protocol"} or the
    \`[`Cross-Origin-Resource-Policy`{#coep:cross-origin-resource-policy}](https://fetch.spec.whatwg.org/#http-cross-origin-resource-policy){x-internal="cross-origin-resource-policy"}\`
    header.

\"[`require-corp`]{#coep-require-corp .dfn dfn-for="embedder policy value" export=""}\"
:   When this value is used, fetching cross-origin resources requires
    the server\'s explicit permission through the [CORS
    protocol](https://fetch.spec.whatwg.org/#http-cors-protocol){#coep:cors-protocol-2
    x-internal="cors-protocol"} or the
    \`[`Cross-Origin-Resource-Policy`{#coep:cross-origin-resource-policy-2}](https://fetch.spec.whatwg.org/#http-cross-origin-resource-policy){x-internal="cross-origin-resource-policy"}\`
    header.

\"[`credentialless`]{#coep-credentialless .dfn dfn-for="embedder policy value" export=""}\"

:   When this value is used, fetching cross-origin no-CORS resources
    omits credentials. In exchange, an explicit
    \`[`Cross-Origin-Resource-Policy`{#coep:cross-origin-resource-policy-3}](https://fetch.spec.whatwg.org/#http-cross-origin-resource-policy){x-internal="cross-origin-resource-policy"}\`
    header is not required. Other requests sent with credentials require
    the server\'s explicit permission through the [CORS
    protocol](https://fetch.spec.whatwg.org/#http-cors-protocol){#coep:cors-protocol-3
    x-internal="cors-protocol"} or the
    \`[`Cross-Origin-Resource-Policy`{#coep:cross-origin-resource-policy-4}](https://fetch.spec.whatwg.org/#http-cross-origin-resource-policy){x-internal="cross-origin-resource-policy"}\`
    header.

::: warning
Before supporting
\"[`credentialless`{#coep:coep-credentialless}](#coep-credentialless)\",
implementers are strongly encouraged to support both:

- [Private Network
  Access](https://wicg.github.io/private-network-access/)
- [Opaque Response Blocking](https://github.com/annevk/orb)

Otherwise, it would allow attackers to leverage the client\'s network
position to read non public resources, using the [cross-origin isolated
capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#coep:concept-settings-object-cross-origin-isolated-capability}.
:::

An [embedder policy
value](#embedder-policy-value){#coep:embedder-policy-value} is
[compatible with cross-origin
isolation]{#compatible-with-cross-origin-isolation .dfn} if it is
\"[`credentialless`{#coep:coep-credentialless-2}](#coep-credentialless)\"
or \"[`require-corp`{#coep:coep-require-corp}](#coep-require-corp)\".

An [embedder policy]{#embedder-policy .dfn export=""} consists of:

- A [value]{#embedder-policy-value-2 .dfn dfn-for="embedder policy"
  export=""}, which is an [embedder policy
  value](#embedder-policy-value){#coep:embedder-policy-value-2},
  initially
  \"[`unsafe-none`{#coep:coep-unsafe-none}](#coep-unsafe-none)\".

- A [reporting endpoint]{#embedder-policy-reporting-endpoint .dfn
  dfn-for="embedder policy" export=""} string, initially the empty
  string.

- A [report only value]{#embedder-policy-report-only-value .dfn
  dfn-for="embedder policy" export=""}, which is an [embedder policy
  value](#embedder-policy-value){#coep:embedder-policy-value-3},
  initially
  \"[`unsafe-none`{#coep:coep-unsafe-none-2}](#coep-unsafe-none)\".

- A [report only reporting
  endpoint]{#embedder-policy-report-only-reporting-endpoint .dfn
  dfn-for="embedder
     policy" export=""} string, initially the empty string.

The [\"`coep`\" report type]{#coep-report-type .dfn export=""} is a
[report
type](https://w3c.github.io/reporting/#report-type){#coep:report-type
x-internal="report-type"} whose value is \"`coep`\". It is [visible to
`ReportingObserver`s](https://w3c.github.io/reporting/#visible-to-reportingobservers){#coep:visible-to-reportingobservers
x-internal="visible-to-reportingobservers"}.

##### [7.1.4.1]{.secno} The headers[](#the-coep-headers){.self-link} {#the-coep-headers}

The \`[`Cross-Origin-Embedder-Policy`]{#cross-origin-embedder-policy
.dfn dfn-type="http-header"}\` and
\`[`Cross-Origin-Embedder-Policy-Report-Only`]{#cross-origin-embedder-policy-report-only
.dfn dfn-type="http-header"}\` HTTP response headers allow a server to
declare an [embedder
policy](#embedder-policy){#the-coep-headers:embedder-policy} for an
[environment settings
object](webappapis.html#environment-settings-object){#the-coep-headers:environment-settings-object}.
These headers are [structured
headers](https://httpwg.org/specs/rfc8941.html){#the-coep-headers:http-structured-header
x-internal="http-structured-header"} whose values must be
[token](https://httpwg.org/specs/rfc8941.html#token){#the-coep-headers:http-structured-header-token
x-internal="http-structured-header-token"}.
[\[STRUCTURED-FIELDS\]](references.html#refsSTRUCTURED-FIELDS)

The valid
[token](https://httpwg.org/specs/rfc8941.html#token){#the-coep-headers:http-structured-header-token-2
x-internal="http-structured-header-token"} values are the [embedder
policy
values](#embedder-policy-value){#the-coep-headers:embedder-policy-value}.
The token may also have attached
[parameters](https://httpwg.org/specs/rfc8941.html#param){#the-coep-headers:http-structured-header-parameters
x-internal="http-structured-header-parameters"}; of these, the
\"[`report-to`]{#coep-report-to .dfn}\" parameter can have a [valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#the-coep-headers:valid-url-string
x-internal="valid-url-string"} identifying an appropriate reporting
endpoint. [\[REPORTING\]](references.html#refsREPORTING)

::: note
The [processing
model](#obtain-an-embedder-policy){#the-coep-headers:obtain-an-embedder-policy}
fails open (by defaulting to
\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none}](#coep-unsafe-none)\")
in the presence of a header that cannot be parsed as a token. This
includes inadvertent lists created by combining multiple instances of
the
\`[`Cross-Origin-Embedder-Policy`{#the-coep-headers:cross-origin-embedder-policy}](#cross-origin-embedder-policy)\`
header present in a given response:

\`[`Cross-Origin-Embedder-Policy`{#the-coep-headers:cross-origin-embedder-policy-2}](#cross-origin-embedder-policy)\`

Final [embedder policy
value](#embedder-policy-value){#the-coep-headers:embedder-policy-value-2}

*No header delivered*

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-2}](#coep-unsafe-none)\"

\``require-corp`\`

\"[`require-corp`{#the-coep-headers:coep-require-corp}](#coep-require-corp)\"

\``unknown-value`\`

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-3}](#coep-unsafe-none)\"

\``require-corp, unknown-value`\`

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-4}](#coep-unsafe-none)\"

\``unknown-value, unknown-value`\`

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-5}](#coep-unsafe-none)\"

\``unknown-value, require-corp`\`

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-6}](#coep-unsafe-none)\"

\``require-corp, require-corp`\`

\"[`unsafe-none`{#the-coep-headers:coep-unsafe-none-7}](#coep-unsafe-none)\"

(The same applies to
\`[`Cross-Origin-Embedder-Policy-Report-Only`{#the-coep-headers:cross-origin-embedder-policy-report-only}](#cross-origin-embedder-policy-report-only)\`.)
:::

------------------------------------------------------------------------

To [obtain an embedder policy]{#obtain-an-embedder-policy .dfn
export=""} from a
[response](https://fetch.spec.whatwg.org/#concept-response){#the-coep-headers:concept-response
x-internal="concept-response"} `response`{.variable} and an
[environment](webappapis.html#environment){#the-coep-headers:environment}
`environment`{.variable}:

1.  Let `policy`{.variable} be a new [embedder
    policy](#embedder-policy){#the-coep-headers:embedder-policy-2}.

2.  If `environment`{.variable} is a [non-secure
    context](webappapis.html#non-secure-context){#the-coep-headers:non-secure-context},
    then return `policy`{.variable}.

3.  Let `parsedItem`{.variable} be the result of [getting a structured
    field
    value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header){#the-coep-headers:getting-a-structured-field-value
    x-internal="getting-a-structured-field-value"} with
    \`[`Cross-Origin-Embedder-Policy`{#the-coep-headers:cross-origin-embedder-policy-3}](#cross-origin-embedder-policy)\`
    and \"`item`\" from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-coep-headers:concept-response-header-list
    x-internal="concept-response-header-list"}.

4.  If `parsedItem`{.variable} is non-null and
    `parsedItem`{.variable}\[0\] is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#the-coep-headers:compatible-with-cross-origin-isolation}:

    1.  Set `policy`{.variable}\'s
        [value](#embedder-policy-value-2){#the-coep-headers:embedder-policy-value-2-2}
        to `parsedItem`{.variable}\[0\].

    2.  If
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coep-headers:coep-report-to}](#coep-report-to)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#the-coep-headers:map-exists
        x-internal="map-exists"}, then set `policy`{.variable}\'s
        [endpoint](#embedder-policy-reporting-endpoint){#the-coep-headers:embedder-policy-reporting-endpoint}
        to
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coep-headers:coep-report-to-2}](#coep-report-to)\"\].

5.  Set `parsedItem`{.variable} to the result of [getting a structured
    field
    value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header){#the-coep-headers:getting-a-structured-field-value-2
    x-internal="getting-a-structured-field-value"} with
    \`[`Cross-Origin-Embedder-Policy-Report-Only`{#the-coep-headers:cross-origin-embedder-policy-report-only-2}](#cross-origin-embedder-policy-report-only)\`
    and \"`item`\" from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-coep-headers:concept-response-header-list-2
    x-internal="concept-response-header-list"}.

6.  If `parsedItem`{.variable} is non-null and
    `parsedItem`{.variable}\[0\] is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#the-coep-headers:compatible-with-cross-origin-isolation-2}:

    1.  Set `policy`{.variable}\'s [report only
        value](#embedder-policy-report-only-value){#the-coep-headers:embedder-policy-report-only-value}
        to `parsedItem`{.variable}\[0\].

    2.  If
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coep-headers:coep-report-to-3}](#coep-report-to)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#the-coep-headers:map-exists-2
        x-internal="map-exists"}, then set `policy`{.variable}\'s
        [endpoint](#embedder-policy-report-only-reporting-endpoint){#the-coep-headers:embedder-policy-report-only-reporting-endpoint}
        to
        `parsedItem`{.variable}\[1\]\[\"[`report-to`{#the-coep-headers:coep-report-to-4}](#coep-report-to)\"\].

7.  Return `policy`{.variable}.

##### [7.1.4.2]{.secno} Embedder policy checks[](#embedder-policy-checks){.self-link}

To [check a navigation response\'s adherence to its embedder
policy]{#check-a-navigation-response's-adherence-to-its-embedder-policy
.dfn} given a
[response](https://fetch.spec.whatwg.org/#concept-response){#embedder-policy-checks:concept-response
x-internal="concept-response"} `response`{.variable}, a
[navigable](document-sequences.html#navigable){#embedder-policy-checks:navigable}
`navigable`{.variable}, and an [embedder
policy](#embedder-policy){#embedder-policy-checks:embedder-policy}
`responsePolicy`{.variable}:

1.  If `navigable`{.variable} is not a [child
    navigable](document-sequences.html#child-navigable){#embedder-policy-checks:child-navigable},
    then return true.

2.  Let `parentPolicy`{.variable} be `navigable`{.variable}\'s
    [container
    document](document-sequences.html#nav-container-document){#embedder-policy-checks:nav-container-document}\'s
    [policy
    container](dom.html#concept-document-policy-container){#embedder-policy-checks:concept-document-policy-container}\'s
    [embedder
    policy](#policy-container-embedder-policy){#embedder-policy-checks:policy-container-embedder-policy}.

3.  If `parentPolicy`{.variable}\'s [report-only
    value](#embedder-policy-report-only-value){#embedder-policy-checks:embedder-policy-report-only-value}
    is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation}
    and `responsePolicy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2}
    is not, then [queue a cross-origin embedder policy inheritance
    violation](#queue-a-cross-origin-embedder-policy-inheritance-violation){#embedder-policy-checks:queue-a-cross-origin-embedder-policy-inheritance-violation}
    with `response`{.variable}, \"`navigation`\",
    `parentPolicy`{.variable}\'s [report only reporting
    endpoint](#embedder-policy-report-only-reporting-endpoint){#embedder-policy-checks:embedder-policy-report-only-reporting-endpoint},
    \"`reporting`\", and `navigable`{.variable}\'s [container
    document](document-sequences.html#nav-container-document){#embedder-policy-checks:nav-container-document-2}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#embedder-policy-checks:relevant-settings-object}.

4.  If `parentPolicy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2-2}
    is not [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation-2}
    or `responsePolicy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2-3}
    is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation-3},
    then return true.

5.  [Queue a cross-origin embedder policy inheritance
    violation](#queue-a-cross-origin-embedder-policy-inheritance-violation){#embedder-policy-checks:queue-a-cross-origin-embedder-policy-inheritance-violation-2}
    with `response`{.variable}, \"`navigation`\",
    `parentPolicy`{.variable}\'s [reporting
    endpoint](#embedder-policy-reporting-endpoint){#embedder-policy-checks:embedder-policy-reporting-endpoint},
    \"`enforce`\", and `navigable`{.variable}\'s [container
    document](document-sequences.html#nav-container-document){#embedder-policy-checks:nav-container-document-3}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#embedder-policy-checks:relevant-settings-object-2}.

6.  Return false.

To [check a global object\'s embedder
policy]{#check-a-global-object's-embedder-policy .dfn} given a
[`WorkerGlobalScope`{#embedder-policy-checks:workerglobalscope}](workers.html#workerglobalscope)
`workerGlobalScope`{.variable}, an [environment settings
object](webappapis.html#environment-settings-object){#embedder-policy-checks:environment-settings-object}
`owner`{.variable}, and a
[response](https://fetch.spec.whatwg.org/#concept-response){#embedder-policy-checks:concept-response-2
x-internal="concept-response"} `response`{.variable}:

1.  If `workerGlobalScope`{.variable} is not a
    [`DedicatedWorkerGlobalScope`{#embedder-policy-checks:dedicatedworkerglobalscope}](workers.html#dedicatedworkerglobalscope)
    object, then return true.

2.  Let `policy`{.variable} be `workerGlobalScope`{.variable}\'s
    [embedder
    policy](workers.html#concept-workerglobalscope-embedder-policy){#embedder-policy-checks:concept-workerglobalscope-embedder-policy}.

3.  Let `ownerPolicy`{.variable} be `owner`{.variable}\'s [policy
    container](webappapis.html#concept-settings-object-policy-container){#embedder-policy-checks:concept-settings-object-policy-container}\'s
    [embedder
    policy](#policy-container-embedder-policy){#embedder-policy-checks:policy-container-embedder-policy-2}.

4.  If `ownerPolicy`{.variable}\'s [report-only
    value](#embedder-policy-report-only-value){#embedder-policy-checks:embedder-policy-report-only-value-2}
    is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation-4}
    and `policy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2-4}
    is not, then [queue a cross-origin embedder policy inheritance
    violation](#queue-a-cross-origin-embedder-policy-inheritance-violation){#embedder-policy-checks:queue-a-cross-origin-embedder-policy-inheritance-violation-3}
    with `response`{.variable}, \"`worker initialization`\",
    `ownerPolicy`{.variable}\'s [report only reporting
    endpoint](#embedder-policy-report-only-reporting-endpoint){#embedder-policy-checks:embedder-policy-report-only-reporting-endpoint-2},
    \"`reporting`\", and `owner`{.variable}.

5.  If `ownerPolicy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2-5}
    is not [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation-5}
    or `policy`{.variable}\'s
    [value](#embedder-policy-value-2){#embedder-policy-checks:embedder-policy-value-2-6}
    is [compatible with cross-origin
    isolation](#compatible-with-cross-origin-isolation){#embedder-policy-checks:compatible-with-cross-origin-isolation-6},
    then return true.

6.  [Queue a cross-origin embedder policy inheritance
    violation](#queue-a-cross-origin-embedder-policy-inheritance-violation){#embedder-policy-checks:queue-a-cross-origin-embedder-policy-inheritance-violation-4}
    with `response`{.variable}, \"`worker initialization`\",
    `ownerPolicy`{.variable}\'s [reporting
    endpoint](#embedder-policy-reporting-endpoint){#embedder-policy-checks:embedder-policy-reporting-endpoint-2},
    \"`enforce`\", and `owner`{.variable}.

7.  Return false.

To [queue a cross-origin embedder policy inheritance
violation]{#queue-a-cross-origin-embedder-policy-inheritance-violation
.dfn} given a
[response](https://fetch.spec.whatwg.org/#concept-response){#embedder-policy-checks:concept-response-3
x-internal="concept-response"} `response`{.variable}, a string
`type`{.variable}, a string `endpoint`{.variable}, a string
`disposition`{.variable}, and an [environment settings
object](webappapis.html#environment-settings-object){#embedder-policy-checks:environment-settings-object-2}
`settings`{.variable}:

1.  Let `serialized`{.variable} be the result of [serializing a response
    URL for
    reporting](https://fetch.spec.whatwg.org/#serialize-a-response-url-for-reporting){#embedder-policy-checks:serialize-a-response-url-for-reporting
    x-internal="serialize-a-response-url-for-reporting"} with
    `response`{.variable}.

2.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    type

    `type`{.variable}

    blockedURL

    `serialized`{.variable}

    disposition

    `disposition`{.variable}

3.  [Queue](https://w3c.github.io/reporting/#queue-report){#embedder-policy-checks:queue-a-report
    x-internal="queue-a-report"} `body`{.variable} as the [\"`coep`\"
    report
    type](#coep-report-type){#embedder-policy-checks:coep-report-type}
    for `endpoint`{.variable} on `settings`{.variable}.

#### [7.1.5]{.secno} Sandboxing[](#sandboxing){.self-link}

A [sandboxing flag set]{#sandboxing-flag-set .dfn export=""} is a set of
zero or more of the following flags, which are used to restrict the
abilities that potentially untrusted resources have:

The [sandboxed navigation browsing context flag]{#sandboxed-navigation-browsing-context-flag .dfn export=""}

:   This flag [prevents content from navigating browsing contexts other
    than the sandboxed browsing context
    itself](browsing-the-web.html#sandboxLinks) (or browsing contexts
    further nested inside it), [auxiliary browsing
    contexts](document-sequences.html#auxiliary-browsing-context){#sandboxing:auxiliary-browsing-context}
    (which are protected by the [sandboxed auxiliary navigation browsing
    context
    flag](#sandboxed-auxiliary-navigation-browsing-context-flag){#sandboxing:sandboxed-auxiliary-navigation-browsing-context-flag}
    defined next), and the [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#sandboxing:top-level-browsing-context}
    (which is protected by the [sandboxed top-level navigation without
    user activation browsing context
    flag](#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-without-user-activation-browsing-context-flag}
    and [sandboxed top-level navigation with user activation browsing
    context
    flag](#sandboxed-top-level-navigation-with-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-with-user-activation-browsing-context-flag}
    defined below).

    If the [sandboxed auxiliary navigation browsing context
    flag](#sandboxed-auxiliary-navigation-browsing-context-flag){#sandboxing:sandboxed-auxiliary-navigation-browsing-context-flag-2}
    is not set, then in certain cases the restrictions nonetheless allow
    popups (new [top-level browsing
    contexts](document-sequences.html#top-level-browsing-context){#sandboxing:top-level-browsing-context-2})
    to be opened. These [browsing
    contexts](document-sequences.html#browsing-context){#sandboxing:browsing-context}
    always have [one permitted sandboxed
    navigator]{#one-permitted-sandboxed-navigator .dfn}, set when the
    browsing context is created, which allows the [browsing
    context](document-sequences.html#browsing-context){#sandboxing:browsing-context-2}
    that created them to actually navigate them. (Otherwise, the
    [sandboxed navigation browsing context
    flag](#sandboxed-navigation-browsing-context-flag){#sandboxing:sandboxed-navigation-browsing-context-flag}
    would prevent them from being navigated even if they were opened.)

The [sandboxed auxiliary navigation browsing context flag]{#sandboxed-auxiliary-navigation-browsing-context-flag .dfn export=""}

:   This flag [prevents content from creating new auxiliary browsing
    contexts](document-sequences.html#sandboxWindowOpen), e.g. using the
    [`target`{#sandboxing:attr-hyperlink-target}](links.html#attr-hyperlink-target)
    attribute or the
    [`window.open()`{#sandboxing:dom-open}](nav-history-apis.html#dom-open)
    method.

The [sandboxed top-level navigation without user activation browsing context flag]{#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag .dfn export=""}

:   This flag [prevents content from navigating their top-level browsing
    context](browsing-the-web.html#sandboxLinks) and [prevents content
    from closing their top-level browsing
    context](nav-history-apis.html#sandboxClose). It is consulted only
    when the sandboxed browsing context\'s [active
    window](document-sequences.html#active-window){#sandboxing:active-window}
    does not have [transient
    activation](interaction.html#transient-activation){#sandboxing:transient-activation}.

    When the [sandboxed top-level navigation without user activation
    browsing context
    flag](#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-without-user-activation-browsing-context-flag-2}
    is *not* set, content can navigate its [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#sandboxing:top-level-browsing-context-3},
    but other [browsing
    contexts](document-sequences.html#browsing-context){#sandboxing:browsing-context-3}
    are still protected by the [sandboxed navigation browsing context
    flag](#sandboxed-navigation-browsing-context-flag){#sandboxing:sandboxed-navigation-browsing-context-flag-2}
    and possibly the [sandboxed auxiliary navigation browsing context
    flag](#sandboxed-auxiliary-navigation-browsing-context-flag){#sandboxing:sandboxed-auxiliary-navigation-browsing-context-flag-3}.

The [sandboxed top-level navigation with user activation browsing context flag]{#sandboxed-top-level-navigation-with-user-activation-browsing-context-flag .dfn export=""}

:   This flag [prevents content from navigating their top-level browsing
    context](browsing-the-web.html#sandboxLinks) and [prevents content
    from closing their top-level browsing
    context](nav-history-apis.html#sandboxClose). It is consulted only
    when the sandboxed browsing context\'s [active
    window](document-sequences.html#active-window){#sandboxing:active-window-2}
    has [transient
    activation](interaction.html#transient-activation){#sandboxing:transient-activation-2}.

    As with the [sandboxed top-level navigation without user activation
    browsing context
    flag](#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-without-user-activation-browsing-context-flag-3},
    this flag only affects the [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#sandboxing:top-level-browsing-context-4};
    if it is not set, other [browsing
    contexts](document-sequences.html#browsing-context){#sandboxing:browsing-context-4}
    might still be protected by other flags.

The [sandboxed origin browsing context flag]{#sandboxed-origin-browsing-context-flag .dfn export=""}

:   This flag [forces content into an opaque
    origin](document-sequences.html#sandboxOrigin), thus preventing it
    from accessing other content from the same
    [origin](#concept-origin){#sandboxing:concept-origin}.

    This flag also [prevents script from reading from or writing to the
    `document.cookie` IDL attribute](dom.html#sandboxCookies), and
    blocks access to
    [`localStorage`{#sandboxing:dom-localstorage}](webstorage.html#dom-localstorage).

The [sandboxed forms browsing context flag]{#sandboxed-forms-browsing-context-flag .dfn export=""}

:   This flag [blocks form
    submission](form-control-infrastructure.html#sandboxSubmitBlocked).

The [sandboxed pointer lock browsing context flag]{#sandboxed-pointer-lock-browsing-context-flag .dfn export=""}

:   This flag disables the Pointer Lock API.
    [\[POINTERLOCK\]](references.html#refsPOINTERLOCK)

The [sandboxed scripts browsing context flag]{#sandboxed-scripts-browsing-context-flag .dfn export=""}

:   This flag [blocks script
    execution](webappapis.html#sandboxScriptBlocked).

The [sandboxed automatic features browsing context flag]{#sandboxed-automatic-features-browsing-context-flag .dfn export=""}

:   This flag blocks features that trigger automatically, such as
    [automatically playing a
    video](media.html#attr-media-autoplay){#sandboxing:attr-media-autoplay}
    or [automatically focusing a form
    control](interaction.html#attr-fe-autofocus){#sandboxing:attr-fe-autofocus}.

The [sandboxed `document.domain` browsing context flag]{#sandboxed-document.domain-browsing-context-flag .dfn export=""}

:   This flag prevents content from using the
    [`document.domain`{#sandboxing:dom-document-domain}](#dom-document-domain)
    setter.

The [sandbox propagates to auxiliary browsing contexts flag]{#sandbox-propagates-to-auxiliary-browsing-contexts-flag .dfn export=""}

:   This flag prevents content from escaping the sandbox by ensuring
    that any [auxiliary browsing
    context](document-sequences.html#auxiliary-browsing-context){#sandboxing:auxiliary-browsing-context-2}
    it creates inherits the content\'s [active sandboxing flag
    set](#active-sandboxing-flag-set){#sandboxing:active-sandboxing-flag-set}.

The [sandboxed modals flag]{#sandboxed-modals-flag .dfn export=""}

:   This flag prevents content from using any of the following features
    to produce modal dialogs:

    - [`window.alert()`{#sandboxing:dom-alert}](timers-and-user-prompts.html#dom-alert)
    - [`window.confirm()`{#sandboxing:dom-confirm}](timers-and-user-prompts.html#dom-confirm)
    - [`window.print()`{#sandboxing:dom-print}](timers-and-user-prompts.html#dom-print)
    - [`window.prompt()`{#sandboxing:dom-prompt}](timers-and-user-prompts.html#dom-prompt)
    - the
      [`beforeunload`{#sandboxing:event-beforeunload}](indices.html#event-beforeunload)
      event

The [sandboxed orientation lock browsing context flag]{#sandboxed-orientation-lock-browsing-context-flag .dfn export=""}

:   This flag disables the ability to lock the screen orientation.
    [\[SCREENORIENTATION\]](references.html#refsSCREENORIENTATION)

The [sandboxed presentation browsing context flag]{#sandboxed-presentation-browsing-context-flag .dfn export=""}

:   This flag disables the Presentation API.
    [\[PRESENTATION\]](references.html#refsPRESENTATION)

The [sandboxed downloads browsing context flag]{#sandboxed-downloads-browsing-context-flag .dfn export=""}

:   This flag prevents content from initiating or instantiating
    downloads, whether through [downloading
    hyperlinks](links.html#downloading-hyperlinks){#sandboxing:downloading-hyperlinks}
    or through
    [navigation](browsing-the-web.html#navigation-as-a-download) that
    gets [handled as a
    download](links.html#handle-as-a-download){#sandboxing:handle-as-a-download}.

The [sandboxed custom protocols navigation browsing context flag]{#sandboxed-custom-protocols-navigation-browsing-context-flag .dfn export=""}

:   This flag prevents navigations toward non [fetch
    schemes](https://fetch.spec.whatwg.org/#fetch-scheme){#sandboxing:fetch-scheme
    x-internal="fetch-scheme"} from being [handed off to external
    software](browsing-the-web.html#hand-off-to-external-software){#sandboxing:hand-off-to-external-software}.

When the user agent is to [parse a sandboxing
directive]{#parse-a-sandboxing-directive .dfn export=""}, given a string
`input`{.variable} and a [sandboxing flag
set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set}
`output`{.variable}, it must run the following steps:

1.  [Split `input`{.variable} on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#sandboxing:split-a-string-on-spaces
    x-internal="split-a-string-on-spaces"}, to obtain
    `tokens`{.variable}.

2.  Let `output`{.variable} be empty.

3.  Add the following flags to `output`{.variable}:

    - The [sandboxed navigation browsing context
      flag](#sandboxed-navigation-browsing-context-flag){#sandboxing:sandboxed-navigation-browsing-context-flag-3}.

    - The [sandboxed auxiliary navigation browsing context
      flag](#sandboxed-auxiliary-navigation-browsing-context-flag){#sandboxing:sandboxed-auxiliary-navigation-browsing-context-flag-4},
      unless `tokens`{.variable} contains the
      [`allow-popups`]{#attr-iframe-sandbox-allow-popups .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed top-level navigation without user activation
      browsing context
      flag](#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-without-user-activation-browsing-context-flag-4},
      unless `tokens`{.variable} contains the
      [`allow-top-navigation`]{#attr-iframe-sandbox-allow-top-navigation
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed top-level navigation with user activation browsing
      context
      flag](#sandboxed-top-level-navigation-with-user-activation-browsing-context-flag){#sandboxing:sandboxed-top-level-navigation-with-user-activation-browsing-context-flag-2},
      unless `tokens`{.variable} contains either the
      [`allow-top-navigation-by-user-activation`]{#attr-iframe-sandbox-allow-top-navigation-by-user-activation
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword or
      the
      [`allow-top-navigation`{#sandboxing:attr-iframe-sandbox-allow-top-navigation}](#attr-iframe-sandbox-allow-top-navigation)
      keyword.

      This means that if the
      [`allow-top-navigation`{#sandboxing:attr-iframe-sandbox-allow-top-navigation-2}](#attr-iframe-sandbox-allow-top-navigation)
      is present, the
      [`allow-top-navigation-by-user-activation`{#sandboxing:attr-iframe-sandbox-allow-top-navigation-by-user-activation}](#attr-iframe-sandbox-allow-top-navigation-by-user-activation)
      keyword will have no effect. For this reason, specifying both is a
      document conformance error.

    - The [sandboxed origin browsing context
      flag](#sandboxed-origin-browsing-context-flag){#sandboxing:sandboxed-origin-browsing-context-flag},
      unless the `tokens`{.variable} contains the
      [`allow-same-origin`]{#attr-iframe-sandbox-allow-same-origin .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

      ::: note
      The
      [`allow-same-origin`{#sandboxing:attr-iframe-sandbox-allow-same-origin}](#attr-iframe-sandbox-allow-same-origin)
      keyword is intended for two cases.

      First, it can be used to allow content from the same site to be
      sandboxed to disable scripting, while still allowing access to the
      DOM of the sandboxed content.

      Second, it can be used to embed content from a third-party site,
      sandboxed to prevent that site from opening popups, etc, without
      preventing the embedded page from communicating back to its
      originating site, using the database APIs to store data, etc.
      :::

    - The [sandboxed forms browsing context
      flag](#sandboxed-forms-browsing-context-flag){#sandboxing:sandboxed-forms-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-forms`]{#attr-iframe-sandbox-allow-forms .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed pointer lock browsing context
      flag](#sandboxed-pointer-lock-browsing-context-flag){#sandboxing:sandboxed-pointer-lock-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-pointer-lock`]{#attr-iframe-sandbox-allow-pointer-lock
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed scripts browsing context
      flag](#sandboxed-scripts-browsing-context-flag){#sandboxing:sandboxed-scripts-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-scripts`]{#attr-iframe-sandbox-allow-scripts .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed automatic features browsing context
      flag](#sandboxed-automatic-features-browsing-context-flag){#sandboxing:sandboxed-automatic-features-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-scripts`{#sandboxing:attr-iframe-sandbox-allow-scripts}](#attr-iframe-sandbox-allow-scripts)
      keyword (defined above).

      This flag is relaxed by the same keyword as scripts, because when
      scripts are enabled these features are trivially possible anyway,
      and it would be unfortunate to force authors to use script to do
      them when sandboxed rather than allowing them to use the
      declarative features.

    - The [sandboxed `document.domain` browsing context
      flag](#sandboxed-document.domain-browsing-context-flag){#sandboxing:sandboxed-document.domain-browsing-context-flag}.

    - The [sandbox propagates to auxiliary browsing contexts
      flag](#sandbox-propagates-to-auxiliary-browsing-contexts-flag){#sandboxing:sandbox-propagates-to-auxiliary-browsing-contexts-flag},
      unless `tokens`{.variable} contains the
      [`allow-popups-to-escape-sandbox`]{#attr-iframe-sandbox-allow-popups-to-escape-sandbox
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed modals
      flag](#sandboxed-modals-flag){#sandboxing:sandboxed-modals-flag},
      unless `tokens`{.variable} contains the
      [`allow-modals`]{#attr-iframe-sandbox-allow-modals .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed orientation lock browsing context
      flag](#sandboxed-orientation-lock-browsing-context-flag){#sandboxing:sandboxed-orientation-lock-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-orientation-lock`]{#attr-iframe-sandbox-allow-orientation-lock
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed presentation browsing context
      flag](#sandboxed-presentation-browsing-context-flag){#sandboxing:sandboxed-presentation-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-presentation`]{#attr-iframe-sandbox-allow-presentation
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed downloads browsing context
      flag](#sandboxed-downloads-browsing-context-flag){#sandboxing:sandboxed-downloads-browsing-context-flag},
      unless `tokens`{.variable} contains the
      [`allow-downloads`]{#attr-iframe-sandbox-allow-downloads .dfn
      dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword.

    - The [sandboxed custom protocols navigation browsing context
      flag](#sandboxed-custom-protocols-navigation-browsing-context-flag){#sandboxing:sandboxed-custom-protocols-navigation-browsing-context-flag},
      unless `tokens`{.variable} contains either the
      [`allow-top-navigation-to-custom-protocols`]{#attr-iframe-sandbox-allow-top-navigation-to-custom-protocols
      .dfn dfn-for="iframe/sandbox" dfn-type="attr-value"} keyword, the
      [`allow-popups`{#sandboxing:attr-iframe-sandbox-allow-popups}](#attr-iframe-sandbox-allow-popups)
      keyword, or the
      [`allow-top-navigation`{#sandboxing:attr-iframe-sandbox-allow-top-navigation-3}](#attr-iframe-sandbox-allow-top-navigation)
      keyword.

------------------------------------------------------------------------

Every [top-level browsing
context](document-sequences.html#top-level-browsing-context){#sandboxing:top-level-browsing-context-5}
has a [popup sandboxing flag set]{#popup-sandboxing-flag-set .dfn},
which is a [sandboxing flag
set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-2}. When a
[browsing
context](document-sequences.html#browsing-context){#sandboxing:browsing-context-5}
is created, its [popup sandboxing flag
set](#popup-sandboxing-flag-set){#sandboxing:popup-sandboxing-flag-set}
must be empty. It is populated by [the rules for choosing a
navigable](document-sequences.html#the-rules-for-choosing-a-navigable){#sandboxing:the-rules-for-choosing-a-navigable}
and the [obtain a browsing context to use for a navigation
response](#obtain-browsing-context-navigation){#sandboxing:obtain-browsing-context-navigation}
algorithm.

Every
[`iframe`{#sandboxing:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element has an [`iframe` sandboxing flag
set]{#iframe-sandboxing-flag-set .dfn}, which is a [sandboxing flag
set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-3}. Which
flags in an [`iframe` sandboxing flag
set](#iframe-sandboxing-flag-set){#sandboxing:iframe-sandboxing-flag-set}
are set at any particular time is determined by the
[`iframe`{#sandboxing:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
element\'s
[`sandbox`{#sandboxing:attr-iframe-sandbox}](iframe-embed-object.html#attr-iframe-sandbox)
attribute.

Every [`Document`{#sandboxing:document}](dom.html#document) has an
[active sandboxing flag set]{#active-sandboxing-flag-set .dfn
dfn-for="Document" export=""}, which is a [sandboxing flag
set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-4}. When the
[`Document`{#sandboxing:document-2}](dom.html#document) is created, its
[active sandboxing flag
set](#active-sandboxing-flag-set){#sandboxing:active-sandboxing-flag-set-2}
must be empty. It is populated by the [navigation
algorithm](browsing-the-web.html#navigate){#sandboxing:navigate}.

Every [CSP
list](https://w3c.github.io/webappsec-csp/#csp-list){#sandboxing:concept-csp-list
x-internal="concept-csp-list"} `cspList`{.variable} has [CSP-derived
sandboxing flags]{#csp-derived-sandboxing-flags .dfn}, which is a
[sandboxing flag
set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-5}. It is the
return value of the following algorithm:

1.  Let `directives`{.variable} be an empty [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#sandboxing:set
    x-internal="set"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#sandboxing:list-iterate
    x-internal="list-iterate"} policy in `cspList`{.variable}:

    1.  If `policy`{.variable}\'s
        [disposition](https://w3c.github.io/webappsec-csp/#policy-disposition){#sandboxing:csp-disposition
        x-internal="csp-disposition"} is not \"`enforce`\", then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#sandboxing:continue
        x-internal="continue"}.

    2.  If `policy`{.variable}\'s [directive
        set](https://w3c.github.io/webappsec-csp/#policy-directive-set){#sandboxing:csp-directive-set
        x-internal="csp-directive-set"}
        [contains](https://infra.spec.whatwg.org/#list-contain){#sandboxing:list-contains
        x-internal="list-contains"} a
        [directive](https://w3c.github.io/webappsec-csp/#directives){#sandboxing:content-security-policy-directive
        x-internal="content-security-policy-directive"} whose name is
        \"[`sandbox`{#sandboxing:sandbox-directive}](https://w3c.github.io/webappsec-csp/#sandbox){x-internal="sandbox-directive"}\",
        then
        [append](https://infra.spec.whatwg.org/#list-append){#sandboxing:list-append
        x-internal="list-append"} that directive to
        `directives`{.variable}.

3.  If `directives`{.variable} is empty, then return an empty
    [sandboxing flag
    set](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-6}.

4.  Let `directive`{.variable} be
    `directives`{.variable}\[`directives`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#sandboxing:list-size
    x-internal="list-size"} − 1\].

5.  Return the result of [parsing the sandboxing
    directive](#parse-a-sandboxing-directive){#sandboxing:parse-a-sandboxing-directive}
    `directive`{.variable}.

------------------------------------------------------------------------

To [determine the creation sandboxing
flags]{#determining-the-creation-sandboxing-flags .dfn} for a [browsing
context](document-sequences.html#concept-document-bc){#sandboxing:concept-document-bc}
`browsing context`{.variable}, given null or an element
`embedder`{.variable}, return the
[union](https://infra.spec.whatwg.org/#set-union){#sandboxing:set-union
x-internal="set-union"} of the flags that are present in the following
[sandboxing flag
sets](#sandboxing-flag-set){#sandboxing:sandboxing-flag-set-7}:

- If `embedder`{.variable} is null, then: the flags set on
  `browsing context`{.variable}\'s [popup sandboxing flag
  set](#popup-sandboxing-flag-set){#sandboxing:popup-sandboxing-flag-set-2}.

- If `embedder`{.variable} is an element, then: the flags set on
  `embedder`{.variable}\'s [`iframe` sandboxing flag
  set](#iframe-sandboxing-flag-set){#sandboxing:iframe-sandboxing-flag-set-2}.

- If `embedder`{.variable} is an element, then: the flags set on
  `embedder`{.variable}\'s [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#sandboxing:node-document
  x-internal="node-document"}\'s [active sandboxing flag
  set](#active-sandboxing-flag-set){#sandboxing:active-sandboxing-flag-set-3}.

#### [7.1.6]{.secno} Policy containers[](#policy-containers){.self-link}

A [policy container]{#policy-container .dfn export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#policy-containers:struct
x-internal="struct"} containing policies that apply to a
[`Document`{#policy-containers:document}](dom.html#document), a
[`WorkerGlobalScope`{#policy-containers:workerglobalscope}](workers.html#workerglobalscope),
or a
[`WorkletGlobalScope`{#policy-containers:workletglobalscope}](worklets.html#workletglobalscope).
It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#policy-containers:struct-item
x-internal="struct-item"}:

- A [CSP list]{#policy-container-csp-list .dfn
  dfn-for="policy container" export=""}, which is a [CSP
  list](https://w3c.github.io/webappsec-csp/#csp-list){#policy-containers:concept-csp-list
  x-internal="concept-csp-list"}. It is initially empty.

- An [embedder policy]{#policy-container-embedder-policy .dfn
  dfn-for="policy container" export=""}, which is an [embedder
  policy](#embedder-policy){#policy-containers:embedder-policy}. It is
  initially a new [embedder
  policy](#embedder-policy){#policy-containers:embedder-policy-2}.

- A [referrer policy]{#policy-container-referrer-policy .dfn
  dfn-for="policy container" export=""}, which is a [referrer
  policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#policy-containers:referrer-policy
  x-internal="referrer-policy"}. It is initially the [default referrer
  policy](https://w3c.github.io/webappsec-referrer-policy/#default-referrer-policy){#policy-containers:default-referrer-policy
  x-internal="default-referrer-policy"}.

Move other policies into the policy container.

To [clone a policy container]{#clone-a-policy-container .dfn export=""}
given a [policy
container](#policy-container){#policy-containers:policy-container}
`policyContainer`{.variable}:

1.  Let `clone`{.variable} be a new [policy
    container](#policy-container){#policy-containers:policy-container-2}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#policy-containers:list-iterate
    x-internal="list-iterate"} `policy`{.variable} in
    `policyContainer`{.variable}\'s [CSP
    list](#policy-container-csp-list){#policy-containers:policy-container-csp-list},
    [append](https://infra.spec.whatwg.org/#list-append){#policy-containers:list-append
    x-internal="list-append"} a copy of `policy`{.variable} into
    `clone`{.variable}\'s [CSP
    list](#policy-container-csp-list){#policy-containers:policy-container-csp-list-2}.

3.  Set `clone`{.variable}\'s [embedder
    policy](#policy-container-embedder-policy){#policy-containers:policy-container-embedder-policy}
    to a copy of `policyContainer`{.variable}\'s [embedder
    policy](#policy-container-embedder-policy){#policy-containers:policy-container-embedder-policy-2}.

4.  Set `clone`{.variable}\'s [referrer
    policy](#policy-container-referrer-policy){#policy-containers:policy-container-referrer-policy}
    to `policyContainer`{.variable}\'s [referrer
    policy](#policy-container-referrer-policy){#policy-containers:policy-container-referrer-policy-2}.

5.  Return `clone`{.variable}.

To determine whether a
[URL](https://url.spec.whatwg.org/#concept-url){#policy-containers:url
x-internal="url"} `url`{.variable} [requires storing the policy
container in history]{#requires-storing-the-policy-container-in-history
.dfn}:

1.  If `url`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#policy-containers:concept-url-scheme
    x-internal="concept-url-scheme"} is \"`blob`\", then return false.

2.  If `url`{.variable} [is
    local](https://fetch.spec.whatwg.org/#is-local){#policy-containers:is-local
    x-internal="is-local"}, then return true.

3.  Return false.

To [create a policy container from a fetch
response]{#creating-a-policy-container-from-a-fetch-response .dfn
lt="creating a policy container from a fetch response" export=""} given
a
[response](https://fetch.spec.whatwg.org/#concept-response){#policy-containers:concept-response
x-internal="concept-response"} `response`{.variable} and an
[environment](webappapis.html#environment){#policy-containers:environment}-or-null
`environment`{.variable}:

1.  If `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#policy-containers:concept-response-url
    x-internal="concept-response-url"}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#policy-containers:concept-url-scheme-2
    x-internal="concept-url-scheme"} is \"`blob`\", then return a
    [clone](#clone-a-policy-container){#policy-containers:clone-a-policy-container}
    of `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#policy-containers:concept-response-url-2
    x-internal="concept-response-url"}\'s [blob URL
    entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#policy-containers:concept-url-blob-entry
    x-internal="concept-url-blob-entry"}\'s
    [environment](https://w3c.github.io/FileAPI/#blob-url-entry-environment){#policy-containers:blob-url-entry-environment
    x-internal="blob-url-entry-environment"}\'s [policy
    container](#policy-container){#policy-containers:policy-container-3}.

2.  Let `result`{.variable} be a new [policy
    container](#policy-container){#policy-containers:policy-container-4}.

3.  Set `result`{.variable}\'s [CSP
    list](#policy-container-csp-list){#policy-containers:policy-container-csp-list-3}
    to the result of [parsing a response\'s Content Security
    Policies](https://w3c.github.io/webappsec-csp/#parse-response-csp){#policy-containers:parse-response-csp
    x-internal="parse-response-csp"} given `response`{.variable}.

4.  If `environment`{.variable} is non-null, then set
    `result`{.variable}\'s [embedder
    policy](#policy-container-embedder-policy){#policy-containers:policy-container-embedder-policy-3}
    to the result of [obtaining an embedder
    policy](#obtain-an-embedder-policy){#policy-containers:obtain-an-embedder-policy}
    given `response`{.variable} and `environment`{.variable}. Otherwise,
    set it to
    \"[`unsafe-none`{#policy-containers:coep-unsafe-none}](#coep-unsafe-none)\".

5.  Set `result`{.variable}\'s [referrer
    policy](#policy-container-referrer-policy){#policy-containers:policy-container-referrer-policy-3}
    to the result of [parsing the \``Referrer-Policy`\`
    header](https://w3c.github.io/webappsec-referrer-policy/#parse-referrer-policy-from-header){#policy-containers:parse-referrer-policy-header
    x-internal="parse-referrer-policy-header"} given
    `response`{.variable}.
    [\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

6.  Return `result`{.variable}.

To [determine navigation params policy
container]{#determining-navigation-params-policy-container .dfn} given a
[URL](https://url.spec.whatwg.org/#concept-url){#policy-containers:url-2
x-internal="url"} `responseURL`{.variable} and four [policy
container](#policy-container){#policy-containers:policy-container-5}-or-nulls
`historyPolicyContainer`{.variable},
`initiatorPolicyContainer`{.variable},
`parentPolicyContainer`{.variable}, and
`responsePolicyContainer`{.variable}:

1.  If `historyPolicyContainer`{.variable} is not null, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#policy-containers:assert
        x-internal="assert"}: `responseURL`{.variable} [requires storing
        the policy container in
        history](#requires-storing-the-policy-container-in-history){#policy-containers:requires-storing-the-policy-container-in-history}.

    2.  Return a
        [clone](#clone-a-policy-container){#policy-containers:clone-a-policy-container-2}
        of `historyPolicyContainer`{.variable}.

2.  If `responseURL`{.variable} is
    [`about:srcdoc`{#policy-containers:about:srcdoc}](urls-and-fetching.html#about:srcdoc),
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#policy-containers:assert-2
        x-internal="assert"}: `parentPolicyContainer`{.variable} is not
        null.

    2.  Return a
        [clone](#clone-a-policy-container){#policy-containers:clone-a-policy-container-3}
        of `parentPolicyContainer`{.variable}.

3.  If `responseURL`{.variable} [is
    local](https://fetch.spec.whatwg.org/#is-local){#policy-containers:is-local-2
    x-internal="is-local"} and `initiatorPolicyContainer`{.variable} is
    not null, then return a
    [clone](#clone-a-policy-container){#policy-containers:clone-a-policy-container-4}
    of `initiatorPolicyContainer`{.variable}.

4.  If `responsePolicyContainer`{.variable} is not null, then return
    `responsePolicyContainer`{.variable}.

5.  Return a new [policy
    container](#policy-container){#policy-containers:policy-container-6}.

To [initialize a worker global scope\'s policy
container]{#initialize-worker-policy-container .dfn} given a
[`WorkerGlobalScope`{#policy-containers:workerglobalscope-2}](workers.html#workerglobalscope)
`workerGlobalScope`{.variable}, a
[response](https://fetch.spec.whatwg.org/#concept-response){#policy-containers:concept-response-2
x-internal="concept-response"} `response`{.variable}, and an
[environment](webappapis.html#environment){#policy-containers:environment-2}
`environment`{.variable}:

1.  If `workerGlobalScope`{.variable}\'s
    [url](workers.html#concept-workerglobalscope-url){#policy-containers:concept-workerglobalscope-url}
    [is
    local](https://fetch.spec.whatwg.org/#is-local){#policy-containers:is-local-3
    x-internal="is-local"} but its
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#policy-containers:concept-url-scheme-3
    x-internal="concept-url-scheme"} is not \"`blob`\":

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#policy-containers:assert-3
        x-internal="assert"}: `workerGlobalScope`{.variable}\'s [owner
        set](workers.html#concept-WorkerGlobalScope-owner-set){#policy-containers:concept-WorkerGlobalScope-owner-set}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#policy-containers:list-size
        x-internal="list-size"} is 1.

    2.  Set `workerGlobalScope`{.variable}\'s [policy
        container](workers.html#concept-workerglobalscope-policy-container){#policy-containers:concept-workerglobalscope-policy-container}
        to a
        [clone](#clone-a-policy-container){#policy-containers:clone-a-policy-container-5}
        of `workerGlobalScope`{.variable}\'s [owner
        set](workers.html#concept-WorkerGlobalScope-owner-set){#policy-containers:concept-WorkerGlobalScope-owner-set-2}\[0\]\'s
        [relevant settings
        object](webappapis.html#relevant-settings-object){#policy-containers:relevant-settings-object}\'s
        [policy
        container](webappapis.html#concept-settings-object-policy-container){#policy-containers:concept-settings-object-policy-container}.

2.  Otherwise, set `workerGlobalScope`{.variable}\'s [policy
    container](workers.html#concept-workerglobalscope-policy-container){#policy-containers:concept-workerglobalscope-policy-container-2}
    to the result of [creating a policy container from a fetch
    response](#creating-a-policy-container-from-a-fetch-response){#policy-containers:creating-a-policy-container-from-a-fetch-response}
    given `response`{.variable} and `environment`{.variable}.

[← 6.12 The popover attribute](popover.html) --- [Table of Contents](./)
--- [7.2 APIs related to navigation and session history
→](nav-history-apis.html)
