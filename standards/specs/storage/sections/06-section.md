## [6. ]{.secno}[Usage and quota]{.content}[](#usage-and-quota){.self-link} {#usage-and-quota .heading .settled level="6"}

The [storage usage]{#storage-usage .dfn .dfn-paneled dfn-type="dfn"
export=""} of a [storage shelf](#storage-shelf){#ref-for-storage-shelf⑥
link-type="dfn"} is an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①
link-type="dfn"} rough estimate of the amount of bytes used by it.

This cannot be an exact amount as user agents might, and are encouraged
to, use deduplication, compression, and other techniques that obscure
exactly how much bytes a [storage
shelf](#storage-shelf){#ref-for-storage-shelf⑦ link-type="dfn"} uses.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg "There is a tracking vector here."){.darkmode-aware
crossorigin="" height="64"
width="46"}](https://infra.spec.whatwg.org/#tracking-vector){.tracking-vector
style="color: currentcolor"} The [storage quota]{#storage-quota .dfn
.dfn-paneled dfn-type="dfn" export=""} of a [storage
shelf](#storage-shelf){#ref-for-storage-shelf⑧ link-type="dfn"} is an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined②
link-type="dfn"} conservative estimate of the total amount of bytes it
can hold. This amount should be less than the total storage space on the
device. It must not be a function of the available storage space on the
device.

::: {.note role="note"}
User agents are strongly encouraged to consider navigation frequency,
recency of visits, bookmarking, and [permission](#persistence) for
\"`persistent-storage`\" when determining quotas.

Directly or indirectly revealing available storage space can lead to
fingerprinting and leaking information outside the scope of the
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin④
link-type="dfn"} involved.
:::

