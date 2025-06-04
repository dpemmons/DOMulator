## [1. ]{.secno}[Preface]{.content}[](#preface){.self-link} {#preface .short .heading .settled level="1"}

At a high level, fetching a resource is a fairly simple operation. A
request goes in, a response comes out. The details of that operation are
however quite involved and used to not be written down carefully and
differ from one API to the next.

Numerous APIs provide the ability to fetch a resource, e.g. HTML's `img`
and `script` element, CSS\' `cursor` and `list-style-image`, the
`navigator.sendBeacon()` and `self.importScripts()` JavaScript APIs. The
Fetch Standard provides a unified architecture for these features so
they are all consistent when it comes to various aspects of fetching,
such as redirects and the CORS protocol.

The Fetch Standard also defines the
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch .idl-code
link-type="method"} JavaScript API, which exposes most of the networking
functionality at a fairly low level of abstraction.

