## [9. ]{.secno}[XSLT]{.content}[](#xslt){.self-link} {#xslt .heading .settled level="9"}

XSL Transformations (XSLT) is a language for transforming XML documents
into other XML documents. The APIs defined in this section have been
widely implemented, and are maintained here so that they can be updated
when Web IDL changes. Complete definitions of these APIs remain
necessary and such work is tracked and can be contributed to in
[whatwg/dom#181](https://github.com/whatwg/dom/issues/181).
[\[XSLT\]](#biblio-xslt "XSL Transformations (XSLT) Version 1.0"){link-type="biblio"}

### [9.1. ]{.secno}[Interface [`XSLTProcessor`{.idl}](#xsltprocessor){#ref-for-xsltprocessor link-type="idl"}]{.content}[](#interface-xsltprocessor){.self-link} {#interface-xsltprocessor .heading .settled level="9.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XSLTProcessor {
  constructor();
  undefined importStylesheet(Node style);
  [CEReactions] DocumentFragment transformToFragment(Node source, Document output);
  [CEReactions] Document transformToDocument(Node source);
  undefined setParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName, any value);
  any getParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName);
  undefined removeParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName);
  undefined clearParameters();
  undefined reset();
};
```

