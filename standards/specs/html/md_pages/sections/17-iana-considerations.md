## [17]{.secno} IANA considerations[](#iana){.self-link} {#iana}

### [17.1]{.secno} [`text/html`]{.dfn}[](#text/html){.self-link} {#text/html}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   text

Subtype name:
:   html

Required parameters:
:   No required parameters

Optional parameters:

:   

    `charset`

    :   The `charset` parameter may be provided to specify the
        [document\'s character
        encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#text/html:document's-character-encoding
        x-internal="document's-character-encoding"}, overriding any
        [character encoding
        declarations](semantics.html#character-encoding-declaration){#text/html:character-encoding-declaration}
        in the document other than a Byte Order Mark (BOM). The
        parameter\'s value must be an [ASCII
        case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#text/html:ascii-case-insensitive
        x-internal="ascii-case-insensitive"} match for the string
        \"`utf-8`\". [\[ENCODING\]](references.html#refsENCODING)

Encoding considerations:
:   8bit (see the section on [character encoding
    declarations](semantics.html#character-encoding-declaration){#text/html:character-encoding-declaration-2})

Security considerations:

:   Entire novels have been written about the security considerations
    that apply to HTML documents. Many are listed in this document, to
    which the reader is referred for more details. Some general concerns
    bear mentioning here, however:

    HTML is scripted language, and has a large number of APIs (some of
    which are described in this document). Script can expose the user to
    potential risks of information leakage, credential leakage,
    cross-site scripting attacks, cross-site request forgeries, and a
    host of other problems. While the designs in this specification are
    intended to be safe if implemented correctly, a full implementation
    is a massive undertaking and, as with any software, user agents are
    likely to have security bugs.

    Even without scripting, there are specific features in HTML which,
    for historical reasons, are required for broad compatibility with
    legacy content but that expose the user to unfortunate security
    problems. In particular, the
    [`img`{#text/html:the-img-element}](embedded-content.html#the-img-element)
    element can be used in conjunction with some other features as a way
    to effect a port scan from the user\'s location on the Internet.
    This can expose local network topologies that the attacker would
    otherwise not be able to determine.

    HTML relies on a compartmentalization scheme sometimes known as the
    *same-origin policy*. An
    [origin](browsers.html#concept-origin){#text/html:concept-origin} in
    most cases consists of all the pages served from the same host, on
    the same port, using the same protocol.

    It is critical, therefore, to ensure that any untrusted content that
    forms part of a site be hosted on a different
    [origin](browsers.html#concept-origin){#text/html:concept-origin-2}
    than any sensitive content on that site. Untrusted content can
    easily spoof any other page on the same origin, read data from that
    origin, cause scripts in that origin to execute, submit forms to and
    from that origin even if they are protected from cross-site request
    forgery attacks by unique tokens, and make use of any third-party
    resources exposed to or rights granted to that origin.

Interoperability considerations:
:   Rules for processing both conforming and non-conforming content are
    defined in this specification.

Published specification:
:   This document is the relevant specification. Labeling a resource
    with the [`text/html`{#text/html:text/html}](#text/html) type
    asserts that the resource is an [HTML
    document](https://dom.spec.whatwg.org/#html-document){#text/html:html-documents
    x-internal="html-documents"} using [the HTML
    syntax](syntax.html#syntax){#text/html:syntax}.

Applications that use this media type:
:   Web browsers, tools for processing web content, HTML authoring
    tools, search engines, validators.

Additional information:

:   

    Magic number(s):
    :   No sequence of bytes can uniquely identify an HTML document.
        More information on detecting HTML documents is available in
        MIME Sniffing. [\[MIMESNIFF\]](references.html#refsMIMESNIFF)

    File extension(s):
    :   \"`html`\" and \"`htm`\" are commonly, but certainly not
        exclusively, used as the extension for HTML documents.

    Macintosh file type code(s):
    :   `TEXT`

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   No restrictions apply.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#text/html:concept-url-fragment
x-internal="concept-url-fragment"} used with
[`text/html`{#text/html:text/html-2}](#text/html) resources either refer
to the [indicated
part](browsing-the-web.html#the-indicated-part-of-the-document){#text/html:the-indicated-part-of-the-document}
of the corresponding
[`Document`{#text/html:document}](dom.html#document), or provide state
information for in-page scripts.

### [17.2]{.secno} [`multipart/x-mixed-replace`]{.dfn}[](#multipart/x-mixed-replace){.self-link} {#multipart/x-mixed-replace}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   multipart

Subtype name:
:   x-mixed-replace

Required parameters:

:   - `boundary` (defined in RFC2046)
      [\[RFC2046\]](references.html#refsRFC2046)

Optional parameters:
:   No optional parameters.

Encoding considerations:
:   binary

Security considerations:
:   Subresources of a
    [`multipart/x-mixed-replace`{#multipart/x-mixed-replace:multipart/x-mixed-replace}](#multipart/x-mixed-replace)
    resource can be of any type, including types with non-trivial
    security implications such as
    [`text/html`{#multipart/x-mixed-replace:text/html}](#text/html).

Interoperability considerations:
:   None.

Published specification:
:   This specification describes processing rules for web browsers.
    Conformance requirements for generating resources with this type are
    the same as for
    [`multipart/mixed`{#multipart/x-mixed-replace:multipart/mixed}](indices.html#multipart/mixed).
    [\[RFC2046\]](references.html#refsRFC2046)

Applications that use this media type:
:   This type is intended to be used in resources generated by web
    servers, for consumption by web browsers.

Additional information:

:   

    Magic number(s):
    :   No sequence of bytes can uniquely identify a
        [`multipart/x-mixed-replace`{#multipart/x-mixed-replace:multipart/x-mixed-replace-2}](#multipart/x-mixed-replace)
        resource.

    File extension(s):
    :   No specific file extensions are recommended for this type.

    Macintosh file type code(s):
    :   No specific Macintosh file type codes are recommended for this
        type.

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   No restrictions apply.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#multipart/x-mixed-replace:concept-url-fragment
x-internal="concept-url-fragment"} used with
[`multipart/x-mixed-replace`{#multipart/x-mixed-replace:multipart/x-mixed-replace-3}](#multipart/x-mixed-replace)
resources apply to each body part as defined by the type used by that
body part.

### [17.3]{.secno} [`application/xhtml+xml`]{.dfn}[](#application/xhtml+xml){.self-link} {#application/xhtml+xml}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   application

Subtype name:
:   xhtml+xml

Required parameters:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Optional parameters:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml-2}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Encoding considerations:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml-3}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Security considerations:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml-4}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Interoperability considerations:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml-5}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Published specification:
:   Labeling a resource with the
    [`application/xhtml+xml`{#application/xhtml+xml:application/xhtml+xml}](#application/xhtml+xml)
    type asserts that the resource is an XML document that likely has a
    [document
    element](https://dom.spec.whatwg.org/#document-element){#application/xhtml+xml:document-element
    x-internal="document-element"} from the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#application/xhtml+xml:html-namespace-2
    x-internal="html-namespace-2"}. Thus, the relevant specifications
    are XML, Namespaces in XML, and this specification.
    [\[XML\]](references.html#refsXML)
    [\[XMLNS\]](references.html#refsXMLNS)

Applications that use this media type:
:   Same as for
    [`application/xml`{#application/xhtml+xml:application/xml-6}](indices.html#application/xml)
    [\[RFC7303\]](references.html#refsRFC7303)

Additional information:

:   

    Magic number(s):
    :   Same as for
        [`application/xml`{#application/xhtml+xml:application/xml-7}](indices.html#application/xml)
        [\[RFC7303\]](references.html#refsRFC7303)

    File extension(s):
    :   \"`xhtml`\" and \"`xht`\" are sometimes used as extensions for
        XML resources that have a [document
        element](https://dom.spec.whatwg.org/#document-element){#application/xhtml+xml:document-element-2
        x-internal="document-element"} from the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#application/xhtml+xml:html-namespace-2-2
        x-internal="html-namespace-2"}.

    Macintosh file type code(s):
    :   `TEXT`

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   No restrictions apply.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#application/xhtml+xml:concept-url-fragment
x-internal="concept-url-fragment"} used with
[`application/xhtml+xml`{#application/xhtml+xml:application/xhtml+xml-2}](#application/xhtml+xml)
resources have the same semantics as with any [XML MIME
type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#application/xhtml+xml:xml-mime-type
x-internal="xml-mime-type"}. [\[RFC7303\]](references.html#refsRFC7303)

### [17.4]{.secno} [`text/ping`]{.dfn}[](#text/ping){.self-link} {#text/ping}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   text

Subtype name:
:   ping

Required parameters:
:   No parameters

Optional parameters:

:   

    `charset`

    :   The `charset` parameter may be provided. The parameter\'s value
        must be \"`utf-8`\". This parameter serves no purpose; it is
        only allowed for compatibility with legacy servers.

Encoding considerations:
:   Not applicable.

Security considerations:

:   If used exclusively in the fashion described in the context of
    [hyperlink
    auditing](links.html#hyperlink-auditing){#text/ping:hyperlink-auditing},
    this type introduces no new security concerns.

Interoperability considerations:
:   Rules applicable to this type are defined in this specification.

Published specification:
:   This document is the relevant specification.

Applications that use this media type:
:   Web browsers.

Additional information:

:   

    Magic number(s):
    :   [`text/ping`{#text/ping:text/ping}](#text/ping) resources always
        consist of the four bytes 0x50 0x49 0x4E 0x47 (\``PING`\`).

    File extension(s):
    :   No specific file extension is recommended for this type.

    Macintosh file type code(s):
    :   No specific Macintosh file type codes are recommended for this
        type.

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   Only intended for use with HTTP POST requests generated as part of a
    web browser\'s processing of the
    [`ping`{#text/ping:ping}](links.html#ping) attribute.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#text/ping:concept-url-fragment
x-internal="concept-url-fragment"} have no meaning with
[`text/ping`{#text/ping:text/ping-2}](#text/ping) resources.

### [17.5]{.secno} [`application/microdata+json`]{.dfn}[](#application/microdata+json){.self-link} {#application/microdata+json}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   application

Subtype name:
:   microdata+json

Required parameters:
:   Same as for
    [`application/json`{#application/microdata+json:application/json}](indices.html#application/json)
    [\[JSON\]](references.html#refsJSON)

Optional parameters:
:   Same as for
    [`application/json`{#application/microdata+json:application/json-2}](indices.html#application/json)
    [\[JSON\]](references.html#refsJSON)

Encoding considerations:
:   8bit (always UTF-8)

Security considerations:
:   Same as for
    [`application/json`{#application/microdata+json:application/json-3}](indices.html#application/json)
    [\[JSON\]](references.html#refsJSON)

Interoperability considerations:
:   Same as for
    [`application/json`{#application/microdata+json:application/json-4}](indices.html#application/json)
    [\[JSON\]](references.html#refsJSON)

Published specification:
:   Labeling a resource with the
    [`application/microdata+json`{#application/microdata+json:application/microdata+json}](#application/microdata+json)
    type asserts that the resource is a JSON text that consists of an
    object with a single entry called \"`items`\" consisting of an array
    of entries, each of which consists of an object with an entry called
    \"`id`\" whose value is a string, an entry called \"`type`\" whose
    value is another string, and an entry called \"`properties`\" whose
    value is an object whose entries each have a value consisting of an
    array of either objects or strings, the objects being of the same
    form as the objects in the aforementioned \"`items`\" entry. Thus,
    the relevant specifications are JSON and this specification.
    [\[JSON\]](references.html#refsJSON)

Applications that use this media type:

:   Applications that transfer data intended for use with HTML\'s
    microdata feature, especially in the context of drag-and-drop, are
    the primary application class for this type.

Additional information:

:   

    Magic number(s):
    :   Same as for
        [`application/json`{#application/microdata+json:application/json-5}](indices.html#application/json)
        [\[JSON\]](references.html#refsJSON)

    File extension(s):
    :   Same as for
        [`application/json`{#application/microdata+json:application/json-6}](indices.html#application/json)
        [\[JSON\]](references.html#refsJSON)

    Macintosh file type code(s):
    :   Same as for
        [`application/json`{#application/microdata+json:application/json-7}](indices.html#application/json)
        [\[JSON\]](references.html#refsJSON)

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   No restrictions apply.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#application/microdata+json:concept-url-fragment
x-internal="concept-url-fragment"} used with
[`application/microdata+json`{#application/microdata+json:application/microdata+json-2}](#application/microdata+json)
resources have the same semantics as when used with
[`application/json`{#application/microdata+json:application/json-8}](indices.html#application/json)
(namely, at the time of writing, no semantics at all).
[\[JSON\]](references.html#refsJSON)

### [17.6]{.secno} [`text/event-stream`]{.dfn}[](#text/event-stream){.self-link} {#text/event-stream}

This registration is for community review and will be submitted to the
IESG for review, approval, and registration with IANA.

Type name:
:   text

Subtype name:
:   event-stream

Required parameters:
:   No parameters

Optional parameters:

:   

    `charset`

    :   The `charset` parameter may be provided. The parameter\'s value
        must be \"`utf-8`\". This parameter serves no purpose; it is
        only allowed for compatibility with legacy servers.

Encoding considerations:
:   8bit (always UTF-8)

Security considerations:

:   An event stream from an origin distinct from the origin of the
    content consuming the event stream can result in information
    leakage. To avoid this, user agents are required to apply CORS
    semantics. [\[FETCH\]](references.html#refsFETCH)

    Event streams can overwhelm a user agent; a user agent is expected
    to apply suitable restrictions to avoid depleting local resources
    because of an overabundance of information from an event stream.

    Servers can be overwhelmed if a situation develops in which the
    server is causing clients to reconnect rapidly. Servers should use a
    5xx status code to indicate capacity problems, as this will prevent
    conforming clients from reconnecting automatically.

Interoperability considerations:
:   Rules for processing both conforming and non-conforming content are
    defined in this specification.

Published specification:
:   This document is the relevant specification.

Applications that use this media type:
:   Web browsers and tools using web services.

Additional information:

:   

    Magic number(s):
    :   No sequence of bytes can uniquely identify an event stream.

    File extension(s):
    :   No specific file extensions are recommended for this type.

    Macintosh file type code(s):
    :   No specific Macintosh file type codes are recommended for this
        type.

Person & email address to contact for further information:
:   Ian Hickson \<ian@hixie.ch\>

Intended usage:
:   Common

Restrictions on usage:
:   This format is only expected to be used by dynamic open-ended
    streams served using HTTP or a similar protocol. Finite resources
    are not expected to be labeled with this type.

Author:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   W3C

[Fragments](https://url.spec.whatwg.org/#concept-url-fragment){#text/event-stream:concept-url-fragment
x-internal="concept-url-fragment"} have no meaning with
[`text/event-stream`{#text/event-stream:text/event-stream}](#text/event-stream)
resources.

### [17.7]{.secno} [`web+` scheme prefix]{.dfn}[](#web+-scheme-prefix){.self-link} {#web+-scheme-prefix}

This section describes a convention for use with the IANA URI scheme
registry. It does not itself register a specific scheme.
[\[RFC7595\]](references.html#refsRFC7595)

Scheme name:
:   Schemes starting with the four characters \"`web+`\" followed by one
    or more letters in the range `a`-`z`.

Status:
:   Permanent

Scheme syntax:
:   Scheme-specific.

Scheme semantics:
:   Scheme-specific.

Encoding considerations:
:   All \"`web+`\" schemes should use UTF-8 encodings where relevant.

Applications/protocols that use this scheme name:
:   Scheme-specific.

Interoperability considerations:
:   The scheme is expected to be used in the context of web
    applications.

Security considerations:
:   Any web page is able to register a handler for all \"`web+`\"
    schemes. As such, these schemes must not be used for features
    intended to be core platform features (e.g., HTTP). Similarly, such
    schemes must not store confidential information in their URLs, such
    as usernames, passwords, personal information, or confidential
    project names.

Contact:
:   Ian Hickson \<ian@hixie.ch\>

Change controller:
:   Ian Hickson \<ian@hixie.ch\>

References:
:   Custom scheme handlers, HTML Living Standard:
    [https://html.spec.whatwg.org/#custom-handlers](system-state.html#custom-handlers)

[← 16 Obsolete features](obsolete.html) --- [Table of Contents](./) ---
[Index →](indices.html)
