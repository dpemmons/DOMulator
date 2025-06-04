## [2. ]{.secno}[Security considerations]{.content}[](#security-considerations){.self-link} {#security-considerations .heading .settled level="2"}

The security of a [URL](#concept-url){#ref-for-concept-url①
link-type="dfn"} is a function of its environment. Care is to be taken
when rendering, interpreting, and passing
[URLs](#concept-url){#ref-for-concept-url② link-type="dfn"} around.

When rendering and allocating new
[URLs](#concept-url){#ref-for-concept-url③ link-type="dfn"} \"spoofing\"
needs to be considered. An attack whereby one
[host](#concept-host){#ref-for-concept-host② link-type="dfn"} or
[URL](#concept-url){#ref-for-concept-url④ link-type="dfn"} can be
confused for another. For instance, consider how 1/l/I, m/rn/rri, 0/O,
and а/a can all appear eerily similar. Or worse, consider how U+202A
LEFT-TO-RIGHT EMBEDDING and similar [code
points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point④
link-type="dfn"} are invisible.
[\[UTR36\]](#biblio-utr36 "Unicode Security Considerations"){link-type="biblio"}

When passing a [URL](#concept-url){#ref-for-concept-url⑤
link-type="dfn"} from party `A`{.variable} to `B`{.variable}, both need
to carefully consider what is happening. `A`{.variable} might end up
leaking data it does not want to leak. `B`{.variable} might receive
input it did not expect and take an action that harms the user. In
particular, `B`{.variable} should never trust `A`{.variable}, as at some
point [URLs](#concept-url){#ref-for-concept-url⑥ link-type="dfn"} from
`A`{.variable} can come from untrusted sources.

