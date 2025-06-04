## [1. ]{.secno}[Introduction]{.content}[](#introduction){.self-link} {#introduction .heading .settled level="1"}

Over the years the web has grown various APIs that can be used for
storage, e.g., IndexedDB, `localStorage`, and `showNotification()`. The
Storage Standard consolidates these APIs by defining:

- A bucket, the primitive these APIs store their data in
- A way of making that bucket persistent
- A way of getting usage and quota estimates for an
  [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin
  link-type="dfn"}

Traditionally, as the user runs out of storage space on their device,
the data stored with these APIs gets lost without the user being able to
intervene. However, persistent buckets cannot be cleared without consent
by the user. This thus brings data guarantees users have enjoyed on
native platforms to the web.

::: {#example-3a7051a8 .example}
[](#example-3a7051a8){.self-link}

A simple way to make storage persistent is through invoking the
[`persist()`{.idl}](#dom-storagemanager-persist){#ref-for-dom-storagemanager-persist
link-type="idl"} method. It simultaneously requests the end user for
permission and changes the storage to be persistent once granted:

``` {.lang-javascript .highlight}
navigator.storage.persist().then(persisted => {
  if (persisted) {
    /* … */
  }
});
```

To not show user-agent-driven dialogs to the end user unannounced
slightly more involved code can be written:

``` {.lang-javascript .highlight}
Promise.all([
  navigator.storage.persisted(),
  navigator.permissions.query({name: "persistent-storage"})
]).then(([persisted, permission]) => {
  if (!persisted && permission.state == "granted") {
    navigator.storage.persist().then( /* … */ );
  } else if (!persisted && permission.state == "prompt") {
    showPersistentStorageExplanation();
  }
});
```

The
[`estimate()`{.idl}](#dom-storagemanager-estimate){#ref-for-dom-storagemanager-estimate
link-type="idl"} method can be used to determine whether there is enough
space left to store content for an application:

``` {.lang-javascript .highlight}
function retrieveNextChunk(nextChunkInfo) {
  return navigator.storage.estimate().then(info => {
    if (info.quota - info.usage > nextChunkInfo.size) {
      return fetch(nextChunkInfo.url);
    } else {
      throw new Error("insufficient space to store next chunk");
    }
  }).then( /* … */ );
}
```
:::

