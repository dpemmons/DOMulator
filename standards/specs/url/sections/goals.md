## [Goals]{.content} {#goals .no-num .heading .settled}

The URL standard takes the following approach towards making URLs fully
interoperable:

- Align RFC 3986 and RFC 3987 with contemporary implementations and
  obsolete the RFCs in the process. (E.g., spaces, other \"illegal\"
  code points, query encoding, equality, canonicalization, are all
  concepts not entirely shared, or defined.) URL parsing needs to become
  as solid as HTML parsing.
  [\[RFC3986\]](#biblio-rfc3986 "Uniform Resource Identifier (URI): Generic Syntax"){link-type="biblio"}
  [\[RFC3987\]](#biblio-rfc3987 "Internationalized Resource Identifiers (IRIs)"){link-type="biblio"}

- Standardize on the term URL. URI and IRI are just confusing. In
  practice a single algorithm is used for both so keeping them distinct
  is not helping anyone. URL also easily wins the [search result
  popularity
  contest](https://trends.google.com/trends/explore?q=url,uri).

- Supplanting [Origin of a URI
  \[sic\]](https://tools.ietf.org/html/rfc6454#section-4).
  [\[RFC6454\]](#biblio-rfc6454 "The Web Origin Concept"){link-type="biblio"}

- Define URL's existing JavaScript API in full detail and add
  enhancements to make it easier to work with. Add a new
  [`URL`](#url){#ref-for-url .idl-code link-type="interface"} object as
  well for URL manipulation without usage of HTML elements. (Useful for
  JavaScript worker environments.)

- Ensure the combination of parser, serializer, and API guarantee
  idempotence. For example, a non-failure result of a
  parse-then-serialize operation will not change with any further
  parse-then-serialize operations applied to it. Similarly, manipulating
  a non-failure result through the API will not change from applying any
  number of serialize-then-parse operations to it.

As the editors learn more about the subject matter the goals might
increase in scope somewhat.

