# DOM Standard Compliance Matrix

**Standard**: WHATWG DOM Living Standard
**Version**: Living Standard as of 2025-06-03
**Last Updated**: 2025-06-03

## Feature Compliance

| Feature/Section | Spec Section | Implemented | Compliance Level | Test Coverage | Notes |
|---|---|---|---|---|---|
| **1. Infrastructure** |  |  |  |  |  |
| 1.1 Trees | [dom.spec.whatwg.org/#trees](https://dom.spec.whatwg.org/#trees) | Yes | PartialCompliance |  | Basic tree structure (parent/child/sibling relationships) is implemented via Node interface and its implementations. Full compliance with all tree concepts (e.g., shadow trees, retargeting) needs verification. |
| 1.2 Ordered Sets | [dom.spec.whatwg.org/#ordered-sets](https://dom.spec.whatwg.org/#ordered-sets) | Yes | PartialCompliance |  | `NodeList` is an ordered collection. Full compliance with spec's definition of ordered sets needs verification. |
| 1.3 Selectors | [dom.spec.whatwg.org/#selectors](https://dom.spec.whatwg.org/#selectors) | Yes | PartialCompliance |  | Basic CSS selectors (tag, ID, class, descendant) are implemented in `internal/css/selectors.go`. `QuerySelector` and `QuerySelectorAll` are available. Missing: attribute selectors, pseudo-classes, pseudo-elements, combinators other than descendant, etc. |
| 1.4 Namespaces | [dom.spec.whatwg.org/#namespaces](https://dom.spec.whatwg.org/#namespaces) | No | NotImplemented |  | No explicit namespace handling (e.g., `namespaceURI`, `prefix`, `localName` on Attr/Element) is implemented. |
| **2. Events** |  |  |  |  |  |
| 2.1 Introduction | [dom.spec.whatwg.org/#introduction-to-dom-events](https://dom.spec.whatwg.org/#introduction-to-dom-events) | Yes | PartialCompliance |  | Basic event concepts are present. |
| 2.2 Interface `Event` | [dom.spec.whatwg.org/#interface-event](https://dom.spec.whatwg.org/#interface-event) | Yes | PartialCompliance |  | `BaseEvent` in `internal/dom/events.go` implements most properties (type, target, currentTarget, eventPhase, bubbles, cancelable, defaultPrevented, isTrusted, timeStamp) and methods (preventDefault, stopPropagation, stopImmediatePropagation). `composed` property and `composedPath()` method are missing. Timestamp is not yet properly implemented. |
| 2.4 Interface `CustomEvent` | [dom.spec.whatwg.org/#interface-customevent](https://dom.spec.whatwg.org/#interface-customevent) | Yes | PartialCompliance |  | Implemented in `internal/browser/events/customevent.go`. Includes `detail` property and constructor. `initCustomEvent` method is also present. |
| 2.5 Constructing events | [dom.spec.whatwg.org/#constructing-events](https://dom.spec.whatwg.org/#constructing-events) | Yes | PartialCompliance |  | `NewEvent` and `NewCustomEvent` allow event construction. Full compliance with event constructor dictionary (e.g. `EventInit`) needs verification. |
| 2.7 Interface `EventTarget` | [dom.spec.whatwg.org/#interface-eventtarget](https://dom.spec.whatwg.org/#interface-eventtarget) | Yes | PartialCompliance |  | Implemented on `Node` (`nodeImpl` in `internal/dom/node.go`) with `addEventListener`, `removeEventListener`, `dispatchEvent`. |
| 2.9 Dispatching events | [dom.spec.whatwg.org/#dispatching-events](https://dom.spec.whatwg.org/#dispatching-events) | Yes | PartialCompliance |  | `dispatchEvent` in `internal/dom/node.go` implements capturing and bubbling phases. Full algorithm compliance (e.g., shadow DOM, activation behavior) needs verification. |
| **3. Aborting ongoing activities** |  |  |  |  |  |
| 3.1 Interface `AbortController` | [dom.spec.whatwg.org/#interface-abortcontroller](https://dom.spec.whatwg.org/#interface-abortcontroller) | No | NotImplemented |  |  |
| 3.2 Interface `AbortSignal` | [dom.spec.whatwg.org/#interface-abortsignal](https://dom.spec.whatwg.org/#interface-abortsignal) | No | NotImplemented |  |  |
| **4. Nodes** |  |  |  |  |  |
| 4.2 Node tree | [dom.spec.whatwg.org/#node-trees](https://dom.spec.whatwg.org/#node-trees) | Yes | PartialCompliance |  | Covered by 1.1 and individual Node type implementations. Specific algorithms (e.g., adopting, cloning steps) need detailed verification. |
| 4.2.2 Shadow tree | [dom.spec.whatwg.org/#shadow-trees](https://dom.spec.whatwg.org/#shadow-trees) | No | NotImplemented |  | Covered by 4.8 Interface `ShadowRoot`. |
| 4.2.3 Mutation algorithms | [dom.spec.whatwg.org/#mutation-algorithms](https://dom.spec.whatwg.org/#mutation-algorithms) | No | NotImplemented |  | Algorithms like pre-insert, pre-remove, replace, etc., are not explicitly implemented per spec details. Basic operations exist but not full algorithm compliance. |
| 4.3 Mutation observers | [dom.spec.whatwg.org/#mutation-observers](https://dom.spec.whatwg.org/#mutation-observers) | No | NotImplemented |  |  |
| 4.4 Interface `Node` | [dom.spec.whatwg.org/#interface-node](https://dom.spec.whatwg.org/#interface-node) | Yes | PartialCompliance |  | Core properties (nodeType, nodeName, nodeValue, parentNode, childNodes, firstChild, lastChild, previousSibling, nextSibling, ownerDocument) and methods (appendChild, removeChild, insertBefore, replaceChild, cloneNode, addEventListener, removeEventListener, dispatchEvent) are implemented. Full spec compliance for all attributes (e.g., baseURI, isConnected, textContent) and methods (e.g., normalize, isEqualNode, compareDocumentPosition, contains) needs detailed verification. |
| 4.5 Interface `Document` | [dom.spec.whatwg.org/#interface-document](https://dom.spec.whatwg.org/#interface-document) | Yes | PartialCompliance |  | Implemented methods: CreateElement, CreateTextNode, CreateComment, CreateDocumentFragment, GetElementById, GetElementsByTagName, GetElementsByClassName. Implemented properties: DocumentElement, Body, Head. Many properties (e.g., URL, compatMode, characterSet, contentType, doctype, implementation) and methods (e.g., querySelector, querySelectorAll, importNode, adoptNode, createAttribute, createEvent) are not yet implemented. |
| 4.6 Interface `DocumentType` | [dom.spec.whatwg.org/#interface-documenttype](https://dom.spec.whatwg.org/#interface-documenttype) | Yes | PartialCompliance |  | Implemented properties: name, publicId, systemId. Methods like `before()`, `after()`, `replaceWith()`, `remove()` are inherited from Node but not specifically tested for DocumentType. |
| 4.7 Interface `DocumentFragment` | [dom.spec.whatwg.org/#interface-documentfragment](https://dom.spec.whatwg.org/#interface-documentfragment) | Yes | PartialCompliance |  | Basic structure implemented. It inherits Node properties and methods. Specific DocumentFragment methods like `getElementById`, `querySelector`, `querySelectorAll` are not implemented. |
| 4.8 Interface `ShadowRoot` | [dom.spec.whatwg.org/#interface-shadowroot](https://dom.spec.whatwg.org/#interface-shadowroot) | No | NotImplemented |  | Shadow DOM is not currently implemented. |
| 4.9 Interface `Element` | [dom.spec.whatwg.org/#interface-element](https://dom.spec.whatwg.org/#interface-element) | Yes | PartialCompliance |  | Implemented: tagName, attributes (via Get/Set/Has/RemoveAttribute), HasClass, GetElementsByTagName, GetElementsByClassName, InnerHTML (basic), OuterHTML (basic), TextContent (basic), CloneNode, InsertAdjacentHTML (basic). Missing: id, className, classList, dataset, style, querySelector, querySelectorAll, closest, matches, and many other properties/methods. Attribute handling is basic (map[string]string). HTML parsing for InnerHTML/InsertAdjacentHTML is very rudimentary. |
| 4.9.1 Interface `NamedNodeMap` | [dom.spec.whatwg.org/#interface-namednodemap](https://dom.spec.whatwg.org/#interface-namednodemap) | No | NotImplemented |  | Attributes are currently handled as a simple `map[string]string` in `Element`, not as a `NamedNodeMap` object. |
| 4.9.2 Interface `Attr` | [dom.spec.whatwg.org/#interface-attr](https://dom.spec.whatwg.org/#interface-attr) | Yes | PartialCompliance |  | Implemented properties: name, value, ownerElement (though `ownerElement` is not always set correctly by `AttributeMap`). Missing: namespaceURI, prefix, localName, specified. The `Attr` nodes are created but not directly exposed or managed via `Element.attributes` as a `NamedNodeMap`. |
| 4.10 Interface `CharacterData` | [dom.spec.whatwg.org/#interface-characterdata](https://dom.spec.whatwg.org/#interface-characterdata) | Yes | PartialCompliance |  | Implemented: data, length, substringData, appendData, insertData, deleteData, replaceData. Missing: DOMException throwing for out-of-bounds errors. Methods like `before()`, `after()`, `replaceWith()`, `remove()` are inherited from Node but not specifically tested for CharacterData types. `previousElementSibling` and `nextElementSibling` are not implemented. |
| 4.11 Interface `Text` | [dom.spec.whatwg.org/#interface-text](https://dom.spec.whatwg.org/#interface-text) | Yes | PartialCompliance |  | Implements CharacterData methods. Implemented `splitText()`. Missing: `wholeText` property, `assignedSlot` property. DOMException throwing for `splitText` out-of-bounds offset is not fully spec compliant (currently creates empty new text instead of throwing). |
| 4.12 Interface `CDATASection` | [dom.spec.whatwg.org/#interface-cdatasection](https://dom.spec.whatwg.org/#interface-cdatasection) | Yes | PartialCompliance |  | Basic structure implemented, inherits from CharacterData. Specific CDATASection behavior or properties beyond CharacterData are not explicitly implemented or tested. |
| 4.13 Interface `ProcessingInstruction` | [dom.spec.whatwg.org/#interface-processinginstruction](https://dom.spec.whatwg.org/#interface-processinginstruction) | Yes | PartialCompliance |  | Implemented: `target` property. Inherits from CharacterData. `sheet` property is not applicable/implemented. |
| 4.14 Interface `Comment` | [dom.spec.whatwg.org/#interface-comment](https://dom.spec.whatwg.org/#interface-comment) | Yes | PartialCompliance |  | Basic structure implemented. Inherits from CharacterData (implicitly, as `Comment` embeds `nodeImpl` which holds `nodeValue` used by `CharacterData` methods, though `Comment` itself doesn't directly embed `characterDataImpl`). |
| **5. Ranges** |  |  |  |  |  |
| 5.2 Boundary points | [dom.spec.whatwg.org/#boundary-points](https://dom.spec.whatwg.org/#boundary-points) | No | NotImplemented |  |  |
| 5.3 Interface `AbstractRange` | [dom.spec.whatwg.org/#interface-abstractrange](https://dom.spec.whatwg.org/#interface-abstractrange) | No | NotImplemented |  |  |
| 5.4 Interface `StaticRange` | [dom.spec.whatwg.org/#interface-staticrange](https://dom.spec.whatwg.org/#interface-staticrange) | No | NotImplemented |  |  |
| 5.5 Interface `Range` | [dom.spec.whatwg.org/#interface-range](https://dom.spec.whatwg.org/#interface-range) | No | NotImplemented |  |  |
| **6. Traversal** |  |  |  |  |  |
| 6.1 Interface `NodeIterator` | [dom.spec.whatwg.org/#interface-nodeiterator](https://dom.spec.whatwg.org/#interface-nodeiterator) | No | NotImplemented |  |  |
| 6.2 Interface `TreeWalker` | [dom.spec.whatwg.org/#interface-treewalker](https://dom.spec.whatwg.org/#interface-treewalker) | No | NotImplemented |  |  |
| 6.3 Interface `NodeFilter` | [dom.spec.whatwg.org/#interface-nodefilter](https://dom.spec.whatwg.org/#interface-nodefilter) | No | NotImplemented |  |  |
| **7. Sets** |  |  |  |  |  |
| 7.1 Interface `DOMTokenList` | [dom.spec.whatwg.org/#interface-domtokenlist](https://dom.spec.whatwg.org/#interface-domtokenlist) | No | NotImplemented |  |  |

## Compliance Levels
- **FullCompliance**: Implemented and passes all relevant tests.
- **PartialCompliance**: Partially implemented or some tests failing.
- **IntentionalDeviation**: Implemented differently for specific reasons (e.g., performance, simplification).
- **NotImplemented**: Feature not implemented.
- **NotApplicable**: Feature not applicable to DOMulator's use case.
