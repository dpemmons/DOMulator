HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 7.4 Navigation and session history](browsing-the-web.html) --- [Table
of Contents](./) --- [8 Web application APIs →](webappapis.html)

1.  ::: {#toc-browsers}
    1.  [[7.5]{.secno} Document
        lifecycle](document-lifecycle.html#document-lifecycle)
        1.  [[7.5.1]{.secno} Shared document creation
            infrastructure](document-lifecycle.html#shared-document-creation-infrastructure)
        2.  [[7.5.2]{.secno} Loading HTML
            documents](document-lifecycle.html#read-html)
        3.  [[7.5.3]{.secno} Loading XML
            documents](document-lifecycle.html#read-xml)
        4.  [[7.5.4]{.secno} Loading text
            documents](document-lifecycle.html#read-text)
        5.  [[7.5.5]{.secno} Loading `multipart/x-mixed-replace`
            documents](document-lifecycle.html#read-multipart-x-mixed-replace)
        6.  [[7.5.6]{.secno} Loading media
            documents](document-lifecycle.html#read-media)
        7.  [[7.5.7]{.secno} Loading a document for inline content that
            doesn\'t have a DOM](document-lifecycle.html#read-ua-inline)
        8.  [[7.5.8]{.secno} Finishing the loading
            process](document-lifecycle.html#loading-documents)
        9.  [[7.5.9]{.secno} Unloading
            documents](document-lifecycle.html#unloading-documents)
        10. [[7.5.10]{.secno} Destroying
            documents](document-lifecycle.html#destroying-documents)
        11. [[7.5.11]{.secno} Aborting a document
            load](document-lifecycle.html#aborting-a-document-load)
    2.  [[7.6]{.secno} The \``X-Frame-Options`\`
        header](document-lifecycle.html#the-x-frame-options-header)
    3.  [[7.7]{.secno} The \``Refresh`\`
        header](document-lifecycle.html#the-refresh-header)
    4.  [[7.8]{.secno} Browser user interface
        considerations](document-lifecycle.html#nav-traversal-ui)
    :::

### [7.5]{.secno} Document lifecycle[](#document-lifecycle){.self-link}

#### [7.5.1]{.secno} Shared document creation infrastructure[](#shared-document-creation-infrastructure){.self-link}

When loading a document using one of the below algorithms, we use the
following steps to [create and initialize a `Document`
object]{#initialise-the-document-object .dfn}, given a
[type](https://dom.spec.whatwg.org/#concept-document-type){#shared-document-creation-infrastructure:concept-document-type
x-internal="concept-document-type"} `type`{.variable}, [content
type](https://dom.spec.whatwg.org/#concept-document-content-type){#shared-document-creation-infrastructure:concept-document-content-type
x-internal="concept-document-content-type"} `contentType`{.variable},
and [navigation
params](browsing-the-web.html#navigation-params){#shared-document-creation-infrastructure:navigation-params}
`navigationParams`{.variable}:

[`Document`{#shared-document-creation-infrastructure:document}](dom.html#document)
objects are also created when [creating a new browsing context and
document](document-sequences.html#creating-a-new-browsing-context){#shared-document-creation-infrastructure:creating-a-new-browsing-context};
such [initial
`about:blank`](dom.html#is-initial-about:blank){#shared-document-creation-infrastructure:is-initial-about:blank}
[`Document`{#shared-document-creation-infrastructure:document-2}](dom.html#document)
are never created by this algorithm. Also, [browsing
context](document-sequences.html#concept-document-bc){#shared-document-creation-infrastructure:concept-document-bc}-less
[`Document`{#shared-document-creation-infrastructure:document-3}](dom.html#document)
objects can be created via various APIs, such as
[`document.implementation.createHTMLDocument()`{#shared-document-creation-infrastructure:dom-domimplementation-createhtmldocument}](https://dom.spec.whatwg.org/#dom-domimplementation-createhtmldocument){x-internal="dom-domimplementation-createhtmldocument"}.

1.  Let `browsingContext`{.variable} be the result of [obtaining a
    browsing context to use for a navigation
    response](browsers.html#obtain-browsing-context-navigation){#shared-document-creation-infrastructure:obtain-browsing-context-navigation}
    given `navigationParams`{.variable}.

    This can result in a [browsing context group
    switch](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy),
    in which case `browsingContext`{.variable} will be a
    [newly-created](document-sequences.html#creating-a-new-browsing-context){#shared-document-creation-infrastructure:creating-a-new-browsing-context-2}
    [browsing
    context](document-sequences.html#browsing-context){#shared-document-creation-infrastructure:browsing-context}
    instead of being `navigationParams`{.variable}\'s
    [navigable](browsing-the-web.html#navigation-params-navigable){#shared-document-creation-infrastructure:navigation-params-navigable}\'s
    [active browsing
    context](document-sequences.html#nav-bc){#shared-document-creation-infrastructure:nav-bc}.
    In such a case, the created
    [`Window`{#shared-document-creation-infrastructure:window}](nav-history-apis.html#window),
    [`Document`{#shared-document-creation-infrastructure:document-4}](dom.html#document),
    and
    [agent](https://tc39.es/ecma262/#sec-agents){#shared-document-creation-infrastructure:agent
    x-internal="agent"} will not end up being used; because the created
    [`Document`{#shared-document-creation-infrastructure:document-5}](dom.html#document)\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#shared-document-creation-infrastructure:concept-document-origin
    x-internal="concept-document-origin"} is
    [opaque](browsers.html#concept-origin-opaque){#shared-document-creation-infrastructure:concept-origin-opaque},
    we will end up creating a new
    [agent](https://tc39.es/ecma262/#sec-agents){#shared-document-creation-infrastructure:agent-2
    x-internal="agent"} and
    [`Window`{#shared-document-creation-infrastructure:window-2}](nav-history-apis.html#window)
    [later in this algorithm](#create-new-agent-and-window) to go along
    with the new
    [`Document`{#shared-document-creation-infrastructure:document-6}](dom.html#document).

2.  Let `permissionsPolicy`{.variable} be the result of [creating a
    permissions policy from a
    response](https://w3c.github.io/webappsec-feature-policy/#create-from-response){#shared-document-creation-infrastructure:creating-a-permissions-policy-from-a-response
    x-internal="creating-a-permissions-policy-from-a-response"} given
    `navigationParams`{.variable}\'s
    [navigable](browsing-the-web.html#navigation-params-navigable){#shared-document-creation-infrastructure:navigation-params-navigable-2}\'s
    [container](document-sequences.html#nav-container){#shared-document-creation-infrastructure:nav-container},
    `navigationParams`{.variable}\'s
    [origin](browsing-the-web.html#navigation-params-origin){#shared-document-creation-infrastructure:navigation-params-origin},
    and `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response}.
    [\[PERMISSIONSPOLICY\]](references.html#refsPERMISSIONSPOLICY)

    ::: note
    The [creating a permissions policy from a
    response](https://w3c.github.io/webappsec-feature-policy/#create-from-response){#shared-document-creation-infrastructure:creating-a-permissions-policy-from-a-response-2
    x-internal="creating-a-permissions-policy-from-a-response"}
    algorithm makes use of the passed
    [origin](browsers.html#concept-origin){#shared-document-creation-infrastructure:concept-origin}.
    If
    [`document.domain`{#shared-document-creation-infrastructure:dom-document-domain}](browsers.html#dom-document-domain)
    has been used for `navigationParams`{.variable}\'s
    [navigable](browsing-the-web.html#navigation-params-navigable){#shared-document-creation-infrastructure:navigation-params-navigable-3}\'s
    [container
    document](document-sequences.html#nav-container-document){#shared-document-creation-infrastructure:nav-container-document},
    then its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#shared-document-creation-infrastructure:concept-document-origin-2
    x-internal="concept-document-origin"} cannot be [same
    origin-domain](browsers.html#same-origin-domain){#shared-document-creation-infrastructure:same-origin-domain}
    with the passed origin, because these steps run before the
    `document`{.variable} is created, so it cannot itself yet have used
    [`document.domain`{#shared-document-creation-infrastructure:dom-document-domain-2}](browsers.html#dom-document-domain).
    Note that this means that Permissions Policy checks are less
    permissive compared to doing a [same
    origin](browsers.html#same-origin){#shared-document-creation-infrastructure:same-origin}
    check instead.

    See below for some examples of this in action.
    :::

3.  Let `creationURL`{.variable} be `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-2}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#shared-document-creation-infrastructure:concept-response-url
    x-internal="concept-response-url"}.

4.  If `navigationParams`{.variable}\'s
    [request](browsing-the-web.html#navigation-params-request){#shared-document-creation-infrastructure:navigation-params-request}
    is non-null, then set `creationURL`{.variable} to
    `navigationParams`{.variable}\'s
    [request](browsing-the-web.html#navigation-params-request){#shared-document-creation-infrastructure:navigation-params-request-2}\'s
    [current
    URL](https://fetch.spec.whatwg.org/#concept-request-current-url){#shared-document-creation-infrastructure:concept-request-current-url
    x-internal="concept-request-current-url"}.

5.  Let `window`{.variable} be null.

6.  If `browsingContext`{.variable}\'s [active
    document](document-sequences.html#active-document){#shared-document-creation-infrastructure:active-document}\'s
    [is initial
    `about:blank`](dom.html#is-initial-about:blank){#shared-document-creation-infrastructure:is-initial-about:blank-2}
    is true, and `browsingContext`{.variable}\'s [active
    document](document-sequences.html#active-document){#shared-document-creation-infrastructure:active-document-2}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#shared-document-creation-infrastructure:concept-document-origin-3
    x-internal="concept-document-origin"} is [same
    origin-domain](browsers.html#same-origin-domain){#shared-document-creation-infrastructure:same-origin-domain-2}
    with `navigationParams`{.variable}\'s
    [origin](browsing-the-web.html#navigation-params-origin){#shared-document-creation-infrastructure:navigation-params-origin-2},
    then set `window`{.variable} to `browsingContext`{.variable}\'s
    [active
    window](document-sequences.html#active-window){#shared-document-creation-infrastructure:active-window}.

    This means that both the [initial
    `about:blank`](dom.html#is-initial-about:blank){#shared-document-creation-infrastructure:is-initial-about:blank-3}
    [`Document`{#shared-document-creation-infrastructure:document-7}](dom.html#document),
    and the new
    [`Document`{#shared-document-creation-infrastructure:document-8}](dom.html#document)
    that is about to be created, will share the same
    [`Window`{#shared-document-creation-infrastructure:window-3}](nav-history-apis.html#window)
    object.

7.  ::: {#create-new-agent-and-window}
    Otherwise:

    1.  Let `oacHeader`{.variable} be the result of [getting a
        structured field
        value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header){#shared-document-creation-infrastructure:getting-a-structured-field-value
        x-internal="getting-a-structured-field-value"} given
        \`[`Origin-Agent-Cluster`{#shared-document-creation-infrastructure:origin-agent-cluster}](browsers.html#origin-agent-cluster)\`
        and \"`item`\" from `navigationParams`{.variable}\'s
        [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-3}\'s
        [header
        list](https://fetch.spec.whatwg.org/#concept-response-header-list){#shared-document-creation-infrastructure:concept-response-header-list
        x-internal="concept-response-header-list"}.

    2.  Let `requestsOAC`{.variable} be true if `oacHeader`{.variable}
        is not null and `oacHeader`{.variable}\[0\] is the
        [boolean](https://httpwg.org/specs/rfc8941.html#boolean){#shared-document-creation-infrastructure:http-structured-header-boolean
        x-internal="http-structured-header-boolean"} true; otherwise
        false.

    3.  If `navigationParams`{.variable}\'s [reserved
        environment](browsing-the-web.html#navigation-params-reserved-environment){#shared-document-creation-infrastructure:navigation-params-reserved-environment}
        is a [non-secure
        context](webappapis.html#non-secure-context){#shared-document-creation-infrastructure:non-secure-context},
        then set `requestsOAC`{.variable} to false.

    4.  Let `agent`{.variable} be the result of [obtaining a
        similar-origin window
        agent](webappapis.html#obtain-similar-origin-window-agent){#shared-document-creation-infrastructure:obtain-similar-origin-window-agent}
        given `navigationParams`{.variable}\'s
        [origin](browsing-the-web.html#navigation-params-origin){#shared-document-creation-infrastructure:navigation-params-origin-3},
        `browsingContext`{.variable}\'s
        [group](document-sequences.html#tlbc-group){#shared-document-creation-infrastructure:tlbc-group},
        and `requestsOAC`{.variable}.

    5.  Let `realmExecutionContext`{.variable} be the result of
        [creating a new
        realm](webappapis.html#creating-a-new-javascript-realm){#shared-document-creation-infrastructure:creating-a-new-javascript-realm}
        given `agent`{.variable} and the following customizations:

        - For the global object, create a new
          [`Window`{#shared-document-creation-infrastructure:window-4}](nav-history-apis.html#window)
          object.

        - For the global **this** binding, use
          `browsingContext`{.variable}\'s
          [`WindowProxy`{#shared-document-creation-infrastructure:windowproxy}](nav-history-apis.html#windowproxy)
          object.

    6.  Set `window`{.variable} to the [global
        object](webappapis.html#concept-realm-global){#shared-document-creation-infrastructure:concept-realm-global}
        of `realmExecutionContext`{.variable}\'s Realm component.

    7.  Let `topLevelCreationURL`{.variable} be
        `creationURL`{.variable}.

    8.  Let `topLevelOrigin`{.variable} be
        `navigationParams`{.variable}\'s
        [origin](browsing-the-web.html#navigation-params-origin){#shared-document-creation-infrastructure:navigation-params-origin-4}.

    9.  If `navigable`{.variable}\'s
        [container](document-sequences.html#nav-container){#shared-document-creation-infrastructure:nav-container-2}
        is not null, then:

        1.  Let `parentEnvironment`{.variable} be
            `navigable`{.variable}\'s
            [container](document-sequences.html#nav-container){#shared-document-creation-infrastructure:nav-container-3}\'s
            [relevant settings
            object](webappapis.html#relevant-settings-object){#shared-document-creation-infrastructure:relevant-settings-object}.

        2.  Set `topLevelCreationURL`{.variable} to
            `parentEnvironment`{.variable}\'s [top-level creation
            URL](webappapis.html#concept-environment-top-level-creation-url){#shared-document-creation-infrastructure:concept-environment-top-level-creation-url}.

        3.  Set `topLevelOrigin`{.variable} to
            `parentEnvironment`{.variable}\'s [top-level
            origin](webappapis.html#concept-environment-top-level-origin){#shared-document-creation-infrastructure:concept-environment-top-level-origin}.

    10. [Set up a window environment settings
        object](nav-history-apis.html#set-up-a-window-environment-settings-object){#shared-document-creation-infrastructure:set-up-a-window-environment-settings-object}
        with `creationURL`{.variable},
        `realmExecutionContext`{.variable},
        `navigationParams`{.variable}\'s [reserved
        environment](browsing-the-web.html#navigation-params-reserved-environment){#shared-document-creation-infrastructure:navigation-params-reserved-environment-2},
        `topLevelCreationURL`{.variable}, and
        `topLevelOrigin`{.variable}.

    This is the usual case, where the new
    [`Document`{#shared-document-creation-infrastructure:document-9}](dom.html#document)
    we\'re about to create gets a new
    [`Window`{#shared-document-creation-infrastructure:window-5}](nav-history-apis.html#window)
    to go along with it.
    :::

8.  Let `loadTimingInfo`{.variable} be a new [document load timing
    info](dom.html#document-load-timing-info){#shared-document-creation-infrastructure:document-load-timing-info}
    with its [navigation start
    time](dom.html#navigation-start-time){#shared-document-creation-infrastructure:navigation-start-time}
    set to `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-4}\'s
    [timing
    info](https://fetch.spec.whatwg.org/#concept-response-timing-info){#shared-document-creation-infrastructure:concept-response-timing-info
    x-internal="concept-response-timing-info"}\'s [start
    time](https://fetch.spec.whatwg.org/#fetch-timing-info-start-time){#shared-document-creation-infrastructure:fetch-timing-info-start-time
    x-internal="fetch-timing-info-start-time"}.

9.  Let `document`{.variable} be a new
    [`Document`{#shared-document-creation-infrastructure:document-10}](dom.html#document),
    with

    [type](https://dom.spec.whatwg.org/#concept-document-type){#shared-document-creation-infrastructure:concept-document-type-2 x-internal="concept-document-type"}
    :   `type`{.variable}

    [content type](https://dom.spec.whatwg.org/#concept-document-content-type){#shared-document-creation-infrastructure:concept-document-content-type-2 x-internal="concept-document-content-type"}
    :   `contentType`{.variable}

    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#shared-document-creation-infrastructure:concept-document-origin-4 x-internal="concept-document-origin"}
    :   `navigationParams`{.variable}\'s
        [origin](browsing-the-web.html#navigation-params-origin){#shared-document-creation-infrastructure:navigation-params-origin-5}

    [browsing context](document-sequences.html#concept-document-bc){#shared-document-creation-infrastructure:concept-document-bc-2}
    :   `browsingContext`{.variable}

    [policy container](dom.html#concept-document-policy-container){#shared-document-creation-infrastructure:concept-document-policy-container}
    :   `navigationParams`{.variable}\'s [policy
        container](browsing-the-web.html#navigation-params-policy-container){#shared-document-creation-infrastructure:navigation-params-policy-container}

    [permissions policy](dom.html#concept-document-permissions-policy){#shared-document-creation-infrastructure:concept-document-permissions-policy}
    :   `permissionsPolicy`{.variable}

    [active sandboxing flag set](browsers.html#active-sandboxing-flag-set){#shared-document-creation-infrastructure:active-sandboxing-flag-set}
    :   `navigationParams`{.variable}\'s [final sandboxing flag
        set](browsing-the-web.html#navigation-params-sandboxing){#shared-document-creation-infrastructure:navigation-params-sandboxing}

    [opener policy](dom.html#concept-document-coop){#shared-document-creation-infrastructure:concept-document-coop}
    :   `navigationParams`{.variable}\'s [cross-origin opener
        policy](browsing-the-web.html#navigation-params-coop){#shared-document-creation-infrastructure:navigation-params-coop}

    [load timing info](dom.html#load-timing-info){#shared-document-creation-infrastructure:load-timing-info}
    :   `loadTimingInfo`{.variable}

    [was created via cross-origin redirects](dom.html#was-created-via-cross-origin-redirects){#shared-document-creation-infrastructure:was-created-via-cross-origin-redirects}
    :   `navigationParams`{.variable}\'s
        [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-5}\'s
        [has cross-origin
        redirects](https://fetch.spec.whatwg.org/#response-has-cross-origin-redirects){#shared-document-creation-infrastructure:concept-response-has-cross-origin-redirects
        x-internal="concept-response-has-cross-origin-redirects"}

    [during-loading navigation ID for WebDriver BiDi](dom.html#concept-document-navigation-id){#shared-document-creation-infrastructure:concept-document-navigation-id}
    :   `navigationParams`{.variable}\'s
        [id](browsing-the-web.html#navigation-params-id){#shared-document-creation-infrastructure:navigation-params-id}

    [URL](https://dom.spec.whatwg.org/#concept-document-url){#shared-document-creation-infrastructure:the-document's-address x-internal="the-document's-address"}
    :   `creationURL`{.variable}

    [current document readiness](dom.html#current-document-readiness){#shared-document-creation-infrastructure:current-document-readiness}
    :   \"`loading`\"

    [about base URL](dom.html#concept-document-about-base-url){#shared-document-creation-infrastructure:concept-document-about-base-url}
    :   `navigationParams`{.variable}\'s [about base
        URL](browsing-the-web.html#navigation-params-about-base-url){#shared-document-creation-infrastructure:navigation-params-about-base-url}

    [allow declarative shadow roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots){#shared-document-creation-infrastructure:concept-document-allow-declarative-shadow-roots x-internal="concept-document-allow-declarative-shadow-roots"}
    :   true

    [custom element registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#shared-document-creation-infrastructure:document-custom-element-registry x-internal="document-custom-element-registry"}
    :   a new
        [`CustomElementRegistry`{#shared-document-creation-infrastructure:customelementregistry}](custom-elements.html#customelementregistry)
        object

10. Set `window`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#shared-document-creation-infrastructure:concept-document-window}
    to `document`{.variable}.

11. [Run CSP initialization for a
    `Document`](https://w3c.github.io/webappsec-csp/#run-document-csp-initialization){#shared-document-creation-infrastructure:run-csp-initialization-for-a-document
    x-internal="run-csp-initialization-for-a-document"} given
    `document`{.variable}. [\[CSP\]](references.html#refsCSP)

12. If `navigationParams`{.variable}\'s
    [request](browsing-the-web.html#navigation-params-request){#shared-document-creation-infrastructure:navigation-params-request-3}
    is non-null, then:

    1.  Set `document`{.variable}\'s
        [referrer](dom.html#the-document's-referrer){#shared-document-creation-infrastructure:the-document's-referrer}
        to the empty string.

    2.  Let `referrer`{.variable} be `navigationParams`{.variable}\'s
        [request](browsing-the-web.html#navigation-params-request){#shared-document-creation-infrastructure:navigation-params-request-4}\'s
        [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#shared-document-creation-infrastructure:concept-request-referrer
        x-internal="concept-request-referrer"}.

    3.  If `referrer`{.variable} is a [URL
        record](https://url.spec.whatwg.org/#concept-url){#shared-document-creation-infrastructure:url-record
        x-internal="url-record"}, then set `document`{.variable}\'s
        [referrer](dom.html#the-document's-referrer){#shared-document-creation-infrastructure:the-document's-referrer-2}
        to the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#shared-document-creation-infrastructure:concept-url-serializer
        x-internal="concept-url-serializer"} of `referrer`{.variable}.

        Per Fetch, `referrer`{.variable} will be either a [URL
        record](https://url.spec.whatwg.org/#concept-url){#shared-document-creation-infrastructure:url-record-2
        x-internal="url-record"} or \"`no-referrer`\" at this point.

13. If `navigationParams`{.variable}\'s [fetch
    controller](browsing-the-web.html#navigation-params-fetch-controller){#shared-document-creation-infrastructure:navigation-params-fetch-controller}
    is not null, then:

    1.  Let `fullTimingInfo`{.variable} be the result of [extracting the
        full timing
        info](https://fetch.spec.whatwg.org/#extract-full-timing-info){#shared-document-creation-infrastructure:extract-full-timing-info
        x-internal="extract-full-timing-info"} from
        `navigationParams`{.variable}\'s [fetch
        controller](browsing-the-web.html#navigation-params-fetch-controller){#shared-document-creation-infrastructure:navigation-params-fetch-controller-2}.

    2.  Let `redirectCount`{.variable} be 0 if
        `navigationParams`{.variable}\'s
        [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-6}\'s
        [has cross-origin
        redirects](https://fetch.spec.whatwg.org/#response-has-cross-origin-redirects){#shared-document-creation-infrastructure:concept-response-has-cross-origin-redirects-2
        x-internal="concept-response-has-cross-origin-redirects"} is
        true; otherwise `navigationParams`{.variable}\'s
        [request](browsing-the-web.html#navigation-params-request){#shared-document-creation-infrastructure:navigation-params-request-5}\'s
        [redirect
        count](https://fetch.spec.whatwg.org/#concept-request-redirect-count){#shared-document-creation-infrastructure:concept-request-redirect-count
        x-internal="concept-request-redirect-count"}.

    3.  [Create the navigation timing
        entry](https://w3c.github.io/navigation-timing/#dfn-create-the-navigation-timing-entry){#shared-document-creation-infrastructure:create-the-navigation-timing-entry
        x-internal="create-the-navigation-timing-entry"} for
        `document`{.variable}, given `fullTimingInfo`{.variable},
        `redirectCount`{.variable}, `navigationTimingType`{.variable},
        `navigationParams`{.variable}\'s
        [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-7}\'s
        [service worker timing
        info](https://fetch.spec.whatwg.org/#response-service-worker-timing-info){#shared-document-creation-infrastructure:concept-response-service-worker-timing-info
        x-internal="concept-response-service-worker-timing-info"}, and
        `navigationParams`{.variable}\'s
        [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-8}\'s
        [body
        info](https://fetch.spec.whatwg.org/#concept-response-body-info){#shared-document-creation-infrastructure:concept-response-body-info
        x-internal="concept-response-body-info"}.

14. [Create the navigation timing
    entry](https://w3c.github.io/navigation-timing/#dfn-create-the-navigation-timing-entry){#shared-document-creation-infrastructure:create-the-navigation-timing-entry-2
    x-internal="create-the-navigation-timing-entry"} for
    `document`{.variable}, with `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-9}\'s
    [timing
    info](https://fetch.spec.whatwg.org/#concept-response-timing-info){#shared-document-creation-infrastructure:concept-response-timing-info-2
    x-internal="concept-response-timing-info"},
    `redirectCount`{.variable}, `navigationParams`{.variable}\'s
    [navigation timing
    type](browsing-the-web.html#navigation-params-nav-timing-type){#shared-document-creation-infrastructure:navigation-params-nav-timing-type},
    and `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-10}\'s
    [service worker timing
    info](https://fetch.spec.whatwg.org/#response-service-worker-timing-info){#shared-document-creation-infrastructure:concept-response-service-worker-timing-info-2
    x-internal="concept-response-service-worker-timing-info"}.

15. If `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-11}
    has a
    \`[`Refresh`{#shared-document-creation-infrastructure:refresh}](#refresh)\`
    header, then:

    1.  Let `value`{.variable} be the [isomorphic
        decoding](https://infra.spec.whatwg.org/#isomorphic-decode){#shared-document-creation-infrastructure:isomorphic-decode
        x-internal="isomorphic-decode"} of the value of the header.

    2.  Run the [shared declarative refresh
        steps](semantics.html#shared-declarative-refresh-steps){#shared-document-creation-infrastructure:shared-declarative-refresh-steps}
        with `document`{.variable} and `value`{.variable}.

    We do not currently have a spec for how to handle multiple
    \`[`Refresh`{#shared-document-creation-infrastructure:refresh-2}](#refresh)\`
    headers. This is tracked as [issue
    #2900](https://github.com/whatwg/html/issues/2900).

16. If `navigationParams`{.variable}\'s [commit early
    hints](browsing-the-web.html#navigation-params-commit-early-hints){#shared-document-creation-infrastructure:navigation-params-commit-early-hints}
    is not null, then call `navigationParams`{.variable}\'s [commit
    early
    hints](browsing-the-web.html#navigation-params-commit-early-hints){#shared-document-creation-infrastructure:navigation-params-commit-early-hints-2}
    with `document`{.variable}.

17. [Process link
    headers](semantics.html#process-link-headers){#shared-document-creation-infrastructure:process-link-headers}
    given `document`{.variable}, `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#shared-document-creation-infrastructure:navigation-params-response-12},
    and \"`pre-media`\".

18. Return `document`{.variable}.

::: example
In this example, the child document is not allowed to use
[`PaymentRequest`{#shared-document-creation-infrastructure:paymentrequest}](https://w3c.github.io/payment-request/#dom-paymentrequest){x-internal="paymentrequest"},
despite being [same
origin-domain](browsers.html#same-origin-domain){#shared-document-creation-infrastructure:same-origin-domain-3}
at the time the child document tries to use it. At the time the child
document is initialized, only the parent document has set
[`document.domain`{#shared-document-creation-infrastructure:dom-document-domain-3}](browsers.html#dom-document-domain),
and the child document has not.

``` html
<!-- https://foo.example.com/a.html -->
<!doctype html>
<script>
document.domain = 'example.com';
</script>
<iframe src=b.html></iframe>
```

``` html
<!-- https://bar.example.com/b.html -->
<!doctype html>
<script>
document.domain = 'example.com'; // This happens after the document is initialized
new PaymentRequest(…); // Not allowed to use
</script>
```
:::

::: example
In this example, the child document *is* allowed to use
[`PaymentRequest`{#shared-document-creation-infrastructure:paymentrequest-2}](https://w3c.github.io/payment-request/#dom-paymentrequest){x-internal="paymentrequest"},
despite not being [same
origin-domain](browsers.html#same-origin-domain){#shared-document-creation-infrastructure:same-origin-domain-4}
at the time the child document tries to use it. At the time the child
document is initialized, none of the documents have set
[`document.domain`{#shared-document-creation-infrastructure:dom-document-domain-4}](browsers.html#dom-document-domain)
yet so [same
origin-domain](browsers.html#same-origin-domain){#shared-document-creation-infrastructure:same-origin-domain-5}
falls back to a normal [same
origin](browsers.html#same-origin){#shared-document-creation-infrastructure:same-origin-2}
check.

``` html
<!-- https://example.com/a.html -->
<!doctype html>
<iframe src=b.html></iframe>
<!-- The child document is now initialized, before the script below is run. -->
<script>
document.domain = 'example.com';
</script>
```

``` html
<!-- https://example.com/b.html -->
<!doctype html>
<script>
new PaymentRequest(…); // Allowed to use
</script>
```
:::

To [populate with `html`/`head`/`body`]{#populate-with-html/head/body
.dfn} given a
[`Document`{#shared-document-creation-infrastructure:document-11}](dom.html#document)
`document`{.variable}:

1.  Let `html`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#shared-document-creation-infrastructure:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    \"`html`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#shared-document-creation-infrastructure:html-namespace-2
    x-internal="html-namespace-2"}.

2.  Let `head`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#shared-document-creation-infrastructure:create-an-element-2
    x-internal="create-an-element"} given `document`{.variable},
    \"`head`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#shared-document-creation-infrastructure:html-namespace-2-2
    x-internal="html-namespace-2"}.

3.  Let `body`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#shared-document-creation-infrastructure:create-an-element-3
    x-internal="create-an-element"} given `document`{.variable},
    \"`body`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#shared-document-creation-infrastructure:html-namespace-2-3
    x-internal="html-namespace-2"}.

4.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#shared-document-creation-infrastructure:concept-node-append
    x-internal="concept-node-append"} `html`{.variable} to
    `document`{.variable}.

5.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#shared-document-creation-infrastructure:concept-node-append-2
    x-internal="concept-node-append"} `head`{.variable} to
    `html`{.variable}.

6.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#shared-document-creation-infrastructure:concept-node-append-3
    x-internal="concept-node-append"} `body`{.variable} to
    `html`{.variable}.

#### [7.5.2]{.secno} Loading HTML documents[](#read-html){.self-link} {#read-html}

To [load an HTML document]{#navigate-html .dfn}, given [navigation
params](browsing-the-web.html#navigation-params){#read-html:navigation-params}
`navigationParams`{.variable}:

1.  Let `document`{.variable} be the result of [creating and
    initializing a `Document`
    object](#initialise-the-document-object){#read-html:initialise-the-document-object}
    given \"`html`\", \"`text/html`\", and
    `navigationParams`{.variable}.

2.  If `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#read-html:the-document's-address
    x-internal="the-document's-address"} is
    [`about:blank`{#read-html:about:blank}](infrastructure.html#about:blank),
    then [populate with
    `html`/`head`/`body`](#populate-with-html/head/body){#read-html:populate-with-html/head/body}
    given `document`{.variable}.

    This special case, where even non-[initial
    `about:blank`](dom.html#is-initial-about:blank){#read-html:is-initial-about:blank}
    [`Document`{#read-html:document}](dom.html#document)s are
    synchronously given their element nodes, is necessary for compatible
    with deployed content. In other words, it is not compatible to
    instead go down the \"otherwise\" branch and feed the empty [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#read-html:byte-sequence
    x-internal="byte-sequence"} into an [HTML
    parser](parsing.html#html-parser){#read-html:html-parser} to
    asynchronously populate `document`{.variable}.

3.  Otherwise, create an [HTML
    parser](parsing.html#html-parser){#read-html:html-parser-2} and
    associate it with the `document`{.variable}. Each
    [task](webappapis.html#concept-task){#read-html:concept-task} that
    the [networking task
    source](webappapis.html#networking-task-source){#read-html:networking-task-source}
    places on the [task
    queue](webappapis.html#task-queue){#read-html:task-queue} while
    fetching runs must then fill the parser\'s [input byte
    stream](parsing.html#the-input-byte-stream){#read-html:the-input-byte-stream}
    with the fetched bytes and cause the [HTML
    parser](parsing.html#html-parser){#read-html:html-parser-3} to
    perform the appropriate processing of the input stream.

    The first
    [task](webappapis.html#concept-task){#read-html:concept-task-2} that
    the [networking task
    source](webappapis.html#networking-task-source){#read-html:networking-task-source-2}
    places on the [task
    queue](webappapis.html#task-queue){#read-html:task-queue-2} while
    fetching runs must [process link
    headers](semantics.html#process-link-headers){#read-html:process-link-headers}
    given `document`{.variable}, `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-html:navigation-params-response},
    and \"`media`\", after the task has been processed by the [HTML
    parser](parsing.html#html-parser){#read-html:html-parser-4}.

    Before any script execution occurs, the user agent must wait for
    [scripts may run for the newly-created
    document](browsing-the-web.html#scripts-may-run-for-the-newly-created-document){#read-html:scripts-may-run-for-the-newly-created-document}
    to be true for `document`{.variable}.

    The [input byte
    stream](parsing.html#the-input-byte-stream){#read-html:the-input-byte-stream-2}
    converts bytes into characters for use in the
    [tokenizer](parsing.html#tokenization){#read-html:tokenization}.
    This process relies, in part, on character encoding information
    found in the real [Content-Type
    metadata](urls-and-fetching.html#content-type){#read-html:content-type}
    of the resource; the computed type is not used for this purpose.

    When no more bytes are available, the user agent must [queue a
    global
    task](webappapis.html#queue-a-global-task){#read-html:queue-a-global-task}
    on the [networking task
    source](webappapis.html#networking-task-source){#read-html:networking-task-source-3}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#read-html:concept-relevant-global}
    to have the parser process the implied EOF character, which
    eventually causes a
    [`load`{#read-html:event-load}](indices.html#event-load) event to be
    fired.

4.  Return `document`{.variable}.

#### [7.5.3]{.secno} [Loading XML documents]{.dfn}[](#read-xml){.self-link} {#read-xml}

When faced with displaying an XML file inline, provided [navigation
params](browsing-the-web.html#navigation-params){#read-xml:navigation-params}
`navigationParams`{.variable} and a string `type`{.variable}, user
agents must follow the requirements defined in XML and Namespaces in
XML, XML Media Types, DOM, and other relevant specifications to [create
and initialize a `Document`
object](#initialise-the-document-object){#read-xml:initialise-the-document-object}
`document`{.variable}, given \"`xml`\", `type`{.variable}, and
`navigationParams`{.variable}, and return that
[`Document`{#read-xml:document}](dom.html#document). They must also
create a corresponding [XML
parser](xhtml.html#xml-parser){#read-xml:xml-parser}.
[\[XML\]](references.html#refsXML)
[\[XMLNS\]](references.html#refsXMLNS)
[\[RFC7303\]](references.html#refsRFC7303)
[\[DOM\]](references.html#refsDOM)

At the time of writing, the XML specification community had not actually
yet specified how XML and the DOM interact.

The first [task](webappapis.html#concept-task){#read-xml:concept-task}
that the [networking task
source](webappapis.html#networking-task-source){#read-xml:networking-task-source}
places on the [task
queue](webappapis.html#task-queue){#read-xml:task-queue} while fetching
runs must [process link
headers](semantics.html#process-link-headers){#read-xml:process-link-headers}
given `document`{.variable}, `navigationParams`{.variable}\'s
[response](browsing-the-web.html#navigation-params-response){#read-xml:navigation-params-response},
and \"`media`\", after the task has been processed by the [XML
parser](xhtml.html#xml-parser){#read-xml:xml-parser-2}.

The actual HTTP headers and other metadata, not the headers as mutated
or implied by the algorithms given in this specification, are the ones
that must be used when determining the character encoding according to
the rules given in the above specifications. Once the character encoding
is established, the [document\'s character
encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#read-xml:document's-character-encoding
x-internal="document's-character-encoding"} must be set to that
character encoding.

Before any script execution occurs, the user agent must wait for
[scripts may run for the newly-created
document](browsing-the-web.html#scripts-may-run-for-the-newly-created-document){#read-xml:scripts-may-run-for-the-newly-created-document}
to be true for the newly-created
[`Document`{#read-xml:document-2}](dom.html#document).

Once parsing is complete, the user agent must set
`document`{.variable}\'s [during-loading navigation ID for WebDriver
BiDi](dom.html#concept-document-navigation-id){#read-xml:concept-document-navigation-id}
to null.

For HTML documents this is reset when parsing is complete, after firing
the load event.

Error messages from the parse process (e.g., XML namespace
well-formedness errors) may be reported inline by mutating the
[`Document`{#read-xml:document-3}](dom.html#document).

#### [7.5.4]{.secno} Loading text documents[](#read-text){.self-link} {#read-text}

To [load a text document]{#navigate-text .dfn}, given a [navigation
params](browsing-the-web.html#navigation-params){#read-text:navigation-params}
`navigationParams`{.variable} and a string `type`{.variable}:

1.  Let `document`{.variable} be the result of [creating and
    initializing a `Document`
    object](#initialise-the-document-object){#read-text:initialise-the-document-object}
    given \"`html`\", `type`{.variable}, and
    `navigationParams`{.variable}.

2.  Set `document`{.variable}\'s [parser cannot change the mode
    flag](parsing.html#parser-cannot-change-the-mode-flag){#read-text:parser-cannot-change-the-mode-flag}
    to true.

3.  Set `document`{.variable}\'s
    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#read-text:concept-document-mode
    x-internal="concept-document-mode"} to \"`no-quirks`\".

4.  Create an [HTML
    parser](parsing.html#html-parser){#read-text:html-parser} and
    associate it with the `document`{.variable}. Act as if the tokenizer
    had emitted a start tag token with the tag name \"pre\" followed by
    a single U+000A LINE FEED (LF) character, and switch the [HTML
    parser](parsing.html#html-parser){#read-text:html-parser-2}\'s
    tokenizer to the [PLAINTEXT
    state](parsing.html#plaintext-state){#read-text:plaintext-state}.
    Each [task](webappapis.html#concept-task){#read-text:concept-task}
    that the [networking task
    source](webappapis.html#networking-task-source){#read-text:networking-task-source}
    places on the [task
    queue](webappapis.html#task-queue){#read-text:task-queue} while
    fetching runs must then fill the parser\'s [input byte
    stream](parsing.html#the-input-byte-stream){#read-text:the-input-byte-stream}
    with the fetched bytes and cause the [HTML
    parser](parsing.html#html-parser){#read-text:html-parser-3} to
    perform the appropriate processing of the input stream.

    `document`{.variable}\'s
    [encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#read-text:document's-character-encoding
    x-internal="document's-character-encoding"} must be set to the
    character encoding used to decode the document during parsing.

    The first
    [task](webappapis.html#concept-task){#read-text:concept-task-2} that
    the [networking task
    source](webappapis.html#networking-task-source){#read-text:networking-task-source-2}
    places on the [task
    queue](webappapis.html#task-queue){#read-text:task-queue-2} while
    fetching runs must [process link
    headers](semantics.html#process-link-headers){#read-text:process-link-headers}
    given `document`{.variable}, `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-text:navigation-params-response},
    and \"`media`\", after the task has been processed by the [HTML
    parser](parsing.html#html-parser){#read-text:html-parser-4}.

    Before any script execution occurs, the user agent must wait for
    [scripts may run for the newly-created
    document](browsing-the-web.html#scripts-may-run-for-the-newly-created-document){#read-text:scripts-may-run-for-the-newly-created-document}
    to be true for `document`{.variable}.

    When no more bytes are available, the user agent must [queue a
    global
    task](webappapis.html#queue-a-global-task){#read-text:queue-a-global-task}
    on the [networking task
    source](webappapis.html#networking-task-source){#read-text:networking-task-source-3}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#read-text:concept-relevant-global}
    to have the parser process the implied EOF character, which
    eventually causes a
    [`load`{#read-text:event-load}](indices.html#event-load) event to be
    fired.

5.  User agents may add content to the
    [`head`{#read-text:the-head-element}](semantics.html#the-head-element)
    element of `document`{.variable}, e.g., linking to a style sheet,
    providing script, or giving the document a
    [`title`{#read-text:the-title-element}](semantics.html#the-title-element).

    In particular, if the user agent supports the `Format=Flowed`
    feature of RFC 3676 then the user agent would need to apply extra
    styling to cause the text to wrap correctly and to handle the
    quoting feature. This could be performed using, e.g., a CSS
    extension.

6.  Return `document`{.variable}.

The rules for how to convert the bytes of the plain text document into
actual characters, and the rules for actually rendering the text to the
user, are defined by the specifications for the [computed MIME
type](https://mimesniff.spec.whatwg.org/#computed-mime-type){#read-text:content-type-sniffing-2
x-internal="content-type-sniffing-2"} of the resource (i.e.,
`type`{.variable}).

#### [7.5.5]{.secno} Loading [`multipart/x-mixed-replace`{#read-multipart-x-mixed-replace:multipart/x-mixed-replace}](iana.html#multipart/x-mixed-replace) documents[](#read-multipart-x-mixed-replace){.self-link} {#read-multipart-x-mixed-replace}

To [load a `multipart/x-mixed-replace`
document]{#navigate-multipart-x-mixed-replace .dfn}, given [navigation
params](browsing-the-web.html#navigation-params){#read-multipart-x-mixed-replace:navigation-params}
`navigationParams`{.variable}, [source snapshot
params](browsing-the-web.html#source-snapshot-params){#read-multipart-x-mixed-replace:source-snapshot-params}
`sourceSnapshotParams`{.variable}, and
[origin](browsers.html#concept-origin){#read-multipart-x-mixed-replace:concept-origin}
`initiatorOrigin`{.variable}:

1.  Parse `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-multipart-x-mixed-replace:navigation-params-response}\'s
    [body](https://fetch.spec.whatwg.org/#concept-response-body){#read-multipart-x-mixed-replace:concept-response-body
    x-internal="concept-response-body"} using the rules for multipart
    types. [\[RFC2046\]](references.html#refsRFC2046)

2.  Let `firstPartNavigationParams`{.variable} be a copy of
    `navigationParams`{.variable}.

3.  Set `firstPartNavigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-multipart-x-mixed-replace:navigation-params-response-2}
    to a new
    [response](https://fetch.spec.whatwg.org/#concept-response){#read-multipart-x-mixed-replace:concept-response
    x-internal="concept-response"} representing the first part of
    `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-multipart-x-mixed-replace:navigation-params-response-3}\'s
    [body](https://fetch.spec.whatwg.org/#concept-response-body){#read-multipart-x-mixed-replace:concept-response-body-2
    x-internal="concept-response-body"}\'s multipart stream.

4.  Let `document`{.variable} be the result of [loading a
    document](browsing-the-web.html#loading-a-document){#read-multipart-x-mixed-replace:loading-a-document}
    given `firstPartNavigationParams`{.variable},
    `sourceSnapshotParams`{.variable}, and `initiatorOrigin`{.variable}.

    For each additional body part obtained from
    `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-multipart-x-mixed-replace:navigation-params-response-4},
    the user agent must
    [navigate](browsing-the-web.html#navigate){#read-multipart-x-mixed-replace:navigate}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#read-multipart-x-mixed-replace:node-navigable}
    to `navigationParams`{.variable}\'s
    [request](browsing-the-web.html#navigation-params-request){#read-multipart-x-mixed-replace:navigation-params-request}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#read-multipart-x-mixed-replace:concept-request-url
    x-internal="concept-request-url"}, using `document`{.variable}, with
    *[response](browsing-the-web.html#navigation-response)* set to
    `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-multipart-x-mixed-replace:navigation-params-response-5}
    and *[historyHandling](browsing-the-web.html#navigation-hh)* set to
    \"[`replace`{#read-multipart-x-mixed-replace:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

5.  Return `document`{.variable}.

For the purposes of algorithms processing these body parts as if they
were complete stand-alone resources, the user agent must act as if there
were no more bytes for those resources whenever the boundary following
the body part is reached.

Thus,
[`load`{#read-multipart-x-mixed-replace:event-load}](indices.html#event-load)
events (and for that matter
[`unload`{#read-multipart-x-mixed-replace:event-unload}](indices.html#event-unload)
events) do fire for each body part loaded.

#### [7.5.6]{.secno} Loading media documents[](#read-media){.self-link} {#read-media}

To [load a media document]{#navigate-media .dfn}, given
`navigationParams`{.variable} and a string `type`{.variable}:

1.  Let `document`{.variable} be the result of [creating and
    initializing a `Document`
    object](#initialise-the-document-object){#read-media:initialise-the-document-object}
    given \"`html`\", `type`{.variable}, and
    `navigationParams`{.variable}.

2.  Set `document`{.variable}\'s
    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#read-media:concept-document-mode
    x-internal="concept-document-mode"} to \"`no-quirks`\".

3.  [Populate with
    `html`/`head`/`body`](#populate-with-html/head/body){#read-media:populate-with-html/head/body}
    given `document`{.variable}.

4.  Append an element `host element`{.variable} for the media, as
    described below, to the
    [`body`{#read-media:the-body-element}](sections.html#the-body-element)
    element.

5.  Set the appropriate attribute of the element
    `host element`{.variable}, as described below, to the address of the
    image, video, or audio resource.

6.  User agents may add content to the
    [`head`{#read-media:the-head-element}](semantics.html#the-head-element)
    element of `document`{.variable}, or attributes to
    `host element`{.variable}, e.g., to link to a style sheet, to
    provide a script, to give the document a
    [`title`{#read-media:the-title-element}](semantics.html#the-title-element),
    or to make the media
    [autoplay](media.html#attr-media-autoplay){#read-media:attr-media-autoplay}.

7.  [Process link
    headers](semantics.html#process-link-headers){#read-media:process-link-headers}
    given `document`{.variable}, `navigationParams`{.variable}\'s
    [response](browsing-the-web.html#navigation-params-response){#read-media:navigation-params-response},
    and \"`media`\".

8.  Act as if the user agent had [stopped
    parsing](parsing.html#stop-parsing){#read-media:stop-parsing}
    `document`{.variable}.

9.  Return `document`{.variable}.

The element `host element`{.variable} to create for the media is the
element given in the table below in the second cell of the row whose
first cell describes the media. The appropriate attribute to set is the
one given by the third cell in that same row.

Type of media

Element for the media

Appropriate attribute

Image

[`img`{#read-media:the-img-element}](embedded-content.html#the-img-element)

[`src`{#read-media:attr-img-src}](embedded-content.html#attr-img-src)

Video

[`video`{#read-media:the-video-element}](media.html#the-video-element)

[`src`{#read-media:attr-media-src}](media.html#attr-media-src)

Audio

[`audio`{#read-media:the-audio-element}](media.html#the-audio-element)

[`src`{#read-media:attr-media-src-2}](media.html#attr-media-src)

Before any script execution occurs, the user agent must wait for
[scripts may run for the newly-created
document](browsing-the-web.html#scripts-may-run-for-the-newly-created-document){#read-media:scripts-may-run-for-the-newly-created-document}
to be true for the
[`Document`{#read-media:document}](dom.html#document).

#### [7.5.7]{.secno} []{#read-ua-plugin}[]{#navigate-plugin}[Loading a document for inline content that doesn\'t have a DOM]{.dfn}[](#read-ua-inline){.self-link} {#read-ua-inline}

When the user agent is to create a document to display a user agent page
or PDF viewer inline, provided a
[navigable](document-sequences.html#navigable){#read-ua-inline:navigable}
`navigable`{.variable}, a [navigation
ID](browsing-the-web.html#navigation-id){#read-ua-inline:navigation-id}
`navigationId`{.variable}, a
[`NavigationTimingType`{#read-ua-inline:navigationtimingtype}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
`navTimingType`{.variable}, and a [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#read-ua-inline:user-navigation-involvement}
`userInvolvement`{.variable}, the user agent should:

1.  Let `origin`{.variable} be a new [opaque
    origin](browsers.html#concept-origin-opaque){#read-ua-inline:concept-origin-opaque}.

2.  Let `coop`{.variable} be a new [opener
    policy](browsers.html#cross-origin-opener-policy){#read-ua-inline:cross-origin-opener-policy}.

3.  Let `coopEnforcementResult`{.variable} be a new [opener policy
    enforcement
    result](browsers.html#coop-enforcement-result){#read-ua-inline:coop-enforcement-result}
    with

    [url](browsers.html#coop-enforcement-url){#read-ua-inline:coop-enforcement-url}
    :   `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#read-ua-inline:concept-response-url
        x-internal="concept-response-url"}

    [origin](browsers.html#coop-enforcement-origin){#read-ua-inline:coop-enforcement-origin}
    :   `origin`{.variable}

    [opener policy](browsers.html#coop-enforcement-coop){#read-ua-inline:coop-enforcement-coop}
    :   `coop`{.variable}

4.  Let `navigationParams`{.variable} be a new [navigation
    params](browsing-the-web.html#navigation-params){#read-ua-inline:navigation-params}
    with

    [id](browsing-the-web.html#navigation-params-id){#read-ua-inline:navigation-params-id}
    :   `navigationId`{.variable}

    [navigable](browsing-the-web.html#navigation-params-navigable){#read-ua-inline:navigation-params-navigable}
    :   `navigable`{.variable}

    [request](browsing-the-web.html#navigation-params-request){#read-ua-inline:navigation-params-request}
    :   null

    [response](browsing-the-web.html#navigation-params-response){#read-ua-inline:navigation-params-response}
    :   a new
        [response](https://fetch.spec.whatwg.org/#concept-response){#read-ua-inline:concept-response
        x-internal="concept-response"}

    [origin](browsing-the-web.html#navigation-params-origin){#read-ua-inline:navigation-params-origin}
    :   `origin`{.variable}

    [fetch controller](browsing-the-web.html#navigation-params-fetch-controller){#read-ua-inline:navigation-params-fetch-controller}
    :   null

    [commit early hints](browsing-the-web.html#navigation-params-commit-early-hints){#read-ua-inline:navigation-params-commit-early-hints}
    :   null

    [COOP enforcement result](browsing-the-web.html#navigation-params-coop-enforcement-result){#read-ua-inline:navigation-params-coop-enforcement-result}
    :   `coopEnforcementResult`{.variable}

    [reserved environment](browsing-the-web.html#navigation-params-reserved-environment){#read-ua-inline:navigation-params-reserved-environment}
    :   null

    [policy container](browsing-the-web.html#navigation-params-policy-container){#read-ua-inline:navigation-params-policy-container}
    :   a new [policy
        container](browsers.html#policy-container){#read-ua-inline:policy-container}

    [final sandboxing flag set](browsing-the-web.html#navigation-params-sandboxing){#read-ua-inline:navigation-params-sandboxing}
    :   an empty set

    [opener policy](browsing-the-web.html#navigation-params-coop){#read-ua-inline:navigation-params-coop}
    :   `coop`{.variable}

    [navigation timing type](browsing-the-web.html#navigation-params-nav-timing-type){#read-ua-inline:navigation-params-nav-timing-type}
    :   `navTimingType`{.variable}

    [about base URL](browsing-the-web.html#navigation-params-about-base-url){#read-ua-inline:navigation-params-about-base-url}
    :   null

    [user involvement](browsing-the-web.html#navigation-params-user-involvement){#read-ua-inline:navigation-params-user-involvement}
    :   `userInvolvement`{.variable}

5.  Let `document`{.variable} be the result of [creating and
    initializing a `Document`
    object](#initialise-the-document-object){#read-ua-inline:initialise-the-document-object}
    given \"`html`\", \"`text/html`\", and
    `navigationParams`{.variable}.

6.  Either associate `document`{.variable} with a custom rendering that
    is not rendered using the normal
    [`Document`{#read-ua-inline:document}](dom.html#document) rendering
    rules, or mutate `document`{.variable} until it represents the
    content the user agent wants to render.

7.  Return `document`{.variable}.

Because we ensure the resulting
[`Document`{#read-ua-inline:document-2}](dom.html#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#read-ua-inline:concept-document-origin
x-internal="concept-document-origin"} is
[opaque](browsers.html#concept-origin-opaque){#read-ua-inline:concept-origin-opaque-2},
and the resulting
[`Document`{#read-ua-inline:document-3}](dom.html#document) cannot run
script with access to the DOM, the existence and properties of this
[`Document`{#read-ua-inline:document-4}](dom.html#document) are not
observable to web developer code. This means that most of the above
values, e.g., the
[`text/html`{#read-ua-inline:text/html}](iana.html#text/html)
[type](https://dom.spec.whatwg.org/#concept-document-type){#read-ua-inline:concept-document-type
x-internal="concept-document-type"}, do not matter. Similarly, most of
the items in `navigationParams`{.variable} don\'t have any observable
effect, besides preventing the [`Document`-creation
algorithm](#initialise-the-document-object){#read-ua-inline:initialise-the-document-object-2}
from getting confused, and so are set to default values.

Once the page has been set up, the user agent must act as if it had
[stopped
parsing](parsing.html#stop-parsing){#read-ua-inline:stop-parsing}.

#### [7.5.8]{.secno} Finishing the loading process[](#loading-documents){.self-link} {#loading-documents}

A [`Document`{#loading-documents:document}](dom.html#document) has a
[completely loaded time]{#completely-loaded-time .dfn} (a time or null),
which is initially null.

A [`Document`{#loading-documents:document-2}](dom.html#document) is
considered [completely loaded]{#completely-loaded .dfn} if its
[completely loaded
time](#completely-loaded-time){#loading-documents:completely-loaded-time}
is non-null.

To [completely finish loading]{#completely-finish-loading .dfn} a
[`Document`{#loading-documents:document-3}](dom.html#document)
`document`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#loading-documents:assert
    x-internal="assert"}: `document`{.variable}\'s [browsing
    context](document-sequences.html#concept-document-bc){#loading-documents:concept-document-bc}
    is non-null.

2.  Set `document`{.variable}\'s [completely loaded
    time](#completely-loaded-time){#loading-documents:completely-loaded-time-2}
    to the current time.

3.  Let `container`{.variable} be `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#loading-documents:node-navigable}\'s
    [container](document-sequences.html#nav-container){#loading-documents:nav-container}.

    ::: note
    This will be null in the case where `document`{.variable} is the
    [initial
    `about:blank`](dom.html#is-initial-about:blank){#loading-documents:is-initial-about:blank}
    [`Document`{#loading-documents:document-4}](dom.html#document) in a
    [`frame`{#loading-documents:frame}](obsolete.html#frame) or
    [`iframe`{#loading-documents:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
    since at the point of [browsing context
    creation](document-sequences.html#creating-a-new-browsing-context){#loading-documents:creating-a-new-browsing-context}
    which calls this algorithm, the container relationship has not yet
    been established. (That happens in a subsequent step of [create a
    new child
    navigable](document-sequences.html#create-a-new-child-navigable){#loading-documents:create-a-new-child-navigable}.)

    The consequence of this is that the following steps do nothing,
    i.e., we do not fire an asynchronous
    [`load`{#loading-documents:event-load}](indices.html#event-load)
    event on the container element for such cases. Instead, a
    synchronous
    [`load`{#loading-documents:event-load-2}](indices.html#event-load)
    event is fired in a special initial-insertion case when [processing
    the `iframe`
    attributes](iframe-embed-object.html#process-the-iframe-attributes){#loading-documents:process-the-iframe-attributes}.
    :::

4.  If `container`{.variable} is an
    [`iframe`{#loading-documents:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
    element, then [queue an element
    task](webappapis.html#queue-an-element-task){#loading-documents:queue-an-element-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#loading-documents:dom-manipulation-task-source}
    given `container`{.variable} to run the [iframe load event
    steps](iframe-embed-object.html#iframe-load-event-steps){#loading-documents:iframe-load-event-steps}
    given `container`{.variable}.

5.  Otherwise, if `container`{.variable} is non-null, then [queue an
    element
    task](webappapis.html#queue-an-element-task){#loading-documents:queue-an-element-task-2}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#loading-documents:dom-manipulation-task-source-2}
    given `container`{.variable} to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-documents:concept-event-fire
    x-internal="concept-event-fire"} named
    [`load`{#loading-documents:event-load-3}](indices.html#event-load)
    at `container`{.variable}.

#### [7.5.9]{.secno} Unloading documents[](#unloading-documents){.self-link}

A [`Document`{#unloading-documents:document}](dom.html#document) has a
[*salvageable*]{#concept-document-salvageable .dfn} state, which must
initially be true, and a [page showing]{#page-showing .dfn} boolean,
which is initially false. The [page
showing](#page-showing){#unloading-documents:page-showing} boolean is
used to ensure that scripts receive
[`pageshow`{#unloading-documents:event-pageshow}](indices.html#event-pageshow)
and
[`pagehide`{#unloading-documents:event-pagehide}](indices.html#event-pagehide)
events in a consistent manner (e.g., that they never receive two
[`pagehide`{#unloading-documents:event-pagehide-2}](indices.html#event-pagehide)
events in a row without an intervening
[`pageshow`{#unloading-documents:event-pageshow-2}](indices.html#event-pageshow),
or vice versa).

A [`Document`{#unloading-documents:document-2}](dom.html#document) has a
[`DOMHighResTimeStamp`{#unloading-documents:domhighrestimestamp}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
[suspension time]{#suspension-time .dfn}, initially 0.

A [`Document`{#unloading-documents:document-3}](dom.html#document) has a
[list](https://infra.spec.whatwg.org/#list){#unloading-documents:list
x-internal="list"} of [suspended timer handles]{#suspended-timer-handles
.dfn}, initially empty.

[Event
loops](webappapis.html#event-loop){#unloading-documents:event-loop} have
a [termination nesting level]{#termination-nesting-level .dfn} counter,
which must initially be 0.

[`Document`{#unloading-documents:document-4}](dom.html#document) objects
have an [unload counter]{#unload-counter .dfn}, which is used to ignore
certain operations while the below algorithms run. Initially, the
counter must be set to zero.

To [unload]{#unload-a-document .dfn} a
[`Document`{#unloading-documents:document-5}](dom.html#document)
`oldDocument`{.variable}, given an optional
[`Document`{#unloading-documents:document-6}](dom.html#document)
`newDocument`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#unloading-documents:assert
    x-internal="assert"}: this is running as part of a
    [task](webappapis.html#concept-task){#unloading-documents:concept-task}
    queued on `oldDocument`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#unloading-documents:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#unloading-documents:concept-agent-event-loop}.

2.  Let `unloadTimingInfo`{.variable} be a new [document unload timing
    info](dom.html#document-unload-timing-info){#unloading-documents:document-unload-timing-info}.

3.  If `newDocument`{.variable} is not given, then set
    `unloadTimingInfo`{.variable} to null.

    In this case there is no new document that needs to know about how
    long it took `oldDocument`{.variable} to unload.

4.  Otherwise, if `newDocument`{.variable}\'s [event
    loop](webappapis.html#event-loop){#unloading-documents:event-loop-2}
    is not `oldDocument`{.variable}\'s [event
    loop](webappapis.html#event-loop){#unloading-documents:event-loop-3},
    then the user agent may be
    [unloading](#unload-a-document){#unloading-documents:unload-a-document}
    `oldDocument`{.variable} [in
    parallel](infrastructure.html#in-parallel){#unloading-documents:in-parallel}.
    In that case, the user agent should set
    `unloadTimingInfo`{.variable} to null.

    In this case `newDocument`{.variable}\'s loading is not impacted by
    how long it takes to unload `oldDocument`{.variable}, so it would be
    meaningless to communicate that timing info.

5.  Let `intendToKeepInBfcache`{.variable} be true if the user agent
    intends to keep `oldDocument`{.variable} alive in a [session history
    entry](browsing-the-web.html#session-history-entry){#unloading-documents:session-history-entry},
    such that it can later be [used for history
    traversal](browsing-the-web.html#note-bfcache).

    This must be false if `oldDocument`{.variable} is not
    *[salvageable](#concept-document-salvageable)*, or if there are any
    descendants of `oldDocument`{.variable} which the user agent does
    not intend to keep alive in the same way (including due to their
    lack of *[salvageability](#concept-document-salvageable)*).

6.  Let `eventLoop`{.variable} be `oldDocument`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#unloading-documents:relevant-agent-2}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#unloading-documents:concept-agent-event-loop-2}.

7.  Increase `eventLoop`{.variable}\'s [termination nesting
    level](#termination-nesting-level){#unloading-documents:termination-nesting-level}
    by 1.

8.  Increase `oldDocument`{.variable}\'s [unload
    counter](#unload-counter){#unloading-documents:unload-counter} by 1.

9.  If `intendToKeepInBfcache`{.variable} is false, then set
    `oldDocument`{.variable}\'s
    *[salvageable](#concept-document-salvageable)* state to false.

10. If `oldDocument`{.variable}\'s [page
    showing](#page-showing){#unloading-documents:page-showing-2} is
    true:

    1.  Set `oldDocument`{.variable}\'s [page
        showing](#page-showing){#unloading-documents:page-showing-3} to
        false.

    2.  [Fire a page transition
        event](nav-history-apis.html#fire-a-page-transition-event){#unloading-documents:fire-a-page-transition-event}
        named
        [`pagehide`{#unloading-documents:event-pagehide-3}](indices.html#event-pagehide)
        at `oldDocument`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global}
        with `oldDocument`{.variable}\'s
        *[salvageable](#concept-document-salvageable)* state.

    3.  [Update the visibility
        state](interaction.html#update-the-visibility-state){#unloading-documents:update-the-visibility-state}
        of `oldDocument`{.variable} to \"`hidden`\".

11. If `unloadTimingInfo`{.variable} is not null, then set
    `unloadTimingInfo`{.variable}\'s [unload event start
    time](dom.html#unload-event-start-time){#unloading-documents:unload-event-start-time}
    to the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#unloading-documents:current-high-resolution-time
    x-internal="current-high-resolution-time"} given
    `newDocument`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-2},
    [coarsened](https://w3c.github.io/hr-time/#dfn-coarsen-time){#unloading-documents:coarsen-time
    x-internal="coarsen-time"} given `oldDocument`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#unloading-documents:relevant-settings-object}\'s
    [cross-origin isolated
    capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#unloading-documents:concept-settings-object-cross-origin-isolated-capability}.

12. If `oldDocument`{.variable}\'s
    *[salvageable](#concept-document-salvageable)* state is false, then
    [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#unloading-documents:concept-event-fire
    x-internal="concept-event-fire"} named
    [`unload`{#unloading-documents:event-unload}](indices.html#event-unload)
    at `oldDocument`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-3},
    with *legacy target override flag* set.

13. If `unloadTimingInfo`{.variable} is not null, then set
    `unloadTimingInfo`{.variable}\'s [unload event end
    time](dom.html#unload-event-end-time){#unloading-documents:unload-event-end-time}
    to the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#unloading-documents:current-high-resolution-time-2
    x-internal="current-high-resolution-time"} given
    `newDocument`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-4},
    [coarsened](https://w3c.github.io/hr-time/#dfn-coarsen-time){#unloading-documents:coarsen-time-2
    x-internal="coarsen-time"} given `oldDocument`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#unloading-documents:relevant-settings-object-2}\'s
    [cross-origin isolated
    capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#unloading-documents:concept-settings-object-cross-origin-isolated-capability-2}.

14. Decrease `eventLoop`{.variable}\'s [termination nesting
    level](#termination-nesting-level){#unloading-documents:termination-nesting-level-2}
    by 1.

15. Set `oldDocument`{.variable}\'s [suspension
    time](#suspension-time){#unloading-documents:suspension-time} to the
    [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#unloading-documents:current-high-resolution-time-3
    x-internal="current-high-resolution-time"} given
    `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-5}.

16. Set `oldDocument`{.variable}\'s [suspended timer
    handles](#suspended-timer-handles){#unloading-documents:suspended-timer-handles}
    to the result of [getting the
    keys](https://infra.spec.whatwg.org/#map-getting-the-keys){#unloading-documents:map-get-the-keys
    x-internal="map-get-the-keys"} for the [map of active
    timers](timers-and-user-prompts.html#map-of-active-timers){#unloading-documents:map-of-active-timers}.

17. Set `oldDocument`{.variable}\'s [has been scrolled by the
    user](browsing-the-web.html#has-been-scrolled-by-the-user){#unloading-documents:has-been-scrolled-by-the-user}
    to false.

18. Run any [unloading document cleanup
    steps](#unloading-document-cleanup-steps){#unloading-documents:unloading-document-cleanup-steps}
    for `oldDocument`{.variable} that are defined by this specification
    and [other applicable
    specifications](infrastructure.html#other-applicable-specifications){#unloading-documents:other-applicable-specifications}.

19. If `oldDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#unloading-documents:node-navigable}
    is a [top-level
    traversable](document-sequences.html#top-level-traversable){#unloading-documents:top-level-traversable},
    [build not restored reasons for a top-level traversable and its
    descendants](browsing-the-web.html#build-not-restored-reasons-for-a-top-level-traversable-and-its-descendants){#unloading-documents:build-not-restored-reasons-for-a-top-level-traversable-and-its-descendants}
    given `oldDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#unloading-documents:node-navigable-2}.

20. If `oldDocument`{.variable}\'s
    *[salvageable](#concept-document-salvageable)* state is false, then
    [destroy](#destroy-a-document){#unloading-documents:destroy-a-document}
    `oldDocument`{.variable}.

21. Decrease `oldDocument`{.variable}\'s [unload
    counter](#unload-counter){#unloading-documents:unload-counter-2} by
    1.

22. If `newDocument`{.variable} is given, `newDocument`{.variable}\'s
    [was created via cross-origin
    redirects](dom.html#was-created-via-cross-origin-redirects){#unloading-documents:was-created-via-cross-origin-redirects}
    is false, and `newDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#unloading-documents:concept-document-origin
    x-internal="concept-document-origin"} is the
    [same](browsers.html#same-origin){#unloading-documents:same-origin}
    as `oldDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#unloading-documents:concept-document-origin-2
    x-internal="concept-document-origin"}, then set
    `newDocument`{.variable}\'s [previous document unload
    timing](dom.html#previous-document-unload-timing){#unloading-documents:previous-document-unload-timing}
    to `unloadTimingInfo`{.variable}.

To [unload a document and its
descendants]{#unload-a-document-and-its-descendants .dfn}, given a
[`Document`{#unloading-documents:document-7}](dom.html#document)
`document`{.variable}, an optional
[`Document`{#unloading-documents:document-8}](dom.html#document)-or-null
`newDocument`{.variable} (default null), an optional set of steps
`afterAllUnloads`{.variable}, and an optional set of steps
`firePageSwapSteps`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#unloading-documents:assert-2
    x-internal="assert"}: this is running within
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#unloading-documents:node-navigable-3}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#unloading-documents:nav-traversable}\'s
    [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#unloading-documents:tn-session-history-traversal-queue}.

2.  Let `childNavigables`{.variable} be `document`{.variable}\'s [child
    navigables](document-sequences.html#child-navigable){#unloading-documents:child-navigable}.

3.  Let `numberUnloaded`{.variable} be 0.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#unloading-documents:list-iterate
    x-internal="list-iterate"} `childNavigable`{.variable} of
    `childNavigables`{.variable} [in what order?]{.XXX}, [queue a global
    task](webappapis.html#queue-a-global-task){#unloading-documents:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#unloading-documents:navigation-and-traversal-task-source}
    given `childNavigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#unloading-documents:nav-window}
    to perform the following steps:

    1.  Let `incrementUnloaded`{.variable} be an algorithm step which
        increments `numberUnloaded`{.variable}.

    2.  [Unload a document and its
        descendants](#unload-a-document-and-its-descendants){#unloading-documents:unload-a-document-and-its-descendants}
        given `childNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#unloading-documents:nav-document},
        null, and `incrementUnloaded`{.variable}.

5.  Wait until `numberUnloaded`{.variable} equals
    `childNavigable`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#unloading-documents:list-size
    x-internal="list-size"}.

6.  [Queue a global
    task](webappapis.html#queue-a-global-task){#unloading-documents:queue-a-global-task-2}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#unloading-documents:navigation-and-traversal-task-source-2}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-6}
    to perform the following steps:

    1.  If `firePageSwapSteps`{.variable} is given, then run
        `firePageSwapSteps`{.variable}.

    2.  [Unload](#unload-a-document){#unloading-documents:unload-a-document-2}
        `document`{.variable}, passing along `newDocument`{.variable} if
        it is not null.

    3.  If `afterAllUnloads`{.variable} was given, then run it.

This specification defines the following [unloading document cleanup
steps]{#unloading-document-cleanup-steps .dfn export=""}. Other
specifications can define more. Given a
[`Document`{#unloading-documents:document-9}](dom.html#document)
`document`{.variable}:

1.  Let `window`{.variable} be `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-7}.

2.  For each
    [`WebSocket`{#unloading-documents:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
    object `webSocket`{.variable} whose [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-8}
    is `window`{.variable}, [make
    disappear](https://websockets.spec.whatwg.org/#make-disappear){#unloading-documents:make-disappear
    x-internal="make-disappear"} `webSocket`{.variable}.

    If this affected any
    [`WebSocket`{#unloading-documents:websocket-2-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
    objects, then [make document
    unsalvageable](browsing-the-web.html#make-document-unsalvageable){#unloading-documents:make-document-unsalvageable}
    given `document`{.variable} and
    \"[`websocket`{#unloading-documents:blocking-websocket}](nav-history-apis.html#blocking-websocket)\".

3.  For each
    [`WebTransport`{#unloading-documents:webtransport}](https://w3c.github.io/webtransport/#webtransport){x-internal="webtransport"}
    object `transport`{.variable} whose [relevant global
    object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-9}
    is `window`{.variable}, run the [context cleanup
    steps](https://w3c.github.io/webtransport/#context-cleanup-steps){#unloading-documents:context-cleanup-steps
    x-internal="context-cleanup-steps"} given `transport`{.variable}.

4.  If `document`{.variable}\'s
    *[salvageable](#concept-document-salvageable)* state is false, then:

    1.  For each
        [`EventSource`{#unloading-documents:eventsource}](server-sent-events.html#eventsource)
        object `eventSource`{.variable} whose [relevant global
        object](webappapis.html#concept-relevant-global){#unloading-documents:concept-relevant-global-10}
        is equal to `window`{.variable}, [forcibly
        close](server-sent-events.html#concept-eventsource-forcibly-close){#unloading-documents:concept-eventsource-forcibly-close}
        `eventSource`{.variable}.

    2.  [Clear](https://infra.spec.whatwg.org/#map-clear){#unloading-documents:map-clear
        x-internal="map-clear"} `window`{.variable}\'s [map of active
        timers](timers-and-user-prompts.html#map-of-active-timers){#unloading-documents:map-of-active-timers-2}.

::: XXX
It would be better if specification authors sent a pull request to add
calls from here into their specifications directly, instead of using the
[unloading document cleanup
steps](#unloading-document-cleanup-steps){#unloading-documents:unloading-document-cleanup-steps-2}
hook, to ensure well-defined cross-specification call order. As of the
time of this writing the following specifications are known to have
[unloading document cleanup
steps](#unloading-document-cleanup-steps){#unloading-documents:unloading-document-cleanup-steps-3},
which will be run in an unspecified order: Fullscreen API, Web NFC,
WebDriver BiDi, Compute Pressure, File API, Media Capture and Streams,
Picture-in-Picture, Screen Orientation, Service Workers, WebLocks API,
WebAudio API, WebRTC. [\[FULLSCREEN\]](references.html#refsFULLSCREEN)
[\[WEBNFC\]](references.html#refsWEBNFC)
[\[WEBDRIVERBIDI\]](references.html#refsWEBDRIVERBIDI)
[\[COMPUTEPRESSURE\]](references.html#refsCOMPUTEPRESSURE)
[\[FILEAPI\]](references.html#refsFILEAPI)
[\[MEDIASTREAM\]](references.html#refsMEDIASTREAM)
[\[PICTUREINPICTURE\]](references.html#refsPICTUREINPICTURE)
[\[SCREENORIENTATION\]](references.html#refsSCREENORIENTATION)
[\[SW\]](references.html#refsSW)
[\[WEBLOCKS\]](references.html#refsWEBLOCKS)
[\[WEBAUDIO\]](references.html#refsWEBAUDIO)
[\[WEBRTC\]](references.html#refsWEBRTC)

[Issue #8906](https://github.com/whatwg/html/issues/8906) tracks the
work to make the order of these steps clear.
:::

#### [7.5.10]{.secno} Destroying documents[](#destroying-documents){.self-link}

To [destroy]{#destroy-a-document .dfn} a
[`Document`{#destroying-documents:document}](dom.html#document)
`document`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#destroying-documents:assert
    x-internal="assert"}: this is running as part of a
    [task](webappapis.html#concept-task){#destroying-documents:concept-task}
    queued on `document`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#destroying-documents:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#destroying-documents:concept-agent-event-loop}.

2.  [Abort](#abort-a-document){#destroying-documents:abort-a-document}
    `document`{.variable}.

3.  Set `document`{.variable}\'s
    *[salvageable](#concept-document-salvageable)* state to false.

4.  Let `ports`{.variable} be the list of
    [`MessagePort`{#destroying-documents:messageport}](web-messaging.html#messageport)s
    whose [relevant global
    object](webappapis.html#concept-relevant-global){#destroying-documents:concept-relevant-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#destroying-documents:concept-document-window}
    is `document`{.variable}.

5.  For each `port`{.variable} in `ports`{.variable},
    [disentangle](web-messaging.html#disentangle){#destroying-documents:disentangle}
    `port`{.variable}.

6.  Run any [unloading document cleanup
    steps](#unloading-document-cleanup-steps){#destroying-documents:unloading-document-cleanup-steps}
    for `document`{.variable} that are defined by this specification and
    [other applicable
    specifications](infrastructure.html#other-applicable-specifications){#destroying-documents:other-applicable-specifications}.

7.  Remove any
    [tasks](webappapis.html#concept-task){#destroying-documents:concept-task-2}
    whose
    [document](webappapis.html#concept-task-document){#destroying-documents:concept-task-document}
    is `document`{.variable} from any [task
    queue](webappapis.html#task-queue){#destroying-documents:task-queue}
    (without running those tasks).

8.  Set `document`{.variable}\'s [browsing
    context](document-sequences.html#concept-document-bc){#destroying-documents:concept-document-bc}
    to null.

9.  Set `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#destroying-documents:node-navigable}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#destroying-documents:nav-active-history-entry}\'s
    [document
    state](browsing-the-web.html#she-document-state){#destroying-documents:she-document-state}\'s
    [document](browsing-the-web.html#document-state-document){#destroying-documents:document-state-document}
    to null.

10. [Remove](https://infra.spec.whatwg.org/#list-remove){#destroying-documents:list-remove
    x-internal="list-remove"} `document`{.variable} from the [owner
    set](workers.html#concept-WorkerGlobalScope-owner-set){#destroying-documents:concept-WorkerGlobalScope-owner-set}
    of each
    [`WorkerGlobalScope`{#destroying-documents:workerglobalscope}](workers.html#workerglobalscope)
    object whose set
    [contains](https://infra.spec.whatwg.org/#list-contain){#destroying-documents:list-contains
    x-internal="list-contains"} `document`{.variable}.

11. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#destroying-documents:list-iterate
    x-internal="list-iterate"} `workletGlobalScope`{.variable} in
    `document`{.variable}\'s [worklet global
    scopes](worklets.html#concept-document-worklet-global-scopes){#destroying-documents:concept-document-worklet-global-scopes},
    [terminate](worklets.html#terminate-a-worklet-global-scope){#destroying-documents:terminate-a-worklet-global-scope}
    `workletGlobalScope`{.variable}.

Even after destruction, the
[`Document`{#destroying-documents:document-2}](dom.html#document) object
itself might still be accessible to script, in the case where we are
[destroying a child
navigable](document-sequences.html#destroy-a-child-navigable){#destroying-documents:destroy-a-child-navigable}.

To [destroy a document and its
descendants]{#destroy-a-document-and-its-descendants .dfn} given a
[`Document`{#destroying-documents:document-3}](dom.html#document)
`document`{.variable} and an optional set of steps
`afterAllDestruction`{.variable}, perform the following steps [in
parallel](infrastructure.html#in-parallel){#destroying-documents:in-parallel}:

1.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#destroying-documents:fully-active},
    then:

    1.  Let `reason`{.variable} be a string from [user-agent specific
        blocking
        reasons](nav-history-apis.html#ua-specific-blocking-reasons){#destroying-documents:ua-specific-blocking-reasons}.
        If none apply, then let `reason`{.variable} be
        \"[`masked`{#destroying-documents:blocking-masked}](nav-history-apis.html#blocking-masked)\".

    2.  [Make document
        unsalvageable](browsing-the-web.html#make-document-unsalvageable){#destroying-documents:make-document-unsalvageable}
        given `document`{.variable} and `reason`{.variable}.

    3.  If `document`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#destroying-documents:node-navigable-2}
        is a [top-level
        traversable](document-sequences.html#top-level-traversable){#destroying-documents:top-level-traversable},
        [build not restored reasons for a top-level traversable and its
        descendants](browsing-the-web.html#build-not-restored-reasons-for-a-top-level-traversable-and-its-descendants){#destroying-documents:build-not-restored-reasons-for-a-top-level-traversable-and-its-descendants}
        given `document`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#destroying-documents:node-navigable-3}.

2.  Let `childNavigables`{.variable} be `document`{.variable}\'s [child
    navigables](document-sequences.html#child-navigable){#destroying-documents:child-navigable}.

3.  Let `numberDestroyed`{.variable} be 0.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#destroying-documents:list-iterate-2
    x-internal="list-iterate"} `childNavigable`{.variable} of
    `childNavigables`{.variable} [in what order?]{.XXX}, [queue a global
    task](webappapis.html#queue-a-global-task){#destroying-documents:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#destroying-documents:navigation-and-traversal-task-source}
    given `childNavigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#destroying-documents:nav-window}
    to perform the following steps:

    1.  Let `incrementDestroyed`{.variable} be an algorithm step which
        increments `numberDestroyed`{.variable}.

    2.  [Destroy a document and its
        descendants](#destroy-a-document-and-its-descendants){#destroying-documents:destroy-a-document-and-its-descendants}
        given `childNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#destroying-documents:nav-document}
        and `incrementDestroyed`{.variable}.

5.  Wait until `numberDestroyed`{.variable} equals
    `childNavigable`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#destroying-documents:list-size
    x-internal="list-size"}.

6.  [Queue a global
    task](webappapis.html#queue-a-global-task){#destroying-documents:queue-a-global-task-2}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#destroying-documents:navigation-and-traversal-task-source-2}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#destroying-documents:concept-relevant-global-2}
    to perform the following steps:

    1.  [Destroy](#destroy-a-document){#destroying-documents:destroy-a-document}
        `document`{.variable}.

    2.  If `afterAllDestruction`{.variable} was given, then run it.

#### [7.5.11]{.secno} Aborting a document load[](#aborting-a-document-load){.self-link}

To [abort]{#abort-a-document .dfn} a
[`Document`{#aborting-a-document-load:document}](dom.html#document)
`document`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#aborting-a-document-load:assert
    x-internal="assert"}: this is running as part of a
    [task](webappapis.html#concept-task){#aborting-a-document-load:concept-task}
    queued on `document`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#aborting-a-document-load:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#aborting-a-document-load:concept-agent-event-loop}.

2.  Cancel any instances of the
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#aborting-a-document-load:concept-fetch
    x-internal="concept-fetch"} algorithm in the context of
    `document`{.variable}, discarding any
    [tasks](webappapis.html#concept-task){#aborting-a-document-load:concept-task-2}
    [queued](webappapis.html#queue-a-task){#aborting-a-document-load:queue-a-task}
    for them, and discarding any further data received from the network
    for them. If this resulted in any instances of the
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#aborting-a-document-load:concept-fetch-2
    x-internal="concept-fetch"} algorithm being canceled or any
    [queued](webappapis.html#queue-a-task){#aborting-a-document-load:queue-a-task-2}
    [tasks](webappapis.html#concept-task){#aborting-a-document-load:concept-task-3}
    or any network data getting discarded, then [make document
    unsalvageable](browsing-the-web.html#make-document-unsalvageable){#aborting-a-document-load:make-document-unsalvageable}
    given `document`{.variable} and
    \"[`fetch`{#aborting-a-document-load:blocking-fetch}](nav-history-apis.html#blocking-fetch)\".

3.  If `document`{.variable}\'s [during-loading navigation ID for
    WebDriver
    BiDi](dom.html#concept-document-navigation-id){#aborting-a-document-load:concept-document-navigation-id}
    is non-null, then:

    1.  Invoke [WebDriver BiDi navigation
        aborted](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-aborted){#aborting-a-document-load:webdriver-bidi-navigation-aborted
        x-internal="webdriver-bidi-navigation-aborted"} with
        `document`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#aborting-a-document-load:node-navigable}
        and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#aborting-a-document-load:webdriver-bidi-navigation-status
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#aborting-a-document-load:navigation-status-id
        x-internal="navigation-status-id"} is `document`{.variable}\'s
        [during-loading navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#aborting-a-document-load:concept-document-navigation-id-2},
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#aborting-a-document-load:navigation-status-status
        x-internal="navigation-status-status"} is
        \"[`canceled`{#aborting-a-document-load:navigation-status-canceled}](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled){x-internal="navigation-status-canceled"}\",
        and
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#aborting-a-document-load:navigation-status-url
        x-internal="navigation-status-url"} is `document`{.variable}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#aborting-a-document-load:the-document's-address
        x-internal="the-document's-address"}.

    2.  Set `document`{.variable}\'s [during-loading navigation ID for
        WebDriver
        BiDi](dom.html#concept-document-navigation-id){#aborting-a-document-load:concept-document-navigation-id-3}
        to null.

4.  If `document`{.variable} has an [active
    parser](dom.html#active-parser){#aborting-a-document-load:active-parser},
    then:

    1.  Set `document`{.variable}\'s [active parser was
        aborted](dynamic-markup-insertion.html#active-parser-was-aborted){#aborting-a-document-load:active-parser-was-aborted}
        to true.

    2.  [Abort that
        parser](parsing.html#abort-a-parser){#aborting-a-document-load:abort-a-parser}.

    3.  Set `document`{.variable}\'s
        *[salvageable](#concept-document-salvageable)* to false.

    4.  [Make document
        unsalvageable](browsing-the-web.html#make-document-unsalvageable){#aborting-a-document-load:make-document-unsalvageable-2}
        given `document`{.variable} and
        \"[`parser-aborted`{#aborting-a-document-load:blocking-parser-aborted}](nav-history-apis.html#blocking-parser-aborted)\".

To [abort a document and its
descendants]{#abort-a-document-and-its-descendants .dfn} given a
[`Document`{#aborting-a-document-load:document-2}](dom.html#document)
`document`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#aborting-a-document-load:assert-2
    x-internal="assert"}: this is running as part of a
    [task](webappapis.html#concept-task){#aborting-a-document-load:concept-task-4}
    queued on `document`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#aborting-a-document-load:relevant-agent-2}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#aborting-a-document-load:concept-agent-event-loop-2}.

2.  Let `descendantNavigables`{.variable} be `document`{.variable}\'s
    [descendant
    navigables](document-sequences.html#descendant-navigables){#aborting-a-document-load:descendant-navigables}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#aborting-a-document-load:list-iterate
    x-internal="list-iterate"} `descendantNavigable`{.variable} of
    `descendantNavigables`{.variable} [in what order?]{.XXX}, [queue a
    global
    task](webappapis.html#queue-a-global-task){#aborting-a-document-load:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#aborting-a-document-load:navigation-and-traversal-task-source}
    given `descendantNavigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#aborting-a-document-load:nav-window}
    to perform the following steps:

    1.  [Abort](#abort-a-document){#aborting-a-document-load:abort-a-document}
        `descendantNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#aborting-a-document-load:nav-document}.

    2.  If `descendantNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#aborting-a-document-load:nav-document-2}\'s
        *[salvageable](#concept-document-salvageable)* is false, then
        set `document`{.variable}\'s
        *[salvageable](#concept-document-salvageable)* to false.

4.  [Abort](#abort-a-document){#aborting-a-document-load:abort-a-document-2}
    `document`{.variable}.

To [stop loading]{#nav-stop .dfn} a
[navigable](document-sequences.html#navigable){#aborting-a-document-load:navigable}
`navigable`{.variable}:

1.  Let `document`{.variable} be `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#aborting-a-document-load:nav-document-3}.

2.  If `document`{.variable}\'s [unload
    counter](#unload-counter){#aborting-a-document-load:unload-counter}
    is 0, and `navigable`{.variable}\'s [ongoing
    navigation](browsing-the-web.html#ongoing-navigation){#aborting-a-document-load:ongoing-navigation}
    is a [navigation
    ID](browsing-the-web.html#navigation-id){#aborting-a-document-load:navigation-id},
    then [set the ongoing
    navigation](browsing-the-web.html#set-the-ongoing-navigation){#aborting-a-document-load:set-the-ongoing-navigation}
    for `navigable`{.variable} to null.

    This will have the effect of aborting any ongoing navigations of
    `navigable`{.variable}, since at certain points during navigation,
    changes to the [ongoing
    navigation](browsing-the-web.html#ongoing-navigation){#aborting-a-document-load:ongoing-navigation-2}
    will cause further work to be abandoned.

3.  [Abort a document and its
    descendants](#abort-a-document-and-its-descendants){#aborting-a-document-load:abort-a-document-and-its-descendants}
    given `document`{.variable}.

Through their [user interface](#nav-traversal-ui), user agents also
allow stopping traversals, i.e. cases where the [ongoing
navigation](browsing-the-web.html#ongoing-navigation){#aborting-a-document-load:ongoing-navigation-3}
is \"`traversal`\". The above algorithm does not account for this. (On
the other hand, user agents do not allow
[`window.stop()`{#aborting-a-document-load:dom-window-stop}](nav-history-apis.html#dom-window-stop)
to stop traversals, so the above algorithm is correct for that caller.)
See [issue #6905](https://github.com/whatwg/html/issues/6905).

### [7.6]{.secno} The \`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options}](#x-frame-options)\` header[](#the-x-frame-options-header){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Headers/X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options "The X-Frame-Options HTTP response header can be used to indicate whether or not a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>. Sites can use this to avoid click-jacking attacks, by ensuring that their content is not embedded into other sites.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox AndroidYes]{.firefox_android .yes}[Safari iOSYes]{.safari_ios
.yes}[Chrome AndroidYes]{.chrome_android .yes}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The \`[`X-Frame-Options`]{#x-frame-options .dfn
dfn-type="http-header"}\` HTTP response header is a way of controlling
whether and how a
[`Document`{#the-x-frame-options-header:document}](dom.html#document)
may be loaded inside of a [child
navigable](document-sequences.html#child-navigable){#the-x-frame-options-header:child-navigable}.
For sites using CSP, the
[`frame-ancestors`{#the-x-frame-options-header:frame-ancestors-directive}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"}
directive provides more granular control over the same situations. It
was originally defined in HTTP Header Field X-Frame-Options, but the
definition and processing model here supersedes that document.
[\[CSP\]](references.html#refsCSP)
[\[RFC7034\]](references.html#refsRFC7034)

In particular, HTTP Header Field X-Frame-Options specified an
\``ALLOW-FROM`\` variant of the header, but that is not to be
implemented.

Per the below processing model, if both a CSP
[`frame-ancestors`{#the-x-frame-options-header:frame-ancestors-directive-2}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"}
directive and an
\`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-2}](#x-frame-options)\`
header are used in the same
[response](https://fetch.spec.whatwg.org/#concept-response){#the-x-frame-options-header:concept-response
x-internal="concept-response"}, then
\`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-3}](#x-frame-options)\`
is ignored.

For web developers and conformance checkers, its value
[ABNF](https://fetch.spec.whatwg.org/#abnf){#the-x-frame-options-header:header-abnf
x-internal="header-abnf"} is:

``` abnf
X-Frame-Options = "DENY" / "SAMEORIGIN"
```

To [check a navigation response\'s adherence to
\``X-Frame-Options`\`]{#check-a-navigation-response's-adherence-to-x-frame-options
.dfn}, given a
[response](https://fetch.spec.whatwg.org/#concept-response){#the-x-frame-options-header:concept-response-2
x-internal="concept-response"} `response`{.variable}, a
[navigable](document-sequences.html#navigable){#the-x-frame-options-header:navigable}
`navigable`{.variable}, a [CSP
list](https://w3c.github.io/webappsec-csp/#csp-list){#the-x-frame-options-header:concept-csp-list
x-internal="concept-csp-list"} `cspList`{.variable}, and an
[origin](browsers.html#concept-origin){#the-x-frame-options-header:concept-origin}
`destinationOrigin`{.variable}:

1.  If `navigable`{.variable} is not a [child
    navigable](document-sequences.html#child-navigable){#the-x-frame-options-header:child-navigable-2},
    then return true.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-x-frame-options-header:list-iterate
    x-internal="list-iterate"} `policy`{.variable} of
    `cspList`{.variable}:

    1.  If `policy`{.variable}\'s
        [disposition](https://w3c.github.io/webappsec-csp/#policy-disposition){#the-x-frame-options-header:csp-disposition
        x-internal="csp-disposition"} is not \"`enforce`\", then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-x-frame-options-header:continue
        x-internal="continue"}.

    2.  If `policy`{.variable}\'s [directive
        set](https://w3c.github.io/webappsec-csp/#policy-directive-set){#the-x-frame-options-header:csp-directive-set
        x-internal="csp-directive-set"}
        [contains](https://infra.spec.whatwg.org/#list-contain){#the-x-frame-options-header:list-contains
        x-internal="list-contains"} a
        [`frame-ancestors`{#the-x-frame-options-header:frame-ancestors-directive-3}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"}
        directive, then return true.

3.  Let `rawXFrameOptions`{.variable} be the result of [getting,
    decoding, and
    splitting](https://fetch.spec.whatwg.org/#concept-header-list-get-decode-split){#the-x-frame-options-header:concept-header-list-get-decode-split
    x-internal="concept-header-list-get-decode-split"}
    \`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-4}](#x-frame-options)\`
    from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-x-frame-options-header:concept-response-header-list
    x-internal="concept-response-header-list"}.

4.  Let `xFrameOptions`{.variable} be a new
    [set](https://infra.spec.whatwg.org/#ordered-set){#the-x-frame-options-header:set
    x-internal="set"}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-x-frame-options-header:list-iterate-2
    x-internal="list-iterate"} `value`{.variable} of
    `rawXFrameOptions`{.variable},
    [append](https://infra.spec.whatwg.org/#set-append){#the-x-frame-options-header:set-append
    x-internal="set-append"} `value`{.variable}, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#the-x-frame-options-header:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}, to
    `xFrameOptions`{.variable}.

6.  If `xFrameOptions`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#the-x-frame-options-header:list-size
    x-internal="list-size"} is greater than 1, and
    `xFrameOptions`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#the-x-frame-options-header:list-contains-2
    x-internal="list-contains"} any of \"`deny`\", \"`allowall`\", or
    \"`sameorigin`\", then return false.

    The intention here is to block any attempts at applying
    \`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-5}](#x-frame-options)\`
    which were trying to do something valid, but appear confused.

    This is the only impact of the legacy \``ALLOWALL`\` value on the
    processing model.

7.  If `xFrameOptions`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#the-x-frame-options-header:list-size-2
    x-internal="list-size"} is greater than 1, then return true.

    This means it contains multiple invalid values, which we treat the
    same way as if the header was omitted entirely.

8.  If `xFrameOptions`{.variable}\[0\] is \"`deny`\", then return false.

9.  If `xFrameOptions`{.variable}\[0\] is \"`sameorigin`\", then:

    1.  Let `containerDocument`{.variable} be `navigable`{.variable}\'s
        [container
        document](document-sequences.html#nav-container-document){#the-x-frame-options-header:nav-container-document}.

    2.  [While](https://infra.spec.whatwg.org/#iteration-while){#the-x-frame-options-header:while
        x-internal="while"} `containerDocument`{.variable} is not null:

        1.  If `containerDocument`{.variable}\'s
            [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-x-frame-options-header:concept-document-origin
            x-internal="concept-document-origin"} is not [same
            origin](browsers.html#same-origin){#the-x-frame-options-header:same-origin}
            with `destinationOrigin`{.variable}, then return false.

        2.  Set `containerDocument`{.variable} to
            `containerDocument`{.variable}\'s [container
            document](document-sequences.html#doc-container-document){#the-x-frame-options-header:doc-container-document}.

10. Return true.

    If we\'ve reached this point then we have a lone invalid value
    (which could potentially be one the legacy \``ALLOWALL`\` or
    \``ALLOW-FROM`\` forms). These are treated as if the header were
    omitted entirely.

------------------------------------------------------------------------

::: example
The following table illustrates the processing of various values for the
header, including non-conformant ones:

\`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-6}](#x-frame-options)\`

Valid

Result

\``DENY`\`

✅

embedding disallowed

\``SAMEORIGIN`\`

✅

same-origin embedding allowed

\``INVALID`\`

❌

embedding allowed

\``ALLOWALL`\`

❌

embedding allowed

\``ALLOW-FROM=https://example.com/`\`

❌

embedding allowed (from anywhere)
:::

::: example
The following table illustrates how various non-conformant cases
involving multiple values are processed:

\`[`X-Frame-Options`{#the-x-frame-options-header:x-frame-options-7}](#x-frame-options)\`

Result

\``SAMEORIGIN, SAMEORIGIN`\`

same-origin embedding allowed

\``SAMEORIGIN, DENY`\`

embedding disallowed

\``SAMEORIGIN,`\`

embedding disallowed

\``SAMEORIGIN, ALLOWALL`\`

embedding disallowed

\``SAMEORIGIN, INVALID`\`

embedding disallowed

\``ALLOWALL, INVALID`\`

embedding disallowed

\``ALLOWALL,`\`

embedding disallowed

\``INVALID, INVALID`\`

embedding allowed

The same results are obtained whether the values are delivered in a
single header whose value is comma-delimited, or in multiple headers.
:::

### [7.7]{.secno} The \`[`Refresh`{#the-refresh-header:refresh}](#refresh)\` header[](#the-refresh-header){.self-link}

The \`[`Refresh`]{#refresh .dfn dfn-type="http-header"}\` HTTP response
header is the HTTP-equivalent to a
[`meta`{#the-refresh-header:the-meta-element}](semantics.html#the-meta-element)
element with an
[`http-equiv`{#the-refresh-header:attr-meta-http-equiv}](semantics.html#attr-meta-http-equiv)
attribute in the [Refresh
state](semantics.html#attr-meta-http-equiv-refresh){#the-refresh-header:attr-meta-http-equiv-refresh}.
It takes [the same
value](semantics.html#conformance-attr-meta-http-equiv-refresh) and
works largely the same. Its processing model is detailed in [create and
initialize a `Document`
object](#initialise-the-document-object){#the-refresh-header:initialise-the-document-object}.

### [7.8]{.secno} []{#history-notes}Browser user interface considerations[](#nav-traversal-ui){.self-link} {#nav-traversal-ui}

Browser user agents should provide the ability to
[navigate](browsing-the-web.html#navigate){#nav-traversal-ui:navigate},
[reload](browsing-the-web.html#reload){#nav-traversal-ui:reload}, and
[stop loading](#nav-stop){#nav-traversal-ui:nav-stop} any [top-level
traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable}
in their [top-level traversable
set](document-sequences.html#top-level-traversable-set){#nav-traversal-ui:top-level-traversable-set}.

For example, via a location bar and reload/stop button UI.

Browser user agents should provide the ability to [traverse by a
delta](browsing-the-web.html#traverse-the-history-by-a-delta){#nav-traversal-ui:traverse-the-history-by-a-delta}
any [top-level
traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-2}
in their [top-level traversable
set](document-sequences.html#top-level-traversable-set){#nav-traversal-ui:top-level-traversable-set-2}.

For example, via back and forward buttons, possibly including long-press
abilities to change the delta.

It is suggested that such user agents allow traversal by deltas greater
than one, to avoid letting a page \"trap\" the user by stuffing the
session history with spurious entries. (For example, via repeated calls
to
[`history.pushState()`{#nav-traversal-ui:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate)
or [fragment
navigations](browsing-the-web.html#navigate-fragid){#nav-traversal-ui:navigate-fragid}.)

Some user agents have heuristics for translating a single \"back\" or
\"forward\" button press into a larger delta, specifically to overcome
such abuses. We are contemplating specifying these heuristics in [issue
#7832](https://github.com/whatwg/html/issues/7832).

Browser user agents should offer users the ability to [create a fresh
top-level
traversable](document-sequences.html#create-a-fresh-top-level-traversable){#nav-traversal-ui:create-a-fresh-top-level-traversable},
given a user-provided or user agent-determined initial
[URL](https://url.spec.whatwg.org/#concept-url){#nav-traversal-ui:url
x-internal="url"}.

For example, via a \"new tab\" or \"new window\" button.

Browser user agents should offer users the ability to arbitrarily
[close](document-sequences.html#close-a-top-level-traversable){#nav-traversal-ui:close-a-top-level-traversable}
any [top-level
traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-3}
in their [top-level traversable
set](document-sequences.html#top-level-traversable-set){#nav-traversal-ui:top-level-traversable-set-3}.

For example, by clicking a \"close tab\" button.

------------------------------------------------------------------------

Browser user agents may provide ways for the user to explicitly cause
any
[navigable](document-sequences.html#navigable){#nav-traversal-ui:navigable}
(not just a [top-level
traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-4})
to
[navigate](browsing-the-web.html#navigate){#nav-traversal-ui:navigate-2},
[reload](browsing-the-web.html#reload){#nav-traversal-ui:reload-2}, or
[stop loading](#nav-stop){#nav-traversal-ui:nav-stop-2}.

For example, via a context menu.

Browser user agents may provide the ability for users to [destroy a
top-level
traversable](document-sequences.html#destroy-a-top-level-traversable){#nav-traversal-ui:destroy-a-top-level-traversable}.

For example, by force-closing a window containing one or more such
[top-level
traversables](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-5}.

------------------------------------------------------------------------

When a user requests a
[reload](browsing-the-web.html#reload){#nav-traversal-ui:reload-3} of a
[navigable](document-sequences.html#navigable){#nav-traversal-ui:navigable-2}
whose [active session history
entry](document-sequences.html#nav-active-history-entry){#nav-traversal-ui:nav-active-history-entry}\'s
[document
state](browsing-the-web.html#she-document-state){#nav-traversal-ui:she-document-state}\'s
[resource](browsing-the-web.html#document-state-resource){#nav-traversal-ui:document-state-resource}
is a [POST
resource](browsing-the-web.html#post-resource){#nav-traversal-ui:post-resource},
the user agent should prompt the user to confirm the operation first,
since otherwise transactions (e.g., purchases or database modifications)
could be repeated.

When a user requests a
[reload](browsing-the-web.html#reload){#nav-traversal-ui:reload-4} of a
[navigable](document-sequences.html#navigable){#nav-traversal-ui:navigable-3},
user agents may provide a mechanism for ignoring any caches when
reloading.

------------------------------------------------------------------------

All calls to
[navigate](browsing-the-web.html#navigate){#nav-traversal-ui:navigate-3}
initiated by the mechanisms mentioned above must have the
*[userInvolvement](browsing-the-web.html#navigation-user-involvement)*
argument set to
\"[`browser UI`{#nav-traversal-ui:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\".

All calls to
[reload](browsing-the-web.html#reload){#nav-traversal-ui:reload-5}
initiated by the mechanisms mentioned above must have the
*[userInvolvement](browsing-the-web.html#reload-user-involvement)*
argument set to
\"[`browser UI`{#nav-traversal-ui:uni-browser-ui-2}](browsing-the-web.html#uni-browser-ui)\".

All calls to [traverse the history by a
delta](browsing-the-web.html#traverse-the-history-by-a-delta){#nav-traversal-ui:traverse-the-history-by-a-delta-2}
initiated by the mechanisms mentioned above must not pass a value for
the *[sourceDocument](browsing-the-web.html#traverse-sourcedocument)*
argument.

------------------------------------------------------------------------

The above recommendations, and the data structures in this
specification, are not meant to place restrictions on how user agents
represent the session history to the user.

For example, although a [top-level
traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-6}\'s
[session history
entries](document-sequences.html#tn-session-history-entries){#nav-traversal-ui:tn-session-history-entries}
are stored and maintained as a list, and the user agent is recommended
to give an interface for [traversing that list by a
delta](browsing-the-web.html#traverse-the-history-by-a-delta){#nav-traversal-ui:traverse-the-history-by-a-delta-3},
a novel user agent could instead or in addition present a tree-like
view, with each page having multiple \"forward\" pages that the user can
choose between.

Similarly, although session history for all descendant
[navigables](document-sequences.html#navigable){#nav-traversal-ui:navigable-4}
is stored in their [traversable
navigable](document-sequences.html#traversable-navigable){#nav-traversal-ui:traversable-navigable},
user agents could present the user with a more nuanced
per-[navigable](document-sequences.html#navigable){#nav-traversal-ui:navigable-5}
view of the session history.

------------------------------------------------------------------------

Browser user agents may use a [top-level browsing
context](document-sequences.html#top-level-browsing-context){#nav-traversal-ui:top-level-browsing-context}\'s
[is popup](document-sequences.html#is-popup){#nav-traversal-ui:is-popup}
boolean for the following purposes:

- Deciding whether or not to provide a minimal web browser user
  interface for the corresponding [top-level
  traversable](document-sequences.html#top-level-traversable){#nav-traversal-ui:top-level-traversable-7}.

- Performing the optional steps in [set up browsing context
  features](https://drafts.csswg.org/cssom-view/#set-up-browsing-context-features){#nav-traversal-ui:set-up-browsing-context-features
  x-internal="set-up-browsing-context-features"}.

In both cases user agents might additionally incorporate user
preferences, or present a choice as to whether to go down the popup
route.

User agents that provides a minimal user interface for such popups are
encouraged to not hide the browser\'s location bar.

[← 7.4 Navigation and session history](browsing-the-web.html) --- [Table
of Contents](./) --- [8 Web application APIs →](webappapis.html)
