## [11. ]{.secno}[Historical]{.content}[](#historical){.self-link} {#historical .heading .settled level="11"}

This standard used to contain several interfaces and interface members
that have been removed.

These interfaces have been removed:

- [`DOMConfiguration`]{#domconfiguration .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMError`]{#domerror .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMErrorHandler`]{#domerrorhandler .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMImplementationList`]{#domimplementationlist .dfn .dfn-paneled
  .idl-code dfn-type="interface" export=""}
- [`DOMImplementationSource`]{#domimplementationsource .dfn .dfn-paneled
  .idl-code dfn-type="interface" export=""}
- [`DOMLocator`]{#domlocator .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMObject`]{#domobject .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMUserData`]{#domuserdata .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`Entity`]{#entity .dfn .dfn-paneled .idl-code dfn-type="interface"
  export=""}
- [`EntityReference`]{#entityreference .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`MutationEvent`]{#mutationevent .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`MutationNameEvent`]{#mutationnameevent .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`NameList`]{#namelist .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`Notation`]{#notation .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`RangeException`]{#rangeexception .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`TypeInfo`]{#typeinfo .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`UserDataHandler`]{#userdatahandler .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}

And these interface members have been removed:

[`Attr`{.idl}](#attr){#ref-for-attr③⑨ link-type="idl"}

:   - [`schemaTypeInfo`]{#dom-attr-schematypeinfo .dfn .dfn-paneled
      .idl-code dfn-for="Attr" dfn-type="attribute" export=""}
    - [`isId`]{#dom-attr-isid .dfn .dfn-paneled .idl-code dfn-for="Attr"
      dfn-type="attribute" export=""}

[`Document`{.idl}](#document){#ref-for-document②⑧ link-type="idl"}

:   - [`createEntityReference()`]{#dom-document-createentityreference
      .dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
      export=""}
    - [`xmlEncoding`]{#dom-document-xmlencoding .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`xmlStandalone`]{#dom-document-xmlstandalone .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`xmlVersion`]{#dom-document-xmlversion .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`strictErrorChecking`]{#dom-document-stricterrorchecking .dfn
      .dfn-paneled .idl-code dfn-for="Document" dfn-type="attribute"
      export=""}
    - [`domConfig`]{#dom-document-domconfig .dfn .dfn-paneled .idl-code
      dfn-for="Document" dfn-type="attribute" export=""}
    - [`normalizeDocument()`]{#dom-document-normalizedocument .dfn
      .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
      export=""}
    - [`renameNode()`]{#dom-document-renamenode .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="method" export=""}

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②⑤ link-type="idl"}

:   - [`entities`]{#dom-documenttype-entities .dfn .dfn-paneled
      .idl-code dfn-for="DocumentType" dfn-type="attribute" export=""}
    - [`notations`]{#dom-documenttype-notations .dfn .dfn-paneled
      .idl-code dfn-for="DocumentType" dfn-type="attribute" export=""}
    - [`internalSubset`]{#dom-documenttype-internalsubset .dfn
      .dfn-paneled .idl-code dfn-for="DocumentType" dfn-type="attribute"
      export=""}

[`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation⑤ link-type="idl"}

:   - [`getFeature()`]{#dom-domimplementation-getfeature .dfn
      .dfn-paneled .idl-code dfn-for="DOMImplementation"
      dfn-type="method" export=""}

[`Element`{.idl}](#element){#ref-for-element⑤⑦ link-type="idl"}

:   - [`schemaTypeInfo`]{#dom-element-schematypeinfo .dfn .dfn-paneled
      .idl-code dfn-for="Element" dfn-type="attribute" export=""}
    - [`setIdAttribute()`]{#dom-element-setidattribute .dfn .dfn-paneled
      .idl-code dfn-for="Element" dfn-type="method" export=""}
    - [`setIdAttributeNS()`]{#dom-element-setidattributens .dfn
      .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
      export=""}
    - [`setIdAttributeNode()`]{#dom-element-setidattributenode .dfn
      .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
      export=""}

[`Node`{.idl}](#node){#ref-for-node①⓪⑧ link-type="idl"}

:   - [`isSupported`]{#dom-node-issupported .dfn .dfn-paneled .idl-code
      dfn-for="Node" dfn-type="attribute" export=""}
    - [`getFeature()`]{#dom-node-getfeature .dfn .dfn-paneled .idl-code
      dfn-for="Node" dfn-type="method" export=""}
    - [`getUserData()`]{#dom-node-getuserdata .dfn .dfn-paneled
      .idl-code dfn-for="Node" dfn-type="method" export=""}
    - [`setUserData()`]{#dom-node-setuserdata .dfn .dfn-paneled
      .idl-code dfn-for="Node" dfn-type="method" export=""}

[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①⑦ link-type="idl"}

:   - [`expandEntityReferences`]{#dom-nodeiterator-expandentityreferences
      .dfn .dfn-paneled .idl-code dfn-for="NodeIterator"
      dfn-type="attribute" export=""}

[`Text`{.idl}](#text){#ref-for-text⑤⑨ link-type="idl"}

:   - [`isElementContentWhitespace`]{#dom-text-iselementcontentwhitespace
      .dfn .dfn-paneled .idl-code dfn-for="Text" dfn-type="attribute"
      export=""}
    - [`replaceWholeText()`]{#dom-text-replacewholetext .dfn
      .dfn-paneled .idl-code dfn-for="Text" dfn-type="method" export=""}

[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker①① link-type="idl"}

:   - [`expandEntityReferences`]{#dom-treewalker-expandentityreferences
      .dfn .dfn-paneled .idl-code dfn-for="TreeWalker"
      dfn-type="attribute" export=""}

