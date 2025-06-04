HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 7.3 Infrastructure for sequences of
documents](document-sequences.html) --- [Table of Contents](./) --- [7.5
Document lifecycle →](document-lifecycle.html)

1.  ::: {#toc-browsers}
    1.  [[7.4]{.secno} Navigation and session
        history](browsing-the-web.html#navigation-and-session-history)
        1.  [[7.4.1]{.secno} Session
            history](browsing-the-web.html#session-history-infrastructure)
            1.  [[7.4.1.1]{.secno} Session history
                entries](browsing-the-web.html#session-history-entries)
            2.  [[7.4.1.2]{.secno} Document
                state](browsing-the-web.html#document-state)
            3.  [[7.4.1.3]{.secno} Centralized modifications of session
                history](browsing-the-web.html#centralized-modifications-of-session-history)
            4.  [[7.4.1.4]{.secno} Low-level operations on session
                history](browsing-the-web.html#low-level-operations-on-session-history)
        2.  [[7.4.2]{.secno}
            Navigation](browsing-the-web.html#navigating-across-documents)
            1.  [[7.4.2.1]{.secno} Supporting
                concepts](browsing-the-web.html#navigation-supporting-concepts)
            2.  [[7.4.2.2]{.secno} Beginning
                navigation](browsing-the-web.html#beginning-navigation)
            3.  [[7.4.2.3]{.secno} Ending
                navigation](browsing-the-web.html#ending-navigation)
                1.  [[7.4.2.3.1]{.secno} The usual cross-document
                    navigation
                    case](browsing-the-web.html#the-usual-cross-document-navigation-case)
                2.  [[7.4.2.3.2]{.secno} The `javascript:` URL special
                    case](browsing-the-web.html#the-javascript:-url-special-case)
                3.  [[7.4.2.3.3]{.secno} Fragment
                    navigations](browsing-the-web.html#scroll-to-fragid)
                4.  [[7.4.2.3.4]{.secno} Non-fetch schemes and external
                    software](browsing-the-web.html#non-fetch-schemes-and-external-software)
            4.  [[7.4.2.4]{.secno} Preventing
                navigation](browsing-the-web.html#preventing-navigation)
            5.  [[7.4.2.5]{.secno} Aborting
                navigation](browsing-the-web.html#aborting-navigation)
        3.  [[7.4.3]{.secno} Reloading and
            traversing](browsing-the-web.html#reloading-and-traversing)
        4.  [[7.4.4]{.secno} Non-fragment synchronous
            \"navigations\"](browsing-the-web.html#navigate-non-frag-sync)
        5.  [[7.4.5]{.secno} Populating a session history
            entry](browsing-the-web.html#populating-a-session-history-entry)
        6.  [[7.4.6]{.secno} Applying the history
            step](browsing-the-web.html#applying-the-history-step)
            1.  [[7.4.6.1]{.secno} Updating the
                traversable](browsing-the-web.html#updating-the-traversable)
            2.  [[7.4.6.2]{.secno} Updating the
                document](browsing-the-web.html#updating-the-document)
            3.  [[7.4.6.3]{.secno} Revealing the
                document](browsing-the-web.html#revealing-the-document)
            4.  [[7.4.6.4]{.secno} Scrolling to a
                fragment](browsing-the-web.html#scrolling-to-a-fragment)
            5.  [[7.4.6.5]{.secno} Persisted history entry
                state](browsing-the-web.html#persisted-user-state-restoration)
    :::

### [7.4]{.secno} []{#history}Navigation and session history[](#navigation-and-session-history){.self-link}

Welcome to the dragon\'s maw. Navigation, session history, and the
traversal through that session history are some of the most complex
parts of this standard.

The basic concept may not seem so difficult:

- The user is looking at a
  [navigable](document-sequences.html#navigable){#navigation-and-session-history:navigable}
  that is presenting its [active
  document](document-sequences.html#nav-document){#navigation-and-session-history:nav-document}.
  They [navigate](#navigate){#navigation-and-session-history:navigate}
  it to another
  [URL](https://url.spec.whatwg.org/#concept-url){#navigation-and-session-history:url
  x-internal="url"}.

- The browser fetches the given URL from the network, using it to
  [populate](#attempt-to-populate-the-history-entry's-document){#navigation-and-session-history:attempt-to-populate-the-history-entry's-document}
  a new [session history
  entry](#session-history-entry){#navigation-and-session-history:session-history-entry}
  with a
  newly-[created](document-lifecycle.html#initialise-the-document-object){#navigation-and-session-history:initialise-the-document-object}
  [`Document`{#navigation-and-session-history:document}](dom.html#document).

- The browser updates the
  [navigable](document-sequences.html#navigable){#navigation-and-session-history:navigable-2}\'s
  [active session history
  entry](document-sequences.html#nav-active-history-entry){#navigation-and-session-history:nav-active-history-entry}
  to the newly-populated one, and thus updates the [active
  document](document-sequences.html#nav-document){#navigation-and-session-history:nav-document-2}
  that it is showing to the user.

- At some point later, the user [presses the browser back
  button](#traverse-the-history-by-a-delta){#navigation-and-session-history:traverse-the-history-by-a-delta}
  to go back to the previous [session history
  entry](#session-history-entry){#navigation-and-session-history:session-history-entry-2}.

- The browser looks at the
  [URL](#she-url){#navigation-and-session-history:she-url} stored in
  that [session history
  entry](#session-history-entry){#navigation-and-session-history:session-history-entry-3},
  and uses it to re-fetch and
  [populate](#attempt-to-populate-the-history-entry's-document){#navigation-and-session-history:attempt-to-populate-the-history-entry's-document-2}
  that entry\'s
  [document](#she-document){#navigation-and-session-history:she-document}.

- The browser again updates the
  [navigable](document-sequences.html#navigable){#navigation-and-session-history:navigable-3}\'s
  [active session history
  entry](document-sequences.html#nav-active-history-entry){#navigation-and-session-history:nav-active-history-entry-2}.

You can see some of the intertwined complexity peeking through here, in
how traversal can cause a navigation (i.e., a network fetch to a stored
URL), and how a navigation necessarily needs to interface with the
session history list to ensure that when it finishes the user is looking
at the right thing. But the real problems come in with the various edge
cases and interacting web platform features:

- [Child
  navigables](document-sequences.html#child-navigable){#navigation-and-session-history:child-navigable}
  (e.g., those contained in
  [`iframe`{#navigation-and-session-history:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s)
  can also navigate and traverse, but those navigations need to be
  linearized into [a single session history
  list](document-sequences.html#tn-session-history-entries){#navigation-and-session-history:tn-session-history-entries}
  since the user only has a single back/forward interface for the entire
  [traversable
  navigable](document-sequences.html#traversable-navigable){#navigation-and-session-history:traversable-navigable}
  (e.g., browser tab).

- Since the user can traverse back more than a single step in the
  session history (e.g., by holding down their back button), they can
  end up traversing multiple
  [navigables](document-sequences.html#navigable){#navigation-and-session-history:navigable-4}
  at the same time when [child
  navigables](document-sequences.html#child-navigable){#navigation-and-session-history:child-navigable-2}
  are involved. This needs to be synchronized across all of the involved
  navigables, which might involve multiple [event
  loops](webappapis.html#event-loop){#navigation-and-session-history:event-loop}
  or even [agent
  clusters](https://tc39.es/ecma262/#sec-agent-clusters){#navigation-and-session-history:agent-cluster
  x-internal="agent-cluster"}.

- During navigation, servers can respond with 204 or 205 status codes or
  with
  \`[`Content-Disposition: attachment`{#navigation-and-session-history:http-content-disposition}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
  headers, which cause navigation to abort and the
  [navigable](document-sequences.html#navigable){#navigation-and-session-history:navigable-5}
  to stay on its original [active
  document](document-sequences.html#active-document){#navigation-and-session-history:active-document}.
  (This is much worse if it happens during a traversal-initiated
  navigation!)

- Various other HTTP headers, such as \``Location`\`,
  \`[`Refresh`{#navigation-and-session-history:refresh}](document-lifecycle.html#refresh)\`,
  \`[`X-Frame-Options`{#navigation-and-session-history:x-frame-options}](document-lifecycle.html#x-frame-options)\`,
  and those for Content Security Policy, contribute to either the
  [fetching
  process](#create-navigation-params-by-fetching){#navigation-and-session-history:create-navigation-params-by-fetching},
  or the [`Document`-creation
  process](document-lifecycle.html#initialise-the-document-object){#navigation-and-session-history:initialise-the-document-object-2},
  or both. The
  \`[`Cross-Origin-Opener-Policy`{#navigation-and-session-history:cross-origin-opener-policy-2}](browsers.html#cross-origin-opener-policy-2)\`
  header even contributes to the [browsing context selection and
  creation](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy)
  process!

- Some navigations (namely [fragment navigations](#scroll-to-fragid) and
  [single-page app navigations](#navigate-non-frag-sync)) are
  synchronous, meaning that JavaScript code expects to observe the
  navigation\'s results instantly. This then needs to be synchronized
  with the view of the session history that all other
  [navigables](document-sequences.html#navigable){#navigation-and-session-history:navigable-6}
  in the tree see, which can be subject to race conditions and
  necessitate resolving conflicting views of the session history.

- The platform has accumulated various exciting navigation-related
  features that need special-casing, such as
  [`javascript:`{#navigation-and-session-history:the-javascript:-url-special-case}](#the-javascript:-url-special-case)
  URLs,
  [`srcdoc`{#navigation-and-session-history:attr-iframe-srcdoc}](iframe-embed-object.html#attr-iframe-srcdoc)
  [`iframe`{#navigation-and-session-history:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)s,
  and the
  [`beforeunload`{#navigation-and-session-history:event-beforeunload}](indices.html#event-beforeunload)
  event.

In what follows, we have attempted to guide the reader through these
complexities by appropriately cordoning them off into labeled sections
and algorithms, and giving appropriate words of introduction where
possible. Nevertheless, if you wish to truly understand navigation and
session history, [the usual
advice](introduction.html#how-to-read-this-specification) will be
invaluable.

#### [7.4.1]{.secno} []{#the-session-history-of-browsing-contexts}Session history[](#session-history-infrastructure){.self-link} {#session-history-infrastructure}

##### [7.4.1.1]{.secno} Session history entries[](#session-history-entries){.self-link}

A [session history entry]{#session-history-entry .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#session-history-entries:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#session-history-entries:struct-item
x-internal="struct-item"}:

- [step]{#she-step .dfn}, a non-negative integer or \"`pending`\",
  initially \"`pending`\".

- [URL]{#she-url .dfn}, a
  [URL](https://url.spec.whatwg.org/#concept-url){#session-history-entries:url
  x-internal="url"}

- [document state]{#she-document-state .dfn}, a [document
  state](#document-state-2){#session-history-entries:document-state-2}.

- [classic history API state]{#she-classic-history-api-state .dfn},
  which is [serialized
  state](#serialized-state){#session-history-entries:serialized-state},
  initially
  [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#session-history-entries:structuredserializeforstorage}(null).

- [navigation API state]{#she-navigation-api-state .dfn}, which is a
  [serialized
  state](#serialized-state){#session-history-entries:serialized-state-2},
  initially
  [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#session-history-entries:structuredserializeforstorage-2}(undefined).

- [navigation API key]{#she-navigation-api-key .dfn}, which is a string,
  initially set to the result of [generating a random
  UUID](https://w3c.github.io/webcrypto/#dfn-generate-a-random-uuid){#session-history-entries:generating-a-random-uuid
  x-internal="generating-a-random-uuid"}.

- [navigation API ID]{#she-navigation-api-id .dfn}, which is a string,
  initially set to the result of [generating a random
  UUID](https://w3c.github.io/webcrypto/#dfn-generate-a-random-uuid){#session-history-entries:generating-a-random-uuid-2
  x-internal="generating-a-random-uuid"}.

- [scroll restoration mode]{#she-scroll-restoration-mode .dfn}, a
  [scroll restoration
  mode](#scroll-restoration-mode){#session-history-entries:scroll-restoration-mode},
  initially
  \"[`auto`{#session-history-entries:dom-scrollrestoration-auto}](#dom-scrollrestoration-auto)\".

- [scroll position data]{#she-scroll-position .dfn}, which is scroll
  position data for the
  [document](#she-document){#session-history-entries:she-document}\'s
  [restorable scrollable
  regions](#restorable-scrollable-regions){#session-history-entries:restorable-scrollable-regions}.

- [persisted user state]{#she-other .dfn}, which is
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#session-history-entries:implementation-defined
  x-internal="implementation-defined"}, initially null

  For example, some user agents might want to persist the values of form
  controls.

  User agents that persist the value of form controls are encouraged to
  also persist their directionality (the value of the element\'s
  [`dir`{#session-history-entries:attr-dir}](dom.html#attr-dir)
  attribute). This prevents values from being displayed incorrectly
  after a history traversal when the user had originally entered the
  values with an explicit, non-default directionality.

To get a [session history
entry](#session-history-entry){#session-history-entries:session-history-entry}\'s
[document]{#she-document .dfn}, return its [document
state](#she-document-state){#session-history-entries:she-document-state}\'s
[document](#document-state-document){#session-history-entries:document-state-document}.

------------------------------------------------------------------------

[Serialized state]{#serialized-state .dfn} is a serialization (via
[StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#session-history-entries:structuredserializeforstorage-3})
of an object representing a user interface state. We sometimes
informally refer to \"state objects\", which are the objects
representing user interface state supplied by the author, or alternately
the objects created by deserializing (via
[StructuredDeserialize](structured-data.html#structureddeserialize){#session-history-entries:structureddeserialize})
serialized state.

Pages can
[add](nav-history-apis.html#dom-history-pushstate){#session-history-entries:dom-history-pushstate}
[serialized
state](#serialized-state){#session-history-entries:serialized-state-3}
to the session history. These are then
[deserialized](structured-data.html#structureddeserialize){#session-history-entries:structureddeserialize-2}
and [returned to the
script](indices.html#event-popstate){#session-history-entries:event-popstate}
when the user (or script) goes back in the history, thus enabling
authors to use the \"navigation\" metaphor even in one-page
applications.

::: note
[Serialized
state](#serialized-state){#session-history-entries:serialized-state-4}
is intended to be used for two main purposes: first, storing a preparsed
description of the state in the
[URL](https://url.spec.whatwg.org/#concept-url){#session-history-entries:url-2
x-internal="url"} so that in the simple case an author doesn\'t have to
do the parsing (though one would still need the parsing for handling
[URLs](https://url.spec.whatwg.org/#concept-url){#session-history-entries:url-3
x-internal="url"} passed around by users, so it\'s only a minor
optimization). Second, so that the author can store state that one
wouldn\'t store in the URL because it only applies to the current
[`Document`{#session-history-entries:document}](dom.html#document)
instance and it would have to be reconstructed if a new
[`Document`{#session-history-entries:document-2}](dom.html#document)
were opened.

An example of the latter would be something like keeping track of the
precise coordinate from which a popup
[`div`{#session-history-entries:the-div-element}](grouping-content.html#the-div-element)
was made to animate, so that if the user goes back, it can be made to
animate to the same location. Or alternatively, it could be used to keep
a pointer into a cache of data that would be fetched from the server
based on the information in the
[URL](https://url.spec.whatwg.org/#concept-url){#session-history-entries:url-4
x-internal="url"}, so that when going back and forward, the information
doesn\'t have to be fetched again.
:::

------------------------------------------------------------------------

A [scroll restoration mode]{#scroll-restoration-mode .dfn} indicates
whether the user agent should restore the persisted scroll position (if
any) when traversing to an
[entry](#session-history-entry){#session-history-entries:session-history-entry-2}.
A scroll restoration mode is one of the following:

\"[`auto`]{#dom-scrollrestoration-auto .dfn dfn-for="ScrollRestoration" dfn-type="enum-value"}\"
:   The user agent is responsible for restoring the scroll position upon
    navigation.

\"[`manual`]{#dom-scrollrestoration-manual .dfn dfn-for="ScrollRestoration" dfn-type="enum-value"}\"
:   The page is responsible for restoring the scroll position and the
    user agent does not attempt to do so automatically

##### [7.4.1.2]{.secno} Document state[](#document-state){.self-link}

[Document state]{#document-state-2 .dfn} holds state inside a [session
history
entry](#session-history-entry){#document-state:session-history-entry}
regarding how to present and, if necessary, recreate, a
[`Document`{#document-state:document}](dom.html#document). It has:

- A [document]{#document-state-document .dfn}, a
  [`Document`{#document-state:document-2}](dom.html#document) or null,
  initially null.

  ::: {#note-bfcache .note}
  When a history entry is
  [active](document-sequences.html#nav-active-history-entry){#document-state:nav-active-history-entry},
  it has a [`Document`{#document-state:document-3}](dom.html#document)
  in its [document
  state](#she-document-state){#document-state:she-document-state}.
  However, when a
  [`Document`{#document-state:document-4}](dom.html#document) is not
  [fully
  active](document-sequences.html#fully-active){#document-state:fully-active},
  it\'s possible for it to be
  [destroyed](document-lifecycle.html#destroy-a-document){#document-state:destroy-a-document}
  to free resources. In such cases, this
  [document](#document-state-document){#document-state:document-state-document}
  item will be nulled out. The [URL](#she-url){#document-state:she-url}
  and other data in the [session history
  entry](#session-history-entry){#document-state:session-history-entry-2}
  and [document
  state](#she-document-state){#document-state:she-document-state-2} is
  then used to bring a new
  [`Document`{#document-state:document-5}](dom.html#document) into being
  to take the place of the original, in the case where the user agent
  finds itself having to traverse to the entry.

  If the [`Document`{#document-state:document-6}](dom.html#document) is
  *not*
  [destroyed](document-lifecycle.html#destroy-a-document){#document-state:destroy-a-document-2},
  then during [history
  traversal](#traverse-the-history-by-a-delta){#document-state:traverse-the-history-by-a-delta},
  it can be
  [reactivated](#reactivate-a-document){#document-state:reactivate-a-document}.
  The cache in which browsers store such
  [`Document`{#document-state:document-7}](dom.html#document)s is often
  called a *back-forward cache*, or *bfcache* (or perhaps [\"blazingly
  fast\"](https://bugzilla.mozilla.org/show_bug.cgi?id=274784) cache).
  :::

- A [history policy container]{#document-state-history-policy-container
  .dfn}, a [policy
  container](browsers.html#policy-container){#document-state:policy-container}
  or null, initially null.

- A [request referrer]{#document-state-request-referrer .dfn}, which is
  \"`no-referrer`\", \"`client`\", or a
  [URL](https://url.spec.whatwg.org/#concept-url){#document-state:url
  x-internal="url"}, initially \"`client`\".

- A [request referrer policy]{#document-state-request-referrer-policy
  .dfn}, which is a [referrer
  policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#document-state:referrer-policy
  x-internal="referrer-policy"}, initially the [default referrer
  policy](https://w3c.github.io/webappsec-referrer-policy/#default-referrer-policy){#document-state:default-referrer-policy
  x-internal="default-referrer-policy"}.

  The [request referrer
  policy](#document-state-request-referrer-policy){#document-state:document-state-request-referrer-policy}
  is distinct from the [history policy
  container](#document-state-history-policy-container){#document-state:document-state-history-policy-container}\'s
  [referrer
  policy](browsers.html#policy-container-referrer-policy){#document-state:policy-container-referrer-policy}.
  The former is used for fetches *of* this document, whereas the latter
  controls fetches *by* this document.

- An [initiator origin]{#document-state-initiator-origin .dfn}, which is
  an
  [origin](browsers.html#concept-origin){#document-state:concept-origin}
  or null, initially null.

- An [origin]{#document-state-origin .dfn}, which is an
  [origin](browsers.html#concept-origin){#document-state:concept-origin-2}
  or null, initially null.

  This is the origin that we set \"`about:`\"-schemed
  [`Document`{#document-state:document-8}](dom.html#document)s\'
  [origin](https://dom.spec.whatwg.org/#concept-document-origin){#document-state:concept-document-origin
  x-internal="concept-document-origin"} to. We store it here because it
  is also used when restoring these
  [`Document`{#document-state:document-9}](dom.html#document)s during
  traversal, since they are reconstructed locally without visiting the
  network. It is also used to compare the origin before and after the
  [session history
  entry](#session-history-entry){#document-state:session-history-entry-3}
  is
  [repopulated](#attempt-to-populate-the-history-entry's-document){#document-state:attempt-to-populate-the-history-entry's-document}.
  If the origins change, the [navigable target
  name](#document-state-nav-target-name){#document-state:document-state-nav-target-name}
  is cleared.

- An [about base URL]{#document-state-about-base-url .dfn}, which is a
  [URL](https://url.spec.whatwg.org/#concept-url){#document-state:url-2
  x-internal="url"} or null, initially null.

  This will be populated only for \"`about:`\"-schemed
  [`Document`{#document-state:document-10}](dom.html#document)s and will
  be the [fallback base
  URL](urls-and-fetching.html#fallback-base-url){#document-state:fallback-base-url}
  for those
  [`Document`{#document-state:document-11}](dom.html#document)s. It is a
  snapshot of the initiator
  [`Document`{#document-state:document-12}](dom.html#document)\'s
  [document base
  URL](urls-and-fetching.html#document-base-url){#document-state:document-base-url}.

- [Nested histories]{#document-state-nested-histories .dfn}, a
  [list](https://infra.spec.whatwg.org/#list){#document-state:list
  x-internal="list"} of [nested
  histories](#nested-history){#document-state:nested-history}, initially
  an empty
  [list](https://infra.spec.whatwg.org/#list){#document-state:list-2
  x-internal="list"}.

- A [resource]{#document-state-resource .dfn}, a string, [POST
  resource](#post-resource){#document-state:post-resource} or null,
  initially null.

  A string is treated as HTML. It\'s used to store the source of an
  [`iframe` `srcdoc`
  document](iframe-embed-object.html#an-iframe-srcdoc-document){#document-state:an-iframe-srcdoc-document}.

- A [reload pending]{#document-state-reload-pending .dfn} boolean,
  initially false.

- An [ever populated]{#document-state-ever-populated .dfn} boolean,
  initially false.

- A [navigable target name]{#document-state-nav-target-name .dfn}
  string, initially the empty string.

- A [not restored reasons]{#document-state-not-restored-reasons .dfn}, a
  [not restored
  reasons](nav-history-apis.html#nrr-struct){#document-state:nrr-struct}
  or null, initially null.

User agents may [destroy a document and its
descendants](document-lifecycle.html#destroy-a-document-and-its-descendants){#document-state:destroy-a-document-and-its-descendants}
given the
[documents](#document-state-document){#document-state:document-state-document-2}
of [document
states](#document-state-2){#document-state:document-state-2} with
non-null
[documents](#document-state-document){#document-state:document-state-document-3},
as long as the
[`Document`{#document-state:document-13}](dom.html#document) is not
[fully
active](document-sequences.html#fully-active){#document-state:fully-active-2}.

Apart from that restriction, this standard does not specify when user
agents should destroy the
[document](#document-state-document){#document-state:document-state-document-4}
stored in a [document
state](#document-state-2){#document-state:document-state-2-2}, versus
keeping it cached.

------------------------------------------------------------------------

A [POST resource]{#post-resource .dfn export=""} has:

- A [request body]{#post-resource-request-body .dfn
  dfn-for="POST resource" export=""}, a [byte
  sequence](https://infra.spec.whatwg.org/#byte-sequence){#document-state:byte-sequence
  x-internal="byte-sequence"} or failure.

  This is only ever accessed [in
  parallel](infrastructure.html#in-parallel){#document-state:in-parallel},
  so it doesn\'t need to be stored in memory. However, it must return
  the same [byte
  sequence](https://infra.spec.whatwg.org/#byte-sequence){#document-state:byte-sequence-2
  x-internal="byte-sequence"} each time. If this isn\'t possible due to
  resources changing on disk, or if resources can no longer be accessed,
  then this must be set to failure.

- A [request content-type]{#post-resource-request-content-type .dfn
  dfn-for="POST resource" export=""}, which is
  \`[`application/x-www-form-urlencoded`{#document-state:application/x-www-form-urlencoded}](https://url.spec.whatwg.org/#concept-urlencoded){x-internal="application/x-www-form-urlencoded"}\`,
  \`[`multipart/form-data`{#document-state:multipart/form-data}](indices.html#multipart/form-data)\`,
  or
  \`[`text/plain`{#document-state:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\`.

------------------------------------------------------------------------

A [nested history]{#nested-history .dfn} has:

- An [id]{#nested-history-id .dfn}, a [unique internal
  value](common-microsyntaxes.html#unique-internal-value){#document-state:unique-internal-value}.

  This is used to associate the [nested
  history](#nested-history){#document-state:nested-history-2} with a
  [navigable](document-sequences.html#navigable){#document-state:navigable}.

- [Entries]{#nested-history-entries .dfn}, a
  [list](https://infra.spec.whatwg.org/#list){#document-state:list-3
  x-internal="list"} of [session history
  entries](#session-history-entry){#document-state:session-history-entry-4}.

This will later contain ways to identify a child navigable across
reloads.

------------------------------------------------------------------------

------------------------------------------------------------------------

Several contiguous entries in a session history can share the same
[document
state](#she-document-state){#document-state:she-document-state-3}. This
can occur when the initial entry is reached via normal
[navigation](#navigate){#document-state:navigate}, and the following
entry is added via
[`history.pushState()`{#document-state:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate).
Or it can occur via [navigation to a
fragment](#navigate-fragid){#document-state:navigate-fragid}.

All entries that share the same [document
state](#she-document-state){#document-state:she-document-state-4} (and
that are therefore merely different states of one particular document)
are contiguous by construction.

------------------------------------------------------------------------

A [`Document`{#document-state:document-14}](dom.html#document) has a
[latest entry]{#latest-entry .dfn}, a [session history
entry](#session-history-entry){#document-state:session-history-entry-5}
or null.

This is the entry that was most recently represented by a given
[`Document`{#document-state:document-15}](dom.html#document). A single
[`Document`{#document-state:document-16}](dom.html#document) can
represent many [session history
entries](#session-history-entry){#document-state:session-history-entry-6}
over time, as many contiguous [session history
entries](#session-history-entry){#document-state:session-history-entry-7}
can share the same [document
state](#she-document-state){#document-state:she-document-state-5} as
explained above.

##### [7.4.1.3]{.secno} Centralized modifications of session history[](#centralized-modifications-of-session-history){.self-link}

To maintain a single source of truth, all modifications to a
[traversable
navigable](document-sequences.html#traversable-navigable){#centralized-modifications-of-session-history:traversable-navigable}\'s
[session history
entries](document-sequences.html#tn-session-history-entries){#centralized-modifications-of-session-history:tn-session-history-entries}
need to be synchronized. This is especially important due to how session
history is influenced by all of the descendant
[navigables](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable},
and thus by multiple [event
loops](webappapis.html#event-loop){#centralized-modifications-of-session-history:event-loop}.
To accomplish this, we use the [session history traversal parallel
queue](#session-history-traversal-parallel-queue){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue}
structure.

A [session history traversal parallel
queue]{#session-history-traversal-parallel-queue .dfn} is very similar
to a [parallel
queue](infrastructure.html#parallel-queue){#centralized-modifications-of-session-history:parallel-queue}.
It has an [algorithm
set]{#session-history-traversal-parallel-queue-algorithm-set .dfn}, an
[ordered
set](https://infra.spec.whatwg.org/#ordered-set){#centralized-modifications-of-session-history:set
x-internal="set"}.

The
[items](https://infra.spec.whatwg.org/#list-item){#centralized-modifications-of-session-history:list-item
x-internal="list-item"} in a [session history traversal parallel
queue](#session-history-traversal-parallel-queue){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-2}\'s
[algorithm
set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set}
are either algorithm steps, or [synchronous navigation
steps]{#session-history-traversal-parallel-queue-sync-nav-steps .dfn},
which are a particular brand of algorithm steps involving a [target
navigable]{#session-history-traversal-parallel-queue-sync-nav-steps-target-nav
.dfn} (a
[navigable](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable-2}).

To [append session history traversal
steps]{#tn-append-session-history-traversal-steps .dfn} to a
[traversable
navigable](document-sequences.html#traversable-navigable){#centralized-modifications-of-session-history:traversable-navigable-2}
`traversable`{.variable} given algorithm steps `steps`{.variable},
[append](https://infra.spec.whatwg.org/#list-append){#centralized-modifications-of-session-history:list-append
x-internal="list-append"} `steps`{.variable} to
`traversable`{.variable}\'s [session history traversal
queue](document-sequences.html#tn-session-history-traversal-queue){#centralized-modifications-of-session-history:tn-session-history-traversal-queue}\'s
[algorithm
set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set-2}.

To [append session history synchronous navigation
steps]{#tn-append-session-history-sync-nav-steps .dfn} involving a
[navigable](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable-3}
`targetNavigable`{.variable} to a [traversable
navigable](document-sequences.html#traversable-navigable){#centralized-modifications-of-session-history:traversable-navigable-3}
`traversable`{.variable} given algorithm steps `steps`{.variable},
[append](https://infra.spec.whatwg.org/#list-append){#centralized-modifications-of-session-history:list-append-2
x-internal="list-append"} `steps`{.variable} as [synchronous navigation
steps](#session-history-traversal-parallel-queue-sync-nav-steps){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-sync-nav-steps}
targeting [target
navigable](#session-history-traversal-parallel-queue-sync-nav-steps-target-nav){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-sync-nav-steps-target-nav}
`targetNavigable`{.variable} to `traversable`{.variable}\'s [session
history traversal
queue](document-sequences.html#tn-session-history-traversal-queue){#centralized-modifications-of-session-history:tn-session-history-traversal-queue-2}\'s
[algorithm
set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set-3}.

To [start a new session history traversal parallel
queue]{#starting-a-new-session-history-traversal-parallel-queue .dfn}:

1.  Let `sessionHistoryTraversalQueue`{.variable} be a new [session
    history traversal parallel
    queue](#session-history-traversal-parallel-queue){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-3}.

2.  Run the following steps [in
    parallel](infrastructure.html#in-parallel){#centralized-modifications-of-session-history:in-parallel}:

    1.  While true:

        1.  If `sessionHistoryTraversalQueue`{.variable}\'s [algorithm
            set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set-4}
            is empty, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#centralized-modifications-of-session-history:continue
            x-internal="continue"}.

        2.  Let `steps`{.variable} be the result of
            [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#centralized-modifications-of-session-history:dequeue
            x-internal="dequeue"} from
            `sessionHistoryTraversalQueue`{.variable}\'s [algorithm
            set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set-5}.

        3.  Run `steps`{.variable}.

3.  Return `sessionHistoryTraversalQueue`{.variable}.

[Synchronous navigation
steps](#session-history-traversal-parallel-queue-sync-nav-steps){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-sync-nav-steps-2}
are tagged in the [algorithm
set](#session-history-traversal-parallel-queue-algorithm-set){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-algorithm-set-6}
to allow them to conditionally \"jump the queue\". This is handled
[within apply the history step](#sync-navigations-jump-queue).

::: {#example-sync-navigation-steps-queue-jumping-basic .example}
Imagine the joint session history depicted by this [Jake
diagram](document-sequences.html#jake-diagram){#centralized-modifications-of-session-history:jake-diagram}:

0

1

`top`

/a

/b

And the following code runs at the top level:

``` js
history.back();
location.href = '#foo';
```

The desired result is:

0

1

2

`top`

/a

/b

/b#foo

This isn\'t straightforward, as the sync navigation wins the race in
terms of being observable, whereas the traversal wins the race in terms
of queuing steps on the [session history traversal parallel
queue](#session-history-traversal-parallel-queue){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-4}.
To achieve this result, the following happens:

1.  [`history.back()`{#centralized-modifications-of-session-history:dom-history-back}](nav-history-apis.html#dom-history-back)
    [appends
    steps](#tn-append-session-history-traversal-steps){#centralized-modifications-of-session-history:tn-append-session-history-traversal-steps}
    intended to traverse by a delta of −1.

2.  [`location.href = '#foo'`{#centralized-modifications-of-session-history:dom-location-href}](nav-history-apis.html#dom-location-href)
    synchronously changes the [active session history
    entry](document-sequences.html#nav-active-history-entry){#centralized-modifications-of-session-history:nav-active-history-entry}
    entry to a newly-created one, with the URL `/b#foo`, and [appends
    synchronous
    steps](#tn-append-session-history-sync-nav-steps){#centralized-modifications-of-session-history:tn-append-session-history-sync-nav-steps}
    to notify the central source of truth about that new entry. Note
    that this does *not* yet update the [current session history
    entry](document-sequences.html#nav-current-history-entry){#centralized-modifications-of-session-history:nav-current-history-entry},
    [current session history
    step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step},
    or the [session history
    entries](document-sequences.html#tn-session-history-entries){#centralized-modifications-of-session-history:tn-session-history-entries-2}
    list; those updates cannot be done synchronously, and instead must
    be done as part of the queued steps.

3.  On the [session history traversal parallel
    queue](#session-history-traversal-parallel-queue){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-5},
    the steps queued by
    [`history.back()`{#centralized-modifications-of-session-history:dom-history-back-2}](nav-history-apis.html#dom-history-back)
    run:

    1.  The target history step is determined to be 0: the [current
        session history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-2}
        (i.e., 1) plus the intended delta of −1.

    2.  We enter the main [apply the history
        step](#apply-the-history-step){#centralized-modifications-of-session-history:apply-the-history-step}
        algorithm.

        The entry at step 0, for the `/a` URL, has its
        [document](#she-document){#centralized-modifications-of-session-history:she-document}
        [populated](#attempt-to-populate-the-history-entry's-document){#centralized-modifications-of-session-history:attempt-to-populate-the-history-entry's-document}.

        Meanwhile, the queue is checked for [synchronous navigation
        steps](#session-history-traversal-parallel-queue-sync-nav-steps){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-sync-nav-steps-3}.
        The steps queued by the
        [`location.href`{#centralized-modifications-of-session-history:dom-location-href-2}](nav-history-apis.html#dom-location-href)
        setter now run, and block the traversal from performing effects
        beyond document population (such as, unloading documents and
        switching active history entries) until they are finished. Those
        steps cause the following to happen:

        1.  The entry with URL `/b#foo` is added, with its
            [step](#she-step){#centralized-modifications-of-session-history:she-step}
            determined to be 2: the [current session history
            step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-3}
            (i.e., 1) plus 1.

        2.  We fully switch to that newly added entry, including a
            nested call to [apply the history
            step](#apply-the-history-step){#centralized-modifications-of-session-history:apply-the-history-step-2}.
            This ultimately results in [updating the
            document](#update-document-for-history-step-application){#centralized-modifications-of-session-history:update-document-for-history-step-application}
            by dispatching events like
            [`hashchange`{#centralized-modifications-of-session-history:event-hashchange}](indices.html#event-hashchange).

        Only once that is all complete, and the `/a` history entry has
        been fully populated with a
        [document](#she-document){#centralized-modifications-of-session-history:she-document-2},
        do we move on with applying the history step given the target
        step of 0.

        At this point, the
        [`Document`{#centralized-modifications-of-session-history:document}](dom.html#document)
        with URL `/b#foo`
        [unloads](document-lifecycle.html#unload-a-document){#centralized-modifications-of-session-history:unload-a-document},
        and we finish moving to our target history step 0, which makes
        the entry with URL `/a` become the [active session history
        entry](document-sequences.html#nav-active-history-entry){#centralized-modifications-of-session-history:nav-active-history-entry-2}
        and 0 become the [current session history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-4}.
:::

::: {#example-sync-navigation-steps-queue-jumping-complex .example}
Here is another more complex example, involving races between populating
two different
[`iframe`{#centralized-modifications-of-session-history:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s,
and a synchronous navigation once one of those iframes loads. We start
with this setup:

0

1

2

`top`

/t

`frames[0]`

/i-0-a

/i-0-b

`frames[1]`

/i-1-a

/i-1-b

and then call
[`history.go(-2)`{#centralized-modifications-of-session-history:dom-history-go}](nav-history-apis.html#dom-history-go).
The following then occurs:

1.  [`history.go(-2)`{#centralized-modifications-of-session-history:dom-history-go-2}](nav-history-apis.html#dom-history-go)
    [appends
    steps](#tn-append-session-history-traversal-steps){#centralized-modifications-of-session-history:tn-append-session-history-traversal-steps-2}
    intended to traverse by a delta of −2. Once those steps run:

    1.  The target step is determined to be 2 + (−2) = 0.

    2.  In parallel, the fetches are made to
        [populate](#attempt-to-populate-the-history-entry's-document){#centralized-modifications-of-session-history:attempt-to-populate-the-history-entry's-document-2}
        the two iframes, fetching `/i-0-a` and `/i-1-a` respectively.

        Meanwhile, the queue is checked for [synchronous navigation
        steps](#session-history-traversal-parallel-queue-sync-nav-steps){#centralized-modifications-of-session-history:session-history-traversal-parallel-queue-sync-nav-steps-4}.
        There aren\'t any right now.

    3.  In the fetch race, the fetch for `/i-0-a` wins. We proceed
        onward to finish all of [apply the history
        step](#apply-the-history-step){#centralized-modifications-of-session-history:apply-the-history-step-3}\'s
        work for how the traversal impacts the `frames[0]`
        [navigable](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable-4},
        including updating its [active session history
        entry](document-sequences.html#nav-active-history-entry){#centralized-modifications-of-session-history:nav-active-history-entry-3}
        to the entry with URL `/i-0-a`.

    4.  Before the fetch for `/i-1-a` finishes, we reach the point where
        [scripts may run for the newly-created
        document](#scripts-may-run-for-the-newly-created-document){#centralized-modifications-of-session-history:scripts-may-run-for-the-newly-created-document}
        in the `frames[0]`
        [navigable](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable-5}\'s
        [active
        document](document-sequences.html#nav-document){#centralized-modifications-of-session-history:nav-document}.
        Some such script does run:

        ``` js
        location.href = '#foo'
        ```

        This synchronously changes the `frames[0]` navigable\'s [active
        session history
        entry](document-sequences.html#nav-active-history-entry){#centralized-modifications-of-session-history:nav-active-history-entry-4}
        entry to a newly-created one, with the URL `/i-0-a#foo`, and
        [appends synchronous
        steps](#tn-append-session-history-sync-nav-steps){#centralized-modifications-of-session-history:tn-append-session-history-sync-nav-steps-2}
        to notify the central source of truth about that new entry.

        Unlike in the [previous
        example](#example-sync-navigation-steps-queue-jumping-basic),
        these synchronous steps do *not* \"jump the queue\" and update
        the
        [traversable](document-sequences.html#traversable-navigable){#centralized-modifications-of-session-history:traversable-navigable-4}
        before we finish the fetch for `/i-1-a`. This is because the
        navigable in question, `frames[0]`, has already been altered as
        part of the traversal, so we know that with the [current session
        history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-5}
        being 2, adding the new entry as a step 3 doesn\'t make sense.

    5.  Once the fetch for `/i-1-a` finally finishes, we proceed to
        finish updating the `frames[1]`
        [navigable](document-sequences.html#navigable){#centralized-modifications-of-session-history:navigable-6}
        for the traversal, including updating its [active session
        history
        entry](document-sequences.html#nav-active-history-entry){#centralized-modifications-of-session-history:nav-active-history-entry-5}
        to the entry with URL `/i-1-a`.

    6.  Now that both navigables have finished processing the traversal,
        we update the [current session history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-6}
        to the target step of 0.

2.  Now we can process the steps that were queued for the synchronous
    navigation:

    1.  The `/i-0-a#foo` entry is added, with its
        [step](#she-step){#centralized-modifications-of-session-history:she-step-2}
        determined to be 1: the [current session history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-7}
        (i.e., 0) plus 1. This also [clears existing forward
        history](#clear-the-forward-session-history){#centralized-modifications-of-session-history:clear-the-forward-session-history}.

    2.  We fully switch to that newly added entry, including calling
        [apply the history
        step](#apply-the-history-step){#centralized-modifications-of-session-history:apply-the-history-step-4}.
        This ultimately results in [updating the
        document](#update-document-for-history-step-application){#centralized-modifications-of-session-history:update-document-for-history-step-application-2}
        by dispatching events like
        [`hashchange`{#centralized-modifications-of-session-history:event-hashchange-2}](indices.html#event-hashchange),
        as well as updating the [current session history
        step](document-sequences.html#tn-current-session-history-step){#centralized-modifications-of-session-history:tn-current-session-history-step-8}
        to the target step of 1.

The end result is:

0

1

`top`

/t

`frames[0]`

/i-0-a

/i-0-a#foo

`frames[1]`

/i-1-a
:::

##### [7.4.1.4]{.secno} Low-level operations on session history[](#low-level-operations-on-session-history){.self-link}

This section contains a miscellaneous grab-bag of operations that we
perform throughout the standard when manipulating session history. The
best way to get a sense of what they do is to look at their call sites.

To [get session history entries]{#getting-session-history-entries .dfn}
of a
[navigable](document-sequences.html#navigable){#low-level-operations-on-session-history:navigable}
`navigable`{.variable}:

1.  Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#low-level-operations-on-session-history:nav-traversable}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#low-level-operations-on-session-history:assert
    x-internal="assert"}: this is running within
    `traversable`{.variable}\'s [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#low-level-operations-on-session-history:tn-session-history-traversal-queue}.

3.  If `navigable`{.variable} is `traversable`{.variable}, return
    `traversable`{.variable}\'s [session history
    entries](document-sequences.html#tn-session-history-entries){#low-level-operations-on-session-history:tn-session-history-entries}.

4.  Let `docStates`{.variable} be an empty [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#low-level-operations-on-session-history:set
    x-internal="set"} of [document
    states](#document-state-2){#low-level-operations-on-session-history:document-state-2}.

5.  For each `entry`{.variable} of `traversable`{.variable}\'s [session
    history
    entries](document-sequences.html#tn-session-history-entries){#low-level-operations-on-session-history:tn-session-history-entries-2},
    [append](https://infra.spec.whatwg.org/#set-append){#low-level-operations-on-session-history:set-append
    x-internal="set-append"} `entry`{.variable}\'s [document
    state](#she-document-state){#low-level-operations-on-session-history:she-document-state}
    to `docStates`{.variable}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate
    x-internal="list-iterate"} `docState`{.variable} of
    `docStates`{.variable}:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-2
        x-internal="list-iterate"} `nestedHistory`{.variable} of
        `docState`{.variable}\'s [nested
        histories](#document-state-nested-histories){#low-level-operations-on-session-history:document-state-nested-histories}:

        1.  If `nestedHistory`{.variable}\'s
            [id](#nested-history-id){#low-level-operations-on-session-history:nested-history-id}
            equals `navigable`{.variable}\'s
            [id](document-sequences.html#nav-id){#low-level-operations-on-session-history:nav-id},
            return `nestedHistory`{.variable}\'s
            [entries](#nested-history-entries){#low-level-operations-on-session-history:nested-history-entries}.

        2.  For each `entry`{.variable} of `nestedHistory`{.variable}\'s
            [entries](#nested-history-entries){#low-level-operations-on-session-history:nested-history-entries-2},
            [append](https://infra.spec.whatwg.org/#set-append){#low-level-operations-on-session-history:set-append-2
            x-internal="set-append"} `entry`{.variable}\'s [document
            state](#she-document-state){#low-level-operations-on-session-history:she-document-state-2}
            to `docStates`{.variable}.

7.  [Assert](https://infra.spec.whatwg.org/#assert){#low-level-operations-on-session-history:assert-2
    x-internal="assert"}: this step is not reached.

To [get session history entries for the navigation
API]{#getting-session-history-entries-for-the-navigation-api .dfn} of a
[navigable](document-sequences.html#navigable){#low-level-operations-on-session-history:navigable-2}
`navigable`{.variable} given an integer `targetStep`{.variable}:

1.  Let `rawEntries`{.variable} be the result of [getting session
    history
    entries](#getting-session-history-entries){#low-level-operations-on-session-history:getting-session-history-entries}
    for `navigable`{.variable}.

2.  Let `entriesForNavigationAPI`{.variable} be a new empty
    [list](https://infra.spec.whatwg.org/#list){#low-level-operations-on-session-history:list
    x-internal="list"}.

3.  Let `startingIndex`{.variable} be the index of the [session history
    entry](#session-history-entry){#low-level-operations-on-session-history:session-history-entry}
    in `rawEntries`{.variable} who has the greatest
    [step](#she-step){#low-level-operations-on-session-history:she-step}
    less than or equal to `targetStep`{.variable}.

    See [this example](#example-getting-the-target-history-entry) to
    understand why it\'s the greatest step less than or equal to
    `targetStep`{.variable}.

4.  [Append](https://infra.spec.whatwg.org/#list-append){#low-level-operations-on-session-history:list-append
    x-internal="list-append"}
    `rawEntries`{.variable}\[`startingIndex`{.variable}\] to
    `entriesForNavigationAPI`{.variable}.

5.  Let `startingOrigin`{.variable} be
    `rawEntries`{.variable}\[`startingIndex`{.variable}\]\'s [document
    state](#she-document-state){#low-level-operations-on-session-history:she-document-state-3}\'s
    [origin](#document-state-origin){#low-level-operations-on-session-history:document-state-origin}.

6.  Let `i`{.variable} be `startingIndex`{.variable} − 1.

7.  While `i`{.variable} \> 0:

    1.  If `rawEntries`{.variable}\[`i`{.variable}\]\'s [document
        state](#she-document-state){#low-level-operations-on-session-history:she-document-state-4}\'s
        [origin](#document-state-origin){#low-level-operations-on-session-history:document-state-origin-2}
        is not [same
        origin](browsers.html#same-origin){#low-level-operations-on-session-history:same-origin}
        with `startingOrigin`{.variable}, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#low-level-operations-on-session-history:break
        x-internal="break"}.

    2.  [Prepend](https://infra.spec.whatwg.org/#list-prepend){#low-level-operations-on-session-history:list-prepend
        x-internal="list-prepend"}
        `rawEntries`{.variable}\[`i`{.variable}\] to
        `entriesForNavigationAPI`{.variable}.

    3.  Set `i`{.variable} to `i`{.variable} − 1.

8.  Set `i`{.variable} to `startingIndex`{.variable} + 1.

9.  While `i`{.variable} \< `rawEntries`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#low-level-operations-on-session-history:list-size
    x-internal="list-size"}:

    1.  If `rawEntries`{.variable}\[`i`{.variable}\]\'s [document
        state](#she-document-state){#low-level-operations-on-session-history:she-document-state-5}\'s
        [origin](#document-state-origin){#low-level-operations-on-session-history:document-state-origin-3}
        is not [same
        origin](browsers.html#same-origin){#low-level-operations-on-session-history:same-origin-2}
        with `startingOrigin`{.variable}, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#low-level-operations-on-session-history:break-2
        x-internal="break"}.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#low-level-operations-on-session-history:list-append-2
        x-internal="list-append"}
        `rawEntries`{.variable}\[`i`{.variable}\] to
        `entriesForNavigationAPI`{.variable}.

    3.  Set `i`{.variable} to `i`{.variable} + 1.

10. Return `entriesForNavigationAPI`{.variable}.

To [clear the forward session
history]{#clear-the-forward-session-history .dfn} of a [traversable
navigable](document-sequences.html#traversable-navigable){#low-level-operations-on-session-history:traversable-navigable}
`navigable`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#low-level-operations-on-session-history:assert-3
    x-internal="assert"}: this is running within
    `navigable`{.variable}\'s [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#low-level-operations-on-session-history:tn-session-history-traversal-queue-2}.

2.  Let `step`{.variable} be the `navigable`{.variable}\'s [current
    session history
    step](document-sequences.html#tn-current-session-history-step){#low-level-operations-on-session-history:tn-current-session-history-step}.

3.  Let `entryLists`{.variable} be the [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#low-level-operations-on-session-history:set-2
    x-internal="set"} « `navigable`{.variable}\'s [session history
    entries](document-sequences.html#tn-session-history-entries){#low-level-operations-on-session-history:tn-session-history-entries-3}
    ».

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-3
    x-internal="list-iterate"} `entryList`{.variable} of
    `entryLists`{.variable}:

    1.  [Remove](https://infra.spec.whatwg.org/#list-remove){#low-level-operations-on-session-history:list-remove
        x-internal="list-remove"} every [session history
        entry](#session-history-entry){#low-level-operations-on-session-history:session-history-entry-2}
        from `entryList`{.variable} that has a
        [step](#she-step){#low-level-operations-on-session-history:she-step-2}
        greater than `step`{.variable}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-4
        x-internal="list-iterate"} `entry`{.variable} of
        `entryList`{.variable}:

        1.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-5
            x-internal="list-iterate"} `nestedHistory`{.variable} of
            `entry`{.variable}\'s [document
            state](#she-document-state){#low-level-operations-on-session-history:she-document-state-6}\'s
            [nested
            histories](#document-state-nested-histories){#low-level-operations-on-session-history:document-state-nested-histories-2},
            [append](https://infra.spec.whatwg.org/#set-append){#low-level-operations-on-session-history:set-append-3
            x-internal="set-append"} `nestedHistory`{.variable}\'s
            [entries
            list](#nested-history-entries){#low-level-operations-on-session-history:nested-history-entries-3}
            to `entryLists`{.variable}.

To [get all used history steps]{#getting-all-used-history-steps .dfn}
that are part of [traversable
navigable](document-sequences.html#traversable-navigable){#low-level-operations-on-session-history:traversable-navigable-2}
`traversable`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#low-level-operations-on-session-history:assert-4
    x-internal="assert"}: this is running within
    `traversable`{.variable}\'s [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#low-level-operations-on-session-history:tn-session-history-traversal-queue-3}.

2.  Let `steps`{.variable} be an empty [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#low-level-operations-on-session-history:set-3
    x-internal="set"} of non-negative integers.

3.  Let `entryLists`{.variable} be the [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#low-level-operations-on-session-history:set-4
    x-internal="set"} « `traversable`{.variable}\'s [session history
    entries](document-sequences.html#tn-session-history-entries){#low-level-operations-on-session-history:tn-session-history-entries-4}
    ».

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-6
    x-internal="list-iterate"} `entryList`{.variable} of
    `entryLists`{.variable}:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-7
        x-internal="list-iterate"} `entry`{.variable} of
        `entryList`{.variable}:

        1.  [Append](https://infra.spec.whatwg.org/#set-append){#low-level-operations-on-session-history:set-append-4
            x-internal="set-append"} `entry`{.variable}\'s
            [step](#she-step){#low-level-operations-on-session-history:she-step-3}
            to `steps`{.variable}.

        2.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#low-level-operations-on-session-history:list-iterate-8
            x-internal="list-iterate"} `nestedHistory`{.variable} of
            `entry`{.variable}\'s [document
            state](#she-document-state){#low-level-operations-on-session-history:she-document-state-7}\'s
            [nested
            histories](#document-state-nested-histories){#low-level-operations-on-session-history:document-state-nested-histories-3},
            [append](https://infra.spec.whatwg.org/#set-append){#low-level-operations-on-session-history:set-append-5
            x-internal="set-append"} `nestedHistory`{.variable}\'s
            [entries
            list](#nested-history-entries){#low-level-operations-on-session-history:nested-history-entries-4}
            to `entryLists`{.variable}.

5.  Return `steps`{.variable},
    [sorted](https://infra.spec.whatwg.org/#list-sort-in-ascending-order){#low-level-operations-on-session-history:list-sort
    x-internal="list-sort"}.

#### [7.4.2]{.secno} []{#browsing-the-web}Navigation[](#navigating-across-documents){.self-link} {#navigating-across-documents}

Certain actions cause a
[navigable](document-sequences.html#navigable){#navigating-across-documents:navigable}
to *[navigate](#navigate)* to a new resource.

For example, [following a
hyperlink](links.html#following-hyperlinks-2){#navigating-across-documents:following-hyperlinks-2},
[form
submission](form-control-infrastructure.html#concept-form-submit){#navigating-across-documents:concept-form-submit},
and the
[`window.open()`{#navigating-across-documents:dom-open}](nav-history-apis.html#dom-open)
and
[`location.assign()`{#navigating-across-documents:dom-location-assign}](nav-history-apis.html#dom-location-assign)
methods can all cause navigation.

::: {#note-meaning-of-navigate .note}
Although in this standard the word \"navigation\" refers specifically to
the [navigate](#navigate){#navigating-across-documents:navigate-2}
algorithm, this doesn\'t always line up with web developer or user
perceptions. For example:

- The [URL and history update
  steps](#url-and-history-update-steps){#navigating-across-documents:url-and-history-update-steps}
  are often used during so-called \"single-page app navigations\" or
  \"same-document navigations\", but they do not trigger the
  [navigate](#navigate){#navigating-across-documents:navigate-3}
  algorithm.

- [Reloads](#reload){#navigating-across-documents:reload} and
  [traversals](#traverse-the-history-by-a-delta){#navigating-across-documents:traverse-the-history-by-a-delta}
  are sometimes talked about as a type of navigation, since all three
  will often [attempt to populate the history entry\'s
  document](#attempt-to-populate-the-history-entry's-document){#navigating-across-documents:attempt-to-populate-the-history-entry's-document}
  and thus could perform navigational fetches. See, e.g., the APIs
  exposed Navigation Timing. But they have their own entry point
  algorithms, separate from the
  [navigate](#navigate){#navigating-across-documents:navigate-4}
  algorithm.
  [\[NAVIGATIONTIMING\]](references.html#refsNAVIGATIONTIMING)

- Although [fragment
  navigations](#navigate-fragid){#navigating-across-documents:navigate-fragid}
  are always done through the
  [navigate](#navigate){#navigating-across-documents:navigate-5}
  algorithm, a user might perceive them as more like jumping around a
  single page, than as a true navigation.
:::

##### [7.4.2.1]{.secno} Supporting concepts[](#navigation-supporting-concepts){.self-link} {#navigation-supporting-concepts}

Before we can jump into the [navigation
algorithm](#navigate){#navigation-supporting-concepts:navigate} itself,
we need to establish several important structures that it uses.

The [source snapshot params]{#source-snapshot-params .dfn}
[struct](https://infra.spec.whatwg.org/#struct){#navigation-supporting-concepts:struct
x-internal="struct"} is used to capture data from a
[`Document`{#navigation-supporting-concepts:document}](dom.html#document)
initiating a navigation. It is snapshotted at the beginning of a
navigation and used throughout the navigation\'s lifetime. It has the
following
[items](https://infra.spec.whatwg.org/#struct-item){#navigation-supporting-concepts:struct-item
x-internal="struct-item"}:

[has transient activation]{#source-snapshot-params-activation .dfn}
:   a boolean

[sandboxing flags]{#source-snapshot-params-sandbox .dfn}
:   a [sandboxing flag
    set](browsers.html#sandboxing-flag-set){#navigation-supporting-concepts:sandboxing-flag-set}

[allows downloading]{#source-snapshot-params-download .dfn}
:   a boolean

[fetch client]{#source-snapshot-params-client .dfn}
:   an [environment settings
    object](webappapis.html#environment-settings-object){#navigation-supporting-concepts:environment-settings-object}
    or null, only to be used as a [request
    client](https://fetch.spec.whatwg.org/#concept-request-client){#navigation-supporting-concepts:concept-request-client
    x-internal="concept-request-client"}

[source policy container]{#source-snapshot-params-policy-container .dfn}
:   a [policy
    container](browsers.html#policy-container){#navigation-supporting-concepts:policy-container}

To [snapshot source snapshot
params]{#snapshotting-source-snapshot-params .dfn} given a
[`Document`{#navigation-supporting-concepts:document-2}](dom.html#document)-or-null
`sourceDocument`{.variable}:

1.  If `sourceDocument`{.variable} is null, then return a new [source
    snapshot
    params](#source-snapshot-params){#navigation-supporting-concepts:source-snapshot-params}
    with

    [has transient activation](#source-snapshot-params-activation){#navigation-supporting-concepts:source-snapshot-params-activation}
    :   true

    [sandboxing flags](#source-snapshot-params-sandbox){#navigation-supporting-concepts:source-snapshot-params-sandbox}
    :   an empty [sandboxing flag
        set](browsers.html#sandboxing-flag-set){#navigation-supporting-concepts:sandboxing-flag-set-2}

    [allows downloading](#source-snapshot-params-download){#navigation-supporting-concepts:source-snapshot-params-download}
    :   true

    [fetch client](#source-snapshot-params-client){#navigation-supporting-concepts:source-snapshot-params-client}
    :   null

    [source policy container](#source-snapshot-params-policy-container){#navigation-supporting-concepts:source-snapshot-params-policy-container}
    :   a new [policy
        container](browsers.html#policy-container){#navigation-supporting-concepts:policy-container-2}

    This [only occurs](#assert-null-sourcedocument) in the case of a
    browser UI-initiated navigation.

2.  Return a new [source snapshot
    params](#source-snapshot-params){#navigation-supporting-concepts:source-snapshot-params-2}
    with

    [has transient activation](#source-snapshot-params-activation){#navigation-supporting-concepts:source-snapshot-params-activation-2}
    :   true if `sourceDocument`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#navigation-supporting-concepts:concept-relevant-global}
        has [transient
        activation](interaction.html#transient-activation){#navigation-supporting-concepts:transient-activation};
        otherwise false

    [sandboxing flags](#source-snapshot-params-sandbox){#navigation-supporting-concepts:source-snapshot-params-sandbox-2}
    :   `sourceDocument`{.variable}\'s [active sandboxing flag
        set](browsers.html#active-sandboxing-flag-set){#navigation-supporting-concepts:active-sandboxing-flag-set}

    [allows downloading](#source-snapshot-params-download){#navigation-supporting-concepts:source-snapshot-params-download-2}
    :   false if `sourceDocument`{.variable}\'s [active sandboxing flag
        set](browsers.html#active-sandboxing-flag-set){#navigation-supporting-concepts:active-sandboxing-flag-set-2}
        has the [sandboxed downloads browsing context
        flag](browsers.html#sandboxed-downloads-browsing-context-flag){#navigation-supporting-concepts:sandboxed-downloads-browsing-context-flag}
        set; otherwise true

    [fetch client](#source-snapshot-params-client){#navigation-supporting-concepts:source-snapshot-params-client-2}
    :   `sourceDocument`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#navigation-supporting-concepts:relevant-settings-object}

    [source policy container](#source-snapshot-params-policy-container){#navigation-supporting-concepts:source-snapshot-params-policy-container-2}
    :   a
        [clone](browsers.html#clone-a-policy-container){#navigation-supporting-concepts:clone-a-policy-container}
        of `sourceDocument`{.variable}\'s [policy
        container](dom.html#concept-document-policy-container){#navigation-supporting-concepts:concept-document-policy-container}

------------------------------------------------------------------------

The [target snapshot params]{#target-snapshot-params .dfn}
[struct](https://infra.spec.whatwg.org/#struct){#navigation-supporting-concepts:struct-2
x-internal="struct"} is used to capture data from a
[navigable](document-sequences.html#navigable){#navigation-supporting-concepts:navigable}
being navigated. Like [source snapshot
params](#source-snapshot-params){#navigation-supporting-concepts:source-snapshot-params-3},
it is snapshotted at the beginning of a navigation and used throughout
the navigation\'s lifetime. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#navigation-supporting-concepts:struct-item-2
x-internal="struct-item"}:

[sandboxing flags]{#target-snapshot-params-sandbox .dfn}
:   a [sandboxing flag
    set](browsers.html#sandboxing-flag-set){#navigation-supporting-concepts:sandboxing-flag-set-3}

To [snapshot target snapshot
params]{#snapshotting-target-snapshot-params .dfn} given a
[navigable](document-sequences.html#navigable){#navigation-supporting-concepts:navigable-2}
`targetNavigable`{.variable}, return a new [target snapshot
params](#target-snapshot-params){#navigation-supporting-concepts:target-snapshot-params}
with [sandboxing
flags](#target-snapshot-params-sandbox){#navigation-supporting-concepts:target-snapshot-params-sandbox}
set to the result of [determining the creation sandboxing
flags](browsers.html#determining-the-creation-sandboxing-flags){#navigation-supporting-concepts:determining-the-creation-sandboxing-flags}
given `targetNavigable`{.variable}\'s [active browsing
context](document-sequences.html#nav-bc){#navigation-supporting-concepts:nav-bc}
and `targetNavigable`{.variable}\'s
[container](document-sequences.html#nav-container){#navigation-supporting-concepts:nav-container}.

------------------------------------------------------------------------

Much of the navigation process is concerned with determining how to
create a new
[`Document`{#navigation-supporting-concepts:document-3}](dom.html#document),
which ultimately happens in the [create and initialize a `Document`
object](document-lifecycle.html#initialise-the-document-object){#navigation-supporting-concepts:initialise-the-document-object}
algorithm. The parameters to that algorithm are tracked via a
[navigation params]{#navigation-params .dfn}
[struct](https://infra.spec.whatwg.org/#struct){#navigation-supporting-concepts:struct-3
x-internal="struct"}, which has the following
[items](https://infra.spec.whatwg.org/#struct-item){#navigation-supporting-concepts:struct-item-3
x-internal="struct-item"}:

[id]{#navigation-params-id .dfn}
:   null or a [navigation
    ID](#navigation-id){#navigation-supporting-concepts:navigation-id}

[navigable]{#navigation-params-navigable .dfn}
:   the
    [navigable](document-sequences.html#navigable){#navigation-supporting-concepts:navigable-3}
    to be navigated

[request]{#navigation-params-request .dfn}
:   null or a
    [request](https://fetch.spec.whatwg.org/#concept-request){#navigation-supporting-concepts:concept-request
    x-internal="concept-request"} that started the navigation

[response]{#navigation-params-response .dfn}
:   a
    [response](https://fetch.spec.whatwg.org/#concept-response){#navigation-supporting-concepts:concept-response
    x-internal="concept-response"} that ultimately was navigated to
    (potentially a [network
    error](https://fetch.spec.whatwg.org/#concept-network-error){#navigation-supporting-concepts:network-error
    x-internal="network-error"})

[fetch controller]{#navigation-params-fetch-controller .dfn}
:   null or a [fetch
    controller](https://fetch.spec.whatwg.org/#fetch-controller){#navigation-supporting-concepts:fetch-controller
    x-internal="fetch-controller"}

[commit early hints]{#navigation-params-commit-early-hints .dfn}
:   null or an algorithm accepting a
    [`Document`{#navigation-supporting-concepts:document-4}](dom.html#document),
    once it has been created

[COOP enforcement result]{#navigation-params-coop-enforcement-result .dfn}
:   an [opener policy enforcement
    result](browsers.html#coop-enforcement-result){#navigation-supporting-concepts:coop-enforcement-result},
    used for reporting and potentially for causing a [browsing context
    group
    switch](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy)

[reserved environment]{#navigation-params-reserved-environment .dfn}
:   null or an
    [environment](webappapis.html#environment){#navigation-supporting-concepts:environment}
    reserved for the new
    [`Document`{#navigation-supporting-concepts:document-5}](dom.html#document)

[origin]{#navigation-params-origin .dfn}
:   an
    [origin](browsers.html#concept-origin){#navigation-supporting-concepts:concept-origin}
    to use for the new
    [`Document`{#navigation-supporting-concepts:document-6}](dom.html#document)

[policy container]{#navigation-params-policy-container .dfn}
:   a [policy
    container](browsers.html#policy-container){#navigation-supporting-concepts:policy-container-3}
    to use for the new
    [`Document`{#navigation-supporting-concepts:document-7}](dom.html#document)

[final sandboxing flag set]{#navigation-params-sandboxing .dfn}
:   a [sandboxing flag
    set](browsers.html#sandboxing-flag-set){#navigation-supporting-concepts:sandboxing-flag-set-4}
    to impose on the new
    [`Document`{#navigation-supporting-concepts:document-8}](dom.html#document)

[opener policy]{#navigation-params-coop .dfn}
:   an [opener
    policy](browsers.html#cross-origin-opener-policy){#navigation-supporting-concepts:cross-origin-opener-policy}
    to use for the new
    [`Document`{#navigation-supporting-concepts:document-9}](dom.html#document)

[navigation timing type]{#navigation-params-nav-timing-type .dfn}
:   a
    [`NavigationTimingType`{#navigation-supporting-concepts:navigationtimingtype}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
    used for [creating the navigation timing
    entry](https://w3c.github.io/navigation-timing/#dfn-create-the-navigation-timing-entry){#navigation-supporting-concepts:create-the-navigation-timing-entry
    x-internal="create-the-navigation-timing-entry"} for the new
    [`Document`{#navigation-supporting-concepts:document-10}](dom.html#document)

[about base URL]{#navigation-params-about-base-url .dfn}
:   a
    [URL](https://url.spec.whatwg.org/#concept-url){#navigation-supporting-concepts:url
    x-internal="url"} or null used to populate the new
    [`Document`{#navigation-supporting-concepts:document-11}](dom.html#document)\'s
    [about base
    URL](dom.html#concept-document-about-base-url){#navigation-supporting-concepts:concept-document-about-base-url}

[user involvement]{#navigation-params-user-involvement .dfn}
:   a [user navigation
    involvement](#user-navigation-involvement){#navigation-supporting-concepts:user-navigation-involvement}
    used when [obtaining a browsing
    context](browsers.html#obtain-browsing-context-navigation){#navigation-supporting-concepts:obtain-browsing-context-navigation}
    for the new
    [`Document`{#navigation-supporting-concepts:document-12}](dom.html#document)

Once a [navigation
params](#navigation-params){#navigation-supporting-concepts:navigation-params}
struct is created, this standard does not mutate any of its
[items](https://infra.spec.whatwg.org/#struct-item){#navigation-supporting-concepts:struct-item-4
x-internal="struct-item"}. They are only passed onward to other
algorithms.

------------------------------------------------------------------------

A [navigation ID]{#navigation-id .dfn} is a UUID string generated during
navigation. It is used to interface with the WebDriver BiDi
specification as well as to track the [ongoing
navigation](#ongoing-navigation){#navigation-supporting-concepts:ongoing-navigation}.
[\[WEBDRIVERBIDI\]](references.html#refsWEBDRIVERBIDI)

------------------------------------------------------------------------

After
[`Document`{#navigation-supporting-concepts:document-13}](dom.html#document)
creation, the relevant [traversable
navigable](document-sequences.html#traversable-navigable){#navigation-supporting-concepts:traversable-navigable}\'s
[session
history](document-sequences.html#tn-session-history-entries){#navigation-supporting-concepts:tn-session-history-entries}
gets updated. The
[`NavigationHistoryBehavior`{#navigation-supporting-concepts:navigationhistorybehavior}](nav-history-apis.html#navigationhistorybehavior)
enumeration is used to indicate the desired type of session history
update to the
[navigate](#navigate){#navigation-supporting-concepts:navigate-2}
algorithm. It is one of the following:

[]{#hh-push}\"[`push`]{#navigationhistorybehavior-push .dfn dfn-for="NavigationHistoryBehavior" dfn-type="enum-value"}\"
:   A regular navigation which adds a new [session history
    entry](#session-history-entry){#navigation-supporting-concepts:session-history-entry},
    and will [clear the forward session
    history](#clear-the-forward-session-history){#navigation-supporting-concepts:clear-the-forward-session-history}.

\"[`replace`]{#navigationhistorybehavior-replace .dfn dfn-for="NavigationHistoryBehavior" dfn-type="enum-value"}\"
:   A navigation that will replace the [active session history
    entry](document-sequences.html#nav-active-history-entry){#navigation-supporting-concepts:nav-active-history-entry}.

\"[`auto`]{#navigationhistorybehavior-auto .dfn dfn-for="NavigationHistoryBehavior" dfn-type="enum-value"}\"
:   The default value, which will be converted very early in the
    [navigate](#navigate){#navigation-supporting-concepts:navigate-3}
    algorithm into
    \"[`push`{#navigation-supporting-concepts:navigationhistorybehavior-push}](#navigationhistorybehavior-push)\"
    or
    \"[`replace`{#navigation-supporting-concepts:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\".
    Usually it becomes
    \"[`push`{#navigation-supporting-concepts:navigationhistorybehavior-push-2}](#navigationhistorybehavior-push)\",
    but under [certain circumstances](#navigate-convert-to-replace) it
    becomes
    \"[`replace`{#navigation-supporting-concepts:navigationhistorybehavior-replace-2}](#navigationhistorybehavior-replace)\"
    instead.

A [history handling behavior]{#history-handling-behavior .dfn} is a
[`NavigationHistoryBehavior`{#navigation-supporting-concepts:navigationhistorybehavior-2}](nav-history-apis.html#navigationhistorybehavior)
that is either
\"[`push`{#navigation-supporting-concepts:navigationhistorybehavior-push-3}](#navigationhistorybehavior-push)\"
or
\"[`replace`{#navigation-supporting-concepts:navigationhistorybehavior-replace-3}](#navigationhistorybehavior-replace)\",
i.e., that has been resolved away from any initial
\"[`auto`{#navigation-supporting-concepts:navigationhistorybehavior-auto}](#navigationhistorybehavior-auto)\"
value.

[The navigation must be a replace]{#the-navigation-must-be-a-replace
.dfn}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#navigation-supporting-concepts:url-2
x-internal="url"} `url`{.variable} and a
[`Document`{#navigation-supporting-concepts:document-14}](dom.html#document)
`document`{.variable}, if any of the following are true:

- `url`{.variable}\'s
  [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#navigation-supporting-concepts:concept-url-scheme
  x-internal="concept-url-scheme"} is
  \"[`javascript`{#navigation-supporting-concepts:the-javascript:-url-special-case}](#the-javascript:-url-special-case)\";
  or

- `document`{.variable}\'s [is initial
  `about:blank`](dom.html#is-initial-about:blank){#navigation-supporting-concepts:is-initial-about:blank}
  is true.

::: note
Other cases that often, but not always, force a
\"[`replace`{#navigation-supporting-concepts:navigationhistorybehavior-replace-4}](#navigationhistorybehavior-replace)\"
navigation are:

- if the
  [`Document`{#navigation-supporting-concepts:document-15}](dom.html#document)
  is not [completely
  loaded](document-lifecycle.html#completely-loaded){#navigation-supporting-concepts:completely-loaded};
  or

- if the target
  [URL](https://url.spec.whatwg.org/#concept-url){#navigation-supporting-concepts:url-3
  x-internal="url"} equals the
  [`Document`{#navigation-supporting-concepts:document-16}](dom.html#document)\'s
  [URL](https://dom.spec.whatwg.org/#concept-document-url){#navigation-supporting-concepts:the-document's-address
  x-internal="the-document's-address"}.
:::

------------------------------------------------------------------------

Various parts of the platform track whether a user is involved in a
navigation. A [user navigation involvement]{#user-navigation-involvement
.dfn export=""} is one of the following:

\"[`browser UI`]{#uni-browser-ui .dfn dfn-for="user navigation involvement" export=""}\"
:   The navigation was initiated by the user via browser UI mechanisms.

\"[`activation`]{#uni-activation .dfn dfn-for="user navigation involvement" export=""}\"
:   The navigation was initiated by the user via the [activation
    behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#navigation-supporting-concepts:activation-behaviour
    x-internal="activation-behaviour"} of an element.

\"[`none`]{#uni-none .dfn dfn-for="user navigation involvement" export=""}\"
:   The navigation was not initiated by the user.

For convenience at certain call sites, the [user navigation
involvement]{#event-uni .dfn dfn-for="Event"} for an
[`Event`{#navigation-supporting-concepts:event}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}
`event`{.variable} is defined as follows:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-supporting-concepts:assert
    x-internal="assert"}: this algorithm is being called as part of an
    [activation
    behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#navigation-supporting-concepts:activation-behaviour-2
    x-internal="activation-behaviour"} definition.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-supporting-concepts:assert-2
    x-internal="assert"}: `event`{.variable}\'s
    [`type`{#navigation-supporting-concepts:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`click`{#navigation-supporting-concepts:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}\".

3.  If `event`{.variable}\'s
    [`isTrusted`{#navigation-supporting-concepts:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    is initialized to true, then return
    \"[`activation`{#navigation-supporting-concepts:uni-activation}](#uni-activation)\".

4.  Return
    \"[`none`{#navigation-supporting-concepts:uni-none}](#uni-none)\".

##### [7.4.2.2]{.secno} Beginning navigation[](#beginning-navigation){.self-link}

To [navigate]{#navigate .dfn export=""} a
[navigable](document-sequences.html#navigable){#beginning-navigation:navigable}
`navigable`{.variable} to a
[URL](https://url.spec.whatwg.org/#concept-url){#beginning-navigation:url
x-internal="url"} `url`{.variable} using an optional
[`Document`{#beginning-navigation:document}](dom.html#document)-or-null
`sourceDocument`{#source-browsing-context .variable} (default null),
with an optional [POST
resource](#post-resource){#beginning-navigation:post-resource}, string,
or null [`documentResource`{.variable}]{#navigation-resource .dfn
dfn-for="navigate" export=""} (default null), an optional
[response](https://fetch.spec.whatwg.org/#concept-response){#beginning-navigation:concept-response
x-internal="concept-response"}-or-null
[`response`{.variable}]{#navigation-response .dfn dfn-for="navigate"
export=""} (default null), an optional boolean
[`exceptionsEnabled`{.variable}]{#exceptions-enabled .dfn
dfn-for="navigate" export=""} (default false), an optional
[`NavigationHistoryBehavior`{#beginning-navigation:navigationhistorybehavior}](nav-history-apis.html#navigationhistorybehavior)
[`historyHandling`{.variable}]{#navigation-hh .dfn dfn-for="navigate"
export=""} (default
\"[`auto`{#beginning-navigation:navigationhistorybehavior-auto}](#navigationhistorybehavior-auto)\"),
an optional [serialized
state](#serialized-state){#beginning-navigation:serialized-state}-or-null
[`navigationAPIState`{.variable}]{#navigation-navigation-api-state .dfn
dfn-for="navigate" export=""} (default null), an optional [entry
list](form-control-infrastructure.html#entry-list){#beginning-navigation:entry-list}
or null
[`formDataEntryList`{.variable}]{#navigation-form-data-entry-list .dfn
dfn-for="navigate" export=""} (default null), an optional [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#beginning-navigation:referrer-policy
x-internal="referrer-policy"}
[`referrerPolicy`{.variable}]{#navigation-referrer-policy .dfn} (default
the empty string), an optional [user navigation
involvement](#user-navigation-involvement){#beginning-navigation:user-navigation-involvement}
[`userInvolvement`{.variable}]{#navigation-user-involvement .dfn}
(default \"[`none`{#beginning-navigation:uni-none}](#uni-none)\"), an
optional
[`Element`{#beginning-navigation:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
[`sourceElement`{.variable}]{#navigation-source-element .dfn} (default
null), and an optional boolean
[`initialInsertion`{.variable}]{#navigation-initial-insertion .dfn
dfn-for="navigate"} (default false):

1.  Let `cspNavigationType`{.variable} be \"`form-submission`\" if
    `formDataEntryList`{.variable} is non-null; otherwise \"`other`\".

2.  Let `sourceSnapshotParams`{.variable} be the result of [snapshotting
    source snapshot
    params](#snapshotting-source-snapshot-params){#beginning-navigation:snapshotting-source-snapshot-params}
    given `sourceDocument`{.variable}.

3.  Let `initiatorOriginSnapshot`{.variable} be a new [opaque
    origin](browsers.html#concept-origin-opaque){#beginning-navigation:concept-origin-opaque}.

4.  Let `initiatorBaseURLSnapshot`{.variable} be
    [`about:blank`{#beginning-navigation:about:blank}](infrastructure.html#about:blank).

5.  If `sourceDocument`{.variable} is null:

    1.  [[Assert](https://infra.spec.whatwg.org/#assert){#beginning-navigation:assert
        x-internal="assert"}: `userInvolvement`{.variable} is
        \"[`browser UI`{#beginning-navigation:uni-browser-ui}](#uni-browser-ui)\".]{#assert-null-sourcedocument}

    2.  If `url`{.variable}\'s
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#beginning-navigation:concept-url-scheme
        x-internal="concept-url-scheme"} is
        \"[`javascript`{#beginning-navigation:the-javascript:-url-special-case}](#the-javascript:-url-special-case)\",
        then set `initiatorOriginSnapshot`{.variable} to
        `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#beginning-navigation:nav-document}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#beginning-navigation:concept-document-origin
        x-internal="concept-document-origin"}.

6.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#beginning-navigation:assert-2
        x-internal="assert"}: `userInvolvement`{.variable} is not
        \"[`browser UI`{#beginning-navigation:uni-browser-ui-2}](#uni-browser-ui)\".

    2.  ::: {#sandboxLinks}
        If `sourceDocument`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#beginning-navigation:node-navigable}
        is not [allowed by sandboxing to
        navigate](#allowed-to-navigate){#beginning-navigation:allowed-to-navigate}
        `navigable`{.variable} given `sourceSnapshotParams`{.variable}:

        1.  If `exceptionsEnabled`{.variable} is true, then throw a
            [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#beginning-navigation:securityerror
            x-internal="securityerror"}
            [`DOMException`{#beginning-navigation:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Return.
        :::

    3.  Set `initiatorOriginSnapshot`{.variable} to
        `sourceDocument`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#beginning-navigation:concept-document-origin-2
        x-internal="concept-document-origin"}.

    4.  Set `initiatorBaseURLSnapshot`{.variable} to
        `sourceDocument`{.variable}\'s [document base
        URL](urls-and-fetching.html#document-base-url){#beginning-navigation:document-base-url}.

7.  Let `navigationId`{.variable} be the result of [generating a random
    UUID](https://w3c.github.io/webcrypto/#dfn-generate-a-random-uuid){#beginning-navigation:generating-a-random-uuid
    x-internal="generating-a-random-uuid"}.
    [\[WEBCRYPTO\]](references.html#refsWEBCRYPTO)

8.  If the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#beginning-navigation:surrounding-agent
    x-internal="surrounding-agent"} is equal to
    `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#beginning-navigation:nav-document-2}\'s
    [relevant
    agent](webappapis.html#relevant-agent){#beginning-navigation:relevant-agent},
    then continue these steps. Otherwise, [queue a global
    task](webappapis.html#queue-a-global-task){#beginning-navigation:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#beginning-navigation:navigation-and-traversal-task-source}
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#beginning-navigation:nav-window}
    to continue these steps.

    ::: note
    We do this because we are about to look at a lot of properties of
    `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#beginning-navigation:nav-document-3},
    which are in theory only accessible over in the appropriate [event
    loop](webappapis.html#event-loop){#beginning-navigation:event-loop}.
    (But, we do not want to unconditionally queue a task, since --- for
    example --- same-event-loop [fragment
    navigations](#navigate-fragid){#beginning-navigation:navigate-fragid}
    need to take effect synchronously.)

    Another implementation strategy would be to replicate the relevant
    information across event loops, or into a canonical \"browser
    process\", so that it can be consulted without queueing a task. This
    could give different results than what we specify here in edge
    cases, where the relevant properties have changed over in the target
    event loop but not yet been replicated. Further testing is needed to
    determine which of these strategies best matches browser behavior,
    in such racy edge cases.
    :::

9.  If `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#beginning-navigation:nav-document-4}\'s
    [unload
    counter](document-lifecycle.html#unload-counter){#beginning-navigation:unload-counter}
    is greater than 0, then invoke [WebDriver BiDi navigation
    failed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-failed){#beginning-navigation:webdriver-bidi-navigation-failed
    x-internal="webdriver-bidi-navigation-failed"} with
    `navigable`{.variable} and a [WebDriver BiDi navigation
    status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#beginning-navigation:webdriver-bidi-navigation-status
    x-internal="webdriver-bidi-navigation-status"} whose
    [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#beginning-navigation:navigation-status-id
    x-internal="navigation-status-id"} is `navigationId`{.variable},
    [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#beginning-navigation:navigation-status-status
    x-internal="navigation-status-status"} is
    \"[`canceled`{#beginning-navigation:navigation-status-canceled}](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled){x-internal="navigation-status-canceled"}\",
    and
    [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#beginning-navigation:navigation-status-url
    x-internal="navigation-status-url"} is `url`{.variable}, and return.

10. Let `container`{.variable} be `navigable`{.variable}\'s
    [container](document-sequences.html#nav-container){#beginning-navigation:nav-container}.

11. If `container`{.variable} is an
    [`iframe`{#beginning-navigation:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    element and [will lazy load element
    steps](urls-and-fetching.html#will-lazy-load-element-steps){#beginning-navigation:will-lazy-load-element-steps}
    given `container`{.variable} returns true, then [stop
    intersection-observing a lazy loading
    element](urls-and-fetching.html#stop-intersection-observing-a-lazy-loading-element){#beginning-navigation:stop-intersection-observing-a-lazy-loading-element}
    `container`{.variable} and set `container`{.variable}\'s [lazy load
    resumption
    steps](urls-and-fetching.html#lazy-load-resumption-steps){#beginning-navigation:lazy-load-resumption-steps}
    to null.

12. ::: {#navigate-convert-to-replace}
    If `historyHandling`{.variable} is
    \"[`auto`{#beginning-navigation:navigationhistorybehavior-auto-2}](#navigationhistorybehavior-auto)\",
    then:

    1.  If `url`{.variable}
        [equals](https://url.spec.whatwg.org/#concept-url-equals){#beginning-navigation:concept-url-equals
        x-internal="concept-url-equals"} `navigable`{.variable}\'s
        [active
        document](document-sequences.html#nav-document){#beginning-navigation:nav-document-5}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#beginning-navigation:the-document's-address
        x-internal="the-document's-address"}, and
        `initiatorOriginSnapshot`{.variable} is [same
        origin](browsers.html#same-origin){#beginning-navigation:same-origin}
        with `targetNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#beginning-navigation:nav-document-6}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#beginning-navigation:concept-document-origin-3
        x-internal="concept-document-origin"}, then set
        `historyHandling`{.variable} to
        \"[`replace`{#beginning-navigation:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\".

    2.  Otherwise, set `historyHandling`{.variable} to
        \"[`push`{#beginning-navigation:navigationhistorybehavior-push}](#navigationhistorybehavior-push)\".
    :::

13. If [the navigation must be a
    replace](#the-navigation-must-be-a-replace){#beginning-navigation:the-navigation-must-be-a-replace}
    given `url`{.variable} and `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#beginning-navigation:nav-document-7},
    then set `historyHandling`{.variable} to
    \"[`replace`{#beginning-navigation:navigationhistorybehavior-replace-2}](#navigationhistorybehavior-replace)\".

14. ::: {#navigate-fragid-step}
    If all of the following are true:

    - `documentResource`{.variable} is null;

    - `response`{.variable} is null;

    - `url`{.variable}
      [equals](https://url.spec.whatwg.org/#concept-url-equals){#beginning-navigation:concept-url-equals-2
      x-internal="concept-url-equals"} `navigable`{.variable}\'s [active
      session history
      entry](document-sequences.html#nav-active-history-entry){#beginning-navigation:nav-active-history-entry}\'s
      [URL](#she-url){#beginning-navigation:she-url} with [*exclude
      fragments*](https://url.spec.whatwg.org/#url-equals-exclude-fragments){#beginning-navigation:url-equals-exclude-fragments
      x-internal="url-equals-exclude-fragments"} set to true; and

    - `url`{.variable}\'s
      [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#beginning-navigation:concept-url-fragment
      x-internal="concept-url-fragment"} is non-null,

    then:

    1.  [Navigate to a
        fragment](#navigate-fragid){#beginning-navigation:navigate-fragid-2}
        given `navigable`{.variable}, `url`{.variable},
        `historyHandling`{.variable}, `userInvolvement`{.variable},
        `sourceElement`{.variable}, `navigationAPIState`{.variable}, and
        `navigationId`{.variable}.

    2.  Return.
    :::

15. If `navigable`{.variable}\'s
    [parent](document-sequences.html#nav-parent){#beginning-navigation:nav-parent}
    is non-null, then set `navigable`{.variable}\'s [is delaying `load`
    events](document-sequences.html#delaying-load-events-mode){#beginning-navigation:delaying-load-events-mode}
    to true.

16. Let `targetSnapshotParams`{.variable} be the result of [snapshotting
    target snapshot
    params](#snapshotting-target-snapshot-params){#beginning-navigation:snapshotting-target-snapshot-params}
    given `navigable`{.variable}.

17. Invoke [WebDriver BiDi navigation
    started](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-started){#beginning-navigation:webdriver-bidi-navigation-started
    x-internal="webdriver-bidi-navigation-started"} with
    `navigable`{.variable} and a new [WebDriver BiDi navigation
    status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#beginning-navigation:webdriver-bidi-navigation-status-2
    x-internal="webdriver-bidi-navigation-status"} whose
    [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#beginning-navigation:navigation-status-id-2
    x-internal="navigation-status-id"} is `navigationId`{.variable},
    [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#beginning-navigation:navigation-status-status-2
    x-internal="navigation-status-status"} is
    \"[`pending`{#beginning-navigation:navigation-status-pending}](https://w3c.github.io/webdriver-bidi/#navigation-status-pending){x-internal="navigation-status-pending"}\",
    and
    [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#beginning-navigation:navigation-status-url-2
    x-internal="navigation-status-url"} is `url`{.variable}.

18. If `navigable`{.variable}\'s [ongoing
    navigation](#ongoing-navigation){#beginning-navigation:ongoing-navigation}
    is \"`traversal`\", then:

    1.  Invoke [WebDriver BiDi navigation
        failed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-failed){#beginning-navigation:webdriver-bidi-navigation-failed-2
        x-internal="webdriver-bidi-navigation-failed"} with
        `navigable`{.variable} and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#beginning-navigation:webdriver-bidi-navigation-status-3
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#beginning-navigation:navigation-status-id-3
        x-internal="navigation-status-id"} is `navigationId`{.variable},
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#beginning-navigation:navigation-status-status-3
        x-internal="navigation-status-status"} is
        \"[`canceled`{#beginning-navigation:navigation-status-canceled-2}](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled){x-internal="navigation-status-canceled"}\",
        and
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#beginning-navigation:navigation-status-url-3
        x-internal="navigation-status-url"} is `url`{.variable}.

    2.  Return.

    Any attempts to navigate a
    [navigable](document-sequences.html#navigable){#beginning-navigation:navigable-2}
    that is currently
    [traversing](#apply-the-traverse-history-step){#beginning-navigation:apply-the-traverse-history-step}
    are ignored.

19. [Set the ongoing
    navigation](#set-the-ongoing-navigation){#beginning-navigation:set-the-ongoing-navigation}
    for `navigable`{.variable} to `navigationId`{.variable}.

    This will have the effect of aborting other ongoing navigations of
    `navigable`{.variable}, since at certain points during navigation
    changes to the [ongoing
    navigation](#ongoing-navigation){#beginning-navigation:ongoing-navigation-2}
    will cause further work to be abandoned.

20. If `url`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#beginning-navigation:concept-url-scheme-2
    x-internal="concept-url-scheme"} is
    \"[`javascript`{#beginning-navigation:the-javascript:-url-special-case-2}](#the-javascript:-url-special-case)\",
    then:

    1.  [Queue a global
        task](webappapis.html#queue-a-global-task){#beginning-navigation:queue-a-global-task-2}
        on the [navigation and traversal task
        source](webappapis.html#navigation-and-traversal-task-source){#beginning-navigation:navigation-and-traversal-task-source-2}
        given `navigable`{.variable}\'s [active
        window](document-sequences.html#nav-window){#beginning-navigation:nav-window-2}
        to [navigate to a `javascript:`
        URL](#navigate-to-a-javascript:-url){#beginning-navigation:navigate-to-a-javascript:-url}
        given `navigable`{.variable}, `url`{.variable},
        `historyHandling`{.variable}, `sourceSnapshotParams`{.variable},
        `initiatorOriginSnapshot`{.variable},
        `userInvolvement`{.variable}, `cspNavigationType`{.variable},
        and `initialInsertion`{.variable}.

    2.  Return.

21. If all of the following are true:

    - `userInvolvement`{.variable} is not
      \"[`browser UI`{#beginning-navigation:uni-browser-ui-3}](#uni-browser-ui)\";

    - `navigable`{.variable}\'s [active
      document](document-sequences.html#nav-document){#beginning-navigation:nav-document-8}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#beginning-navigation:concept-document-origin-4
      x-internal="concept-document-origin"} is [same
      origin-domain](browsers.html#same-origin-domain){#beginning-navigation:same-origin-domain}
      with `sourceDocument`{.variable}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#beginning-navigation:concept-document-origin-5
      x-internal="concept-document-origin"};

    - `navigable`{.variable}\'s [active
      document](document-sequences.html#nav-document){#beginning-navigation:nav-document-9}\'s
      [is initial
      `about:blank`](dom.html#is-initial-about:blank){#beginning-navigation:is-initial-about:blank}
      is false; and

    - `url`{.variable}\'s
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#beginning-navigation:concept-url-scheme-3
      x-internal="concept-url-scheme"} is a [fetch
      scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#beginning-navigation:fetch-scheme
      x-internal="fetch-scheme"},

    then:

    1.  Let `navigation`{.variable} be `navigable`{.variable}\'s [active
        window](document-sequences.html#nav-window){#beginning-navigation:nav-window-3}\'s
        [navigation
        API](nav-history-apis.html#window-navigation-api){#beginning-navigation:window-navigation-api}.

    2.  Let `entryListForFiring`{.variable} be
        `formDataEntryList`{.variable} if `documentResource`{.variable}
        is a [POST
        resource](#post-resource){#beginning-navigation:post-resource-2};
        otherwise, null.

    3.  Let `navigationAPIStateForFiring`{.variable} be
        `navigationAPIState`{.variable} if
        `navigationAPIState`{.variable} is not null; otherwise,
        [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#beginning-navigation:structuredserializeforstorage}(undefined).

    4.  Let `continue`{.variable} be the result of [firing a
        push/replace/reload `navigate`
        event](nav-history-apis.html#fire-a-push/replace/reload-navigate-event){#beginning-navigation:fire-a-push/replace/reload-navigate-event}
        at `navigation`{.variable} with
        *[navigationType](nav-history-apis.html#fire-navigate-prr-navigationtype)*
        set to `historyHandling`{.variable},
        *[isSameDocument](nav-history-apis.html#fire-navigate-prr-issamedocument)*
        set to false,
        *[userInvolvement](nav-history-apis.html#fire-navigate-prr-userinvolvement)*
        set to `userInvolvement`{.variable},
        *[sourceElement](nav-history-apis.html#fire-navigate-prr-sourceelement)*
        set to `sourceElement`{.variable},
        *[formDataEntryList](nav-history-apis.html#fire-navigate-prr-formdataentrylist)*
        set to `entryListForFiring`{.variable},
        *[destinationURL](nav-history-apis.html#fire-navigate-prr-destinationurl)*
        set to `url`{.variable}, and
        *[navigationAPIState](nav-history-apis.html#fire-navigate-prr-navigationapistate)*
        set to `navigationAPIStateForFiring`{.variable}.

    5.  If `continue`{.variable} is false, then return.

    It is possible for navigations with `userInvolvement`{.variable} of
    \"[`browser UI`{#beginning-navigation:uni-browser-ui-4}](#uni-browser-ui)\"
    or initiated by a [cross
    origin-domain](browsers.html#same-origin-domain){#beginning-navigation:same-origin-domain-2}
    `sourceDocument`{.variable} to fire
    [`navigate`{#beginning-navigation:event-navigate}](indices.html#event-navigate)
    events, if they go through the earlier [navigate to a
    fragment](#navigate-fragid){#beginning-navigation:navigate-fragid-3}
    path.

22. [In
    parallel](infrastructure.html#in-parallel){#beginning-navigation:in-parallel},
    run these steps:

    1.  Let `unloadPromptCanceled`{.variable} be the result of [checking
        if unloading is
        canceled](#checking-if-unloading-is-canceled){#beginning-navigation:checking-if-unloading-is-canceled}
        for `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#beginning-navigation:nav-document-10}\'s
        [inclusive descendant
        navigables](document-sequences.html#inclusive-descendant-navigables){#beginning-navigation:inclusive-descendant-navigables}.

    2.  If `unloadPromptCanceled`{.variable} is true, or
        `navigable`{.variable}\'s [ongoing
        navigation](#ongoing-navigation){#beginning-navigation:ongoing-navigation-3}
        is no longer `navigationId`{.variable}, then:

        1.  Invoke [WebDriver BiDi navigation
            failed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-failed){#beginning-navigation:webdriver-bidi-navigation-failed-3
            x-internal="webdriver-bidi-navigation-failed"} with
            `navigable`{.variable} and a new [WebDriver BiDi navigation
            status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#beginning-navigation:webdriver-bidi-navigation-status-4
            x-internal="webdriver-bidi-navigation-status"} whose
            [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#beginning-navigation:navigation-status-id-4
            x-internal="navigation-status-id"} is
            `navigationId`{.variable},
            [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#beginning-navigation:navigation-status-status-4
            x-internal="navigation-status-status"} is
            \"[`canceled`{#beginning-navigation:navigation-status-canceled-3}](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled){x-internal="navigation-status-canceled"}\",
            and
            [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#beginning-navigation:navigation-status-url-4
            x-internal="navigation-status-url"} is `url`{.variable}.

        2.  Abort these steps.

    3.  [Queue a global
        task](webappapis.html#queue-a-global-task){#beginning-navigation:queue-a-global-task-3}
        on the [navigation and traversal task
        source](webappapis.html#navigation-and-traversal-task-source){#beginning-navigation:navigation-and-traversal-task-source-3}
        given `navigable`{.variable}\'s [active
        window](document-sequences.html#nav-window){#beginning-navigation:nav-window-4}
        to [abort a document and its
        descendants](document-lifecycle.html#abort-a-document-and-its-descendants){#beginning-navigation:abort-a-document-and-its-descendants}
        given `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#beginning-navigation:nav-document-11}.

    4.  ::: {#navigation-create-document-state}
        Let `documentState`{.variable} be a new [document
        state](#document-state-2){#beginning-navigation:document-state-2}
        with

        [request referrer policy](#document-state-request-referrer-policy){#beginning-navigation:document-state-request-referrer-policy}
        :   `referrerPolicy`{.variable}

        [initiator origin](#document-state-initiator-origin){#beginning-navigation:document-state-initiator-origin}
        :   `initiatorOriginSnapshot`{.variable}

        [resource](#document-state-resource){#beginning-navigation:document-state-resource}
        :   `documentResource`{.variable}

        [navigable target name](#document-state-nav-target-name){#beginning-navigation:document-state-nav-target-name}
        :   `navigable`{.variable}\'s [target
            name](document-sequences.html#nav-target){#beginning-navigation:nav-target}

        The [navigable target
        name](#document-state-nav-target-name){#beginning-navigation:document-state-nav-target-name-2}
        can get cleared under various conditions later in the navigation
        process, before the document state is finalized.
        :::

    5.  If `url`{.variable} [matches
        `about:blank`](urls-and-fetching.html#matches-about:blank){#beginning-navigation:matches-about:blank}
        or is
        [`about:srcdoc`{#beginning-navigation:about:srcdoc}](urls-and-fetching.html#about:srcdoc),
        then:

        1.  Set `documentState`{.variable}\'s
            [origin](#document-state-origin){#beginning-navigation:document-state-origin}
            to `initiatorOriginSnapshot`{.variable}.

        2.  Set `documentState`{.variable}\'s [about base
            URL](#document-state-about-base-url){#beginning-navigation:document-state-about-base-url}
            to `initiatorBaseURLSnapshot`{.variable}.

    6.  Let `historyEntry`{.variable} be a new [session history
        entry](#session-history-entry){#beginning-navigation:session-history-entry},
        with its [URL](#she-url){#beginning-navigation:she-url-2} set to
        `url`{.variable} and its [document
        state](#she-document-state){#beginning-navigation:she-document-state}
        set to `documentState`{.variable}.

    7.  Let `navigationParams`{.variable} be null.

    8.  If `response`{.variable} is non-null:

        The [navigate](#navigate){#beginning-navigation:navigate}
        algorithm is only supplied with a
        [response](https://fetch.spec.whatwg.org/#concept-response){#beginning-navigation:concept-response-2
        x-internal="concept-response"} as part of the
        [`object`{#beginning-navigation:the-object-element}](iframe-embed-object.html#the-object-element)
        and
        [`embed`{#beginning-navigation:the-embed-element}](iframe-embed-object.html#the-embed-element)
        processing models, or for processing parts of
        [`multipart/x-mixed-replace`
        responses](document-lifecycle.html#navigate-multipart-x-mixed-replace){#beginning-navigation:navigate-multipart-x-mixed-replace}
        after the initial response.

        1.  Let `sourcePolicyContainer`{.variable} be a
            [clone](browsers.html#clone-a-policy-container){#beginning-navigation:clone-a-policy-container}
            of the `sourceDocument`{.variable}\'s [policy
            container](dom.html#concept-document-policy-container){#beginning-navigation:concept-document-policy-container},
            if `sourceDocument`{.variable} is not null; otherwise null.

        2.  Let `policyContainer`{.variable} be the result of
            [determining navigation params policy
            container](browsers.html#determining-navigation-params-policy-container){#beginning-navigation:determining-navigation-params-policy-container}
            given `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#beginning-navigation:concept-response-url
            x-internal="concept-response-url"}, null,
            `sourcePolicyContainer`{.variable},
            `navigable`{.variable}\'s [container
            document](document-sequences.html#nav-container-document){#beginning-navigation:nav-container-document}\'s
            [policy
            container](dom.html#concept-document-policy-container){#beginning-navigation:concept-document-policy-container-2},
            and null.

        3.  Let `finalSandboxFlags`{.variable} be the
            [union](https://infra.spec.whatwg.org/#set-union){#beginning-navigation:set-union
            x-internal="set-union"} of
            `targetSnapshotParams`{.variable}\'s [sandboxing
            flags](#target-snapshot-params-sandbox){#beginning-navigation:target-snapshot-params-sandbox}
            and `policyContainer`{.variable}\'s [CSP
            list](browsers.html#policy-container-csp-list){#beginning-navigation:policy-container-csp-list}\'s
            [CSP-derived sandboxing
            flags](browsers.html#csp-derived-sandboxing-flags){#beginning-navigation:csp-derived-sandboxing-flags}.

        4.  Let `responseOrigin`{.variable} be the result of
            [determining the
            origin](document-sequences.html#determining-the-origin){#beginning-navigation:determining-the-origin}
            given `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#beginning-navigation:concept-response-url-2
            x-internal="concept-response-url"},
            `finalSandboxFlags`{.variable}, and
            `documentState`{.variable}\'s [initiator
            origin](#document-state-initiator-origin){#beginning-navigation:document-state-initiator-origin-2}.

        5.  Let `coop`{.variable} be a new [opener
            policy](browsers.html#cross-origin-opener-policy){#beginning-navigation:cross-origin-opener-policy}.

        6.  Let `coopEnforcementResult`{.variable} be a new [opener
            policy enforcement
            result](browsers.html#coop-enforcement-result){#beginning-navigation:coop-enforcement-result}
            with

            [url](browsers.html#coop-enforcement-url){#beginning-navigation:coop-enforcement-url}
            :   `response`{.variable}\'s
                [URL](https://fetch.spec.whatwg.org/#concept-response-url){#beginning-navigation:concept-response-url-3
                x-internal="concept-response-url"}

            [origin](browsers.html#coop-enforcement-origin){#beginning-navigation:coop-enforcement-origin}
            :   `responseOrigin`{.variable}

            [opener policy](browsers.html#coop-enforcement-coop){#beginning-navigation:coop-enforcement-coop}
            :   `coop`{.variable}

        7.  Set `navigationParams`{.variable} to a new [navigation
            params](#navigation-params){#beginning-navigation:navigation-params},
            with

            [id](#navigation-params-id){#beginning-navigation:navigation-params-id}
            :   `navigationId`{.variable}

            [navigable](#navigation-params-navigable){#beginning-navigation:navigation-params-navigable}
            :   `navigable`{.variable}

            [request](#navigation-params-request){#beginning-navigation:navigation-params-request}
            :   null

            [response](#navigation-params-response){#beginning-navigation:navigation-params-response}
            :   `response`{.variable}

            [fetch controller](#navigation-params-fetch-controller){#beginning-navigation:navigation-params-fetch-controller}
            :   null

            [commit early hints](#navigation-params-commit-early-hints){#beginning-navigation:navigation-params-commit-early-hints}
            :   null

            [COOP enforcement result](#navigation-params-coop-enforcement-result){#beginning-navigation:navigation-params-coop-enforcement-result}
            :   `coopEnforcementResult`{.variable}

            [reserved environment](#navigation-params-reserved-environment){#beginning-navigation:navigation-params-reserved-environment}
            :   null

            [origin](#navigation-params-origin){#beginning-navigation:navigation-params-origin}
            :   `responseOrigin`{.variable}

            [policy container](#navigation-params-policy-container){#beginning-navigation:navigation-params-policy-container}
            :   `policyContainer`{.variable}

            [final sandboxing flag set](#navigation-params-sandboxing){#beginning-navigation:navigation-params-sandboxing}
            :   `finalSandboxFlags`{.variable}

            [opener policy](#navigation-params-coop){#beginning-navigation:navigation-params-coop}
            :   `coop`{.variable}

            [navigation timing type](#navigation-params-nav-timing-type){#beginning-navigation:navigation-params-nav-timing-type}
            :   \"[`navigate`{#beginning-navigation:dom-navigationtimingtype-navigate}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-navigate){x-internal="dom-navigationtimingtype-navigate"}\"

            [about base URL](#navigation-params-about-base-url){#beginning-navigation:navigation-params-about-base-url}
            :   `documentState`{.variable}\'s [about base
                URL](#document-state-about-base-url){#beginning-navigation:document-state-about-base-url-2}

            [user involvement](#navigation-params-user-involvement){#beginning-navigation:navigation-params-user-involvement}
            :   `userInvolvement`{.variable}

    9.  [Attempt to populate the history entry\'s
        document](#attempt-to-populate-the-history-entry's-document){#beginning-navigation:attempt-to-populate-the-history-entry's-document}
        for `historyEntry`{.variable}, given `navigable`{.variable},
        \"[`navigate`{#beginning-navigation:dom-navigationtimingtype-navigate-2}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-navigate){x-internal="dom-navigationtimingtype-navigate"}\",
        `sourceSnapshotParams`{.variable},
        `targetSnapshotParams`{.variable}, `userInvolvement`{.variable},
        `navigationId`{.variable}, `navigationParams`{.variable},
        `cspNavigationType`{.variable}, with
        *[allowPOST](#attempt-to-populate-allow-post)* set to true and
        *[completionSteps](#attempt-to-populate-completion-steps)* set
        to the following step:

        1.  [Append session history traversal
            steps](#tn-append-session-history-traversal-steps){#beginning-navigation:tn-append-session-history-traversal-steps}
            to `navigable`{.variable}\'s
            [traversable](document-sequences.html#nav-traversable){#beginning-navigation:nav-traversable}
            to [finalize a cross-document
            navigation](#finalize-a-cross-document-navigation){#beginning-navigation:finalize-a-cross-document-navigation}
            given `navigable`{.variable}, `historyHandling`{.variable},
            `userInvolvement`{.variable}, and `historyEntry`{.variable}.

##### [7.4.2.3]{.secno} Ending navigation[](#ending-navigation){.self-link}

Although the usual cross-document navigation case will first foray into
[populating a session history
entry](#populating-a-session-history-entry) with a
[`Document`{#ending-navigation:document}](dom.html#document), all
navigations that don\'t get aborted will ultimately end up calling into
one of the below algorithms.

###### [7.4.2.3.1]{.secno} The usual cross-document navigation case[](#the-usual-cross-document-navigation-case){.self-link}

To [finalize a cross-document
navigation]{#finalize-a-cross-document-navigation .dfn} given a
[navigable](document-sequences.html#navigable){#the-usual-cross-document-navigation-case:navigable}
`navigable`{.variable}, a [history handling
behavior](#history-handling-behavior){#the-usual-cross-document-navigation-case:history-handling-behavior}
`historyHandling`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#the-usual-cross-document-navigation-case:user-navigation-involvement}
`userInvolvement`{.variable}, and a [session history
entry](#session-history-entry){#the-usual-cross-document-navigation-case:session-history-entry}
`historyEntry`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-usual-cross-document-navigation-case:assert
    x-internal="assert"}: this is running on `navigable`{.variable}\'s
    [traversable
    navigable\'s](document-sequences.html#nav-traversable){#the-usual-cross-document-navigation-case:nav-traversable}
    [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#the-usual-cross-document-navigation-case:tn-session-history-traversal-queue}.

2.  Set `navigable`{.variable}\'s [is delaying `load`
    events](document-sequences.html#delaying-load-events-mode){#the-usual-cross-document-navigation-case:delaying-load-events-mode}
    to false.

3.  If `historyEntry`{.variable}\'s
    [document](#she-document){#the-usual-cross-document-navigation-case:she-document}
    is null, then return.

    This means that [attempting to populate the history entry\'s
    document](#attempt-to-populate-the-history-entry's-document){#the-usual-cross-document-navigation-case:attempt-to-populate-the-history-entry's-document}
    ended up not creating a document, as a result of e.g., the
    navigation being canceled by a subsequent navigation, a 204 No
    Content response, etc.

4.  ::: {#resetBCName}
    If all of the following are true:

    - `navigable`{.variable}\'s
      [parent](document-sequences.html#nav-parent){#the-usual-cross-document-navigation-case:nav-parent}
      is null;

    - `historyEntry`{.variable}\'s
      [document](#she-document){#the-usual-cross-document-navigation-case:she-document-2}\'s
      [browsing
      context](document-sequences.html#concept-document-bc){#the-usual-cross-document-navigation-case:concept-document-bc}
      is not an [auxiliary browsing
      context](document-sequences.html#auxiliary-browsing-context){#the-usual-cross-document-navigation-case:auxiliary-browsing-context}
      whose [opener browsing
      context](document-sequences.html#opener-browsing-context){#the-usual-cross-document-navigation-case:opener-browsing-context}
      is non-null; and

    - `historyEntry`{.variable}\'s
      [document](#she-document){#the-usual-cross-document-navigation-case:she-document-3}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-usual-cross-document-navigation-case:concept-document-origin
      x-internal="concept-document-origin"} is not
      `navigable`{.variable}\'s [active
      document](document-sequences.html#nav-document){#the-usual-cross-document-navigation-case:nav-document}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-usual-cross-document-navigation-case:concept-document-origin-2
      x-internal="concept-document-origin"},

    then set `historyEntry`{.variable}\'s [document
    state](#she-document-state){#the-usual-cross-document-navigation-case:she-document-state}\'s
    [navigable target
    name](#document-state-nav-target-name){#the-usual-cross-document-navigation-case:document-state-nav-target-name}
    to the empty string.
    :::

5.  Let `entryToReplace`{.variable} be `navigable`{.variable}\'s [active
    session history
    entry](document-sequences.html#nav-active-history-entry){#the-usual-cross-document-navigation-case:nav-active-history-entry}
    if `historyHandling`{.variable} is
    \"[`replace`{#the-usual-cross-document-navigation-case:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\",
    otherwise null.

6.  Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#the-usual-cross-document-navigation-case:nav-traversable-2}.

7.  Let `targetStep`{.variable} be null.

8.  Let `targetEntries`{.variable} be the result of [getting session
    history
    entries](#getting-session-history-entries){#the-usual-cross-document-navigation-case:getting-session-history-entries}
    for `navigable`{.variable}.

9.  If `entryToReplace`{.variable} is null, then:

    1.  [Clear the forward session
        history](#clear-the-forward-session-history){#the-usual-cross-document-navigation-case:clear-the-forward-session-history}
        of `traversable`{.variable}.

    2.  Set `targetStep`{.variable} to `traversable`{.variable}\'s
        [current session history
        step](document-sequences.html#tn-current-session-history-step){#the-usual-cross-document-navigation-case:tn-current-session-history-step} +
        1.

    3.  Set `historyEntry`{.variable}\'s
        [step](#she-step){#the-usual-cross-document-navigation-case:she-step}
        to `targetStep`{.variable}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#the-usual-cross-document-navigation-case:list-append
        x-internal="list-append"} `historyEntry`{.variable} to
        `targetEntries`{.variable}.

    Otherwise:

    1.  [Replace](https://infra.spec.whatwg.org/#list-replace){#the-usual-cross-document-navigation-case:list-replace
        x-internal="list-replace"} `entryToReplace`{.variable} with
        `historyEntry`{.variable} in `targetEntries`{.variable}.

    2.  Set `historyEntry`{.variable}\'s
        [step](#she-step){#the-usual-cross-document-navigation-case:she-step-2}
        to `entryToReplace`{.variable}\'s
        [step](#she-step){#the-usual-cross-document-navigation-case:she-step-3}.

    3.  If `historyEntry`{.variable}\'s [document
        state](#she-document-state){#the-usual-cross-document-navigation-case:she-document-state-2}\'s
        [origin](#document-state-origin){#the-usual-cross-document-navigation-case:document-state-origin}
        is [same
        origin](browsers.html#same-origin){#the-usual-cross-document-navigation-case:same-origin}
        with `entryToReplace`{.variable}\'s [document
        state](#she-document-state){#the-usual-cross-document-navigation-case:she-document-state-3}\'s
        [origin](#document-state-origin){#the-usual-cross-document-navigation-case:document-state-origin-2},
        then set `historyEntry`{.variable}\'s [navigation API
        key](#she-navigation-api-key){#the-usual-cross-document-navigation-case:she-navigation-api-key}
        to `entryToReplace`{.variable}\'s [navigation API
        key](#she-navigation-api-key){#the-usual-cross-document-navigation-case:she-navigation-api-key-2}.

    4.  Set `targetStep`{.variable} to `traversable`{.variable}\'s
        [current session history
        step](document-sequences.html#tn-current-session-history-step){#the-usual-cross-document-navigation-case:tn-current-session-history-step-2}.

10. [Apply the push/replace history
    step](#apply-the-push/replace-history-step){#the-usual-cross-document-navigation-case:apply-the-push/replace-history-step}
    `targetStep`{.variable} to `traversable`{.variable} given
    `historyHandling`{.variable} and `userInvolvement`{.variable}.

###### [7.4.2.3.2]{.secno} [The `javascript:` URL special case]{.dfn}[](#the-javascript:-url-special-case){.self-link} {#the-javascript:-url-special-case}

[`javascript:`{#the-javascript:-url-special-case:the-javascript:-url-special-case}](#the-javascript:-url-special-case)
URLs have a [dedicated
label](https://github.com/whatwg/html/labels/topic%3A%20javascript%3A%20URLs)
on the issue tracker documenting various problems with their
specification.

To [navigate to a `javascript:` URL]{#navigate-to-a-javascript:-url
.dfn}, given a
[navigable](document-sequences.html#navigable){#the-javascript:-url-special-case:navigable}
`targetNavigable`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#the-javascript:-url-special-case:url
x-internal="url"} `url`{.variable}, a [history handling
behavior](#history-handling-behavior){#the-javascript:-url-special-case:history-handling-behavior}
`historyHandling`{.variable}, a [source snapshot
params](#source-snapshot-params){#the-javascript:-url-special-case:source-snapshot-params}
`sourceSnapshotParams`{.variable}, an
[origin](browsers.html#concept-origin){#the-javascript:-url-special-case:concept-origin}
`initiatorOrigin`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#the-javascript:-url-special-case:user-navigation-involvement}
`userInvolvement`{.variable}, a string `cspNavigationType`{.variable},
and a boolean `initialInsertion`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-javascript:-url-special-case:assert
    x-internal="assert"}: `historyHandling`{.variable} is
    \"[`replace`{#the-javascript:-url-special-case:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\".

2.  [Set the ongoing
    navigation](#set-the-ongoing-navigation){#the-javascript:-url-special-case:set-the-ongoing-navigation}
    for `targetNavigable`{.variable} to null.

3.  If `initiatorOrigin`{.variable} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-javascript:-url-special-case:same-origin-domain}
    with `targetNavigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-javascript:-url-special-case:concept-document-origin
    x-internal="concept-document-origin"}, then return.

4.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#the-javascript:-url-special-case:concept-request
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#the-javascript:-url-special-case:concept-request-url
    x-internal="concept-request-url"} is `url`{.variable} and whose
    [policy
    container](https://fetch.spec.whatwg.org/#concept-request-policy-container){#the-javascript:-url-special-case:concept-request-policy-container
    x-internal="concept-request-policy-container"} is
    `sourceSnapshotParams`{.variable}\'s [source policy
    container](#source-snapshot-params-policy-container){#the-javascript:-url-special-case:source-snapshot-params-policy-container}.

    This is a synthetic
    [request](https://fetch.spec.whatwg.org/#concept-request){#the-javascript:-url-special-case:concept-request-2
    x-internal="concept-request"} solely for plumbing into the next
    step. It will never hit the network.

5.  If the result of [should navigation request of type be blocked by
    Content Security
    Policy?](https://w3c.github.io/webappsec-csp/#should-block-navigation-request){#the-javascript:-url-special-case:should-navigation-request-of-type-be-blocked-by-content-security-policy
    x-internal="should-navigation-request-of-type-be-blocked-by-content-security-policy"}
    given `request`{.variable} and `cspNavigationType`{.variable} is
    \"`Blocked`\", then return. [\[CSP\]](references.html#refsCSP)

6.  Let `newDocument`{.variable} be the result of [evaluating a
    `javascript:`
    URL](#evaluate-a-javascript:-url){#the-javascript:-url-special-case:evaluate-a-javascript:-url}
    given `targetNavigable`{.variable}, `url`{.variable},
    `initiatorOrigin`{.variable}, and `userInvolvement`{.variable}.

7.  If `newDocument`{.variable} is null:

    1.  If `initialInsertion`{.variable} is true and
        `targetNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-2}\'s
        [is initial
        `about:blank`](dom.html#is-initial-about:blank){#the-javascript:-url-special-case:is-initial-about:blank}
        is true, then run the [iframe load event
        steps](iframe-embed-object.html#iframe-load-event-steps){#the-javascript:-url-special-case:iframe-load-event-steps}
        given `targetNavigable`{.variable}\'s
        [container](document-sequences.html#nav-container){#the-javascript:-url-special-case:nav-container}.

    2.  Return.

    In this case, some JavaScript code was executed, but no new
    [`Document`{#the-javascript:-url-special-case:document}](dom.html#document)
    was created, so we will not perform a navigation.

8.  [Assert](https://infra.spec.whatwg.org/#assert){#the-javascript:-url-special-case:assert-2
    x-internal="assert"}: `initiatorOrigin`{.variable} is
    `newDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-javascript:-url-special-case:concept-document-origin-2
    x-internal="concept-document-origin"}.

9.  Let `entryToReplace`{.variable} be `targetNavigable`{.variable}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#the-javascript:-url-special-case:nav-active-history-entry}.

10. Let `oldDocState`{.variable} be `entryToReplace`{.variable}\'s
    [document
    state](#she-document-state){#the-javascript:-url-special-case:she-document-state}.

11. Let `documentState`{.variable} be a new [document
    state](#document-state-2){#the-javascript:-url-special-case:document-state-2}
    with

    [document](#document-state-document){#the-javascript:-url-special-case:document-state-document}
    :   `newDocument`{.variable}

    [history policy container](#document-state-history-policy-container){#the-javascript:-url-special-case:document-state-history-policy-container}
    :   a
        [clone](browsers.html#clone-a-policy-container){#the-javascript:-url-special-case:clone-a-policy-container}
        of the `oldDocState`{.variable}\'s [history policy
        container](#document-state-history-policy-container){#the-javascript:-url-special-case:document-state-history-policy-container-2}
        if it is non-null; null otherwise

    [request referrer](#document-state-request-referrer){#the-javascript:-url-special-case:document-state-request-referrer}
    :   `oldDocState`{.variable}\'s [request
        referrer](#document-state-request-referrer){#the-javascript:-url-special-case:document-state-request-referrer-2}

    [request referrer policy](#document-state-request-referrer-policy){#the-javascript:-url-special-case:document-state-request-referrer-policy}
    :   `oldDocState`{.variable}\'s [request referrer
        policy](#document-state-request-referrer-policy){#the-javascript:-url-special-case:document-state-request-referrer-policy-2}
        [or should this be the `referrerPolicy`{.variable} that was
        passed to
        [navigate](#navigate){#the-javascript:-url-special-case:navigate}?]{.XXX}

    [initiator origin](#document-state-initiator-origin){#the-javascript:-url-special-case:document-state-initiator-origin}
    :   `initiatorOrigin`{.variable}

    [origin](#document-state-origin){#the-javascript:-url-special-case:document-state-origin}
    :   `initiatorOrigin`{.variable}

    [about base URL](#document-state-about-base-url){#the-javascript:-url-special-case:document-state-about-base-url}
    :   `oldDocState`{.variable}\'s [about base
        URL](#document-state-about-base-url){#the-javascript:-url-special-case:document-state-about-base-url-2}

    [resource](#document-state-resource){#the-javascript:-url-special-case:document-state-resource}
    :   null

    [ever populated](#document-state-ever-populated){#the-javascript:-url-special-case:document-state-ever-populated}
    :   true

    [navigable target name](#document-state-nav-target-name){#the-javascript:-url-special-case:document-state-nav-target-name}
    :   `oldDocState`{.variable}\'s [navigable target
        name](#document-state-nav-target-name){#the-javascript:-url-special-case:document-state-nav-target-name-2}

12. Let `historyEntry`{.variable} be a new [session history
    entry](#session-history-entry){#the-javascript:-url-special-case:session-history-entry},
    with

    [URL](#she-url){#the-javascript:-url-special-case:she-url}
    :   `entryToReplace`{.variable}\'s
        [URL](#she-url){#the-javascript:-url-special-case:she-url-2}

    [document state](#she-document-state){#the-javascript:-url-special-case:she-document-state-2}
    :   `documentState`{.variable}

    For the
    [URL](#she-url){#the-javascript:-url-special-case:she-url-3}, we do
    *not* use `url`{.variable}, i.e. the actual
    [`javascript:`{#the-javascript:-url-special-case:the-javascript:-url-special-case-2}](#the-javascript:-url-special-case)
    URL that the
    [navigate](#navigate){#the-javascript:-url-special-case:navigate-2}
    algorithm was called with. This means
    [`javascript:`{#the-javascript:-url-special-case:the-javascript:-url-special-case-3}](#the-javascript:-url-special-case)
    URLs are never stored in session history, and so can never be
    traversed to.

13. [Append session history traversal
    steps](#tn-append-session-history-traversal-steps){#the-javascript:-url-special-case:tn-append-session-history-traversal-steps}
    to `targetNavigable`{.variable}\'s
    [traversable](document-sequences.html#nav-traversable){#the-javascript:-url-special-case:nav-traversable}
    to [finalize a cross-document
    navigation](#finalize-a-cross-document-navigation){#the-javascript:-url-special-case:finalize-a-cross-document-navigation}
    with `targetNavigable`{.variable}, `historyHandling`{.variable},
    `userInvolvement`{.variable}, and `historyEntry`{.variable}.

To [evaluate a `javascript:` URL]{#evaluate-a-javascript:-url .dfn}
given a
[navigable](document-sequences.html#navigable){#the-javascript:-url-special-case:navigable-2}
`targetNavigable`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#the-javascript:-url-special-case:url-2
x-internal="url"} `url`{.variable}, an
[origin](browsers.html#concept-origin){#the-javascript:-url-special-case:concept-origin-2}
`newDocumentOrigin`{.variable}, and a [user navigation
involvement](#user-navigation-involvement){#the-javascript:-url-special-case:user-navigation-involvement-2}
`userInvolvement`{.variable}:

1.  Let `urlString`{.variable} be the result of running the [URL
    serializer](https://url.spec.whatwg.org/#concept-url-serializer){#the-javascript:-url-special-case:concept-url-serializer
    x-internal="concept-url-serializer"} on `url`{.variable}.

2.  Let `encodedScriptSource`{.variable} be the result of removing the
    leading \"`javascript:`\" from `urlString`{.variable}.

3.  Let `scriptSource`{.variable} be the [UTF-8
    decoding](https://encoding.spec.whatwg.org/#utf-8-decode){#the-javascript:-url-special-case:utf-8-decode
    x-internal="utf-8-decode"} of the
    [percent-decoding](https://url.spec.whatwg.org/#string-percent-decode){#the-javascript:-url-special-case:percent-decode
    x-internal="percent-decode"} of `encodedScriptSource`{.variable}.

4.  Let `settings`{.variable} be `targetNavigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-3}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#the-javascript:-url-special-case:relevant-settings-object}.

5.  Let `baseURL`{.variable} be `settings`{.variable}\'s [API base
    URL](webappapis.html#api-base-url){#the-javascript:-url-special-case:api-base-url}.

6.  Let `script`{.variable} be the result of [creating a classic
    script](webappapis.html#creating-a-classic-script){#the-javascript:-url-special-case:creating-a-classic-script}
    given `scriptSource`{.variable}, `settings`{.variable},
    `baseURL`{.variable}, and the [default script fetch
    options](webappapis.html#default-script-fetch-options){#the-javascript:-url-special-case:default-script-fetch-options}.

7.  Let `evaluationStatus`{.variable} be the result of [running the
    classic
    script](webappapis.html#run-a-classic-script){#the-javascript:-url-special-case:run-a-classic-script}
    `script`{.variable}.

8.  Let `result`{.variable} be null.

9.  If `evaluationStatus`{.variable} is a normal completion, and
    `evaluationStatus`{.variable}.\[\[Value\]\] is a String, then set
    `result`{.variable} to `evaluationStatus`{.variable}.\[\[Value\]\].

10. Otherwise, return null.

11. Let `response`{.variable} be a new
    [response](https://fetch.spec.whatwg.org/#concept-response){#the-javascript:-url-special-case:concept-response
    x-internal="concept-response"} with

    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#the-javascript:-url-special-case:concept-response-url x-internal="concept-response-url"}
    :   `targetNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-4}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-javascript:-url-special-case:the-document's-address
        x-internal="the-document's-address"}

    [header list](https://fetch.spec.whatwg.org/#concept-response-header-list){#the-javascript:-url-special-case:concept-response-header-list x-internal="concept-response-header-list"}
    :   «
        (\`[`Content-Type`{#the-javascript:-url-special-case:content-type}](urls-and-fetching.html#content-type)\`,
        \``text/html;charset=utf-8`\`) »

    [body](https://fetch.spec.whatwg.org/#concept-response-body){#the-javascript:-url-special-case:concept-response-body x-internal="concept-response-body"}
    :   the [UTF-8
        encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#the-javascript:-url-special-case:utf-8-encode
        x-internal="utf-8-encode"} of `result`{.variable}, [as a
        body](https://fetch.spec.whatwg.org/#byte-sequence-as-a-body){#the-javascript:-url-special-case:as-a-body
        x-internal="as-a-body"}

    The encoding to UTF-8 means that unpaired
    [surrogates](https://infra.spec.whatwg.org/#surrogate){#the-javascript:-url-special-case:surrogate
    x-internal="surrogate"} will not roundtrip, once the HTML parser
    decodes the response body.

12. Let `policyContainer`{.variable} be `targetNavigable`{.variable}\'s
    [active
    document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-5}\'s
    [policy
    container](dom.html#concept-document-policy-container){#the-javascript:-url-special-case:concept-document-policy-container}.

13. Let `finalSandboxFlags`{.variable} be
    `policyContainer`{.variable}\'s [CSP
    list](browsers.html#policy-container-csp-list){#the-javascript:-url-special-case:policy-container-csp-list}\'s
    [CSP-derived sandboxing
    flags](browsers.html#csp-derived-sandboxing-flags){#the-javascript:-url-special-case:csp-derived-sandboxing-flags}.

14. Let `coop`{.variable} be `targetNavigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-6}\'s
    [opener
    policy](dom.html#concept-document-coop){#the-javascript:-url-special-case:concept-document-coop}.

15. Let `coopEnforcementResult`{.variable} be a new [opener policy
    enforcement
    result](browsers.html#coop-enforcement-result){#the-javascript:-url-special-case:coop-enforcement-result}
    with

    [url](browsers.html#coop-enforcement-url){#the-javascript:-url-special-case:coop-enforcement-url}
    :   `url`{.variable}

    [origin](browsers.html#coop-enforcement-origin){#the-javascript:-url-special-case:coop-enforcement-origin}
    :   `newDocumentOrigin`{.variable}

    [opener policy](browsers.html#coop-enforcement-coop){#the-javascript:-url-special-case:coop-enforcement-coop}
    :   `coop`{.variable}

16. Let `navigationParams`{.variable} be a new [navigation
    params](#navigation-params){#the-javascript:-url-special-case:navigation-params},
    with

    [id](#navigation-params-id){#the-javascript:-url-special-case:navigation-params-id}
    :   `navigationId`{.variable}

    [navigable](#navigation-params-navigable){#the-javascript:-url-special-case:navigation-params-navigable}
    :   `targetNavigable`{.variable}

    [request](#navigation-params-request){#the-javascript:-url-special-case:navigation-params-request}
    :   null [this will cause the referrer of the resulting
        [`Document`{#the-javascript:-url-special-case:document-2}](dom.html#document)
        to be null; is that correct?]{.XXX}

    [response](#navigation-params-response){#the-javascript:-url-special-case:navigation-params-response}
    :   `response`{.variable}

    [fetch controller](#navigation-params-fetch-controller){#the-javascript:-url-special-case:navigation-params-fetch-controller}
    :   null

    [commit early hints](#navigation-params-commit-early-hints){#the-javascript:-url-special-case:navigation-params-commit-early-hints}
    :   null

    [COOP enforcement result](#navigation-params-coop-enforcement-result){#the-javascript:-url-special-case:navigation-params-coop-enforcement-result}
    :   `coopEnforcementResult`{.variable}

    [reserved environment](#navigation-params-reserved-environment){#the-javascript:-url-special-case:navigation-params-reserved-environment}
    :   null

    [origin](#navigation-params-origin){#the-javascript:-url-special-case:navigation-params-origin}
    :   `newDocumentOrigin`{.variable}

    [policy container](#navigation-params-policy-container){#the-javascript:-url-special-case:navigation-params-policy-container}
    :   `policyContainer`{.variable}

    [final sandboxing flag set](#navigation-params-sandboxing){#the-javascript:-url-special-case:navigation-params-sandboxing}
    :   `finalSandboxFlags`{.variable}

    [opener policy](#navigation-params-coop){#the-javascript:-url-special-case:navigation-params-coop}
    :   `coop`{.variable}

    [navigation timing type](#navigation-params-nav-timing-type){#the-javascript:-url-special-case:navigation-params-nav-timing-type}
    :   \"[`navigate`{#the-javascript:-url-special-case:dom-navigationtimingtype-navigate}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-navigate){x-internal="dom-navigationtimingtype-navigate"}\"

    [about base URL](#navigation-params-about-base-url){#the-javascript:-url-special-case:navigation-params-about-base-url}
    :   `targetNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#the-javascript:-url-special-case:nav-document-7}\'s
        [about base
        URL](dom.html#concept-document-about-base-url){#the-javascript:-url-special-case:concept-document-about-base-url}

    [user involvement](#navigation-params-user-involvement){#the-javascript:-url-special-case:navigation-params-user-involvement}
    :   `userInvolvement`{.variable}

17. Return the result of [loading an HTML
    document](document-lifecycle.html#navigate-html){#the-javascript:-url-special-case:navigate-html}
    given `navigationParams`{.variable}.

###### [7.4.2.3.3]{.secno} Fragment navigations[](#scroll-to-fragid){.self-link} {#scroll-to-fragid}

To [navigate to a fragment]{#navigate-fragid .dfn} given a
[navigable](document-sequences.html#navigable){#scroll-to-fragid:navigable}
`navigable`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#scroll-to-fragid:url
x-internal="url"} `url`{.variable}, a [history handling
behavior](#history-handling-behavior){#scroll-to-fragid:history-handling-behavior}
`historyHandling`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#scroll-to-fragid:user-navigation-involvement}
`userInvolvement`{.variable}, an
[`Element`{#scroll-to-fragid:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}-or-null
`sourceElement`{.variable}, a [serialized
state](#serialized-state){#scroll-to-fragid:serialized-state}-or-null
`navigationAPIState`{.variable}, and a [navigation
ID](#navigation-id){#scroll-to-fragid:navigation-id}
`navigationId`{.variable}:

1.  Let `navigation`{.variable} be `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#scroll-to-fragid:nav-window}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#scroll-to-fragid:window-navigation-api}.

2.  Let `destinationNavigationAPIState`{.variable} be
    `navigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry}\'s
    [navigation API
    state](#she-navigation-api-state){#scroll-to-fragid:she-navigation-api-state}.

3.  If `navigationAPIState`{.variable} is not null, then set
    `destinationNavigationAPIState`{.variable} to
    `navigationAPIState`{.variable}.

4.  Let `continue`{.variable} be the result of [firing a
    push/replace/reload `navigate`
    event](nav-history-apis.html#fire-a-push/replace/reload-navigate-event){#scroll-to-fragid:fire-a-push/replace/reload-navigate-event}
    at `navigation`{.variable} with
    *[navigationType](nav-history-apis.html#fire-navigate-prr-navigationtype)*
    set to `historyHandling`{.variable},
    *[isSameDocument](nav-history-apis.html#fire-navigate-prr-issamedocument)*
    set to true,
    *[userInvolvement](nav-history-apis.html#fire-navigate-prr-userinvolvement)*
    set to `userInvolvement`{.variable},
    *[sourceElement](nav-history-apis.html#fire-navigate-prr-sourceelement)*
    set to `sourceElement`{.variable},
    *[destinationURL](nav-history-apis.html#fire-navigate-prr-destinationurl)*
    set to `url`{.variable}, and
    *[navigationAPIState](nav-history-apis.html#fire-navigate-prr-navigationapistate)*
    set to `destinationNavigationAPIState`{.variable}.

5.  If `continue`{.variable} is false, then return.

6.  Let `historyEntry`{.variable} be a new [session history
    entry](#session-history-entry){#scroll-to-fragid:session-history-entry},
    with

    [URL](#she-url){#scroll-to-fragid:she-url}
    :   `url`{.variable}

    [document state](#she-document-state){#scroll-to-fragid:she-document-state}
    :   `navigable`{.variable}\'s [active session history
        entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry-2}\'s
        [document
        state](#she-document-state){#scroll-to-fragid:she-document-state-2}

    [navigation API state](#she-navigation-api-state){#scroll-to-fragid:she-navigation-api-state-2}
    :   `destinationNavigationAPIState`{.variable}

    [scroll restoration mode](#she-scroll-restoration-mode){#scroll-to-fragid:she-scroll-restoration-mode}
    :   `navigable`{.variable}\'s [active session history
        entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry-3}\'s
        [scroll restoration
        mode](#she-scroll-restoration-mode){#scroll-to-fragid:she-scroll-restoration-mode-2}

    ::: note
    For navigations performed with
    [`navigation.navigate()`{#scroll-to-fragid:dom-navigation-navigate}](nav-history-apis.html#dom-navigation-navigate),
    the value provided by the
    [`state`{#scroll-to-fragid:dom-navigationnavigateoptions-state}](nav-history-apis.html#dom-navigationnavigateoptions-state)
    option is used for the new [navigation API
    state](#she-navigation-api-state){#scroll-to-fragid:she-navigation-api-state-3}.
    (This will set it to the serialization of undefined, if no value is
    provided for that option.) For other fragment navigations, including
    user-initiated ones, the [navigation API
    state](#she-navigation-api-state){#scroll-to-fragid:she-navigation-api-state-4}
    is carried over from the previous entry.

    The [classic history API
    state](#she-classic-history-api-state){#scroll-to-fragid:she-classic-history-api-state}
    is never carried over.
    :::

7.  Let `entryToReplace`{.variable} be `navigable`{.variable}\'s [active
    session history
    entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry-4}
    if `historyHandling`{.variable} is
    \"[`replace`{#scroll-to-fragid:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\",
    otherwise null.

8.  Let `history`{.variable} be `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#scroll-to-fragid:nav-document}\'s
    [history
    object](nav-history-apis.html#doc-history){#scroll-to-fragid:doc-history}.

9.  Let `scriptHistoryIndex`{.variable} be `history`{.variable}\'s
    [index](nav-history-apis.html#concept-history-index){#scroll-to-fragid:concept-history-index}.

10. Let `scriptHistoryLength`{.variable} be `history`{.variable}\'s
    [length](nav-history-apis.html#concept-history-length){#scroll-to-fragid:concept-history-length}.

11. If `historyHandling`{.variable} is
    \"[`push`{#scroll-to-fragid:navigationhistorybehavior-push}](#navigationhistorybehavior-push)\",
    then:

    1.  Set `history`{.variable}\'s
        [state](nav-history-apis.html#concept-history-state){#scroll-to-fragid:concept-history-state}
        to null.

    2.  Increment `scriptHistoryIndex`{.variable}.

    3.  Set `scriptHistoryLength`{.variable} to
        `scriptHistoryIndex`{.variable} + 1.

12. Set `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#scroll-to-fragid:nav-document-2}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#scroll-to-fragid:the-document's-address
    x-internal="the-document's-address"} to `url`{.variable}.

13. Set `navigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry-5}
    to `historyEntry`{.variable}.

14. [Update document for history step
    application](#update-document-for-history-step-application){#scroll-to-fragid:update-document-for-history-step-application}
    given `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#scroll-to-fragid:nav-document-3},
    `historyEntry`{.variable}, true, `scriptHistoryIndex`{.variable},
    `scriptHistoryLength`{.variable}, and `historyHandling`{.variable}.

    This algorithm will be called twice as a result of a single fragment
    navigation: once synchronously, where best-guess values
    `scriptHistoryIndex`{.variable} and `scriptHistoryLength`{.variable}
    are set,
    [`history.state`{#scroll-to-fragid:dom-history-state}](nav-history-apis.html#dom-history-state)
    is nulled out, and various events are fired; and once
    asynchronously, where the final values for index and length are set,
    [`history.state`{#scroll-to-fragid:dom-history-state-2}](nav-history-apis.html#dom-history-state)
    remains untouched, and no events are fired.

15. [Scroll to the
    fragment](#scroll-to-the-fragment-identifier){#scroll-to-fragid:scroll-to-the-fragment-identifier}
    given `navigable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#scroll-to-fragid:nav-document-4}.

    If the scrolling fails because the
    [`Document`{#scroll-to-fragid:document}](dom.html#document) is new
    and the relevant
    [ID](https://dom.spec.whatwg.org/#concept-id){#scroll-to-fragid:concept-id
    x-internal="concept-id"} has not yet been parsed, then the second
    asynchronous call to [update document for history step
    application](#update-document-for-history-step-application){#scroll-to-fragid:update-document-for-history-step-application-2}
    will take care of scrolling.

16. Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#scroll-to-fragid:nav-traversable}.

17. [Append the following session history synchronous navigation
    steps](#tn-append-session-history-sync-nav-steps){#scroll-to-fragid:tn-append-session-history-sync-nav-steps}
    involving `navigable`{.variable} to `traversable`{.variable}:

    1.  [Finalize a same-document
        navigation](#finalize-a-same-document-navigation){#scroll-to-fragid:finalize-a-same-document-navigation}
        given `traversable`{.variable}, `navigable`{.variable},
        `historyEntry`{.variable}, `entryToReplace`{.variable},
        `historyHandling`{.variable}, and `userInvolvement`{.variable}.

    2.  Invoke [WebDriver BiDi fragment
        navigated](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-fragment-navigated){#scroll-to-fragid:webdriver-bidi-fragment-navigated
        x-internal="webdriver-bidi-fragment-navigated"} with
        `navigable`{.variable} and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#scroll-to-fragid:webdriver-bidi-navigation-status
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#scroll-to-fragid:navigation-status-id
        x-internal="navigation-status-id"} is `navigationId`{.variable},
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#scroll-to-fragid:navigation-status-url
        x-internal="navigation-status-url"} is `url`{.variable}, and
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#scroll-to-fragid:navigation-status-status
        x-internal="navigation-status-status"} is
        \"[`complete`{#scroll-to-fragid:navigation-status-complete}](https://w3c.github.io/webdriver-bidi/#navigation-status-complete){x-internal="navigation-status-complete"}\".

To [finalize a same-document
navigation]{#finalize-a-same-document-navigation .dfn} given a
[traversable
navigable](document-sequences.html#traversable-navigable){#scroll-to-fragid:traversable-navigable}
`traversable`{.variable}, a
[navigable](document-sequences.html#navigable){#scroll-to-fragid:navigable-2}
`targetNavigable`{.variable}, a [session history
entry](#session-history-entry){#scroll-to-fragid:session-history-entry-2}
`targetEntry`{.variable}, a [session history
entry](#session-history-entry){#scroll-to-fragid:session-history-entry-3}-or-null
`entryToReplace`{.variable}, a [history handling
behavior](#history-handling-behavior){#scroll-to-fragid:history-handling-behavior-2}
`historyHandling`{.variable}, and a [user navigation
involvement](#user-navigation-involvement){#scroll-to-fragid:user-navigation-involvement-2}
`userInvolvement`{.variable}:

This is used by both [fragment
navigations](#navigate-fragid){#scroll-to-fragid:navigate-fragid} and by
the [URL and history update
steps](#url-and-history-update-steps){#scroll-to-fragid:url-and-history-update-steps},
which are the only synchronous updates to session history. By virtue of
being synchronous, those algorithms are performed outside of the
[top-level
traversable](document-sequences.html#top-level-traversable){#scroll-to-fragid:top-level-traversable}\'s
[session history traversal
queue](document-sequences.html#tn-session-history-traversal-queue){#scroll-to-fragid:tn-session-history-traversal-queue}.
This puts them out of sync with the [top-level
traversable](document-sequences.html#top-level-traversable){#scroll-to-fragid:top-level-traversable-2}\'s
[current session history
step](document-sequences.html#tn-current-session-history-step){#scroll-to-fragid:tn-current-session-history-step},
so this algorithm is used to resolve conflicts due to race conditions.

1.  [Assert](https://infra.spec.whatwg.org/#assert){#scroll-to-fragid:assert
    x-internal="assert"}: this is running on `traversable`{.variable}\'s
    [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#scroll-to-fragid:tn-session-history-traversal-queue-2}.

2.  If `targetNavigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#scroll-to-fragid:nav-active-history-entry-6}
    is not `targetEntry`{.variable}, then return.

3.  Let `targetStep`{.variable} be null.

4.  Let `targetEntries`{.variable} be the result of [getting session
    history
    entries](#getting-session-history-entries){#scroll-to-fragid:getting-session-history-entries}
    for `targetNavigable`{.variable}.

5.  If `entryToReplace`{.variable} is null, then:

    1.  [Clear the forward session
        history](#clear-the-forward-session-history){#scroll-to-fragid:clear-the-forward-session-history}
        of `traversable`{.variable}.

    2.  Set `targetStep`{.variable} to `traversable`{.variable}\'s
        [current session history
        step](document-sequences.html#tn-current-session-history-step){#scroll-to-fragid:tn-current-session-history-step-2} +
        1.

    3.  Set `targetEntry`{.variable}\'s
        [step](#she-step){#scroll-to-fragid:she-step} to
        `targetStep`{.variable}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#scroll-to-fragid:list-append
        x-internal="list-append"} `targetEntry`{.variable} to
        `targetEntries`{.variable}.

    Otherwise:

    1.  [Replace](https://infra.spec.whatwg.org/#list-replace){#scroll-to-fragid:list-replace
        x-internal="list-replace"} `entryToReplace`{.variable} with
        `targetEntry`{.variable} in `targetEntries`{.variable}.

    2.  Set `targetEntry`{.variable}\'s
        [step](#she-step){#scroll-to-fragid:she-step-2} to
        `entryToReplace`{.variable}\'s
        [step](#she-step){#scroll-to-fragid:she-step-3}.

    3.  Set `targetStep`{.variable} to `traversable`{.variable}\'s
        [current session history
        step](document-sequences.html#tn-current-session-history-step){#scroll-to-fragid:tn-current-session-history-step-3}.

6.  [Apply the push/replace history
    step](#apply-the-push/replace-history-step){#scroll-to-fragid:apply-the-push/replace-history-step}
    `targetStep`{.variable} to `traversable`{.variable} given
    `historyHandling`{.variable} and `userInvolvement`{.variable}.

    This is done even for
    \"[`replace`{#scroll-to-fragid:navigationhistorybehavior-replace-2}](#navigationhistorybehavior-replace)\"
    navigations, as it resolves race conditions across multiple
    synchronous navigations.

###### [7.4.2.3.4]{.secno} Non-fetch schemes and external software[](#non-fetch-schemes-and-external-software){.self-link}

The input to [attempt to create a non-fetch scheme
document](#attempt-to-create-a-non-fetch-scheme-document){#non-fetch-schemes-and-external-software:attempt-to-create-a-non-fetch-scheme-document}
is the [non-fetch scheme navigation
params]{#non-fetch-scheme-navigation-params .dfn}
[struct](https://infra.spec.whatwg.org/#struct){#non-fetch-schemes-and-external-software:struct
x-internal="struct"}. It is a lightweight version of [navigation
params](#navigation-params){#non-fetch-schemes-and-external-software:navigation-params}
which only carries parameters relevant to the non-[fetch
scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#non-fetch-schemes-and-external-software:fetch-scheme
x-internal="fetch-scheme"} navigation case. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#non-fetch-schemes-and-external-software:struct-item
x-internal="struct-item"}:

[id]{#non-fetch-scheme-params-id .dfn}
:   null or a [navigation
    ID](#navigation-id){#non-fetch-schemes-and-external-software:navigation-id}

[navigable]{#non-fetch-scheme-params-navigable .dfn}
:   the
    [navigable](document-sequences.html#navigable){#non-fetch-schemes-and-external-software:navigable}
    experiencing the navigation

[URL]{#non-fetch-scheme-params-url .dfn}
:   a
    [URL](https://url.spec.whatwg.org/#concept-url){#non-fetch-schemes-and-external-software:url
    x-internal="url"}

[target snapshot sandboxing flags]{#non-fetch-scheme-params-target-sandbox .dfn}
:   the [target snapshot
    params](#target-snapshot-params){#non-fetch-schemes-and-external-software:target-snapshot-params}\'s
    [sandboxing
    flags](#target-snapshot-params-sandbox){#non-fetch-schemes-and-external-software:target-snapshot-params-sandbox}
    present during navigation

[source snapshot has transient activation]{#non-fetch-scheme-params-source-activation .dfn}
:   a copy of the [source snapshot
    params](#source-snapshot-params){#non-fetch-schemes-and-external-software:source-snapshot-params}\'s
    [has transient
    activation](#source-snapshot-params-activation){#non-fetch-schemes-and-external-software:source-snapshot-params-activation}
    boolean present during activation

[initiator origin]{#non-fetch-scheme-params-initiator-origin .dfn}

:   an
    [origin](browsers.html#concept-origin){#non-fetch-schemes-and-external-software:concept-origin}
    possibly for use in a user-facing prompt to confirm the invocation
    of an external software package

    This differs slightly from a [document
    state](#she-document-state){#non-fetch-schemes-and-external-software:she-document-state}\'s
    [initiator
    origin](#document-state-initiator-origin){#non-fetch-schemes-and-external-software:document-state-initiator-origin}
    in that a [non-fetch scheme navigation
    params](#non-fetch-scheme-navigation-params){#non-fetch-schemes-and-external-software:non-fetch-scheme-navigation-params}\'s
    [initiator
    origin](#non-fetch-scheme-params-initiator-origin){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-initiator-origin}
    follows redirects up to the last [fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#non-fetch-schemes-and-external-software:fetch-scheme-2
    x-internal="fetch-scheme"} URL in a redirect chain that ends in a
    non-[fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#non-fetch-schemes-and-external-software:fetch-scheme-3
    x-internal="fetch-scheme"} URL.

[navigation timing type]{#non-fetch-scheme-params-nav-timing-type .dfn}
:   a
    [`NavigationTimingType`{#non-fetch-schemes-and-external-software:navigationtimingtype}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
    used for [creating the navigation timing
    entry](https://w3c.github.io/navigation-timing/#dfn-create-the-navigation-timing-entry){#non-fetch-schemes-and-external-software:create-the-navigation-timing-entry
    x-internal="create-the-navigation-timing-entry"} for the new
    [`Document`{#non-fetch-schemes-and-external-software:document}](dom.html#document)
    (if one is created)

[user involvement]{#non-fetch-scheme-params-user-involvement .dfn}
:   a [user navigation
    involvement](#user-navigation-involvement){#non-fetch-schemes-and-external-software:user-navigation-involvement}
    used when [obtaining a browsing
    context](browsers.html#obtain-browsing-context-navigation){#non-fetch-schemes-and-external-software:obtain-browsing-context-navigation}
    for the new
    [`Document`{#non-fetch-schemes-and-external-software:document-2}](dom.html#document)
    (if one is created)

To [attempt to create a non-fetch scheme
document]{#attempt-to-create-a-non-fetch-scheme-document .dfn}, given a
[non-fetch scheme navigation
params](#non-fetch-scheme-navigation-params){#non-fetch-schemes-and-external-software:non-fetch-scheme-navigation-params-2}
`navigationParams`{.variable}:

1.  Let `url`{.variable} be `navigationParams`{.variable}\'s
    [URL](#non-fetch-scheme-params-url){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-url}.

2.  Let `navigable`{.variable} be `navigationParams`{.variable}\'s
    [navigable](#non-fetch-scheme-params-navigable){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-navigable}.

3.  If `url`{.variable} is to be handled using a mechanism that does not
    affect `navigable`{.variable}, e.g., because `url`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#non-fetch-schemes-and-external-software:concept-url-scheme
    x-internal="concept-url-scheme"} is handled externally, then:

    1.  [Hand-off to external
        software](#hand-off-to-external-software){#non-fetch-schemes-and-external-software:hand-off-to-external-software}
        given `url`{.variable}, `navigable`{.variable},
        `navigationParams`{.variable}\'s [target snapshot sandboxing
        flags](#non-fetch-scheme-params-target-sandbox){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-target-sandbox},
        `navigationParams`{.variable}\'s [source snapshot has transient
        activation](#non-fetch-scheme-params-source-activation){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-source-activation},
        and `navigationParams`{.variable}\'s [initiator
        origin](#non-fetch-scheme-params-initiator-origin){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-initiator-origin-2}.

    2.  Return null.

4.  Handle `url`{.variable} by displaying some sort of inline content,
    e.g., an error message because the specified scheme is not one of
    the supported protocols, or an inline prompt to allow the user to
    select [a registered
    handler](system-state.html#dom-navigator-registerprotocolhandler){#non-fetch-schemes-and-external-software:dom-navigator-registerprotocolhandler}
    for the given scheme. Return the result of [displaying the inline
    content](document-lifecycle.html#read-ua-inline){#non-fetch-schemes-and-external-software:read-ua-inline}
    given `navigable`{.variable}, `navigationParams`{.variable}\'s
    [id](#non-fetch-scheme-params-id){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-id},
    `navigationParams`{.variable}\'s [navigation timing
    type](#non-fetch-scheme-params-nav-timing-type){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-nav-timing-type},
    and `navigationParams`{.variable}\'s [user
    involvement](#non-fetch-scheme-params-user-involvement){#non-fetch-schemes-and-external-software:non-fetch-scheme-params-user-involvement}.

    In the case of a registered handler being used,
    [navigate](#navigate){#non-fetch-schemes-and-external-software:navigate}
    will be invoked with a new URL.

To [hand-off to external software]{#hand-off-to-external-software .dfn}
given a
[URL](https://url.spec.whatwg.org/#concept-url){#non-fetch-schemes-and-external-software:url-2
x-internal="url"} or
[response](https://fetch.spec.whatwg.org/#concept-response){#non-fetch-schemes-and-external-software:concept-response
x-internal="concept-response"} `resource`{.variable}, a
[navigable](document-sequences.html#navigable){#non-fetch-schemes-and-external-software:navigable-2}
`navigable`{.variable}, a [sandboxing flag
set](browsers.html#sandboxing-flag-set){#non-fetch-schemes-and-external-software:sandboxing-flag-set}
`sandboxFlags`{.variable}, a boolean
`hasTransientActivation`{.variable}, and an
[origin](browsers.html#concept-origin){#non-fetch-schemes-and-external-software:concept-origin-2}
`initiatorOrigin`{.variable}, user agents should:

1.  If all of the following are true:

    - `navigable`{.variable} is not a [top-level
      traversable](document-sequences.html#top-level-traversable){#non-fetch-schemes-and-external-software:top-level-traversable};

    - `sandboxFlags`{.variable} has its [sandboxed custom protocols
      navigation browsing context
      flag](browsers.html#sandboxed-custom-protocols-navigation-browsing-context-flag){#non-fetch-schemes-and-external-software:sandboxed-custom-protocols-navigation-browsing-context-flag}
      set; and

    - `sandboxFlags`{.variable} has its [sandboxed top-level navigation
      with user activation browsing context
      flag](browsers.html#sandboxed-top-level-navigation-with-user-activation-browsing-context-flag){#non-fetch-schemes-and-external-software:sandboxed-top-level-navigation-with-user-activation-browsing-context-flag}
      set, or `hasTransientActivation`{.variable} is false,

    then return without invoking the external software package.

    Navigation inside an iframe toward external software can be seen by
    users as a new popup or a new top-level navigation. That\'s why its
    is allowed in sandboxed
    [`iframe`{#non-fetch-schemes-and-external-software:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    only when one of
    [`allow-popups`{#non-fetch-schemes-and-external-software:attr-iframe-sandbox-allow-popups}](browsers.html#attr-iframe-sandbox-allow-popups),
    [`allow-top-navigation`{#non-fetch-schemes-and-external-software:attr-iframe-sandbox-allow-top-navigation}](browsers.html#attr-iframe-sandbox-allow-top-navigation),
    [`allow-top-navigation-by-user-activation`{#non-fetch-schemes-and-external-software:attr-iframe-sandbox-allow-top-navigation-by-user-activation}](browsers.html#attr-iframe-sandbox-allow-top-navigation-by-user-activation),
    or
    [`allow-top-navigation-to-custom-protocols`{#non-fetch-schemes-and-external-software:attr-iframe-sandbox-allow-top-navigation-to-custom-protocols}](browsers.html#attr-iframe-sandbox-allow-top-navigation-to-custom-protocols)
    is specified.

2.  Perform the appropriate handoff of `resource`{.variable} while
    attempting to mitigate the risk that this is an attempt to exploit
    the target software. For example, user agents could prompt the user
    to confirm that `initiatorOrigin`{.variable} is to be allowed to
    invoke the external software in question. In particular, if
    `hasTransientActivation`{.variable} is false, then the user agent
    should not invoke the external software package without prior user
    confirmation.

    For example, there could be a vulnerability in the target
    software\'s URL handler which a hostile page would attempt to
    exploit by tricking a user into clicking a link.

##### [7.4.2.4]{.secno} Preventing navigation[](#preventing-navigation){.self-link}

A couple of scenarios can intervene early in the navigation process and
put the whole thing to a halt. This can be especially exciting when
multiple
[navigables](document-sequences.html#navigable){#preventing-navigation:navigable}
are navigating at the same time, due to a session history traversal.

A
[navigable](document-sequences.html#navigable){#preventing-navigation:navigable-2}
`source`{.variable} is [allowed by sandboxing to
navigate]{#allowed-to-navigate .dfn} a second
[navigable](document-sequences.html#navigable){#preventing-navigation:navigable-3}
`target`{.variable}, given a [source snapshot
params](#source-snapshot-params){#preventing-navigation:source-snapshot-params}
`sourceSnapshotParams`{.variable}, if the following steps return true:

1.  If `source`{.variable} is `target`{.variable}, then return true.

2.  If `source`{.variable} is an ancestor of `target`{.variable}, then
    return true.

3.  If `target`{.variable} is an ancestor of `source`{.variable}, then:

    1.  If `target`{.variable} is not a [top-level
        traversable](document-sequences.html#top-level-traversable){#preventing-navigation:top-level-traversable},
        then return true.

    2.  If `sourceSnapshotParams`{.variable}\'s [has transient
        activation](#source-snapshot-params-activation){#preventing-navigation:source-snapshot-params-activation}
        is true, and `sourceSnapshotParams`{.variable}\'s [sandboxing
        flags](#source-snapshot-params-sandbox){#preventing-navigation:source-snapshot-params-sandbox}\'s
        [sandboxed top-level navigation with user activation browsing
        context
        flag](browsers.html#sandboxed-top-level-navigation-with-user-activation-browsing-context-flag){#preventing-navigation:sandboxed-top-level-navigation-with-user-activation-browsing-context-flag}
        is set, then return false.

    3.  If `sourceSnapshotParams`{.variable}\'s [has transient
        activation](#source-snapshot-params-activation){#preventing-navigation:source-snapshot-params-activation-2}
        is false, and `sourceSnapshotParams`{.variable}\'s [sandboxing
        flags](#source-snapshot-params-sandbox){#preventing-navigation:source-snapshot-params-sandbox-2}\'s
        [sandboxed top-level navigation without user activation browsing
        context
        flag](browsers.html#sandboxed-top-level-navigation-without-user-activation-browsing-context-flag){#preventing-navigation:sandboxed-top-level-navigation-without-user-activation-browsing-context-flag}
        is set, then return false.

    4.  Return true.

4.  If `target`{.variable} is a [top-level
    traversable](document-sequences.html#top-level-traversable){#preventing-navigation:top-level-traversable-2}:

    1.  If `source`{.variable} is the [one permitted sandboxed
        navigator](browsers.html#one-permitted-sandboxed-navigator){#preventing-navigation:one-permitted-sandboxed-navigator}
        of `target`{.variable}, then return true.

    2.  If `sourceSnapshotParams`{.variable}\'s [sandboxing
        flags](#source-snapshot-params-sandbox){#preventing-navigation:source-snapshot-params-sandbox-3}\'s
        [sandboxed navigation browsing context
        flag](browsers.html#sandboxed-navigation-browsing-context-flag){#preventing-navigation:sandboxed-navigation-browsing-context-flag}
        is set, then return false.

    3.  Return true.

5.  If `sourceSnapshotParams`{.variable}\'s [sandboxing
    flags](#source-snapshot-params-sandbox){#preventing-navigation:source-snapshot-params-sandbox-4}\'s
    [sandboxed navigation browsing context
    flag](browsers.html#sandboxed-navigation-browsing-context-flag){#preventing-navigation:sandboxed-navigation-browsing-context-flag-2}
    is set, then return false.

6.  Return true.

[]{#prompt-to-unload}[]{#checking-if-unloading-is-user-canceled}To
[check if unloading is canceled]{#checking-if-unloading-is-canceled
.dfn} for a
[list](https://infra.spec.whatwg.org/#list){#preventing-navigation:list
x-internal="list"} of
[navigables](document-sequences.html#navigable){#preventing-navigation:navigable-4}
`navigablesThatNeedBeforeUnload`{.variable}, given an optional
[traversable
navigable](document-sequences.html#traversable-navigable){#preventing-navigation:traversable-navigable}
`traversable`{.variable}, an optional integer `targetStep`{.variable},
and an optional [user navigation
involvement](#user-navigation-involvement){#preventing-navigation:user-navigation-involvement}
`userInvolvementForNavigateEvent`{.variable}, run these steps. They
return \"`canceled-by-beforeunload`\", \"`canceled-by-navigate`\", or
\"`continue`\".

1.  Let `documentsToFireBeforeunload`{.variable} be the [active
    document](document-sequences.html#nav-document){#preventing-navigation:nav-document}
    of each
    [item](https://infra.spec.whatwg.org/#list-item){#preventing-navigation:list-item
    x-internal="list-item"} in
    `navigablesThatNeedBeforeUnload`{.variable}.

2.  Let `unloadPromptShown`{.variable} be false.

3.  Let `finalStatus`{.variable} be \"`continue`\".

4.  If `traversable`{.variable} was given, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#preventing-navigation:assert
        x-internal="assert"}: `targetStep`{.variable} and
        `userInvolvementForNavigateEvent`{.variable} were given.

    2.  Let `targetEntry`{.variable} be the result of [getting the
        target history
        entry](#getting-the-target-history-entry){#preventing-navigation:getting-the-target-history-entry}
        given `traversable`{.variable} and `targetStep`{.variable}.

    3.  If `targetEntry`{.variable} is not `traversable`{.variable}\'s
        [current session history
        entry](document-sequences.html#nav-current-history-entry){#preventing-navigation:nav-current-history-entry},
        and `targetEntry`{.variable}\'s [document
        state](#she-document-state){#preventing-navigation:she-document-state}\'s
        [origin](#document-state-origin){#preventing-navigation:document-state-origin}
        is the
        [same](browsers.html#same-origin){#preventing-navigation:same-origin}
        as `traversable`{.variable}\'s [current session history
        entry](document-sequences.html#nav-current-history-entry){#preventing-navigation:nav-current-history-entry-2}\'s
        [document
        state](#she-document-state){#preventing-navigation:she-document-state-2}\'s
        [origin](#document-state-origin){#preventing-navigation:document-state-origin-2},
        then:

        ::: note
        In this case, we\'re going to fire the
        [`navigate`{#preventing-navigation:event-navigate}](indices.html#event-navigate)
        event for `traversable`{.variable} here. Because [under some
        circumstances](nav-history-apis.html#navigate-event-traverse-can-be-canceled)
        it might be canceled, we need to do this separately from [other
        traversal `navigate`
        events](#descendant-navigable-traversal-navigate-events), which
        happen later.

        Additionally, because we want
        [`beforeunload`{#preventing-navigation:event-beforeunload}](indices.html#event-beforeunload)
        events to fire before
        [`navigate`{#preventing-navigation:event-navigate-2}](indices.html#event-navigate)
        events, this means we need to fire
        [`beforeunload`{#preventing-navigation:event-beforeunload-2}](indices.html#event-beforeunload)
        for `traversable`{.variable} here (if applicable), instead of
        doing it as part of the below loop over
        `documentsToFireBeforeunload`{.variable}.
        :::

        1.  Let `eventsFired`{.variable} be false.

        2.  Let `needsBeforeunload`{.variable} be true if
            `navigablesThatNeedBeforeUnload`{.variable}
            [contains](https://infra.spec.whatwg.org/#list-contain){#preventing-navigation:list-contains
            x-internal="list-contains"} `traversable`{.variable};
            otherwise false.

        3.  If `needsBeforeunload`{.variable} is true, then
            [remove](https://infra.spec.whatwg.org/#list-remove){#preventing-navigation:list-remove
            x-internal="list-remove"} `traversable`{.variable}\'s
            [active
            document](document-sequences.html#nav-document){#preventing-navigation:nav-document-2}
            from `documentsToFireBeforeunload`{.variable}.

        4.  [Queue a global
            task](webappapis.html#queue-a-global-task){#preventing-navigation:queue-a-global-task}
            on the [navigation and traversal task
            source](webappapis.html#navigation-and-traversal-task-source){#preventing-navigation:navigation-and-traversal-task-source}
            given `traversable`{.variable}\'s [active
            window](document-sequences.html#nav-window){#preventing-navigation:nav-window}
            to perform the following steps:

            1.  If `needsBeforeunload`{.variable} is true, then:

                1.  Let (`unloadPromptShownForThisDocument`{.variable},
                    `unloadPromptCanceledByThisDocument`{.variable}) be
                    the result of running the [steps to fire
                    `beforeunload`](#steps-to-fire-beforeunload){#preventing-navigation:steps-to-fire-beforeunload}
                    given `traversable`{.variable}\'s [active
                    document](document-sequences.html#nav-document){#preventing-navigation:nav-document-3}
                    and false.

                2.  If `unloadPromptShownForThisDocument`{.variable} is
                    true, then set `unloadPromptShown`{.variable} to
                    true.

                3.  If `unloadPromptCanceledByThisDocument`{.variable}
                    is true, then set `finalStatus`{.variable} to
                    \"`canceled-by-beforeunload`\".

            2.  If `finalStatus`{.variable} is
                \"`canceled-by-beforeunload`\", then abort these steps.

            3.  Let `navigation`{.variable} be
                `traversable`{.variable}\'s [active
                window](document-sequences.html#nav-window){#preventing-navigation:nav-window-2}\'s
                [navigation
                API](nav-history-apis.html#window-navigation-api){#preventing-navigation:window-navigation-api}.

            4.  Let `navigateEventResult`{.variable} be the result of
                [firing a traverse `navigate`
                event](nav-history-apis.html#fire-a-traverse-navigate-event){#preventing-navigation:fire-a-traverse-navigate-event}
                at `navigation`{.variable} given
                `targetEntry`{.variable} and
                `userInvolvementForNavigateEvent`{.variable}.

            5.  If `navigateEventResult`{.variable} is false, then set
                `finalStatus`{.variable} to \"`canceled-by-navigate`\".

            6.  Set `eventsFired`{.variable} to true.

        5.  Wait until `eventsFired`{.variable} is true.

        6.  If `finalStatus`{.variable} is not \"`continue`\", then
            return `finalStatus`{.variable}.

5.  Let `totalTasks`{.variable} be the
    [size](https://infra.spec.whatwg.org/#list-size){#preventing-navigation:list-size
    x-internal="list-size"} of `documentsToFireBeforeunload`{.variable}.

6.  Let `completedTasks`{.variable} be 0.

7.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#preventing-navigation:list-iterate
    x-internal="list-iterate"} `document`{.variable} of
    `documentsToFireBeforeunload`{.variable}, [queue a global
    task](webappapis.html#queue-a-global-task){#preventing-navigation:queue-a-global-task-2}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#preventing-navigation:navigation-and-traversal-task-source-2}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#preventing-navigation:concept-relevant-global}
    to run the steps:

    1.  Let (`unloadPromptShownForThisDocument`{.variable},
        `unloadPromptCanceledByThisDocument`{.variable}) be the result
        of running the [steps to fire
        `beforeunload`](#steps-to-fire-beforeunload){#preventing-navigation:steps-to-fire-beforeunload-2}
        given `document`{.variable} and `unloadPromptShown`{.variable}.

    2.  If `unloadPromptShownForThisDocument`{.variable} is true, then
        set `unloadPromptShown`{.variable} to true.

    3.  If `unloadPromptCanceledByThisDocument`{.variable} is true, then
        set `finalStatus`{.variable} to \"`canceled-by-beforeunload`\".

    4.  Increment `completedTasks`{.variable}.

8.  Wait for `completedTasks`{.variable} to be `totalTasks`{.variable}.

9.  Return `finalStatus`{.variable}.

The [steps to fire `beforeunload`]{#steps-to-fire-beforeunload .dfn}
given a [`Document`{#preventing-navigation:document}](dom.html#document)
`document`{.variable} and a boolean `unloadPromptShown`{.variable} are:

1.  Let `unloadPromptCanceled`{.variable} be false.

2.  Increase the `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#preventing-navigation:unload-counter}
    by 1.

3.  Increase `document`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#preventing-navigation:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#preventing-navigation:concept-agent-event-loop}\'s
    [termination nesting
    level](document-lifecycle.html#termination-nesting-level){#preventing-navigation:termination-nesting-level}
    by 1.

4.  Let `eventFiringResult`{.variable} be the result of [firing an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#preventing-navigation:concept-event-fire
    x-internal="concept-event-fire"} named
    [`beforeunload`{#preventing-navigation:event-beforeunload-3}](indices.html#event-beforeunload)
    at `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#preventing-navigation:concept-relevant-global-2},
    using
    [`BeforeUnloadEvent`{#preventing-navigation:beforeunloadevent}](nav-history-apis.html#beforeunloadevent),
    with the
    [`cancelable`{#preventing-navigation:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attribute initialized to true.

5.  Decrease `document`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#preventing-navigation:relevant-agent-2}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#preventing-navigation:concept-agent-event-loop-2}\'s
    [termination nesting
    level](document-lifecycle.html#termination-nesting-level){#preventing-navigation:termination-nesting-level-2}
    by 1.

6.  If all of the following are true:

    - `unloadPromptShown`{.variable} is false;

    - `document`{.variable}\'s [active sandboxing flag
      set](browsers.html#active-sandboxing-flag-set){#preventing-navigation:active-sandboxing-flag-set}
      does not have its [sandboxed modals
      flag](browsers.html#sandboxed-modals-flag){#preventing-navigation:sandboxed-modals-flag}
      set;

    - `document`{.variable}\'s [relevant global
      object](webappapis.html#concept-relevant-global){#preventing-navigation:concept-relevant-global-3}
      has [sticky
      activation](interaction.html#sticky-activation){#preventing-navigation:sticky-activation};

    - `eventFiringResult`{.variable} is false, or the
      [`returnValue`{#preventing-navigation:dom-beforeunloadevent-returnvalue}](nav-history-apis.html#dom-beforeunloadevent-returnvalue)
      attribute of `event`{.variable} is not the empty string; and

    - showing an unload prompt is unlikely to be annoying, deceptive, or
      pointless,

    then:

    1.  Set `unloadPromptShown`{.variable} to true.

    2.  Let `userPromptHandler`{.variable} be the result of [WebDriver
        BiDi user prompt
        opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-opened){#preventing-navigation:webdriver-bidi-user-prompt-opened
        x-internal="webdriver-bidi-user-prompt-opened"} with
        `document`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#preventing-navigation:concept-relevant-global-4},
        \"`beforeunload`\", and \"\".

    3.  If `userPromptHandler`{.variable} is \"`dismiss`\", then set
        `unloadPromptCanceled`{.variable} to true.

    4.  If `userPromptHandler`{.variable} is \"`none`\", then:

        1.  Ask the user to confirm that they wish to unload the
            document, and
            [pause](webappapis.html#pause){#preventing-navigation:pause}
            while waiting for the user\'s response.

            The message shown to the user is not customizable, but
            instead determined by the user agent. In particular, the
            actual value of the
            [`returnValue`{#preventing-navigation:dom-beforeunloadevent-returnvalue-2}](nav-history-apis.html#dom-beforeunloadevent-returnvalue)
            attribute is ignored.

        2.  If the user did not confirm the page navigation, then set
            `unloadPromptCanceled`{.variable} to true.

    5.  Invoke [WebDriver BiDi user prompt
        closed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-closed){#preventing-navigation:webdriver-bidi-user-prompt-closed
        x-internal="webdriver-bidi-user-prompt-closed"} with
        `document`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#preventing-navigation:concept-relevant-global-5},
        \"`beforeunload`\", and true if
        `unloadPromptCanceled`{.variable} is false or false otherwise.

7.  Decrease `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#preventing-navigation:unload-counter-2}
    by 1.

8.  Return (`unloadPromptShown`{.variable},
    `unloadPromptCanceled`{.variable}).

##### [7.4.2.5]{.secno} Aborting navigation[](#aborting-navigation){.self-link}

Each
[navigable](document-sequences.html#navigable){#aborting-navigation:navigable}
has an [ongoing navigation]{#ongoing-navigation .dfn}, which is a
[navigation ID](#navigation-id){#aborting-navigation:navigation-id},
\"`traversal`\", or null, initially null. It is used to track navigation
aborting and to prevent any navigations from taking place during
[traversal](#apply-the-traverse-history-step){#aborting-navigation:apply-the-traverse-history-step}.

To [set the ongoing navigation]{#set-the-ongoing-navigation .dfn} for a
[navigable](document-sequences.html#navigable){#aborting-navigation:navigable-2}
`navigable`{.variable} to `newValue`{.variable}:

1.  If `navigable`{.variable}\'s [ongoing
    navigation](#ongoing-navigation){#aborting-navigation:ongoing-navigation}
    is equal to `newValue`{.variable}, then return.

2.  [Inform the navigation API about aborting
    navigation](nav-history-apis.html#inform-the-navigation-api-about-aborting-navigation){#aborting-navigation:inform-the-navigation-api-about-aborting-navigation}
    given `navigable`{.variable}.

3.  Set `navigable`{.variable}\'s [ongoing
    navigation](#ongoing-navigation){#aborting-navigation:ongoing-navigation-2}
    to `newValue`{.variable}.

#### [7.4.3]{.secno} Reloading and traversing[](#reloading-and-traversing){.self-link}

To [reload]{#reload .dfn} a
[navigable](document-sequences.html#navigable){#reloading-and-traversing:navigable}
`navigable`{.variable} given an optional [serialized
state](#serialized-state){#reloading-and-traversing:serialized-state}-or-null
[`navigationAPIState`{.variable}]{#reload-navigation-api-state .dfn}
(default null) and an optional [user navigation
involvement](#user-navigation-involvement){#reloading-and-traversing:user-navigation-involvement}
[`userInvolvement`{.variable}]{#reload-user-involvement .dfn} (default
\"[`none`{#reloading-and-traversing:uni-none}](#uni-none)\"):

1.  If `userInvolvement`{.variable} is not
    \"[`browser UI`{#reloading-and-traversing:uni-browser-ui}](#uni-browser-ui)\",
    then:

    1.  Let `navigation`{.variable} be `navigable`{.variable}\'s [active
        window](document-sequences.html#nav-window){#reloading-and-traversing:nav-window}\'s
        [navigation
        API](nav-history-apis.html#window-navigation-api){#reloading-and-traversing:window-navigation-api}.

    2.  Let `destinationNavigationAPIState`{.variable} be
        `navigable`{.variable}\'s [active session history
        entry](document-sequences.html#nav-active-history-entry){#reloading-and-traversing:nav-active-history-entry}\'s
        [navigation API
        state](#she-navigation-api-state){#reloading-and-traversing:she-navigation-api-state}.

    3.  If `navigationAPIState`{.variable} is not null, then set
        `destinationNavigationAPIState`{.variable} to
        `navigationAPIState`{.variable}.

    4.  Let `continue`{.variable} be the result of [firing a
        push/replace/reload `navigate`
        event](nav-history-apis.html#fire-a-push/replace/reload-navigate-event){#reloading-and-traversing:fire-a-push/replace/reload-navigate-event}
        at `navigation`{.variable} with
        *[navigationType](nav-history-apis.html#fire-navigate-prr-navigationtype)*
        set to
        \"[`reload`{#reloading-and-traversing:dom-navigationtype-reload}](nav-history-apis.html#dom-navigationtype-reload)\",
        *[isSameDocument](nav-history-apis.html#fire-navigate-prr-issamedocument)*
        set to false,
        *[userInvolvement](nav-history-apis.html#fire-navigate-prr-userinvolvement)*
        set to `userInvolvement`{.variable},
        *[destinationURL](nav-history-apis.html#fire-navigate-prr-destinationurl)*
        set to `navigable`{.variable}\'s [active session history
        entry](document-sequences.html#nav-active-history-entry){#reloading-and-traversing:nav-active-history-entry-2}\'s
        [URL](#she-url){#reloading-and-traversing:she-url}, and
        *[navigationAPIState](nav-history-apis.html#fire-navigate-prr-navigationapistate)*
        set to `destinationNavigationAPIState`{.variable}.

    5.  If `continue`{.variable} is false, then return.

2.  Set `navigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#reloading-and-traversing:nav-active-history-entry-3}\'s
    [document
    state](#she-document-state){#reloading-and-traversing:she-document-state}\'s
    [reload
    pending](#document-state-reload-pending){#reloading-and-traversing:document-state-reload-pending}
    to true.

3.  Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#reloading-and-traversing:nav-traversable}.

4.  [Append the following session history traversal
    steps](#tn-append-session-history-traversal-steps){#reloading-and-traversing:tn-append-session-history-traversal-steps}
    to `traversable`{.variable}:

    1.  [Apply the reload history
        step](#apply-the-reload-history-step){#reloading-and-traversing:apply-the-reload-history-step}
        to `traversable`{.variable} given `userInvolvement`{.variable}.

To [traverse the history by a delta]{#traverse-the-history-by-a-delta
.dfn} given a [traversable
navigable](document-sequences.html#traversable-navigable){#reloading-and-traversing:traversable-navigable}
`traversable`{.variable}, an integer `delta`{.variable}, and an optional
[`Document`{#reloading-and-traversing:document}](dom.html#document)
[`sourceDocument`{.variable}]{#traverse-sourcedocument .dfn}:

1.  Let `sourceSnapshotParams`{.variable} and
    `initiatorToCheck`{.variable} be null.

2.  Let `userInvolvement`{.variable} be
    \"[`browser UI`{#reloading-and-traversing:uni-browser-ui-2}](#uni-browser-ui)\".

3.  If `sourceDocument`{.variable} is given, then:

    1.  Set `sourceSnapshotParams`{.variable} to the result of
        [snapshotting source snapshot
        params](#snapshotting-source-snapshot-params){#reloading-and-traversing:snapshotting-source-snapshot-params}
        given `sourceDocument`{.variable}.

    2.  Set `initiatorToCheck`{.variable} to
        `sourceDocument`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#reloading-and-traversing:node-navigable}.

    3.  Set `userInvolvement`{.variable} to
        \"[`none`{#reloading-and-traversing:uni-none-2}](#uni-none)\".

4.  [Append the following session history traversal
    steps](#tn-append-session-history-traversal-steps){#reloading-and-traversing:tn-append-session-history-traversal-steps-2}
    to `traversable`{.variable}:

    1.  Let `allSteps`{.variable} be the result of [getting all used
        history
        steps](#getting-all-used-history-steps){#reloading-and-traversing:getting-all-used-history-steps}
        for `traversable`{.variable}.

    2.  Let `currentStepIndex`{.variable} be the index of
        `traversable`{.variable}\'s [current session history
        step](document-sequences.html#tn-current-session-history-step){#reloading-and-traversing:tn-current-session-history-step}
        within `allSteps`{.variable}.

    3.  Let `targetStepIndex`{.variable} be
        `currentStepIndex`{.variable} plus `delta`{.variable}.

    4.  If `allSteps`{.variable}\[`targetStepIndex`{.variable}\] does
        not
        [exist](https://infra.spec.whatwg.org/#list-contain){#reloading-and-traversing:list-contains
        x-internal="list-contains"}, then abort these steps.

    5.  [Apply the traverse history
        step](#apply-the-traverse-history-step){#reloading-and-traversing:apply-the-traverse-history-step}
        `allSteps`{.variable}\[`targetStepIndex`{.variable}\] to
        `traversable`{.variable}, given
        `sourceSnapshotParams`{.variable},
        `initiatorToCheck`{.variable}, and `userInvolvement`{.variable}.

#### [7.4.4]{.secno} Non-fragment synchronous \"navigations\"[](#navigate-non-frag-sync){.self-link} {#navigate-non-frag-sync}

Apart from the [navigate](#navigate){#navigate-non-frag-sync:navigate}
algorithm, [session history
entries](#session-history-entry){#navigate-non-frag-sync:session-history-entry}
can be pushed or replaced via one more mechanism, the [URL and history
update
steps](#url-and-history-update-steps){#navigate-non-frag-sync:url-and-history-update-steps}.
The most well-known callers of these steps are the
[`history.replaceState()`{#navigate-non-frag-sync:dom-history-replacestate}](nav-history-apis.html#dom-history-replacestate)
and
[`history.pushState()`{#navigate-non-frag-sync:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate)
APIs, but various other parts of the standard also need to perform
updates to the [active history
entry](document-sequences.html#nav-active-history-entry){#navigate-non-frag-sync:nav-active-history-entry},
and they use these steps to do so.

The [URL and history update steps]{#url-and-history-update-steps .dfn},
given a
[`Document`{#navigate-non-frag-sync:document}](dom.html#document)
`document`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#navigate-non-frag-sync:url
x-internal="url"} `newURL`{.variable}, an optional [serialized
state](#serialized-state){#navigate-non-frag-sync:serialized-state}-or-null
[`serializedData`{.variable}]{#uhus-serializeddata .dfn} (default null),
and an optional [history handling
behavior](#history-handling-behavior){#navigate-non-frag-sync:history-handling-behavior}
[`historyHandling`{#uhus-ispush .variable}]{#uhus-historyhandling .dfn}
(default
\"[`replace`{#navigate-non-frag-sync:navigationhistorybehavior-replace}](#navigationhistorybehavior-replace)\"),
are:

1.  Let `navigable`{.variable} be `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#navigate-non-frag-sync:node-navigable}.

2.  Let `activeEntry`{.variable} be `navigable`{.variable}\'s [active
    session history
    entry](document-sequences.html#nav-active-history-entry){#navigate-non-frag-sync:nav-active-history-entry-2}.

3.  Let `newEntry`{.variable} be a new [session history
    entry](#session-history-entry){#navigate-non-frag-sync:session-history-entry-2},
    with

    [URL](#she-url){#navigate-non-frag-sync:she-url}
    :   `newURL`{.variable}

    [serialized state](#she-classic-history-api-state){#navigate-non-frag-sync:she-classic-history-api-state}
    :   if `serializedData`{.variable} is not null,
        `serializedData`{.variable}; otherwise
        `activeEntry`{.variable}\'s [classic history API
        state](#she-classic-history-api-state){#navigate-non-frag-sync:she-classic-history-api-state-2}

    [document state](#she-document-state){#navigate-non-frag-sync:she-document-state}
    :   `activeEntry`{.variable}\'s [document
        state](#she-document-state){#navigate-non-frag-sync:she-document-state-2}

    [scroll restoration mode](#she-scroll-restoration-mode){#navigate-non-frag-sync:she-scroll-restoration-mode}
    :   `activeEntry`{.variable}\'s [scroll restoration
        mode](#she-scroll-restoration-mode){#navigate-non-frag-sync:she-scroll-restoration-mode-2}

    [persisted user state](#she-other){#navigate-non-frag-sync:she-other}
    :   `activeEntry`{.variable}\'s [persisted user
        state](#she-other){#navigate-non-frag-sync:she-other-2}

4.  If `document`{.variable}\'s [is initial
    `about:blank`](dom.html#is-initial-about:blank){#navigate-non-frag-sync:is-initial-about:blank}
    is true, then set `historyHandling`{.variable} to
    \"[`replace`{#navigate-non-frag-sync:navigationhistorybehavior-replace-2}](#navigationhistorybehavior-replace)\".

    This means that
    [`pushState()`{#navigate-non-frag-sync:dom-history-pushstate-2}](nav-history-apis.html#dom-history-pushstate)
    on an [initial
    `about:blank`](dom.html#is-initial-about:blank){#navigate-non-frag-sync:is-initial-about:blank-2}
    [`Document`{#navigate-non-frag-sync:document-2}](dom.html#document)
    behaves as a
    [`replaceState()`{#navigate-non-frag-sync:dom-history-replacestate-2}](nav-history-apis.html#dom-history-replacestate)
    call.

5.  Let `entryToReplace`{.variable} be `activeEntry`{.variable} if
    `historyHandling`{.variable} is
    \"[`replace`{#navigate-non-frag-sync:navigationhistorybehavior-replace-3}](#navigationhistorybehavior-replace)\",
    otherwise null.

6.  If `historyHandling`{.variable} is
    \"[`push`{#navigate-non-frag-sync:navigationhistorybehavior-push}](#navigationhistorybehavior-push)\",
    then:

    1.  Increment `document`{.variable}\'s [history
        object](nav-history-apis.html#doc-history){#navigate-non-frag-sync:doc-history}\'s
        [index](nav-history-apis.html#concept-history-index){#navigate-non-frag-sync:concept-history-index}.

    2.  Set `document`{.variable}\'s [history
        object](nav-history-apis.html#doc-history){#navigate-non-frag-sync:doc-history-2}\'s
        [length](nav-history-apis.html#concept-history-length){#navigate-non-frag-sync:concept-history-length}
        to its
        [index](nav-history-apis.html#concept-history-index){#navigate-non-frag-sync:concept-history-index-2} +
        1.

    These are temporary best-guess values for immediate synchronous
    access.

7.  If `serializedData`{.variable} is not null, then [restore the
    history object
    state](#restore-the-history-object-state){#navigate-non-frag-sync:restore-the-history-object-state}
    given `document`{.variable} and `newEntry`{.variable}.

8.  Set `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#navigate-non-frag-sync:the-document's-address
    x-internal="the-document's-address"} to `newURL`{.variable}.

    Since this is neither a
    [navigation](#navigate){#navigate-non-frag-sync:navigate-2} nor a
    [history
    traversal](#traverse-the-history-by-a-delta){#navigate-non-frag-sync:traverse-the-history-by-a-delta},
    it does not cause a
    [`hashchange`{#navigate-non-frag-sync:event-hashchange}](indices.html#event-hashchange)
    event to be fired.

9.  Set `document`{.variable}\'s [latest
    entry](#latest-entry){#navigate-non-frag-sync:latest-entry} to
    `newEntry`{.variable}.

10. Set `navigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#navigate-non-frag-sync:nav-active-history-entry-3}
    to `newEntry`{.variable}.

11. [Update the navigation API entries for a same-document
    navigation](nav-history-apis.html#update-the-navigation-api-entries-for-a-same-document-navigation){#navigate-non-frag-sync:update-the-navigation-api-entries-for-a-same-document-navigation}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigate-non-frag-sync:concept-relevant-global}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#navigate-non-frag-sync:window-navigation-api},
    `newEntry`{.variable}, and `historyHandling`{.variable}.

12. Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#navigate-non-frag-sync:nav-traversable}.

13. [Append the following session history synchronous navigation
    steps](#tn-append-session-history-sync-nav-steps){#navigate-non-frag-sync:tn-append-session-history-sync-nav-steps}
    involving `navigable`{.variable} to `traversable`{.variable}:

    1.  [Finalize a same-document
        navigation](#finalize-a-same-document-navigation){#navigate-non-frag-sync:finalize-a-same-document-navigation}
        given `traversable`{.variable}, `navigable`{.variable},
        `newEntry`{.variable}, `entryToReplace`{.variable},
        `historyHandling`{.variable}, and
        \"[`none`{#navigate-non-frag-sync:uni-none}](#uni-none)\".

    2.  Invoke [WebDriver BiDi history
        updated](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-history-updated){#navigate-non-frag-sync:webdriver-bidi-history-updated
        x-internal="webdriver-bidi-history-updated"} with
        `navigable`{.variable}.

Although both [fragment
navigation](#navigate-fragid){#navigate-non-frag-sync:navigate-fragid}
and the [URL and history update
steps](#url-and-history-update-steps){#navigate-non-frag-sync:url-and-history-update-steps-2}
perform synchronous history updates, only fragment navigation contains a
synchronous call to [update document for history step
application](#update-document-for-history-step-application){#navigate-non-frag-sync:update-document-for-history-step-application}.
The [URL and history update
steps](#url-and-history-update-steps){#navigate-non-frag-sync:url-and-history-update-steps-3}
instead perform a few select updates inside the above algorithm,
omitting others. This is somewhat of an unfortunate historical accident,
and generally leads to [web-developer
sadness](https://github.com/whatwg/html/issues/5562) about the
inconsistency. For example, this means that
[`popstate`{#navigate-non-frag-sync:event-popstate}](indices.html#event-popstate)
events fire for fragment navigations, but not for
[`history.pushState()`{#navigate-non-frag-sync:dom-history-pushstate-3}](nav-history-apis.html#dom-history-pushstate)
calls.

#### [7.4.5]{.secno} []{#history-traversal}[]{#traverse-the-history}Populating a session history entry[](#populating-a-session-history-entry){.self-link}

As explained in [the overview](#history), both
[navigation](#navigating-across-documents) and
[traversal](#reloading-and-traversing) involve creating a [session
history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry}
and then attempting to populate its
[document](#she-document){#populating-a-session-history-entry:she-document}
member, so that it can be presented inside the
[navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable}.

This involves either: using [an already-given
response](#note-navigate-called-with-response); using the [srcdoc
resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource}
stored in the [session history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-2};
or
[fetching](#create-navigation-params-by-fetching){#populating-a-session-history-entry:create-navigation-params-by-fetching}.
The process has several failure modes, which can either result in doing
nothing (leaving the
[navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable-2}
on its
currently-[active](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document}
[`Document`{#populating-a-session-history-entry:document}](dom.html#document))
or can result in populating the [session history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-3}
with an [error
document](document-lifecycle.html#read-ua-inline){#populating-a-session-history-entry:read-ua-inline}.

To [attempt to populate the history entry\'s
document]{#attempt-to-populate-the-history-entry's-document .dfn} for a
[session history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-4}
`entry`{.variable}, given a
[navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable-3}
`navigable`{.variable}, a
[`NavigationTimingType`{#populating-a-session-history-entry:navigationtimingtype}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
`navTimingType`{.variable}, a [source snapshot
params](#source-snapshot-params){#populating-a-session-history-entry:source-snapshot-params}
`sourceSnapshotParams`{.variable}, a [target snapshot
params](#target-snapshot-params){#populating-a-session-history-entry:target-snapshot-params}
`targetSnapshotParams`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#populating-a-session-history-entry:user-navigation-involvement}
`userInvolvement`{.variable}, an optional [navigation
ID](#navigation-id){#populating-a-session-history-entry:navigation-id}-or-null
`navigationId`{.variable} (default null), an optional [navigation
params](#navigation-params){#populating-a-session-history-entry:navigation-params}-or-null
`navigationParams`{.variable} (default null), an optional string
`cspNavigationType`{.variable} (default \"`other`\"), an optional
boolean [`allowPOST`{.variable}]{#attempt-to-populate-allow-post .dfn}
(default false), and optional algorithm steps
[`completionSteps`{.variable}]{#attempt-to-populate-completion-steps
.dfn} (default an empty algorithm):

1.  [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert
    x-internal="assert"}: this is running [in
    parallel](infrastructure.html#in-parallel){#populating-a-session-history-entry:in-parallel}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-2
    x-internal="assert"}: if `navigationParams`{.variable} is non-null,
    then `navigationParams`{.variable}\'s
    [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response}
    is non-null.

3.  Let `documentResource`{.variable} be `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state}\'s
    [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-2}.

4.  If `navigationParams`{.variable} is null, then:

    1.  If `documentResource`{.variable} is a string, then set
        `navigationParams`{.variable} to the result of [creating
        navigation params from a srcdoc
        resource](#create-navigation-params-from-a-srcdoc-resource){#populating-a-session-history-entry:create-navigation-params-from-a-srcdoc-resource}
        given `entry`{.variable}, `navigable`{.variable},
        `targetSnapshotParams`{.variable}, `userInvolvement`{.variable},
        `navigationId`{.variable}, and `navTimingType`{.variable}.

    2.  Otherwise, if all of the following are true:

        - `entry`{.variable}\'s
          [URL](#she-url){#populating-a-session-history-entry:she-url}\'s
          [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme
          x-internal="concept-url-scheme"} is a [fetch
          scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme
          x-internal="fetch-scheme"}; and

        - `documentResource`{.variable} is null, or
          `allowPOST`{.variable} is true and
          `documentResource`{.variable}\'s [request
          body](#post-resource-request-body){#populating-a-session-history-entry:post-resource-request-body}
          is not failure,

        then set `navigationParams`{.variable} to the result of
        [creating navigation params by
        fetching](#create-navigation-params-by-fetching){#populating-a-session-history-entry:create-navigation-params-by-fetching-2}
        given `entry`{.variable}, `navigable`{.variable},
        `sourceSnapshotParams`{.variable},
        `targetSnapshotParams`{.variable},
        `cspNavigationType`{.variable}, `userInvolvement`{.variable},
        `navigationId`{.variable}, and `navTimingType`{.variable}.

    3.  Otherwise, if `entry`{.variable}\'s
        [URL](#she-url){#populating-a-session-history-entry:she-url-2}\'s
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-2
        x-internal="concept-url-scheme"} is not a [fetch
        scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-2
        x-internal="fetch-scheme"}, then set
        `navigationParams`{.variable} to a new [non-fetch scheme
        navigation
        params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params},
        with

        [id](#non-fetch-scheme-params-id){#populating-a-session-history-entry:non-fetch-scheme-params-id}
        :   `navigationId`{.variable}

        [navigable](#non-fetch-scheme-params-navigable){#populating-a-session-history-entry:non-fetch-scheme-params-navigable}
        :   `navigable`{.variable}

        [URL](#non-fetch-scheme-params-url){#populating-a-session-history-entry:non-fetch-scheme-params-url}
        :   `entry`{.variable}\'s
            [URL](#she-url){#populating-a-session-history-entry:she-url-3}

        [target snapshot sandboxing flags](#non-fetch-scheme-params-target-sandbox){#populating-a-session-history-entry:non-fetch-scheme-params-target-sandbox}
        :   `targetSnapshotParams`{.variable}\'s [sandboxing
            flags](#target-snapshot-params-sandbox){#populating-a-session-history-entry:target-snapshot-params-sandbox}

        [source snapshot has transient activation](#non-fetch-scheme-params-source-activation){#populating-a-session-history-entry:non-fetch-scheme-params-source-activation}
        :   `sourceSnapshotParams`{.variable}\'s [has transient
            activation](#source-snapshot-params-activation){#populating-a-session-history-entry:source-snapshot-params-activation}

        [initiator origin](#non-fetch-scheme-params-initiator-origin){#populating-a-session-history-entry:non-fetch-scheme-params-initiator-origin}
        :   `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-2}\'s
            [initiator
            origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin}

        [navigation timing type](#non-fetch-scheme-params-nav-timing-type){#populating-a-session-history-entry:non-fetch-scheme-params-nav-timing-type}
        :   `navTimingType`{.variable}

        [user involvement](#non-fetch-scheme-params-user-involvement){#populating-a-session-history-entry:non-fetch-scheme-params-user-involvement}
        :   `userInvolvement`{.variable}

5.  ::: {#process-a-navigate-response}
    [Queue a global
    task](webappapis.html#queue-a-global-task){#populating-a-session-history-entry:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#populating-a-session-history-entry:navigation-and-traversal-task-source},
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#active-window){#populating-a-session-history-entry:active-window},
    to run these steps:

    1.  If `navigable`{.variable}\'s [ongoing
        navigation](#ongoing-navigation){#populating-a-session-history-entry:ongoing-navigation}
        no longer equals `navigationId`{.variable}, then run
        `completionSteps`{.variable} and abort these steps.

    2.  Let `saveExtraDocumentState`{.variable} be true.

        ::: note
        Usually, in the cases where we end up populating
        `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-3}\'s
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document},
        we then want to save some of the state from that
        [`Document`{#populating-a-session-history-entry:document-2}](dom.html#document)
        into `entry`{.variable}. This ensures that if there are future
        traversals to `entry`{.variable} where its
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-2}
        [has been destroyed](#note-bfcache), we can use that state when
        creating a new
        [`Document`{#populating-a-session-history-entry:document-3}](dom.html#document).

        However, in some specific cases, saving the state would be
        unhelpful. For those, we set `saveExtraDocumentState`{.variable}
        to false later in this algorithm.
        :::

    3.  If `navigationParams`{.variable} is a [non-fetch scheme
        navigation
        params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params-2},
        then:

        1.  Set `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-4}\'s
            [document](#document-state-document){#populating-a-session-history-entry:document-state-document-3}
            to the result of running [attempt to create a non-fetch
            scheme
            document](#attempt-to-create-a-non-fetch-scheme-document){#populating-a-session-history-entry:attempt-to-create-a-non-fetch-scheme-document}
            given `navigationParams`{.variable}.

            This can result in setting `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-5}\'s
            [document](#document-state-document){#populating-a-session-history-entry:document-state-document-4}
            to null, e.g., when [handing-off to external
            software](#hand-off-to-external-software){#populating-a-session-history-entry:hand-off-to-external-software}.

        2.  Set `saveExtraDocumentState`{.variable} to false.

    4.  Otherwise, if any of the following are true:

        - `navigationParams`{.variable} is null;

        - the result of [should navigation response to navigation
          request of type in target be blocked by Content Security
          Policy?](https://w3c.github.io/webappsec-csp/#should-block-navigation-response){#populating-a-session-history-entry:should-navigation-response-to-navigation-request-of-type-in-target-be-blocked-by-content-security-policy
          x-internal="should-navigation-response-to-navigation-request-of-type-in-target-be-blocked-by-content-security-policy"}
          given `navigationParams`{.variable}\'s
          [request](#navigation-params-request){#populating-a-session-history-entry:navigation-params-request},
          `navigationParams`{.variable}\'s
          [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-2},
          `navigationParams`{.variable}\'s [policy
          container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container}\'s
          [CSP
          list](browsers.html#policy-container-csp-list){#populating-a-session-history-entry:policy-container-csp-list},
          `cspNavigationType`{.variable}, and `navigable`{.variable} is
          \"`Blocked`\";

        - `navigationParams`{.variable}\'s [reserved
          environment](#navigation-params-reserved-environment){#populating-a-session-history-entry:navigation-params-reserved-environment}
          is non-null and the result of [checking a navigation
          response\'s adherence to its embedder
          policy](browsers.html#check-a-navigation-response's-adherence-to-its-embedder-policy){#populating-a-session-history-entry:check-a-navigation-response's-adherence-to-its-embedder-policy}
          given `navigationParams`{.variable}\'s
          [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-3},
          `navigable`{.variable}, and `navigationParams`{.variable}\'s
          [policy
          container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container-2}\'s
          [embedder
          policy](browsers.html#policy-container-embedder-policy){#populating-a-session-history-entry:policy-container-embedder-policy}
          is false; or

        - the result of [checking a navigation response\'s adherence to
          \``X-Frame-Options`\`](document-lifecycle.html#check-a-navigation-response's-adherence-to-x-frame-options){#populating-a-session-history-entry:check-a-navigation-response's-adherence-to-x-frame-options}
          given `navigationParams`{.variable}\'s
          [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-4},
          `navigable`{.variable}, `navigationParams`{.variable}\'s
          [policy
          container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container-3}\'s
          [CSP
          list](browsers.html#policy-container-csp-list){#populating-a-session-history-entry:policy-container-csp-list-2},
          and `navigationParams`{.variable}\'s
          [origin](#navigation-params-origin){#populating-a-session-history-entry:navigation-params-origin}
          is false,

        then:

        1.  Set `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-6}\'s
            [document](#document-state-document){#populating-a-session-history-entry:document-state-document-5}
            to the result of [creating a document for inline content
            that doesn\'t have a
            DOM](document-lifecycle.html#read-ua-inline){#populating-a-session-history-entry:read-ua-inline-2},
            given `navigable`{.variable}, null,
            `navTimingType`{.variable}, and
            `userInvolvement`{.variable}. The inline content should
            indicate to the user the sort of error that occurred.

        2.  [Make document
            unsalvageable](#make-document-unsalvageable){#populating-a-session-history-entry:make-document-unsalvageable}
            given `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-7}\'s
            [document](#document-state-document){#populating-a-session-history-entry:document-state-document-6}
            and
            \"[`navigation-failure`{#populating-a-session-history-entry:blocking-navigation-failure}](nav-history-apis.html#blocking-navigation-failure)\".

        3.  Set `saveExtraDocumentState`{.variable} to false.

        4.  If `navigationParams`{.variable} is not null, then:

            1.  Run the [environment discarding
                steps](webappapis.html#environment-discarding-steps){#populating-a-session-history-entry:environment-discarding-steps}
                for `navigationParams`{.variable}\'s [reserved
                environment](#navigation-params-reserved-environment){#populating-a-session-history-entry:navigation-params-reserved-environment-2}.

            2.  Invoke [WebDriver BiDi navigation
                failed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-failed){#populating-a-session-history-entry:webdriver-bidi-navigation-failed
                x-internal="webdriver-bidi-navigation-failed"} with
                `navigable`{.variable} and a new [WebDriver BiDi
                navigation
                status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#populating-a-session-history-entry:webdriver-bidi-navigation-status
                x-internal="webdriver-bidi-navigation-status"} whose
                [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#populating-a-session-history-entry:navigation-status-id
                x-internal="navigation-status-id"} is
                `navigationId`{.variable},
                [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#populating-a-session-history-entry:navigation-status-status
                x-internal="navigation-status-status"} is
                \"[`canceled`{#populating-a-session-history-entry:navigation-status-canceled}](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled){x-internal="navigation-status-canceled"}\",
                and
                [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#populating-a-session-history-entry:navigation-status-url
                x-internal="navigation-status-url"} is
                `navigationParams`{.variable}\'s
                [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-5}\'s
                [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url
                x-internal="concept-response-url"}.

    5.  ::: {#navigation-as-a-download}
        Otherwise, if `navigationParams`{.variable}\'s
        [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-6}
        has a
        \`[`Content-Disposition`{#populating-a-session-history-entry:http-content-disposition}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
        header specifying the `attachment` disposition type, then:

        1.  Let `sourceAllowsDownloading`{.variable} be
            `sourceSnapshotParams`{.variable}\'s [allows
            downloading](#source-snapshot-params-download){#populating-a-session-history-entry:source-snapshot-params-download}.

        2.  Let `targetAllowsDownloading`{.variable} be false if
            `navigationParams`{.variable}\'s [final sandboxing flag
            set](#navigation-params-sandboxing){#populating-a-session-history-entry:navigation-params-sandboxing}
            has the [sandboxed downloads browsing context
            flag](browsers.html#sandboxed-downloads-browsing-context-flag){#populating-a-session-history-entry:sandboxed-downloads-browsing-context-flag}
            set; otherwise true.

        3.  Let `uaAllowsDownloading`{.variable} be true.

        4.  Optionally, the user agent may set
            `uaAllowsDownloading`{.variable} to false, if it believes
            doing so would safeguard the user from a potentially hostile
            download.

        5.  If `sourceAllowsDownloading`{.variable},
            `targetAllowsDownloading`{.variable}, and
            `uaAllowsDownloading`{.variable} are true, then:

            1.  Let `suggestedFilename`{.variable} be the result of
                [handling as a
                download](links.html#handle-as-a-download){#populating-a-session-history-entry:handle-as-a-download}
                `navigationParams`{.variable}\'s
                [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-7}.

            2.  Invoke [WebDriver BiDi download
                started](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-download-started){#populating-a-session-history-entry:webdriver-bidi-download-started
                x-internal="webdriver-bidi-download-started"} with
                `navigable`{.variable} and a new [WebDriver BiDi
                navigation
                status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#populating-a-session-history-entry:webdriver-bidi-navigation-status-2
                x-internal="webdriver-bidi-navigation-status"} whose
                [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#populating-a-session-history-entry:navigation-status-id-2
                x-internal="navigation-status-id"} is
                `navigationId`{.variable},
                [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#populating-a-session-history-entry:navigation-status-status-2
                x-internal="navigation-status-status"} is
                \"[`complete`{#populating-a-session-history-entry:navigation-status-complete}](https://w3c.github.io/webdriver-bidi/#navigation-status-complete){x-internal="navigation-status-complete"}\",
                [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#populating-a-session-history-entry:navigation-status-url-2
                x-internal="navigation-status-url"} is
                `navigationParams`{.variable}\'s
                [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-8}\'s
                [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-2
                x-internal="concept-response-url"}, and
                [suggestedFilename](https://w3c.github.io/webdriver-bidi/#navigation-status-suggested-filename){#populating-a-session-history-entry:navigation-status-suggested-filename
                x-internal="navigation-status-suggested-filename"} is
                `suggestedFilename`{.variable}.

        This branch leaves `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-8}\'s
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-7}
        as null.
        :::

    6.  Otherwise, if `navigationParams`{.variable}\'s
        [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-9}\'s
        [status](https://fetch.spec.whatwg.org/#concept-response-status){#populating-a-session-history-entry:concept-response-status
        x-internal="concept-response-status"} is not 204 and is not 205,
        then set `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-9}\'s
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-8}
        to the result of [loading a
        document](#loading-a-document){#populating-a-session-history-entry:loading-a-document}
        given `navigationParams`{.variable},
        `sourceSnapshotParams`{.variable}, and `entry`{.variable}\'s
        [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-10}\'s
        [initiator
        origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin-2}.

        This can result in setting `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-11}\'s
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-9}
        to null, e.g., when [handing-off to external
        software](#hand-off-to-external-software){#populating-a-session-history-entry:hand-off-to-external-software-2}.

    7.  If `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-12}\'s
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-10}
        is not null, then:

        1.  Set `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-13}\'s
            [ever
            populated](#document-state-ever-populated){#populating-a-session-history-entry:document-state-ever-populated}
            to true.

        2.  If `saveExtraDocumentState`{.variable} is true:

            1.  Let `document`{.variable} be `entry`{.variable}\'s
                [document
                state](#she-document-state){#populating-a-session-history-entry:she-document-state-14}\'s
                [document](#document-state-document){#populating-a-session-history-entry:document-state-document-11}.

            2.  Set `entry`{.variable}\'s [document
                state](#she-document-state){#populating-a-session-history-entry:she-document-state-15}\'s
                [origin](#document-state-origin){#populating-a-session-history-entry:document-state-origin}
                to `document`{.variable}\'s
                [origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin
                x-internal="concept-document-origin"}.

            3.  If `document`{.variable}\'s
                [URL](https://dom.spec.whatwg.org/#concept-document-url){#populating-a-session-history-entry:the-document's-address
                x-internal="the-document's-address"} [requires storing
                the policy container in
                history](browsers.html#requires-storing-the-policy-container-in-history){#populating-a-session-history-entry:requires-storing-the-policy-container-in-history},
                then:

                1.  [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-3
                    x-internal="assert"}: `navigationParams`{.variable}
                    is a [navigation
                    params](#navigation-params){#populating-a-session-history-entry:navigation-params-2}
                    (i.e., neither null nor a [non-fetch scheme
                    navigation
                    params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params-3}).

                2.  Set `entry`{.variable}\'s [document
                    state](#she-document-state){#populating-a-session-history-entry:she-document-state-16}\'s
                    [history policy
                    container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container}
                    to `navigationParams`{.variable}\'s [policy
                    container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container-4}.

        3.  If `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-17}\'s
            [request
            referrer](#document-state-request-referrer){#populating-a-session-history-entry:document-state-request-referrer}
            is \"`client`\", and `navigationParams`{.variable} is a
            [navigation
            params](#navigation-params){#populating-a-session-history-entry:navigation-params-3}
            (i.e., neither null nor a [non-fetch scheme navigation
            params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params-4}),
            then:

            1.  [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-4
                x-internal="assert"}: `navigationParams`{.variable}\'s
                [request](#navigation-params-request){#populating-a-session-history-entry:navigation-params-request-2}
                is not null.

            2.  Set `entry`{.variable}\'s [document
                state](#she-document-state){#populating-a-session-history-entry:she-document-state-18}\'s
                [request
                referrer](#document-state-request-referrer){#populating-a-session-history-entry:document-state-request-referrer-2}
                to `navigationParams`{.variable}\'s
                [request](#navigation-params-request){#populating-a-session-history-entry:navigation-params-request-3}\'s
                [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#populating-a-session-history-entry:concept-request-referrer
                x-internal="concept-request-referrer"}.

    8.  Run `completionSteps`{.variable}.
    :::

To [create navigation params from a srcdoc
resource]{#create-navigation-params-from-a-srcdoc-resource .dfn} given a
[session history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-5}
`entry`{.variable}, a
[navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable-4}
`navigable`{.variable}, a [target snapshot
params](#target-snapshot-params){#populating-a-session-history-entry:target-snapshot-params-2}
`targetSnapshotParams`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#populating-a-session-history-entry:user-navigation-involvement-2}
`userInvolvement`{.variable}, a [navigation
ID](#navigation-id){#populating-a-session-history-entry:navigation-id-2}-or-null
`navigationId`{.variable}, and a
[`NavigationTimingType`{#populating-a-session-history-entry:navigationtimingtype-2}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
`navTimingType`{.variable}:

1.  Let `documentResource`{.variable} be `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-19}\'s
    [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-3}.

2.  Let `response`{.variable} be a new
    [response](https://fetch.spec.whatwg.org/#concept-response){#populating-a-session-history-entry:concept-response
    x-internal="concept-response"} with

    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-3 x-internal="concept-response-url"}
    :   [`about:srcdoc`{#populating-a-session-history-entry:about:srcdoc}](urls-and-fetching.html#about:srcdoc)

    [header list](https://fetch.spec.whatwg.org/#concept-response-header-list){#populating-a-session-history-entry:concept-response-header-list x-internal="concept-response-header-list"}
    :   «
        (\`[`Content-Type`{#populating-a-session-history-entry:content-type}](urls-and-fetching.html#content-type)\`,
        \``text/html`\`) »

    [body](https://fetch.spec.whatwg.org/#concept-response-body){#populating-a-session-history-entry:concept-response-body x-internal="concept-response-body"}
    :   the [UTF-8
        encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#populating-a-session-history-entry:utf-8-encode
        x-internal="utf-8-encode"} of `documentResource`{.variable}, [as
        a
        body](https://fetch.spec.whatwg.org/#byte-sequence-as-a-body){#populating-a-session-history-entry:as-a-body
        x-internal="as-a-body"}

3.  Let `responseOrigin`{.variable} be the result of [determining the
    origin](document-sequences.html#determining-the-origin){#populating-a-session-history-entry:determining-the-origin}
    given `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-4
    x-internal="concept-response-url"},
    `targetSnapshotParams`{.variable}\'s [sandboxing
    flags](#target-snapshot-params-sandbox){#populating-a-session-history-entry:target-snapshot-params-sandbox-2},
    and `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-20}\'s
    [origin](#document-state-origin){#populating-a-session-history-entry:document-state-origin-2}.

4.  Let `coop`{.variable} be a new [opener
    policy](browsers.html#cross-origin-opener-policy){#populating-a-session-history-entry:cross-origin-opener-policy}.

5.  Let `coopEnforcementResult`{.variable} be a new [opener policy
    enforcement
    result](browsers.html#coop-enforcement-result){#populating-a-session-history-entry:coop-enforcement-result}
    with

    [url](browsers.html#coop-enforcement-url){#populating-a-session-history-entry:coop-enforcement-url}
    :   `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-5
        x-internal="concept-response-url"}

    [origin](browsers.html#coop-enforcement-origin){#populating-a-session-history-entry:coop-enforcement-origin}
    :   `responseOrigin`{.variable}

    [opener policy](browsers.html#coop-enforcement-coop){#populating-a-session-history-entry:coop-enforcement-coop}
    :   `coop`{.variable}

6.  Let `policyContainer`{.variable} be the result of [determining
    navigation params policy
    container](browsers.html#determining-navigation-params-policy-container){#populating-a-session-history-entry:determining-navigation-params-policy-container}
    given `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-6
    x-internal="concept-response-url"}, `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-21}\'s
    [history policy
    container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container-2},
    null, `navigable`{.variable}\'s [container
    document](document-sequences.html#nav-container-document){#populating-a-session-history-entry:nav-container-document}\'s
    [policy
    container](dom.html#concept-document-policy-container){#populating-a-session-history-entry:concept-document-policy-container},
    and null.

7.  Return a new [navigation
    params](#navigation-params){#populating-a-session-history-entry:navigation-params-4},
    with

    [id](#navigation-params-id){#populating-a-session-history-entry:navigation-params-id}
    :   `navigationId`{.variable}

    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable}
    :   `navigable`{.variable}

    [request](#navigation-params-request){#populating-a-session-history-entry:navigation-params-request-4}
    :   null

    [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-10}
    :   `response`{.variable}

    [fetch controller](#navigation-params-fetch-controller){#populating-a-session-history-entry:navigation-params-fetch-controller}
    :   null

    [commit early hints](#navigation-params-commit-early-hints){#populating-a-session-history-entry:navigation-params-commit-early-hints}
    :   null

    [COOP enforcement result](#navigation-params-coop-enforcement-result){#populating-a-session-history-entry:navigation-params-coop-enforcement-result}
    :   `coopEnforcementResult`{.variable}

    [reserved environment](#navigation-params-reserved-environment){#populating-a-session-history-entry:navigation-params-reserved-environment-3}
    :   null

    [origin](#navigation-params-origin){#populating-a-session-history-entry:navigation-params-origin-2}
    :   `responseOrigin`{.variable}

    [policy container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container-5}
    :   `policyContainer`{.variable}

    [final sandboxing flag set](#navigation-params-sandboxing){#populating-a-session-history-entry:navigation-params-sandboxing-2}
    :   `targetSnapshotParams`{.variable}\'s [sandboxing
        flags](#target-snapshot-params-sandbox){#populating-a-session-history-entry:target-snapshot-params-sandbox-3}

    [opener policy](#navigation-params-coop){#populating-a-session-history-entry:navigation-params-coop}
    :   `coop`{.variable}

    [navigation timing type](#navigation-params-nav-timing-type){#populating-a-session-history-entry:navigation-params-nav-timing-type}
    :   `navTimingType`{.variable}

    [about base URL](#navigation-params-about-base-url){#populating-a-session-history-entry:navigation-params-about-base-url}
    :   `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-22}\'s
        [about base
        URL](#document-state-about-base-url){#populating-a-session-history-entry:document-state-about-base-url}

    [user involvement](#navigation-params-user-involvement){#populating-a-session-history-entry:navigation-params-user-involvement}
    :   `userInvolvement`{.variable}

To [create navigation params by
fetching]{#create-navigation-params-by-fetching .dfn} given a [session
history
entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-6}
`entry`{.variable}, a
[navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable-5}
`navigable`{.variable}, a [source snapshot
params](#source-snapshot-params){#populating-a-session-history-entry:source-snapshot-params-2}
`sourceSnapshotParams`{.variable}, a [target snapshot
params](#target-snapshot-params){#populating-a-session-history-entry:target-snapshot-params-3}
`targetSnapshotParams`{.variable}, a string
`cspNavigationType`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#populating-a-session-history-entry:user-navigation-involvement-3}
`userInvolvement`{.variable}, a [navigation
ID](#navigation-id){#populating-a-session-history-entry:navigation-id-3}-or-null
`navigationId`{.variable}, and a
[`NavigationTimingType`{#populating-a-session-history-entry:navigationtimingtype-3}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype){x-internal="navigationtimingtype"}
`navTimingType`{.variable}, perform the following steps. They return a
[navigation
params](#navigation-params){#populating-a-session-history-entry:navigation-params-5},
a [non-fetch scheme navigation
params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params-5},
or null.

This algorithm mutates `entry`{.variable}.

1.  [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-5
    x-internal="assert"}: this is running [in
    parallel](infrastructure.html#in-parallel){#populating-a-session-history-entry:in-parallel-2}.

2.  Let `documentResource`{.variable} be `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-23}\'s
    [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-4}.

3.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#populating-a-session-history-entry:concept-request
    x-internal="concept-request"}, with

    [url](https://fetch.spec.whatwg.org/#concept-request-url){#populating-a-session-history-entry:concept-request-url x-internal="concept-request-url"}
    :   `entry`{.variable}\'s
        [URL](#she-url){#populating-a-session-history-entry:she-url-4}

    [client](https://fetch.spec.whatwg.org/#concept-request-client){#populating-a-session-history-entry:concept-request-client x-internal="concept-request-client"}
    :   `sourceSnapshotParams`{.variable}\'s [fetch
        client](#source-snapshot-params-client){#populating-a-session-history-entry:source-snapshot-params-client}

    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#populating-a-session-history-entry:concept-request-destination x-internal="concept-request-destination"}
    :   \"`document`\" [The destination is updated below when
        `navigable`{.variable} has a
        [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container}.]{.note}

    [credentials mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#populating-a-session-history-entry:concept-request-credentials-mode x-internal="concept-request-credentials-mode"}
    :   \"`include`\"

    [use-URL-credentials flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#populating-a-session-history-entry:use-url-credentials-flag x-internal="use-url-credentials-flag"}
    :   set

    [redirect mode](https://fetch.spec.whatwg.org/#concept-request-redirect-mode){#populating-a-session-history-entry:concept-request-redirect-mode x-internal="concept-request-redirect-mode"}
    :   \"`manual`\"

    [replaces client id](https://fetch.spec.whatwg.org/#concept-request-replaces-client-id){#populating-a-session-history-entry:concept-request-replaces-client-id x-internal="concept-request-replaces-client-id"}
    :   `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-2}\'s
        [relevant settings
        object](webappapis.html#relevant-settings-object){#populating-a-session-history-entry:relevant-settings-object}\'s
        [id](webappapis.html#concept-environment-id){#populating-a-session-history-entry:concept-environment-id}

    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#populating-a-session-history-entry:concept-request-mode x-internal="concept-request-mode"}
    :   \"`navigate`\"

    [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#populating-a-session-history-entry:concept-request-referrer-2 x-internal="concept-request-referrer"}
    :   `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-24}\'s
        [request
        referrer](#document-state-request-referrer){#populating-a-session-history-entry:document-state-request-referrer-3}

    [referrer policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#populating-a-session-history-entry:concept-request-referrer-policy x-internal="concept-request-referrer-policy"}
    :   `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-25}\'s
        [request referrer
        policy](#document-state-request-referrer-policy){#populating-a-session-history-entry:document-state-request-referrer-policy}

    [policy container](https://fetch.spec.whatwg.org/#concept-request-policy-container){#populating-a-session-history-entry:concept-request-policy-container x-internal="concept-request-policy-container"}
    :   `sourceSnapshotParams`{.variable}\'s [source policy
        container](#source-snapshot-params-policy-container){#populating-a-session-history-entry:source-snapshot-params-policy-container}

    [traversable for user prompts](https://fetch.spec.whatwg.org/#concept-request-window){#populating-a-session-history-entry:concept-request-traversable-for-user-prompts x-internal="concept-request-traversable-for-user-prompts"}
    :   `navigable`{.variable}\'s [top-level
        traversable](document-sequences.html#nav-top){#populating-a-session-history-entry:nav-top}

4.  If `request`{.variable}\'s
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#populating-a-session-history-entry:concept-request-client-2
    x-internal="concept-request-client"} is null:

    This [only occurs](#assert-null-sourcedocument) in the case of a
    browser UI-initiated navigation.

    1.  Set `request`{.variable}\'s
        [origin](https://fetch.spec.whatwg.org/#concept-request-origin){#populating-a-session-history-entry:concept-request-origin
        x-internal="concept-request-origin"} to a new [opaque
        origin](browsers.html#concept-origin-opaque){#populating-a-session-history-entry:concept-origin-opaque}.

    2.  Set `request`{.variable}\'s [service-workers
        mode](https://fetch.spec.whatwg.org/#request-service-workers-mode){#populating-a-session-history-entry:concept-request-service-workers-mode
        x-internal="concept-request-service-workers-mode"} to \"`all`\".

    3.  Set `request`{.variable}\'s
        [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#populating-a-session-history-entry:concept-request-referrer-3
        x-internal="concept-request-referrer"} to \"`no-referrer`\".

5.  If `documentResource`{.variable} is a [POST
    resource](#post-resource){#populating-a-session-history-entry:post-resource}:

    1.  Set `request`{.variable}\'s
        [method](https://fetch.spec.whatwg.org/#concept-request-method){#populating-a-session-history-entry:concept-request-method
        x-internal="concept-request-method"} to \``POST`\`.

    2.  Set `request`{.variable}\'s
        [body](https://fetch.spec.whatwg.org/#concept-request-body){#populating-a-session-history-entry:concept-request-body
        x-internal="concept-request-body"} to
        `documentResource`{.variable}\'s [request
        body](#post-resource-request-body){#populating-a-session-history-entry:post-resource-request-body-2}.

    3.  [Set](https://fetch.spec.whatwg.org/#concept-header-list-set){#populating-a-session-history-entry:concept-header-list-set
        x-internal="concept-header-list-set"} \``Content-Type`\` to
        `documentResource`{.variable}\'s [request
        content-type](#post-resource-request-content-type){#populating-a-session-history-entry:post-resource-request-content-type}
        in `request`{.variable}\'s [header
        list](https://fetch.spec.whatwg.org/#concept-request-header-list){#populating-a-session-history-entry:concept-request-header-list
        x-internal="concept-request-header-list"}.

6.  If `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-26}\'s
    [reload
    pending](#document-state-reload-pending){#populating-a-session-history-entry:document-state-reload-pending}
    is true, then set `request`{.variable}\'s [reload-navigation
    flag](https://fetch.spec.whatwg.org/#concept-request-reload-navigation-flag){#populating-a-session-history-entry:concept-request-reload-navigation-flag
    x-internal="concept-request-reload-navigation-flag"}.

7.  Otherwise, if `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-27}\'s
    [ever
    populated](#document-state-ever-populated){#populating-a-session-history-entry:document-state-ever-populated-2}
    is true, then set `request`{.variable}\'s [history-navigation
    flag](https://fetch.spec.whatwg.org/#concept-request-history-navigation-flag){#populating-a-session-history-entry:concept-request-history-navigation-flag
    x-internal="concept-request-history-navigation-flag"}.

8.  If `sourceSnapshotParams`{.variable}\'s [has transient
    activation](#source-snapshot-params-activation){#populating-a-session-history-entry:source-snapshot-params-activation-2}
    is true, then set `request`{.variable}\'s
    [user-activation](https://fetch.spec.whatwg.org/#request-user-activation){#populating-a-session-history-entry:concept-request-user-activation
    x-internal="concept-request-user-activation"} to true.

9.  If `navigable`{.variable}\'s
    [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-2}
    is non-null:

    1.  If the `navigable`{.variable}\'s
        [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-3}
        has a [browsing context scope
        origin](#browsing-context-scope-origin){#populating-a-session-history-entry:browsing-context-scope-origin},
        then set `request`{.variable}\'s
        [origin](https://fetch.spec.whatwg.org/#concept-request-origin){#populating-a-session-history-entry:concept-request-origin-2
        x-internal="concept-request-origin"} to that [browsing context
        scope
        origin](#browsing-context-scope-origin){#populating-a-session-history-entry:browsing-context-scope-origin-2}.

    2.  Set `request`{.variable}\'s
        [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#populating-a-session-history-entry:concept-request-destination-2
        x-internal="concept-request-destination"} to
        `navigable`{.variable}\'s
        [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-4}\'s
        [local
        name](https://dom.spec.whatwg.org/#concept-element-local-name){#populating-a-session-history-entry:concept-element-local-name
        x-internal="concept-element-local-name"}.

    3.  If `sourceSnapshotParams`{.variable}\'s [fetch
        client](#source-snapshot-params-client){#populating-a-session-history-entry:source-snapshot-params-client-2}
        is `navigable`{.variable}\'s [container
        document](document-sequences.html#nav-container-document){#populating-a-session-history-entry:nav-container-document-2}\'s
        [relevant settings
        object](webappapis.html#relevant-settings-object){#populating-a-session-history-entry:relevant-settings-object-2},
        then set `request`{.variable}\'s [initiator
        type](https://fetch.spec.whatwg.org/#request-initiator-type){#populating-a-session-history-entry:concept-request-initiator-type
        x-internal="concept-request-initiator-type"} to
        `navigable`{.variable}\'s
        [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-5}\'s
        [local
        name](https://dom.spec.whatwg.org/#concept-element-local-name){#populating-a-session-history-entry:concept-element-local-name-2
        x-internal="concept-element-local-name"}.

        This ensure that only container-initiated navigations are
        reported to resource timing.

10. Let `response`{.variable} be null.

11. Let `responseOrigin`{.variable} be null.

12. Let `fetchController`{.variable} be null.

13. Let `coopEnforcementResult`{.variable} be a new [opener policy
    enforcement
    result](browsers.html#coop-enforcement-result){#populating-a-session-history-entry:coop-enforcement-result-2},
    with

    [url](browsers.html#coop-enforcement-url){#populating-a-session-history-entry:coop-enforcement-url-2}
    :   `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-3}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#populating-a-session-history-entry:the-document's-address-2
        x-internal="the-document's-address"}

    [origin](browsers.html#coop-enforcement-origin){#populating-a-session-history-entry:coop-enforcement-origin-2}
    :   `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-4}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-2
        x-internal="concept-document-origin"}

    [opener policy](browsers.html#coop-enforcement-coop){#populating-a-session-history-entry:coop-enforcement-coop-2}
    :   `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-5}\'s
        [opener
        policy](dom.html#concept-document-coop){#populating-a-session-history-entry:concept-document-coop}

    [current context is navigation source](browsers.html#coop-enforcement-source){#populating-a-session-history-entry:coop-enforcement-source}
    :   true if `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-6}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-3
        x-internal="concept-document-origin"} is [same
        origin](browsers.html#same-origin){#populating-a-session-history-entry:same-origin}
        with `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-28}\'s
        [initiator
        origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin-3}
        otherwise false

14. Let `finalSandboxFlags`{.variable} be an empty [sandboxing flag
    set](browsers.html#sandboxing-flag-set){#populating-a-session-history-entry:sandboxing-flag-set}.

15. Let `responsePolicyContainer`{.variable} be null.

16. Let `responseCOOP`{.variable} be a new [opener
    policy](browsers.html#cross-origin-opener-policy){#populating-a-session-history-entry:cross-origin-opener-policy-2}.

17. Let `locationURL`{.variable} be null.

18. Let `currentURL`{.variable} be `request`{.variable}\'s [current
    URL](https://fetch.spec.whatwg.org/#concept-request-current-url){#populating-a-session-history-entry:concept-request-current-url
    x-internal="concept-request-current-url"}.

19. Let `commitEarlyHints`{.variable} be null.

20. While true:

    1.  If `request`{.variable}\'s [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client
        x-internal="concept-request-reserved-client"} is not null and
        `currentURL`{.variable}\'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#populating-a-session-history-entry:concept-url-origin
        x-internal="concept-url-origin"} is not the
        [same](browsers.html#same-origin){#populating-a-session-history-entry:same-origin-2}
        as `request`{.variable}\'s [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-2
        x-internal="concept-request-reserved-client"}\'s [creation
        URL](webappapis.html#concept-environment-creation-url){#populating-a-session-history-entry:concept-environment-creation-url}\'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#populating-a-session-history-entry:concept-url-origin-2
        x-internal="concept-url-origin"}, then:

        1.  Run the [environment discarding
            steps](webappapis.html#environment-discarding-steps){#populating-a-session-history-entry:environment-discarding-steps-2}
            for `request`{.variable}\'s [reserved
            client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-3
            x-internal="concept-request-reserved-client"}.

        2.  Set `request`{.variable}\'s [reserved
            client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-4
            x-internal="concept-request-reserved-client"} to null.

        3.  Set `commitEarlyHints`{.variable} to null.

            Preloaded links from [early hint
            headers](semantics.html#early-hints-2){#populating-a-session-history-entry:early-hints-2}
            remain in the preload cache after a [same
            origin](browsers.html#same-origin){#populating-a-session-history-entry:same-origin-3}
            redirect, but get discarded when the redirect is
            cross-origin.

    2.  If `request`{.variable}\'s [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-5
        x-internal="concept-request-reserved-client"} is null, then:

        1.  Let `topLevelCreationURL`{.variable} be
            `currentURL`{.variable}.

        2.  Let `topLevelOrigin`{.variable} be null.

        3.  If `navigable`{.variable} is not a [top-level
            traversable](document-sequences.html#top-level-traversable){#populating-a-session-history-entry:top-level-traversable},
            then:

            1.  Let `parentEnvironment`{.variable} be
                `navigable`{.variable}\'s
                [parent](document-sequences.html#nav-parent){#populating-a-session-history-entry:nav-parent}\'s
                [active
                document](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-7}\'s
                [relevant settings
                object](webappapis.html#relevant-settings-object){#populating-a-session-history-entry:relevant-settings-object-3}.

            2.  Set `topLevelCreationURL`{.variable} to
                `parentEnvironment`{.variable}\'s [top-level creation
                URL](webappapis.html#concept-environment-top-level-creation-url){#populating-a-session-history-entry:concept-environment-top-level-creation-url}.

            3.  Set `topLevelOrigin`{.variable} to
                `parentEnvironment`{.variable}\'s [top-level
                origin](webappapis.html#concept-environment-top-level-origin){#populating-a-session-history-entry:concept-environment-top-level-origin}.

        4.  Set `request`{.variable}\'s [reserved
            client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-6
            x-internal="concept-request-reserved-client"} to a new
            [environment](webappapis.html#environment){#populating-a-session-history-entry:environment}
            whose
            [id](webappapis.html#concept-environment-id){#populating-a-session-history-entry:concept-environment-id-2}
            is a unique opaque string, [target browsing
            context](webappapis.html#concept-environment-target-browsing-context){#populating-a-session-history-entry:concept-environment-target-browsing-context}
            is `navigable`{.variable}\'s [active browsing
            context](document-sequences.html#nav-bc){#populating-a-session-history-entry:nav-bc},
            [creation
            URL](webappapis.html#concept-environment-creation-url){#populating-a-session-history-entry:concept-environment-creation-url-2}
            is `currentURL`{.variable}, [top-level creation
            URL](webappapis.html#concept-environment-top-level-creation-url){#populating-a-session-history-entry:concept-environment-top-level-creation-url-2}
            is `topLevelCreationURL`{.variable}, and [top-level
            origin](webappapis.html#concept-environment-top-level-origin){#populating-a-session-history-entry:concept-environment-top-level-origin-2}
            is `topLevelOrigin`{.variable}.

            The created environment\'s [active service
            worker](webappapis.html#concept-environment-active-service-worker){#populating-a-session-history-entry:concept-environment-active-service-worker}
            is set in the [Handle
            Fetch](https://w3c.github.io/ServiceWorker/#on-fetch-request-algorithm){#populating-a-session-history-entry:on-fetch-request-algorithm
            x-internal="on-fetch-request-algorithm"} algorithm during
            the fetch if the request URL matches a service worker
            registration. [\[SW\]](references.html#refsSW)

    3.  If the result of [should navigation request of type be blocked
        by Content Security
        Policy?](https://w3c.github.io/webappsec-csp/#should-block-navigation-request){#populating-a-session-history-entry:should-navigation-request-of-type-be-blocked-by-content-security-policy
        x-internal="should-navigation-request-of-type-be-blocked-by-content-security-policy"}
        given `request`{.variable} and `cspNavigationType`{.variable} is
        \"`Blocked`\", then set `response`{.variable} to a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error
        x-internal="network-error"} and
        [break](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break
        x-internal="break"}. [\[CSP\]](references.html#refsCSP)

    4.  Set `response`{.variable} to null.

    5.  If `fetchController`{.variable} is null, then set
        `fetchController`{.variable} to the result of
        [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#populating-a-session-history-entry:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}, with
        *[processEarlyHintsResponse](https://fetch.spec.whatwg.org/#fetch-processearlyhintsresponse){x-internal="processearlyhintsresponse"}*
        set to `processEarlyHintsResponse`{.variable} as defined below,
        *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
        set to `processResponse`{.variable} as defined below, and
        *[useParallelQueue](https://fetch.spec.whatwg.org/#fetch-useparallelqueue){x-internal="useparallelqueue"}*
        set to true.

        Let `processEarlyHintsResponse`{.variable} be the following
        algorithm given a
        [response](https://fetch.spec.whatwg.org/#concept-response){#populating-a-session-history-entry:concept-response-2
        x-internal="concept-response"} `earlyResponse`{.variable}:

        1.  If `commitEarlyHints`{.variable} is null, then set
            `commitEarlyHints`{.variable} to the result of [processing
            early hint
            headers](semantics.html#process-early-hint-headers){#populating-a-session-history-entry:process-early-hint-headers}
            given `earlyResponse`{.variable} and `request`{.variable}\'s
            [reserved
            client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-7
            x-internal="concept-request-reserved-client"}.

        Let `processResponse`{.variable} be the following algorithm
        given a
        [response](https://fetch.spec.whatwg.org/#concept-response){#populating-a-session-history-entry:concept-response-3
        x-internal="concept-response"} `fetchedResponse`{.variable}:

        1.  Set `response`{.variable} to `fetchedResponse`{.variable}.

    6.  Otherwise, [process the next manual
        redirect](https://fetch.spec.whatwg.org/#fetch-controller-process-the-next-manual-redirect){#populating-a-session-history-entry:process-the-next-manual-redirect
        x-internal="process-the-next-manual-redirect"} for
        `fetchController`{.variable}.

        This will result in calling the
        *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
        we supplied above, during our first iteration through the loop,
        and thus setting `response`{.variable}.

        Navigation handles redirects manually as navigation is the only
        place in the web platform that cares for redirects to
        [`mailto:`{#populating-a-session-history-entry:mailto-protocol}](https://www.rfc-editor.org/rfc/rfc6068#section-2){x-internal="mailto-protocol"}
        URLs and such.

    7.  Wait until either `response`{.variable} is non-null, or
        `navigable`{.variable}\'s [ongoing
        navigation](#ongoing-navigation){#populating-a-session-history-entry:ongoing-navigation-2}
        changes to no longer equal `navigationId`{.variable}.

        If the latter condition occurs, then
        [abort](https://fetch.spec.whatwg.org/#fetch-controller-abort){#populating-a-session-history-entry:fetch-controller-abort
        x-internal="fetch-controller-abort"}
        `fetchController`{.variable}, and return.

        Otherwise, proceed onward.

    8.  If `request`{.variable}\'s
        [body](https://fetch.spec.whatwg.org/#concept-request-body){#populating-a-session-history-entry:concept-request-body-2
        x-internal="concept-request-body"} is null, then set
        `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-29}\'s
        [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-5}
        to null.

        Fetch unsets the
        [body](https://fetch.spec.whatwg.org/#concept-request-body){#populating-a-session-history-entry:concept-request-body-3
        x-internal="concept-request-body"} for particular redirects.

    9.  Set `responsePolicyContainer`{.variable} to the result of
        [creating a policy container from a fetch
        response](browsers.html#creating-a-policy-container-from-a-fetch-response){#populating-a-session-history-entry:creating-a-policy-container-from-a-fetch-response}
        given `response`{.variable} and `request`{.variable}\'s
        [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-8
        x-internal="concept-request-reserved-client"}.

    10. Set `finalSandboxFlags`{.variable} to the
        [union](https://infra.spec.whatwg.org/#set-union){#populating-a-session-history-entry:set-union
        x-internal="set-union"} of `targetSnapshotParams`{.variable}\'s
        [sandboxing
        flags](#target-snapshot-params-sandbox){#populating-a-session-history-entry:target-snapshot-params-sandbox-4}
        and `responsePolicyContainer`{.variable}\'s [CSP
        list](browsers.html#policy-container-csp-list){#populating-a-session-history-entry:policy-container-csp-list-3}\'s
        [CSP-derived sandboxing
        flags](browsers.html#csp-derived-sandboxing-flags){#populating-a-session-history-entry:csp-derived-sandboxing-flags}.

    11. Set `responseOrigin`{.variable} to the result of [determining
        the
        origin](document-sequences.html#determining-the-origin){#populating-a-session-history-entry:determining-the-origin-2}
        given `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-7
        x-internal="concept-response-url"},
        `finalSandboxFlags`{.variable}, and `entry`{.variable}\'s
        [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-30}\'s
        [initiator
        origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin-4}.

        If `response`{.variable} is a redirect, then
        `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-8
        x-internal="concept-response-url"} will be the URL that led to
        the redirect to `response`{.variable}\'s [location
        URL](https://fetch.spec.whatwg.org/#concept-response-location-url){#populating-a-session-history-entry:concept-response-location-url
        x-internal="concept-response-location-url"}; it will not be the
        [location
        URL](https://fetch.spec.whatwg.org/#concept-response-location-url){#populating-a-session-history-entry:concept-response-location-url-2
        x-internal="concept-response-location-url"} itself.

    12. If `navigable`{.variable} is a [top-level
        traversable](document-sequences.html#top-level-traversable){#populating-a-session-history-entry:top-level-traversable-2},
        then:

        1.  Set `responseCOOP`{.variable} to the result of [obtaining an
            opener
            policy](browsers.html#obtain-coop){#populating-a-session-history-entry:obtain-coop}
            given `response`{.variable} and `request`{.variable}\'s
            [reserved
            client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-9
            x-internal="concept-request-reserved-client"}.

        2.  Set `coopEnforcementResult`{.variable} to the result of
            [enforcing the response\'s opener
            policy](browsers.html#coop-enforce){#populating-a-session-history-entry:coop-enforce}
            given `navigable`{.variable}\'s [active browsing
            context](document-sequences.html#nav-bc){#populating-a-session-history-entry:nav-bc-2},
            `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-9
            x-internal="concept-response-url"},
            `responseOrigin`{.variable}, `responseCOOP`{.variable},
            `coopEnforcementResult`{.variable}, and
            `request`{.variable}\'s
            [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#populating-a-session-history-entry:concept-request-referrer-4
            x-internal="concept-request-referrer"}.

        3.  If `finalSandboxFlags`{.variable} is not empty and
            `responseCOOP`{.variable}\'s
            [value](browsers.html#coop-struct-value){#populating-a-session-history-entry:coop-struct-value}
            is not
            \"[`unsafe-none`{#populating-a-session-history-entry:coop-unsafe-none}](browsers.html#coop-unsafe-none)\",
            then set `response`{.variable} to an appropriate [network
            error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-2
            x-internal="network-error"} and
            [break](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break-2
            x-internal="break"}.

            This results in a network error as one cannot simultaneously
            provide a clean slate to a response using opener policy and
            sandbox the result of navigating to that response.

    13. If `response`{.variable} is not a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-3
        x-internal="network-error"}, `navigable`{.variable} is a [child
        navigable](document-sequences.html#child-navigable){#populating-a-session-history-entry:child-navigable},
        and the result of performing a [cross-origin resource policy
        check](https://fetch.spec.whatwg.org/#cross-origin-resource-policy-check){#populating-a-session-history-entry:cross-origin-resource-policy-check
        x-internal="cross-origin-resource-policy-check"} with
        `navigable`{.variable}\'s [container
        document](document-sequences.html#nav-container-document){#populating-a-session-history-entry:nav-container-document-3}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-4
        x-internal="concept-document-origin"}, `navigable`{.variable}\'s
        [container
        document](document-sequences.html#nav-container-document){#populating-a-session-history-entry:nav-container-document-4}\'s
        [relevant settings
        object](webappapis.html#relevant-settings-object){#populating-a-session-history-entry:relevant-settings-object-4},
        `request`{.variable}\'s
        [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#populating-a-session-history-entry:concept-request-destination-3
        x-internal="concept-request-destination"},
        `response`{.variable}, and true is **blocked**, then set
        `response`{.variable} to a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-4
        x-internal="network-error"} and
        [break](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break-3
        x-internal="break"}.

        Here we\'re running the [cross-origin resource policy
        check](https://fetch.spec.whatwg.org/#cross-origin-resource-policy-check){#populating-a-session-history-entry:cross-origin-resource-policy-check-2
        x-internal="cross-origin-resource-policy-check"} against the
        [parent
        navigable](document-sequences.html#nav-parent){#populating-a-session-history-entry:nav-parent-2}
        rather than `navigable`{.variable} itself. This is because we
        care about the same-originness of the embedded content against
        the parent context, not the navigation source.

    14. Set `locationURL`{.variable} to `response`{.variable}\'s
        [location
        URL](https://fetch.spec.whatwg.org/#concept-response-location-url){#populating-a-session-history-entry:concept-response-location-url-3
        x-internal="concept-response-location-url"} given
        `currentURL`{.variable}\'s
        [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#populating-a-session-history-entry:concept-url-fragment
        x-internal="concept-url-fragment"}.

    15. If `locationURL`{.variable} is failure or null, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break-4
        x-internal="break"}.

    16. [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-6
        x-internal="assert"}: `locationURL`{.variable} is a
        [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url
        x-internal="url"}.

    17. Set `entry`{.variable}\'s [classic history API
        state](#she-classic-history-api-state){#populating-a-session-history-entry:she-classic-history-api-state}
        to
        [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#populating-a-session-history-entry:structuredserializeforstorage}(null).

    18. Let `oldDocState`{.variable} be `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-31}.

    19. Set `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-32}
        to a new [document
        state](#document-state-2){#populating-a-session-history-entry:document-state-2},
        with

        [history policy container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container-3}
        :   a
            [clone](browsers.html#clone-a-policy-container){#populating-a-session-history-entry:clone-a-policy-container}
            of the `oldDocState`{.variable}\'s [history policy
            container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container-4}
            if it is non-null; null otherwise

        [request referrer](#document-state-request-referrer){#populating-a-session-history-entry:document-state-request-referrer-4}
        :   `oldDocState`{.variable}\'s [request
            referrer](#document-state-request-referrer){#populating-a-session-history-entry:document-state-request-referrer-5}

        [request referrer policy](#document-state-request-referrer-policy){#populating-a-session-history-entry:document-state-request-referrer-policy-2}
        :   `oldDocState`{.variable}\'s [request referrer
            policy](#document-state-request-referrer-policy){#populating-a-session-history-entry:document-state-request-referrer-policy-3}

        [initiator origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin-5}
        :   `oldDocState`{.variable}\'s [initiator
            origin](#document-state-initiator-origin){#populating-a-session-history-entry:document-state-initiator-origin-6}

        [origin](#document-state-origin){#populating-a-session-history-entry:document-state-origin-3}
        :   `oldDocState`{.variable}\'s
            [origin](#document-state-origin){#populating-a-session-history-entry:document-state-origin-4}

        [about base URL](#document-state-about-base-url){#populating-a-session-history-entry:document-state-about-base-url-2}
        :   `oldDocState`{.variable}\'s [about base
            URL](#document-state-about-base-url){#populating-a-session-history-entry:document-state-about-base-url-3}

        [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-6}
        :   `oldDocState`{.variable}\'s
            [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-7}

        [ever populated](#document-state-ever-populated){#populating-a-session-history-entry:document-state-ever-populated-3}
        :   `oldDocState`{.variable}\'s [ever
            populated](#document-state-ever-populated){#populating-a-session-history-entry:document-state-ever-populated-4}

        [navigable target name](#document-state-nav-target-name){#populating-a-session-history-entry:document-state-nav-target-name}
        :   `oldDocState`{.variable}\'s [navigable target
            name](#document-state-nav-target-name){#populating-a-session-history-entry:document-state-nav-target-name-2}

        For the navigation case, only `entry`{.variable} referenced
        `oldDocState`{.variable}, which was created [early in the
        navigate algorithm](#navigation-create-document-state). So for
        navigations, this is functionally just an update to
        `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-33}.
        For the traversal case, it\'s possible adjacent [session history
        entries](#session-history-entry){#populating-a-session-history-entry:session-history-entry-7}
        also reference `oldDocState`{.variable}, in which case they will
        continue doing so even after we\'ve updated
        `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-34}.

        `oldDocState`{.variable}\'s [history policy
        container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container-5}
        is only ever non-null here in the traversal case, after we\'ve
        populated it during a navigation to a URL that [requires storing
        the policy container in
        history](browsers.html#requires-storing-the-policy-container-in-history){#populating-a-session-history-entry:requires-storing-the-policy-container-in-history-2}.

        ::: example
        The setup is given by the following [Jake
        diagram](document-sequences.html#jake-diagram){#populating-a-session-history-entry:jake-diagram}:

        0
        1
        2
        3
        `top`
        /a
        /a#foo
        /a#bar
        /b
        Also assume that the [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-35}
        shared by the entries in steps 0, 1, and 2 has a null
        [document](#document-state-document){#populating-a-session-history-entry:document-state-document-12},
        i.e., [bfcache](#note-bfcache) is not in play.

        Now consider the scenario where we traverse back to step 2, but
        this time when fetching `/a`, the server responds with a
        \``Location`\` header pointing to `/c`. That is,
        `locationURL`{.variable} points to `/c` and so we have reached
        this step instead of
        [breaking](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break-5
        x-internal="break"} out of the loop.

        In this case, we replace the [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-36}
        of the [session history
        entry](#session-history-entry){#populating-a-session-history-entry:session-history-entry-8}
        occupying step 2, but we do *not* replace the document state of
        the entries occupying steps 0 and 1. The resulting [Jake
        diagram](document-sequences.html#jake-diagram){#populating-a-session-history-entry:jake-diagram-2}
        looks like this:

        0
        1
        2
        3
        `top`
        /a
        /a#foo
        /c#bar
        /b
        Note that we perform this replacement even if we end up in a
        redirect chain back to the original URL, for example if `/c`
        itself had a \``Location`\` header pointing to `/a`. Such a case
        would end up like so:

        0
        1
        2
        3
        `top`
        /a
        /a#foo
        /a#bar
        /b
        :::

    20. If `locationURL`{.variable}\'s
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-3
        x-internal="concept-url-scheme"} is not an [HTTP(S)
        scheme](https://fetch.spec.whatwg.org/#http-scheme){#populating-a-session-history-entry:http(s)-scheme
        x-internal="http(s)-scheme"}, then:

        1.  Set `entry`{.variable}\'s [document
            state](#she-document-state){#populating-a-session-history-entry:she-document-state-37}\'s
            [resource](#document-state-resource){#populating-a-session-history-entry:document-state-resource-8}
            to null.

        2.  [Break](https://infra.spec.whatwg.org/#iteration-break){#populating-a-session-history-entry:break-6
            x-internal="break"}.

    21. Set `currentURL`{.variable} to `locationURL`{.variable}.

    22. Set `entry`{.variable}\'s
        [URL](#she-url){#populating-a-session-history-entry:she-url-5}
        to `currentURL`{.variable}.

    ::: note
    By the end of this loop we will be in one of these scenarios:

    - `locationURL`{.variable} is failure, because of an unparseable
      \``Location`\` header.

    - `locationURL`{.variable} is null, either because
      `response`{.variable} is a [network
      error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-5
      x-internal="network-error"} or because we successfully fetched a
      non-[network
      error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-6
      x-internal="network-error"} HTTP(S) response with no
      \``Location`\` header.

    - `locationURL`{.variable} is a
      [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-2
      x-internal="url"} with a
      non-[HTTP(S)](https://fetch.spec.whatwg.org/#http-scheme){#populating-a-session-history-entry:http(s)-scheme-2
      x-internal="http(s)-scheme"}
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-4
      x-internal="concept-url-scheme"}.
    :::

21. If `locationURL`{.variable} is a
    [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-3
    x-internal="url"} whose
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-5
    x-internal="concept-url-scheme"} is not a [fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-3
    x-internal="fetch-scheme"}, then return a new [non-fetch scheme
    navigation
    params](#non-fetch-scheme-navigation-params){#populating-a-session-history-entry:non-fetch-scheme-navigation-params-6},
    with

    [id](#non-fetch-scheme-params-id){#populating-a-session-history-entry:non-fetch-scheme-params-id-2}
    :   `navigationId`{.variable}

    [navigable](#non-fetch-scheme-params-navigable){#populating-a-session-history-entry:non-fetch-scheme-params-navigable-2}
    :   `navigable`{.variable}

    [URL](#non-fetch-scheme-params-url){#populating-a-session-history-entry:non-fetch-scheme-params-url-2}
    :   `locationURL`{.variable}

    [target snapshot sandboxing flags](#non-fetch-scheme-params-target-sandbox){#populating-a-session-history-entry:non-fetch-scheme-params-target-sandbox-2}
    :   `targetSnapshotParams`{.variable}\'s [sandboxing
        flags](#target-snapshot-params-sandbox){#populating-a-session-history-entry:target-snapshot-params-sandbox-5}

    [source snapshot has transient activation](#non-fetch-scheme-params-source-activation){#populating-a-session-history-entry:non-fetch-scheme-params-source-activation-2}
    :   `sourceSnapshotParams`{.variable}\'s [has transient
        activation](#source-snapshot-params-activation){#populating-a-session-history-entry:source-snapshot-params-activation-3}

    [initiator origin](#non-fetch-scheme-params-initiator-origin){#populating-a-session-history-entry:non-fetch-scheme-params-initiator-origin-2}
    :   `responseOrigin`{.variable}

    [navigation timing type](#non-fetch-scheme-params-nav-timing-type){#populating-a-session-history-entry:non-fetch-scheme-params-nav-timing-type-2}
    :   `navTimingType`{.variable}

    [user involvement](#non-fetch-scheme-params-user-involvement){#populating-a-session-history-entry:non-fetch-scheme-params-user-involvement-2}
    :   `userInvolvement`{.variable}

    At this point, `request`{.variable}\'s [current
    URL](https://fetch.spec.whatwg.org/#concept-request-current-url){#populating-a-session-history-entry:concept-request-current-url-2
    x-internal="concept-request-current-url"} is the last
    [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-4
    x-internal="url"} in the redirect chain with a
    [fetch](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-4
    x-internal="fetch-scheme"}
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-6
    x-internal="concept-url-scheme"} before redirecting to a non-[fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-5
    x-internal="fetch-scheme"}
    [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-5
    x-internal="url"}. It is this
    [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-6
    x-internal="url"}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#populating-a-session-history-entry:concept-url-origin-3
    x-internal="concept-url-origin"} that will be used as the initiator
    origin for navigations to non-[fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-6
    x-internal="fetch-scheme"}
    [URLs](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-7
    x-internal="url"}.

22. If any of the following are true:

    - `response`{.variable} is a [network
      error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-7
      x-internal="network-error"};

    - `locationURL`{.variable} is failure; or

    - `locationURL`{.variable} is a
      [URL](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-8
      x-internal="url"} whose
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#populating-a-session-history-entry:concept-url-scheme-7
      x-internal="concept-url-scheme"} is a [fetch
      scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-7
      x-internal="fetch-scheme"},

    then return null.

    We allow redirects to non-[fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-8
    x-internal="fetch-scheme"}
    [URLs](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-9
    x-internal="url"}, but redirects to [fetch
    scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#populating-a-session-history-entry:fetch-scheme-9
    x-internal="fetch-scheme"}
    [URLs](https://url.spec.whatwg.org/#concept-url){#populating-a-session-history-entry:url-10
    x-internal="url"} that aren\'t
    [HTTP(S)](https://fetch.spec.whatwg.org/#http-scheme){#populating-a-session-history-entry:http(s)-scheme-3
    x-internal="http(s)-scheme"} are treated like network errors.

23. [Assert](https://infra.spec.whatwg.org/#assert){#populating-a-session-history-entry:assert-7
    x-internal="assert"}: `locationURL`{.variable} is null and
    `response`{.variable} is not a [network
    error](https://fetch.spec.whatwg.org/#concept-network-error){#populating-a-session-history-entry:network-error-8
    x-internal="network-error"}.

24. Let `resultPolicyContainer`{.variable} be the result of [determining
    navigation params policy
    container](browsers.html#determining-navigation-params-policy-container){#populating-a-session-history-entry:determining-navigation-params-policy-container-2}
    given `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#populating-a-session-history-entry:concept-response-url-10
    x-internal="concept-response-url"}, `entry`{.variable}\'s [document
    state](#she-document-state){#populating-a-session-history-entry:she-document-state-38}\'s
    [history policy
    container](#document-state-history-policy-container){#populating-a-session-history-entry:document-state-history-policy-container-6},
    `sourceSnapshotParams`{.variable}\'s [source policy
    container](#source-snapshot-params-policy-container){#populating-a-session-history-entry:source-snapshot-params-policy-container-2},
    null, and `responsePolicyContainer`{.variable}.

25. If `navigable`{.variable}\'s
    [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-6}
    is an
    [`iframe`{#populating-a-session-history-entry:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
    and `response`{.variable}\'s [timing allow passed
    flag](https://fetch.spec.whatwg.org/#concept-response-timing-allow-passed){#populating-a-session-history-entry:concept-response-timing-allow-passed
    x-internal="concept-response-timing-allow-passed"} is set, then set
    `navigable`{.variable}\'s
    [container](document-sequences.html#nav-container){#populating-a-session-history-entry:nav-container-7}\'s
    [pending resource-timing start
    time](iframe-embed-object.html#iframe-pending-resource-timing-start-time){#populating-a-session-history-entry:iframe-pending-resource-timing-start-time}
    to null.

    If the
    [`iframe`{#populating-a-session-history-entry:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
    is allowed to report to resource timing, we don\'t need to run its
    fallback steps as the normal reporting would happen.

26. Return a new [navigation
    params](#navigation-params){#populating-a-session-history-entry:navigation-params-6},
    with

    [id](#navigation-params-id){#populating-a-session-history-entry:navigation-params-id-2}
    :   `navigationId`{.variable}

    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-2}
    :   `navigable`{.variable}

    [request](#navigation-params-request){#populating-a-session-history-entry:navigation-params-request-5}
    :   `request`{.variable}

    [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-11}
    :   `response`{.variable}

    [fetch controller](#navigation-params-fetch-controller){#populating-a-session-history-entry:navigation-params-fetch-controller-2}
    :   `fetchController`{.variable}

    [commit early hints](#navigation-params-commit-early-hints){#populating-a-session-history-entry:navigation-params-commit-early-hints-2}
    :   `commitEarlyHints`{.variable}

    [opener policy](#navigation-params-coop){#populating-a-session-history-entry:navigation-params-coop-2}
    :   `responseCOOP`{.variable}

    [reserved environment](#navigation-params-reserved-environment){#populating-a-session-history-entry:navigation-params-reserved-environment-4}
    :   `request`{.variable}\'s [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#populating-a-session-history-entry:concept-request-reserved-client-10
        x-internal="concept-request-reserved-client"}

    [origin](#navigation-params-origin){#populating-a-session-history-entry:navigation-params-origin-3}
    :   `responseOrigin`{.variable}

    [policy container](#navigation-params-policy-container){#populating-a-session-history-entry:navigation-params-policy-container-6}
    :   `resultPolicyContainer`{.variable}

    [final sandboxing flag set](#navigation-params-sandboxing){#populating-a-session-history-entry:navigation-params-sandboxing-3}
    :   `finalSandboxFlags`{.variable}

    [COOP enforcement result](#navigation-params-coop-enforcement-result){#populating-a-session-history-entry:navigation-params-coop-enforcement-result-2}
    :   `coopEnforcementResult`{.variable}

    [navigation timing type](#navigation-params-nav-timing-type){#populating-a-session-history-entry:navigation-params-nav-timing-type-2}
    :   `navTimingType`{.variable}

    [about base URL](#navigation-params-about-base-url){#populating-a-session-history-entry:navigation-params-about-base-url-2}
    :   `entry`{.variable}\'s [document
        state](#she-document-state){#populating-a-session-history-entry:she-document-state-39}\'s
        [about base
        URL](#document-state-about-base-url){#populating-a-session-history-entry:document-state-about-base-url-4}

    [user involvement](#navigation-params-user-involvement){#populating-a-session-history-entry:navigation-params-user-involvement-2}
    :   `userInvolvement`{.variable}

An element has a [browsing context scope
origin]{#browsing-context-scope-origin .dfn} if its
[`Document`{#populating-a-session-history-entry:document-4}](dom.html#document)\'s
[node
navigable](document-sequences.html#node-navigable){#populating-a-session-history-entry:node-navigable}
is a [top-level
traversable](document-sequences.html#top-level-traversable){#populating-a-session-history-entry:top-level-traversable-3}
or if all of its
[`Document`{#populating-a-session-history-entry:document-5}](dom.html#document)\'s
[ancestor
navigables](document-sequences.html#ancestor-navigables){#populating-a-session-history-entry:ancestor-navigables}
all have [active
documents](document-sequences.html#nav-document){#populating-a-session-history-entry:nav-document-8}
whose
[origins](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-5
x-internal="concept-document-origin"} are the [same
origin](browsers.html#same-origin){#populating-a-session-history-entry:same-origin-4}
as the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#populating-a-session-history-entry:node-document
x-internal="node-document"}\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-6
x-internal="concept-document-origin"}. If an element has a [browsing
context scope
origin](#browsing-context-scope-origin){#populating-a-session-history-entry:browsing-context-scope-origin-3},
then its value is the
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#populating-a-session-history-entry:concept-document-origin-7
x-internal="concept-document-origin"} of the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#populating-a-session-history-entry:node-document-2
x-internal="node-document"}.

This definition is broken and needs investigation to see what it was
intended to express: see [issue
#4703](https://github.com/whatwg/html/issues/4703).

To [load a document]{#loading-a-document .dfn} given [navigation
params](#navigation-params){#populating-a-session-history-entry:navigation-params-7}
`navigationParams`{.variable}, [source snapshot
params](#source-snapshot-params){#populating-a-session-history-entry:source-snapshot-params-3}
`sourceSnapshotParams`{.variable}, and
[origin](browsers.html#concept-origin){#populating-a-session-history-entry:concept-origin}
`initiatorOrigin`{.variable}, perform the following steps. They return a
[`Document`{#populating-a-session-history-entry:document-6}](dom.html#document)
or null.

1.  Let `type`{.variable} be the [computed
    type](https://mimesniff.spec.whatwg.org/#computed-mime-type){#populating-a-session-history-entry:content-type-sniffing-2
    x-internal="content-type-sniffing-2"} of
    `navigationParams`{.variable}\'s
    [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-12}.

2.  If the user agent has been configured to process resources of the
    given `type`{.variable} using some mechanism other than rendering
    the content in a
    [navigable](document-sequences.html#navigable){#populating-a-session-history-entry:navigable-6},
    then skip this step. Otherwise, if the `type`{.variable} is one of
    the following types:

    an [HTML MIME type](https://mimesniff.spec.whatwg.org/#html-mime-type){#populating-a-session-history-entry:html-mime-type x-internal="html-mime-type"}
    :   Return the result of [loading an HTML
        document](document-lifecycle.html#navigate-html){#populating-a-session-history-entry:navigate-html},
        given `navigationParams`{.variable}.

    an [XML MIME type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#populating-a-session-history-entry:xml-mime-type x-internal="xml-mime-type"} that is not an [explicitly supported XML MIME type](#explicitly-supported-xml-mime-type){#populating-a-session-history-entry:explicitly-supported-xml-mime-type}
    :   Return the result of [loading an XML
        document](document-lifecycle.html#read-xml){#populating-a-session-history-entry:read-xml}
        given `navigationParams`{.variable} and `type`{.variable}.

    a [JavaScript MIME type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#populating-a-session-history-entry:javascript-mime-type x-internal="javascript-mime-type"}\
    a [JSON MIME type](https://mimesniff.spec.whatwg.org/#json-mime-type){#populating-a-session-history-entry:json-mime-type x-internal="json-mime-type"} that is not an [explicitly supported JSON MIME type](#explicitly-supported-json-mime-type){#populating-a-session-history-entry:explicitly-supported-json-mime-type}\
    \"[`text/css`{#populating-a-session-history-entry:text/css}](indices.html#text/css)\"\
    \"[`text/plain`{#populating-a-session-history-entry:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\"\
    \"[`text/vtt`{#populating-a-session-history-entry:text/vtt}](indices.html#text/vtt)\"
    :   Return the result of [loading a text
        document](document-lifecycle.html#navigate-text){#populating-a-session-history-entry:navigate-text}
        given `navigationParams`{.variable} and `type`{.variable}.

    \"[`multipart/x-mixed-replace`{#populating-a-session-history-entry:multipart/x-mixed-replace}](iana.html#multipart/x-mixed-replace)\"
    :   Return the result of [loading a `multipart/x-mixed-replace`
        document](document-lifecycle.html#navigate-multipart-x-mixed-replace){#populating-a-session-history-entry:navigate-multipart-x-mixed-replace},
        given `navigationParams`{.variable},
        `sourceSnapshotParams`{.variable}, and
        `initiatorOrigin`{.variable}.

    a supported image, video, or audio type
    :   Return the result of [loading a media
        document](document-lifecycle.html#navigate-media){#populating-a-session-history-entry:navigate-media}
        given `navigationParams`{.variable} and `type`{.variable}.

    \"`application/pdf`\"\
    \"`text/pdf`\"
    :   If the user agent\'s [PDF viewer
        supported](system-state.html#pdf-viewer-supported){#populating-a-session-history-entry:pdf-viewer-supported}
        is true, return the result of [creating a document for inline
        content that doesn\'t have a
        DOM](document-lifecycle.html#read-ua-inline){#populating-a-session-history-entry:read-ua-inline-3}
        given `navigationParams`{.variable}\'s
        [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-3},
        `navigationParams`{.variable}\'s
        [id](#navigation-params-id){#populating-a-session-history-entry:navigation-params-id-3},
        `navigationParams`{.variable}\'s [navigation timing
        type](#navigation-params-nav-timing-type){#populating-a-session-history-entry:navigation-params-nav-timing-type-3},
        and `navigationParams`{.variable}\'s [user
        involvement](#navigation-params-user-involvement){#populating-a-session-history-entry:navigation-params-user-involvement-3}.

    Otherwise, proceed onward.

    An [explicitly supported XML MIME
    type]{#explicitly-supported-xml-mime-type .dfn} is an [XML MIME
    type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#populating-a-session-history-entry:xml-mime-type-2
    x-internal="xml-mime-type"} for which the user agent is configured
    to use an external application to render the content, or for which
    the user agent has dedicated processing rules. For example, a web
    browser with a built-in Atom feed viewer would be said to explicitly
    support the
    [`application/atom+xml`{#populating-a-session-history-entry:application/atom+xml}](indices.html#application/atom+xml)
    MIME type.

    An [explicitly supported JSON MIME
    type]{#explicitly-supported-json-mime-type .dfn} is a [JSON MIME
    type](https://mimesniff.spec.whatwg.org/#json-mime-type){#populating-a-session-history-entry:json-mime-type-2
    x-internal="json-mime-type"} for which the user agent is configured
    to use an external application to render the content, or for which
    the user agent has dedicated processing rules.

    In both cases, the external application or user agent will either
    [display the content
    inline](document-lifecycle.html#read-ua-inline){#populating-a-session-history-entry:read-ua-inline-4}
    directly in `navigationParams`{.variable}\'s
    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-4},
    or [hand it off to external
    software](#hand-off-to-external-software){#populating-a-session-history-entry:hand-off-to-external-software-3}.
    Both happen in the steps below.

3.  [If, given `type`{.variable}, the new resource is to be handled by
    displaying some sort of inline content, e.g., a native rendering of
    the content or an error message because the specified type is not
    supported, then return the result of [creating a document for inline
    content that doesn\'t have a
    DOM](document-lifecycle.html#read-ua-inline){#populating-a-session-history-entry:read-ua-inline-5}
    given `navigationParams`{.variable}\'s
    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-5},
    `navigationParams`{.variable}\'s
    [id](#navigation-params-id){#populating-a-session-history-entry:navigation-params-id-4},
    `navigationParams`{.variable}\'s [navigation timing
    type](#navigation-params-nav-timing-type){#populating-a-session-history-entry:navigation-params-nav-timing-type-4},
    and `navigationParams`{.variable}\'s [user
    involvement](#navigation-params-user-involvement){#populating-a-session-history-entry:navigation-params-user-involvement-4}.]{#navigate-non-Document}

4.  Otherwise, the document\'s `type`{.variable} is such that the
    resource will not affect `navigationParams`{.variable}\'s
    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-6},
    e.g., because the resource is to be handed to an external
    application or because it is an unknown type that will be processed
    by [handle as a
    download](links.html#handle-as-a-download){#populating-a-session-history-entry:handle-as-a-download-2}.
    [Hand-off to external
    software](#hand-off-to-external-software){#populating-a-session-history-entry:hand-off-to-external-software-4}
    given `navigationParams`{.variable}\'s
    [response](#navigation-params-response){#populating-a-session-history-entry:navigation-params-response-13},
    `navigationParams`{.variable}\'s
    [navigable](#navigation-params-navigable){#populating-a-session-history-entry:navigation-params-navigable-7},
    `navigationParams`{.variable}\'s [final sandboxing flag
    set](#navigation-params-sandboxing){#populating-a-session-history-entry:navigation-params-sandboxing-4},
    `sourceSnapshotParams`{.variable}\'s [has transient
    activation](#source-snapshot-params-activation){#populating-a-session-history-entry:source-snapshot-params-activation-4},
    and `initiatorOrigin`{.variable}.

5.  Return null.

#### [7.4.6]{.secno} Applying the history step[](#applying-the-history-step){.self-link}

For both navigation and traversal, once we have an idea of where we want
to head to in the session history, much of the work comes about in
applying that notion to the [traversable
navigable](document-sequences.html#traversable-navigable){#applying-the-history-step:traversable-navigable}
and the relevant
[`Document`{#applying-the-history-step:document}](dom.html#document).
For navigations, this work generally occurs toward the end of the
process; for traversals, it is the beginning.

##### [7.4.6.1]{.secno} Updating the traversable[](#updating-the-traversable){.self-link}

Ensuring a
[traversable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable}
ends up at the right session history step is particularly complex, as it
can involve coordinating across multiple
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable}
descendants of the traversable,
[populating](#populating-a-session-history-entry) them in parallel, and
then synchronizing back up to ensure everyone has the same view of the
result. This is further complicated by the existence of synchronous
same-document navigations being mixed together with cross-document
navigations, and how web pages have come to have certain relative timing
expectations.

A [changing navigable continuation
state]{#changing-navigable-continuation-state .dfn} is used to store
information during the [apply the history
step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step}
algorithm, allowing parts of the algorithm to continue only after other
parts have finished. It is a
[struct](https://infra.spec.whatwg.org/#struct){#updating-the-traversable:struct
x-internal="struct"} with:

[displayed document]{#changing-nav-continuation-displayed-document .dfn}
:   A
    [`Document`{#updating-the-traversable:document}](dom.html#document)

[target entry]{#changing-nav-continuation-target-entry .dfn}
:   A [session history
    entry](#session-history-entry){#updating-the-traversable:session-history-entry}

[navigable]{#changing-nav-continuation-navigable .dfn}
:   A
    [navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-2}

[update only]{#changing-nav-continuation-update-only .dfn}
:   A boolean

------------------------------------------------------------------------

Although all updates to the [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-2}
end up in the same [apply the history
step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-2}
algorithm, each possible entry point comes along with some minor
customizations:

To [update for navigable
creation/destruction]{#update-for-navigable-creation/destruction .dfn}
given a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-3}
`traversable`{.variable}:

1.  Let `step`{.variable} be `traversable`{.variable}\'s [current
    session history
    step](document-sequences.html#tn-current-session-history-step){#updating-the-traversable:tn-current-session-history-step}.

2.  Return the result of [applying the history
    step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-3}
    `step`{.variable} to `traversable`{.variable} given false, null,
    null, \"[`none`{#updating-the-traversable:uni-none}](#uni-none)\",
    and null.

To [apply the push/replace history
step]{#apply-the-push/replace-history-step .dfn} given a non-negative
integer `step`{.variable} to a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-4}
`traversable`{.variable}, given a [history handling
behavior](#history-handling-behavior){#updating-the-traversable:history-handling-behavior}
`historyHandling`{.variable} and a [user navigation
involvement](#user-navigation-involvement){#updating-the-traversable:user-navigation-involvement}
`userInvolvement`{.variable}:

1.  Return the result of [applying the history
    step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-4}
    `step`{.variable} to `traversable`{.variable} given false, null,
    null, `userInvolvement`{.variable}, and
    `historyHandling`{.variable}.

[Apply the push/replace history
step](#apply-the-push/replace-history-step){#updating-the-traversable:apply-the-push/replace-history-step}
never passes [source snapshot
params](#source-snapshot-params){#updating-the-traversable:source-snapshot-params}
or an initiator
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-3}
to [apply the history
step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-5}.
This is because those checks are done earlier in the
[navigation](#navigate){#updating-the-traversable:navigate} algorithm.

To [apply the reload history step]{#apply-the-reload-history-step .dfn}
to a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-5}
`traversable`{.variable} given [user navigation
involvement](#user-navigation-involvement){#updating-the-traversable:user-navigation-involvement-2}
`userInvolvement`{.variable}:

1.  Let `step`{.variable} be `traversable`{.variable}\'s [current
    session history
    step](document-sequences.html#tn-current-session-history-step){#updating-the-traversable:tn-current-session-history-step-2}.

2.  Return the result of [applying the history
    step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-6}
    `step`{.variable} to `traversable`{.variable} given true, null,
    null, `userInvolvement`{.variable}, and
    \"[`reload`{#updating-the-traversable:dom-navigationtype-reload}](nav-history-apis.html#dom-navigationtype-reload)\".

[Apply the reload history
step](#apply-the-reload-history-step){#updating-the-traversable:apply-the-reload-history-step}
never passes [source snapshot
params](#source-snapshot-params){#updating-the-traversable:source-snapshot-params-2}
or an initiator
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-4}
to [apply the history
step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-7}.
This is because reloading is always treated as if it were done by the
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-5}
itself, even in cases like `parent.location.reload()`.

To [apply the traverse history step]{#apply-the-traverse-history-step
.dfn} given a non-negative integer `step`{.variable} to a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-6}
`traversable`{.variable}, with [source snapshot
params](#source-snapshot-params){#updating-the-traversable:source-snapshot-params-3}
`sourceSnapshotParams`{.variable},
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-6}
`initiatorToCheck`{.variable}, and [user navigation
involvement](#user-navigation-involvement){#updating-the-traversable:user-navigation-involvement-3}
`userInvolvement`{.variable}:

1.  Return the result of [applying the history
    step](#apply-the-history-step){#updating-the-traversable:apply-the-history-step-8}
    `step`{.variable} to `traversable`{.variable} given true,
    `sourceSnapshotParams`{.variable}, `initiatorToCheck`{.variable},
    `userInvolvement`{.variable}, and
    \"[`traverse`{#updating-the-traversable:dom-navigationtype-traverse}](nav-history-apis.html#dom-navigationtype-traverse)\".

------------------------------------------------------------------------

Now for the algorithm itself.

To [apply the history step]{#apply-the-history-step .dfn} given a
non-negative integer `step`{.variable} to a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-7}
`traversable`{.variable}, with boolean `checkForCancelation`{.variable},
[source snapshot
params](#source-snapshot-params){#updating-the-traversable:source-snapshot-params-4}-or-null
`sourceSnapshotParams`{.variable},
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-7}-or-null
`initiatorToCheck`{.variable}, [user navigation
involvement](#user-navigation-involvement){#updating-the-traversable:user-navigation-involvement-4}
`userInvolvement`{.variable}, and
[`NavigationType`{#updating-the-traversable:navigationtype}](nav-history-apis.html#navigationtype)-or-null
`navigationType`{.variable}, perform the following steps. They return
\"`initiator-disallowed`\", \"`canceled-by-beforeunload`\",
\"`canceled-by-navigate`\", or \"`applied`\".

1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert
    x-internal="assert"}: This is running within
    `traversable`{.variable}\'s [session history traversal
    queue](document-sequences.html#tn-session-history-traversal-queue){#updating-the-traversable:tn-session-history-traversal-queue}.

2.  Let `targetStep`{.variable} be the result of [getting the used
    step](#getting-the-used-step){#updating-the-traversable:getting-the-used-step}
    given `traversable`{.variable} and `step`{.variable}.

3.  If `initiatorToCheck`{.variable} is not null, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-2
        x-internal="assert"}: `sourceSnapshotParams`{.variable} is not
        null.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate
        x-internal="list-iterate"} `navigable`{.variable} of [get all
        navigables whose current session history entry will change or
        reload](#get-all-navigables-whose-current-session-history-entry-will-change-or-reload){#updating-the-traversable:get-all-navigables-whose-current-session-history-entry-will-change-or-reload}:
        if `initiatorToCheck`{.variable} is not [allowed by sandboxing
        to
        navigate](#allowed-to-navigate){#updating-the-traversable:allowed-to-navigate}
        `navigable`{.variable} given `sourceSnapshotParams`{.variable},
        then return \"`initiator-disallowed`\".

4.  Let `navigablesCrossingDocuments`{.variable} be the result of
    [getting all navigables that might experience a cross-document
    traversal](#getting-all-navigables-that-might-experience-a-cross-document-traversal){#updating-the-traversable:getting-all-navigables-that-might-experience-a-cross-document-traversal}
    given `traversable`{.variable} and `targetStep`{.variable}.

5.  If `checkForCancelation`{.variable} is true, and the result of
    [checking if unloading is
    canceled](#checking-if-unloading-is-canceled){#updating-the-traversable:checking-if-unloading-is-canceled}
    given `navigablesCrossingDocuments`{.variable},
    `traversable`{.variable}, `targetStep`{.variable}, and
    `userInvolvement`{.variable} is not \"`continue`\", then return that
    result.

6.  Let `changingNavigables`{.variable} be the result of [get all
    navigables whose current session history entry will change or
    reload](#get-all-navigables-whose-current-session-history-entry-will-change-or-reload){#updating-the-traversable:get-all-navigables-whose-current-session-history-entry-will-change-or-reload-2}
    given `traversable`{.variable} and `targetStep`{.variable}.

7.  Let `nonchangingNavigablesThatStillNeedUpdates`{.variable} be the
    result of [getting all navigables that only need history object
    length/index
    update](#getting-all-navigables-that-only-need-history-object-length/index-update){#updating-the-traversable:getting-all-navigables-that-only-need-history-object-length/index-update}
    given `traversable`{.variable} and `targetStep`{.variable}.

8.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-2
    x-internal="list-iterate"} `navigable`{.variable} of
    `changingNavigables`{.variable}:

    1.  Let `targetEntry`{.variable} be the result of [getting the
        target history
        entry](#getting-the-target-history-entry){#updating-the-traversable:getting-the-target-history-entry}
        given `navigable`{.variable} and `targetStep`{.variable}.

    2.  Set `navigable`{.variable}\'s [current session history
        entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry}
        to `targetEntry`{.variable}.

    3.  [Set the ongoing
        navigation](#set-the-ongoing-navigation){#updating-the-traversable:set-the-ongoing-navigation}
        for `navigable`{.variable} to \"`traversal`\".

9.  Let `totalChangeJobs`{.variable} be the
    [size](https://infra.spec.whatwg.org/#list-size){#updating-the-traversable:list-size
    x-internal="list-size"} of `changingNavigables`{.variable}.

10. Let `completedChangeJobs`{.variable} be 0.

11. Let `changingNavigableContinuations`{.variable} be an empty
    [queue](https://infra.spec.whatwg.org/#queue){#updating-the-traversable:queue
    x-internal="queue"} of [changing navigable continuation
    states](#changing-navigable-continuation-state){#updating-the-traversable:changing-navigable-continuation-state}.

    This queue is used to split the operations on
    `changingNavigables`{.variable} into two parts. Specifically,
    `changingNavigableContinuations`{.variable} holds data for the
    [second part](#continuations-dequeue).

12. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-3
    x-internal="list-iterate"} `navigable`{.variable} of
    `changingNavigables`{.variable}, [queue a global
    task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source}
    of `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#updating-the-traversable:nav-window}
    to run the steps:

    This set of steps are split into two parts to allow synchronous
    navigations to be processed before documents unload. State is stored
    in `changingNavigableContinuations`{.variable} for the [second
    part](#continuations-dequeue).

    1.  Let `displayedEntry`{.variable} be `navigable`{.variable}\'s
        [active session history
        entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry}.

    2.  Let `targetEntry`{.variable} be `navigable`{.variable}\'s
        [current session history
        entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-2}.

    3.  Let `changingNavigableContinuation`{.variable} be a [changing
        navigable continuation
        state](#changing-navigable-continuation-state){#updating-the-traversable:changing-navigable-continuation-state-2}
        with:

        [displayed document](#changing-nav-continuation-displayed-document){#updating-the-traversable:changing-nav-continuation-displayed-document}
        :   `displayedEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document}

        [target entry](#changing-nav-continuation-target-entry){#updating-the-traversable:changing-nav-continuation-target-entry}
        :   `targetEntry`{.variable}

        [navigable](#changing-nav-continuation-navigable){#updating-the-traversable:changing-nav-continuation-navigable}
        :   `navigable`{.variable}

        [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only}
        :   false

    4.  If `displayedEntry`{.variable} is `targetEntry`{.variable} and
        `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending}
        is false, then:

        1.  Set `changingNavigableContinuation`{.variable}\'s
            [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only-2}
            to true.

        2.  [Enqueue](https://infra.spec.whatwg.org/#queue-enqueue){#updating-the-traversable:enqueue
            x-internal="enqueue"}
            `changingNavigableContinuation`{.variable} on
            `changingNavigableContinuations`{.variable}.

        3.  Abort these steps.

        This case occurs due to a [synchronous
        navigation](#finalize-a-same-document-navigation){#updating-the-traversable:finalize-a-same-document-navigation}
        which already updated the [active session history
        entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-2}.

    5.  Switch on `navigationType`{.variable}:

        \"[`reload`{#updating-the-traversable:dom-navigationtype-reload-2}](nav-history-apis.html#dom-navigationtype-reload)\"
        :   [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-3
            x-internal="assert"}: `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-2}\'s
            [reload
            pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-2}
            is true.

        \"[`traverse`{#updating-the-traversable:dom-navigationtype-traverse-2}](nav-history-apis.html#dom-navigationtype-traverse)\"
        :   [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-4
            x-internal="assert"}: `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-3}\'s
            [ever
            populated](#document-state-ever-populated){#updating-the-traversable:document-state-ever-populated}
            is true.

        \"[`replace`{#updating-the-traversable:dom-navigationtype-replace}](nav-history-apis.html#dom-navigationtype-replace)\"
        :   [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-5
            x-internal="assert"}: `targetEntry`{.variable}\'s
            [step](#she-step){#updating-the-traversable:she-step} is
            `displayedEntry`{.variable}\'s
            [step](#she-step){#updating-the-traversable:she-step-2} and
            `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-4}\'s
            [ever
            populated](#document-state-ever-populated){#updating-the-traversable:document-state-ever-populated-2}
            is false.

        \"[`push`{#updating-the-traversable:dom-navigationtype-push}](nav-history-apis.html#dom-navigationtype-push)\"

        :   [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-6
            x-internal="assert"}: `targetEntry`{.variable}\'s
            [step](#she-step){#updating-the-traversable:she-step-3} is
            `displayedEntry`{.variable}\'s
            [step](#she-step){#updating-the-traversable:she-step-4} + 1
            and `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-5}\'s
            [ever
            populated](#document-state-ever-populated){#updating-the-traversable:document-state-ever-populated-3}
            is false.

    6.  Let `oldOrigin`{.variable} be `targetEntry`{.variable}\'s
        [document
        state](#she-document-state){#updating-the-traversable:she-document-state-6}\'s
        [origin](#document-state-origin){#updating-the-traversable:document-state-origin}.

    7.  ::: {#descendant-navigable-traversal-navigate-events}
        If all of the following are true:

        - `navigable`{.variable} is not `traversable`{.variable};

        - `targetEntry`{.variable} is not `navigable`{.variable}\'s
          [current session history
          entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-3};
          and

        - `oldOrigin`{.variable} is the
          [same](browsers.html#same-origin){#updating-the-traversable:same-origin}
          as `navigable`{.variable}\'s [current session history
          entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-4}\'s
          [document
          state](#she-document-state){#updating-the-traversable:she-document-state-7}\'s
          [origin](#document-state-origin){#updating-the-traversable:document-state-origin-2},

        then:

        1.  Let `navigation`{.variable} be `navigable`{.variable}\'s
            [active
            window](document-sequences.html#nav-window){#updating-the-traversable:nav-window-2}\'s
            [navigation
            API](nav-history-apis.html#window-navigation-api){#updating-the-traversable:window-navigation-api}.

        2.  [Fire a traverse `navigate`
            event](nav-history-apis.html#fire-a-traverse-navigate-event){#updating-the-traversable:fire-a-traverse-navigate-event}
            at `navigation`{.variable} given `targetEntry`{.variable}
            and `userInvolvement`{.variable}.
        :::

    8.  If `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-2}
        is null, or `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state-8}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-3}
        is true, then:

        1.  Let `navTimingType`{.variable} be
            \"[`back_forward`{#updating-the-traversable:dom-navigationtimingtype-back_forward}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-back_forward){x-internal="dom-navigationtimingtype-back_forward"}\"
            if `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-3}
            is null; otherwise
            \"[`reload`{#updating-the-traversable:dom-navigationtimingtype-back_forward-2}](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-back_forward){x-internal="dom-navigationtimingtype-back_forward"}\".

        2.  Let `targetSnapshotParams`{.variable} be the result of
            [snapshotting target snapshot
            params](#snapshotting-target-snapshot-params){#updating-the-traversable:snapshotting-target-snapshot-params}
            given `navigable`{.variable}.

        3.  Let
            `potentiallyTargetSpecificSourceSnapshotParams`{.variable}
            be `sourceSnapshotParams`{.variable}.

        4.  If
            `potentiallyTargetSpecificSourceSnapshotParams`{.variable}
            is null, then set it to the result of [snapshotting source
            snapshot
            params](#snapshotting-source-snapshot-params){#updating-the-traversable:snapshotting-source-snapshot-params}
            given `navigable`{.variable}\'s [active
            document](document-sequences.html#nav-document){#updating-the-traversable:nav-document}.

            In this case there is no clear source of the
            traversal/reload. We treat this situation as if
            `navigable`{.variable} navigated itself, but note that some
            properties of `targetEntry`{.variable}\'s original initiator
            are preserved in `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-9},
            such as the [initiator
            origin](#document-state-initiator-origin){#updating-the-traversable:document-state-initiator-origin}
            and
            [referrer](#document-state-request-referrer){#updating-the-traversable:document-state-request-referrer},
            which will appropriately influence the navigation.

        5.  Set `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-10}\'s
            [reload
            pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-4}
            to false.

        6.  Let `allowPOST`{.variable} be `targetEntry`{.variable}\'s
            [document
            state](#she-document-state){#updating-the-traversable:she-document-state-11}\'s
            [reload
            pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-5}.

        7.  [In
            parallel](infrastructure.html#in-parallel){#updating-the-traversable:in-parallel},
            [attempt to populate the history entry\'s
            document](#attempt-to-populate-the-history-entry's-document){#updating-the-traversable:attempt-to-populate-the-history-entry's-document}
            for `targetEntry`{.variable}, given `navigable`{.variable},
            `potentiallyTargetSpecificSourceSnapshotParams`{.variable},
            `targetSnapshotParams`{.variable},
            `userInvolvement`{.variable}, with
            *[allowPOST](#attempt-to-populate-allow-post)* set to
            `allowPOST`{.variable} and
            [*completionSteps*](#attempt-to-populate-completion-steps){#updating-the-traversable:attempt-to-populate-completion-steps}
            set to [queue a global
            task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task-2}
            on the [navigation and traversal task
            source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source-2}
            given `navigable`{.variable}\'s [active
            window](document-sequences.html#nav-window){#updating-the-traversable:nav-window-3}
            to run `afterDocumentPopulated`{.variable}.

        Otherwise, run `afterDocumentPopulated`{.variable}
        [immediately](infrastructure.html#immediately){#updating-the-traversable:immediately}.

        In both cases, let `afterDocumentPopulated`{.variable} be the
        following steps:

        1.  If `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-4}
            is null, then set
            `changingNavigableContinuation`{.variable}\'s
            [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only-3}
            to true.

            This means we tried to populate the document, but were
            unable to do so, e.g. because of the server returning a 204.

            ::: note
            These kinds of failed navigations or traversals will not be
            signaled to the [navigation
            API](nav-history-apis.html#navigation-api) (e.g., through
            the promises of any [navigation API method
            tracker](nav-history-apis.html#navigation-api-method-tracker){#updating-the-traversable:navigation-api-method-tracker},
            or the
            [`navigateerror`{#updating-the-traversable:event-navigateerror}](indices.html#event-navigateerror)
            event). Doing so would leak information about the timing of
            responses from other origins, in the cross-origin case, and
            providing different results in the cross-origin vs.
            same-origin cases was deemed too confusing.

            However, implementations could use this opportunity to clear
            any promise handlers for the
            [`navigation.transition.finished`{#updating-the-traversable:dom-navigationtransition-finished}](nav-history-apis.html#dom-navigationtransition-finished)
            promise, as they are guaranteed at this point to never run.
            And, they might wish to [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#updating-the-traversable:report-a-warning-to-the-console
            x-internal="report-a-warning-to-the-console"} if any part of
            the navigation API initiated these navigations, to make it
            clear to the web developer why their promises will never
            settle and events will never fire.
            :::

        2.  If `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-5}\'s
            [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-traversable:concept-document-origin
            x-internal="concept-document-origin"} is not
            `oldOrigin`{.variable}, then set `targetEntry`{.variable}\'s
            [classic history API
            state](#she-classic-history-api-state){#updating-the-traversable:she-classic-history-api-state}
            to
            [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#updating-the-traversable:structuredserializeforstorage}(null).

            This clears history state when the origin changed vs a
            previous load of `targetEntry`{.variable} without a redirect
            occuring. This can happen due to a change in CSP sandbox
            headers.

        3.  If all of the following are true:

            - `navigable`{.variable}\'s
              [parent](document-sequences.html#nav-parent){#updating-the-traversable:nav-parent}
              is null;

            - `targetEntry`{.variable}\'s
              [document](#she-document){#updating-the-traversable:she-document-6}\'s
              [browsing
              context](document-sequences.html#concept-document-bc){#updating-the-traversable:concept-document-bc}
              is not an [auxiliary browsing
              context](document-sequences.html#auxiliary-browsing-context){#updating-the-traversable:auxiliary-browsing-context}
              whose [opener browsing
              context](document-sequences.html#opener-browsing-context){#updating-the-traversable:opener-browsing-context}
              is non-null; and

            - `targetEntry`{.variable}\'s
              [document](#she-document){#updating-the-traversable:she-document-7}\'s
              [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-traversable:concept-document-origin-2
              x-internal="concept-document-origin"} is not
              `oldOrigin`{.variable},

            then set `targetEntry`{.variable}\'s [document
            state](#she-document-state){#updating-the-traversable:she-document-state-12}\'s
            [navigable target
            name](#document-state-nav-target-name){#updating-the-traversable:document-state-nav-target-name}
            to the empty string.

        4.  [Enqueue](https://infra.spec.whatwg.org/#queue-enqueue){#updating-the-traversable:enqueue-2
            x-internal="enqueue"}
            `changingNavigableContinuation`{.variable} on
            `changingNavigableContinuations`{.variable}.

            The rest of this job [runs later](#continuations-dequeue) in
            this algorithm.

13. Let `navigablesThatMustWaitBeforeHandlingSyncNavigation`{.variable}
    be an empty
    [set](https://infra.spec.whatwg.org/#ordered-set){#updating-the-traversable:set
    x-internal="set"}.

14. While `completedChangeJobs`{.variable} does not equal
    `totalChangeJobs`{.variable}:

    1.  If `traversable`{.variable}\'s [running nested apply history
        step](document-sequences.html#tn-running-nested-apply-history-step){#updating-the-traversable:tn-running-nested-apply-history-step}
        is false, then:

        1.  ::: {#sync-navigations-jump-queue}
            While `traversable`{.variable}\'s [session history traversal
            queue](document-sequences.html#tn-session-history-traversal-queue){#updating-the-traversable:tn-session-history-traversal-queue-2}\'s
            [algorithm
            set](#session-history-traversal-parallel-queue-algorithm-set){#updating-the-traversable:session-history-traversal-parallel-queue-algorithm-set}
            [contains](https://infra.spec.whatwg.org/#list-contain){#updating-the-traversable:list-contains
            x-internal="list-contains"} one or more [synchronous
            navigation
            steps](#session-history-traversal-parallel-queue-sync-nav-steps){#updating-the-traversable:session-history-traversal-parallel-queue-sync-nav-steps}
            with a [target
            navigable](#session-history-traversal-parallel-queue-sync-nav-steps-target-nav){#updating-the-traversable:session-history-traversal-parallel-queue-sync-nav-steps-target-nav}
            not
            [contained](https://infra.spec.whatwg.org/#list-contain){#updating-the-traversable:list-contains-2
            x-internal="list-contains"} in
            `navigablesThatMustWaitBeforeHandlingSyncNavigation`{.variable}:

            1.  Let `steps`{.variable} be the first
                [item](https://infra.spec.whatwg.org/#list-item){#updating-the-traversable:list-item
                x-internal="list-item"} in `traversable`{.variable}\'s
                [session history traversal
                queue](document-sequences.html#tn-session-history-traversal-queue){#updating-the-traversable:tn-session-history-traversal-queue-3}\'s
                [algorithm
                set](#session-history-traversal-parallel-queue-algorithm-set){#updating-the-traversable:session-history-traversal-parallel-queue-algorithm-set-2}
                that is [synchronous navigation
                steps](#session-history-traversal-parallel-queue-sync-nav-steps){#updating-the-traversable:session-history-traversal-parallel-queue-sync-nav-steps-2}
                with a [target
                navigable](#session-history-traversal-parallel-queue-sync-nav-steps-target-nav){#updating-the-traversable:session-history-traversal-parallel-queue-sync-nav-steps-target-nav-2}
                not
                [contained](https://infra.spec.whatwg.org/#list-contain){#updating-the-traversable:list-contains-3
                x-internal="list-contains"} in
                `navigablesThatMustWaitBeforeHandlingSyncNavigation`{.variable}.

            2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#updating-the-traversable:list-remove
                x-internal="list-remove"} `steps`{.variable} from
                `traversable`{.variable}\'s [session history traversal
                queue](document-sequences.html#tn-session-history-traversal-queue){#updating-the-traversable:tn-session-history-traversal-queue-4}\'s
                [algorithm
                set](#session-history-traversal-parallel-queue-algorithm-set){#updating-the-traversable:session-history-traversal-parallel-queue-algorithm-set-3}.

            3.  Set `traversable`{.variable}\'s [running nested apply
                history
                step](document-sequences.html#tn-running-nested-apply-history-step){#updating-the-traversable:tn-running-nested-apply-history-step-2}
                to true.

            4.  Run `steps`{.variable}.

            5.  Set `traversable`{.variable}\'s [running nested apply
                history
                step](document-sequences.html#tn-running-nested-apply-history-step){#updating-the-traversable:tn-running-nested-apply-history-step-3}
                to false.

            Synchronous navigations that are intended to take place
            before this traversal jump the queue at this point, so they
            can be added to the correct place in
            `traversable`{.variable}\'s [session history
            entries](document-sequences.html#tn-session-history-entries){#updating-the-traversable:tn-session-history-entries}
            before this traversal potentially unloads their document.
            [More details can be found
            here](#sync-navigation-steps-queue-jumping-examples).
            :::

    2.  [Let `changingNavigableContinuation`{.variable} be the result of
        [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#updating-the-traversable:dequeue
        x-internal="dequeue"} from
        `changingNavigableContinuations`{.variable}.]{#continuations-dequeue}

    3.  If `changingNavigableContinuation`{.variable} is nothing, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-traversable:continue
        x-internal="continue"}.

    4.  Let `displayedDocument`{.variable} be
        `changingNavigableContinuation`{.variable}\'s [displayed
        document](#changing-nav-continuation-displayed-document){#updating-the-traversable:changing-nav-continuation-displayed-document-2}.

    5.  Let `targetEntry`{.variable} be
        `changingNavigableContinuation`{.variable}\'s [target
        entry](#changing-nav-continuation-target-entry){#updating-the-traversable:changing-nav-continuation-target-entry-2}.

    6.  Let `navigable`{.variable} be
        `changingNavigableContinuation`{.variable}\'s
        [navigable](#changing-nav-continuation-navigable){#updating-the-traversable:changing-nav-continuation-navigable-2}.

    7.  Let (`scriptHistoryLength`{.variable},
        `scriptHistoryIndex`{.variable}) be the result of [getting the
        history object length and
        index](#getting-the-history-object-length-and-index){#updating-the-traversable:getting-the-history-object-length-and-index}
        given `traversable`{.variable} and `targetStep`{.variable}.

        These values might have changed since they were last calculated.

    8.  [Append](https://infra.spec.whatwg.org/#list-append){#updating-the-traversable:list-append
        x-internal="list-append"} `navigable`{.variable} to
        `navigablesThatMustWaitBeforeHandlingSyncNavigation`{.variable}.

        Once a navigable has reached this point in traversal,
        additionally queued synchronous navigation steps are likely to
        be intended to occur after this traversal rather than before it,
        so they no longer jump the queue. [More details can be found
        here](#sync-navigation-steps-queue-jumping-examples).

    9.  Let `entriesForNavigationAPI`{.variable} be the result of
        [getting session history entries for the navigation
        API](#getting-session-history-entries-for-the-navigation-api){#updating-the-traversable:getting-session-history-entries-for-the-navigation-api}
        given `navigable`{.variable} and `targetStep`{.variable}.

    10. If `changingNavigableContinuation`{.variable}\'s
        [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only-4}
        is true, or `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-8}
        is `displayedDocument`{.variable}, then:

        This is a same-document navigation: we proceed without
        unloading.

        1.  [Set the ongoing
            navigation](#set-the-ongoing-navigation){#updating-the-traversable:set-the-ongoing-navigation-2}
            for `navigable`{.variable} to null.

            This allows new
            [navigations](#navigate){#updating-the-traversable:navigate-2}
            of `navigable`{.variable} to start, whereas during the
            traversal they were blocked.

        2.  [Queue a global
            task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task-3}
            on the [navigation and traversal task
            source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source-3}
            given `navigable`{.variable}\'s [active
            window](document-sequences.html#nav-window){#updating-the-traversable:nav-window-4}
            to perform `afterPotentialUnloads`{.variable}.

    11. Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-7
            x-internal="assert"}: `navigationType`{.variable} is not
            null.

        2.  [Deactivate](#deactivate-a-document-for-a-cross-document-navigation){#updating-the-traversable:deactivate-a-document-for-a-cross-document-navigation}
            `displayedDocument`{.variable}, given
            `userInvolvement`{.variable}, `targetEntry`{.variable},
            `navigationType`{.variable}, and
            `afterPotentialUnloads`{.variable}.

    12. In both cases, let `afterPotentialUnloads`{.variable} be the
        following steps:

        1.  Let `previousEntry`{.variable} be `navigable`{.variable}\'s
            [active session history
            entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-3}.

        2.  If `changingNavigableContinuation`{.variable}\'s
            [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only-5}
            is false, then [activate history
            entry](#activate-history-entry){#updating-the-traversable:activate-history-entry}
            `targetEntry`{.variable} for `navigable`{.variable}.

        3.  Let `updateDocument`{.variable} be an algorithm step which
            performs [update document for history step
            application](#update-document-for-history-step-application){#updating-the-traversable:update-document-for-history-step-application}
            given `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-9},
            `targetEntry`{.variable},
            `changingNavigableContinuation`{.variable}\'s
            [update-only](#changing-nav-continuation-update-only){#updating-the-traversable:changing-nav-continuation-update-only-6},
            `scriptHistoryLength`{.variable},
            `scriptHistoryIndex`{.variable},
            `navigationType`{.variable},
            `entriesForNavigationAPI`{.variable}, and
            `previousEntry`{.variable}.

        4.  If `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-10}
            is equal to `displayedDocument`{.variable}, then perform
            `updateDocument`{.variable}.

        5.  Otherwise, [queue a global
            task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task-4}
            on the [navigation and traversal task
            source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source-4}
            given `targetEntry`{.variable}\'s
            [document](#she-document){#updating-the-traversable:she-document-11}\'s
            [relevant global
            object](webappapis.html#concept-relevant-global){#updating-the-traversable:concept-relevant-global}
            to perform `updateDocument`{.variable}.

        6.  Increment `completedChangeJobs`{.variable}.

15. Let `totalNonchangingJobs`{.variable} be the
    [size](https://infra.spec.whatwg.org/#list-size){#updating-the-traversable:list-size-2
    x-internal="list-size"} of
    `nonchangingNavigablesThatStillNeedUpdates`{.variable}.

    This step onwards deliberately waits for all the previous operations
    to complete, as they include [processing synchronous
    navigations](#sync-navigations-jump-queue) which will also post
    tasks to update history length and index.

16. Let `completedNonchangingJobs`{.variable} be 0.

17. Let (`scriptHistoryLength`{.variable},
    `scriptHistoryIndex`{.variable}) be the result of [getting the
    history object length and
    index](#getting-the-history-object-length-and-index){#updating-the-traversable:getting-the-history-object-length-and-index-2}
    given `traversable`{.variable} and `targetStep`{.variable}.

18. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-4
    x-internal="list-iterate"} `navigable`{.variable} of
    `nonchangingNavigablesThatStillNeedUpdates`{.variable}, [queue a
    global
    task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task-5}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source-5}
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#updating-the-traversable:nav-window-5}
    to run the steps:

    1.  Let `document`{.variable} be `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-2}.

    2.  Set `document`{.variable}\'s [history
        object](nav-history-apis.html#doc-history){#updating-the-traversable:doc-history}\'s
        [index](nav-history-apis.html#concept-history-index){#updating-the-traversable:concept-history-index}
        to `scriptHistoryIndex`{.variable}.

    3.  Set `document`{.variable}\'s [history
        object](nav-history-apis.html#doc-history){#updating-the-traversable:doc-history-2}\'s
        [length](nav-history-apis.html#concept-history-length){#updating-the-traversable:concept-history-length}
        to `scriptHistoryLength`{.variable}.

    4.  Increment `completedNonchangingJobs`{.variable}.

19. Wait for `completedNonchangingJobs`{.variable} to equal
    `totalNonchangingJobs`{.variable}.

20. Set `traversable`{.variable}\'s [current session history
    step](document-sequences.html#tn-current-session-history-step){#updating-the-traversable:tn-current-session-history-step-3}
    to `targetStep`{.variable}.

21. Return \"`applied`\".

To [deactivate a document for a cross-document
navigation]{#deactivate-a-document-for-a-cross-document-navigation .dfn}
given a
[`Document`{#updating-the-traversable:document-2}](dom.html#document)
`displayedDocument`{.variable}, a [user navigation
involvement](#user-navigation-involvement){#updating-the-traversable:user-navigation-involvement-5}
`userNavigationInvolvement`{.variable}, a [session history
entry](#session-history-entry){#updating-the-traversable:session-history-entry-2}
`targetEntry`{.variable}, a
[`NavigationType`{#updating-the-traversable:navigationtype-2}](nav-history-apis.html#navigationtype)
`navigationType`{.variable}, and `afterPotentialUnloads`{.variable},
which is an algorithm that receives no arguments:

1.  Let `navigable`{.variable} be `displayedDocument`{.variable}\'s
    [node
    navigable](document-sequences.html#node-navigable){#updating-the-traversable:node-navigable}.

2.  Let `potentiallyTriggerViewTransition`{.variable} be false.

3.  Let `isBrowserUINavigation`{.variable} be true if
    `userNavigationInvolvement`{.variable} is
    \"[`browser UI`{#updating-the-traversable:uni-browser-ui}](#uni-browser-ui)\";
    otherwise false.

4.  Set `potentiallyTriggerViewTransition`{.variable} to the result of
    calling [can navigation trigger a cross-document
    view-transition?](https://drafts.csswg.org/css-view-transitions-2/#can-navigation-trigger-a-cross-document-view-transition){#updating-the-traversable:can-navigation-trigger-a-cross-document-view-transition
    x-internal="can-navigation-trigger-a-cross-document-view-transition"}
    given `displayedDocument`{.variable}, `targetEntry`{.variable}\'s
    [document](#she-document){#updating-the-traversable:she-document-12},
    `navigationType`{.variable}, and `isBrowserUINavigation`{.variable}.

5.  If `potentiallyTriggerViewTransition`{.variable} is false, then:

    1.  Let `firePageSwapBeforeUnload`{.variable} be the following step:

        1.  [Fire the `pageswap`
            event](#fire-the-pageswap-event){#updating-the-traversable:fire-the-pageswap-event}
            given `displayedDocument`{.variable},
            `targetEntry`{.variable}, `navigationType`{.variable}, and
            null.

    2.  [Set the ongoing
        navigation](#set-the-ongoing-navigation){#updating-the-traversable:set-the-ongoing-navigation-3}
        for `navigable`{.variable} to null.

        This allows new
        [navigations](#navigate){#updating-the-traversable:navigate-3}
        of `navigable`{.variable} to start, whereas during the traversal
        they were blocked.

    3.  [Unload a document and its
        descendants](document-lifecycle.html#unload-a-document-and-its-descendants){#updating-the-traversable:unload-a-document-and-its-descendants}
        given `displayedDocument`{.variable},
        `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-13},
        `afterPotentialUnloads`{.variable}, and
        `firePageSwapBeforeUnload`{.variable}.

6.  Otherwise, [queue a global
    task](webappapis.html#queue-a-global-task){#updating-the-traversable:queue-a-global-task-6}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#updating-the-traversable:navigation-and-traversal-task-source-6}
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#updating-the-traversable:nav-window-6}
    to run the steps:

    1.  Let `proceedWithNavigationAfterViewTransitionCapture`{.variable}
        be the following step:

        1.  [Append the following session history traversal
            steps](#tn-append-session-history-traversal-steps){#updating-the-traversable:tn-append-session-history-traversal-steps}
            to `navigable`{.variable}\'s [traversable
            navigable](document-sequences.html#nav-traversable){#updating-the-traversable:nav-traversable}:

            1.  [Set the ongoing
                navigation](#set-the-ongoing-navigation){#updating-the-traversable:set-the-ongoing-navigation-4}
                for `navigable`{.variable} to null.

                This allows new
                [navigations](#navigate){#updating-the-traversable:navigate-4}
                of `navigable`{.variable} to start, whereas during the
                traversal they were blocked.

            2.  [Unload a document and its
                descendants](document-lifecycle.html#unload-a-document-and-its-descendants){#updating-the-traversable:unload-a-document-and-its-descendants-2}
                given `displayedDocument`{.variable},
                `targetEntry`{.variable}\'s
                [document](#she-document){#updating-the-traversable:she-document-14},
                and `afterPotentialUnloads`{.variable}.

    2.  Let `viewTransition`{.variable} be the result of [setting up a
        cross-document
        view-transition](https://drafts.csswg.org/css-view-transitions-2/#setup-cross-document-view-transition){#updating-the-traversable:setting-up-a-cross-document-view-transition
        x-internal="setting-up-a-cross-document-view-transition"} given
        `displayedDocument`{.variable}, `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-15},
        `navigationType`{.variable}, and
        `proceedWithNavigationAfterViewTransitionCapture`{.variable}.

    3.  [Fire the `pageswap`
        event](#fire-the-pageswap-event){#updating-the-traversable:fire-the-pageswap-event-2}
        given `displayedDocument`{.variable}, `targetEntry`{.variable},
        `navigationType`{.variable}, and `viewTransition`{.variable}.

    4.  If `viewTransition`{.variable} is null, then run
        `proceedWithNavigationAfterViewTransitionCapture`{.variable}.

        In the case where a view transition started, the view
        transitions algorithms are responsible for calling
        `proceedWithNavigationAfterViewTransitionCapture`{.variable}.

To [fire the `pageswap` event]{#fire-the-pageswap-event .dfn} given a
[`Document`{#updating-the-traversable:document-3}](dom.html#document)
`displayedDocument`{.variable}, a [session history
entry](#session-history-entry){#updating-the-traversable:session-history-entry-3}
`targetEntry`{.variable}, a
[`NavigationType`{#updating-the-traversable:navigationtype-3}](nav-history-apis.html#navigationtype)
`navigationType`{.variable}, and a
[`ViewTransition`{#updating-the-traversable:viewtransition}](https://drafts.csswg.org/css-view-transitions/#viewtransition){x-internal="viewtransition"}-or-null
`viewTransition`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-8
    x-internal="assert"}: this is running as part of a
    [task](webappapis.html#concept-task){#updating-the-traversable:concept-task}
    queued on `displayedDocument`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#updating-the-traversable:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#updating-the-traversable:concept-agent-event-loop}.

2.  Let `navigation`{.variable} be `displayedDocument`{.variable}\'s
    [relevant global
    object](webappapis.html#concept-relevant-global){#updating-the-traversable:concept-relevant-global-2}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#updating-the-traversable:window-navigation-api-2}.

3.  Let `activation`{.variable} be null.

4.  If all of the following are true:

    - `targetEntry`{.variable}\'s
      [document](#she-document){#updating-the-traversable:she-document-16}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-traversable:concept-document-origin-3
      x-internal="concept-document-origin"} is [same
      origin](browsers.html#same-origin){#updating-the-traversable:same-origin-2}
      with `displayedDocument`{.variable}\'s
      [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-traversable:concept-document-origin-4
      x-internal="concept-document-origin"}; and

    - `targetEntry`{.variable}\'s
      [document](#she-document){#updating-the-traversable:she-document-17}\'s
      [was created via cross-origin
      redirects](dom.html#was-created-via-cross-origin-redirects){#updating-the-traversable:was-created-via-cross-origin-redirects}
      is false, or `targetEntry`{.variable}\'s
      [document](#she-document){#updating-the-traversable:she-document-18}\'s
      [latest
      entry](#latest-entry){#updating-the-traversable:latest-entry} is
      not null,

    then:

    1.  Let `destinationEntry`{.variable} be determined by switching on
        `navigationType`{.variable}:

        \"[`reload`{#updating-the-traversable:dom-navigationtype-reload-3}](nav-history-apis.html#dom-navigationtype-reload)\"
        :   The [current
            entry](nav-history-apis.html#navigation-current-entry){#updating-the-traversable:navigation-current-entry}
            of `navigation`{.variable}

        \"[`traverse`{#updating-the-traversable:dom-navigationtype-traverse-3}](nav-history-apis.html#dom-navigationtype-traverse)\"
        :   The
            [`NavigationHistoryEntry`{#updating-the-traversable:navigationhistoryentry}](nav-history-apis.html#navigationhistoryentry)
            in `navigation`{.variable}\'s [entry
            list](nav-history-apis.html#navigation-entry-list){#updating-the-traversable:navigation-entry-list}
            whose [session history
            entry](nav-history-apis.html#nhe-she){#updating-the-traversable:nhe-she}
            is `targetEntry`{.variable}

        \"[`push`{#updating-the-traversable:dom-navigationtype-push-2}](nav-history-apis.html#dom-navigationtype-push)\"\
        \"[`replace`{#updating-the-traversable:dom-navigationtype-replace-2}](nav-history-apis.html#dom-navigationtype-replace)\"

        :   A new
            [`NavigationHistoryEntry`{#updating-the-traversable:navigationhistoryentry-2}](nav-history-apis.html#navigationhistoryentry)
            in `displayedDocument`{.variable}\'s [relevant
            realm](webappapis.html#concept-relevant-realm){#updating-the-traversable:concept-relevant-realm}
            with its [session history
            entry](nav-history-apis.html#nhe-she){#updating-the-traversable:nhe-she-2}
            set to `targetEntry`{.variable}

    2.  Set `activation`{.variable} to a
        [new](https://webidl.spec.whatwg.org/#new){#updating-the-traversable:new
        x-internal="new"}
        [`NavigationActivation`{#updating-the-traversable:navigationactivation}](nav-history-apis.html#navigationactivation)
        created in `displayedDocument`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#updating-the-traversable:concept-relevant-realm-2},
        with

        [old entry](nav-history-apis.html#nav-activation-old-entry){#updating-the-traversable:nav-activation-old-entry}
        :   the [current
            entry](nav-history-apis.html#navigation-current-entry){#updating-the-traversable:navigation-current-entry-2}
            of `navigation`{.variable}

        [new entry](nav-history-apis.html#nav-activation-new-entry){#updating-the-traversable:nav-activation-new-entry}
        :   `destinationEntry`{.variable}

        [navigation type](nav-history-apis.html#nav-activation-navigation-type){#updating-the-traversable:nav-activation-navigation-type}
        :   `navigationType`{.variable}

    This means that a cross-origin redirect during a navigation would
    result in a null
    [`activation`{#updating-the-traversable:dom-pageswapevent-activation}](nav-history-apis.html#dom-pageswapevent-activation)
    in the old document\'s
    [`PageSwapEvent`{#updating-the-traversable:pageswapevent}](nav-history-apis.html#pageswapevent),
    unless the new document is being restored from
    [bfcache](#note-bfcache).

5.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-traversable:concept-event-fire
    x-internal="concept-event-fire"} named
    [`pageswap`{#updating-the-traversable:event-pageswap}](indices.html#event-pageswap)
    at `displayedDocument`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#updating-the-traversable:concept-relevant-global-3},
    using
    [`PageSwapEvent`{#updating-the-traversable:pageswapevent-2}](nav-history-apis.html#pageswapevent)
    with its
    [`activation`{#updating-the-traversable:dom-pageswapevent-activation-2}](nav-history-apis.html#dom-pageswapevent-activation)
    set to `activation`{.variable}, and its
    [`viewTransition`{#updating-the-traversable:dom-pageswapevent-viewtransition}](nav-history-apis.html#dom-pageswapevent-viewtransition)
    set to `viewTransition`{.variable}.

To [activate history entry]{#activate-history-entry .dfn} [session
history
entry](#session-history-entry){#updating-the-traversable:session-history-entry-4}
`entry`{.variable} for
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-8}
`navigable`{.variable}:

1.  [Save persisted
    state](#save-persisted-state){#updating-the-traversable:save-persisted-state}
    to the
    [navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-9}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-4}.

2.  Let `newDocument`{.variable} be `entry`{.variable}\'s
    [document](#she-document){#updating-the-traversable:she-document-19}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-9
    x-internal="assert"}: `newDocument`{.variable}\'s [is initial
    `about:blank`](dom.html#is-initial-about:blank){#updating-the-traversable:is-initial-about:blank}
    is false, i.e., we never traverse back to the [initial
    `about:blank`](dom.html#is-initial-about:blank){#updating-the-traversable:is-initial-about:blank-2}
    [`Document`{#updating-the-traversable:document-4}](dom.html#document)
    because it always gets [replaced](#navigate-convert-to-replace) when
    we navigate away from it.

4.  Set `navigable`{.variable}\'s [active session history
    entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-5}
    to `entry`{.variable}.

5.  [Make active](#make-active){#updating-the-traversable:make-active}
    `newDocument`{.variable}.

To [get the used step]{#getting-the-used-step .dfn} given a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-8}
`traversable`{.variable}, and a non-negative integer `step`{.variable},
perform the following steps. They return a non-negative integer.

1.  Let `steps`{.variable} be the result of [getting all used history
    steps](#getting-all-used-history-steps){#updating-the-traversable:getting-all-used-history-steps}
    within `traversable`{.variable}.

2.  Return the greatest
    [item](https://infra.spec.whatwg.org/#list-item){#updating-the-traversable:list-item-2
    x-internal="list-item"} in `steps`{.variable} that is less than or
    equal to `step`{.variable}.

    This caters for situations where there\'s no [session history
    entry](#session-history-entry){#updating-the-traversable:session-history-entry-5}
    with [step](#she-step){#updating-the-traversable:she-step-5}
    `step`{.variable}, due to the removal of a
    [navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-10}.

To [get the history object length and
index]{#getting-the-history-object-length-and-index .dfn} given a
[traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-9}
`traversable`{.variable}, and a non-negative integer `step`{.variable},
perform the following steps. They return a
[tuple](https://infra.spec.whatwg.org/#tuple){#updating-the-traversable:tuple
x-internal="tuple"} of two non-negative integers.

1.  Let `steps`{.variable} be the result of [getting all used history
    steps](#getting-all-used-history-steps){#updating-the-traversable:getting-all-used-history-steps-2}
    within `traversable`{.variable}.

2.  Let `scriptHistoryLength`{.variable} be the
    [size](https://infra.spec.whatwg.org/#list-size){#updating-the-traversable:list-size-3
    x-internal="list-size"} of `steps`{.variable}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-traversable:assert-10
    x-internal="assert"}: `steps`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#updating-the-traversable:list-contains-4
    x-internal="list-contains"} `step`{.variable}.

    It is assumed that `step`{.variable} has been adjusted by [getting
    the used
    step](#getting-the-used-step){#updating-the-traversable:getting-the-used-step-2}.

4.  Let `scriptHistoryIndex`{.variable} be the index of
    `step`{.variable} in `steps`{.variable}.

5.  Return (`scriptHistoryLength`{.variable},
    `scriptHistoryIndex`{.variable}).

To [get all navigables whose current session history entry will change
or
reload]{#get-all-navigables-whose-current-session-history-entry-will-change-or-reload
.dfn} given a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-10}
`traversable`{.variable}, and a non-negative integer
`targetStep`{.variable}, perform the following steps. They return a
[list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list
x-internal="list"} of
[navigables](document-sequences.html#navigable){#updating-the-traversable:navigable-11}.

1.  Let `results`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list-2
    x-internal="list"}.

2.  Let `navigablesToCheck`{.variable} be « `traversable`{.variable} ».

    This list is extended in the loop below.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-5
    x-internal="list-iterate"} `navigable`{.variable} of
    `navigablesToCheck`{.variable}:

    1.  Let `targetEntry`{.variable} be the result of [getting the
        target history
        entry](#getting-the-target-history-entry){#updating-the-traversable:getting-the-target-history-entry-2}
        given `navigable`{.variable} and `targetStep`{.variable}.

    2.  If `targetEntry`{.variable} is not `navigable`{.variable}\'s
        [current session history
        entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-5}
        or `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state-13}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-6}
        is true, then
        [append](https://infra.spec.whatwg.org/#list-append){#updating-the-traversable:list-append-2
        x-internal="list-append"} `navigable`{.variable} to
        `results`{.variable}.

    3.  If `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-20}
        is `navigable`{.variable}\'s
        [document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-3},
        and `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state-14}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-7}
        is false, then
        [extend](https://infra.spec.whatwg.org/#list-extend){#updating-the-traversable:list-extend
        x-internal="list-extend"} `navigablesToCheck`{.variable} with
        the [child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable}
        of `navigable`{.variable}.

        Adding [child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-2}
        to `navigablesToCheck`{.variable} means those navigables will
        also be checked by this loop. [Child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-3}
        are only checked if the `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-4}
        will not change as part of this traversal.

4.  Return `results`{.variable}.

To [get all navigables that only need history object length/index
update]{#getting-all-navigables-that-only-need-history-object-length/index-update
.dfn} given a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-11}
`traversable`{.variable}, and a non-negative integer
`targetStep`{.variable}, perform the following steps. They return a
[list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list-3
x-internal="list"} of
[navigables](document-sequences.html#navigable){#updating-the-traversable:navigable-12}.

Other
[navigables](document-sequences.html#navigable){#updating-the-traversable:navigable-13}
might not be impacted by the traversal. For example, if the response is
a 204, the currently active document will remain. Additionally, going
\'back\' after a 204 will change the [current session history
entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-6},
but the [active session history
entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-6}
will already be correct.

1.  Let `results`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list-4
    x-internal="list"}.

2.  Let `navigablesToCheck`{.variable} be « `traversable`{.variable} ».

    This list is extended in the loop below.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-6
    x-internal="list-iterate"} `navigable`{.variable} of
    `navigablesToCheck`{.variable}:

    1.  Let `targetEntry`{.variable} be the result of [getting the
        target history
        entry](#getting-the-target-history-entry){#updating-the-traversable:getting-the-target-history-entry-3}
        given `navigable`{.variable} and `targetStep`{.variable}.

    2.  If `targetEntry`{.variable} is `navigable`{.variable}\'s
        [current session history
        entry](document-sequences.html#nav-current-history-entry){#updating-the-traversable:nav-current-history-entry-7}
        and `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state-15}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-8}
        is false, then:

        1.  [Append](https://infra.spec.whatwg.org/#list-append){#updating-the-traversable:list-append-3
            x-internal="list-append"} `navigable`{.variable} to
            `results`{.variable}.

        2.  [Extend](https://infra.spec.whatwg.org/#list-extend){#updating-the-traversable:list-extend-2
            x-internal="list-extend"} `navigablesToCheck`{.variable}
            with `navigable`{.variable}\'s [child
            navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-4}.

            Adding [child
            navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-5}
            to `navigablesToCheck`{.variable} means those navigables
            will also be checked by this loop. [child
            navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-6}
            are only checked if the `navigable`{.variable}\'s [active
            document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-5}
            will not change as part of this traversal.

4.  Return `results`{.variable}.

To [get the target history entry]{#getting-the-target-history-entry
.dfn} given a
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-14}
`navigable`{.variable}, and a non-negative integer `step`{.variable},
perform the following steps. They return a [session history
entry](#session-history-entry){#updating-the-traversable:session-history-entry-6}.

1.  Let `entries`{.variable} be the result of [getting session history
    entries](#getting-session-history-entries){#updating-the-traversable:getting-session-history-entries}
    for `navigable`{.variable}.

2.  Return the
    [item](https://infra.spec.whatwg.org/#list-item){#updating-the-traversable:list-item-3
    x-internal="list-item"} in `entries`{.variable} that has the
    greatest [step](#she-step){#updating-the-traversable:she-step-6}
    less than or equal to `step`{.variable}.

::: {#example-getting-the-target-history-entry .example}
To see why [getting the target history
entry](#getting-the-target-history-entry){#updating-the-traversable:getting-the-target-history-entry-4}
returns the entry with the greatest
[step](#she-step){#updating-the-traversable:she-step-7} less than or
equal to the input step, consider the following [Jake
diagram](document-sequences.html#jake-diagram){#updating-the-traversable:jake-diagram}:

0

1

2

3

`top`

/t

/t#foo

`frames[0]`

/i-0-a

/i-0-b

For the input step 1, the target history entry for the `top` navigable
is the `/t` entry, whose
[step](#she-step){#updating-the-traversable:she-step-8} is 0, while the
target history entry for the `frames[0]` navigable is the `/i-0-b`
entry, whose [step](#she-step){#updating-the-traversable:she-step-9} is
1:

0

1

2

3

`top`

/t

/t#foo

`frames[0]`

/i-0-a

/i-0-b

Similarly, given the input step 3 we get the `top` entry whose
[step](#she-step){#updating-the-traversable:she-step-10} is 3, and the
`frames[0]` entry whose
[step](#she-step){#updating-the-traversable:she-step-11} is 1:

0

1

2

3

`top`

/t

/t#foo

`frames[0]`

/i-0-a

/i-0-b
:::

To [get all navigables that might experience a cross-document
traversal]{#getting-all-navigables-that-might-experience-a-cross-document-traversal
.dfn} given a [traversable
navigable](document-sequences.html#traversable-navigable){#updating-the-traversable:traversable-navigable-12}
`traversable`{.variable}, and a non-negative integer
`targetStep`{.variable}, perform the following steps. They return a
[list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list-5
x-internal="list"} of
[navigables](document-sequences.html#navigable){#updating-the-traversable:navigable-15}.

::: note
From `traversable`{.variable}\'s [session history traversal
queue](document-sequences.html#tn-session-history-traversal-queue){#updating-the-traversable:tn-session-history-traversal-queue-5}\'s
perspective, these documents are candidates for going cross-document
during the traversal described by `targetStep`{.variable}. They will not
experience a cross-document traversal if the status code for their
target document is HTTP 204 No Content.

Note that if a given
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-16}
might experience a cross-document traversal, this algorithm will return
[navigable](document-sequences.html#navigable){#updating-the-traversable:navigable-17}
but not its [child
navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-7}.
Those would end up
[unloaded](document-lifecycle.html#unload-a-document){#updating-the-traversable:unload-a-document},
not traversed.
:::

1.  Let `results`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#updating-the-traversable:list-6
    x-internal="list"}.

2.  Let `navigablesToCheck`{.variable} be « `traversable`{.variable} ».

    This list is extended in the loop below.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-traversable:list-iterate-7
    x-internal="list-iterate"} `navigable`{.variable} of
    `navigablesToCheck`{.variable}:

    1.  Let `targetEntry`{.variable} be the result of [getting the
        target history
        entry](#getting-the-target-history-entry){#updating-the-traversable:getting-the-target-history-entry-5}
        given `navigable`{.variable} and `targetStep`{.variable}.

    2.  If `targetEntry`{.variable}\'s
        [document](#she-document){#updating-the-traversable:she-document-21}
        is not `navigable`{.variable}\'s
        [document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-6}
        or `targetEntry`{.variable}\'s [document
        state](#she-document-state){#updating-the-traversable:she-document-state-16}\'s
        [reload
        pending](#document-state-reload-pending){#updating-the-traversable:document-state-reload-pending-9}
        is true, then
        [append](https://infra.spec.whatwg.org/#list-append){#updating-the-traversable:list-append-4
        x-internal="list-append"} `navigable`{.variable} to
        `results`{.variable}.

        Although `navigable`{.variable}\'s [active history
        entry](document-sequences.html#nav-active-history-entry){#updating-the-traversable:nav-active-history-entry-7}
        can change synchronously, the new entry will always have the
        same
        [`Document`{#updating-the-traversable:document-5}](dom.html#document),
        so accessing `navigable`{.variable}\'s
        [document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-7}
        is reliable.

    3.  Otherwise,
        [extend](https://infra.spec.whatwg.org/#list-extend){#updating-the-traversable:list-extend-3
        x-internal="list-extend"} `navigablesToCheck`{.variable} with
        `navigable`{.variable}\'s [child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-8}.

        Adding [child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-9}
        to `navigablesToCheck`{.variable} means those navigables will
        also be checked by this loop. [Child
        navigables](document-sequences.html#child-navigable){#updating-the-traversable:child-navigable-10}
        are only checked if the `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-traversable:nav-document-8}
        will not change as part of this traversal.

4.  Return `results`{.variable}.

##### [7.4.6.2]{.secno} Updating the document[](#updating-the-document){.self-link}

To [update document for history step
application]{#update-document-for-history-step-application .dfn} given a
[`Document`{#updating-the-document:document}](dom.html#document)
`document`{.variable}, a [session history
entry](#session-history-entry){#updating-the-document:session-history-entry}
`entry`{.variable}, a boolean `doNotReactivate`{.variable}, integers
`scriptHistoryLength`{.variable} and `scriptHistoryIndex`{.variable},
[`NavigationType`{#updating-the-document:navigationtype}](nav-history-apis.html#navigationtype)-or-null
`navigationType`{.variable}, an optional
[list](https://infra.spec.whatwg.org/#list){#updating-the-document:list
x-internal="list"} of [session history
entries](#session-history-entry){#updating-the-document:session-history-entry-2}
`entriesForNavigationAPI`{.variable}, and an optional [session history
entry](#session-history-entry){#updating-the-document:session-history-entry-3}
`previousEntryForActivation`{.variable}:

1.  Let `documentIsNew`{.variable} be true if `document`{.variable}\'s
    [latest entry](#latest-entry){#updating-the-document:latest-entry}
    is null; otherwise false.

2.  Let `documentsEntryChanged`{.variable} be true if
    `document`{.variable}\'s [latest
    entry](#latest-entry){#updating-the-document:latest-entry-2} is not
    `entry`{.variable}; otherwise false.

3.  Set `document`{.variable}\'s [history
    object](nav-history-apis.html#doc-history){#updating-the-document:doc-history}\'s
    [index](nav-history-apis.html#concept-history-index){#updating-the-document:concept-history-index}
    to `scriptHistoryIndex`{.variable}.

4.  Set `document`{.variable}\'s [history
    object](nav-history-apis.html#doc-history){#updating-the-document:doc-history-2}\'s
    [length](nav-history-apis.html#concept-history-length){#updating-the-document:concept-history-length}
    to `scriptHistoryLength`{.variable}.

5.  Let `navigation`{.variable} be `history`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#updating-the-document:window-navigation-api}.

6.  If `documentsEntryChanged`{.variable} is true, then:

    1.  Let `oldURL`{.variable} be `document`{.variable}\'s [latest
        entry](#latest-entry){#updating-the-document:latest-entry-3}\'s
        [URL](#she-url){#updating-the-document:she-url}.

    2.  Set `document`{.variable}\'s [latest
        entry](#latest-entry){#updating-the-document:latest-entry-4} to
        `entry`{.variable}.

    3.  [Restore the history object
        state](#restore-the-history-object-state){#updating-the-document:restore-the-history-object-state}
        given `document`{.variable} and `entry`{.variable}.

    4.  If `documentIsNew`{.variable} is false, then:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-document:assert
            x-internal="assert"}: `navigationType`{.variable} is not
            null.

        2.  [Update the navigation API entries for a same-document
            navigation](nav-history-apis.html#update-the-navigation-api-entries-for-a-same-document-navigation){#updating-the-document:update-the-navigation-api-entries-for-a-same-document-navigation}
            given `navigation`{.variable}, `entry`{.variable}, and
            `navigationType`{.variable}.

        3.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-document:concept-event-fire
            x-internal="concept-event-fire"} named
            [`popstate`{#updating-the-document:event-popstate}](indices.html#event-popstate)
            at `document`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-2},
            using
            [`PopStateEvent`{#updating-the-document:popstateevent}](nav-history-apis.html#popstateevent),
            with the
            [`state`{#updating-the-document:dom-popstateevent-state}](nav-history-apis.html#dom-popstateevent-state)
            attribute initialized to `document`{.variable}\'s [history
            object](nav-history-apis.html#doc-history){#updating-the-document:doc-history-3}\'s
            [state](nav-history-apis.html#concept-history-state){#updating-the-document:concept-history-state}
            and
            [`hasUAVisualTransition`{#updating-the-document:dom-popstateevent-hasuavisualtransition}](nav-history-apis.html#dom-popstateevent-hasuavisualtransition)
            initialized to true if a visual transition, to display a
            cached rendered state of the [latest
            entry](#latest-entry){#updating-the-document:latest-entry-5},
            was done by the user agent.

        4.  [Restore persisted
            state](#restore-persisted-state){#updating-the-document:restore-persisted-state}
            given `entry`{.variable}.

        5.  If `oldURL`{.variable}\'s
            [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#updating-the-document:concept-url-fragment
            x-internal="concept-url-fragment"} is not equal to
            `entry`{.variable}\'s
            [URL](#she-url){#updating-the-document:she-url-2}\'s
            [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#updating-the-document:concept-url-fragment-2
            x-internal="concept-url-fragment"}, then [queue a global
            task](webappapis.html#queue-a-global-task){#updating-the-document:queue-a-global-task}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#updating-the-document:dom-manipulation-task-source}
            given `document`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-3}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-document:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`hashchange`{#updating-the-document:event-hashchange}](indices.html#event-hashchange)
            at `document`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-4},
            using
            [`HashChangeEvent`{#updating-the-document:hashchangeevent}](nav-history-apis.html#hashchangeevent),
            with the
            [`oldURL`{#updating-the-document:dom-hashchangeevent-oldurl}](nav-history-apis.html#dom-hashchangeevent-oldurl)
            attribute initialized to the
            [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#updating-the-document:concept-url-serializer
            x-internal="concept-url-serializer"} of `oldURL`{.variable}
            and the
            [`newURL`{#updating-the-document:dom-hashchangeevent-newurl}](nav-history-apis.html#dom-hashchangeevent-newurl)
            attribute initialized to the
            [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#updating-the-document:concept-url-serializer-2
            x-internal="concept-url-serializer"} of
            `entry`{.variable}\'s
            [URL](#she-url){#updating-the-document:she-url-3}.

    5.  Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-document:assert-2
            x-internal="assert"}: `entriesForNavigationAPI`{.variable}
            is given.

        2.  [Restore persisted
            state](#restore-persisted-state){#updating-the-document:restore-persisted-state-2}
            given `entry`{.variable}.

        3.  [Initialize the navigation API entries for a new
            document](nav-history-apis.html#initialize-the-navigation-api-entries-for-a-new-document){#updating-the-document:initialize-the-navigation-api-entries-for-a-new-document}
            given `navigation`{.variable},
            `entriesForNavigationAPI`{.variable}, and
            `entry`{.variable}.

7.  If all the following are true:

    - `previousEntryForActivation`{.variable} is given;

    - `navigationType`{.variable} is non-null; and

    - `navigationType`{.variable} is
      \"[`reload`{#updating-the-document:dom-navigationtype-reload}](nav-history-apis.html#dom-navigationtype-reload)\"
      or `previousEntryForActivation`{.variable}\'s
      [document](#she-document){#updating-the-document:she-document} is
      not `document`{.variable},

    then:

    1.  If `navigation`{.variable}\'s
        [activation](nav-history-apis.html#navigation-activation){#updating-the-document:navigation-activation}
        is null, then set `navigation`{.variable}\'s
        [activation](nav-history-apis.html#navigation-activation){#updating-the-document:navigation-activation-2}
        to a new
        [`NavigationActivation`{#updating-the-document:navigationactivation}](nav-history-apis.html#navigationactivation)
        object in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#updating-the-document:concept-relevant-realm}.

    2.  Let `previousEntryIndex`{.variable} be the result of [getting
        the navigation API entry
        index](nav-history-apis.html#getting-the-navigation-api-entry-index){#updating-the-document:getting-the-navigation-api-entry-index}
        of `previousEntryForActivation`{.variable} within
        `navigation`{.variable}.

    3.  If `previousEntryIndex`{.variable} is non-negative, then set
        `activation`{.variable}\'s [old
        entry](nav-history-apis.html#nav-activation-old-entry){#updating-the-document:nav-activation-old-entry}
        to `navigation`{.variable}\'s [entry
        list](nav-history-apis.html#navigation-entry-list){#updating-the-document:navigation-entry-list}\[`previousEntryIndex`{.variable}\].

    4.  Otherwise, if all the following are true:

        - `navigationType`{.variable} is
          \"[`replace`{#updating-the-document:dom-navigationtype-replace}](nav-history-apis.html#dom-navigationtype-replace)\";

        - `previousEntryForActivation`{.variable}\'s [document
          state](#she-document-state){#updating-the-document:she-document-state}\'s
          [origin](#document-state-origin){#updating-the-document:document-state-origin}
          is [same
          origin](browsers.html#same-origin){#updating-the-document:same-origin}
          with `document`{.variable}\'s
          [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-document:concept-document-origin
          x-internal="concept-document-origin"}; and

        - `previousEntryForActivation`{.variable}\'s
          [document](#she-document){#updating-the-document:she-document-2}\'s
          [initial
          `about:blank`](dom.html#is-initial-about:blank){#updating-the-document:is-initial-about:blank}
          is false,

        then set `activation`{.variable}\'s [old
        entry](nav-history-apis.html#nav-activation-old-entry){#updating-the-document:nav-activation-old-entry-2}
        to a new
        [`NavigationHistoryEntry`{#updating-the-document:navigationhistoryentry}](nav-history-apis.html#navigationhistoryentry)
        in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#updating-the-document:concept-relevant-realm-2},
        whose [session history
        entry](nav-history-apis.html#nhe-she){#updating-the-document:nhe-she}
        is `previousEntryForActivation`{.variable}.

    5.  Set `activation`{.variable}\'s [new
        entry](nav-history-apis.html#nav-activation-new-entry){#updating-the-document:nav-activation-new-entry}
        to `navigation`{.variable}\'s [current
        entry](nav-history-apis.html#navigation-current-entry){#updating-the-document:navigation-current-entry}.

    6.  Set `activation`{.variable}\'s [navigation
        type](nav-history-apis.html#nav-activation-navigation-type){#updating-the-document:nav-activation-navigation-type}
        to `navigationType`{.variable}.

8.  If `documentIsNew`{.variable} is true, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-document:assert-3
        x-internal="assert"}: `document`{.variable}\'s [during-loading
        navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#updating-the-document:concept-document-navigation-id}
        is not null.

    2.  Invoke [WebDriver BiDi navigation
        committed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-committed){#updating-the-document:webdriver-bidi-navigation-committed
        x-internal="webdriver-bidi-navigation-committed"} with
        `navigable`{.variable} and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#updating-the-document:webdriver-bidi-navigation-status
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#updating-the-document:navigation-status-id
        x-internal="navigation-status-id"} is `document`{.variable}\'s
        [during-loading navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#updating-the-document:concept-document-navigation-id-2},
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#updating-the-document:navigation-status-status
        x-internal="navigation-status-status"} is
        \"[`committed`{#updating-the-document:navigation-status-committed}](https://w3c.github.io/webdriver-bidi/#navigation-status-committed){x-internal="navigation-status-committed"}\",
        and
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#updating-the-document:navigation-status-url
        x-internal="navigation-status-url"} is `document`{.variable}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#updating-the-document:the-document's-address
        x-internal="the-document's-address"}.

    3.  [Try to scroll to the
        fragment](#try-to-scroll-to-the-fragment){#updating-the-document:try-to-scroll-to-the-fragment}
        for `document`{.variable}.

    4.  At this point [scripts may run for the newly-created
        document]{#scripts-may-run-for-the-newly-created-document .dfn}
        `document`{.variable}.

9.  Otherwise, if `documentsEntryChanged`{.variable} is false and
    `doNotReactivate`{.variable} is false, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-document:assert-4
        x-internal="assert"}: `entriesForNavigationAPI`{.variable} is
        given.

    2.  [Reactivate](#reactivate-a-document){#updating-the-document:reactivate-a-document}
        `document`{.variable} given `entry`{.variable} and
        `entriesForNavigationAPI`{.variable}.

    `documentsEntryChanged`{.variable} can be false for one of two
    reasons: either we are restoring from [bfcache](#note-bfcache), or
    we are asynchronously finishing up a synchronous navigation which
    already synchronously set `document`{.variable}\'s [latest
    entry](#latest-entry){#updating-the-document:latest-entry-6}. The
    `doNotReactivate`{.variable} argument distinguishes between these
    two cases.

To [restore the history object state]{#restore-the-history-object-state
.dfn} given
[`Document`{#updating-the-document:document-2}](dom.html#document)
`document`{.variable} and [session history
entry](#session-history-entry){#updating-the-document:session-history-entry-4}
`entry`{.variable}:

1.  Let `targetRealm`{.variable} be `document`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#updating-the-document:concept-relevant-realm-3}.

2.  Let `state`{.variable} be
    [StructuredDeserialize](structured-data.html#structureddeserialize){#updating-the-document:structureddeserialize}(`entry`{.variable}\'s
    [classic history API
    state](#she-classic-history-api-state){#updating-the-document:she-classic-history-api-state},
    `targetRealm`{.variable}). If this throws an exception, catch it and
    let `state`{.variable} be null.

3.  Set `document`{.variable}\'s [history
    object](nav-history-apis.html#doc-history){#updating-the-document:doc-history-4}\'s
    [state](nav-history-apis.html#concept-history-state){#updating-the-document:concept-history-state-2}
    to `state`{.variable}.

To [make active]{#make-active .dfn} a
[`Document`{#updating-the-document:document-3}](dom.html#document)
`document`{.variable}:

1.  Let `window`{.variable} be `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-5}.

2.  Set `document`{.variable}\'s [browsing
    context](document-sequences.html#concept-document-bc){#updating-the-document:concept-document-bc}\'s
    [`WindowProxy`{#updating-the-document:windowproxy}](nav-history-apis.html#windowproxy)\'s
    [\[\[Window\]\]](nav-history-apis.html#concept-windowproxy-window){#updating-the-document:concept-windowproxy-window}
    internal slot value to `window`{.variable}.

3.  Set `document`{.variable}\'s [visibility
    state](interaction.html#visibility-state){#updating-the-document:visibility-state}
    to `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#updating-the-document:node-navigable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#updating-the-document:nav-traversable}\'s
    [system visibility
    state](document-sequences.html#system-visibility-state){#updating-the-document:system-visibility-state}.

4.  [Queue](https://w3c.github.io/performance-timeline/#queue-a-performanceentry){#updating-the-document:queue-a-performance-entry
    x-internal="queue-a-performance-entry"} a new
    [`VisibilityStateEntry`{#updating-the-document:visibilitystateentry}](interaction.html#visibilitystateentry)
    whose [visibility
    state](interaction.html#visibilitystateentry-state){#updating-the-document:visibilitystateentry-state}
    is `document`{.variable}\'s [visibility
    state](interaction.html#visibility-state){#updating-the-document:visibility-state-2}
    and whose
    [timestamp](interaction.html#visibilitystateentry-timestamp){#updating-the-document:visibilitystateentry-timestamp}
    is zero.

5.  Set `window`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#updating-the-document:relevant-settings-object}\'s
    [execution ready
    flag](webappapis.html#concept-environment-execution-ready-flag){#updating-the-document:concept-environment-execution-ready-flag}.

To [reactivate]{#reactivate-a-document .dfn dfn-for="Document"
export=""} a
[`Document`{#updating-the-document:document-4}](dom.html#document)
`document`{.variable} given a [session history
entry](#session-history-entry){#updating-the-document:session-history-entry-5}
`reactivatedEntry`{.variable} and a
[list](https://infra.spec.whatwg.org/#list){#updating-the-document:list-2
x-internal="list"} of [session history
entries](#session-history-entry){#updating-the-document:session-history-entry-6}
`entriesForNavigationAPI`{.variable}:

This algorithm updates `document`{.variable} after it has come out of
[bfcache](#note-bfcache), i.e., after it has been made [fully
active](document-sequences.html#fully-active){#updating-the-document:fully-active}
again. Other specifications that want to watch for this change to the
[fully
active](document-sequences.html#fully-active){#updating-the-document:fully-active-2}
state are encouraged to add steps into this algorithm, so that the
ordering of events that happen in effect of the change is clear.

1.  [[For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-document:list-iterate
    x-internal="list-iterate"} `formControl`{.variable} of form controls
    in `document`{.variable} with an [autofill field
    name](form-control-infrastructure.html#autofill-field-name){#updating-the-document:autofill-field-name}
    of
    \"[`off`{#updating-the-document:attr-fe-autocomplete-off}](form-control-infrastructure.html#attr-fe-autocomplete-off)\",
    invoke the [reset
    algorithm](form-control-infrastructure.html#concept-form-reset-control){#updating-the-document:concept-form-reset-control}
    for `formControl`{.variable}.]{#history-autocomplete}

2.  If `document`{.variable}\'s [suspended timer
    handles](document-lifecycle.html#suspended-timer-handles){#updating-the-document:suspended-timer-handles}
    is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#updating-the-document:list-is-empty
    x-internal="list-is-empty"}:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#updating-the-document:assert-5
        x-internal="assert"}: `document`{.variable}\'s [suspension
        time](document-lifecycle.html#suspension-time){#updating-the-document:suspension-time}
        is not zero.

    2.  Let `suspendDuration`{.variable} be the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#updating-the-document:current-high-resolution-time
        x-internal="current-high-resolution-time"} minus
        `document`{.variable}\'s [suspension
        time](document-lifecycle.html#suspension-time){#updating-the-document:suspension-time-2}.

    3.  Let `activeTimers`{.variable} be `document`{.variable}\'s
        [relevant global
        object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-6}\'s
        [map of active
        timers](timers-and-user-prompts.html#map-of-active-timers){#updating-the-document:map-of-active-timers}.

    4.  For each `handle`{.variable} in `document`{.variable}\'s
        [suspended timer
        handles](document-lifecycle.html#suspended-timer-handles){#updating-the-document:suspended-timer-handles-2},
        if `activeTimers`{.variable}\[`handle`{.variable}\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#updating-the-document:map-exists
        x-internal="map-exists"}, then increase
        `activeTimers`{.variable}\[`handle`{.variable}\] by
        `suspendDuration`{.variable}.

3.  [Update the navigation API entries for
    reactivation](nav-history-apis.html#update-the-navigation-api-entries-for-reactivation){#updating-the-document:update-the-navigation-api-entries-for-reactivation}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-7}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#updating-the-document:window-navigation-api-2},
    `entriesForNavigationAPI`{.variable}, and
    `reactivatedEntry`{.variable}.

4.  If `document`{.variable}\'s [current document
    readiness](dom.html#current-document-readiness){#updating-the-document:current-document-readiness}
    is \"`complete`\", and `document`{.variable}\'s [page
    showing](document-lifecycle.html#page-showing){#updating-the-document:page-showing}
    is false:

    1.  Set `document`{.variable}\'s [page
        showing](document-lifecycle.html#page-showing){#updating-the-document:page-showing-2}
        to true.

    2.  Set `document`{.variable}\'s [has been
        revealed](#has-been-revealed){#updating-the-document:has-been-revealed}
        to false.

    3.  [Update the visibility
        state](interaction.html#update-the-visibility-state){#updating-the-document:update-the-visibility-state}
        of `document`{.variable} to \"`visible`\".

    4.  [Fire a page transition
        event](nav-history-apis.html#fire-a-page-transition-event){#updating-the-document:fire-a-page-transition-event}
        named
        [`pageshow`{#updating-the-document:event-pageshow}](indices.html#event-pageshow)
        at `document`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-8}
        with true.

To [try to scroll to the fragment]{#try-to-scroll-to-the-fragment .dfn}
for a [`Document`{#updating-the-document:document-5}](dom.html#document)
`document`{.variable}, perform the following steps [in
parallel](infrastructure.html#in-parallel){#updating-the-document:in-parallel}:

1.  Wait for an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#updating-the-document:implementation-defined
    x-internal="implementation-defined"} amount of time. (This is
    intended to allow the user agent to optimize the user experience in
    the face of performance concerns.)

2.  [Queue a global
    task](webappapis.html#queue-a-global-task){#updating-the-document:queue-a-global-task-2}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#updating-the-document:navigation-and-traversal-task-source}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#updating-the-document:concept-relevant-global-9}
    to run these steps:

    1.  If `document`{.variable} has no parser, or its parser has
        [stopped
        parsing](parsing.html#stop-parsing){#updating-the-document:stop-parsing},
        or the user agent has reason to believe the user is no longer
        interested in scrolling to the
        [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#updating-the-document:concept-url-fragment-3
        x-internal="concept-url-fragment"}, then abort these steps.

    2.  [Scroll to the
        fragment](#scroll-to-the-fragment-identifier){#updating-the-document:scroll-to-the-fragment-identifier}
        given `document`{.variable}.

    3.  If `document`{.variable}\'s [indicated
        part](#the-indicated-part-of-the-document){#updating-the-document:the-indicated-part-of-the-document}
        is still null, then [try to scroll to the
        fragment](#try-to-scroll-to-the-fragment){#updating-the-document:try-to-scroll-to-the-fragment-2}
        for `document`{.variable}.

To [make document unsalvageable]{#make-document-unsalvageable .dfn
export=""}, given a
[`Document`{#updating-the-document:document-6}](dom.html#document)
`document`{.variable} and a string `reason`{.variable}:

1.  Let `details`{.variable} be a new [not restored reason
    details](nav-history-apis.html#nrr-details-struct){#updating-the-document:nrr-details-struct}
    whose
    [reason](nav-history-apis.html#nrr-details-reason){#updating-the-document:nrr-details-reason}
    is `reason`{.variable}.

2.  [Append](https://infra.spec.whatwg.org/#set-append){#updating-the-document:set-append
    x-internal="set-append"} `details`{.variable} to
    `document`{.variable}\'s [bfcache blocking
    details](dom.html#concept-document-bfcache-blocking-details){#updating-the-document:concept-document-bfcache-blocking-details}.

3.  Set `document`{.variable}\'s
    *[salvageable](document-lifecycle.html#concept-document-salvageable)*
    state to false.

To [build not restored reasons for document
state]{#build-not-restored-reasons-for-document-state .dfn} given
[`Document`{#updating-the-document:document-7}](dom.html#document)
`document`{.variable}:

1.  Let `notRestoredReasonsForDocument`{.variable} be a new [not
    restored
    reasons](nav-history-apis.html#nrr-struct){#updating-the-document:nrr-struct}.

2.  Set `notRestoredReasonsForDocument`{.variable}\'s
    [URL](nav-history-apis.html#nrr-url){#updating-the-document:nrr-url}
    to `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#updating-the-document:the-document's-address-2
    x-internal="the-document's-address"}.

3.  Let `container`{.variable} be `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#updating-the-document:node-navigable-2}\'s
    [container](document-sequences.html#nav-container){#updating-the-document:nav-container}.

4.  If `container`{.variable} is an
    [`iframe`{#updating-the-document:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    element:

    1.  Let `src`{.variable} be the empty string.

    2.  If `container`{.variable} has a
        [`src`{#updating-the-document:dom-iframe-src}](iframe-embed-object.html#dom-iframe-src)
        attribute:

        1.  Let `src`{.variable} be the result of
            [encoding-parsing-and-serializing a
            URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#updating-the-document:encoding-parsing-and-serializing-a-url}
            given `container`{.variable}\'s
            [`src`{#updating-the-document:attr-iframe-src}](iframe-embed-object.html#attr-iframe-src)
            attribute\'s value, relative to `container`{.variable}\'s
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-document:node-document
            x-internal="node-document"}.

        2.  If `src`{.variable} is failure, then set `src`{.variable} to
            `container`{.variable}\'s
            [`src`{#updating-the-document:attr-iframe-src-2}](iframe-embed-object.html#attr-iframe-src)
            attribute\'s value.

    3.  Set `notRestoredReasonsForDocument`{.variable}\'s
        [src](nav-history-apis.html#nrr-src){#updating-the-document:nrr-src}
        to `src`{.variable}.

    4.  Set `notRestoredReasonsForDocument`{.variable}\'s
        [id](nav-history-apis.html#nrr-id){#updating-the-document:nrr-id}
        to `container`{.variable}\'s
        [`id`{#updating-the-document:the-id-attribute}](dom.html#the-id-attribute)
        attribute\'s value, or the empty string if it has no such
        attribute.

    5.  Set `notRestoredReasonsForDocument`{.variable}\'s
        [name](nav-history-apis.html#nrr-name){#updating-the-document:nrr-name}
        to `container`{.variable}\'s
        [`name`{#updating-the-document:attr-iframe-name}](iframe-embed-object.html#attr-iframe-name)
        attribute\'s value, or the empty string if it has no such
        attribute.

5.  Set `notRestoredReasonsForDocument`{.variable}\'s
    [reasons](nav-history-apis.html#nrr-reasons){#updating-the-document:nrr-reasons}
    to a
    [clone](https://infra.spec.whatwg.org/#list-clone){#updating-the-document:list-clone
    x-internal="list-clone"} of `document`{.variable}\'s [bfcache
    blocking
    details](dom.html#concept-document-bfcache-blocking-details){#updating-the-document:concept-document-bfcache-blocking-details-2}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-document:list-iterate-2
    x-internal="list-iterate"} `navigable`{.variable} of
    `document`{.variable}\'s [document-tree child
    navigables](document-sequences.html#document-tree-child-navigables){#updating-the-document:document-tree-child-navigables}:

    1.  Let `childDocument`{.variable} be `navigable`{.variable}\'s
        [active
        document](document-sequences.html#nav-document){#updating-the-document:nav-document}.

    2.  [Build not restored reasons for document
        state](#build-not-restored-reasons-for-document-state){#updating-the-document:build-not-restored-reasons-for-document-state}
        given `childDocument`{.variable}.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#updating-the-document:list-append
        x-internal="list-append"} `childDocument`{.variable}\'s [not
        restored
        reasons](nav-history-apis.html#concept-document-nrr){#updating-the-document:concept-document-nrr}
        to `notRestoredReasonsForDocument`{.variable}\'s
        [children](nav-history-apis.html#nrr-children){#updating-the-document:nrr-children}.

7.  Set `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#updating-the-document:node-navigable-3}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#updating-the-document:nav-active-history-entry}\'s
    [document
    state](#she-document-state){#updating-the-document:she-document-state-2}\'s
    [not restored
    reasons](#document-state-not-restored-reasons){#updating-the-document:document-state-not-restored-reasons}
    to `notRestoredReasonsForDocument`{.variable}.

To [build not restored reasons for a top-level traversable and its
descendants]{#build-not-restored-reasons-for-a-top-level-traversable-and-its-descendants
.dfn} given [top-level
traversable](document-sequences.html#top-level-traversable){#updating-the-document:top-level-traversable}
`topLevelTraversable`{.variable}:

1.  [Build not restored reasons for document
    state](#build-not-restored-reasons-for-document-state){#updating-the-document:build-not-restored-reasons-for-document-state-2}
    given `topLevelTraversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#updating-the-document:nav-document-2}.

2.  Let `crossOriginDescendants`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#updating-the-document:list-3
    x-internal="list"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-document:list-iterate-3
    x-internal="list-iterate"} `childNavigable`{.variable} of
    `topLevelTraversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#updating-the-document:nav-document-3}\'s
    [descendant
    navigables](document-sequences.html#descendant-navigables){#updating-the-document:descendant-navigables}:

    1.  If `childNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-document:nav-document-4}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-document:concept-document-origin-2
        x-internal="concept-document-origin"} is not [same
        origin](browsers.html#same-origin){#updating-the-document:same-origin-2}
        with `topLevelTraversable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-document:nav-document-5}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-document:concept-document-origin-3
        x-internal="concept-document-origin"}, then
        [append](https://infra.spec.whatwg.org/#list-append){#updating-the-document:list-append-2
        x-internal="list-append"} `childNavigable`{.variable} to
        `crossOriginDescendants`{.variable}.

4.  Let `crossOriginDescendantsPreventsBfcache`{.variable} be false.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-document:list-iterate-4
    x-internal="list-iterate"} `crossOriginNavigable`{.variable} of
    `crossOriginDescendants`{.variable}:

    1.  Let `reasonsForCrossOriginChild`{.variable} be
        `crossOriginNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#updating-the-document:nav-document-6}\'s
        [document
        state](#document-state-2){#updating-the-document:document-state-2}\'s
        [not restored
        reasons](#document-state-not-restored-reasons){#updating-the-document:document-state-not-restored-reasons-2}.

    2.  If `reasonsForCrossOriginChild`{.variable}\'s
        [reasons](nav-history-apis.html#nrr-reasons){#updating-the-document:nrr-reasons-2}
        is not empty, set
        `crossOriginDescendantsPreventsBfcache`{.variable} to true.

    3.  Set `reasonsForCrossOriginChild`{.variable}\'s
        [URL](nav-history-apis.html#nrr-url){#updating-the-document:nrr-url-2}
        to null.

    4.  Set `reasonsForCrossOriginChild`{.variable}\'s
        [reasons](nav-history-apis.html#nrr-reasons){#updating-the-document:nrr-reasons-3}
        to null.

    5.  Set `reasonsForCrossOriginChild`{.variable}\'s
        [children](nav-history-apis.html#nrr-children){#updating-the-document:nrr-children-2}
        to null.

6.  If `crossOriginDescendantsPreventsBfcache`{.variable} is true, [make
    document
    unsalvageable](#make-document-unsalvageable){#updating-the-document:make-document-unsalvageable}
    given `topLevelTraversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#updating-the-document:nav-document-7}
    and
    \"[`masked`{#updating-the-document:blocking-masked}](nav-history-apis.html#blocking-masked)\".

##### [7.4.6.3]{.secno} Revealing the document[](#revealing-the-document){.self-link}

A [`Document`{#revealing-the-document:document}](dom.html#document) has
a boolean [has been revealed]{#has-been-revealed .dfn}, initially false.
It is used to ensure that the
[`pagereveal`{#revealing-the-document:event-pagereveal}](indices.html#event-pagereveal)
event is fired once for each activation of the
[`Document`{#revealing-the-document:document-2}](dom.html#document)
(once when it\'s rendered initially, and once for each
[reactivation](#reactivate-a-document){#revealing-the-document:reactivate-a-document}).

To [reveal]{#reveal .dfn} a
[`Document`{#revealing-the-document:document-3}](dom.html#document)
`document`{.variable}:

1.  If `document`{.variable}\'s [has been
    revealed](#has-been-revealed){#revealing-the-document:has-been-revealed}
    is true, then return.

2.  Set `document`{.variable}\'s [has been
    revealed](#has-been-revealed){#revealing-the-document:has-been-revealed-2}
    to true.

3.  Let `transition`{.variable} be the result of [resolving inbound
    cross-document
    view-transition](https://drafts.csswg.org/css-view-transitions-2/#resolve-inbound-cross-document-view-transition){#revealing-the-document:resolving-inbound-cross-document-view-transition
    x-internal="resolving-inbound-cross-document-view-transition"} for
    `document`{.variable}.

4.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#revealing-the-document:concept-event-fire
    x-internal="concept-event-fire"} named
    [`pagereveal`{#revealing-the-document:event-pagereveal-2}](indices.html#event-pagereveal)
    at `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#revealing-the-document:concept-relevant-global},
    using
    [`PageRevealEvent`{#revealing-the-document:pagerevealevent}](nav-history-apis.html#pagerevealevent)
    with its
    [`viewTransition`{#revealing-the-document:dom-pagerevealevent-viewtransition}](nav-history-apis.html#dom-pagerevealevent-viewtransition)
    set to `transition`{.variable}.

5.  If `transition`{.variable} is not null, then:

    1.  [Prepare to run
        script](webappapis.html#prepare-to-run-script){#revealing-the-document:prepare-to-run-script}
        given `document`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#revealing-the-document:relevant-settings-object}.

    2.  [Activate](https://drafts.csswg.org/css-view-transitions/#activate-view-transition){#revealing-the-document:activate-view-transition
        x-internal="activate-view-transition"} `transition`{.variable}.

    3.  [Clean up after running
        script](webappapis.html#clean-up-after-running-script){#revealing-the-document:clean-up-after-running-script}
        given `document`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#revealing-the-document:relevant-settings-object-2}.

    Activating a view transition might resolve/reject promises, so by
    wrapping the activation with prepare/cleanup we ensure those
    promises are handled before the next rendering step.

Though
[`pagereveal`{#revealing-the-document:event-pagereveal-3}](indices.html#event-pagereveal)
is guaranteed to be fired during the first [update the
rendering](webappapis.html#update-the-rendering){#revealing-the-document:update-the-rendering}
step that displays an up-to-date version of the page, user agents are
free to display a cached frame of the page before firing it. This
prevents the presence of a
[`pagereveal`{#revealing-the-document:event-pagereveal-4}](indices.html#event-pagereveal)
handler from delaying the presentation of such cached frame.

##### [7.4.6.4]{.secno} Scrolling to a fragment[](#scrolling-to-a-fragment){.self-link}

To [scroll to the fragment]{#scroll-to-the-fragment-identifier .dfn}
given a
[`Document`{#scrolling-to-a-fragment:document}](dom.html#document)
`document`{.variable}:

1.  If `document`{.variable}\'s [indicated
    part](#the-indicated-part-of-the-document){#scrolling-to-a-fragment:the-indicated-part-of-the-document}
    is null, then set `document`{.variable}\'s [target
    element](#target-element){#scrolling-to-a-fragment:target-element}
    to null.

2.  Otherwise, if `document`{.variable}\'s [indicated
    part](#the-indicated-part-of-the-document){#scrolling-to-a-fragment:the-indicated-part-of-the-document-2}
    is [top of the
    document](#top-of-the-document){#scrolling-to-a-fragment:top-of-the-document},
    then:

    1.  Set `document`{.variable}\'s [target
        element](#target-element){#scrolling-to-a-fragment:target-element-2}
        to null.

    2.  [Scroll to the beginning of the
        document](https://drafts.csswg.org/cssom-view/#scroll-to-the-beginning-of-the-document){#scrolling-to-a-fragment:scroll-to-the-beginning-of-the-document
        x-internal="scroll-to-the-beginning-of-the-document"} for
        `document`{.variable}.
        [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    3.  Return.

3.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#scrolling-to-a-fragment:assert
        x-internal="assert"}: `document`{.variable}\'s [indicated
        part](#the-indicated-part-of-the-document){#scrolling-to-a-fragment:the-indicated-part-of-the-document-3}
        is an element.

    2.  Let `target`{.variable} be `document`{.variable}\'s [indicated
        part](#the-indicated-part-of-the-document){#scrolling-to-a-fragment:the-indicated-part-of-the-document-4}.

    3.  Set `document`{.variable}\'s [target
        element](#target-element){#scrolling-to-a-fragment:target-element-3}
        to `target`{.variable}.

    4.  Run the [ancestor details revealing
        algorithm](interactive-elements.html#ancestor-details-revealing-algorithm){#scrolling-to-a-fragment:ancestor-details-revealing-algorithm}
        on `target`{.variable}.

    5.  Run the [ancestor hidden-until-found revealing
        algorithm](interaction.html#ancestor-hidden-until-found-revealing-algorithm){#scrolling-to-a-fragment:ancestor-hidden-until-found-revealing-algorithm}
        on `target`{.variable}.

    6.  [Scroll `target`{.variable} into
        view](https://drafts.csswg.org/cssom-view/#scroll-a-target-into-view){#scrolling-to-a-fragment:scroll-a-target-into-view
        x-internal="scroll-a-target-into-view"}, with *behavior* set to
        \"auto\", *block* set to \"start\", and *inline* set to
        \"nearest\". [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    7.  Run the [focusing
        steps](interaction.html#focusing-steps){#scrolling-to-a-fragment:focusing-steps}
        for `target`{.variable}, with the
        [`Document`{#scrolling-to-a-fragment:document-2}](dom.html#document)\'s
        [viewport](https://drafts.csswg.org/css2/#viewport){#scrolling-to-a-fragment:viewport
        x-internal="viewport"} as the `fallback target`{.variable}.

    8.  Move the [sequential focus navigation starting
        point](interaction.html#sequential-focus-navigation-starting-point){#scrolling-to-a-fragment:sequential-focus-navigation-starting-point}
        to `target`{.variable}.

A
[`Document`{#scrolling-to-a-fragment:document-3}](dom.html#document)\'s
[indicated part]{#the-indicated-part-of-the-document .dfn} is the one
that its
[URL](https://dom.spec.whatwg.org/#concept-document-url){#scrolling-to-a-fragment:the-document's-address
x-internal="the-document's-address"}\'s
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#scrolling-to-a-fragment:concept-url-fragment
x-internal="concept-url-fragment"} identifies, or null if the fragment
does not identify anything. The semantics of the
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#scrolling-to-a-fragment:concept-url-fragment-2
x-internal="concept-url-fragment"} in terms of mapping it to a node is
defined by the specification that defines the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#scrolling-to-a-fragment:mime-type
x-internal="mime-type"} used by the
[`Document`{#scrolling-to-a-fragment:document-4}](dom.html#document)
(for example, the processing of
[fragments](https://url.spec.whatwg.org/#concept-url-fragment){#scrolling-to-a-fragment:concept-url-fragment-3
x-internal="concept-url-fragment"} for [XML MIME
types](https://mimesniff.spec.whatwg.org/#xml-mime-type){#scrolling-to-a-fragment:xml-mime-type
x-internal="xml-mime-type"} is the responsibility of RFC7303).
[\[RFC7303\]](references.html#refsRFC7303)

There is also a [target element]{#target-element .dfn} for each
[`Document`{#scrolling-to-a-fragment:document-5}](dom.html#document),
which is used in defining the
[`:target`{#scrolling-to-a-fragment:selector-target}](semantics-other.html#selector-target)
pseudo-class and is updated by the above algorithm. It is initially
null.

For an [HTML
document](https://dom.spec.whatwg.org/#html-document){#scrolling-to-a-fragment:html-documents
x-internal="html-documents"} `document`{.variable}, its [indicated
part](#the-indicated-part-of-the-document){#scrolling-to-a-fragment:the-indicated-part-of-the-document-5}
is the result of [selecting the indicated
part](#select-the-indicated-part){#scrolling-to-a-fragment:select-the-indicated-part}
given `document`{.variable} and `document`{.variable}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#scrolling-to-a-fragment:the-document's-address-2
x-internal="the-document's-address"}.

To [select the indicated part]{#select-the-indicated-part .dfn} given a
[`Document`{#scrolling-to-a-fragment:document-6}](dom.html#document)
`document`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#scrolling-to-a-fragment:url
x-internal="url"} `url`{.variable}:

1.  If `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#scrolling-to-a-fragment:the-document's-address-3
    x-internal="the-document's-address"} does not
    [equal](https://url.spec.whatwg.org/#concept-url-equals){#scrolling-to-a-fragment:concept-url-equals
    x-internal="concept-url-equals"} `url`{.variable} with *[exclude
    fragments](https://url.spec.whatwg.org/#url-equals-exclude-fragments){x-internal="url-equals-exclude-fragments"}*
    set to true, then return null.

2.  Let `fragment`{.variable} be `url`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#scrolling-to-a-fragment:concept-url-fragment-4
    x-internal="concept-url-fragment"}.

3.  If `fragment`{.variable} is the empty string, then return the
    special value [top of the document]{#top-of-the-document .dfn}.

4.  Let `potentialIndicatedElement`{.variable} be the result of [finding
    a potential indicated
    element](#find-a-potential-indicated-element){#scrolling-to-a-fragment:find-a-potential-indicated-element}
    given `document`{.variable} and `fragment`{.variable}.

5.  If `potentialIndicatedElement`{.variable} is not null, then return
    `potentialIndicatedElement`{.variable}.

6.  Let `fragmentBytes`{.variable} be the result of
    [percent-decoding](https://url.spec.whatwg.org/#string-percent-decode){#scrolling-to-a-fragment:percent-decode
    x-internal="percent-decode"} `fragment`{.variable}.

7.  Let `decodedFragment`{.variable} be the result of running [UTF-8
    decode without
    BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#scrolling-to-a-fragment:utf-8-decode-without-bom
    x-internal="utf-8-decode-without-bom"} on
    `fragmentBytes`{.variable}.

8.  Set `potentialIndicatedElement`{.variable} to the result of [finding
    a potential indicated
    element](#find-a-potential-indicated-element){#scrolling-to-a-fragment:find-a-potential-indicated-element-2}
    given `document`{.variable} and `decodedFragment`{.variable}.

9.  If `potentialIndicatedElement`{.variable} is not null, then return
    `potentialIndicatedElement`{.variable}.

10. If `decodedFragment`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#scrolling-to-a-fragment:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the string `top`,
    then return the [top of the
    document](#top-of-the-document){#scrolling-to-a-fragment:top-of-the-document-2}.

11. Return null.

To [find a potential indicated
element]{#find-a-potential-indicated-element .dfn} given a
[`Document`{#scrolling-to-a-fragment:document-7}](dom.html#document)
`document`{.variable} and a string `fragment`{.variable}, run these
steps:

1.  If there is an element [in the document
    tree](https://dom.spec.whatwg.org/#in-a-document-tree){#scrolling-to-a-fragment:in-a-document-tree
    x-internal="in-a-document-tree"} whose
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#scrolling-to-a-fragment:root
    x-internal="root"} is `document`{.variable} and that has an
    [ID](https://dom.spec.whatwg.org/#concept-id){#scrolling-to-a-fragment:concept-id
    x-internal="concept-id"} equal to `fragment`{.variable}, then return
    the first such element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#scrolling-to-a-fragment:tree-order
    x-internal="tree-order"}.

2.  If there is an
    [`a`{#scrolling-to-a-fragment:the-a-element}](text-level-semantics.html#the-a-element)
    element [in the document
    tree](https://dom.spec.whatwg.org/#in-a-document-tree){#scrolling-to-a-fragment:in-a-document-tree-2
    x-internal="in-a-document-tree"} whose
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#scrolling-to-a-fragment:root-2
    x-internal="root"} is `document`{.variable} that has a
    [`name`{#scrolling-to-a-fragment:attr-a-name}](obsolete.html#attr-a-name)
    attribute whose value is equal to `fragment`{.variable}, then return
    the first such element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#scrolling-to-a-fragment:tree-order-2
    x-internal="tree-order"}.

3.  Return null.

##### [7.4.6.5]{.secno} Persisted history entry state[](#persisted-user-state-restoration){.self-link} {#persisted-user-state-restoration}

To [save persisted state]{#save-persisted-state .dfn} to a [session
history
entry](#session-history-entry){#persisted-user-state-restoration:session-history-entry}
`entry`{.variable}:

1.  Set the [scroll position
    data](#she-scroll-position){#persisted-user-state-restoration:she-scroll-position}
    of `entry`{.variable} to contain the scroll positions for all of
    `entry`{.variable}\'s
    [document](#she-document){#persisted-user-state-restoration:she-document}\'s
    [restorable scrollable
    regions](#restorable-scrollable-regions){#persisted-user-state-restoration:restorable-scrollable-regions}.

2.  Optionally, update `entry`{.variable}\'s [persisted user
    state](#she-other){#persisted-user-state-restoration:she-other} to
    reflect any state that the user agent wishes to persist, such as the
    values of form fields.

To [restore persisted state]{#restore-persisted-state .dfn} from a
[session history
entry](#session-history-entry){#persisted-user-state-restoration:session-history-entry-2}
`entry`{.variable}:

1.  If `entry`{.variable}\'s [scroll restoration
    mode](#she-scroll-restoration-mode){#persisted-user-state-restoration:she-scroll-restoration-mode}
    is
    \"[`auto`{#persisted-user-state-restoration:dom-scrollrestoration-auto}](#dom-scrollrestoration-auto)\",
    and `entry`{.variable}\'s
    [document](#she-document){#persisted-user-state-restoration:she-document-2}\'s
    [relevant global
    object](webappapis.html#concept-relevant-global){#persisted-user-state-restoration:concept-relevant-global}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#persisted-user-state-restoration:window-navigation-api}\'s
    [suppress normal scroll restoration during ongoing
    navigation](nav-history-apis.html#suppress-normal-scroll-restoration-during-ongoing-navigation){#persisted-user-state-restoration:suppress-normal-scroll-restoration-during-ongoing-navigation}
    is false, then [restore scroll position
    data](#restore-scroll-position-data){#persisted-user-state-restoration:restore-scroll-position-data}
    given `entry`{.variable}.

    The user agent not restoring scroll positions does not imply that
    scroll positions will be left at any particular value (e.g., (0,0)).
    The actual scroll position depends on the navigation type and the
    user agent\'s particular caching strategy. So web applications
    cannot assume any particular scroll position but rather are urged to
    set it to what they want it to be.

    If [suppress normal scroll restoration during ongoing
    navigation](nav-history-apis.html#suppress-normal-scroll-restoration-during-ongoing-navigation){#persisted-user-state-restoration:suppress-normal-scroll-restoration-during-ongoing-navigation-2}
    is true, then [restoring scroll position
    data](#restore-scroll-position-data){#persisted-user-state-restoration:restore-scroll-position-data-2}
    might still happen at a later point, as part of
    [finishing](nav-history-apis.html#navigateevent-finish){#persisted-user-state-restoration:navigateevent-finish}
    the relevant
    [`NavigateEvent`{#persisted-user-state-restoration:navigateevent}](nav-history-apis.html#navigateevent),
    or via a
    [`navigateEvent.scroll()`{#persisted-user-state-restoration:dom-navigateevent-scroll}](nav-history-apis.html#dom-navigateevent-scroll)
    method call.

2.  Optionally, update other aspects of `entry`{.variable}\'s
    [document](#she-document){#persisted-user-state-restoration:she-document-3}
    and its rendering, for instance values of form fields, that the user
    agent had previously recorded in `entry`{.variable}\'s [persisted
    user
    state](#she-other){#persisted-user-state-restoration:she-other-2}.

    This can even include updating the
    [`dir`{#persisted-user-state-restoration:attr-dir}](dom.html#attr-dir)
    attribute of
    [`textarea`{#persisted-user-state-restoration:the-textarea-element}](form-elements.html#the-textarea-element)
    elements or
    [`input`{#persisted-user-state-restoration:the-input-element}](input.html#the-input-element)
    elements whose
    [`type`{#persisted-user-state-restoration:attr-input-type}](input.html#attr-input-type)
    attribute is in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#persisted-user-state-restoration:text-(type=text)-state-and-search-state-(type=search)},
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#persisted-user-state-restoration:text-(type=text)-state-and-search-state-(type=search)-2},
    [Telephone](input.html#telephone-state-(type=tel)){#persisted-user-state-restoration:telephone-state-(type=tel)},
    [URL](input.html#url-state-(type=url)){#persisted-user-state-restoration:url-state-(type=url)},
    or
    [Email](input.html#email-state-(type=email)){#persisted-user-state-restoration:email-state-(type=email)}
    state, if the persisted state includes the directionality of user
    input in such controls.

    Restoring the value of form controls as part of this process does
    not fire any
    [`input`{#persisted-user-state-restoration:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    or
    [`change`{#persisted-user-state-restoration:event-change}](indices.html#event-change)
    events, but can trigger the `formStateRestoreCallback` of
    [form-associated custom
    elements](custom-elements.html#form-associated-custom-element){#persisted-user-state-restoration:form-associated-custom-element}.

------------------------------------------------------------------------

Each
[`Document`{#persisted-user-state-restoration:document}](dom.html#document)
has a boolean [has been scrolled by the
user]{#has-been-scrolled-by-the-user .dfn}, initially false. If the user
scrolls the document, the user agent must set that document\'s [has been
scrolled by the
user](#has-been-scrolled-by-the-user){#persisted-user-state-restoration:has-been-scrolled-by-the-user}
to true.

The [restorable scrollable regions]{#restorable-scrollable-regions .dfn}
of a
[`Document`{#persisted-user-state-restoration:document-2}](dom.html#document)
`document`{.variable} are `document`{.variable}\'s
[viewport](https://drafts.csswg.org/css2/#viewport){#persisted-user-state-restoration:viewport
x-internal="viewport"}, and all of `document`{.variable}\'s scrollable
regions excepting any [navigable
containers](document-sequences.html#navigable-container){#persisted-user-state-restoration:navigable-container}.

[Child
navigable](document-sequences.html#child-navigable){#persisted-user-state-restoration:child-navigable}
scroll restoration is handled as part of state restoration for the
[session history
entry](#session-history-entry){#persisted-user-state-restoration:session-history-entry-3}
for those
[navigables](document-sequences.html#navigable){#persisted-user-state-restoration:navigable}\'
[`Document`{#persisted-user-state-restoration:document-3}](dom.html#document)s.

To [restore scroll position data]{#restore-scroll-position-data .dfn}
given a [session history
entry](#session-history-entry){#persisted-user-state-restoration:session-history-entry-4}
`entry`{.variable}:

1.  Let `document`{.variable} be `entry`{.variable}\'s
    [document](#she-document){#persisted-user-state-restoration:she-document-4}.

2.  If `document`{.variable}\'s [has been scrolled by the
    user](#has-been-scrolled-by-the-user){#persisted-user-state-restoration:has-been-scrolled-by-the-user-2}
    is true, then the user agent should return.

3.  The user agent should attempt to use `entry`{.variable}\'s [scroll
    position
    data](#she-scroll-position){#persisted-user-state-restoration:she-scroll-position-2}
    to restore the scroll positions of `entry`{.variable}\'s
    [document](#she-document){#persisted-user-state-restoration:she-document-5}\'s
    [restorable scrollable
    regions](#restorable-scrollable-regions){#persisted-user-state-restoration:restorable-scrollable-regions-2}.
    The user agent may continue to attempt to do so periodically, until
    `document`{.variable}\'s [has been scrolled by the
    user](#has-been-scrolled-by-the-user){#persisted-user-state-restoration:has-been-scrolled-by-the-user-3}
    becomes true.

    This is formulated as an *attempt*, which is potentially repeated
    until success or until the user scrolls, due to the fact that
    relevant content indicated by the [scroll position
    data](#she-scroll-position){#persisted-user-state-restoration:she-scroll-position-3}
    might take some time to load from the network.

    Scroll restoration might be affected by scroll anchoring.
    [\[CSSSCROLLANCHORING\]](references.html#refsCSSSCROLLANCHORING)

[← 7.3 Infrastructure for sequences of
documents](document-sequences.html) --- [Table of Contents](./) --- [7.5
Document lifecycle →](document-lifecycle.html)
