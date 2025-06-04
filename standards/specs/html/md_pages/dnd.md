HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 6 User interaction](interaction.html) --- [Table of Contents](./) ---
[6.12 The popover attribute →](popover.html)

1.  ::: {#toc-editing}
    1.  [[6.11]{.secno} Drag and drop](dnd.html#dnd)
        1.  [[6.11.1]{.secno} Introduction](dnd.html#event-drag)
        2.  [[6.11.2]{.secno} The drag data
            store](dnd.html#the-drag-data-store)
        3.  [[6.11.3]{.secno} The `DataTransfer`
            interface](dnd.html#the-datatransfer-interface)
            1.  [[6.11.3.1]{.secno} The `DataTransferItemList`
                interface](dnd.html#the-datatransferitemlist-interface)
            2.  [[6.11.3.2]{.secno} The `DataTransferItem`
                interface](dnd.html#the-datatransferitem-interface)
        4.  [[6.11.4]{.secno} The `DragEvent`
            interface](dnd.html#the-dragevent-interface)
        5.  [[6.11.5]{.secno} Processing
            model](dnd.html#drag-and-drop-processing-model)
        6.  [[6.11.6]{.secno} Events summary](dnd.html#dndevents)
        7.  [[6.11.7]{.secno} The `draggable`
            attribute](dnd.html#the-draggable-attribute)
        8.  [[6.11.8]{.secno} Security risks in the drag-and-drop
            model](dnd.html#security-risks-in-the-drag-and-drop-model)
    :::

### [6.11]{.secno} [Drag and drop]{.dfn}[](#dnd){.self-link} {#dnd}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTML_Drag_and_Drop_API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_Drag_and_Drop_API "HTML Drag and Drop interfaces enable applications to use drag-and-drop features in browsers.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS2+]{.safari_ios
.yes}[Chrome Android18+]{.chrome_android .yes}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

This section defines an event-based drag-and-drop mechanism.

This specification does not define exactly what a *drag-and-drop
operation* actually is.

On a visual medium with a pointing device, a drag operation could be the
default action of a
[`mousedown`{#dnd:event-mousedown}](https://w3c.github.io/uievents/#event-type-mousedown){x-internal="event-mousedown"}
event that is followed by a series of
[`mousemove`{#dnd:event-mousemove}](https://w3c.github.io/uievents/#event-type-mousemove){x-internal="event-mousemove"}
events, and the drop could be triggered by the mouse being released.

When using an input modality other than a pointing device, users would
probably have to explicitly indicate their intention to perform a
drag-and-drop operation, stating what they wish to drag and where they
wish to drop it, respectively.

However it is implemented, drag-and-drop operations must have a starting
point (e.g. where the mouse was clicked, or the start of the selection
or element that was selected for the drag), may have any number of
intermediate steps (elements that the mouse moves over during a drag, or
elements that the user picks as possible drop points as they cycle
through possibilities), and must either have an end point (the element
above which the mouse button was released, or the element that was
finally selected), or be canceled. The end point must be the last
element selected as a possible drop point before the drop occurs (so if
the operation is not canceled, there must be at least one element in the
middle step).

#### [6.11.1]{.secno} Introduction[](#event-drag){.self-link} {#event-drag}

*This section is non-normative.*

To make an element draggable, give the element a
[`draggable`{#event-drag:attr-draggable}](#attr-draggable) attribute,
and set an event listener for
[`dragstart`{#event-drag:event-dnd-dragstart}](#event-dnd-dragstart)
that stores the data being dragged.

The event handler typically needs to check that it\'s not a text
selection that is being dragged, and then needs to store data into the
[`DataTransfer`{#event-drag:datatransfer}](#datatransfer) object and set
the allowed effects (copy, move, link, or some combination).

For example:

``` html
<p>What fruits do you like?</p>
<ol ondragstart="dragStartHandler(event)">
 <li draggable="true" data-value="fruit-apple">Apples</li>
 <li draggable="true" data-value="fruit-orange">Oranges</li>
 <li draggable="true" data-value="fruit-pear">Pears</li>
</ol>
<script>
  var internalDNDType = 'text/x-example'; // set this to something specific to your site
  function dragStartHandler(event) {
    if (event.target instanceof HTMLLIElement) {
      // use the element's data-value="" attribute as the value to be moving:
      event.dataTransfer.setData(internalDNDType, event.target.dataset.value);
      event.dataTransfer.effectAllowed = 'move'; // only allow moves
    } else {
      event.preventDefault(); // don't allow selection to be dragged
    }
  }
</script>
```

------------------------------------------------------------------------

To accept a drop, the drop target has to listen to the following events:

1.  The
    [`dragenter`{#event-drag:event-dnd-dragenter}](#event-dnd-dragenter)
    event handler reports whether or not the drop target is potentially
    willing to accept the drop, by canceling the event.
2.  The
    [`dragover`{#event-drag:event-dnd-dragover}](#event-dnd-dragover)
    event handler specifies what feedback will be shown to the user, by
    setting the
    [`dropEffect`{#event-drag:dom-datatransfer-dropeffect}](#dom-datatransfer-dropeffect)
    attribute of the
    [`DataTransfer`{#event-drag:datatransfer-2}](#datatransfer)
    associated with the event. This event also needs to be canceled.
3.  The [`drop`{#event-drag:event-dnd-drop}](#event-dnd-drop) event
    handler has a final chance to accept or reject the drop. If the drop
    is accepted, the event handler must perform the drop operation on
    the target. This event needs to be canceled, so that the
    [`dropEffect`{#event-drag:dom-datatransfer-dropeffect-2}](#dom-datatransfer-dropeffect)
    attribute\'s value can be used by the source. Otherwise, the drop
    operation is rejected.

For example:

``` html
<p>Drop your favorite fruits below:</p>
<ol ondragenter="dragEnterHandler(event)" ondragover="dragOverHandler(event)"
    ondrop="dropHandler(event)">
</ol>
<script>
  var internalDNDType = 'text/x-example'; // set this to something specific to your site
  function dragEnterHandler(event) {
    var items = event.dataTransfer.items;
    for (var i = 0; i < items.length; ++i) {
      var item = items[i];
      if (item.kind == 'string' && item.type == internalDNDType) {
        event.preventDefault();
        return;
      }
    }
  }
  function dragOverHandler(event) {
    event.dataTransfer.dropEffect = 'move';
    event.preventDefault();
  }
  function dropHandler(event) {
    var li = document.createElement('li');
    var data = event.dataTransfer.getData(internalDNDType);
    if (data == 'fruit-apple') {
      li.textContent = 'Apples';
    } else if (data == 'fruit-orange') {
      li.textContent = 'Oranges';
    } else if (data == 'fruit-pear') {
      li.textContent = 'Pears';
    } else {
      li.textContent = 'Unknown Fruit';
    }
    event.target.appendChild(li);
  }
</script>
```

------------------------------------------------------------------------

To remove the original element (the one that was dragged) from the
display, the
[`dragend`{#event-drag:event-dnd-dragend}](#event-dnd-dragend) event can
be used.

For our example here, that means updating the original markup to handle
that event:

``` html
<p>What fruits do you like?</p>
<ol ondragstart="dragStartHandler(event)" ondragend="dragEndHandler(event)">
 ...as before...
</ol>
<script>
  function dragStartHandler(event) {
    // ...as before...
  }
  function dragEndHandler(event) {
    if (event.dataTransfer.dropEffect == 'move') {
      // remove the dragged element
      event.target.parentNode.removeChild(event.target);
    }
  }
</script>
```

#### [6.11.2]{.secno} The drag data store[](#the-drag-data-store){.self-link}

The data that underlies a drag-and-drop operation, known as the [drag
data store]{#drag-data-store .dfn}, consists of the following
information:

- A [drag data store item list]{#drag-data-store-item-list .dfn}, which
  is a list of items representing the dragged data, each consisting of
  the following information:

  [The drag data item kind]{#the-drag-data-item-kind .dfn}

  :   The kind of data:

      *Text*

      :   Text.

      *File*

      :   Binary data with a filename.

  [The drag data item type string]{#the-drag-data-item-type-string .dfn}

  :   A Unicode string giving the type or format of the data, generally
      given by a [MIME
      type](https://mimesniff.spec.whatwg.org/#mime-type){#the-drag-data-store:mime-type
      x-internal="mime-type"}. Some values that are not [MIME
      types](https://mimesniff.spec.whatwg.org/#mime-type){#the-drag-data-store:mime-type-2
      x-internal="mime-type"} are special-cased for legacy reasons. The
      API does not enforce the use of [MIME
      types](https://mimesniff.spec.whatwg.org/#mime-type){#the-drag-data-store:mime-type-3
      x-internal="mime-type"}; other values can be used as well. In all
      cases, however, the values are all [converted to ASCII
      lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-drag-data-store:converted-to-ascii-lowercase
      x-internal="converted-to-ascii-lowercase"} by the API.

      There is a limit of one *text* item per [item type
      string](#the-drag-data-item-type-string){#the-drag-data-store:the-drag-data-item-type-string}.

  The actual data

  :   A Unicode or binary string, in some cases with a filename (itself
      a Unicode string), as per [the drag data item
      kind](#the-drag-data-item-kind){#the-drag-data-store:the-drag-data-item-kind}.

  The [drag data store item
  list](#drag-data-store-item-list){#the-drag-data-store:drag-data-store-item-list}
  is ordered in the order that the items were added to the list; most
  recently added last.

- The following information, used to generate the UI feedback during the
  drag:

  - User-agent-defined default feedback information, known as the [drag
    data store default feedback]{#drag-data-store-default-feedback
    .dfn}.
  - Optionally, a bitmap image and the coordinate of a point within that
    image, known as the [drag data store bitmap]{#drag-data-store-bitmap
    .dfn} and [drag data store hot spot
    coordinate]{#drag-data-store-hot-spot-coordinate .dfn}.

- A [drag data store mode]{#drag-data-store-mode .dfn}, which is one of
  the following:

  [Read/write mode]{#concept-dnd-rw .dfn}

  :   For the
      [`dragstart`{#the-drag-data-store:event-dnd-dragstart}](#event-dnd-dragstart)
      event. New data can be added to the [drag data
      store](#drag-data-store){#the-drag-data-store:drag-data-store}.

  [Read-only mode]{#concept-dnd-ro .dfn}

  :   For the
      [`drop`{#the-drag-data-store:event-dnd-drop}](#event-dnd-drop)
      event. The list of items representing dragged data can be read,
      including the data. No new data can be added.

  [Protected mode]{#concept-dnd-p .dfn}

  :   For all other events. The formats and kinds in the [drag data
      store](#drag-data-store){#the-drag-data-store:drag-data-store-2}
      list of items representing dragged data can be enumerated, but the
      data itself is unavailable and no new data can be added.

- A [drag data store allowed effects
  state]{#drag-data-store-allowed-effects-state .dfn}, which is a
  string.

When a [drag data
store](#drag-data-store){#the-drag-data-store:drag-data-store-3} is
[created]{#create-a-drag-data-store .dfn}, it must be initialized such
that its [drag data store item
list](#drag-data-store-item-list){#the-drag-data-store:drag-data-store-item-list-2}
is empty, it has no [drag data store default
feedback](#drag-data-store-default-feedback){#the-drag-data-store:drag-data-store-default-feedback},
it has no [drag data store
bitmap](#drag-data-store-bitmap){#the-drag-data-store:drag-data-store-bitmap}
and [drag data store hot spot
coordinate](#drag-data-store-hot-spot-coordinate){#the-drag-data-store:drag-data-store-hot-spot-coordinate},
its [drag data store
mode](#drag-data-store-mode){#the-drag-data-store:drag-data-store-mode}
is [protected mode](#concept-dnd-p){#the-drag-data-store:concept-dnd-p},
and its [drag data store allowed effects
state](#drag-data-store-allowed-effects-state){#the-drag-data-store:drag-data-store-allowed-effects-state}
is the string
\"[`uninitialized`{#the-drag-data-store:dom-datatransfer-effectallowed-uninitialized}](#dom-datatransfer-effectallowed-uninitialized)\".

#### [6.11.3]{.secno} The [`DataTransfer`{#the-datatransfer-interface:datatransfer}](#datatransfer) interface[](#the-datatransfer-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DataTransfer](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer "The DataTransfer object is used to hold the data that is being dragged during a drag and drop operation. It may hold one or more data items, each of one or more data types. For more information about drag and drop, see HTML Drag and Drop API.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`DataTransfer`{#the-datatransfer-interface:datatransfer-2}](#datatransfer)
objects are used to expose the [drag data
store](#drag-data-store){#the-datatransfer-interface:drag-data-store}
that underlies a drag-and-drop operation.

``` idl
[Exposed=Window]
interface DataTransfer {
  constructor();

  attribute DOMString dropEffect;
  attribute DOMString effectAllowed;

  [SameObject] readonly attribute DataTransferItemList items;

  undefined setDragImage(Element image, long x, long y);

  /* old interface */
  readonly attribute FrozenArray<DOMString> types;
  DOMString getData(DOMString format);
  undefined setData(DOMString format, DOMString data);
  undefined clearData(optional DOMString format);
  [SameObject] readonly attribute FileList files;
};
```

`dataTransfer`{.variable}` = new `[`DataTransfer`](#dom-datatransfer){#dom-datatransfer-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/DataTransfer](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/DataTransfer "The DataTransfer constructor creates a new DataTransfer object instance.")

Support in all current engines.

::: support
[Firefox62+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome59+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet8.0+]{.samsunginternet_android .yes}[Opera
Android44+]{.opera_android .yes}
:::
::::
:::::

Creates a new
[`DataTransfer`{#the-datatransfer-interface:datatransfer-3}](#datatransfer)
object with an empty [drag data
store](#drag-data-store){#the-datatransfer-interface:drag-data-store-2}.

`dataTransfer`{.variable}`.`[`dropEffect`](#dom-datatransfer-dropeffect){#dom-datatransfer-dropeffect-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/dropEffect](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/dropEffect "The DataTransfer.dropEffect property controls the feedback (typically visual) the user is given during a drag and drop operation. It will affect which cursor is displayed while dragging. For example, when the user hovers over a target drop element, the browser's cursor may indicate which type of operation will occur.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the kind of operation that is currently selected. If the kind of
operation isn\'t one of those that is allowed by the
[`effectAllowed`{#the-datatransfer-interface:dom-datatransfer-effectallowed-2}](#dom-datatransfer-effectallowed)
attribute, then the operation will fail.

Can be set, to change the selected operation.

The possible values are
\"[`none`{#the-datatransfer-interface:dom-datatransfer-dropeffect-none}](#dom-datatransfer-dropeffect-none)\",
\"[`copy`{#the-datatransfer-interface:dom-datatransfer-dropeffect-copy}](#dom-datatransfer-dropeffect-copy)\",
\"[`link`{#the-datatransfer-interface:dom-datatransfer-dropeffect-link}](#dom-datatransfer-dropeffect-link)\",
and
\"[`move`{#the-datatransfer-interface:dom-datatransfer-dropeffect-move}](#dom-datatransfer-dropeffect-move)\".

`dataTransfer`{.variable}`.`[`effectAllowed`](#dom-datatransfer-effectallowed){#dom-datatransfer-effectallowed-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/effectAllowed](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/effectAllowed "The DataTransfer.effectAllowed property specifies the effect that is allowed for a drag operation. The copy operation is used to indicate that the data being dragged will be copied from its present location to the drop location. The move operation is used to indicate that the data being dragged will be moved, and the link operation is used to indicate that some form of relationship or connection will be created between the source and drop locations.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the kinds of operations that are to be allowed.

Can be set (during the
[`dragstart`{#the-datatransfer-interface:event-dnd-dragstart}](#event-dnd-dragstart)
event), to change the allowed operations.

The possible values are
\"[`none`{#the-datatransfer-interface:dom-datatransfer-effectallowed-none}](#dom-datatransfer-effectallowed-none)\",
\"[`copy`{#the-datatransfer-interface:dom-datatransfer-effectallowed-copy}](#dom-datatransfer-effectallowed-copy)\",
\"[`copyLink`{#the-datatransfer-interface:dom-datatransfer-effectallowed-copylink}](#dom-datatransfer-effectallowed-copylink)\",
\"[`copyMove`{#the-datatransfer-interface:dom-datatransfer-effectallowed-copymove}](#dom-datatransfer-effectallowed-copymove)\",
\"[`link`{#the-datatransfer-interface:dom-datatransfer-effectallowed-link}](#dom-datatransfer-effectallowed-link)\",
\"[`linkMove`{#the-datatransfer-interface:dom-datatransfer-effectallowed-linkmove}](#dom-datatransfer-effectallowed-linkmove)\",
\"[`move`{#the-datatransfer-interface:dom-datatransfer-effectallowed-move}](#dom-datatransfer-effectallowed-move)\",
\"[`all`{#the-datatransfer-interface:dom-datatransfer-effectallowed-all}](#dom-datatransfer-effectallowed-all)\",
and
\"[`uninitialized`{#the-datatransfer-interface:dom-datatransfer-effectallowed-uninitialized}](#dom-datatransfer-effectallowed-uninitialized)\",

`dataTransfer`{.variable}`.`[`items`](#dom-datatransfer-items){#dom-datatransfer-items-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/items](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/items "The read-only DataTransfer property items property is a list of the data transfer items in a drag operation. The list includes one item for each item in the operation and if the operation had no items, the list is empty.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari11.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android52+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`DataTransferItemList`{#the-datatransfer-interface:datatransferitemlist-2}](#datatransferitemlist)
object, with the drag data.

`dataTransfer`{.variable}`.`[`setDragImage`](#dom-datatransfer-setdragimage){#dom-datatransfer-setdragimage-dev}`(``element`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/setDragImage](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/setDragImage "When a drag occurs, a translucent image is generated from the drag target (the element the dragstart event is fired at), and follows the mouse pointer during the drag. This image is created automatically, so you do not need to create it yourself. However, if a custom image is desired, the DataTransfer.setDragImage() method can be used to set the custom image to be used. The image will typically be an <img> element but it can also be a <canvas> or any other visible element.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Uses the given element to update the drag feedback, replacing any
previously specified feedback.

`dataTransfer`{.variable}`.`[`types`](#dom-datatransfer-types){#dom-datatransfer-types-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/types](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/types "The DataTransfer.types read-only property returns the available types that exist in the items.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a [frozen
array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#the-datatransfer-interface:frozen-array
x-internal="frozen-array"} listing the formats that were set in the
[`dragstart`{#the-datatransfer-interface:event-dnd-dragstart-2}](#event-dnd-dragstart)
event. In addition, if any files are being dragged, then one of the
types will be the string \"`Files`\".

`data`{.variable}` = ``dataTransfer`{.variable}`.`[`getData`](#dom-datatransfer-getdata){#dom-datatransfer-getdata-dev}`(``format`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/getData](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/getData "The DataTransfer.getData() method retrieves drag data (as a string) for the specified type. If the drag operation does not include data, this method returns an empty string.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the specified data. If there is no such data, returns the empty
string.

`dataTransfer`{.variable}`.`[`setData`](#dom-datatransfer-setdata){#dom-datatransfer-setdata-dev}`(``format`{.variable}`, ``data`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/setData](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/setData "The DataTransfer.setData() method sets the drag operation's drag data to the specified data and type. If data for the given type does not exist, it is added at the end of the drag data store, such that the last item in the types list will be the new type. If data for the given type already exists, the existing data is replaced in the same position. That is, the order of the types list is not changed when replacing data of the same type.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Adds the specified data.

`dataTransfer`{.variable}`.`[`clearData`](#dom-datatransfer-cleardata){#dom-datatransfer-cleardata-dev}`([ ``format`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/clearData](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/clearData "The DataTransfer.clearData() method removes the drag operation's drag data for the given type. If data for the given type does not exist, this method does nothing.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Removes the data of the specified formats. Removes all data if the
argument is omitted.

`dataTransfer`{.variable}`.`[`files`](#dom-datatransfer-files){#dom-datatransfer-files-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransfer/files](https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer/files "The files property of DataTransfer objects is a list of the files in the drag operation. If the operation includes no files, the list is empty.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`FileList`{#the-datatransfer-interface:filelist-2}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
of the files being dragged, if any.

[`DataTransfer`{#the-datatransfer-interface:datatransfer-4}](#datatransfer)
objects that are created as part of [drag-and-drop events](#dndevents)
are only valid while those events are being fired.

A
[`DataTransfer`{#the-datatransfer-interface:datatransfer-5}](#datatransfer)
object is associated with a [drag data
store](#drag-data-store){#the-datatransfer-interface:drag-data-store-3}
while it is valid.

A
[`DataTransfer`{#the-datatransfer-interface:datatransfer-6}](#datatransfer)
object has an associated [types array]{#concept-datatransfer-types
.dfn}, which is a
[`FrozenArray<DOMString>`{#the-datatransfer-interface:frozen-array-2}](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){x-internal="frozen-array"},
initially empty. When the contents of the
[`DataTransfer`{#the-datatransfer-interface:datatransfer-7}](#datatransfer)
object\'s [drag data store item
list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list}
change, or when the
[`DataTransfer`{#the-datatransfer-interface:datatransfer-8}](#datatransfer)
object becomes no longer associated with a [drag data
store](#drag-data-store){#the-datatransfer-interface:drag-data-store-4},
run the following steps:

1.  Let `L`{.variable} be an empty sequence.

2.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-9}](#datatransfer)
    object is still associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-5},
    then:

    1.  For each item in the
        [`DataTransfer`{#the-datatransfer-interface:datatransfer-10}](#datatransfer)
        object\'s [drag data store item
        list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-2}
        whose
        [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind}
        is *text*, add an entry to `L`{.variable} consisting of the
        item\'s [type
        string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string}.

    2.  If there are any items in the
        [`DataTransfer`{#the-datatransfer-interface:datatransfer-11}](#datatransfer)
        object\'s [drag data store item
        list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-3}
        whose
        [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-2}
        is *File*, then add an entry to `L`{.variable} consisting of the
        string \"`Files`\". (This value can be distinguished from the
        other values because it is not lowercase.)

3.  Set the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-12}](#datatransfer)
    object\'s [types
    array](#concept-datatransfer-types){#the-datatransfer-interface:concept-datatransfer-types}
    to the result of [creating a frozen
    array](https://webidl.spec.whatwg.org/#dfn-create-frozen-array){#the-datatransfer-interface:creating-a-frozen-array
    x-internal="creating-a-frozen-array"} from `L`{.variable}.

The [`DataTransfer()`]{#dom-datatransfer .dfn dfn-for="DataTransfer"
dfn-type="constructor"} constructor, when invoked, must return a newly
created
[`DataTransfer`{#the-datatransfer-interface:datatransfer-13}](#datatransfer)
object initialized as follows:

1.  Set the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-6}\'s
    [item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-4}
    to be an empty list.

2.  Set the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-7}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode}
    to [read/write
    mode](#concept-dnd-rw){#the-datatransfer-interface:concept-dnd-rw}.

3.  Set the
    [`dropEffect`{#the-datatransfer-interface:dom-datatransfer-dropeffect-2}](#dom-datatransfer-dropeffect)
    and
    [`effectAllowed`{#the-datatransfer-interface:dom-datatransfer-effectallowed-3}](#dom-datatransfer-effectallowed)
    to \"none\".

The [`dropEffect`]{#dom-datatransfer-dropeffect .dfn
dfn-for="DataTransfer" dfn-type="attribute"} attribute controls the
drag-and-drop feedback that the user is given during a drag-and-drop
operation. When the
[`DataTransfer`{#the-datatransfer-interface:datatransfer-14}](#datatransfer)
object is created, the
[`dropEffect`{#the-datatransfer-interface:dom-datatransfer-dropeffect-3}](#dom-datatransfer-dropeffect)
attribute is set to a string value. On getting, it must return its
current value. On setting, if the new value is one of
\"[`none`]{#dom-datatransfer-dropeffect-none .dfn}\",
\"[`copy`]{#dom-datatransfer-dropeffect-copy .dfn}\",
\"[`link`]{#dom-datatransfer-dropeffect-link .dfn}\", or
\"[`move`]{#dom-datatransfer-dropeffect-move .dfn}\", then the
attribute\'s current value must be set to the new value. Other values
must be ignored.

The [`effectAllowed`]{#dom-datatransfer-effectallowed .dfn
dfn-for="DataTransfer" dfn-type="attribute"} attribute is used in the
drag-and-drop processing model to initialize the
[`dropEffect`{#the-datatransfer-interface:dom-datatransfer-dropeffect-4}](#dom-datatransfer-dropeffect)
attribute during the
[`dragenter`{#the-datatransfer-interface:event-dnd-dragenter}](#event-dnd-dragenter)
and
[`dragover`{#the-datatransfer-interface:event-dnd-dragover}](#event-dnd-dragover)
events. When the
[`DataTransfer`{#the-datatransfer-interface:datatransfer-15}](#datatransfer)
object is created, the
[`effectAllowed`{#the-datatransfer-interface:dom-datatransfer-effectallowed-4}](#dom-datatransfer-effectallowed)
attribute is set to a string value. On getting, it must return its
current value. On setting, if the [drag data
store](#drag-data-store){#the-datatransfer-interface:drag-data-store-8}\'s
[mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-2}
is the [read/write
mode](#concept-dnd-rw){#the-datatransfer-interface:concept-dnd-rw-2} and
the new value is one of \"[`none`]{#dom-datatransfer-effectallowed-none
.dfn}\", \"[`copy`]{#dom-datatransfer-effectallowed-copy .dfn}\",
\"[`copyLink`]{#dom-datatransfer-effectallowed-copylink .dfn}\",
\"[`copyMove`]{#dom-datatransfer-effectallowed-copymove .dfn}\",
\"[`link`]{#dom-datatransfer-effectallowed-link .dfn}\",
\"[`linkMove`]{#dom-datatransfer-effectallowed-linkmove .dfn}\",
\"[`move`]{#dom-datatransfer-effectallowed-move .dfn}\",
\"[`all`]{#dom-datatransfer-effectallowed-all .dfn}\", or
\"[`uninitialized`]{#dom-datatransfer-effectallowed-uninitialized
.dfn}\", then the attribute\'s current value must be set to the new
value. Otherwise, it must be left unchanged.

The [`items`]{#dom-datatransfer-items .dfn dfn-for="DataTransfer"
dfn-type="attribute"} attribute must return a
[`DataTransferItemList`{#the-datatransfer-interface:datatransferitemlist-3}](#datatransferitemlist)
object associated with the
[`DataTransfer`{#the-datatransfer-interface:datatransfer-16}](#datatransfer)
object.

The
[`setDragImage(``image`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`]{#dom-datatransfer-setdragimage
.dfn dfn-for="DataTransfer" dfn-type="method"} method must run the
following steps:

1.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-17}](#datatransfer)
    object is no longer associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-9},
    return. Nothing happens.

2.  If the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-10}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-3}
    is not the [read/write
    mode](#concept-dnd-rw){#the-datatransfer-interface:concept-dnd-rw-3},
    return. Nothing happens.

3.  If `image`{.variable} is an
    [`img`{#the-datatransfer-interface:the-img-element}](embedded-content.html#the-img-element)
    element, then set the [drag data store
    bitmap](#drag-data-store-bitmap){#the-datatransfer-interface:drag-data-store-bitmap}
    to the element\'s image (at its [natural
    size](https://drafts.csswg.org/css-images/#natural-dimensions){#the-datatransfer-interface:natural-dimensions
    x-internal="natural-dimensions"}); otherwise, set the [drag data
    store
    bitmap](#drag-data-store-bitmap){#the-datatransfer-interface:drag-data-store-bitmap-2}
    to an image generated from the given element (the exact mechanism
    for doing so is not currently specified).

4.  Set the [drag data store hot spot
    coordinate](#drag-data-store-hot-spot-coordinate){#the-datatransfer-interface:drag-data-store-hot-spot-coordinate}
    to the given `x`{.variable}, `y`{.variable} coordinate.

The [`types`]{#dom-datatransfer-types .dfn dfn-for="DataTransfer"
dfn-type="attribute"} attribute must return this
[`DataTransfer`{#the-datatransfer-interface:datatransfer-18}](#datatransfer)
object\'s [types
array](#concept-datatransfer-types){#the-datatransfer-interface:concept-datatransfer-types-2}.

The [`getData(``format`{.variable}`)`]{#dom-datatransfer-getdata .dfn
dfn-for="DataTransfer" dfn-type="method"} method must run the following
steps:

1.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-19}](#datatransfer)
    object is no longer associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-11},
    then return the empty string.

2.  If the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-12}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-4}
    is the [protected
    mode](#concept-dnd-p){#the-datatransfer-interface:concept-dnd-p},
    then return the empty string.

3.  Let `format`{.variable} be the first argument, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransfer-interface:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

4.  Let `convert-to-URL`{.variable} be false.

5.  If `format`{.variable} equals \"`text`\", change it to
    \"`text/plain`\".

6.  If `format`{.variable} equals \"`url`\", change it to
    \"`text/uri-list`\" and set `convert-to-URL`{.variable} to true.

7.  If there is no item in the [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-5}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-3}
    is *text* and whose [type
    string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-2}
    is equal to `format`{.variable}, return the empty string.

8.  Let `result`{.variable} be the data of the item in the [drag data
    store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-6}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-4}
    is *Plain Unicode string* and whose [type
    string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-3}
    is equal to `format`{.variable}.

9.  If `convert-to-URL`{.variable} is true, then parse
    `result`{.variable} as appropriate for `text/uri-list` data, and
    then set `result`{.variable} to the first URL from the list, if any,
    or the empty string otherwise.
    [\[RFC2483\]](references.html#refsRFC2483)

10. Return `result`{.variable}.

The
[`setData(``format`{.variable}`, ``data`{.variable}`)`]{#dom-datatransfer-setdata
.dfn dfn-for="DataTransfer" dfn-type="method"} method must run the
following steps:

1.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-20}](#datatransfer)
    object is no longer associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-13},
    return. Nothing happens.

2.  If the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-14}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-5}
    is not the [read/write
    mode](#concept-dnd-rw){#the-datatransfer-interface:concept-dnd-rw-4},
    return. Nothing happens.

3.  Let `format`{.variable} be the first argument, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransfer-interface:converted-to-ascii-lowercase-2
    x-internal="converted-to-ascii-lowercase"}.

4.  If `format`{.variable} equals \"`text`\", change it to
    \"`text/plain`\".

    If `format`{.variable} equals \"`url`\", change it to
    \"`text/uri-list`\".

5.  Remove the item in the [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-7}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-5}
    is *text* and whose [type
    string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-4}
    is equal to `format`{.variable}, if there is one.

6.  Add an item to the [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-8}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-6}
    is *text*, whose [type
    string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-5}
    is equal to `format`{.variable}, and whose data is the string given
    by the method\'s second argument.

The [`clearData(``format`{.variable}`)`]{#dom-datatransfer-cleardata
.dfn dfn-for="DataTransfer" dfn-type="method"} method must run the
following steps:

1.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-21}](#datatransfer)
    object is no longer associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-15},
    return. Nothing happens.

2.  If the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-16}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-6}
    is not the [read/write
    mode](#concept-dnd-rw){#the-datatransfer-interface:concept-dnd-rw-5},
    return. Nothing happens.

3.  If the method was called with no arguments, remove each item in the
    [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-9}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-7}
    is *Plain Unicode string*, and return.

4.  Set `format`{.variable} to `format`{.variable}, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransfer-interface:converted-to-ascii-lowercase-3
    x-internal="converted-to-ascii-lowercase"}.

5.  If `format`{.variable} equals \"`text`\", change it to
    \"`text/plain`\".

    If `format`{.variable} equals \"`url`\", change it to
    \"`text/uri-list`\".

6.  Remove the item in the [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-10}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-8}
    is *text* and whose [type
    string](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-6}
    is equal to `format`{.variable}, if there is one.

The
[`clearData()`{#the-datatransfer-interface:dom-datatransfer-cleardata-2}](#dom-datatransfer-cleardata)
method does not affect whether any files were included in the drag, so
the
[`types`{#the-datatransfer-interface:dom-datatransfer-types-2}](#dom-datatransfer-types)
attribute\'s list might still not be empty after calling
[`clearData()`{#the-datatransfer-interface:dom-datatransfer-cleardata-3}](#dom-datatransfer-cleardata)
(it would still contain the \"`Files`\" string if any files were
included in the drag).

The [`files`]{#dom-datatransfer-files .dfn dfn-for="DataTransfer"
dfn-type="attribute"} attribute must return a
[live](infrastructure.html#live){#the-datatransfer-interface:live}
[`FileList`{#the-datatransfer-interface:filelist-3}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
sequence consisting of
[`File`{#the-datatransfer-interface:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
objects representing the files found by the following steps.
Furthermore, for a given
[`FileList`{#the-datatransfer-interface:filelist-4}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
object and a given underlying file, the same
[`File`{#the-datatransfer-interface:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
object must be used each time.

1.  Start with an empty list `L`{.variable}.

2.  If the
    [`DataTransfer`{#the-datatransfer-interface:datatransfer-22}](#datatransfer)
    object is no longer associated with a [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-17},
    the
    [`FileList`{#the-datatransfer-interface:filelist-5}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
    is empty. Return the empty list `L`{.variable}.

3.  If the [drag data
    store](#drag-data-store){#the-datatransfer-interface:drag-data-store-18}\'s
    [mode](#drag-data-store-mode){#the-datatransfer-interface:drag-data-store-mode-7}
    is the [protected
    mode](#concept-dnd-p){#the-datatransfer-interface:concept-dnd-p-2},
    return the empty list `L`{.variable}.

4.  For each item in the [drag data store item
    list](#drag-data-store-item-list){#the-datatransfer-interface:drag-data-store-item-list-11}
    whose
    [kind](#the-drag-data-item-kind){#the-datatransfer-interface:the-drag-data-item-kind-9}
    is *File*, add the item\'s data (the file, in particular its name
    and contents, as well as its
    [type](#the-drag-data-item-type-string){#the-datatransfer-interface:the-drag-data-item-type-string-7})
    to the list `L`{.variable}.

5.  The files found by these steps are those in the list `L`{.variable}.

This version of the API does not expose the types of the files during
the drag.

##### [6.11.3.1]{.secno} The [`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist}](#datatransferitemlist) interface[](#the-datatransferitemlist-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DataTransferItemList](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList "The DataTransferItemList object is a list of DataTransferItem objects representing items being dragged. During a drag operation, each DragEvent has a dataTransfer property and that property is a DataTransferItemList.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Each
[`DataTransfer`{#the-datatransferitemlist-interface:datatransfer}](#datatransfer)
object is associated with a
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-2}](#datatransferitemlist)
object.

``` idl
[Exposed=Window]
interface DataTransferItemList {
  readonly attribute unsigned long length;
  getter DataTransferItem (unsigned long index);
  DataTransferItem? add(DOMString data, DOMString type);
  DataTransferItem? add(File data);
  undefined remove(unsigned long index);
  undefined clear();
};
```

`items`{.variable}`.`[`length`](#dom-datatransferitemlist-length){#dom-datatransferitemlist-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItemList/length](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList/length "The read-only length property of the DataTransferItemList interface returns the number of items currently in the drag item list.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Returns the number of items in the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store}.

`items`{.variable}`[``index`{.variable}`]`

Returns the
[`DataTransferItem`{#the-datatransferitemlist-interface:datatransferitem-4}](#datatransferitem)
object representing the `index`{.variable}th entry in the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-2}.

`items`{.variable}`.`[`remove`](#dom-datatransferitemlist-remove){#dom-datatransferitemlist-remove-dev}`(``index`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItemList/remove](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList/remove "The DataTransferItemList.remove() method removes the DataTransferItem at the specified index from the list. If the index is less than zero or greater than one less than the length of the list, the list will not be changed.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome31+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Removes the `index`{.variable}th entry in the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-3}.

`items`{.variable}`.`[`clear`](#dom-datatransferitemlist-clear){#dom-datatransferitemlist-clear-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItemList/clear](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList/clear "The DataTransferItemList method clear() removes all DataTransferItem objects from the drag data items list, leaving the list empty.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Removes all the entries in the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-4}.

`items`{.variable}`.`[`add`](#dom-datatransferitemlist-add){#dom-datatransferitemlist-add-dev}`(``data`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItemList/add](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList/add "The DataTransferItemList.add() method creates a new DataTransferItem using the specified data and adds it to the drag data list. The item may be a File or a string of a given type. If the item is successfully added to the list, the newly-created DataTransferItem object is returned.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

`items`{.variable}`.`[`add`](#dom-datatransferitemlist-add){#the-datatransferitemlist-interface:dom-datatransferitemlist-add-3}`(``data`{.variable}`, ``type`{.variable}`)`

Adds a new entry for the given data to the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-5}.
If the data is plain text then a `type`{.variable} string has to be
provided also.

While the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-3}](#datatransferitemlist)
object\'s
[`DataTransfer`{#the-datatransferitemlist-interface:datatransfer-2}](#datatransfer)
object is associated with a [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-6},
the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-4}](#datatransferitemlist)
object\'s *mode* is the same as the [drag data store
mode](#drag-data-store-mode){#the-datatransferitemlist-interface:drag-data-store-mode}.
When the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-5}](#datatransferitemlist)
object\'s
[`DataTransfer`{#the-datatransferitemlist-interface:datatransfer-3}](#datatransfer)
object is *not* associated with a [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-7},
the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-6}](#datatransferitemlist)
object\'s *mode* is the *disabled mode*. The [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-8}
referenced in this section (which is used only when the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-7}](#datatransferitemlist)
object is not in the *disabled mode*) is the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-9}
with which the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-8}](#datatransferitemlist)
object\'s
[`DataTransfer`{#the-datatransferitemlist-interface:datatransfer-4}](#datatransfer)
object is associated.

The [`length`]{#dom-datatransferitemlist-length .dfn
dfn-for="DataTransferItemList" dfn-type="attribute"} attribute must
return zero if the object is in the *disabled mode*; otherwise it must
return the number of items in the [drag data store item
list](#drag-data-store-item-list){#the-datatransferitemlist-interface:drag-data-store-item-list}.

When a
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-9}](#datatransferitemlist)
object is not in the *disabled mode*, its [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-datatransferitemlist-interface:supported-property-indices
x-internal="supported-property-indices"} are the
[indices](https://infra.spec.whatwg.org/#list-get-the-indices){#the-datatransferitemlist-interface:indices
x-internal="indices"} of the [drag data store item
list](#drag-data-store-item-list){#the-datatransferitemlist-interface:drag-data-store-item-list-2}.

To [determine the value of an indexed
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property){#the-datatransferitemlist-interface:determine-the-value-of-an-indexed-property
x-internal="determine-the-value-of-an-indexed-property"} `i`{.variable}
of a
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-10}](#datatransferitemlist)
object, the user agent must return a
[`DataTransferItem`{#the-datatransferitemlist-interface:datatransferitem-5}](#datatransferitem)
object representing the `i`{.variable}th item in the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-10}.
The same object must be returned each time a particular item is obtained
from this
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-11}](#datatransferitemlist)
object. The
[`DataTransferItem`{#the-datatransferitemlist-interface:datatransferitem-6}](#datatransferitem)
object must be associated with the same
[`DataTransfer`{#the-datatransferitemlist-interface:datatransfer-5}](#datatransfer)
object as the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-12}](#datatransferitemlist)
object when it is first created.

The [`add()`]{#dom-datatransferitemlist-add .dfn
dfn-for="DataTransferItemList" dfn-type="method"} method must run the
following steps:

1.  If the
    [`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-13}](#datatransferitemlist)
    object is not in the *[read/write mode](#concept-dnd-rw)*, return
    null.

2.  Jump to the appropriate set of steps from the following list:

    If the first argument to the method is a string

    :   If there is already an item in the [drag data store item
        list](#drag-data-store-item-list){#the-datatransferitemlist-interface:drag-data-store-item-list-3}
        whose
        [kind](#the-drag-data-item-kind){#the-datatransferitemlist-interface:the-drag-data-item-kind}
        is *text* and whose [type
        string](#the-drag-data-item-type-string){#the-datatransferitemlist-interface:the-drag-data-item-type-string}
        is equal to the value of the method\'s second argument,
        [converted to ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransferitemlist-interface:converted-to-ascii-lowercase
        x-internal="converted-to-ascii-lowercase"}, then throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-datatransferitemlist-interface:notsupportederror
        x-internal="notsupportederror"}
        [`DOMException`{#the-datatransferitemlist-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        Otherwise, add an item to the [drag data store item
        list](#drag-data-store-item-list){#the-datatransferitemlist-interface:drag-data-store-item-list-4}
        whose
        [kind](#the-drag-data-item-kind){#the-datatransferitemlist-interface:the-drag-data-item-kind-2}
        is *text*, whose [type
        string](#the-drag-data-item-type-string){#the-datatransferitemlist-interface:the-drag-data-item-type-string-2}
        is equal to the value of the method\'s second argument,
        [converted to ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransferitemlist-interface:converted-to-ascii-lowercase-2
        x-internal="converted-to-ascii-lowercase"}, and whose data is
        the string given by the method\'s first argument.

    If the first argument to the method is a [`File`{#the-datatransferitemlist-interface:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}

    :   Add an item to the [drag data store item
        list](#drag-data-store-item-list){#the-datatransferitemlist-interface:drag-data-store-item-list-5}
        whose
        [kind](#the-drag-data-item-kind){#the-datatransferitemlist-interface:the-drag-data-item-kind-3}
        is *File*, whose [type
        string](#the-drag-data-item-type-string){#the-datatransferitemlist-interface:the-drag-data-item-type-string-3}
        is the
        [`type`{#the-datatransferitemlist-interface:dom-blob-type}](https://w3c.github.io/FileAPI/#dfn-type){x-internal="dom-blob-type"}
        of the
        [`File`{#the-datatransferitemlist-interface:file-3}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"},
        [converted to ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-datatransferitemlist-interface:converted-to-ascii-lowercase-3
        x-internal="converted-to-ascii-lowercase"}, and whose data is
        the same as the
        [`File`{#the-datatransferitemlist-interface:file-4}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}\'s
        data.

3.  [Determine the value of the indexed
    property](#dom-datatransferitemlist-item) corresponding to the newly
    added item, and return that value (a newly created
    [`DataTransferItem`{#the-datatransferitemlist-interface:datatransferitem-7}](#datatransferitem)
    object).

The [`remove(``index`{.variable}`)`]{#dom-datatransferitemlist-remove
.dfn dfn-for="DataTransferItemList" dfn-type="method"} method must run
these steps:

1.  If the
    [`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-14}](#datatransferitemlist)
    object is not in the *[read/write mode](#concept-dnd-rw)*, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-datatransferitemlist-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-datatransferitemlist-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If the [drag data
    store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-11}
    does not contain an `index`{.variable}th item, then return.

3.  Remove the `index`{.variable}th item from the [drag data
    store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-12}.

The [`clear()`]{#dom-datatransferitemlist-clear .dfn
dfn-for="DataTransferItemList" dfn-type="method"} method, if the
[`DataTransferItemList`{#the-datatransferitemlist-interface:datatransferitemlist-15}](#datatransferitemlist)
object is in the *[read/write mode](#concept-dnd-rw)*, must remove all
the items from the [drag data
store](#drag-data-store){#the-datatransferitemlist-interface:drag-data-store-13}.
Otherwise, it must do nothing.

##### [6.11.3.2]{.secno} The [`DataTransferItem`{#the-datatransferitem-interface:datatransferitem}](#datatransferitem) interface[](#the-datatransferitem-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DataTransferItem](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem "The DataTransferItem object represents one drag data item. During a drag operation, each drag event has a dataTransfer property which contains a list of drag data items. Each item in the list is a DataTransferItem object.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome11+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Each
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-2}](#datatransferitem)
object is associated with a
[`DataTransfer`{#the-datatransferitem-interface:datatransfer}](#datatransfer)
object.

``` idl
[Exposed=Window]
interface DataTransferItem {
  readonly attribute DOMString kind;
  readonly attribute DOMString type;
  undefined getAsString(FunctionStringCallback? _callback);
  File? getAsFile();
};

callback FunctionStringCallback = undefined (DOMString data);
```

`item`{.variable}`.`[`kind`](#dom-datatransferitem-kind){#dom-datatransferitem-kind-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItem/kind](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem/kind "The read-only DataTransferItem.kind property returns a DataTransferItem representing the drag data item kind: some text or some file.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome11+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Returns [the drag data item
kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind},
one of: \"string\", \"file\".

`item`{.variable}`.`[`type`](#dom-datatransferitem-type){#dom-datatransferitem-type-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItem/type](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem/type "The read-only DataTransferItem.type property returns the type (format) of the DataTransferItem object representing the drag data item. The type is a Unicode string generally given by a MIME type, although a MIME type is not required.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome11+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Returns [the drag data item type
string](#the-drag-data-item-type-string){#the-datatransferitem-interface:the-drag-data-item-type-string}.

`item`{.variable}`.`[`getAsString`](#dom-datatransferitem-getasstring){#dom-datatransferitem-getasstring-dev}`(``callback`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItem/getAsString](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem/getAsString "The DataTransferItem.getAsString() method invokes the given callback with the drag data item's string data as the argument if the item's kind is a Plain unicode string (i.e. kind is string).")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome11+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Invokes the callback with the string data as the argument, if [the drag
data item
kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind-2}
is *text*.

`file`{.variable}` = ``item`{.variable}`.`[`getAsFile`](#dom-datatransferitem-getasfile){#dom-datatransferitem-getasfile-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DataTransferItem/getAsFile](https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem/getAsFile "If the item is a file, the DataTransferItem.getAsFile() method returns the drag data item's File object. If the item is not a file, this method returns null.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome11+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`File`{#the-datatransferitem-interface:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
object, if [the drag data item
kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind-3}
is *File*.

While the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-3}](#datatransferitem)
object\'s
[`DataTransfer`{#the-datatransferitem-interface:datatransfer-2}](#datatransfer)
object is associated with a [drag data
store](#drag-data-store){#the-datatransferitem-interface:drag-data-store}
and that [drag data
store](#drag-data-store){#the-datatransferitem-interface:drag-data-store-2}\'s
[drag data store item
list](#drag-data-store-item-list){#the-datatransferitem-interface:drag-data-store-item-list}
still contains the item that the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-4}](#datatransferitem)
object represents, the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-5}](#datatransferitem)
object\'s *mode* is the same as the [drag data store
mode](#drag-data-store-mode){#the-datatransferitem-interface:drag-data-store-mode}.
When the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-6}](#datatransferitem)
object\'s
[`DataTransfer`{#the-datatransferitem-interface:datatransfer-3}](#datatransfer)
object is *not* associated with a [drag data
store](#drag-data-store){#the-datatransferitem-interface:drag-data-store-3},
or if the item that the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-7}](#datatransferitem)
object represents has been removed from the relevant [drag data store
item
list](#drag-data-store-item-list){#the-datatransferitem-interface:drag-data-store-item-list-2},
the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-8}](#datatransferitem)
object\'s *mode* is the *disabled mode*. The [drag data
store](#drag-data-store){#the-datatransferitem-interface:drag-data-store-4}
referenced in this section (which is used only when the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-9}](#datatransferitem)
object is not in the *disabled mode*) is the [drag data
store](#drag-data-store){#the-datatransferitem-interface:drag-data-store-5}
with which the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-10}](#datatransferitem)
object\'s
[`DataTransfer`{#the-datatransferitem-interface:datatransfer-4}](#datatransfer)
object is associated.

The [`kind`]{#dom-datatransferitem-kind .dfn dfn-for="DataTransferItem"
dfn-type="attribute"} attribute must return the empty string if the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-11}](#datatransferitem)
object is in the *disabled mode*; otherwise it must return the string
given in the cell from the second column of the following table from the
row whose cell in the first column contains [the drag data item
kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind-4}
of the item represented by the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-12}](#datatransferitem)
object:

Kind

String

*Text*

\"`string`\"

*File*

\"`file`\"

The [`type`]{#dom-datatransferitem-type .dfn dfn-for="DataTransferItem"
dfn-type="attribute"} attribute must return the empty string if the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-13}](#datatransferitem)
object is in the *disabled mode*; otherwise it must return [the drag
data item type
string](#the-drag-data-item-type-string){#the-datatransferitem-interface:the-drag-data-item-type-string-2}
of the item represented by the
[`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-14}](#datatransferitem)
object.

The
[`getAsString(``callback`{.variable}`)`]{#dom-datatransferitem-getasstring
.dfn dfn-for="DataTransferItem" dfn-type="method"} method must run the
following steps:

1.  If the `callback`{.variable} is null, return.

2.  If the
    [`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-15}](#datatransferitem)
    object is not in the *[read/write mode](#concept-dnd-rw)* or the
    *[read-only mode](#concept-dnd-ro)*, return. The callback is never
    invoked.

3.  If [the drag data item
    kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind-5}
    is not *text*, then return. The callback is never invoked.

4.  Otherwise, [queue a
    task](webappapis.html#queue-a-task){#the-datatransferitem-interface:queue-a-task}
    to invoke `callback`{.variable}, passing the actual data of the item
    represented by the
    [`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-16}](#datatransferitem)
    object as the argument.

The [`getAsFile()`]{#dom-datatransferitem-getasfile .dfn
dfn-for="DataTransferItem" dfn-type="method"} method must run the
following steps:

1.  If the
    [`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-17}](#datatransferitem)
    object is not in the *[read/write mode](#concept-dnd-rw)* or the
    *[read-only mode](#concept-dnd-ro)*, then return null.

2.  If [the drag data item
    kind](#the-drag-data-item-kind){#the-datatransferitem-interface:the-drag-data-item-kind-6}
    is not *File*, then return null.

3.  Return a new
    [`File`{#the-datatransferitem-interface:file-3}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
    object representing the actual data of the item represented by the
    [`DataTransferItem`{#the-datatransferitem-interface:datatransferitem-18}](#datatransferitem)
    object.

#### [6.11.4]{.secno} The [`DragEvent`{#the-dragevent-interface:dragevent}](#dragevent) interface[](#the-dragevent-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DragEvent/DragEvent](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent "This constructor is used to create a synthetic DragEvent object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari14+]{.safari .yes}[Chrome46+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOSNo]{.safari_ios
.no}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[DragEvent](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent "The DragEvent interface is a DOM event that represents a drag and drop interaction. The user initiates a drag by placing a pointer device (such as a mouse) on the touch surface and then dragging the pointer to a new location (such as another DOM element). Applications are free to interpret a drag and drop interaction in an application-specific way.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari14+]{.safari .yes}[Chrome46+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOSNo]{.safari_ios
.no}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

The drag-and-drop processing model involves several events. They all use
the [`DragEvent`{#the-dragevent-interface:dragevent-2}](#dragevent)
interface.

``` idl
[Exposed=Window]
interface DragEvent : MouseEvent {
  constructor(DOMString type, optional DragEventInit eventInitDict = {});

  readonly attribute DataTransfer? dataTransfer;
};

dictionary DragEventInit : MouseEventInit {
  DataTransfer? dataTransfer = null;
};
```

`event`{.variable}`.`[`dataTransfer`](#dom-dragevent-datatransfer){#dom-dragevent-datatransfer-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DragEvent/dataTransfer](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/dataTransfer "The DragEvent.dataTransfer property holds the drag operation's data (as a DataTransfer object).")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari14+]{.safari .yes}[Chrome46+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer[🔰
10+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOSNo]{.safari_ios
.no}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the
[`DataTransfer`{#the-dragevent-interface:datatransfer-3}](#datatransfer)
object for the event.

Although, for consistency with other event interfaces, the
[`DragEvent`{#the-dragevent-interface:dragevent-3}](#dragevent)
interface has a constructor, it is not particularly useful. In
particular, there\'s no way to create a useful
[`DataTransfer`{#the-dragevent-interface:datatransfer-4}](#datatransfer)
object from script, as
[`DataTransfer`{#the-dragevent-interface:datatransfer-5}](#datatransfer)
objects have a processing and security model that is coordinated by the
browser during drag-and-drops.

The [`dataTransfer`]{#dom-dragevent-datatransfer .dfn
dfn-for="DragEvent" dfn-type="attribute"} attribute of the
[`DragEvent`{#the-dragevent-interface:dragevent-4}](#dragevent)
interface must return the value it was initialized to. It represents the
context information for the event.

When a user agent is required to [fire a DND event]{#fire-a-dnd-event
.dfn} named `e`{.variable} at an element, using a particular [drag data
store](#drag-data-store){#the-dragevent-interface:drag-data-store}, and
optionally with a specific `related target`{.variable}, the user agent
must run the following steps:

1.  Let `dataDragStoreWasChanged`{.variable} be false.

2.  If no specific `related target`{.variable} was provided, set
    `related target`{.variable} to null.

3.  Let `window`{.variable} be the [relevant global
    object](webappapis.html#concept-relevant-global){#the-dragevent-interface:concept-relevant-global}
    of the
    [`Document`{#the-dragevent-interface:document}](dom.html#document)
    object of the specified target element.

4.  If `e`{.variable} is
    [`dragstart`{#the-dragevent-interface:event-dnd-dragstart}](#event-dnd-dragstart),
    then set the [drag data store
    mode](#drag-data-store-mode){#the-dragevent-interface:drag-data-store-mode}
    to the [read/write
    mode](#concept-dnd-rw){#the-dragevent-interface:concept-dnd-rw} and
    set `dataDragStoreWasChanged`{.variable} to true.

    If `e`{.variable} is
    [`drop`{#the-dragevent-interface:event-dnd-drop}](#event-dnd-drop),
    set the [drag data store
    mode](#drag-data-store-mode){#the-dragevent-interface:drag-data-store-mode-2}
    to the [read-only
    mode](#concept-dnd-ro){#the-dragevent-interface:concept-dnd-ro}.

5.  Let `dataTransfer`{.variable} be a newly created
    [`DataTransfer`{#the-dragevent-interface:datatransfer-6}](#datatransfer)
    object associated with the given [drag data
    store](#drag-data-store){#the-dragevent-interface:drag-data-store-2}.

6.  Set the
    [`effectAllowed`{#the-dragevent-interface:dom-datatransfer-effectallowed}](#dom-datatransfer-effectallowed)
    attribute to the [drag data
    store](#drag-data-store){#the-dragevent-interface:drag-data-store-3}\'s
    [drag data store allowed effects
    state](#drag-data-store-allowed-effects-state){#the-dragevent-interface:drag-data-store-allowed-effects-state}.

7.  Set the
    [`dropEffect`{#the-dragevent-interface:dom-datatransfer-dropeffect}](#dom-datatransfer-dropeffect)
    attribute to
    \"[`none`{#the-dragevent-interface:dom-datatransfer-dropeffect-none}](#dom-datatransfer-dropeffect-none)\"
    if `e`{.variable} is
    [`dragstart`{#the-dragevent-interface:event-dnd-dragstart-2}](#event-dnd-dragstart),
    [`drag`{#the-dragevent-interface:event-dnd-drag}](#event-dnd-drag),
    or
    [`dragleave`{#the-dragevent-interface:event-dnd-dragleave}](#event-dnd-dragleave);
    to the value corresponding to the [current drag
    operation](#current-drag-operation){#the-dragevent-interface:current-drag-operation}
    if `e`{.variable} is
    [`drop`{#the-dragevent-interface:event-dnd-drop-2}](#event-dnd-drop)
    or
    [`dragend`{#the-dragevent-interface:event-dnd-dragend}](#event-dnd-dragend);
    and to a value based on the
    [`effectAllowed`{#the-dragevent-interface:dom-datatransfer-effectallowed-2}](#dom-datatransfer-effectallowed)
    attribute\'s value and the drag-and-drop source, as given by the
    following table, otherwise (i.e. if `e`{.variable} is
    [`dragenter`{#the-dragevent-interface:event-dnd-dragenter}](#event-dnd-dragenter)
    or
    [`dragover`{#the-dragevent-interface:event-dnd-dragover}](#event-dnd-dragover)):

    [`effectAllowed`{#the-dragevent-interface:dom-datatransfer-effectallowed-3}](#dom-datatransfer-effectallowed)

    [`dropEffect`{#the-dragevent-interface:dom-datatransfer-dropeffect-2}](#dom-datatransfer-dropeffect)

    \"[`none`{#the-dragevent-interface:dom-datatransfer-effectallowed-none}](#dom-datatransfer-effectallowed-none)\"

    \"[`none`{#the-dragevent-interface:dom-datatransfer-dropeffect-none-2}](#dom-datatransfer-dropeffect-none)\"

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-effectallowed-copy}](#dom-datatransfer-effectallowed-copy)\"

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy}](#dom-datatransfer-dropeffect-copy)\"

    \"[`copyLink`{#the-dragevent-interface:dom-datatransfer-effectallowed-copylink}](#dom-datatransfer-effectallowed-copylink)\"

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-2}](#dom-datatransfer-dropeffect-copy)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override},
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link}](#dom-datatransfer-dropeffect-link)\"

    \"[`copyMove`{#the-dragevent-interface:dom-datatransfer-effectallowed-copymove}](#dom-datatransfer-effectallowed-copymove)\"

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-3}](#dom-datatransfer-dropeffect-copy)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-2},
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move}](#dom-datatransfer-dropeffect-move)\"

    \"[`all`{#the-dragevent-interface:dom-datatransfer-effectallowed-all}](#dom-datatransfer-effectallowed-all)\"

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-4}](#dom-datatransfer-dropeffect-copy)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-3},
    either
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-2}](#dom-datatransfer-dropeffect-link)\"
    or
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-2}](#dom-datatransfer-dropeffect-move)\"

    \"[`link`{#the-dragevent-interface:dom-datatransfer-effectallowed-link}](#dom-datatransfer-effectallowed-link)\"

    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-3}](#dom-datatransfer-dropeffect-link)\"

    \"[`linkMove`{#the-dragevent-interface:dom-datatransfer-effectallowed-linkmove}](#dom-datatransfer-effectallowed-linkmove)\"

    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-4}](#dom-datatransfer-dropeffect-link)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-4},
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-3}](#dom-datatransfer-dropeffect-move)\"

    \"[`move`{#the-dragevent-interface:dom-datatransfer-effectallowed-move}](#dom-datatransfer-effectallowed-move)\"

    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-4}](#dom-datatransfer-dropeffect-move)\"

    \"[`uninitialized`{#the-dragevent-interface:dom-datatransfer-effectallowed-uninitialized}](#dom-datatransfer-effectallowed-uninitialized)\"
    when what is being dragged is a selection from a text control

    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-5}](#dom-datatransfer-dropeffect-move)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-5},
    either
    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-5}](#dom-datatransfer-dropeffect-copy)\"
    or
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-5}](#dom-datatransfer-dropeffect-link)\"

    \"[`uninitialized`{#the-dragevent-interface:dom-datatransfer-effectallowed-uninitialized-2}](#dom-datatransfer-effectallowed-uninitialized)\"
    when what is being dragged is a selection

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-6}](#dom-datatransfer-dropeffect-copy)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-6},
    either
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-6}](#dom-datatransfer-dropeffect-link)\"
    or
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-6}](#dom-datatransfer-dropeffect-move)\"

    \"[`uninitialized`{#the-dragevent-interface:dom-datatransfer-effectallowed-uninitialized-3}](#dom-datatransfer-effectallowed-uninitialized)\"
    when what is being dragged is an
    [`a`{#the-dragevent-interface:the-a-element}](text-level-semantics.html#the-a-element)
    element with an
    [`href`{#the-dragevent-interface:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    attribute

    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-7}](#dom-datatransfer-dropeffect-link)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-7},
    either
    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-7}](#dom-datatransfer-dropeffect-copy)\"
    or
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-7}](#dom-datatransfer-dropeffect-move)\"

    Any other case

    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-8}](#dom-datatransfer-dropeffect-copy)\",
    or, [if
    appropriate](#concept-platform-dropeffect-override){#the-dragevent-interface:concept-platform-dropeffect-override-8},
    either
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-8}](#dom-datatransfer-dropeffect-link)\"
    or
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-8}](#dom-datatransfer-dropeffect-move)\"

    Where the table above provides [possibly appropriate
    alternatives]{#concept-platform-dropeffect-override .dfn}, user
    agents may instead use the listed alternative values if platform
    conventions dictate that the user has requested those alternate
    effects.

    For example, Windows platform conventions are such that dragging
    while holding the \"alt\" key indicates a preference for linking the
    data, rather than moving or copying it. Therefore, on a Windows
    system, if
    \"[`link`{#the-dragevent-interface:dom-datatransfer-dropeffect-link-9}](#dom-datatransfer-dropeffect-link)\"
    is an option according to the table above while the \"alt\" key is
    depressed, the user agent could select that instead of
    \"[`copy`{#the-dragevent-interface:dom-datatransfer-dropeffect-copy-9}](#dom-datatransfer-dropeffect-copy)\"
    or
    \"[`move`{#the-dragevent-interface:dom-datatransfer-dropeffect-move-9}](#dom-datatransfer-dropeffect-move)\".

8.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#the-dragevent-interface:creating-an-event
    x-internal="creating-an-event"} using
    [`DragEvent`{#the-dragevent-interface:dragevent-5}](#dragevent).

9.  Initialize `event`{.variable}\'s
    [`type`{#the-dragevent-interface:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    attribute to `e`{.variable}, its
    [`bubbles`{#the-dragevent-interface:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute to true, its
    [`view`{#the-dragevent-interface:dom-uievent-view}](https://w3c.github.io/uievents/#dom-uievent-view){x-internal="dom-uievent-view"}
    attribute to `window`{.variable}, its
    [`relatedTarget`{#the-dragevent-interface:dom-mouseevent-relatedtarget}](https://w3c.github.io/uievents/#dom-mouseevent-relatedtarget){x-internal="dom-mouseevent-relatedtarget"}
    attribute to `related target`{.variable}, and its
    [`dataTransfer`{#the-dragevent-interface:dom-dragevent-datatransfer-2}](#dom-dragevent-datatransfer)
    attribute to `dataTransfer`{.variable}.

10. If `e`{.variable} is not
    [`dragleave`{#the-dragevent-interface:event-dnd-dragleave-2}](#event-dnd-dragleave)
    or
    [`dragend`{#the-dragevent-interface:event-dnd-dragend-2}](#event-dnd-dragend),
    then initialize `event`{.variable}\'s
    [`cancelable`{#the-dragevent-interface:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attribute to true.

11. Initialize `event`{.variable}\'s mouse and key attributes according
    to the state of the input devices as they would be for user
    interaction events.

    If there is no relevant pointing device, then initialize
    `event`{.variable}\'s `screenX`, `screenY`,
    [`clientX`{#the-dragevent-interface:mouseevent-clientx}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clientx){x-internal="mouseevent-clientx"},
    [`clientY`{#the-dragevent-interface:mouseevent-clienty}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clienty){x-internal="mouseevent-clienty"},
    and `button` attributes to 0.

12. [Dispatch](https://dom.spec.whatwg.org/#concept-event-dispatch){#the-dragevent-interface:concept-event-dispatch
    x-internal="concept-event-dispatch"} `event`{.variable} at the
    specified target element.

13. Set the [drag data store allowed effects
    state](#drag-data-store-allowed-effects-state){#the-dragevent-interface:drag-data-store-allowed-effects-state-2}
    to the current value of `dataTransfer`{.variable}\'s
    [`effectAllowed`{#the-dragevent-interface:dom-datatransfer-effectallowed-4}](#dom-datatransfer-effectallowed)
    attribute. (It can only have changed value if `e`{.variable} is
    [`dragstart`{#the-dragevent-interface:event-dnd-dragstart-3}](#event-dnd-dragstart).)

14. If `dataDragStoreWasChanged`{.variable} is true, then set the [drag
    data store
    mode](#drag-data-store-mode){#the-dragevent-interface:drag-data-store-mode-3}
    back to the [protected
    mode](#concept-dnd-p){#the-dragevent-interface:concept-dnd-p}.

15. Break the association between `dataTransfer`{.variable} and the
    [drag data
    store](#drag-data-store){#the-dragevent-interface:drag-data-store-4}.

#### [6.11.5]{.secno} Processing model[](#drag-and-drop-processing-model){.self-link} {#drag-and-drop-processing-model}

When the user attempts to begin a drag operation, the user agent must
run the following steps. User agents must act as if these steps were run
even if the drag actually started in another document or application and
the user agent was not aware that the drag was occurring until it
intersected with a document under the user agent\'s purview.

1.  Determine what is being dragged, as follows:

    If the drag operation was invoked on a selection, then it is the
    selection that is being dragged.

    Otherwise, if the drag operation was invoked on a
    [`Document`{#drag-and-drop-processing-model:document}](dom.html#document),
    it is the first element, going up the ancestor chain, starting at
    the node that the user tried to drag, that has the IDL attribute
    [`draggable`{#drag-and-drop-processing-model:dom-draggable}](#dom-draggable)
    set to true. If there is no such element, then nothing is being
    dragged; return, the drag-and-drop operation is never started.

    Otherwise, the drag operation was invoked outside the user agent\'s
    purview. What is being dragged is defined by the document or
    application where the drag was started.

    [`img`{#drag-and-drop-processing-model:the-img-element}](embedded-content.html#the-img-element)
    elements and
    [`a`{#drag-and-drop-processing-model:the-a-element}](text-level-semantics.html#the-a-element)
    elements with an
    [`href`{#drag-and-drop-processing-model:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    attribute have their
    [`draggable`{#drag-and-drop-processing-model:dom-draggable-2}](#dom-draggable)
    attribute set to true by default.

2.  [Create a drag data
    store](#create-a-drag-data-store){#drag-and-drop-processing-model:create-a-drag-data-store}.
    All the DND events fired subsequently by the steps in this section
    must use this [drag data
    store](#drag-data-store){#drag-and-drop-processing-model:drag-data-store}.

3.  Establish which DOM node is the [source node]{#source-node .dfn}, as
    follows:

    If it is a selection that is being dragged, then the [source
    node](#source-node){#drag-and-drop-processing-model:source-node} is
    the
    [`Text`{#drag-and-drop-processing-model:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node that the user started the drag on (typically the
    [`Text`{#drag-and-drop-processing-model:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node that the user originally clicked). If the user did not specify
    a particular node, for example if the user just told the user agent
    to begin a drag of \"the selection\", then the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-2}
    is the first
    [`Text`{#drag-and-drop-processing-model:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node containing a part of the selection.

    Otherwise, if it is an element that is being dragged, then the
    [source
    node](#source-node){#drag-and-drop-processing-model:source-node-3}
    is the element that is being dragged.

    Otherwise, the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-4}
    is part of another document or application. When this specification
    requires that an event be dispatched at the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-5}
    in this case, the user agent must instead follow the
    platform-specific conventions relevant to that situation.

    Multiple events are fired on the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-6}
    during the course of the drag-and-drop operation.

4.  Determine the [list of dragged nodes]{#list-of-dragged-nodes .dfn},
    as follows:

    If it is a selection that is being dragged, then the [list of
    dragged
    nodes](#list-of-dragged-nodes){#drag-and-drop-processing-model:list-of-dragged-nodes}
    contains, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#drag-and-drop-processing-model:tree-order
    x-internal="tree-order"}, every node that is partially or completely
    included in the selection (including all their ancestors).

    Otherwise, the [list of dragged
    nodes](#list-of-dragged-nodes){#drag-and-drop-processing-model:list-of-dragged-nodes-2}
    contains only the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-7},
    if any.

5.  If it is a selection that is being dragged, then add an item to the
    [drag data store item
    list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list},
    with its properties set as follows:

    [The drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string}
    :   \"[`text/plain`{#drag-and-drop-processing-model:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\"

    [The drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind}
    :   *Text*

    The actual data
    :   The text of the selection

    Otherwise, if any files are being dragged, then add one item per
    file to the [drag data store item
    list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-2},
    with their properties set as follows:

    [The drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-2}
    :   The MIME type of the file, if known, or
        \"[`application/octet-stream`{#drag-and-drop-processing-model:application/octet-stream}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}\"
        otherwise.

    [The drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-2}
    :   *File*

    The actual data
    :   The file\'s contents and name.

    Dragging files can currently only happen from outside a
    [navigable](document-sequences.html#navigable){#drag-and-drop-processing-model:navigable},
    for example from a file system manager application.

    If the drag initiated outside of the application, the user agent
    must add items to the [drag data store item
    list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-3}
    as appropriate for the data being dragged, honoring platform
    conventions where appropriate; however, if the platform conventions
    do not use [MIME
    types](https://mimesniff.spec.whatwg.org/#mime-type){#drag-and-drop-processing-model:mime-type
    x-internal="mime-type"} to label dragged data, the user agent must
    make a best-effort attempt to map the types to MIME types, and, in
    any case, all the [drag data item type
    strings](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-3}
    must be [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#drag-and-drop-processing-model:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

    User agents may also add one or more items representing the
    selection or dragged element(s) in other forms, e.g. as HTML.

6.  If the [list of dragged
    nodes](#list-of-dragged-nodes){#drag-and-drop-processing-model:list-of-dragged-nodes-3}
    is not empty, then [extract the microdata from those nodes into a
    JSON
    form](microdata.html#extracting-json){#drag-and-drop-processing-model:extracting-json},
    and add one item to the [drag data store item
    list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-4},
    with its properties set as follows:

    [The drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-4}
    :   [`application/microdata+json`{#drag-and-drop-processing-model:application/microdata+json}](iana.html#application/microdata+json)

    [The drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-3}
    :   *Text*

    The actual data
    :   The resulting JSON string.

7.  Run the following substeps:

    1.  Let `urls`{.variable} be « ».

    2.  For each `node`{.variable} in the [list of dragged
        nodes](#list-of-dragged-nodes){#drag-and-drop-processing-model:list-of-dragged-nodes-4}:

        If the node is an [`a`{#drag-and-drop-processing-model:the-a-element-2}](text-level-semantics.html#the-a-element) element with an [`href`{#drag-and-drop-processing-model:attr-hyperlink-href-2}](links.html#attr-hyperlink-href) attribute
        :   Add to `urls`{.variable} the result of
            [encoding-parsing-and-serializing a
            URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#drag-and-drop-processing-model:encoding-parsing-and-serializing-a-url}
            given the element\'s
            [`href`{#drag-and-drop-processing-model:attr-hyperlink-href-3}](links.html#attr-hyperlink-href)
            content attribute\'s value, relative to the element\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#drag-and-drop-processing-model:node-document
            x-internal="node-document"}.

        If the node is an [`img`{#drag-and-drop-processing-model:the-img-element-2}](embedded-content.html#the-img-element) element with a [`src`{#drag-and-drop-processing-model:attr-img-src}](embedded-content.html#attr-img-src) attribute
        :   Add to `urls`{.variable} the result of
            [encoding-parsing-and-serializing a
            URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#drag-and-drop-processing-model:encoding-parsing-and-serializing-a-url-2}
            given the element\'s
            [`src`{#drag-and-drop-processing-model:attr-img-src-2}](embedded-content.html#attr-img-src)
            content attribute\'s value, relative to the element\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#drag-and-drop-processing-model:node-document-2
            x-internal="node-document"}.

    3.  If `urls`{.variable} is still empty, then return.

    4.  Let `url string`{.variable} be the result of concatenating the
        strings in `urls`{.variable}, in the order they were added,
        separated by a U+000D CARRIAGE RETURN U+000A LINE FEED character
        pair (CRLF).

    5.  Add one item to the [drag data store item
        list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-5},
        with its properties set as follows:

        [The drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-5}
        :   [`text/uri-list`{#drag-and-drop-processing-model:text/uri-list}](indices.html#text/uri-list)

        [The drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-4}
        :   *Text*

        The actual data
        :   `url string`{.variable}

8.  Update the [drag data store default
    feedback](#drag-data-store-default-feedback){#drag-and-drop-processing-model:drag-data-store-default-feedback}
    as appropriate for the user agent (if the user is dragging the
    selection, then the selection would likely be the basis for this
    feedback; if the user is dragging an element, then that element\'s
    rendering would be used; if the drag began outside the user agent,
    then the platform conventions for determining the drag feedback
    should be used).

9.  [Fire a DND
    event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event}
    named
    [`dragstart`{#drag-and-drop-processing-model:event-dnd-dragstart}](#event-dnd-dragstart)
    at the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-8}.

    If the event is canceled, then the drag-and-drop operation should
    not occur; return.

    Since events with no event listeners registered are, almost by
    definition, never canceled, drag-and-drop is always available to the
    user if the author does not specifically prevent it.

10. [Fire a pointer
    event](https://w3c.github.io/pointerevents/#dfn-fire-a-pointer-event){#drag-and-drop-processing-model:fire-a-pointer-event
    x-internal="fire-a-pointer-event"} at the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-9}
    named
    [`pointercancel`{#drag-and-drop-processing-model:event-pointercancel}](https://w3c.github.io/pointerevents/#the-pointercancel-event){x-internal="event-pointercancel"},
    and fire any other follow-up events as required by Pointer Events.
    [\[POINTEREVENTS\]](references.html#refsPOINTEREVENTS)

11. [Initiate the drag-and-drop
    operation](#initiate-the-drag-and-drop-operation){#drag-and-drop-processing-model:initiate-the-drag-and-drop-operation}
    in a manner consistent with platform conventions, and as described
    below.

    The drag-and-drop feedback must be generated from the first of the
    following sources that is available:

    1.  The [drag data store
        bitmap](#drag-data-store-bitmap){#drag-and-drop-processing-model:drag-data-store-bitmap},
        if any. In this case, the [drag data store hot spot
        coordinate](#drag-data-store-hot-spot-coordinate){#drag-and-drop-processing-model:drag-data-store-hot-spot-coordinate}
        should be used as hints for where to put the cursor relative to
        the resulting image. The values are expressed as distances in
        [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#drag-and-drop-processing-model:'px'
        x-internal="'px'"} from the left side and from the top side of
        the image respectively. [\[CSS\]](references.html#refsCSS)
    2.  The [drag data store default
        feedback](#drag-data-store-default-feedback){#drag-and-drop-processing-model:drag-data-store-default-feedback-2}.

From the moment that the user agent is to [initiate the drag-and-drop
operation]{#initiate-the-drag-and-drop-operation .dfn}, until the end of
the drag-and-drop operation, device input events (e.g. mouse and
keyboard events) must be suppressed.

During the drag operation, the element directly indicated by the user as
the drop target is called the [immediate user
selection]{#immediate-user-selection .dfn}. (Only elements can be
selected by the user; other nodes must not be made available as drop
targets.) However, the [immediate user
selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection}
is not necessarily the [current target element]{#current-target-element
.dfn}, which is the element currently selected for the drop part of the
drag-and-drop operation.

The [immediate user
selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-2}
changes as the user selects different elements (either by pointing at
them with a pointing device, or by selecting them in some other way).
The [current target
element](#current-target-element){#drag-and-drop-processing-model:current-target-element}
changes when the [immediate user
selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-3}
changes, based on the results of event listeners in the document, as
described below.

Both the [current target
element](#current-target-element){#drag-and-drop-processing-model:current-target-element-2}
and the [immediate user
selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-4}
can be null, which means no target element is selected. They can also
both be elements in other (DOM-based) documents, or other (non-web)
programs altogether. (For example, a user could drag text to a
word-processor.) The [current target
element](#current-target-element){#drag-and-drop-processing-model:current-target-element-3}
is initially null.

In addition, there is also a [current drag
operation]{#current-drag-operation .dfn}, which can take on the values
\"[`none`]{#concept-current-drag-operation-none .dfn}\",
\"[`copy`]{#concept-current-drag-operation-copy .dfn}\",
\"[`link`]{#concept-current-drag-operation-link .dfn}\", and
\"[`move`]{#concept-current-drag-operation-move .dfn}\". Initially, it
has the value
\"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none}](#concept-current-drag-operation-none)\".
It is updated by the user agent as described in the steps below.

User agents must, as soon as the drag operation is
[initiated](#initiate-the-drag-and-drop-operation){#drag-and-drop-processing-model:initiate-the-drag-and-drop-operation-2}
and every 350ms (±200ms) thereafter for as long as the drag operation is
ongoing, [queue a
task](webappapis.html#queue-a-task){#drag-and-drop-processing-model:queue-a-task}
to perform the following steps in sequence:

1.  If the user agent is still performing the previous iteration of the
    sequence (if any) when the next iteration becomes due, return for
    this iteration (effectively \"skipping missed frames\" of the
    drag-and-drop operation).

2.  [Fire a DND
    event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-2}
    named
    [`drag`{#drag-and-drop-processing-model:event-dnd-drag}](#event-dnd-drag)
    at the [source
    node](#source-node){#drag-and-drop-processing-model:source-node-10}.
    If this event is canceled, the user agent must set the [current drag
    operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation}
    to
    \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-2}](#concept-current-drag-operation-none)\"
    (no drag operation).

3.  If the
    [`drag`{#drag-and-drop-processing-model:event-dnd-drag-2}](#event-dnd-drag)
    event was not canceled and the user has not ended the drag-and-drop
    operation, check the state of the drag-and-drop operation, as
    follows:

    1.  If the user is indicating a different [immediate user
        selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-5}
        than during the last iteration (or if this is the first
        iteration), and if this [immediate user
        selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-6}
        is not the same as the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-4},
        then update the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-5}
        as follows:

        If the new [immediate user selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-7} is null
        :   Set the [current target
            element](#current-target-element){#drag-and-drop-processing-model:current-target-element-6}
            to null also.

        If the new [immediate user selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-8} is in a non-DOM document or application
        :   Set the [current target
            element](#current-target-element){#drag-and-drop-processing-model:current-target-element-7}
            to the [immediate user
            selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-9}.

        Otherwise

        :   [Fire a DND
            event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-3}
            named
            [`dragenter`{#drag-and-drop-processing-model:event-dnd-dragenter}](#event-dnd-dragenter)
            at the [immediate user
            selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-10}.

            If the event is canceled, then set the [current target
            element](#current-target-element){#drag-and-drop-processing-model:current-target-element-8}
            to the [immediate user
            selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-11}.

            Otherwise, run the appropriate step from the following list:

            If the [immediate user selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-12} is a text control (e.g., [`textarea`{#drag-and-drop-processing-model:the-textarea-element}](form-elements.html#the-textarea-element), or an [`input`{#drag-and-drop-processing-model:the-input-element}](input.html#the-input-element) element whose [`type`{#drag-and-drop-processing-model:attr-input-type}](input.html#attr-input-type) attribute is in the [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#drag-and-drop-processing-model:text-(type=text)-state-and-search-state-(type=search)} state) or an [editing host](interaction.html#editing-host){#drag-and-drop-processing-model:editing-host} or [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#drag-and-drop-processing-model:editable x-internal="editable"} element, and the [drag data store item list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-6} has an item with [the drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-6} \"[`text/plain`{#drag-and-drop-processing-model:text/plain-2}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\" and [the drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-5} *text*
            :   Set the [current target
                element](#current-target-element){#drag-and-drop-processing-model:current-target-element-9}
                to the [immediate user
                selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-13}
                anyway.

            If the [immediate user selection](#immediate-user-selection){#drag-and-drop-processing-model:immediate-user-selection-14} is [the body element](dom.html#the-body-element-2){#drag-and-drop-processing-model:the-body-element-2}
            :   Leave the [current target
                element](#current-target-element){#drag-and-drop-processing-model:current-target-element-10}
                unchanged.

            Otherwise

            :   [Fire a DND
                event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-4}
                named
                [`dragenter`{#drag-and-drop-processing-model:event-dnd-dragenter-2}](#event-dnd-dragenter)
                at [the body
                element](dom.html#the-body-element-2){#drag-and-drop-processing-model:the-body-element-2-2},
                if there is one, or at the
                [`Document`{#drag-and-drop-processing-model:document-2}](dom.html#document)
                object, if not. Then, set the [current target
                element](#current-target-element){#drag-and-drop-processing-model:current-target-element-11}
                to [the body
                element](dom.html#the-body-element-2){#drag-and-drop-processing-model:the-body-element-2-3},
                regardless of whether that event was canceled or not.

    2.  If the previous step caused the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-12}
        to change, and if the previous target element was not null or a
        part of a non-DOM document, then [fire a DND
        event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-5}
        named
        [`dragleave`{#drag-and-drop-processing-model:event-dnd-dragleave}](#event-dnd-dragleave)
        at the previous target element, with the new [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-13}
        as the specific `related target`{.variable}.

    3.  If the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-14}
        is a DOM element, then [fire a DND
        event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-6}
        named
        [`dragover`{#drag-and-drop-processing-model:event-dnd-dragover}](#event-dnd-dragover)
        at this [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-15}.

        If the
        [`dragover`{#drag-and-drop-processing-model:event-dnd-dragover-2}](#event-dnd-dragover)
        event is not canceled, run the appropriate step from the
        following list:

        If the [current target element](#current-target-element){#drag-and-drop-processing-model:current-target-element-16} is a text control (e.g., [`textarea`{#drag-and-drop-processing-model:the-textarea-element-2}](form-elements.html#the-textarea-element), or an [`input`{#drag-and-drop-processing-model:the-input-element-2}](input.html#the-input-element) element whose [`type`{#drag-and-drop-processing-model:attr-input-type-2}](input.html#attr-input-type) attribute is in the [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#drag-and-drop-processing-model:text-(type=text)-state-and-search-state-(type=search)-2} state) or an [editing host](interaction.html#editing-host){#drag-and-drop-processing-model:editing-host-2} or [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#drag-and-drop-processing-model:editable-2 x-internal="editable"} element, and the [drag data store item list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-7} has an item with [the drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-7} \"[`text/plain`{#drag-and-drop-processing-model:text/plain-3}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\" and [the drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-6} *text*
        :   Set the [current drag
            operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-2}
            to either
            \"[`copy`{#drag-and-drop-processing-model:concept-current-drag-operation-copy}](#concept-current-drag-operation-copy)\"
            or
            \"[`move`{#drag-and-drop-processing-model:concept-current-drag-operation-move}](#concept-current-drag-operation-move)\",
            as appropriate given the platform conventions.

        Otherwise

        :   Reset the [current drag
            operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-3}
            to
            \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-3}](#concept-current-drag-operation-none)\".

        Otherwise (if the
        [`dragover`{#drag-and-drop-processing-model:event-dnd-dragover-3}](#event-dnd-dragover)
        event *is* canceled), set the [current drag
        operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-4}
        based on the values of the
        [`effectAllowed`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed}](#dom-datatransfer-effectallowed)
        and
        [`dropEffect`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect}](#dom-datatransfer-dropeffect)
        attributes of the
        [`DragEvent`{#drag-and-drop-processing-model:dragevent}](#dragevent)
        object\'s
        [`dataTransfer`{#drag-and-drop-processing-model:dom-dragevent-datatransfer}](#dom-dragevent-datatransfer)
        object as they stood after the event
        [dispatch](https://dom.spec.whatwg.org/#concept-event-dispatch){#drag-and-drop-processing-model:concept-event-dispatch
        x-internal="concept-event-dispatch"} finished, as per the
        following table:

        [`effectAllowed`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-2}](#dom-datatransfer-effectallowed)

        [`dropEffect`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect-2}](#dom-datatransfer-dropeffect)

        Drag operation

        \"[`uninitialized`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-uninitialized}](#dom-datatransfer-effectallowed-uninitialized)\",
        \"[`copy`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-copy}](#dom-datatransfer-effectallowed-copy)\",
        \"[`copyLink`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-copylink}](#dom-datatransfer-effectallowed-copylink)\",
        \"[`copyMove`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-copymove}](#dom-datatransfer-effectallowed-copymove)\",
        or
        \"[`all`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-all}](#dom-datatransfer-effectallowed-all)\"

        \"[`copy`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect-copy}](#dom-datatransfer-dropeffect-copy)\"

        \"[`copy`{#drag-and-drop-processing-model:concept-current-drag-operation-copy-2}](#concept-current-drag-operation-copy)\"

        \"[`uninitialized`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-uninitialized-2}](#dom-datatransfer-effectallowed-uninitialized)\",
        \"[`link`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-link}](#dom-datatransfer-effectallowed-link)\",
        \"[`copyLink`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-copylink-2}](#dom-datatransfer-effectallowed-copylink)\",
        \"[`linkMove`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-linkmove}](#dom-datatransfer-effectallowed-linkmove)\",
        or
        \"[`all`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-all-2}](#dom-datatransfer-effectallowed-all)\"

        \"[`link`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect-link}](#dom-datatransfer-dropeffect-link)\"

        \"[`link`{#drag-and-drop-processing-model:concept-current-drag-operation-link}](#concept-current-drag-operation-link)\"

        \"[`uninitialized`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-uninitialized-3}](#dom-datatransfer-effectallowed-uninitialized)\",
        \"[`move`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-move}](#dom-datatransfer-effectallowed-move)\",
        \"[`copyMove`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-copymove-2}](#dom-datatransfer-effectallowed-copymove)\",
        \"[`linkMove`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-linkmove-2}](#dom-datatransfer-effectallowed-linkmove)\",
        or
        \"[`all`{#drag-and-drop-processing-model:dom-datatransfer-effectallowed-all-3}](#dom-datatransfer-effectallowed-all)\"

        \"[`move`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect-move}](#dom-datatransfer-dropeffect-move)\"

        \"[`move`{#drag-and-drop-processing-model:concept-current-drag-operation-move-2}](#concept-current-drag-operation-move)\"

        Any other case

        \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-4}](#concept-current-drag-operation-none)\"

    4.  Otherwise, if the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-17}
        is not a DOM element, use platform-specific mechanisms to
        determine what drag operation is being performed (none, copy,
        link, or move), and set the [current drag
        operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-5}
        accordingly.

    5.  Update the drag feedback (e.g. the mouse cursor) to match the
        [current drag
        operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-6},
        as follows:

        Drag operation

        Feedback

        \"[`copy`{#drag-and-drop-processing-model:concept-current-drag-operation-copy-3}](#concept-current-drag-operation-copy)\"

        Data will be copied if dropped here.

        \"[`link`{#drag-and-drop-processing-model:concept-current-drag-operation-link-2}](#concept-current-drag-operation-link)\"

        Data will be linked if dropped here.

        \"[`move`{#drag-and-drop-processing-model:concept-current-drag-operation-move-3}](#concept-current-drag-operation-move)\"

        Data will be moved if dropped here.

        \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-5}](#concept-current-drag-operation-none)\"

        No operation allowed, dropping here will cancel the
        drag-and-drop operation.

4.  Otherwise, if the user ended the drag-and-drop operation (e.g. by
    releasing the mouse button in a mouse-driven drag-and-drop
    interface), or if the
    [`drag`{#drag-and-drop-processing-model:event-dnd-drag-3}](#event-dnd-drag)
    event was canceled, then this will be the last iteration. Run the
    following steps, then stop the drag-and-drop operation:

    1.  If the [current drag
        operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-7}
        is
        \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-6}](#concept-current-drag-operation-none)\"
        (no drag operation), or, if the user ended the drag-and-drop
        operation by canceling it (e.g. by hitting the
        [[Escape]{.kbd}]{.kbd} key), or if the [current target
        element](#current-target-element){#drag-and-drop-processing-model:current-target-element-18}
        is null, then the drag operation failed. Run these substeps:

        1.  Let `dropped`{.variable} be false.

        2.  If the [current target
            element](#current-target-element){#drag-and-drop-processing-model:current-target-element-19}
            is a DOM element, [fire a DND
            event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-7}
            named
            [`dragleave`{#drag-and-drop-processing-model:event-dnd-dragleave-2}](#event-dnd-dragleave)
            at it; otherwise, if it is not null, use platform-specific
            conventions for drag cancelation.

        3.  Set the [current drag
            operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-8}
            to
            \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-7}](#concept-current-drag-operation-none)\".

        Otherwise, the drag operation might be a success; run these
        substeps:

        1.  Let `dropped`{.variable} be true.

        2.  If the [current target
            element](#current-target-element){#drag-and-drop-processing-model:current-target-element-20}
            is a DOM element, [fire a DND
            event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-8}
            named
            [`drop`{#drag-and-drop-processing-model:event-dnd-drop}](#event-dnd-drop)
            at it; otherwise, use platform-specific conventions for
            indicating a drop.

        3.  If the event is canceled, set the [current drag
            operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-9}
            to the value of the
            [`dropEffect`{#drag-and-drop-processing-model:dom-datatransfer-dropeffect-3}](#dom-datatransfer-dropeffect)
            attribute of the
            [`DragEvent`{#drag-and-drop-processing-model:dragevent-2}](#dragevent)
            object\'s
            [`dataTransfer`{#drag-and-drop-processing-model:dom-dragevent-datatransfer-2}](#dom-dragevent-datatransfer)
            object as it stood after the event
            [dispatch](https://dom.spec.whatwg.org/#concept-event-dispatch){#drag-and-drop-processing-model:concept-event-dispatch-2
            x-internal="concept-event-dispatch"} finished.

            Otherwise, the event is not canceled; perform the event\'s
            default action, which depends on the exact target as
            follows:

            If the [current target element](#current-target-element){#drag-and-drop-processing-model:current-target-element-21} is a text control (e.g., [`textarea`{#drag-and-drop-processing-model:the-textarea-element-3}](form-elements.html#the-textarea-element), or an [`input`{#drag-and-drop-processing-model:the-input-element-3}](input.html#the-input-element) element whose [`type`{#drag-and-drop-processing-model:attr-input-type-3}](input.html#attr-input-type) attribute is in the [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#drag-and-drop-processing-model:text-(type=text)-state-and-search-state-(type=search)-3} state) or an [editing host](interaction.html#editing-host){#drag-and-drop-processing-model:editing-host-3} or [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#drag-and-drop-processing-model:editable-3 x-internal="editable"} element, and the [drag data store item list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-8} has an item with [the drag data item type string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-8} \"[`text/plain`{#drag-and-drop-processing-model:text/plain-4}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\" and [the drag data item kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-7} *text*
            :   Insert the actual data of the first item in the [drag
                data store item
                list](#drag-data-store-item-list){#drag-and-drop-processing-model:drag-data-store-item-list-9}
                to have [a drag data item type
                string](#the-drag-data-item-type-string){#drag-and-drop-processing-model:the-drag-data-item-type-string-9}
                of
                \"[`text/plain`{#drag-and-drop-processing-model:text/plain-5}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\"
                and [a drag data item
                kind](#the-drag-data-item-kind){#drag-and-drop-processing-model:the-drag-data-item-kind-8}
                that is *text* into the text control or [editing
                host](interaction.html#editing-host){#drag-and-drop-processing-model:editing-host-4}
                or
                [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#drag-and-drop-processing-model:editable-4
                x-internal="editable"} element in a manner consistent
                with platform-specific conventions (e.g. inserting it at
                the current mouse cursor position, or inserting it at
                the end of the field).

            Otherwise

            :   Reset the [current drag
                operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-10}
                to
                \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-8}](#concept-current-drag-operation-none)\".

    2.  [Fire a DND
        event](#fire-a-dnd-event){#drag-and-drop-processing-model:fire-a-dnd-event-9}
        named
        [`dragend`{#drag-and-drop-processing-model:event-dnd-dragend}](#event-dnd-dragend)
        at the [source
        node](#source-node){#drag-and-drop-processing-model:source-node-11}.

    3.  Run the appropriate steps from the following list as the default
        action of the
        [`dragend`{#drag-and-drop-processing-model:event-dnd-dragend-2}](#event-dnd-dragend)
        event:

        If `dropped`{.variable} is true, the [current target element](#current-target-element){#drag-and-drop-processing-model:current-target-element-22} is a *text control* (see below), the [current drag operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-11} is \"[`move`{#drag-and-drop-processing-model:concept-current-drag-operation-move-4}](#concept-current-drag-operation-move)\", and the source of the drag-and-drop operation is a selection in the DOM that is entirely contained within an [editing host](interaction.html#editing-host){#drag-and-drop-processing-model:editing-host-5}
        :   [Delete the
            selection](https://w3c.github.io/editing/docs/execCommand/#delete-the-selection){#drag-and-drop-processing-model:delete-the-selection
            x-internal="delete-the-selection"}.

        If `dropped`{.variable} is true, the [current target element](#current-target-element){#drag-and-drop-processing-model:current-target-element-23} is a *text control* (see below), the [current drag operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-12} is \"[`move`{#drag-and-drop-processing-model:concept-current-drag-operation-move-5}](#concept-current-drag-operation-move)\", and the source of the drag-and-drop operation is a selection in a text control
        :   The user agent should delete the dragged selection from the
            relevant text control.

        If `dropped`{.variable} is false or if the [current drag operation](#current-drag-operation){#drag-and-drop-processing-model:current-drag-operation-13} is \"[`none`{#drag-and-drop-processing-model:concept-current-drag-operation-none-9}](#concept-current-drag-operation-none)\"
        :   The drag was canceled. If the platform conventions dictate
            that this be represented to the user (e.g. by animating the
            dragged selection going back to the source of the
            drag-and-drop operation), then do so.

        Otherwise

        :   The event has no default action.

        For the purposes of this step, a *text control* is a
        [`textarea`{#drag-and-drop-processing-model:the-textarea-element-4}](form-elements.html#the-textarea-element)
        element or an
        [`input`{#drag-and-drop-processing-model:the-input-element-4}](input.html#the-input-element)
        element whose
        [`type`{#drag-and-drop-processing-model:attr-input-type-4}](input.html#attr-input-type)
        attribute is in one of the
        [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#drag-and-drop-processing-model:text-(type=text)-state-and-search-state-(type=search)-4},
        [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#drag-and-drop-processing-model:text-(type=text)-state-and-search-state-(type=search)-5},
        [Tel](input.html#telephone-state-(type=tel)){#drag-and-drop-processing-model:telephone-state-(type=tel)},
        [URL](input.html#url-state-(type=url)){#drag-and-drop-processing-model:url-state-(type=url)},
        [Email](input.html#email-state-(type=email)){#drag-and-drop-processing-model:email-state-(type=email)},
        [Password](input.html#password-state-(type=password)){#drag-and-drop-processing-model:password-state-(type=password)},
        or
        [Number](input.html#number-state-(type=number)){#drag-and-drop-processing-model:number-state-(type=number)}
        states.

User agents are encouraged to consider how to react to drags near the
edge of scrollable regions. For example, if a user drags a link to the
bottom of the
[viewport](https://drafts.csswg.org/css2/#viewport){#drag-and-drop-processing-model:viewport
x-internal="viewport"} on a long page, it might make sense to scroll the
page so that the user can drop the link lower on the page.

This model is independent of which
[`Document`{#drag-and-drop-processing-model:document-3}](dom.html#document)
object the nodes involved are from; the events are fired as described
above and the rest of the processing model runs as described above,
irrespective of how many documents are involved in the operation.

#### [6.11.6]{.secno} Events summary[](#dndevents){.self-link} {#dndevents}

*This section is non-normative.*

The following events are involved in the drag-and-drop model.

Event name

Target

Cancelable?

[Drag data store
mode](#drag-data-store-mode){#dndevents:drag-data-store-mode}

[`dropEffect`{#dndevents:dom-datatransfer-dropeffect}](#dom-datatransfer-dropeffect)

Default Action

[`dragstart`]{#event-dnd-dragstart .dfn
dfn-for="GlobalEventHandlers,Text" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dragstart_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dragstart_event "The dragstart event is fired when the user starts dragging an element or text selection.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Source node](#source-node){#dndevents:source-node}

✓ Cancelable

[Read/write mode](#concept-dnd-rw){#dndevents:concept-dnd-rw}

\"[`none`{#dndevents:dom-datatransfer-dropeffect-none}](#dom-datatransfer-dropeffect-none)\"

Initiate the drag-and-drop operation

[`drag`]{#event-dnd-drag .dfn dfn-for="GlobalEventHandlers,Text"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/drag_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/drag_event "The drag event is fired every few hundred milliseconds as an element or text selection is being dragged by the user.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Source node](#source-node){#dndevents:source-node-2}

✓ Cancelable

[Protected mode](#concept-dnd-p){#dndevents:concept-dnd-p}

\"[`none`{#dndevents:dom-datatransfer-dropeffect-none-2}](#dom-datatransfer-dropeffect-none)\"

Continue the drag-and-drop operation

[`dragenter`]{#event-dnd-dragenter .dfn dfn-for="GlobalEventHandlers"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dragenter_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dragenter_event "The dragenter event is fired when a dragged element or text selection enters a valid drop target.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Immediate user
selection](#immediate-user-selection){#dndevents:immediate-user-selection}
or [the body
element](dom.html#the-body-element-2){#dndevents:the-body-element-2}

✓ Cancelable

[Protected mode](#concept-dnd-p){#dndevents:concept-dnd-p-2}

[Based on `effectAllowed` value](#dropEffect-initialisation)

Reject [immediate user
selection](#immediate-user-selection){#dndevents:immediate-user-selection-2}
as potential [target
element](#current-target-element){#dndevents:current-target-element}

[`dragleave`]{#event-dnd-dragleave .dfn dfn-for="GlobalEventHandlers"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dragleave_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dragleave_event "The dragleave event is fired when a dragged element or text selection leaves a valid drop target.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Previous target
element](#current-target-element){#dndevents:current-target-element-2}

---

[Protected mode](#concept-dnd-p){#dndevents:concept-dnd-p-3}

\"[`none`{#dndevents:dom-datatransfer-dropeffect-none-3}](#dom-datatransfer-dropeffect-none)\"

None

[`dragover`]{#event-dnd-dragover .dfn dfn-for="GlobalEventHandlers"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dragover_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dragover_event "The dragover event is fired when an element or text selection is being dragged over a valid drop target (every few hundred milliseconds).")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Current target
element](#current-target-element){#dndevents:current-target-element-3}

✓ Cancelable

[Protected mode](#concept-dnd-p){#dndevents:concept-dnd-p-4}

[Based on `effectAllowed` value](#dropEffect-initialisation)

Reset the [current drag
operation](#current-drag-operation){#dndevents:current-drag-operation}
to \"none\"

[`drop`]{#event-dnd-drop .dfn dfn-for="GlobalEventHandlers"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/drop_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/drop_event "The drop event is fired when an element or text selection is dropped on a valid drop target. To ensure that the drop event always fires as expected, you should always include a preventDefault() call in the part of your code which handles the dragover event.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Current target
element](#current-target-element){#dndevents:current-target-element-4}

✓ Cancelable

[Read-only mode](#concept-dnd-ro){#dndevents:concept-dnd-ro}

[Current drag
operation](#current-drag-operation){#dndevents:current-drag-operation-2}

Varies

[`dragend`]{#event-dnd-dragend .dfn dfn-for="GlobalEventHandlers,Text"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dragend_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dragend_event "The dragend event is fired when a drag operation is being ended (by releasing a mouse button or hitting the escape key).")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Source node](#source-node){#dndevents:source-node-3}

---

[Protected mode](#concept-dnd-p){#dndevents:concept-dnd-p-5}

[Current drag
operation](#current-drag-operation){#dndevents:current-drag-operation-3}

Varies

All of these events bubble, are composed, and the
[`effectAllowed`{#dndevents:dom-datatransfer-effectallowed}](#dom-datatransfer-effectallowed)
attribute always has the value it had after the
[`dragstart`{#dndevents:event-dnd-dragstart}](#event-dnd-dragstart)
event, defaulting to
\"[`uninitialized`{#dndevents:dom-datatransfer-effectallowed-uninitialized}](#dom-datatransfer-effectallowed-uninitialized)\"
in the
[`dragstart`{#dndevents:event-dnd-dragstart-2}](#event-dnd-dragstart)
event.

#### [6.11.7]{.secno} The [`draggable`{#the-draggable-attribute:attr-draggable}](#attr-draggable) attribute[](#the-draggable-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/draggable](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/draggable "The draggable global attribute is an enumerated attribute that indicates whether the element can be dragged, either with native browser behavior or the HTML Drag and Drop API.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

All [HTML
elements](infrastructure.html#html-elements){#the-draggable-attribute:html-elements}
may have the [`draggable`]{#attr-draggable .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute set. The
[`draggable`{#the-draggable-attribute:attr-draggable-2}](#attr-draggable)
attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-draggable-attribute:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`true`]{#attr-draggable-true .dfn dfn-for="html-global/draggable"
dfn-type="attr-value"}

[True]{#attr-draggable-true-state .dfn}

The element will be draggable.

[`false`]{#attr-draggable-false .dfn dfn-for="html-global/draggable"
dfn-type="attr-value"}

[False]{#attr-draggable-false-state .dfn}

The element will not be draggable.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Auto]{#attr-draggable-auto-state .dfn} state. The auto state
uses the default behavior of the user agent.

An element with a
[`draggable`{#the-draggable-attribute:attr-draggable-3}](#attr-draggable)
attribute should also have a
[`title`{#the-draggable-attribute:attr-title}](dom.html#attr-title)
attribute that names the element for the purpose of non-visual
interactions.

`element`{.variable}`.`[`draggable`](#dom-draggable){#dom-draggable-dev}` [ = ``value`{.variable}` ]`

:   Returns true if the element is draggable; otherwise, returns false.

    Can be set, to override the default and set the
    [`draggable`{#the-draggable-attribute:attr-draggable-4}](#attr-draggable)
    content attribute.

The [`draggable`]{#dom-draggable .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute, whose value depends on the content
attribute\'s in the way described below, controls whether or not the
element is draggable. Generally, only text selections are draggable, but
elements whose
[`draggable`{#the-draggable-attribute:dom-draggable}](#dom-draggable)
IDL attribute is true become draggable as well.

If an element\'s
[`draggable`{#the-draggable-attribute:attr-draggable-5}](#attr-draggable)
content attribute has the state
[True](#attr-draggable-true-state){#the-draggable-attribute:attr-draggable-true-state},
the
[`draggable`{#the-draggable-attribute:dom-draggable-2}](#dom-draggable)
IDL attribute must return true.

Otherwise, if the element\'s
[`draggable`{#the-draggable-attribute:attr-draggable-6}](#attr-draggable)
content attribute has the state
[False](#attr-draggable-false-state){#the-draggable-attribute:attr-draggable-false-state},
the
[`draggable`{#the-draggable-attribute:dom-draggable-3}](#dom-draggable)
IDL attribute must return false.

Otherwise, the element\'s
[`draggable`{#the-draggable-attribute:attr-draggable-7}](#attr-draggable)
content attribute has the state
[Auto](#attr-draggable-auto-state){#the-draggable-attribute:attr-draggable-auto-state}.
If the element is an
[`img`{#the-draggable-attribute:the-img-element}](embedded-content.html#the-img-element)
element, an
[`object`{#the-draggable-attribute:the-object-element}](iframe-embed-object.html#the-object-element)
element that
[represents](dom.html#represents){#the-draggable-attribute:represents}
an image, or an
[`a`{#the-draggable-attribute:the-a-element}](text-level-semantics.html#the-a-element)
element with an
[`href`{#the-draggable-attribute:attr-hyperlink-href}](links.html#attr-hyperlink-href)
content attribute, the
[`draggable`{#the-draggable-attribute:dom-draggable-4}](#dom-draggable)
IDL attribute must return true; otherwise, the
[`draggable`{#the-draggable-attribute:dom-draggable-5}](#dom-draggable)
IDL attribute must return false.

If the
[`draggable`{#the-draggable-attribute:dom-draggable-6}](#dom-draggable)
IDL attribute is set to the value false, the
[`draggable`{#the-draggable-attribute:attr-draggable-8}](#attr-draggable)
content attribute must be set to the literal value \"`false`\". If the
[`draggable`{#the-draggable-attribute:dom-draggable-7}](#dom-draggable)
IDL attribute is set to the value true, the
[`draggable`{#the-draggable-attribute:attr-draggable-9}](#attr-draggable)
content attribute must be set to the literal value \"`true`\".

#### [6.11.8]{.secno} Security risks in the drag-and-drop model[](#security-risks-in-the-drag-and-drop-model){.self-link}

User agents must not make the data added to the
[`DataTransfer`{#security-risks-in-the-drag-and-drop-model:datatransfer}](#datatransfer)
object during the
[`dragstart`{#security-risks-in-the-drag-and-drop-model:event-dnd-dragstart}](#event-dnd-dragstart)
event available to scripts until the
[`drop`{#security-risks-in-the-drag-and-drop-model:event-dnd-drop}](#event-dnd-drop)
event, because otherwise, if a user were to drag sensitive information
from one document to a second document, crossing a hostile third
document in the process, the hostile document could intercept the data.

For the same reason, user agents must consider a drop to be successful
only if the user specifically ended the drag operation --- if any
scripts end the drag operation, it must be considered unsuccessful
(canceled) and the
[`drop`{#security-risks-in-the-drag-and-drop-model:event-dnd-drop-2}](#event-dnd-drop)
event must not be fired.

User agents should take care to not start drag-and-drop operations in
response to script actions. For example, in a mouse-and-window
environment, if a script moves a window while the user has their mouse
button depressed, the UA would not consider that to start a drag. This
is important because otherwise UAs could cause data to be dragged from
sensitive sources and dropped into hostile documents without the user\'s
consent.

User agents should filter potentially active (scripted) content (e.g.
HTML) when it is dragged and when it is dropped, using a safelist of
known-safe features. Similarly, [relative
URLs](https://url.spec.whatwg.org/#syntax-url-relative){#security-risks-in-the-drag-and-drop-model:relative-url
x-internal="relative-url"} should be turned into absolute URLs to avoid
references changing in unexpected ways. This specification does not
specify how this is performed.

::: example
Consider a hostile page providing some content and getting the user to
select and drag and drop (or indeed, copy and paste) that content to a
victim page\'s
[`contenteditable`{#security-risks-in-the-drag-and-drop-model:attr-contenteditable}](interaction.html#attr-contenteditable)
region. If the browser does not ensure that only safe content is
dragged, potentially unsafe content such as scripts and event handlers
in the selection, once dropped (or pasted) into the victim site, get the
privileges of the victim site. This would thus enable a cross-site
scripting attack.
:::

[← 6 User interaction](interaction.html) --- [Table of Contents](./) ---
[6.12 The popover attribute →](popover.html)
