## [8. ]{.secno}[API]{.content}[](#api){.self-link} {#api .heading .settled level="8"}

``` {.idl .highlight .def}
[SecureContext]
interface mixin NavigatorStorage {
  [SameObject] readonly attribute StorageManager storage;
};
Navigator includes NavigatorStorage;
WorkerNavigator includes NavigatorStorage;
```

Each [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑧
link-type="dfn"} has an associated
[`StorageManager`{.idl}](#storagemanager){#ref-for-storagemanager①
link-type="idl"} object.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

The [`storage`]{#dom-navigatorstorage-storage .dfn .dfn-paneled
.idl-code dfn-for="NavigatorStorage" dfn-type="attribute" export=""}
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this
link-type="dfn"}'s [relevant settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object
link-type="dfn"}'s
[`StorageManager`{.idl}](#storagemanager){#ref-for-storagemanager②
link-type="idl"} object.

``` {.idl .highlight .def}
[SecureContext,
 Exposed=(Window,Worker)]
interface StorageManager {
  Promise<boolean> persisted();
  [Exposed=Window] Promise<boolean> persist();

  Promise<StorageEstimate> estimate();
};

dictionary StorageEstimate {
  unsigned long long usage;
  unsigned long long quota;
};
```

::: {.algorithm algorithm="persisted()" algorithm-for="StorageManager"}
The [`persisted()`]{#dom-storagemanager-persisted .dfn .dfn-paneled
.idl-code dfn-for="StorageManager" dfn-type="method" export=""} method
steps are:

1.  Let `promise`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise
    link-type="dfn"}.

2.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global
    link-type="dfn"}.

3.  Let `shelf`{.variable} be the result of running [obtain a local
    storage
    shelf](#obtain-a-local-storage-shelf){#ref-for-obtain-a-local-storage-shelf①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object①
    link-type="dfn"}.

4.  If `shelf`{.variable} is failure, then
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject
    link-type="dfn"} `promise`{.variable} with a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror
    link-type="idl"}.

5.  Otherwise, run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel
    link-type="dfn"}:

    1.  Let `persisted`{.variable} be true if `shelf`{.variable}'s
        [bucket map](#bucket-map){#ref-for-bucket-map⑦
        link-type="dfn"}\[\"`default`\"\]\'s
        [mode](#bucket-mode){#ref-for-bucket-mode④ link-type="dfn"} is
        \"`persistent`\"; otherwise false.

        It will be false when there's an internal error.

    2.  [Queue a storage
        task](#queue-a-storage-task){#ref-for-queue-a-storage-task
        link-type="dfn"} with `global`{.variable} to
        [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve
        link-type="dfn"} `promise`{.variable} with
        `persisted`{.variable}.

6.  Return `promise`{.variable}.
:::

::: {.algorithm algorithm="persist()" algorithm-for="StorageManager"}
The [`persist()`]{#dom-storagemanager-persist .dfn .dfn-paneled
.idl-code dfn-for="StorageManager" dfn-type="method" export=""} method
steps are:

1.  Let `promise`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise①
    link-type="dfn"}.

2.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global①
    link-type="dfn"}.

3.  Let `shelf`{.variable} be the result of running [obtain a local
    storage
    shelf](#obtain-a-local-storage-shelf){#ref-for-obtain-a-local-storage-shelf②
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object②
    link-type="dfn"}.

4.  If `shelf`{.variable} is failure, then
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject①
    link-type="dfn"} `promise`{.variable} with a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①
    link-type="idl"}.

5.  Otherwise, run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel①
    link-type="dfn"}:

    1.  Let `permission`{.variable} be the result of [requesting
        permission to
        use](https://w3c.github.io/permissions/#dfn-request-permission-to-use){#ref-for-dfn-request-permission-to-use
        link-type="dfn"} \"`persistent-storage`\".

        User agents are encouraged to not let the user answer this
        question twice for the same
        [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑤
        link-type="dfn"} around the same time and this algorithm is not
        equipped to handle such a scenario.

    2.  Let `bucket`{.variable} be `shelf`{.variable}'s [bucket
        map](#bucket-map){#ref-for-bucket-map⑧
        link-type="dfn"}\[\"`default`\"\].

    3.  Let `persisted`{.variable} be true if `bucket`{.variable}'s
        [mode](#bucket-mode){#ref-for-bucket-mode⑤ link-type="dfn"} is
        \"`persistent`\"; otherwise false.

        It will be false when there's an internal error.

    4.  If `persisted`{.variable} is false and `permission`{.variable}
        is
        \"[`granted`{.idl}](https://w3c.github.io/permissions/#dom-permissionstate-granted){#ref-for-dom-permissionstate-granted①
        link-type="idl"}\", then:

        1.  Set `bucket`{.variable}'s
            [mode](#bucket-mode){#ref-for-bucket-mode⑥ link-type="dfn"}
            to \"`persistent`\".

        2.  If there was no internal error, then set
            `persisted`{.variable} to true.

    5.  [Queue a storage
        task](#queue-a-storage-task){#ref-for-queue-a-storage-task①
        link-type="dfn"} with `global`{.variable} to
        [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve①
        link-type="dfn"} `promise`{.variable} with
        `persisted`{.variable}.

6.  Return `promise`{.variable}.
:::

::: {.algorithm algorithm="estimate()" algorithm-for="StorageManager"}
The [`estimate()`]{#dom-storagemanager-estimate .dfn .dfn-paneled
.idl-code dfn-for="StorageManager" dfn-type="method" export=""} method
steps are:

1.  Let `promise`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise②
    link-type="dfn"}.

2.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global②
    link-type="dfn"}.

3.  Let `shelf`{.variable} be the result of running [obtain a local
    storage
    shelf](#obtain-a-local-storage-shelf){#ref-for-obtain-a-local-storage-shelf③
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object③
    link-type="dfn"}.

4.  If `shelf`{.variable} is failure, then
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject②
    link-type="dfn"} `promise`{.variable} with a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②
    link-type="idl"}.

5.  Otherwise, run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel②
    link-type="dfn"}:

    1.  Let `usage`{.variable} be [storage
        usage](#storage-usage){#ref-for-storage-usage link-type="dfn"}
        for `shelf`{.variable}.

    2.  Let `quota`{.variable} be [storage
        quota](#storage-quota){#ref-for-storage-quota① link-type="dfn"}
        for `shelf`{.variable}.

    3.  Let `dictionary`{.variable} be a new
        [`StorageEstimate`{.idl}](#dictdef-storageestimate){#ref-for-dictdef-storageestimate①
        link-type="idl"} dictionary whose
        [`usage`{.idl}](#dom-storageestimate-usage){#ref-for-dom-storageestimate-usage
        link-type="idl"} member is `usage`{.variable} and
        [`quota`{.idl}](#dom-storageestimate-quota){#ref-for-dom-storageestimate-quota
        link-type="idl"} member is `quota`{.variable}.

    4.  If there was an internal error while obtaining
        `usage`{.variable} and `quota`{.variable}, then [queue a storage
        task](#queue-a-storage-task){#ref-for-queue-a-storage-task②
        link-type="dfn"} with `global`{.variable} to
        [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject③
        link-type="dfn"} `promise`{.variable} with a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror③
        link-type="idl"}.

        Internal errors are supposed to be extremely rare and indicate
        some kind of low-level platform or hardware fault. However, at
        the scale of the web with the diversity of implementation and
        platforms, the unexpected does occur.

    5.  Otherwise, [queue a storage
        task](#queue-a-storage-task){#ref-for-queue-a-storage-task③
        link-type="dfn"} with `global`{.variable} to
        [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve②
        link-type="dfn"} `promise`{.variable} with
        `dictionary`{.variable}.

6.  Return `promise`{.variable}.
:::

