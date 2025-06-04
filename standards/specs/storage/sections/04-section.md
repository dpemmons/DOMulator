## [4. ]{.secno}[Model]{.content}[](#model){.self-link} {#model .heading .settled level="4"}

Standards defining local or session storage APIs will define a [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint link-type="dfn"}
and
[register](#registered-storage-endpoints){#ref-for-registered-storage-endpoints
link-type="dfn"} it by changing this standard. They will invoke either
the [obtain a local storage bottle
map](#obtain-a-local-storage-bottle-map){#ref-for-obtain-a-local-storage-bottle-map
link-type="dfn"} or the [obtain a session storage bottle
map](#obtain-a-session-storage-bottle-map){#ref-for-obtain-a-session-storage-bottle-map
link-type="dfn"} algorithm, which will give them:

- Failure, which might mean the API has to throw or otherwise indicate
  there is no storage available for that [environment settings
  object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object
  link-type="dfn"}.

- A [storage proxy map](#storage-proxy-map){#ref-for-storage-proxy-map
  link-type="dfn"} that operates analogously to a
  [map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map
  link-type="dfn"}, which can be used to store data in a manner that
  suits the API. This standard takes care of isolating that data from
  other APIs, [storage keys](#storage-key){#ref-for-storage-key
  link-type="dfn"}, and [storage
  types](#storage-type){#ref-for-storage-type link-type="dfn"}.

If you are defining a standard for such an API, consider filing an issue
against this standard for assistance and review.

![Storage model diagram (described in the next
paragraph).](assets/model-diagram.svg){crossorigin="" height="815"
width="434"}

To isolate this data this standard defines a [storage
shed](#storage-shed){#ref-for-storage-shed link-type="dfn"} which
segments [storage shelves](#storage-shelf){#ref-for-storage-shelf
link-type="dfn"} by a [storage key](#storage-key){#ref-for-storage-key①
link-type="dfn"}. A [storage
shelf](#storage-shelf){#ref-for-storage-shelf① link-type="dfn"} in turn
consists of a [storage bucket](#storage-bucket){#ref-for-storage-bucket
link-type="dfn"} and will likely consist of multiple [storage
buckets](#storage-bucket){#ref-for-storage-bucket① link-type="dfn"} in
the future to allow for different storage policies. And lastly, a
[storage bucket](#storage-bucket){#ref-for-storage-bucket②
link-type="dfn"} consists of [storage
bottles](#storage-bottle){#ref-for-storage-bottle link-type="dfn"}, one
for each [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint①
link-type="dfn"}.

### [4.1. ]{.secno}[Storage endpoints]{.content}[](#storage-endpoints){.self-link} {#storage-endpoints .heading .settled level="4.1"}

A [storage endpoint]{#storage-endpoint .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [local](#local-storage){#ref-for-local-storage
link-type="dfn"} or [session
storage](#session-storage){#ref-for-session-storage link-type="dfn"} API
that uses the infrastructure defined by this standard, most notably
[storage bottles](#storage-bottle){#ref-for-storage-bottle①
link-type="dfn"}, to keep track of its storage needs.

A [storage endpoint](#storage-endpoint){#ref-for-storage-endpoint②
link-type="dfn"} has an [identifier]{#storage-endpoint-identifier .dfn
.dfn-paneled dfn-for="storage endpoint" dfn-type="dfn" noexport=""},
which is a [storage
identifier](#storage-identifier){#ref-for-storage-identifier
link-type="dfn"}.

A [storage endpoint](#storage-endpoint){#ref-for-storage-endpoint③
link-type="dfn"} also has [types]{#storage-endpoint-types .dfn
.dfn-paneled dfn-for="storage endpoint" dfn-type="dfn" noexport=""},
which is a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set
link-type="dfn"} of [storage
types](#storage-type){#ref-for-storage-type① link-type="dfn"}.

A [storage endpoint](#storage-endpoint){#ref-for-storage-endpoint④
link-type="dfn"} also has a [quota]{#storage-endpoint-quota .dfn
.dfn-paneled dfn-for="storage endpoint" dfn-type="dfn" noexport=""},
which is null or a number representing a recommended
[quota](#storage-bottle-quota){#ref-for-storage-bottle-quota
link-type="dfn"} (in bytes) for each [storage
bottle](#storage-bottle){#ref-for-storage-bottle② link-type="dfn"}
corresponding to this [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint⑤
link-type="dfn"}.

A [storage identifier]{#storage-identifier .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string
link-type="dfn"}.

A [storage type]{#storage-type .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is \"`local`\" or \"`session`\".

------------------------------------------------------------------------

The [registered storage endpoints]{#registered-storage-endpoints .dfn
.dfn-paneled dfn-type="dfn" noexport=""} are a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set①
link-type="dfn"} of [storage
endpoints](#storage-endpoint){#ref-for-storage-endpoint⑥
link-type="dfn"} defined by the following table:

[Identifier](#storage-endpoint-identifier){#ref-for-storage-endpoint-identifier
link-type="dfn"}

[Type](#storage-endpoint-types){#ref-for-storage-endpoint-types
link-type="dfn"}

[Quota](#storage-endpoint-quota){#ref-for-storage-endpoint-quota
link-type="dfn"}

\"`caches`\"

« \"`local`\" »

null

\"`indexedDB`\"

« \"`local`\" »

null

\"`localStorage`\"

« \"`local`\" »

5 × 2^20^ (i.e., 5 mebibytes)

\"`serviceWorkerRegistrations`\"

« \"`local`\" »

null

\"`sessionStorage`\"

« \"`session`\" »

5 × 2^20^ (i.e., 5 mebibytes)

As mentioned, standards can use these [storage
identifiers](#storage-identifier){#ref-for-storage-identifier①
link-type="dfn"} with [obtain a local storage bottle
map](#obtain-a-local-storage-bottle-map){#ref-for-obtain-a-local-storage-bottle-map①
link-type="dfn"} and [obtain a session storage bottle
map](#obtain-a-session-storage-bottle-map){#ref-for-obtain-a-session-storage-bottle-map①
link-type="dfn"}. It is anticipated that some APIs will be applicable to
both [storage types](#storage-type){#ref-for-storage-type②
link-type="dfn"} going forward.

### [4.2. ]{.secno}[Storage keys]{.content}[](#storage-keys){.self-link} {#storage-keys .heading .settled level="4.2"}

A [storage key]{#storage-key .dfn .dfn-paneled dfn-type="dfn" export=""}
is a [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple
link-type="dfn"} consisting of an [origin]{#storage-key-origin .dfn
.dfn-paneled dfn-for="storage key" dfn-type="dfn" noexport=""} (an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin①
link-type="dfn"}).
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

This is expected to change; see [Client-Side Storage
Partitioning](https://privacycg.github.io/storage-partitioning/).

::: {.algorithm algorithm="obtain a storage key"}
To [obtain a storage key]{#obtain-a-storage-key .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an
[environment](https://html.spec.whatwg.org/multipage/webappapis.html#environment){#ref-for-environment
link-type="dfn"} `environment`{.variable}:

1.  Let `key`{.variable} be the result of running [obtain a storage key
    for non-storage
    purposes](#obtain-a-storage-key-for-non-storage-purposes){#ref-for-obtain-a-storage-key-for-non-storage-purposes
    link-type="dfn"} with `environment`{.variable}.

2.  If `key`{.variable}'s
    [origin](#storage-key-origin){#ref-for-storage-key-origin
    link-type="dfn"} is an [opaque
    origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque
    link-type="dfn"}, then return failure.

3.  If the user has disabled storage, then return failure.

4.  Return `key`{.variable}.
:::

::: {.algorithm algorithm="obtain a storage key for non-storage purposes"}
To [obtain a storage key for non-storage
purposes]{#obtain-a-storage-key-for-non-storage-purposes .dfn
.dfn-paneled dfn-type="dfn" export=""}, given an
[environment](https://html.spec.whatwg.org/multipage/webappapis.html#environment){#ref-for-environment①
link-type="dfn"} `environment`{.variable}:

1.  Let `origin`{.variable} be `environment`{.variable}'s
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin
    link-type="dfn"} if `environment`{.variable} is an [environment
    settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object①
    link-type="dfn"}; otherwise `environment`{.variable}'s [creation
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#concept-environment-creation-url){#ref-for-concept-environment-creation-url
    link-type="dfn"}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin
    link-type="dfn"}.

2.  Return a
    [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①
    link-type="dfn"} consisting of `origin`{.variable}.
:::

::: {.algorithm algorithm="equal" algorithm-for="storage key"}
To determine whether a [storage key](#storage-key){#ref-for-storage-key②
link-type="dfn"} `A`{.variable} [equals]{#storage-key-equal .dfn
.dfn-paneled dfn-for="storage key" dfn-type="dfn" export="" lt="equal"}
[storage key](#storage-key){#ref-for-storage-key③ link-type="dfn"}
`B`{.variable}:

1.  If `A`{.variable}'s
    [origin](#storage-key-origin){#ref-for-storage-key-origin①
    link-type="dfn"} is not [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin
    link-type="dfn"} with `B`{.variable}'s
    [origin](#storage-key-origin){#ref-for-storage-key-origin②
    link-type="dfn"}, then return false.

2.  Return true.
:::

### [4.3. ]{.secno}[Storage sheds]{.content}[](#storage-sheds){.self-link} {#storage-sheds .heading .settled level="4.3"}

A [storage shed]{#storage-shed .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a
[map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map①
link-type="dfn"} of [storage keys](#storage-key){#ref-for-storage-key④
link-type="dfn"} to [storage
shelves](#storage-shelf){#ref-for-storage-shelf② link-type="dfn"}. It is
initially empty.

------------------------------------------------------------------------

A [user
agent](https://infra.spec.whatwg.org/#user-agent){#ref-for-user-agent①
link-type="dfn"} holds a [storage shed]{#user-agent-storage-shed .dfn
.dfn-paneled dfn-for="user agent" dfn-type="dfn" noexport=""}, which is
a [storage shed](#storage-shed){#ref-for-storage-shed① link-type="dfn"}.
A user agent's [storage
shed](#user-agent-storage-shed){#ref-for-user-agent-storage-shed
link-type="dfn"} holds all [local storage]{#local-storage .dfn
.dfn-paneled dfn-type="dfn" noexport=""} data.

A [traversable
navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable
link-type="dfn"} holds a [[]{#browsing-session-storage-shed
.bs-old-id}storage shed]{#traversable-navigable-storage-shed .dfn
.dfn-paneled dfn-for="traversable navigable" dfn-type="dfn"
noexport=""}, which is a [storage
shed](#storage-shed){#ref-for-storage-shed② link-type="dfn"}. A
[traversable
navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable①
link-type="dfn"}'s [storage
shed](#traversable-navigable-storage-shed){#ref-for-traversable-navigable-storage-shed
link-type="dfn"} holds all [session storage]{#session-storage .dfn
.dfn-paneled dfn-type="dfn" noexport=""} data.

To [[]{#legacy-clone-a-browsing-session-storage-shed
.bs-old-id}legacy-clone a traversable storage
shed]{#legacy-clone-a-traversable-storage-shed .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a [traversable
navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable②
link-type="dfn"} `A`{.variable} and a [traversable
navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable③
link-type="dfn"} `B`{.variable}, run these steps:

1.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `key`{.variable} → `shelf`{.variable} of
    `A`{.variable}'s [storage
    shed](#traversable-navigable-storage-shed){#ref-for-traversable-navigable-storage-shed①
    link-type="dfn"}:

    1.  Let `newShelf`{.variable} be the result of running [create a
        storage
        shelf](#create-a-storage-shelf){#ref-for-create-a-storage-shelf
        link-type="dfn"} with \"`session`\".

    2.  Set `newShelf`{.variable}'s [bucket
        map](#bucket-map){#ref-for-bucket-map
        link-type="dfn"}\[\"`default`\"\]\'s [bottle
        map](#bottle-map){#ref-for-bottle-map
        link-type="dfn"}\[\"`sessionStorage`\"\]\'s
        [map](#storage-bottle-map){#ref-for-storage-bottle-map
        link-type="dfn"} to a
        [clone](https://infra.spec.whatwg.org/#map-clone){#ref-for-map-clone
        link-type="dfn"} of `shelf`{.variable}'s [bucket
        map](#bucket-map){#ref-for-bucket-map①
        link-type="dfn"}\[\"`default`\"\]\'s [bottle
        map](#bottle-map){#ref-for-bottle-map①
        link-type="dfn"}\[\"`sessionStorage`\"\]\'s
        [map](#storage-bottle-map){#ref-for-storage-bottle-map①
        link-type="dfn"}.

    3.  Set `B`{.variable}'s [storage
        shed](#traversable-navigable-storage-shed){#ref-for-traversable-navigable-storage-shed②
        link-type="dfn"}\[`key`{.variable}\] to `newShelf`{.variable}.

This is considered legacy as the benefits, if any, do not outweigh the
implementation complexity. And therefore it will not be expanded or used
outside of HTML.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

### [4.4. ]{.secno}[Storage shelves]{.content}[](#storage-shelves){.self-link} {#storage-shelves .heading .settled level="4.4"}

A [[]{#site-storage-unit .bs-old-id}storage shelf]{#storage-shelf .dfn
.dfn-paneled dfn-type="dfn" lt="storage shelf|storage shelves"
noexport=""} exists for each [storage
key](#storage-key){#ref-for-storage-key⑤ link-type="dfn"} within a
[storage shed](#storage-shed){#ref-for-storage-shed③ link-type="dfn"}.
It holds a [bucket map]{#bucket-map .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, which is a
[map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map②
link-type="dfn"} of
[strings](https://infra.spec.whatwg.org/#string){#ref-for-string
link-type="dfn"} to [storage
buckets](#storage-bucket){#ref-for-storage-bucket③ link-type="dfn"}.

For now \"`default`\" is the only
[key](https://infra.spec.whatwg.org/#map-key){#ref-for-map-key
link-type="dfn"} that exists in a [bucket
map](#bucket-map){#ref-for-bucket-map② link-type="dfn"}. See [issue
#2](https://github.com/whatwg/storage/issues/2). It is given a
[value](https://infra.spec.whatwg.org/#map-value){#ref-for-map-value
link-type="dfn"} when a [storage
shelf](#storage-shelf){#ref-for-storage-shelf③ link-type="dfn"} is
[obtained](#obtain-a-storage-shelf){#ref-for-obtain-a-storage-shelf
link-type="dfn"} for the first time.

To [obtain a storage shelf]{#obtain-a-storage-shelf .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [storage
shed](#storage-shed){#ref-for-storage-shed④ link-type="dfn"}
`shed`{.variable}, an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object②
link-type="dfn"} `environment`{.variable}, and a [storage
type](#storage-type){#ref-for-storage-type③ link-type="dfn"}
`type`{.variable}, run these steps:

1.  Let `key`{.variable} be the result of running [obtain a storage
    key](#obtain-a-storage-key){#ref-for-obtain-a-storage-key
    link-type="dfn"} with `environment`{.variable}.

2.  If `key`{.variable} is failure, then return failure.

3.  If `shed`{.variable}\[`key`{.variable}\] does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists
    link-type="dfn"}, then set `shed`{.variable}\[`key`{.variable}\] to
    the result of running [create a storage
    shelf](#create-a-storage-shelf){#ref-for-create-a-storage-shelf①
    link-type="dfn"} with `type`{.variable}.

4.  Return `shed`{.variable}\[`key`{.variable}\].

To [obtain a local storage shelf]{#obtain-a-local-storage-shelf .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object③
link-type="dfn"} `environment`{.variable}, return the result of running
[obtain a storage
shelf](#obtain-a-storage-shelf){#ref-for-obtain-a-storage-shelf①
link-type="dfn"} with the user agent's [storage
shed](#user-agent-storage-shed){#ref-for-user-agent-storage-shed①
link-type="dfn"}, `environment`{.variable}, and \"`local`\".

To [create a storage shelf]{#create-a-storage-shelf .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [storage
type](#storage-type){#ref-for-storage-type④ link-type="dfn"}
`type`{.variable}, run these steps:

1.  Let `shelf`{.variable} be a new [storage
    shelf](#storage-shelf){#ref-for-storage-shelf④ link-type="dfn"}.

2.  Set `shelf`{.variable}'s [bucket
    map](#bucket-map){#ref-for-bucket-map③
    link-type="dfn"}\[\"`default`\"\] to the result of running [create a
    storage
    bucket](#create-a-storage-bucket){#ref-for-create-a-storage-bucket
    link-type="dfn"} with `type`{.variable}.

3.  Return `shelf`{.variable}.

### [4.5. ]{.secno}[]{#boxes .bs-old-id}[Storage buckets]{.content}[](#buckets){.self-link} {#buckets .heading .settled level="4.5"}

A [storage bucket]{#storage-bucket .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a place for [storage
endpoints](#storage-endpoint){#ref-for-storage-endpoint⑦
link-type="dfn"} to store data.

A [storage bucket](#storage-bucket){#ref-for-storage-bucket④
link-type="dfn"} has a [bottle map]{#bottle-map .dfn .dfn-paneled
dfn-type="dfn" noexport=""} of [storage
identifiers](#storage-identifier){#ref-for-storage-identifier②
link-type="dfn"} to [storage
bottles](#storage-bottle){#ref-for-storage-bottle③ link-type="dfn"}.

------------------------------------------------------------------------

A [[]{#box .bs-old-id}local storage bucket]{#bucket .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is a [storage
bucket](#storage-bucket){#ref-for-storage-bucket⑤ link-type="dfn"} for
[local storage](#local-storage){#ref-for-local-storage① link-type="dfn"}
APIs.

A [local storage bucket](#bucket){#ref-for-bucket link-type="dfn"} has a
[[]{#box-mode .bs-old-id}mode]{#bucket-mode .dfn .dfn-paneled
dfn-for="local storage bucket" dfn-type="dfn" noexport=""}, which is
\"`best-effort`\" or \"`persistent`\". It is initially
\"`best-effort`\".

------------------------------------------------------------------------

A [session storage bucket]{#session-storage-bucket .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is a [storage
bucket](#storage-bucket){#ref-for-storage-bucket⑥ link-type="dfn"} for
[session storage](#session-storage){#ref-for-session-storage①
link-type="dfn"} APIs.

------------------------------------------------------------------------

To [create a storage bucket]{#create-a-storage-bucket .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [storage
type](#storage-type){#ref-for-storage-type⑤ link-type="dfn"}
`type`{.variable}, run these steps:

1.  Let `bucket`{.variable} be null.

2.  If `type`{.variable} is \"`local`\", then set `bucket`{.variable} to
    a new [local storage bucket](#bucket){#ref-for-bucket①
    link-type="dfn"}.

3.  Otherwise:

    1.  Assert: `type`{.variable} is \"`session`\".

    2.  Set `bucket`{.variable} to a new [session storage
        bucket](#session-storage-bucket){#ref-for-session-storage-bucket
        link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate
    link-type="dfn"} `endpoint`{.variable} of [registered storage
    endpoints](#registered-storage-endpoints){#ref-for-registered-storage-endpoints①
    link-type="dfn"} whose
    [types](#storage-endpoint-types){#ref-for-storage-endpoint-types①
    link-type="dfn"}
    [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain
    link-type="dfn"} `type`{.variable}, set `bucket`{.variable}'s
    [bottle map](#bottle-map){#ref-for-bottle-map②
    link-type="dfn"}\[`endpoint`{.variable}'s
    [identifier](#storage-endpoint-identifier){#ref-for-storage-endpoint-identifier①
    link-type="dfn"}\] to a new [storage
    bottle](#storage-bottle){#ref-for-storage-bottle④ link-type="dfn"}
    whose [quota](#storage-bottle-quota){#ref-for-storage-bottle-quota①
    link-type="dfn"} is `endpoint`{.variable}'s
    [quota](#storage-endpoint-quota){#ref-for-storage-endpoint-quota①
    link-type="dfn"}.

5.  Return `bucket`{.variable}.

### [4.6. ]{.secno}[Storage bottles]{.content}[](#storage-bottles){.self-link} {#storage-bottles .heading .settled level="4.6"}

A [storage bottle]{#storage-bottle .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a part of a [storage
bucket](#storage-bucket){#ref-for-storage-bucket⑦ link-type="dfn"}
carved out for a single [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint⑧
link-type="dfn"}. A [storage
bottle](#storage-bottle){#ref-for-storage-bottle⑤ link-type="dfn"} has a
[map]{#storage-bottle-map .dfn .dfn-paneled dfn-for="storage bottle"
dfn-type="dfn" noexport=""}, which is initially an empty
[map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map③
link-type="dfn"}. A [storage
bottle](#storage-bottle){#ref-for-storage-bottle⑥ link-type="dfn"} also
has a [proxy map reference set]{#storage-bottle-proxy-map-reference-set
.dfn .dfn-paneled dfn-for="storage bottle" dfn-type="dfn" noexport=""},
which is initially an empty
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set②
link-type="dfn"}. A [storage
bottle](#storage-bottle){#ref-for-storage-bottle⑦ link-type="dfn"} also
has a [quota]{#storage-bottle-quota .dfn .dfn-paneled
dfn-for="storage bottle" dfn-type="dfn" noexport=""}, which is null or a
number representing a conservative estimate of the total amount of bytes
it can hold. Null indicates the lack of a limit. [It is still bound by
the [storage quota](#storage-quota){#ref-for-storage-quota
link-type="dfn"} of its encompassing [storage
shelf](#storage-shelf){#ref-for-storage-shelf⑤ link-type="dfn"}.
]{.note}

A [storage bottle](#storage-bottle){#ref-for-storage-bottle⑧
link-type="dfn"}'s
[map](#storage-bottle-map){#ref-for-storage-bottle-map② link-type="dfn"}
is where the actual data meant to be stored lives. User agents are
expected to store this data, and make it available across
[agent](https://tc39.github.io/ecma262/#sec-agents){#ref-for-sec-agents
link-type="dfn"} and even [agent
cluster](https://tc39.github.io/ecma262/#sec-agent-clusters){#ref-for-sec-agent-clusters
link-type="dfn"} boundaries, in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined
link-type="dfn"} manner, so that this standard and standards using this
standard can access the contents.

------------------------------------------------------------------------

To [obtain a storage bottle map]{#obtain-a-storage-bottle-map .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given a [storage
type](#storage-type){#ref-for-storage-type⑥ link-type="dfn"}
`type`{.variable}, [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object④
link-type="dfn"} `environment`{.variable}, and [storage
identifier](#storage-identifier){#ref-for-storage-identifier③
link-type="dfn"} `identifier`{.variable}, run these steps:

1.  Let `shed`{.variable} be null.

2.  If `type`{.variable} is \"`local`\", then set `shed`{.variable} to
    the user agent's [storage
    shed](#user-agent-storage-shed){#ref-for-user-agent-storage-shed②
    link-type="dfn"}.

3.  Otherwise:

    1.  Assert: `type`{.variable} is \"`session`\".

    2.  Set `shed`{.variable} to `environment`{.variable}'s [global
        object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object
        link-type="dfn"}'s [associated
        `Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window
        link-type="dfn"}'s [node
        navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#node-navigable){#ref-for-node-navigable
        link-type="dfn"}'s [traversable
        navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#nav-traversable){#ref-for-nav-traversable
        link-type="dfn"}'s [storage
        shed](#traversable-navigable-storage-shed){#ref-for-traversable-navigable-storage-shed③
        link-type="dfn"}.

4.  Let `shelf`{.variable} be the result of running [obtain a storage
    shelf](#obtain-a-storage-shelf){#ref-for-obtain-a-storage-shelf②
    link-type="dfn"}, with `shed`{.variable}, `environment`{.variable},
    and `type`{.variable}.

5.  If `shelf`{.variable} is failure, then return failure.

6.  Let `bucket`{.variable} be `shelf`{.variable}'s [bucket
    map](#bucket-map){#ref-for-bucket-map④
    link-type="dfn"}\[\"`default`\"\].

7.  Let `bottle`{.variable} be `bucket`{.variable}'s [bottle
    map](#bottle-map){#ref-for-bottle-map③
    link-type="dfn"}\[`identifier`{.variable}\].

8.  Let `proxyMap`{.variable} be a new [storage proxy
    map](#storage-proxy-map){#ref-for-storage-proxy-map①
    link-type="dfn"} whose [backing
    map](#storage-proxy-map-backing-map){#ref-for-storage-proxy-map-backing-map
    link-type="dfn"} is `bottle`{.variable}'s
    [map](#storage-bottle-map){#ref-for-storage-bottle-map③
    link-type="dfn"}.

9.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append
    link-type="dfn"} `proxyMap`{.variable} to `bottle`{.variable}'s
    [proxy map reference
    set](#storage-bottle-proxy-map-reference-set){#ref-for-storage-bottle-proxy-map-reference-set
    link-type="dfn"}.

10. Return `proxyMap`{.variable}.

To [obtain a local storage bottle
map]{#obtain-a-local-storage-bottle-map .dfn .dfn-paneled dfn-type="dfn"
export=""}, given an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑤
link-type="dfn"} `environment`{.variable} and [storage
identifier](#storage-identifier){#ref-for-storage-identifier④
link-type="dfn"} `identifier`{.variable}, return the result of running
[obtain a storage bottle
map](#obtain-a-storage-bottle-map){#ref-for-obtain-a-storage-bottle-map
link-type="dfn"} with \"`local`\", `environment`{.variable}, and
`identifier`{.variable}.

To [obtain a session storage bottle
map]{#obtain-a-session-storage-bottle-map .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑥
link-type="dfn"} `environment`{.variable} and [storage
identifier](#storage-identifier){#ref-for-storage-identifier⑤
link-type="dfn"} `identifier`{.variable}, return the result of running
[obtain a storage bottle
map](#obtain-a-storage-bottle-map){#ref-for-obtain-a-storage-bottle-map①
link-type="dfn"} with \"`session`\", `environment`{.variable}, and
`identifier`{.variable}.

### [4.7. ]{.secno}[Storage proxy maps]{.content}[](#storage-proxy-maps){.self-link} {#storage-proxy-maps .heading .settled level="4.7"}

A [storage proxy map]{#storage-proxy-map .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is equivalent to a
[map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map④
link-type="dfn"}, except that all operations are instead performed on
its [backing map]{#storage-proxy-map-backing-map .dfn .dfn-paneled
dfn-for="storage proxy map" dfn-type="dfn" noexport=""}.

This allows for the [backing
map](#storage-proxy-map-backing-map){#ref-for-storage-proxy-map-backing-map①
link-type="dfn"} to be replaced. This is needed for [issue
#4](https://github.com/whatwg/storage/issues/4) and potentially the
[Storage Access API](https://privacycg.github.io/storage-access/).

### [4.8. ]{.secno}[Storage task source]{.content}[](#storage-task-source){.self-link} {#storage-task-source .heading .settled level="4.8"}

The [storage task source]{#task-source .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [task
source](https://html.spec.whatwg.org/multipage/webappapis.html#task-source){#ref-for-task-source
link-type="dfn"} to be used for all
[tasks](https://html.spec.whatwg.org/multipage/webappapis.html#concept-task){#ref-for-concept-task
link-type="dfn"} related to a [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint⑨
link-type="dfn"}. In particular those that relate to a [storage
endpoint](#storage-endpoint){#ref-for-storage-endpoint①⓪
link-type="dfn"}'s
[quota](#storage-endpoint-quota){#ref-for-storage-endpoint-quota②
link-type="dfn"}.

::: {.algorithm algorithm="queue a storage task"}
To [queue a storage task]{#queue-a-storage-task .dfn .dfn-paneled
dfn-type="dfn" export=""} given a [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object①
link-type="dfn"} `global`{.variable} and a series of steps
`steps`{.variable}, [queue a global
task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task
link-type="dfn"} on the [storage task
source](#task-source){#ref-for-task-source① link-type="dfn"} with
`global`{.variable} and `steps`{.variable}.
:::

