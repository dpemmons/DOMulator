## [12]{.secno} Web storage[](#webstorage){.self-link} {#webstorage}

:::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Web_Storage_API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API "The Web Storage API provides mechanisms by which browsers can store key/value pairs, in a much more intuitive fashion than using cookies.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::

::: feature
[Web_Storage_API/Using_the_Web_Storage_API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API/Using_the_Web_Storage_API "The Web Storage API provides mechanisms by which browsers can securely store key/value pairs.")
:::
::::::

### [12.1]{.secno} Introduction[](#introduction-15){.self-link} {#introduction-15}

*This section is non-normative.*

This specification introduces two related mechanisms, similar to HTTP
session cookies, for storing name-value pairs on the client side.
[\[COOKIES\]](references.html#refsCOOKIES)

The first is designed for scenarios where the user is carrying out a
single transaction, but could be carrying out multiple transactions in
different windows at the same time.

Cookies don\'t really handle this case well. For example, a user could
be buying plane tickets in two different windows, using the same site.
If the site used cookies to keep track of which ticket the user was
buying, then as the user clicked from page to page in both windows, the
ticket currently being purchased would \"leak\" from one window to the
other, potentially causing the user to buy two tickets for the same
flight without noticing.

To address this, this specification introduces the
[`sessionStorage`{#introduction-15:dom-sessionstorage}](#dom-sessionstorage)
getter. Sites can add data to the session storage, and it will be
accessible to any page from the same site opened in that window.

::: example
For example, a page could have a checkbox that the user ticks to
indicate that they want insurance:

``` html
<label>
 <input type="checkbox" onchange="sessionStorage.insurance = checked ? 'true' : ''">
  I want insurance on this trip.
</label>
```

A later page could then check, from script, whether the user had checked
the checkbox or not:

``` js
if (sessionStorage.insurance) { ... }
```

If the user had multiple windows opened on the site, each one would have
its own individual copy of the session storage object.
:::

The second storage mechanism is designed for storage that spans multiple
windows, and lasts beyond the current session. In particular, web
applications might wish to store megabytes of user data, such as entire
user-authored documents or a user\'s mailbox, on the client side for
performance reasons.

Again, cookies do not handle this case well, because they are
transmitted with every request.

The
[`localStorage`{#introduction-15:dom-localstorage}](#dom-localstorage)
getter is used to access a page\'s local storage area.

::: example
The site at example.com can display a count of how many times the user
has loaded its page by putting the following at the bottom of its page:

``` html
<p>
  You have viewed this page
  <span id="count">an untold number of</span>
  time(s).
</p>
<script>
  if (!localStorage.pageLoadCount)
    localStorage.pageLoadCount = 0;
  localStorage.pageLoadCount = parseInt(localStorage.pageLoadCount) + 1;
  document.getElementById('count').textContent = localStorage.pageLoadCount;
</script>
```
:::

Each site has its own separate storage area.

The
[`localStorage`{#introduction-15:dom-localstorage-2}](#dom-localstorage)
getter provides access to shared state. This specification does not
define the interaction with other agent clusters in a multiprocess user
agent, and authors are encouraged to assume that there is no locking
mechanism. A site could, for instance, try to read the value of a key,
increment its value, then write it back out, using the new value as a
unique identifier for the session; if the site does this twice in two
different browser windows at the same time, it might end up using the
same \"unique\" identifier for both sessions, with potentially
disastrous effects.

### [12.2]{.secno} The API[](#storage){.self-link} {#storage}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Storage](https://developer.mozilla.org/en-US/docs/Web/API/Storage "The Storage interface of the Web Storage API provides access to a particular domain's session or local storage. It allows, for example, the addition, modification, or deletion of stored data items.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

#### [12.2.1]{.secno} The [`Storage`{#the-storage-interface:storage-2}](#storage-2) interface[](#the-storage-interface){.self-link}

``` idl
[Exposed=Window]
interface Storage {
  readonly attribute unsigned long length;
  DOMString? key(unsigned long index);
  getter DOMString? getItem(DOMString key);
  setter undefined setItem(DOMString key, DOMString value);
  deleter undefined removeItem(DOMString key);
  undefined clear();
};
```

`storage`{.variable}`.`[`length`](#dom-storage-length){#dom-storage-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/length](https://developer.mozilla.org/en-US/docs/Web/API/Storage/length "The length read-only property of the Storage interface returns the number of data items stored in a given Storage object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns the number of key/value pairs.

`storage`{.variable}`.`[`key`](#dom-storage-key){#dom-storage-key-dev}` (``n`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/key](https://developer.mozilla.org/en-US/docs/Web/API/Storage/key "The key() method of the Storage interface, when passed a number n, returns the name of the nth key in a given Storage object. The order of keys is user-agent defined, so you should not rely on it.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns the name of the `n`{.variable}th key, or null if `n`{.variable}
is greater than or equal to the number of key/value pairs.

`value`{.variable}` = ``storage`{.variable}`.`[`getItem`](#dom-storage-getitem){#dom-storage-getitem-dev}` (``key`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/getItem](https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem "The getItem() method of the Storage interface, when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

`value`{.variable}` = ``storage`{.variable}`[``key`{.variable}`]`

Returns the current value associated with the given `key`{.variable}, or
null if the given `key`{.variable} does not exist.

`storage`{.variable}`.`[`setItem`](#dom-storage-setitem){#dom-storage-setitem-dev}` (``key`{.variable}`, ``value`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/setItem](https://developer.mozilla.org/en-US/docs/Web/API/Storage/setItem "The setItem() method of the Storage interface, when passed a key name and value, will add that key to the given Storage object, or update that key's value if it already exists.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

`storage`{.variable}`[``key`{.variable}`] = ``value`{.variable}

Sets the value of the pair identified by `key`{.variable} to
`value`{.variable}, creating a new key/value pair if none existed for
`key`{.variable} previously.

Throws a
[\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror){#the-storage-interface:quotaexceedederror
x-internal="quotaexceedederror"}
[`DOMException`{#the-storage-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the new value couldn\'t be set. (Setting could fail if, e.g., the
user has disabled storage for the site, or if the quota has been
exceeded.)

Dispatches a
[`storage`{#the-storage-interface:event-storage}](indices.html#event-storage)
event on
[`Window`{#the-storage-interface:window}](nav-history-apis.html#window)
objects holding an equivalent
[`Storage`{#the-storage-interface:storage-2-2}](#storage-2) object.

`storage`{.variable}`.`[`removeItem`](#dom-storage-removeitem){#dom-storage-removeitem-dev}` (``key`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/removeItem](https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem "The removeItem() method of the Storage interface, when passed a key name, will remove that key from the given Storage object if it exists. The Storage interface of the Web Storage API provides access to a particular domain's session or local storage.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`delete`{#the-storage-interface:delete}](https://tc39.es/ecma262/#sec-delete-operator){x-internal="delete"}
`storage`{.variable}\[`key`{.variable}\]

Removes the key/value pair with the given `key`{.variable}, if a
key/value pair with the given `key`{.variable} exists.

Dispatches a
[`storage`{#the-storage-interface:event-storage-2}](indices.html#event-storage)
event on
[`Window`{#the-storage-interface:window-2}](nav-history-apis.html#window)
objects holding an equivalent
[`Storage`{#the-storage-interface:storage-2-3}](#storage-2) object.

`storage`{.variable}`.`[`clear`](#dom-storage-clear){#dom-storage-clear-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Storage/clear](https://developer.mozilla.org/en-US/docs/Web/API/Storage/clear "The clear() method of the Storage interface clears all keys stored in a given Storage object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Removes all key/value pairs, if there are any.

Dispatches a
[`storage`{#the-storage-interface:event-storage-3}](indices.html#event-storage)
event on
[`Window`{#the-storage-interface:window-3}](nav-history-apis.html#window)
objects holding an equivalent
[`Storage`{#the-storage-interface:storage-2-4}](#storage-2) object.

A [`Storage`{#the-storage-interface:storage-2-5}](#storage-2) object has
an associated:

[map]{#concept-storage-map .dfn}
:   A [storage proxy
    map](https://storage.spec.whatwg.org/#storage-proxy-map){#the-storage-interface:storage-proxy-map
    x-internal="storage-proxy-map"}.

[type]{#concept-storage-type .dfn}
:   \"`local`\" or \"`session`\".

To [reorder]{#concept-storage-reorder .dfn} a
[`Storage`{#the-storage-interface:storage-2-6}](#storage-2) object
`storage`{.variable}, reorder `storage`{.variable}\'s
[map](#concept-storage-map){#the-storage-interface:concept-storage-map}\'s
[entries](https://infra.spec.whatwg.org/#map-entry){#the-storage-interface:map-entry
x-internal="map-entry"} in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-storage-interface:implementation-defined
x-internal="implementation-defined"} manner.

Unfortunate as it is, iteration order is not defined and can change upon
most mutations.

To [broadcast]{#concept-storage-broadcast .dfn} a
[`Storage`{#the-storage-interface:storage-2-7}](#storage-2) object
`storage`{.variable}, given a `key`{.variable}, `oldValue`{.variable},
and `newValue`{.variable}, run these steps:

1.  Let `thisDocument`{.variable} be `storage`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#the-storage-interface:concept-relevant-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-storage-interface:concept-document-window}.

2.  Let `url`{.variable} be the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-storage-interface:concept-url-serializer
    x-internal="concept-url-serializer"} of `thisDocument`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-storage-interface:the-document's-address
    x-internal="the-document's-address"}.

3.  Let `remoteStorages`{.variable} be all
    [`Storage`{#the-storage-interface:storage-2-8}](#storage-2) objects
    excluding `storage`{.variable} whose:

    - [type](#concept-storage-type){#the-storage-interface:concept-storage-type}
      is `storage`{.variable}\'s
      [type](#concept-storage-type){#the-storage-interface:concept-storage-type-2}
    - [relevant settings
      object](webappapis.html#relevant-settings-object){#the-storage-interface:relevant-settings-object}\'s
      [origin](browsers.html#concept-origin){#the-storage-interface:concept-origin}
      is [same
      origin](browsers.html#same-origin){#the-storage-interface:same-origin}
      with `storage`{.variable}\'s [relevant settings
      object](webappapis.html#relevant-settings-object){#the-storage-interface:relevant-settings-object-2}\'s
      [origin](browsers.html#concept-origin){#the-storage-interface:concept-origin-2}

    and, if
    [type](#concept-storage-type){#the-storage-interface:concept-storage-type-3}
    is \"`session`\", whose [relevant settings
    object](webappapis.html#relevant-settings-object){#the-storage-interface:relevant-settings-object-3}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-storage-interface:concept-document-window-2}\'s
    [node
    navigable](document-sequences.html#node-navigable){#the-storage-interface:node-navigable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#the-storage-interface:nav-traversable}
    is `thisDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-storage-interface:node-navigable-2}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#the-storage-interface:nav-traversable-2}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-storage-interface:list-iterate
    x-internal="list-iterate"} `remoteStorage`{.variable} of
    `remoteStorages`{.variable}: [queue a global
    task](webappapis.html#queue-a-global-task){#the-storage-interface:queue-a-global-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-storage-interface:dom-manipulation-task-source}
    given `remoteStorage`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-storage-interface:concept-relevant-global-2}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-storage-interface:concept-event-fire
    x-internal="concept-event-fire"} named
    [`storage`{#the-storage-interface:event-storage-4}](indices.html#event-storage)
    at `remoteStorage`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-storage-interface:concept-relevant-global-3},
    using
    [`StorageEvent`{#the-storage-interface:storageevent}](#storageevent),
    with
    [`key`{#the-storage-interface:dom-storageevent-key}](#dom-storageevent-key)
    initialized to `key`{.variable},
    [`oldValue`{#the-storage-interface:dom-storageevent-oldvalue}](#dom-storageevent-oldvalue)
    initialized to `oldValue`{.variable},
    [`newValue`{#the-storage-interface:dom-storageevent-newvalue}](#dom-storageevent-newvalue)
    initialized to `newValue`{.variable},
    [`url`{#the-storage-interface:dom-storageevent-url}](#dom-storageevent-url)
    initialized to `url`{.variable}, and
    [`storageArea`{#the-storage-interface:dom-storageevent-storagearea}](#dom-storageevent-storagearea)
    initialized to `remoteStorage`{.variable}.

    The [`Document`{#the-storage-interface:document}](dom.html#document)
    object associated with the resulting
    [task](webappapis.html#concept-task){#the-storage-interface:concept-task}
    is not necessarily [fully
    active](document-sequences.html#fully-active){#the-storage-interface:fully-active},
    but events fired on such objects are ignored by the [event
    loop](webappapis.html#event-loop){#the-storage-interface:event-loop}
    until the
    [`Document`{#the-storage-interface:document-2}](dom.html#document)
    becomes [fully
    active](document-sequences.html#fully-active){#the-storage-interface:fully-active-2}
    again.

------------------------------------------------------------------------

The [`length`]{#dom-storage-length .dfn dfn-for="Storage"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this
x-internal="this"}\'s
[map](#concept-storage-map){#the-storage-interface:concept-storage-map-2}\'s
[size](https://infra.spec.whatwg.org/#map-size){#the-storage-interface:map-size
x-internal="map-size"}.

The [`key(``index`{.variable}`)`]{#dom-storage-key .dfn
dfn-for="Storage" dfn-type="method"} method steps are:

1.  If `index`{.variable} is greater than or equal to
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-2
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-3}\'s
    [size](https://infra.spec.whatwg.org/#map-size){#the-storage-interface:map-size-2
    x-internal="map-size"}, then return null.

2.  Let `keys`{.variable} be the result of running [get the
    keys](https://infra.spec.whatwg.org/#map-getting-the-keys){#the-storage-interface:map-get-the-keys
    x-internal="map-get-the-keys"} on
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-3
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-4}.

3.  Return `keys`{.variable}\[`index`{.variable}\].

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#the-storage-interface:supported-property-names
x-internal="supported-property-names"} on a
[`Storage`{#the-storage-interface:storage-2-9}](#storage-2) object
`storage`{.variable} are the result of running [get the
keys](https://infra.spec.whatwg.org/#map-getting-the-keys){#the-storage-interface:map-get-the-keys-2
x-internal="map-get-the-keys"} on `storage`{.variable}\'s
[map](#concept-storage-map){#the-storage-interface:concept-storage-map-5}.

The [`getItem(``key`{.variable}`)`]{#dom-storage-getitem .dfn
dfn-for="Storage" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-4
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-6}\[`key`{.variable}\]
    does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#the-storage-interface:map-exists
    x-internal="map-exists"}, then return null.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-5
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-7}\[`key`{.variable}\].

The
[`setItem(``key`{.variable}`, ``value`{.variable}`)`]{#dom-storage-setitem
.dfn dfn-for="Storage" dfn-type="method"} method steps are:

1.  Let `oldValue`{.variable} be null.

2.  Let `reorder`{.variable} be true.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-6
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-8}\[`key`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-storage-interface:map-exists-2
    x-internal="map-exists"}:

    1.  Set `oldValue`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-7
        x-internal="this"}\'s
        [map](#concept-storage-map){#the-storage-interface:concept-storage-map-9}\[`key`{.variable}\].

    2.  If `oldValue`{.variable}
        [is](https://infra.spec.whatwg.org/#string-is){#the-storage-interface:is
        x-internal="is"} `value`{.variable}, then return.

    3.  Set `reorder`{.variable} to false.

4.  If `value`{.variable} cannot be stored, then throw a
    [\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror){#the-storage-interface:quotaexceedederror-2
    x-internal="quotaexceedederror"}
    [`DOMException`{#the-storage-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  [Set](https://infra.spec.whatwg.org/#map-set){#the-storage-interface:map-set
    x-internal="map-set"}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-8
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-10}\[`key`{.variable}\]
    to `value`{.variable}.

6.  If `reorder`{.variable} is true, then
    [reorder](#concept-storage-reorder){#the-storage-interface:concept-storage-reorder}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-9
    x-internal="this"}.

7.  [Broadcast](#concept-storage-broadcast){#the-storage-interface:concept-storage-broadcast}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-10
    x-internal="this"} with `key`{.variable}, `oldValue`{.variable}, and
    `value`{.variable}.

The [`removeItem(``key`{.variable}`)`]{#dom-storage-removeitem .dfn
dfn-for="Storage" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-11
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-11}\[`key`{.variable}\]
    does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#the-storage-interface:map-exists-3
    x-internal="map-exists"}, then return.

2.  Set `oldValue`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-12
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-12}\[`key`{.variable}\].

3.  [Remove](https://infra.spec.whatwg.org/#map-remove){#the-storage-interface:map-remove
    x-internal="map-remove"}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-13
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-13}\[`key`{.variable}\].

4.  [Reorder](#concept-storage-reorder){#the-storage-interface:concept-storage-reorder-2}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-14
    x-internal="this"}.

5.  [Broadcast](#concept-storage-broadcast){#the-storage-interface:concept-storage-broadcast-2}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-15
    x-internal="this"} with `key`{.variable}, `oldValue`{.variable}, and
    null.

The [`clear()`]{#dom-storage-clear .dfn dfn-for="Storage"
dfn-type="method"} method steps are:

1.  [Clear](https://infra.spec.whatwg.org/#map-clear){#the-storage-interface:map-clear
    x-internal="map-clear"}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-16
    x-internal="this"}\'s
    [map](#concept-storage-map){#the-storage-interface:concept-storage-map-14}.

2.  [Broadcast](#concept-storage-broadcast){#the-storage-interface:concept-storage-broadcast-3}
    [this](https://webidl.spec.whatwg.org/#this){#the-storage-interface:this-17
    x-internal="this"} with null, null, and null.

#### [12.2.2]{.secno} The [`sessionStorage`{#the-sessionstorage-attribute:dom-sessionstorage}](#dom-sessionstorage) getter[](#the-sessionstorage-attribute){.self-link} {#the-sessionstorage-attribute}

``` idl
interface mixin WindowSessionStorage {
  readonly attribute Storage sessionStorage;
};
Window includes WindowSessionStorage;
```

`window`{.variable}`.`[`sessionStorage`](#dom-sessionstorage){#dom-sessionstorage-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/sessionStorage](https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage "The read-only sessionStorage property accesses a session Storage object for the current origin. sessionStorage is similar to localStorage; the difference is that while data in localStorage doesn't expire, data in sessionStorage is cleared when the page session ends.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`Storage`{#the-sessionstorage-attribute:storage-2-2}](#storage-2)
object associated with that `window`{.variable}\'s origin\'s session
storage area.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-sessionstorage-attribute:securityerror
x-internal="securityerror"}
[`DOMException`{#the-sessionstorage-attribute:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the
[`Document`{#the-sessionstorage-attribute:document}](dom.html#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-sessionstorage-attribute:concept-document-origin
x-internal="concept-document-origin"} is an [opaque
origin](browsers.html#concept-origin-opaque){#the-sessionstorage-attribute:concept-origin-opaque}
or if the request violates a policy decision (e.g., if the user agent is
configured to not allow the page to persist data).

A
[`Document`{#the-sessionstorage-attribute:document-2}](dom.html#document)
object has an associated [session storage
holder]{#session-storage-holder .dfn}, which is null or a
[`Storage`{#the-sessionstorage-attribute:storage-2-3}](#storage-2)
object. It is initially null.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#the-sessionstorage-attribute:tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`sessionStorage`]{#dom-sessionstorage .dfn
dfn-for="WindowSessionStorage" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-sessionstorage-attribute:this
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-sessionstorage-attribute:concept-document-window}\'s
    [session storage
    holder](#session-storage-holder){#the-sessionstorage-attribute:session-storage-holder}
    is non-null, then return
    [this](https://webidl.spec.whatwg.org/#this){#the-sessionstorage-attribute:this-2
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-sessionstorage-attribute:concept-document-window-2}\'s
    [session storage
    holder](#session-storage-holder){#the-sessionstorage-attribute:session-storage-holder-2}.

2.  Let `map`{.variable} be the result of running [obtain a session
    storage bottle
    map](https://storage.spec.whatwg.org/#obtain-a-session-storage-bottle-map){#the-sessionstorage-attribute:obtain-a-session-storage-bottle-map
    x-internal="obtain-a-session-storage-bottle-map"} with
    [this](https://webidl.spec.whatwg.org/#this){#the-sessionstorage-attribute:this-3
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#the-sessionstorage-attribute:relevant-settings-object}
    and \"`sessionStorage`\".

3.  If `map`{.variable} is failure, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-sessionstorage-attribute:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#the-sessionstorage-attribute:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `storage`{.variable} be a new
    [`Storage`{#the-sessionstorage-attribute:storage-2-4}](#storage-2)
    object whose
    [map](#concept-storage-map){#the-sessionstorage-attribute:concept-storage-map}
    is `map`{.variable}.

5.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-sessionstorage-attribute:this-4
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-sessionstorage-attribute:concept-document-window-3}\'s
    [session storage
    holder](#session-storage-holder){#the-sessionstorage-attribute:session-storage-holder-3}
    to `storage`{.variable}.

6.  Return `storage`{.variable}.

After [creating a new auxiliary browsing context and
document](document-sequences.html#creating-a-new-auxiliary-browsing-context){#the-sessionstorage-attribute:creating-a-new-auxiliary-browsing-context},
the session storage [is
copied](document-sequences.html#copy-session-storage) over.

#### [12.2.3]{.secno} The [`localStorage`{#the-localstorage-attribute:dom-localstorage}](#dom-localstorage) getter[](#the-localstorage-attribute){.self-link} {#the-localstorage-attribute}

``` idl
interface mixin WindowLocalStorage {
  readonly attribute Storage localStorage;
};
Window includes WindowLocalStorage;
```

`window`{.variable}`.`[`localStorage`](#dom-localstorage){#dom-localstorage-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/localStorage](https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage "The localStorage read-only property of the window interface allows you to access a Storage object for the Document's origin; the stored data is saved across browser sessions.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`Storage`{#the-localstorage-attribute:storage-2-2}](#storage-2) object
associated with `window`{.variable}\'s origin\'s local storage area.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-localstorage-attribute:securityerror
x-internal="securityerror"}
[`DOMException`{#the-localstorage-attribute:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the
[`Document`{#the-localstorage-attribute:document}](dom.html#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-localstorage-attribute:concept-document-origin
x-internal="concept-document-origin"} is an [opaque
origin](browsers.html#concept-origin-opaque){#the-localstorage-attribute:concept-origin-opaque}
or if the request violates a policy decision (e.g., if the user agent is
configured to not allow the page to persist data).

A
[`Document`{#the-localstorage-attribute:document-2}](dom.html#document)
object has an associated [local storage holder]{#local-storage-holder
.dfn}, which is null or a
[`Storage`{#the-localstorage-attribute:storage-2-3}](#storage-2) object.
It is initially null.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#the-localstorage-attribute:tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`localStorage`]{#dom-localstorage .dfn dfn-for="WindowLocalStorage"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-localstorage-attribute:this
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-localstorage-attribute:concept-document-window}\'s
    [local storage
    holder](#local-storage-holder){#the-localstorage-attribute:local-storage-holder}
    is non-null, then return
    [this](https://webidl.spec.whatwg.org/#this){#the-localstorage-attribute:this-2
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-localstorage-attribute:concept-document-window-2}\'s
    [local storage
    holder](#local-storage-holder){#the-localstorage-attribute:local-storage-holder-2}.

2.  Let `map`{.variable} be the result of running [obtain a local
    storage bottle
    map](https://storage.spec.whatwg.org/#obtain-a-local-storage-bottle-map){#the-localstorage-attribute:obtain-a-local-storage-bottle-map
    x-internal="obtain-a-local-storage-bottle-map"} with
    [this](https://webidl.spec.whatwg.org/#this){#the-localstorage-attribute:this-3
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#the-localstorage-attribute:relevant-settings-object}
    and \"`localStorage`\".

3.  If `map`{.variable} is failure, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-localstorage-attribute:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#the-localstorage-attribute:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `storage`{.variable} be a new
    [`Storage`{#the-localstorage-attribute:storage-2-4}](#storage-2)
    object whose
    [map](#concept-storage-map){#the-localstorage-attribute:concept-storage-map}
    is `map`{.variable}.

5.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-localstorage-attribute:this-4
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-localstorage-attribute:concept-document-window-3}\'s
    [local storage
    holder](#local-storage-holder){#the-localstorage-attribute:local-storage-holder-3}
    to `storage`{.variable}.

6.  Return `storage`{.variable}.

#### [12.2.4]{.secno} The [`StorageEvent`{#the-storageevent-interface:storageevent}](#storageevent) interface[](#the-storageevent-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[StorageEvent](https://developer.mozilla.org/en-US/docs/Web/API/StorageEvent "The StorageEvent interface is implemented by the storage event, which is sent to a window when a storage area the window has access to is changed within the context of another document.")

Support in all current engines.

::: support
[Firefox13+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Window]
interface StorageEvent : Event {
  constructor(DOMString type, optional StorageEventInit eventInitDict = {});

  readonly attribute DOMString? key;
  readonly attribute DOMString? oldValue;
  readonly attribute DOMString? newValue;
  readonly attribute USVString url;
  readonly attribute Storage? storageArea;

  undefined initStorageEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false, optional DOMString? key = null, optional DOMString? oldValue = null, optional DOMString? newValue = null, optional USVString url = "", optional Storage? storageArea = null);
};

dictionary StorageEventInit : EventInit {
  DOMString? key = null;
  DOMString? oldValue = null;
  DOMString? newValue = null;
  USVString url = "";
  Storage? storageArea = null;
};
```

`event`{.variable}`.`[`key`](#dom-storageevent-key){#dom-storageevent-key-dev}
:   Returns the key of the storage item being changed.

`event`{.variable}`.`[`oldValue`](#dom-storageevent-oldvalue){#dom-storageevent-oldvalue-dev}
:   Returns the old value of the key of the storage item whose value is
    being changed.

`event`{.variable}`.`[`newValue`](#dom-storageevent-newvalue){#dom-storageevent-newvalue-dev}
:   Returns the new value of the key of the storage item whose value is
    being changed.

`event`{.variable}`.`[`url`](#dom-storageevent-url){#dom-storageevent-url-dev}
:   Returns the
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-storageevent-interface:the-document's-address
    x-internal="the-document's-address"} of the document whose storage
    item changed.

`event`{.variable}`.`[`storageArea`](#dom-storageevent-storagearea){#dom-storageevent-storagearea-dev}

:   Returns the
    [`Storage`{#the-storageevent-interface:storage-2-3}](#storage-2)
    object that was affected.

The [`key`]{#dom-storageevent-key .dfn dfn-for="StorageEvent"
dfn-type="attribute"}, [`oldValue`]{#dom-storageevent-oldvalue .dfn
dfn-for="StorageEvent" dfn-type="attribute"},
[`newValue`]{#dom-storageevent-newvalue .dfn dfn-for="StorageEvent"
dfn-type="attribute"}, [`url`]{#dom-storageevent-url .dfn
dfn-for="StorageEvent" dfn-type="attribute"}, and
[`storageArea`]{#dom-storageevent-storagearea .dfn
dfn-for="StorageEvent" dfn-type="attribute"} attributes must return the
values they were initialized to.

The
[`initStorageEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`, ``key`{.variable}`, ``oldValue`{.variable}`, ``newValue`{.variable}`, ``url`{.variable}`, ``storageArea`{.variable}`)`]{#dom-storageevent-initstorageevent
.dfn dfn-for="StorageEvent" dfn-type="method"} method must initialize
the event in a manner analogous to the similarly-named
[`initEvent()`{#the-storageevent-interface:dom-event-initevent}](https://dom.spec.whatwg.org/#dom-event-initevent){x-internal="dom-event-initevent"}
method. [\[DOM\]](references.html#refsDOM)

### [12.3]{.secno} Privacy[](#privacy){.self-link}

#### [12.3.1]{.secno} User tracking[](#user-tracking){.self-link}

A third-party advertiser (or any entity capable of getting content
distributed to multiple sites) could use a unique identifier stored in
its local storage area to track a user across multiple sessions,
building a profile of the user\'s interests to allow for highly targeted
advertising. In conjunction with a site that is aware of the user\'s
real identity (for example an e-commerce site that requires
authenticated credentials), this could allow oppressive groups to target
individuals with greater accuracy than in a world with purely anonymous
web usage.

There are a number of techniques that can be used to mitigate the risk
of user tracking:

Blocking third-party storage

:   User agents may restrict access to the
    [`localStorage`{#user-tracking:dom-localstorage}](#dom-localstorage)
    objects to scripts originating at the domain of the [active
    document](document-sequences.html#nav-document){#user-tracking:nav-document}
    of the [top-level
    traversable](document-sequences.html#top-level-traversable){#user-tracking:top-level-traversable},
    for instance denying access to the API for pages from other domains
    running in
    [`iframe`{#user-tracking:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s.

Expiring stored data

:   User agents may, possibly in a manner configured by the user,
    automatically delete stored data after a period of time.

    For example, a user agent could be configured to treat third-party
    local storage areas as session-only storage, deleting the data once
    the user had closed all the
    [navigables](document-sequences.html#navigable){#user-tracking:navigable}
    that could access it.

    This can restrict the ability of a site to track a user, as the site
    would then only be able to track the user across multiple sessions
    when they authenticate with the site itself (e.g. by making a
    purchase or logging in to a service).

    However, this also reduces the usefulness of the API as a long-term
    storage mechanism. It can also put the user\'s data at risk, if the
    user does not fully understand the implications of data expiration.

Treating persistent storage as cookies

:   If users attempt to protect their privacy by clearing cookies
    without also clearing data stored in the local storage area, sites
    can defeat those attempts by using the two features as redundant
    backup for each other. User agents should present the interfaces for
    clearing these in a way that helps users to understand this
    possibility and enables them to delete data in all persistent
    storage features simultaneously.
    [\[COOKIES\]](references.html#refsCOOKIES)

Site-specific safelisting of access to local storage areas

:   User agents may allow sites to access session storage areas in an
    unrestricted manner, but require the user to authorize access to
    local storage areas.

Origin-tracking of stored data

:   User agents may record the
    [origins](browsers.html#concept-origin){#user-tracking:concept-origin}
    of sites that contained content from third-party origins that caused
    data to be stored.

    If this information is then used to present the view of data
    currently in persistent storage, it would allow the user to make
    informed decisions about which parts of the persistent storage to
    prune. Combined with a blocklist (\"delete this data and prevent
    this domain from ever storing data again\"), the user can restrict
    the use of persistent storage to sites that they trust.

Shared blocklists

:   User agents may allow users to share their persistent storage domain
    blocklists.

    This would allow communities to act together to protect their
    privacy.

While these suggestions prevent trivial use of this API for user
tracking, they do not block it altogether. Within a single domain, a
site can continue to track the user during a session, and can then pass
all this information to the third party along with any identifying
information (names, credit card numbers, addresses) obtained by the
site. If a third party cooperates with multiple sites to obtain such
information, a profile can still be created.

However, user tracking is to some extent possible even with no
cooperation from the user agent whatsoever, for instance by using
session identifiers in URLs, a technique already commonly used for
innocuous purposes but easily repurposed for user tracking (even
retroactively). This information can then be shared with other sites,
using visitors\' IP addresses and other user-specific data (e.g.
user-agent headers and configuration settings) to combine separate
sessions into coherent user profiles.

#### [12.3.2]{.secno} Sensitivity of data[](#sensitivity-of-data){.self-link}

User agents should treat persistently stored data as potentially
sensitive; it\'s quite possible for emails, calendar appointments,
health records, or other confidential documents to be stored in this
mechanism.

To this end, user agents should ensure that when deleting data, it is
promptly deleted from the underlying storage.

### [12.4]{.secno} Security[](#security-storage){.self-link} {#security-storage}

#### [12.4.1]{.secno} DNS spoofing attacks[](#dns-spoofing-attacks){.self-link}

Because of the potential for DNS spoofing attacks, one cannot guarantee
that a host claiming to be in a certain domain really is from that
domain. To mitigate this, pages can use TLS. Pages using TLS can be sure
that only the user, software working on behalf of the user, and other
pages using TLS that have certificates identifying them as being from
the same domain, can access their storage areas.

#### [12.4.2]{.secno} Cross-directory attacks[](#cross-directory-attacks){.self-link}

Different authors sharing one host name, for example users hosting
content on the now defunct `geocities.com`, all share one local storage
object. There is no feature to restrict the access by pathname. Authors
on shared hosts are therefore urged to avoid using these features, as it
would be trivial for other authors to read the data and overwrite it.

Even if a path-restriction feature was made available, the usual DOM
scripting security model would make it trivial to bypass this protection
and access the data from any path.

#### [12.4.3]{.secno} Implementation risks[](#implementation-risks){.self-link}

The two primary risks when implementing these persistent storage
features are letting hostile sites read information from other domains,
and letting hostile sites write information that is then read from other
domains.

Letting third-party sites read data that is not supposed to be read from
their domain causes *information leakage*. For example, a user\'s
shopping wishlist on one domain could be used by another domain for
targeted advertising; or a user\'s work-in-progress confidential
documents stored by a word-processing site could be examined by the site
of a competing company.

Letting third-party sites write data to the persistent storage of other
domains can result in *information spoofing*, which is equally
dangerous. For example, a hostile site could add items to a user\'s
wishlist; or a hostile site could set a user\'s session identifier to a
known ID that the hostile site can then use to track the user\'s actions
on the victim site.

Thus, strictly following the
[origin](browsers.html#concept-origin){#implementation-risks:concept-origin}
model described in this specification is important for user security.

[← 11 Worklets](worklets.html) --- [Table of Contents](./) --- [13 The
HTML syntax →](syntax.html)
