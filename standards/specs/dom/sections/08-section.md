## [8. ]{.secno}[XPath]{.content}[](#xpath){.self-link} {#xpath .heading .settled level="8"}

DOM Level 3 XPath defined an API for evaluating XPath 1.0 expressions.
These APIs are widely implemented, but have not been maintained. The
interface definitions are maintained here so that they can be updated
when Web IDL changes. Complete definitions of these APIs remain
necessary and such work is tracked and can be contributed to in
[whatwg/dom#67](https://github.com/whatwg/dom/issues/67).
[\[DOM-Level-3-XPath\]](#biblio-dom-level-3-xpath "Document Object Model (DOM) Level 3 XPath Specification"){link-type="biblio"}
[\[XPath\]](#biblio-xpath "XML Path Language (XPath) Version 1.0"){link-type="biblio"}
[\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}

### [8.1. ]{.secno}[Interface [`XPathResult`{.idl}](#xpathresult){#ref-for-xpathresult link-type="idl"}]{.content}[](#interface-xpathresult){.self-link} {#interface-xpathresult .heading .settled level="8.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathResult {
  const unsigned short ANY_TYPE = 0;
  const unsigned short NUMBER_TYPE = 1;
  const unsigned short STRING_TYPE = 2;
  const unsigned short BOOLEAN_TYPE = 3;
  const unsigned short UNORDERED_NODE_ITERATOR_TYPE = 4;
  const unsigned short ORDERED_NODE_ITERATOR_TYPE = 5;
  const unsigned short UNORDERED_NODE_SNAPSHOT_TYPE = 6;
  const unsigned short ORDERED_NODE_SNAPSHOT_TYPE = 7;
  const unsigned short ANY_UNORDERED_NODE_TYPE = 8;
  const unsigned short FIRST_ORDERED_NODE_TYPE = 9;

  readonly attribute unsigned short resultType;
  readonly attribute unrestricted double numberValue;
  readonly attribute DOMString stringValue;
  readonly attribute boolean booleanValue;
  readonly attribute Node? singleNodeValue;
  readonly attribute boolean invalidIteratorState;
  readonly attribute unsigned long snapshotLength;

  Node? iterateNext();
  Node? snapshotItem(unsigned long index);
};
```

### [8.2. ]{.secno}[Interface [`XPathExpression`{.idl}](#xpathexpression){#ref-for-xpathexpression link-type="idl"}]{.content}[](#interface-xpathexpression){.self-link} {#interface-xpathexpression .heading .settled level="8.2"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathExpression {
  // XPathResult.ANY_TYPE = 0
  XPathResult evaluate(Node contextNode, optional unsigned short type = 0, optional XPathResult? result = null);
};
```

### [8.3. ]{.secno}[Mixin [`XPathEvaluatorBase`{.idl}](#xpathevaluatorbase){#ref-for-xpathevaluatorbase link-type="idl"}]{.content}[](#mixin-xpathevaluatorbase){.self-link} {#mixin-xpathevaluatorbase .heading .settled level="8.3"}

``` {.idl .highlight .def}
callback interface XPathNSResolver {
  DOMString? lookupNamespaceURI(DOMString? prefix);
};

interface mixin XPathEvaluatorBase {
  [NewObject] XPathExpression createExpression(DOMString expression, optional XPathNSResolver? resolver = null);
  Node createNSResolver(Node nodeResolver); // legacy
  // XPathResult.ANY_TYPE = 0
  XPathResult evaluate(DOMString expression, Node contextNode, optional XPathNSResolver? resolver = null, optional unsigned short type = 0, optional XPathResult? result = null);
};
Document includes XPathEvaluatorBase;
```

The
[`createNSResolver(``nodeResolver`{.variable}`)`]{#dom-xpathevaluatorbase-creatensresolver
.dfn .dfn-paneled .idl-code dfn-for="XPathEvaluatorBase"
dfn-type="method" export=""} method steps are to return
`nodeResolver`{.variable}.

This method exists only for historical reasons.

### [8.4. ]{.secno}[Interface [`XPathEvaluator`{.idl}](#xpathevaluator){#ref-for-xpathevaluator link-type="idl"}]{.content}[](#interface-xpathevaluator){.self-link} {#interface-xpathevaluator .heading .settled level="8.4"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathEvaluator {
  constructor();
};

XPathEvaluator includes XPathEvaluatorBase;
```

For historical reasons you can both construct
[`XPathEvaluator`{.idl}](#xpathevaluator){#ref-for-xpathevaluator②
link-type="idl"} and access the same methods on
[`Document`{.idl}](#document){#ref-for-document②⑤ link-type="idl"}.

