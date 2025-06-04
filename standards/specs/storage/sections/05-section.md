## [5. ]{.secno}[Persistence permission]{.content}[](#persistence){.self-link} {#persistence .heading .settled level="5"}

A [local storage bucket](#bucket){#ref-for-bucket② link-type="dfn"} can
only have its [mode](#bucket-mode){#ref-for-bucket-mode link-type="dfn"}
change to \"`persistent`\" if the user (or user agent on behalf of the
user) has granted permission to use the \"`persistent-storage`\"
[powerful
feature](https://w3c.github.io/permissions/#dfn-powerful-feature){#ref-for-dfn-powerful-feature
link-type="dfn"}.

When granted to an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin②
link-type="dfn"}, the persistence permission can be used to protect
storage from the user agent's clearing policies. The user agent cannot
clear storage marked as persistent without involvement from the
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin③
link-type="dfn"} or user. This makes it particularly useful for
resources the user needs to have available while offline or resources
the user creates locally.

The \"`persistent-storage`\" [powerful
feature](https://w3c.github.io/permissions/#dfn-powerful-feature){#ref-for-dfn-powerful-feature①
link-type="dfn"}'s permission-related algorithms, and types are
defaulted, except for:

[permission state](https://w3c.github.io/permissions/#dfn-permission-state){#ref-for-dfn-permission-state link-type="dfn"}

:   \"`persistent-storage`\"\'s [permission
    state](https://w3c.github.io/permissions/#dfn-permission-state){#ref-for-dfn-permission-state①
    link-type="dfn"} must have the same value for all [environment
    settings
    objects](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑦
    link-type="dfn"} with a given
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin①
    link-type="dfn"}.

[permission revocation algorithm](https://w3c.github.io/permissions/#dfn-permission-revocation-algorithm){#ref-for-dfn-permission-revocation-algorithm link-type="dfn"}

:   1.  If the result of [getting the current permission
        state](https://w3c.github.io/permissions/#dfn-getting-the-current-permission-state){#ref-for-dfn-getting-the-current-permission-state
        link-type="dfn"} with \"`persistent-storage`\" is
        \"[`granted`{.idl}](https://w3c.github.io/permissions/#dom-permissionstate-granted){#ref-for-dom-permissionstate-granted
        link-type="idl"}\", then return.

    2.  Let `shelf`{.variable} be the result of running [obtain a local
        storage
        shelf](#obtain-a-local-storage-shelf){#ref-for-obtain-a-local-storage-shelf
        link-type="dfn"} with [current settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#current-settings-object){#ref-for-current-settings-object
        link-type="dfn"}.

    3.  Set `shelf`{.variable}'s [bucket
        map](#bucket-map){#ref-for-bucket-map⑤
        link-type="dfn"}\[\"`default`\"\]\'s
        [mode](#bucket-mode){#ref-for-bucket-mode① link-type="dfn"} to
        \"`best-effort`\".

