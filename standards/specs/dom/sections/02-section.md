## [2. ]{.secno}[Events]{.content}[](#events){.self-link} {#events .heading .settled level="2"}

### [2.1. ]{.secno}[Introduction to \"DOM Events\"]{.content}[](#introduction-to-dom-events){.self-link} {#introduction-to-dom-events .heading .settled level="2.1"}

Throughout the web platform
[events](#concept-event){#ref-for-concept-event link-type="dfn"} are
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch
link-type="dfn"} to objects to signal an occurrence, such as network
activity or user interaction. These objects implement the
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget
link-type="idl"} interface and can therefore add [event
listeners](#concept-event-listener){#ref-for-concept-event-listener
link-type="dfn"} to observe
[events](#concept-event){#ref-for-concept-event① link-type="dfn"} by
calling
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener
link-type="idl"}:

``` {.lang-javascript .highlight}
obj.addEventListener("load", imgFetched)

function imgFetched(ev) {
  // great success
  …
}
```

[Event
listeners](#concept-event-listener){#ref-for-concept-event-listener①
link-type="dfn"} can be removed by utilizing the
[`removeEventListener()`{.idl}](#dom-eventtarget-removeeventlistener){#ref-for-dom-eventtarget-removeeventlistener
link-type="idl"} method, passing the same arguments.

Alternatively, [event
listeners](#concept-event-listener){#ref-for-concept-event-listener②
link-type="dfn"} can be removed by passing an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal
link-type="idl"} to
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener①
link-type="idl"} and calling
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort
link-type="idl"} on the controller owning the signal.

[Events](#concept-event){#ref-for-concept-event② link-type="dfn"} are
objects too and implement the [`Event`{.idl}](#event){#ref-for-event
link-type="idl"} interface (or a derived interface). In the example
above `ev`{.variable} is the
[event](#concept-event){#ref-for-concept-event③ link-type="dfn"}.
`ev`{.variable} is passed as an argument to the [event
listener](#concept-event-listener){#ref-for-concept-event-listener③
link-type="dfn"}'s
[callback](#event-listener-callback){#ref-for-event-listener-callback
link-type="dfn"} (typically a JavaScript Function as shown above).
[Event
listeners](#concept-event-listener){#ref-for-concept-event-listener④
link-type="dfn"} key off the
[event](#concept-event){#ref-for-concept-event④ link-type="dfn"}'s
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type link-type="idl"}
attribute value (\"`load`\" in the above example). The
[event](#concept-event){#ref-for-concept-event⑤ link-type="dfn"}'s
[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target
link-type="idl"} attribute value returns the object to which the
[event](#concept-event){#ref-for-concept-event⑥ link-type="dfn"} was
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①
link-type="dfn"} (`obj`{.variable} above).

Although [events](#concept-event){#ref-for-concept-event⑦
link-type="dfn"} are typically
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch②
link-type="dfn"} by the user agent as the result of user interaction or
the completion of some task, applications can
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch③
link-type="dfn"} [events](#concept-event){#ref-for-concept-event⑧
link-type="dfn"} themselves by using what are commonly known as
synthetic events:

``` {.lang-javascript .highlight}
// add an appropriate event listener
obj.addEventListener("cat", function(e) { process(e.detail) })

// create and dispatch the event
var event = new CustomEvent("cat", {"detail":{"hazcheeseburger":true}})
obj.dispatchEvent(event)
```

Apart from signaling, [events](#concept-event){#ref-for-concept-event⑨
link-type="dfn"} are sometimes also used to let an application control
what happens next in an operation. For instance as part of form
submission an [event](#concept-event){#ref-for-concept-event①⓪
link-type="dfn"} whose
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①
link-type="idl"} attribute value is \"`submit`\" is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch④
link-type="dfn"}. If this
[event](#concept-event){#ref-for-concept-event①① link-type="dfn"}'s
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault
link-type="idl"} method is invoked, form submission will be terminated.
Applications who wish to make use of this functionality through
[events](#concept-event){#ref-for-concept-event①② link-type="dfn"}
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑤
link-type="dfn"} by the application (synthetic events) can make use of
the return value of the
[`dispatchEvent()`{.idl}](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent
link-type="idl"} method:

``` {.lang-javascript .highlight}
if(obj.dispatchEvent(event)) {
  // event was not canceled, time for some magic
  …
}
```

When an [event](#concept-event){#ref-for-concept-event①③
link-type="dfn"} is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑥
link-type="dfn"} to an object that
[participates](#concept-tree-participate){#ref-for-concept-tree-participate①
link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑥
link-type="dfn"} (e.g., an
[element](#concept-element){#ref-for-concept-element link-type="dfn"}),
it can reach [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑤
link-type="dfn"} on that object's
[ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor①
link-type="dfn"} too. Effectively, all the object's [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor
link-type="dfn"} [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑥
link-type="dfn"} whose
[capture](#event-listener-capture){#ref-for-event-listener-capture
link-type="dfn"} is true are invoked, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order②
link-type="dfn"}. And then, if
[event](#concept-event){#ref-for-concept-event①④ link-type="dfn"}'s
[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles
link-type="idl"} is true, all the object's [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①
link-type="dfn"} [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑦
link-type="dfn"} whose
[capture](#event-listener-capture){#ref-for-event-listener-capture①
link-type="dfn"} is false are invoked, now in reverse [tree
order](#concept-tree-order){#ref-for-concept-tree-order③
link-type="dfn"}.

Let's look at an example of how
[events](#concept-event){#ref-for-concept-event①⑤ link-type="dfn"} work
in a [tree](#concept-tree){#ref-for-concept-tree⑦ link-type="dfn"}:

``` {.lang-markup .highlight}
<!doctype html>
<html>
 <head>
  <title>Boring example</title>
 </head>
 <body>
  <p>Hello <span id=x>world</span>!</p>
  <script>
   function test(e) {
     debug(e.target, e.currentTarget, e.eventPhase)
   }
   document.addEventListener("hey", test, {capture: true})
   document.body.addEventListener("hey", test)
   var ev = new Event("hey", {bubbles:true})
   document.getElementById("x").dispatchEvent(ev)
  </script>
 </body>
</html>
```

The `debug` function will be invoked twice. Each time the
[event](#concept-event){#ref-for-concept-event①⑥ link-type="dfn"}'s
[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target①
link-type="idl"} attribute value will be the `span`
[element](#concept-element){#ref-for-concept-element① link-type="dfn"}.
The first time
[`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget
link-type="idl"} attribute's value will be the
[document](#concept-document){#ref-for-concept-document
link-type="dfn"}, the second time the `body`
[element](#concept-element){#ref-for-concept-element② link-type="dfn"}.
[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase
link-type="idl"} attribute's value switches from
[`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase
link-type="idl"} to
[`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase
link-type="idl"}. If an [event
listener](#concept-event-listener){#ref-for-concept-event-listener⑧
link-type="dfn"} was registered for the `span`
[element](#concept-element){#ref-for-concept-element③ link-type="dfn"},
[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①
link-type="idl"} attribute's value would have been
[`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target
link-type="idl"}.

### [2.2. ]{.secno}[Interface [`Event`{.idl}](#event){#ref-for-event① link-type="idl"}]{.content}[](#interface-event){.self-link} {#interface-event .heading .settled level="2.2"}

``` {.idl .highlight .def}
[Exposed=*]
interface Event {
  constructor(DOMString type, optional EventInit eventInitDict = {});

  readonly attribute DOMString type;
  readonly attribute EventTarget? target;
  readonly attribute EventTarget? srcElement; // legacy
  readonly attribute EventTarget? currentTarget;
  sequence<EventTarget> composedPath();

  const unsigned short NONE = 0;
  const unsigned short CAPTURING_PHASE = 1;
  const unsigned short AT_TARGET = 2;
  const unsigned short BUBBLING_PHASE = 3;
  readonly attribute unsigned short eventPhase;

  undefined stopPropagation();
           attribute boolean cancelBubble; // legacy alias of .stopPropagation()
  undefined stopImmediatePropagation();

  readonly attribute boolean bubbles;
  readonly attribute boolean cancelable;
           attribute boolean returnValue;  // legacy
  undefined preventDefault();
  readonly attribute boolean defaultPrevented;
  readonly attribute boolean composed;

  [LegacyUnforgeable] readonly attribute boolean isTrusted;
  readonly attribute DOMHighResTimeStamp timeStamp;

  undefined initEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false); // legacy
};

dictionary EventInit {
  boolean bubbles = false;
  boolean cancelable = false;
  boolean composed = false;
};
```

An [`Event`{.idl}](#event){#ref-for-event② link-type="idl"} object is
simply named an [event]{#concept-event .dfn .dfn-paneled dfn-type="dfn"
export=""}. It allows for signaling that something has occurred, e.g.,
that an image has completed downloading.

A [potential event target]{#potential-event-target .dfn .dfn-paneled
dfn-type="dfn" export=""} is null or an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑤
link-type="idl"} object.

An [event](#concept-event){#ref-for-concept-event①⑦ link-type="dfn"} has
an associated [target]{#event-target .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} (a [potential event
target](#potential-event-target){#ref-for-potential-event-target
link-type="dfn"}). Unless stated otherwise it is null.

An [event](#concept-event){#ref-for-concept-event①⑧ link-type="dfn"} has
an associated [relatedTarget]{#event-relatedtarget .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} (a [potential event
target](#potential-event-target){#ref-for-potential-event-target①
link-type="dfn"}). Unless stated otherwise it is null.

Other specifications use
[relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget
link-type="dfn"} to define a `relatedTarget` attribute.
[\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

An [event](#concept-event){#ref-for-concept-event①⑨ link-type="dfn"} has
an associated [touch target list]{#event-touch-target-list .dfn
.dfn-paneled dfn-for="Event" dfn-type="dfn" export=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list
link-type="dfn"} of zero or more [potential event
targets](#potential-event-target){#ref-for-potential-event-target②
link-type="dfn"}). Unless stated otherwise it is the empty list.

The [touch target
list](#event-touch-target-list){#ref-for-event-touch-target-list
link-type="dfn"} is for the exclusive use of defining the
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent
link-type="idl"} interface and related interfaces.
[\[TOUCH-EVENTS\]](#biblio-touch-events "Touch Events"){link-type="biblio"}

An [event](#concept-event){#ref-for-concept-event②⓪ link-type="dfn"} has
an associated [path]{#event-path .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""}. A [path](#event-path){#ref-for-event-path
link-type="dfn"} is a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①
link-type="dfn"} of
[structs](https://infra.spec.whatwg.org/#struct){#ref-for-struct
link-type="dfn"}. Each
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct①
link-type="dfn"} consists of an [invocation
target]{#event-path-invocation-target .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑥
link-type="idl"} object), an
[invocation-target-in-shadow-tree]{#event-path-invocation-target-in-shadow-tree
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
boolean), a [shadow-adjusted target]{#event-path-shadow-adjusted-target
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[potential event
target](#potential-event-target){#ref-for-potential-event-target③
link-type="dfn"}), a [relatedTarget]{#event-path-relatedtarget .dfn
.dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[potential event
target](#potential-event-target){#ref-for-potential-event-target④
link-type="dfn"}), a [touch target list]{#event-path-touch-target-list
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list②
link-type="dfn"} of [potential event
targets](#potential-event-target){#ref-for-potential-event-target⑤
link-type="dfn"}), a
[root-of-closed-tree]{#event-path-root-of-closed-tree .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (a boolean), and a
[slot-in-closed-tree]{#event-path-slot-in-closed-tree .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (a boolean). A
[path](#event-path){#ref-for-event-path① link-type="dfn"} is initially
the empty list.

`event`{.variable}` = new `[`Event`](#dom-event-event){#ref-for-dom-event-event .idl-code link-type="constructor"}`(``type`{.variable}` [, ``eventInitDict`{.variable}`])`
:   Returns a new `event`{.variable} whose
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type③
    link-type="idl"} attribute value is set to `type`{.variable}. The
    `eventInitDict`{.variable} argument allows for setting the
    [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles②
    link-type="idl"} and
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable①
    link-type="idl"} attributes via object members of the same name.

`event`{.variable}` . `[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type④ link-type="idl"}
:   Returns the type of `event`{.variable}, e.g. \"`click`\",
    \"`hashchange`\", or \"`submit`\".

`event`{.variable}` . `[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target③ link-type="idl"}
:   Returns the object to which `event`{.variable} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑦
    link-type="dfn"} (its [target](#event-target){#ref-for-event-target
    link-type="dfn"}).

`event`{.variable}` . `[`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget② link-type="idl"}
:   Returns the object whose [event
    listener](#concept-event-listener){#ref-for-concept-event-listener⑨
    link-type="dfn"}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①
    link-type="dfn"} is currently being invoked.

`event`{.variable}` . `[`composedPath()`{.idl}](#dom-event-composedpath){#ref-for-dom-event-composedpath① link-type="idl"}
:   Returns the [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target
    link-type="dfn"} objects of `event`{.variable}'s
    [path](#event-path){#ref-for-event-path② link-type="dfn"} (objects
    on which listeners will be invoked), except for any
    [nodes](#concept-node){#ref-for-concept-node link-type="dfn"} in
    [shadow trees](#concept-shadow-tree){#ref-for-concept-shadow-tree
    link-type="dfn"} of which the [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root
    link-type="dfn"}'s [mode](#shadowroot-mode){#ref-for-shadowroot-mode
    link-type="dfn"} is \"`closed`\" that are not reachable from
    `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget③
    link-type="idl"}.

`event`{.variable}` . `[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase③ link-type="idl"}
:   Returns the [event](#concept-event){#ref-for-concept-event②①
    link-type="dfn"}'s phase, which is one of
    [`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none①
    link-type="idl"},
    [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase②
    link-type="idl"},
    [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target②
    link-type="idl"}, and
    [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase②
    link-type="idl"}.

`event`{.variable}` . `[`stopPropagation`](#dom-event-stoppropagation){#ref-for-dom-event-stoppropagation① .idl-code link-type="method"}`()`
:   When
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑧
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑧
    link-type="dfn"}, invoking this method prevents `event`{.variable}
    from reaching any objects other than the current object.

`event`{.variable}` . `[`stopImmediatePropagation`](#dom-event-stopimmediatepropagation){#ref-for-dom-event-stopimmediatepropagation① .idl-code link-type="method"}`()`
:   Invoking this method prevents `event`{.variable} from reaching any
    registered [event
    listeners](#concept-event-listener){#ref-for-concept-event-listener①⓪
    link-type="dfn"} after the current one finishes running and, when
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑨
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑨
    link-type="dfn"}, also prevents `event`{.variable} from reaching any
    other objects.

`event`{.variable}` . `[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles③ link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. True if `event`{.variable} goes through its
    [target](#event-target){#ref-for-event-target① link-type="dfn"}'s
    [ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor②
    link-type="dfn"} in reverse [tree
    order](#concept-tree-order){#ref-for-concept-tree-order④
    link-type="dfn"}; otherwise false.

`event`{.variable}` . `[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable② link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. Its return value does not always carry meaning, but
    true can indicate that part of the operation during which
    `event`{.variable} was
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⓪
    link-type="dfn"}, can be canceled by invoking the
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault②
    link-type="idl"} method.

`event`{.variable}` . `[`preventDefault`](#dom-event-preventdefault){#ref-for-dom-event-preventdefault③ .idl-code link-type="method"}`()`
:   If invoked when the
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable③
    link-type="idl"} attribute value is true, and while executing a
    listener for the `event`{.variable} with
    [`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive
    link-type="idl"} set to false, signals to the operation that caused
    `event`{.variable} to be
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①①
    link-type="dfn"} that it needs to be canceled.

`event`{.variable}` . `[`defaultPrevented`{.idl}](#dom-event-defaultprevented){#ref-for-dom-event-defaultprevented① link-type="idl"}
:   Returns true if
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault④
    link-type="idl"} was invoked successfully to indicate cancelation;
    otherwise false.

`event`{.variable}` . `[`composed`{.idl}](#dom-event-composed){#ref-for-dom-event-composed① link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. True if `event`{.variable} invokes listeners past a
    [`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①
    link-type="dfn"} that is the
    [root](#concept-tree-root){#ref-for-concept-tree-root③
    link-type="dfn"} of its
    [target](#event-target){#ref-for-event-target② link-type="dfn"};
    otherwise false.

`event`{.variable}` . `[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted① link-type="idl"}
:   Returns true if `event`{.variable} was
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①②
    link-type="dfn"} by the user agent, and false otherwise.

`event`{.variable}` . `[`timeStamp`{.idl}](#dom-event-timestamp){#ref-for-dom-event-timestamp① link-type="idl"}
:   Returns the `event`{.variable}'s timestamp as the number of
    milliseconds measured relative to the occurrence.

The [`type`]{#dom-event-type .dfn .dfn-paneled .idl-code dfn-for="Event"
dfn-type="attribute" export=""} attribute must return the value it was
initialized to. When an [event](#concept-event){#ref-for-concept-event②②
link-type="dfn"} is created the attribute must be initialized to the
empty string.

The [`target`]{#dom-event-target .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this
link-type="dfn"}'s [target](#event-target){#ref-for-event-target③
link-type="dfn"}.

The [`srcElement`]{#dom-event-srcelement .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
link-type="dfn"}'s [target](#event-target){#ref-for-event-target④
link-type="dfn"}.

The [`currentTarget`]{#dom-event-currenttarget .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="attribute" export=""} attribute must
return the value it was initialized to. When an
[event](#concept-event){#ref-for-concept-event②③ link-type="dfn"} is
created the attribute must be initialized to null.

::: {.algorithm algorithm="composedPath()" algorithm-for="Event"}
The [`composedPath()`]{#dom-event-composedpath .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are:

1.  Let `composedPath`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list③
    link-type="dfn"}.

2.  Let `path`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
    link-type="dfn"}'s [path](#event-path){#ref-for-event-path③
    link-type="dfn"}.

3.  If `path`{.variable} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty
    link-type="dfn"}, then return `composedPath`{.variable}.

4.  Let `currentTarget`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget④
    link-type="idl"} attribute value.

5.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert
    link-type="dfn"}: `currentTarget`{.variable} is an
    [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑦
    link-type="idl"} object.

6.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append
    link-type="dfn"} `currentTarget`{.variable} to
    `composedPath`{.variable}.

7.  Let `currentTargetIndex`{.variable} be 0.

8.  Let `currentTargetHiddenSubtreeLevel`{.variable} be 0.

9.  Let `index`{.variable} be `path`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size
    link-type="dfn"} − 1.

10. While `index`{.variable} is greater than or equal to 0:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree
        link-type="dfn"} is true, then increase
        `currentTargetHiddenSubtreeLevel`{.variable} by 1.

    2.  If `path`{.variable}\[`index`{.variable}\]\'s [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target①
        link-type="dfn"} is `currentTarget`{.variable}, then set
        `currentTargetIndex`{.variable} to `index`{.variable} and
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break
        link-type="dfn"}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree
        link-type="dfn"} is true, then decrease
        `currentTargetHiddenSubtreeLevel`{.variable} by 1.

    4.  Decrease `index`{.variable} by 1.

11. Let `currentHiddenLevel`{.variable} and `maxHiddenLevel`{.variable}
    be `currentTargetHiddenSubtreeLevel`{.variable}.

12. Set `index`{.variable} to `currentTargetIndex`{.variable} − 1.

13. While `index`{.variable} is greater than or equal to 0:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree①
        link-type="dfn"} is true, then increase
        `currentHiddenLevel`{.variable} by 1.

    2.  If `currentHiddenLevel`{.variable} is less than or equal to
        `maxHiddenLevel`{.variable}, then
        [prepend](https://infra.spec.whatwg.org/#list-prepend){#ref-for-list-prepend
        link-type="dfn"} `path`{.variable}\[`index`{.variable}\]\'s
        [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target②
        link-type="dfn"} to `composedPath`{.variable}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree①
        link-type="dfn"} is true:

        1.  Decrease `currentHiddenLevel`{.variable} by 1.

        2.  If `currentHiddenLevel`{.variable} is less than
            `maxHiddenLevel`{.variable}, then set
            `maxHiddenLevel`{.variable} to
            `currentHiddenLevel`{.variable}.

    4.  Decrease `index`{.variable} by 1.

14. Set `currentHiddenLevel`{.variable} and `maxHiddenLevel`{.variable}
    to `currentTargetHiddenSubtreeLevel`{.variable}.

15. Set `index`{.variable} to `currentTargetIndex`{.variable} + 1.

16. While `index`{.variable} is less than `path`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size①
    link-type="dfn"}:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree②
        link-type="dfn"} is true, then increase
        `currentHiddenLevel`{.variable} by 1.

    2.  If `currentHiddenLevel`{.variable} is less than or equal to
        `maxHiddenLevel`{.variable}, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①
        link-type="dfn"} `path`{.variable}\[`index`{.variable}\]\'s
        [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target③
        link-type="dfn"} to `composedPath`{.variable}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree②
        link-type="dfn"} is true:

        1.  Decrease `currentHiddenLevel`{.variable} by 1.

        2.  If `currentHiddenLevel`{.variable} is less than
            `maxHiddenLevel`{.variable}, then set
            `maxHiddenLevel`{.variable} to
            `currentHiddenLevel`{.variable}.

    4.  Increase `index`{.variable} by 1.

17. Return `composedPath`{.variable}.
:::

The [`eventPhase`]{#dom-event-eventphase .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to, which must be one of the following:

[`NONE`]{#dom-event-none .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 0)
:   [Events](#concept-event){#ref-for-concept-event②④ link-type="dfn"}
    not currently
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①③
    link-type="dfn"} are in this phase.

[`CAPTURING_PHASE`]{#dom-event-capturing_phase .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 1)
:   When an [event](#concept-event){#ref-for-concept-event②⑤
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①④
    link-type="dfn"} to an object that
    [participates](#concept-tree-participate){#ref-for-concept-tree-participate②
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①⓪
    link-type="dfn"} it will be in this phase before it reaches its
    [target](#event-target){#ref-for-event-target⑤ link-type="dfn"}.

[`AT_TARGET`]{#dom-event-at_target .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 2)
:   When an [event](#concept-event){#ref-for-concept-event②⑥
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑤
    link-type="dfn"} it will be in this phase on its
    [target](#event-target){#ref-for-event-target⑥ link-type="dfn"}.

[`BUBBLING_PHASE`]{#dom-event-bubbling_phase .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 3)
:   When an [event](#concept-event){#ref-for-concept-event②⑦
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑥
    link-type="dfn"} to an object that
    [participates](#concept-tree-participate){#ref-for-concept-tree-participate③
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①①
    link-type="dfn"} it will be in this phase after it reaches its
    [target](#event-target){#ref-for-event-target⑦ link-type="dfn"}.

Initially the attribute must be initialized to
[`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none②
link-type="idl"}.

------------------------------------------------------------------------

Each [event](#concept-event){#ref-for-concept-event②⑧ link-type="dfn"}
has the following associated flags that are all initially unset:

- [stop propagation flag]{#stop-propagation-flag .dfn .dfn-paneled
  dfn-for="Event" dfn-type="dfn" export=""}
- [stop immediate propagation flag]{#stop-immediate-propagation-flag
  .dfn .dfn-paneled dfn-for="Event" dfn-type="dfn" export=""}
- [canceled flag]{#canceled-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [in passive listener flag]{#in-passive-listener-flag .dfn .dfn-paneled
  dfn-for="Event" dfn-type="dfn" export=""}
- [composed flag]{#composed-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [initialized flag]{#initialized-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [dispatch flag]{#dispatch-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}

The [`stopPropagation()`]{#dom-event-stoppropagation .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are
to set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag
link-type="dfn"}.

The [`cancelBubble`]{#dom-event-cancelbubble .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag①
link-type="dfn"} is set; otherwise false.

The
[`cancelBubble`{.idl}](#dom-event-cancelbubble){#ref-for-dom-event-cancelbubble①
link-type="idl"} setter steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag②
link-type="dfn"} if the given value is true; otherwise do nothing.

The [`stopImmediatePropagation()`]{#dom-event-stopimmediatepropagation
.dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="method" export=""}
method steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag③
link-type="dfn"} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧
link-type="dfn"}'s [stop immediate propagation
flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag
link-type="dfn"}.

The [`bubbles`]{#dom-event-bubbles .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} and
[`cancelable`]{#dom-event-cancelable .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attributes must return
the values they were initialized to.

To [set the canceled flag]{#set-the-canceled-flag .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an
[event](#concept-event){#ref-for-concept-event②⑨ link-type="dfn"}
`event`{.variable}, if `event`{.variable}'s
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable④
link-type="idl"} attribute value is true and `event`{.variable}'s [in
passive listener
flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag
link-type="dfn"} is unset, then set `event`{.variable}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag link-type="dfn"}, and do
nothing otherwise.

The [`returnValue`]{#dom-event-returnvalue .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return false if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨
link-type="dfn"}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag① link-type="dfn"} is set;
otherwise true.

The
[`returnValue`{.idl}](#dom-event-returnvalue){#ref-for-dom-event-returnvalue①
link-type="idl"} setter steps are to [set the canceled
flag](#set-the-canceled-flag){#ref-for-set-the-canceled-flag
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪
link-type="dfn"} if the given value is false; otherwise do nothing.

The [`preventDefault()`]{#dom-event-preventdefault .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are
to [set the canceled
flag](#set-the-canceled-flag){#ref-for-set-the-canceled-flag①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①
link-type="dfn"}.

There are scenarios where invoking
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑤
link-type="idl"} has no effect. User agents are encouraged to log the
precise cause in a developer console, to aid debugging.

The [`defaultPrevented`]{#dom-event-defaultprevented .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="attribute" export=""} getter steps
are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②
link-type="dfn"}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag② link-type="dfn"} is set;
otherwise false.

The [`composed`]{#dom-event-composed .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③
link-type="dfn"}'s [composed
flag](#composed-flag){#ref-for-composed-flag link-type="dfn"} is set;
otherwise false.

------------------------------------------------------------------------

The [`isTrusted`]{#dom-event-istrusted .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to. When an
[event](#concept-event){#ref-for-concept-event③⓪ link-type="dfn"} is
created the attribute must be initialized to false.

[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted②
link-type="idl"} is a convenience that indicates whether an
[event](#concept-event){#ref-for-concept-event③① link-type="dfn"} is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑦
link-type="dfn"} by the user agent (as opposed to using
[`dispatchEvent()`{.idl}](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent①
link-type="idl"}). The sole legacy exception is
[`click()`{.idl}](https://html.spec.whatwg.org/multipage/interaction.html#dom-click){#ref-for-dom-click
link-type="idl"}, which causes the user agent to dispatch an
[event](#concept-event){#ref-for-concept-event③② link-type="dfn"} whose
[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted③
link-type="idl"} attribute is initialized to false.

The [`timeStamp`]{#dom-event-timestamp .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to.

------------------------------------------------------------------------

::: {.algorithm algorithm="initialize" algorithm-for="Event"}
To [initialize]{#concept-event-initialize .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} an `event`{.variable}, with
`type`{.variable}, `bubbles`{.variable}, and `cancelable`{.variable},
run these steps:

1.  Set `event`{.variable}'s [initialized
    flag](#initialized-flag){#ref-for-initialized-flag link-type="dfn"}.

2.  Unset `event`{.variable}'s [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag④
    link-type="dfn"}, [stop immediate propagation
    flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag①
    link-type="dfn"}, and [canceled
    flag](#canceled-flag){#ref-for-canceled-flag③ link-type="dfn"}.

3.  Set `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted④
    link-type="idl"} attribute to false.

4.  Set `event`{.variable}'s
    [target](#event-target){#ref-for-event-target⑧ link-type="dfn"} to
    null.

5.  Set `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑤
    link-type="idl"} attribute to `type`{.variable}.

6.  Set `event`{.variable}'s
    [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles④
    link-type="idl"} attribute to `bubbles`{.variable}.

7.  Set `event`{.variable}'s
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑤
    link-type="idl"} attribute to `cancelable`{.variable}.
:::

::: {.algorithm algorithm="initEvent(type, bubbles, cancelable)" algorithm-for="Event"}
The
[`initEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`)`]{#dom-event-initevent
.dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="method" export=""
lt="initEvent(type, bubbles, cancelable)|initEvent(type, bubbles)|initEvent(type)"}
method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④
    link-type="dfn"}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag link-type="dfn"} is
    set, then return.

2.  [Initialize](#concept-event-initialize){#ref-for-concept-event-initialize
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤
    link-type="dfn"} with `type`{.variable}, `bubbles`{.variable}, and
    `cancelable`{.variable}.
:::

[`initEvent()`{.idl}](#dom-event-initevent){#ref-for-dom-event-initevent①
link-type="idl"} is redundant with
[event](#concept-event){#ref-for-concept-event③③ link-type="dfn"}
constructors and incapable of setting
[`composed`{.idl}](#dom-event-composed){#ref-for-dom-event-composed②
link-type="idl"}. It has to be supported for legacy content.

### [2.3. ]{.secno}[Legacy extensions to the [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window link-type="idl"} interface]{.content}[](#interface-window-extensions){.self-link} {#interface-window-extensions .heading .settled level="2.3"}

``` {.idl .highlight .def}
partial interface Window {
  [Replaceable] readonly attribute (Event or undefined) event; // legacy
};
```

Each
[`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window②
link-type="idl"} object has an associated [current
event]{#window-current-event .dfn .dfn-paneled dfn-for="Window"
dfn-type="dfn" noexport=""} (undefined or an
[`Event`{.idl}](#event){#ref-for-event④ link-type="idl"} object). Unless
stated otherwise it is undefined.

The [`event`]{#dom-window-event .dfn .dfn-paneled .idl-code
dfn-for="Window" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥
link-type="dfn"}'s [current
event](#window-current-event){#ref-for-window-current-event
link-type="dfn"}.

Web developers are strongly encouraged to instead rely on the
[`Event`{.idl}](#event){#ref-for-event⑤ link-type="idl"} object passed
to event listeners, as that will result in more portable code. This
attribute is not available in workers or worklets, and is inaccurate for
events dispatched in [shadow
trees](#concept-shadow-tree){#ref-for-concept-shadow-tree①
link-type="dfn"}.

### [2.4. ]{.secno}[Interface [`CustomEvent`{.idl}](#customevent){#ref-for-customevent link-type="idl"}]{.content}[](#interface-customevent){.self-link} {#interface-customevent .heading .settled level="2.4"}

``` {.idl .highlight .def}
[Exposed=*]
interface CustomEvent : Event {
  constructor(DOMString type, optional CustomEventInit eventInitDict = {});

  readonly attribute any detail;

  undefined initCustomEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false, optional any detail = null); // legacy
};

dictionary CustomEventInit : EventInit {
  any detail = null;
};
```

[Events](#concept-event){#ref-for-concept-event③④ link-type="dfn"} using
the [`CustomEvent`{.idl}](#customevent){#ref-for-customevent①
link-type="idl"} interface can be used to carry custom data.

`event`{.variable}` = new `[`CustomEvent`](#dom-customevent-customevent){#ref-for-dom-customevent-customevent .idl-code link-type="constructor"}`(``type`{.variable}` [, ``eventInitDict`{.variable}`])`
:   Works analogously to the constructor for
    [`Event`{.idl}](#event){#ref-for-event⑦ link-type="idl"} except that
    the `eventInitDict`{.variable} argument now allows for setting the
    [`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail①
    link-type="idl"} attribute too.

`event`{.variable}` . `[`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail② link-type="idl"}
:   Returns any custom data `event`{.variable} was created with.
    Typically used for synthetic events.

The [`detail`]{#dom-customevent-detail .dfn .dfn-paneled .idl-code
dfn-for="CustomEvent" dfn-type="attribute" export=""} attribute must
return the value it was initialized to.

::: {.algorithm algorithm="initCustomEvent(type, bubbles, cancelable, detail)" algorithm-for="CustomEvent"}
The
[`initCustomEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`, ``detail`{.variable}`)`]{#dom-customevent-initcustomevent
.dfn .dfn-paneled .idl-code dfn-for="CustomEvent" dfn-type="method"
export=""
lt="initCustomEvent(type, bubbles, cancelable, detail)|initCustomEvent(type, bubbles, cancelable)|initCustomEvent(type, bubbles)|initCustomEvent(type)"}
method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦
    link-type="dfn"}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag① link-type="dfn"} is
    set, then return.

2.  [Initialize](#concept-event-initialize){#ref-for-concept-event-initialize①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧
    link-type="dfn"} with `type`{.variable}, `bubbles`{.variable}, and
    `cancelable`{.variable}.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨
    link-type="dfn"}'s
    [`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail③
    link-type="idl"} attribute to `detail`{.variable}.
:::

### [2.5. ]{.secno}[Constructing events]{.content}[](#constructing-events){.self-link} {#constructing-events .heading .settled level="2.5"}

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications
link-type="dfn"} may define [event constructing
steps]{#concept-event-constructor-ext .dfn .dfn-paneled dfn-type="dfn"
export=""} for all or some
[events](#concept-event){#ref-for-concept-event③⑤ link-type="dfn"}. The
algorithm is passed an [event](#concept-event){#ref-for-concept-event③⑥
link-type="dfn"} `event`{.variable} and an
[`EventInit`{.idl}](#dictdef-eventinit){#ref-for-dictdef-eventinit②
link-type="idl"} `eventInitDict`{.variable} as indicated in the [inner
event creation
steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps
link-type="dfn"}.

This construct can be used by [`Event`{.idl}](#event){#ref-for-event⑧
link-type="idl"} subclasses that have a more complex structure than a
simple 1:1 mapping between their initializing dictionary members and IDL
attributes.

::: {.algorithm algorithm="constructor" algorithm-for="Event"}
When a [constructor]{#concept-event-constructor .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} of the
[`Event`{.idl}](#event){#ref-for-event⑨ link-type="idl"} interface, or
of an interface that inherits from the
[`Event`{.idl}](#event){#ref-for-event①⓪ link-type="idl"} interface, is
invoked, these steps must be run, given the arguments `type`{.variable}
and `eventInitDict`{.variable}:

1.  Let `event`{.variable} be the result of running the [inner event
    creation
    steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps①
    link-type="dfn"} with this interface, null, now, and
    `eventInitDict`{.variable}.

2.  Initialize `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑥
    link-type="idl"} attribute to `type`{.variable}.

3.  Return `event`{.variable}.
:::

::: {.algorithm algorithm="creating an event"}
To [create an event]{#concept-event-create .dfn .dfn-paneled
dfn-type="dfn" export="" lt="creating an event|create an event"} using
`eventInterface`{.variable}, which must be either
[`Event`{.idl}](#event){#ref-for-event①① link-type="idl"} or an
interface that inherits from it, and optionally given a
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm link-type="dfn"}
`realm`{.variable}, run these steps:

1.  If `realm`{.variable} is not given, then set it to null.

2.  Let `dictionary`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#dfn-convert-ecmascript-to-idl-value){#ref-for-dfn-convert-ecmascript-to-idl-value
    link-type="dfn"} the JavaScript value undefined to the dictionary
    type accepted by `eventInterface`{.variable}'s constructor. (This
    dictionary type will either be
    [`EventInit`{.idl}](#dictdef-eventinit){#ref-for-dictdef-eventinit③
    link-type="idl"} or a dictionary that inherits from it.)

    This does not work if members are required; see
    [whatwg/dom#600](https://github.com/whatwg/dom/issues/600).

3.  Let `event`{.variable} be the result of running the [inner event
    creation
    steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps②
    link-type="dfn"} with `eventInterface`{.variable},
    `realm`{.variable}, the time of the occurrence that the event is
    signaling, and `dictionary`{.variable}.

    [](#example-timestamp-initialization){.self-link}In macOS the time
    of the occurrence for input actions is available via the `timestamp`
    property of `NSEvent` objects.

4.  Initialize `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑤
    link-type="idl"} attribute to true.

5.  Return `event`{.variable}.
:::

[Create an event](#concept-event-create){#ref-for-concept-event-create
link-type="dfn"} is meant to be used by other specifications which need
to separately
[create](#concept-event-create){#ref-for-concept-event-create①
link-type="dfn"} and
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑧
link-type="dfn"} events, instead of simply
[firing](#concept-event-fire){#ref-for-concept-event-fire
link-type="dfn"} them. It ensures the event's attributes are initialized
to the correct defaults.

::: {.algorithm algorithm="inner event creation steps"}
The [inner event creation steps]{#inner-event-creation-steps .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an
`eventInterface`{.variable}, `realm`{.variable}, `time`{.variable}, and
`dictionary`{.variable}, are as follows:

1.  Let `event`{.variable} be the result of creating a new object using
    `eventInterface`{.variable}. If `realm`{.variable} is non-null, then
    use that realm; otherwise, use the default behavior defined in Web
    IDL.

    As of the time of this writing Web IDL does not yet define any
    default behavior; see
    [whatwg/webidl#135](https://github.com/whatwg/webidl/issues/135).

2.  Set `event`{.variable}'s [initialized
    flag](#initialized-flag){#ref-for-initialized-flag①
    link-type="dfn"}.

3.  Initialize `event`{.variable}'s
    [`timeStamp`{.idl}](#dom-event-timestamp){#ref-for-dom-event-timestamp②
    link-type="idl"} attribute to the [relative high resolution coarse
    time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-coarse-time){#ref-for-dfn-relative-high-resolution-coarse-time
    link-type="dfn"} given `time`{.variable} and `event`{.variable}'s
    [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global
    link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `member`{.variable} → `value`{.variable} of
    `dictionary`{.variable}, if `event`{.variable} has an attribute
    whose
    [identifier](https://webidl.spec.whatwg.org/#dfn-identifier){#ref-for-dfn-identifier
    link-type="dfn"} is `member`{.variable}, then initialize that
    attribute to `value`{.variable}.

5.  Run the [event constructing
    steps](#concept-event-constructor-ext){#ref-for-concept-event-constructor-ext
    link-type="dfn"} with `event`{.variable} and
    `dictionary`{.variable}.

6.  Return `event`{.variable}.
:::

### [2.6. ]{.secno}[Defining event interfaces]{.content}[](#defining-event-interfaces){.self-link} {#defining-event-interfaces .heading .settled level="2.6"}

In general, when defining a new interface that inherits from
[`Event`{.idl}](#event){#ref-for-event①② link-type="idl"} please always
ask feedback from the [WHATWG](https://whatwg.org/) or the [W3C WebApps
WG](https://www.w3.org/2008/webapps/) community.

The [`CustomEvent`{.idl}](#customevent){#ref-for-customevent②
link-type="idl"} interface can be used as starting point. However, do
not introduce any `init``*`{.variable}`Event()` methods as they are
redundant with constructors. Interfaces that inherit from the
[`Event`{.idl}](#event){#ref-for-event①③ link-type="idl"} interface that
have such a method only have it for historical reasons.

### [2.7. ]{.secno}[Interface [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑧ link-type="idl"}]{.content}[](#interface-eventtarget){.self-link} {#interface-eventtarget .heading .settled level="2.7"}

``` {.idl .highlight .def}
[Exposed=*]
interface EventTarget {
  constructor();

  undefined addEventListener(DOMString type, EventListener? callback, optional (AddEventListenerOptions or boolean) options = {});
  undefined removeEventListener(DOMString type, EventListener? callback, optional (EventListenerOptions or boolean) options = {});
  boolean dispatchEvent(Event event);
};

callback interface EventListener {
  undefined handleEvent(Event event);
};

dictionary EventListenerOptions {
  boolean capture = false;
};

dictionary AddEventListenerOptions : EventListenerOptions {
  boolean passive;
  boolean once = false;
  AbortSignal signal;
};
```

An [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑨
link-type="idl"} object represents a target to which an
[event](#concept-event){#ref-for-concept-event③⑦ link-type="dfn"} can be
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑨
link-type="dfn"} when something has occurred.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⓪
link-type="idl"} object has an associated [event listener
list]{#eventtarget-event-listener-list .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" noexport=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list④
link-type="dfn"} of zero or more [event
listeners](#concept-event-listener){#ref-for-concept-event-listener①①
link-type="dfn"}). It is initially the empty list.

An [event listener]{#concept-event-listener .dfn .dfn-paneled
dfn-type="dfn" export=""} can be used to observe a specific
[event](#concept-event){#ref-for-concept-event③⑧ link-type="dfn"} and
consists of:

- [type]{#event-listener-type .dfn .dfn-paneled dfn-for="event listener"
  dfn-type="dfn" noexport=""} (a string)
- [callback]{#event-listener-callback .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" export=""} (null or an
  [`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener②
  link-type="idl"} object)
- [capture]{#event-listener-capture .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (a boolean,
  initially false)
- [passive]{#event-listener-passive .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (null or a
  boolean, initially null)
- [once]{#event-listener-once .dfn .dfn-paneled dfn-for="event listener"
  dfn-type="dfn" noexport=""} (a boolean, initially false)
- [signal]{#event-listener-signal .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (null or an
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②
  link-type="idl"} object)
- [removed]{#event-listener-removed .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (a boolean for
  bookkeeping purposes, initially false)

Although
[callback](#event-listener-callback){#ref-for-event-listener-callback②
link-type="dfn"} is an
[`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener③
link-type="idl"} object, an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①②
link-type="dfn"} is a broader concept as can be seen above.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①①
link-type="idl"} object also has an associated [get the
parent]{#get-the-parent .dfn .dfn-paneled dfn-type="dfn" export=""}
algorithm, which takes an
[event](#concept-event){#ref-for-concept-event③⑨ link-type="dfn"}
`event`{.variable}, and returns an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①②
link-type="idl"} object. Unless specified otherwise it returns null.

[Nodes](#concept-node){#ref-for-concept-node② link-type="dfn"}, [shadow
roots](#concept-shadow-root){#ref-for-concept-shadow-root①
link-type="dfn"}, and
[documents](#concept-document){#ref-for-concept-document①
link-type="dfn"} override the [get the
parent](#get-the-parent){#ref-for-get-the-parent link-type="dfn"}
algorithm.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①③
link-type="idl"} object can have an associated [activation
behavior]{#eventtarget-activation-behavior .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm. The
[activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior
link-type="dfn"} algorithm is passed an
[event](#concept-event){#ref-for-concept-event④⓪ link-type="dfn"}, as
indicated in the
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch②⓪
link-type="dfn"} algorithm.

This exists because user agents perform certain actions for certain
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①④
link-type="idl"} objects, e.g., the
[`area`](https://html.spec.whatwg.org/multipage/image-maps.html#the-area-element){#ref-for-the-area-element
link-type="element"} element, in response to synthetic
[`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent
link-type="idl"} [events](#concept-event){#ref-for-concept-event④①
link-type="dfn"} whose
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑦
link-type="idl"} attribute is `click`. Web compatibility prevented it
from being removed and it is now the enshrined way of defining an
activation of something.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑤
link-type="idl"} object that has [activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior①
link-type="dfn"}, can additionally have both (not either) a
[legacy-pre-activation
behavior]{#eventtarget-legacy-pre-activation-behavior .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm and a
[legacy-canceled-activation
behavior]{#eventtarget-legacy-canceled-activation-behavior .dfn
.dfn-paneled dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm.

These algorithms only exist for checkbox and radio
[`input`](https://html.spec.whatwg.org/multipage/input.html#the-input-element){#ref-for-the-input-element
link-type="element"} elements and are not to be used for anything else.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

`target`{.variable}` = new `[`EventTarget`](#dom-eventtarget-eventtarget){#ref-for-dom-eventtarget-eventtarget① .idl-code link-type="constructor"}`();`

:   Creates a new
    [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑥
    link-type="idl"} object, which can be used by developers to
    [dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch②①
    link-type="dfn"} and listen for
    [events](#concept-event){#ref-for-concept-event④② link-type="dfn"}.

`target`{.variable}` . `[`addEventListener`](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener③ .idl-code link-type="method"}`(``type`{.variable}`, ``callback`{.variable}` [, ``options`{.variable}`])`

:   Appends an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①③
    link-type="dfn"} for
    [events](#concept-event){#ref-for-concept-event④③ link-type="dfn"}
    whose [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑧
    link-type="idl"} attribute value is `type`{.variable}. The
    `callback`{.variable} argument sets the
    [callback](#event-listener-callback){#ref-for-event-listener-callback③
    link-type="dfn"} that will be invoked when the
    [event](#concept-event){#ref-for-concept-event④④ link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch②②
    link-type="dfn"}.

    The `options`{.variable} argument sets listener-specific options.
    For compatibility this can be a boolean, in which case the method
    behaves exactly as if the value was specified as
    `options`{.variable}'s
    [`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture
    link-type="idl"}.

    When set to true, `options`{.variable}'s
    [`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture①
    link-type="idl"} prevents
    [callback](#event-listener-callback){#ref-for-event-listener-callback④
    link-type="dfn"} from being invoked when the
    [event](#concept-event){#ref-for-concept-event④⑤ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase④
    link-type="idl"} attribute value is
    [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase③
    link-type="idl"}. When false (or not present),
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑤
    link-type="dfn"} will not be invoked when
    [event](#concept-event){#ref-for-concept-event④⑥ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑤
    link-type="idl"} attribute value is
    [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase③
    link-type="idl"}. Either way,
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑥
    link-type="dfn"} will be invoked if
    [event](#concept-event){#ref-for-concept-event④⑦ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑥
    link-type="idl"} attribute value is
    [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target③
    link-type="idl"}.

    When set to true, `options`{.variable}'s
    [`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive①
    link-type="idl"} indicates that the
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑦
    link-type="dfn"} will not cancel the event by invoking
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑥
    link-type="idl"}. This is used to enable performance optimizations
    described in [§ 2.8 Observing event
    listeners](#observing-event-listeners).

    When set to true, `options`{.variable}'s
    [`once`{.idl}](#dom-addeventlisteneroptions-once){#ref-for-dom-addeventlisteneroptions-once
    link-type="idl"} indicates that the
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑧
    link-type="dfn"} will only be invoked once after which the event
    listener will be removed.

    If an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③
    link-type="idl"} is passed for `options`{.variable}'s
    [`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal
    link-type="idl"}, then the event listener will be removed when
    signal is aborted.

    The [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①④
    link-type="dfn"} is appended to `target`{.variable}'s [event
    listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list
    link-type="dfn"} and is not appended if it has the same
    [type](#event-listener-type){#ref-for-event-listener-type
    link-type="dfn"},
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑨
    link-type="dfn"}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture②
    link-type="dfn"}.

`target`{.variable}` . `[`removeEventListener`](#dom-eventtarget-removeeventlistener){#ref-for-dom-eventtarget-removeeventlistener② .idl-code link-type="method"}`(``type`{.variable}`, ``callback`{.variable}` [, ``options`{.variable}`])`

:   Removes the [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑤
    link-type="dfn"} in `target`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list①
    link-type="dfn"} with the same `type`{.variable},
    `callback`{.variable}, and `options`{.variable}.

`target`{.variable}` . `[`dispatchEvent`](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent③ .idl-code link-type="method"}`(``event`{.variable}`)`

:   [Dispatches](#concept-event-dispatch){#ref-for-concept-event-dispatch②③
    link-type="dfn"} a synthetic event `event`{.variable} to
    `target`{.variable} and returns true if either `event`{.variable}'s
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑥
    link-type="idl"} attribute value is false or its
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑦
    link-type="idl"} method was not invoked; otherwise false.

::: {.algorithm algorithm="flatten" algorithm-for="Event"}
To [flatten]{#concept-flatten-options .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} `options`{.variable}, run these steps:

1.  If `options`{.variable} is a boolean, then return
    `options`{.variable}.

2.  Return
    `options`{.variable}\[\"[`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture②
    link-type="idl"}\"\].
:::

::: {.algorithm algorithm="flatten more" algorithm-for="Event"}
To [flatten more]{#event-flatten-more .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} `options`{.variable}, run these steps:

1.  Let `capture`{.variable} be the result of
    [flattening](#concept-flatten-options){#ref-for-concept-flatten-options
    link-type="dfn"} `options`{.variable}.

2.  Let `once`{.variable} be false.

3.  Let `passive`{.variable} and `signal`{.variable} be null.

4.  If `options`{.variable} is a
    [dictionary](https://webidl.spec.whatwg.org/#dfn-dictionary){#ref-for-dfn-dictionary
    link-type="dfn"}:

    1.  Set `once`{.variable} to
        `options`{.variable}\[\"[`once`{.idl}](#dom-addeventlisteneroptions-once){#ref-for-dom-addeventlisteneroptions-once①
        link-type="idl"}\"\].

    2.  If
        `options`{.variable}\[\"[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive②
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists
        link-type="dfn"}, then set `passive`{.variable} to
        `options`{.variable}\[\"[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive③
        link-type="idl"}\"\].

    3.  If
        `options`{.variable}\[\"[`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal①
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①
        link-type="dfn"}, then set `signal`{.variable} to
        `options`{.variable}\[\"[`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal②
        link-type="idl"}\"\].

5.  Return `capture`{.variable}, `passive`{.variable},
    `once`{.variable}, and `signal`{.variable}.
:::

The [`new EventTarget()`]{#dom-eventtarget-eventtarget .dfn .dfn-paneled
.idl-code dfn-for="EventTarget" dfn-type="constructor" export=""
lt="EventTarget()|constructor()"} constructor steps are to do nothing.

Because of the defaults stated elsewhere, the returned
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑦
link-type="idl"}'s [get the
parent](#get-the-parent){#ref-for-get-the-parent① link-type="dfn"}
algorithm will return null, and it will have no [activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior②
link-type="dfn"}, [legacy-pre-activation
behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior
link-type="dfn"}, or [legacy-canceled-activation
behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior
link-type="dfn"}.

In the future we could allow custom [get the
parent](#get-the-parent){#ref-for-get-the-parent② link-type="dfn"}
algorithms. Let us know if this would be useful for your programs. For
now, all author-created
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑧
link-type="idl"}s do not participate in a tree structure.

::: {.algorithm algorithm="default passive value"}
The [default passive value]{#default-passive-value .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an event type `type`{.variable} and
an [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑨
link-type="idl"} `eventTarget`{.variable}, is determined as follows:

1.  Return true if all of the following are true:

    - `type`{.variable} is one of \"`touchstart`\", \"`touchmove`\",
      \"`wheel`\", or \"`mousewheel`\".
      [\[TOUCH-EVENTS\]](#biblio-touch-events "Touch Events"){link-type="biblio"}
      [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    - `eventTarget`{.variable} is a
      [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window③
      link-type="idl"} object, or is a
      [node](#concept-node){#ref-for-concept-node③ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document
      link-type="dfn"} is `eventTarget`{.variable}, or is a
      [node](#concept-node){#ref-for-concept-node④ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document①
      link-type="dfn"}'s [document
      element](#document-element){#ref-for-document-element
      link-type="dfn"} is `eventTarget`{.variable}, or is a
      [node](#concept-node){#ref-for-concept-node⑤ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document②
      link-type="dfn"}'s [body
      element](https://html.spec.whatwg.org/multipage/dom.html#the-body-element-2){#ref-for-the-body-element-2
      link-type="dfn"} is `eventTarget`{.variable}.
      [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

2.  Return false.
:::

::: {.algorithm algorithm="add an event listener"}
To [add an event listener]{#add-an-event-listener .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②⓪
link-type="idl"} object `eventTarget`{.variable} and an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①⑥
link-type="dfn"} `listener`{.variable}, run these steps:

1.  If `eventTarget`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope
    link-type="idl"} object, its [service
    worker](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope-service-worker){#ref-for-serviceworkerglobalscope-service-worker
    link-type="dfn"}'s [script
    resource](https://w3c.github.io/ServiceWorker/#dfn-script-resource){#ref-for-dfn-script-resource
    link-type="dfn"}'s [has ever been evaluated
    flag](https://w3c.github.io/ServiceWorker/#dfn-has-ever-been-evaluated-flag){#ref-for-dfn-has-ever-been-evaluated-flag
    link-type="dfn"} is set, and `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type①
    link-type="dfn"} matches the
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑨
    link-type="idl"} attribute value of any of the [service worker
    events](https://w3c.github.io/ServiceWorker/#dfn-service-worker-events){#ref-for-dfn-service-worker-events
    link-type="dfn"}, then [report a warning to the
    console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#ref-for-report-a-warning-to-the-console
    link-type="dfn"} that this might not give the expected results.
    [\[SERVICE-WORKERS\]](#biblio-service-workers "Service Workers"){link-type="biblio"}

2.  If `listener`{.variable}'s
    [signal](#event-listener-signal){#ref-for-event-listener-signal
    link-type="dfn"} is not null and is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted
    link-type="dfn"}, then return.

3.  If `listener`{.variable}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①⓪
    link-type="dfn"} is null, then return.

4.  If `listener`{.variable}'s
    [passive](#event-listener-passive){#ref-for-event-listener-passive
    link-type="dfn"} is null, then set it to the [default passive
    value](#default-passive-value){#ref-for-default-passive-value
    link-type="dfn"} given `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type②
    link-type="dfn"} and `eventTarget`{.variable}.

5.  If `eventTarget`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list②
    link-type="dfn"} does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain
    link-type="dfn"} an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑦
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type③
    link-type="dfn"} is `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type④
    link-type="dfn"},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①①
    link-type="dfn"} is `listener`{.variable}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①②
    link-type="dfn"}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture③
    link-type="dfn"} is `listener`{.variable}'s
    [capture](#event-listener-capture){#ref-for-event-listener-capture④
    link-type="dfn"}, then
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append②
    link-type="dfn"} `listener`{.variable} to `eventTarget`{.variable}'s
    [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list③
    link-type="dfn"}.

6.  If `listener`{.variable}'s
    [signal](#event-listener-signal){#ref-for-event-listener-signal①
    link-type="dfn"} is not null, then [add the
    following](#abortsignal-add){#ref-for-abortsignal-add
    link-type="dfn"} abort steps to it:

    1.  [Remove an event
        listener](#remove-an-event-listener){#ref-for-remove-an-event-listener
        link-type="dfn"} with `eventTarget`{.variable} and
        `listener`{.variable}.

The [add an event
listener](#add-an-event-listener){#ref-for-add-an-event-listener
link-type="dfn"} concept exists to ensure [event
handlers](https://html.spec.whatwg.org/multipage/webappapis.html#event-handlers){#ref-for-event-handlers
link-type="dfn"} use the same code path.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="addEventListener(type, callback, options)" algorithm-for="EventTarget"}
The
[`addEventListener(``type`{.variable}`, ``callback`{.variable}`, ``options`{.variable}`)`]{#dom-eventtarget-addeventlistener
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""
lt="addEventListener(type, callback, options)|addEventListener(type, callback)"}
method steps are:

1.  Let `capture`{.variable}, `passive`{.variable}, `once`{.variable},
    and `signal`{.variable} be the result of [flattening
    more](#event-flatten-more){#ref-for-event-flatten-more
    link-type="dfn"} `options`{.variable}.

2.  [Add an event
    listener](#add-an-event-listener){#ref-for-add-an-event-listener①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪
    link-type="dfn"} and an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑧
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type⑤
    link-type="dfn"} is `type`{.variable},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①③
    link-type="dfn"} is `callback`{.variable},
    [capture](#event-listener-capture){#ref-for-event-listener-capture⑤
    link-type="dfn"} is `capture`{.variable},
    [passive](#event-listener-passive){#ref-for-event-listener-passive①
    link-type="dfn"} is `passive`{.variable},
    [once](#event-listener-once){#ref-for-event-listener-once
    link-type="dfn"} is `once`{.variable}, and
    [signal](#event-listener-signal){#ref-for-event-listener-signal②
    link-type="dfn"} is `signal`{.variable}.
:::

::: {.algorithm algorithm="remove an event listener"}
To [remove an event listener]{#remove-an-event-listener .dfn
.dfn-paneled dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②①
link-type="idl"} object `eventTarget`{.variable} and an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①⑨
link-type="dfn"} `listener`{.variable}, run these steps:

1.  If `eventTarget`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope①
    link-type="idl"} object and its [service
    worker](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope-service-worker){#ref-for-serviceworkerglobalscope-service-worker①
    link-type="dfn"}'s [set of event types to
    handle](https://w3c.github.io/ServiceWorker/#dfn-set-of-event-types-to-handle){#ref-for-dfn-set-of-event-types-to-handle
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain①
    link-type="dfn"} `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type⑥
    link-type="dfn"}, then [report a warning to the
    console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#ref-for-report-a-warning-to-the-console①
    link-type="dfn"} that this might not give the expected results.
    [\[SERVICE-WORKERS\]](#biblio-service-workers "Service Workers"){link-type="biblio"}

2.  Set `listener`{.variable}'s
    [removed](#event-listener-removed){#ref-for-event-listener-removed
    link-type="dfn"} to true and
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove
    link-type="dfn"} `listener`{.variable} from
    `eventTarget`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list④
    link-type="dfn"}.

HTML needs this to define event handlers.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="remove all event listeners"}
To [remove all event listeners]{#remove-all-event-listeners .dfn
.dfn-paneled dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②②
link-type="idl"} object `eventTarget`{.variable}, [for
each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①
link-type="dfn"} `listener`{.variable} of `eventTarget`{.variable}'s
[event listener
list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑤
link-type="dfn"}, [remove an event
listener](#remove-an-event-listener){#ref-for-remove-an-event-listener①
link-type="dfn"} with `eventTarget`{.variable} and
`listener`{.variable}.

HTML needs this to define `document.open()`.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="removeEventListener(type, callback, options)" algorithm-for="EventTarget"}
The
[`removeEventListener(``type`{.variable}`, ``callback`{.variable}`, ``options`{.variable}`)`]{#dom-eventtarget-removeeventlistener
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""
lt="removeEventListener(type, callback, options)|removeEventListener(type, callback)"}
method steps are:

1.  Let `capture`{.variable} be the result of
    [flattening](#concept-flatten-options){#ref-for-concept-flatten-options①
    link-type="dfn"} `options`{.variable}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①
    link-type="dfn"}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑥
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain②
    link-type="dfn"} an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener②⓪
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type⑦
    link-type="dfn"} is `type`{.variable},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①④
    link-type="dfn"} is `callback`{.variable}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture⑥
    link-type="dfn"} is `capture`{.variable}, then [remove an event
    listener](#remove-an-event-listener){#ref-for-remove-an-event-listener②
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②
    link-type="dfn"} and that [event
    listener](#concept-event-listener){#ref-for-concept-event-listener②①
    link-type="dfn"}.

The event listener list will not contain multiple event listeners with
equal `type`{.variable}, `callback`{.variable}, and
`capture`{.variable}, as [add an event
listener](#add-an-event-listener){#ref-for-add-an-event-listener②
link-type="dfn"} prevents that.
:::

::: {.algorithm algorithm="dispatchEvent(event)" algorithm-for="EventTarget"}
The
[`dispatchEvent(``event`{.variable}`)`]{#dom-eventtarget-dispatchevent
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""} method steps are:

1.  If `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag② link-type="dfn"} is
    set, or if its [initialized
    flag](#initialized-flag){#ref-for-initialized-flag② link-type="dfn"}
    is not set, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥
    link-type="dfn"} an
    \"[`InvalidStateError`{.idl}](https://webidl.spec.whatwg.org/#invalidstateerror){#ref-for-invalidstateerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥
    link-type="idl"}.

2.  Initialize `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑥
    link-type="idl"} attribute to false.

3.  Return the result of
    [dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②④
    link-type="dfn"} `event`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③
    link-type="dfn"}.
:::

### [2.8. ]{.secno}[Observing event listeners]{.content}[](#observing-event-listeners){.self-link} {#observing-event-listeners .heading .settled level="2.8"}

In general, developers do not expect the presence of an [event
listener](#concept-event-listener){#ref-for-concept-event-listener②②
link-type="dfn"} to be observable. The impact of an [event
listener](#concept-event-listener){#ref-for-concept-event-listener②③
link-type="dfn"} is determined by its
[callback](#event-listener-callback){#ref-for-event-listener-callback①⑤
link-type="dfn"}. That is, a developer adding a no-op [event
listener](#concept-event-listener){#ref-for-concept-event-listener②④
link-type="dfn"} would not expect it to have any side effects.

Unfortunately, some event APIs have been designed such that implementing
them efficiently requires observing [event
listeners](#concept-event-listener){#ref-for-concept-event-listener②⑤
link-type="dfn"}. This can make the presence of listeners observable in
that even empty listeners can have a dramatic performance impact on the
behavior of the application. For example, touch and wheel events which
can be used to block asynchronous scrolling. In some cases this problem
can be mitigated by specifying the event to be
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑦
link-type="idl"} only when there is at least one
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive④
link-type="idl"} listener. For example,
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑤
link-type="idl"}
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent①
link-type="idl"} listeners must block scrolling, but if all listeners
are
[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑥
link-type="idl"} then scrolling can be allowed to start [in
parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel
link-type="dfn"} by making the
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent②
link-type="idl"} uncancelable (so that calls to
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑧
link-type="idl"} are ignored). So code dispatching an event is able to
observe the absence of
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑦
link-type="idl"} listeners, and use that to clear the
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑧
link-type="idl"} property of the event being dispatched.

Ideally, any new event APIs are defined such that they do not need this
property. (Use [whatwg/dom](https://github.com/whatwg/dom/issues) for
discussion.)

::: {.algorithm algorithm="legacy-obtain service worker fetch event listener callbacks"}
To [legacy-obtain service worker fetch event listener
callbacks]{#legacy-obtain-service-worker-fetch-event-listener-callbacks
.dfn .dfn-paneled dfn-type="dfn" export=""} given a
[`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope②
link-type="idl"} `global`{.variable}, run these steps. They return a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list⑤
link-type="dfn"} of
[`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener④
link-type="idl"} objects.

1.  Let `callbacks`{.variable} be « ».

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②
    link-type="dfn"} `listener`{.variable} of `global`{.variable}'s
    [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑦
    link-type="dfn"}:

    1.  If `listener`{.variable}'s
        [type](#event-listener-type){#ref-for-event-listener-type⑧
        link-type="dfn"} is \"`fetch`\", and `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑥
        link-type="dfn"} is not null, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append③
        link-type="dfn"} `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑦
        link-type="dfn"} to `callbacks`{.variable}.

3.  Return `callbacks`{.variable}.
:::

### [2.9. ]{.secno}[Dispatching events]{.content}[](#dispatching-events){.self-link} {#dispatching-events .heading .settled level="2.9"}

::: {.algorithm algorithm="dispatch"}
To [dispatch]{#concept-event-dispatch .dfn .dfn-paneled dfn-type="dfn"
export=""} an `event`{.variable} to a `target`{.variable}, with an
optional `legacy target override flag`{.variable} and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Set `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag③ link-type="dfn"}.

2.  Let `targetOverride`{.variable} be `target`{.variable}, if
    `legacy target override flag`{.variable} is not given, and
    `target`{.variable}'s [associated
    `Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window
    link-type="dfn"} otherwise.
    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    `legacy target override flag`{.variable} is only used by HTML and
    only when `target`{.variable} is a
    [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window④
    link-type="idl"} object.

3.  Let `activationTarget`{.variable} be null.

4.  Let `relatedTarget`{.variable} be the result of
    [retargeting](#retarget){#ref-for-retarget link-type="dfn"}
    `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget①
    link-type="dfn"} against `target`{.variable}.

5.  Let `clearTargets`{.variable} be false.

6.  If `target`{.variable} is not `relatedTarget`{.variable} or
    `target`{.variable} is `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget②
    link-type="dfn"}:

    1.  Let `touchTargets`{.variable} be a new
        [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑥
        link-type="dfn"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate③
        link-type="dfn"} `touchTarget`{.variable} of
        `event`{.variable}'s [touch target
        list](#event-touch-target-list){#ref-for-event-touch-target-list①
        link-type="dfn"},
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append④
        link-type="dfn"} the result of
        [retargeting](#retarget){#ref-for-retarget① link-type="dfn"}
        `touchTarget`{.variable} against `target`{.variable} to
        `touchTargets`{.variable}.

    3.  [Append to an event
        path](#concept-event-path-append){#ref-for-concept-event-path-append
        link-type="dfn"} with `event`{.variable}, `target`{.variable},
        `targetOverride`{.variable}, `relatedTarget`{.variable},
        `touchTargets`{.variable}, and false.

    4.  Let `isActivationEvent`{.variable} be true, if
        `event`{.variable} is a
        [`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent①
        link-type="idl"} object and `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⓪
        link-type="idl"} attribute is \"`click`\"; otherwise false.

    5.  If `isActivationEvent`{.variable} is true and
        `target`{.variable} has [activation
        behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior③
        link-type="dfn"}, then set `activationTarget`{.variable} to
        `target`{.variable}.

    6.  Let `slottable`{.variable} be `target`{.variable}, if
        `target`{.variable} is a
        [slottable](#concept-slotable){#ref-for-concept-slotable
        link-type="dfn"} and is
        [assigned](#slotable-assigned){#ref-for-slotable-assigned
        link-type="dfn"}, and null otherwise.

    7.  Let `slot-in-closed-tree`{.variable} be false.

    8.  Let `parent`{.variable} be the result of invoking
        `target`{.variable}'s [get the
        parent](#get-the-parent){#ref-for-get-the-parent③
        link-type="dfn"} with `event`{.variable}.

    9.  While `parent`{.variable} is non-null:

        1.  If `slottable`{.variable} is non-null:

            1.  Assert: `parent`{.variable} is a
                [slot](#concept-slot){#ref-for-concept-slot
                link-type="dfn"}.

            2.  Set `slottable`{.variable} to null.

            3.  If `parent`{.variable}'s
                [root](#concept-tree-root){#ref-for-concept-tree-root④
                link-type="dfn"} is a [shadow
                root](#concept-shadow-root){#ref-for-concept-shadow-root②
                link-type="dfn"} whose
                [mode](#shadowroot-mode){#ref-for-shadowroot-mode①
                link-type="dfn"} is \"`closed`\", then set
                `slot-in-closed-tree`{.variable} to true.

        2.  If `parent`{.variable} is a
            [slottable](#concept-slotable){#ref-for-concept-slotable①
            link-type="dfn"} and is
            [assigned](#slotable-assigned){#ref-for-slotable-assigned①
            link-type="dfn"}, then set `slottable`{.variable} to
            `parent`{.variable}.

        3.  Let `relatedTarget`{.variable} be the result of
            [retargeting](#retarget){#ref-for-retarget② link-type="dfn"}
            `event`{.variable}'s
            [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget③
            link-type="dfn"} against `parent`{.variable}.

        4.  Let `touchTargets`{.variable} be a new
            [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑦
            link-type="dfn"}.

        5.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate④
            link-type="dfn"} `touchTarget`{.variable} of
            `event`{.variable}'s [touch target
            list](#event-touch-target-list){#ref-for-event-touch-target-list②
            link-type="dfn"},
            [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑤
            link-type="dfn"} the result of
            [retargeting](#retarget){#ref-for-retarget③ link-type="dfn"}
            `touchTarget`{.variable} against `parent`{.variable} to
            `touchTargets`{.variable}.

        6.  If `parent`{.variable} is a
            [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑤
            link-type="idl"} object, or `parent`{.variable} is a
            [node](#concept-node){#ref-for-concept-node⑥
            link-type="dfn"} and `target`{.variable}'s
            [root](#concept-tree-root){#ref-for-concept-tree-root⑤
            link-type="dfn"} is a [shadow-including inclusive
            ancestor](#concept-shadow-including-inclusive-ancestor){#ref-for-concept-shadow-including-inclusive-ancestor
            link-type="dfn"} of `parent`{.variable}:

            1.  If `isActivationEvent`{.variable} is true,
                `event`{.variable}'s
                [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑤
                link-type="idl"} attribute is true,
                `activationTarget`{.variable} is null, and
                `parent`{.variable} has [activation
                behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior④
                link-type="dfn"}, then set `activationTarget`{.variable}
                to `parent`{.variable}.

            2.  [Append to an event
                path](#concept-event-path-append){#ref-for-concept-event-path-append①
                link-type="dfn"} with `event`{.variable},
                `parent`{.variable}, null, `relatedTarget`{.variable},
                `touchTargets`{.variable}, and
                `slot-in-closed-tree`{.variable}.

        7.  Otherwise, if `parent`{.variable} is
            `relatedTarget`{.variable}, then set `parent`{.variable} to
            null.

        8.  Otherwise:

            1.  Set `target`{.variable} to `parent`{.variable}.

            2.  If `isActivationEvent`{.variable} is true,
                `activationTarget`{.variable} is null, and
                `target`{.variable} has [activation
                behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior⑤
                link-type="dfn"}, then set `activationTarget`{.variable}
                to `target`{.variable}.

            3.  [Append to an event
                path](#concept-event-path-append){#ref-for-concept-event-path-append②
                link-type="dfn"} with `event`{.variable},
                `parent`{.variable}, `target`{.variable},
                `relatedTarget`{.variable}, `touchTargets`{.variable},
                and `slot-in-closed-tree`{.variable}.

        9.  If `parent`{.variable} is non-null, then set
            `parent`{.variable} to the result of invoking
            `parent`{.variable}'s [get the
            parent](#get-the-parent){#ref-for-get-the-parent④
            link-type="dfn"} with `event`{.variable}.

        10. Set `slot-in-closed-tree`{.variable} to false.

    10. Let `clearTargetsStruct`{.variable} be the last struct in
        `event`{.variable}'s [path](#event-path){#ref-for-event-path④
        link-type="dfn"} whose [shadow-adjusted
        target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target
        link-type="dfn"} is non-null.

    11. If `clearTargetsStruct`{.variable}'s [shadow-adjusted
        target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target①
        link-type="dfn"}, `clearTargetsStruct`{.variable}'s
        [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget
        link-type="dfn"}, or an
        [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②③
        link-type="idl"} object in `clearTargetsStruct`{.variable}'s
        [touch target
        list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list
        link-type="dfn"} is a
        [node](#concept-node){#ref-for-concept-node⑦ link-type="dfn"}
        whose [root](#concept-tree-root){#ref-for-concept-tree-root⑥
        link-type="dfn"} is a [shadow
        root](#concept-shadow-root){#ref-for-concept-shadow-root③
        link-type="dfn"}: set `clearTargets`{.variable} to true.

    12. If `activationTarget`{.variable} is non-null and
        `activationTarget`{.variable} has [legacy-pre-activation
        behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior①
        link-type="dfn"}, then run `activationTarget`{.variable}'s
        [legacy-pre-activation
        behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior②
        link-type="dfn"}.

    13. [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑤
        link-type="dfn"} `struct`{.variable} of `event`{.variable}'s
        [path](#event-path){#ref-for-event-path⑤ link-type="dfn"}, in
        reverse order:

        1.  If `struct`{.variable}'s [shadow-adjusted
            target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target②
            link-type="dfn"} is non-null, then set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑦
            link-type="idl"} attribute to
            [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target④
            link-type="idl"}.

        2.  Otherwise, set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑧
            link-type="idl"} attribute to
            [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase④
            link-type="idl"}.

        3.  [Invoke](#concept-event-listener-invoke){#ref-for-concept-event-listener-invoke
            link-type="dfn"} with `struct`{.variable},
            `event`{.variable}, \"`capturing`\", and
            `legacyOutputDidListenersThrowFlag`{.variable} if given.

    14. [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑥
        link-type="dfn"} `struct`{.variable} of `event`{.variable}'s
        [path](#event-path){#ref-for-event-path⑥ link-type="dfn"}:

        1.  If `struct`{.variable}'s [shadow-adjusted
            target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target③
            link-type="dfn"} is non-null, then set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑨
            link-type="idl"} attribute to
            [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target⑤
            link-type="idl"}.

        2.  Otherwise:

            1.  If `event`{.variable}'s
                [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑥
                link-type="idl"} attribute is false, then
                [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue
                link-type="dfn"}.

            2.  Set `event`{.variable}'s
                [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①⓪
                link-type="idl"} attribute to
                [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase④
                link-type="idl"}.

        3.  [Invoke](#concept-event-listener-invoke){#ref-for-concept-event-listener-invoke①
            link-type="dfn"} with `struct`{.variable},
            `event`{.variable}, \"`bubbling`\", and
            `legacyOutputDidListenersThrowFlag`{.variable} if given.

7.  Set `event`{.variable}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①①
    link-type="idl"} attribute to
    [`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none③
    link-type="idl"}.

8.  Set `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑤
    link-type="idl"} attribute to null.

9.  Set `event`{.variable}'s [path](#event-path){#ref-for-event-path⑦
    link-type="dfn"} to the empty list.

10. Unset `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag④ link-type="dfn"},
    [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag⑤
    link-type="dfn"}, and [stop immediate propagation
    flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag②
    link-type="dfn"}.

11. If `clearTargets`{.variable} is true:

    1.  Set `event`{.variable}'s
        [target](#event-target){#ref-for-event-target⑨ link-type="dfn"}
        to null.

    2.  Set `event`{.variable}'s
        [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget④
        link-type="dfn"} to null.

    3.  Set `event`{.variable}'s [touch target
        list](#event-touch-target-list){#ref-for-event-touch-target-list③
        link-type="dfn"} to the empty list.

12. If `activationTarget`{.variable} is non-null:

    1.  If `event`{.variable}'s [canceled
        flag](#canceled-flag){#ref-for-canceled-flag④ link-type="dfn"}
        is unset, then run `activationTarget`{.variable}'s [activation
        behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior⑥
        link-type="dfn"} with `event`{.variable}.

    2.  Otherwise, if `activationTarget`{.variable} has
        [legacy-canceled-activation
        behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior①
        link-type="dfn"}, then run `activationTarget`{.variable}'s
        [legacy-canceled-activation
        behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior②
        link-type="dfn"}.

13. Return false if `event`{.variable}'s [canceled
    flag](#canceled-flag){#ref-for-canceled-flag⑤ link-type="dfn"} is
    set; otherwise true.
:::

::: {.algorithm algorithm="append to an event path"}
To [append to an event path]{#concept-event-path-append .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an `event`{.variable},
`invocationTarget`{.variable}, `shadowAdjustedTarget`{.variable},
`relatedTarget`{.variable}, `touchTargets`{.variable}, and a
`slot-in-closed-tree`{.variable}, run these steps:

1.  Let `invocationTargetInShadowTree`{.variable} be false.

2.  If `invocationTarget`{.variable} is a
    [node](#concept-node){#ref-for-concept-node⑧ link-type="dfn"} and
    its [root](#concept-tree-root){#ref-for-concept-tree-root⑦
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root④
    link-type="dfn"}, then set `invocationTargetInShadowTree`{.variable}
    to true.

3.  Let `root-of-closed-tree`{.variable} be false.

4.  If `invocationTarget`{.variable} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root⑤
    link-type="dfn"} whose
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode② link-type="dfn"}
    is \"`closed`\", then set `root-of-closed-tree`{.variable} to true.

5.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑥
    link-type="dfn"} a new
    [struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct②
    link-type="dfn"} to `event`{.variable}'s
    [path](#event-path){#ref-for-event-path⑧ link-type="dfn"} whose
    [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target④
    link-type="dfn"} is `invocationTarget`{.variable},
    [invocation-target-in-shadow-tree](#event-path-invocation-target-in-shadow-tree){#ref-for-event-path-invocation-target-in-shadow-tree
    link-type="dfn"} is `invocationTargetInShadowTree`{.variable},
    [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target④
    link-type="dfn"} is `shadowAdjustedTarget`{.variable},
    [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget①
    link-type="dfn"} is `relatedTarget`{.variable}, [touch target
    list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list①
    link-type="dfn"} is `touchTargets`{.variable},
    [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree③
    link-type="dfn"} is `root-of-closed-tree`{.variable}, and
    [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree③
    link-type="dfn"} is `slot-in-closed-tree`{.variable}.
:::

::: {.algorithm algorithm="invoke"}
To [invoke]{#concept-event-listener-invoke .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a `struct`{.variable},
`event`{.variable}, `phase`{.variable}, and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Set `event`{.variable}'s
    [target](#event-target){#ref-for-event-target①⓪ link-type="dfn"} to
    the [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target⑤
    link-type="dfn"} of the last struct in `event`{.variable}'s
    [path](#event-path){#ref-for-event-path⑨ link-type="dfn"}, that is
    either `struct`{.variable} or preceding `struct`{.variable}, whose
    [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target⑥
    link-type="dfn"} is non-null.

2.  Set `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget⑤
    link-type="dfn"} to `struct`{.variable}'s
    [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget②
    link-type="dfn"}.

3.  Set `event`{.variable}'s [touch target
    list](#event-touch-target-list){#ref-for-event-touch-target-list④
    link-type="dfn"} to `struct`{.variable}'s [touch target
    list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list②
    link-type="dfn"}.

4.  If `event`{.variable}'s [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag⑥
    link-type="dfn"} is set, then return.

5.  Initialize `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑥
    link-type="idl"} attribute to `struct`{.variable}'s [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target⑤
    link-type="dfn"}.

6.  Let `listeners`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone
    link-type="dfn"} of `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑦
    link-type="idl"} attribute value's [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑧
    link-type="dfn"}.

    This avoids [event
    listeners](#concept-event-listener){#ref-for-concept-event-listener②⑥
    link-type="dfn"} added after this point from being run. Note that
    removal still has an effect due to the
    [removed](#event-listener-removed){#ref-for-event-listener-removed①
    link-type="dfn"} field.

7.  Let `invocationTargetInShadowTree`{.variable} be
    `struct`{.variable}'s
    [invocation-target-in-shadow-tree](#event-path-invocation-target-in-shadow-tree){#ref-for-event-path-invocation-target-in-shadow-tree①
    link-type="dfn"}.

8.  Let `found`{.variable} be the result of running [inner
    invoke](#concept-event-listener-inner-invoke){#ref-for-concept-event-listener-inner-invoke
    link-type="dfn"} with `event`{.variable}, `listeners`{.variable},
    `phase`{.variable}, `invocationTargetInShadowTree`{.variable}, and
    `legacyOutputDidListenersThrowFlag`{.variable} if given.

9.  If `found`{.variable} is false and `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑦
    link-type="idl"} attribute is true:

    1.  Let `originalEventType`{.variable} be `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①①
        link-type="idl"} attribute value.

    2.  If `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①②
        link-type="idl"} attribute value is a match for any of the
        strings in the first column in the following table, set
        `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①③
        link-type="idl"} attribute value to the string in the second
        column on the same row as the matching string, and return
        otherwise.

        Event type

        Legacy event type

        \"`animationend`\"

        \"`webkitAnimationEnd`\"

        \"`animationiteration`\"

        \"`webkitAnimationIteration`\"

        \"`animationstart`\"

        \"`webkitAnimationStart`\"

        \"`transitionend`\"

        \"`webkitTransitionEnd`\"

    3.  [Inner
        invoke](#concept-event-listener-inner-invoke){#ref-for-concept-event-listener-inner-invoke①
        link-type="dfn"} with `event`{.variable},
        `listeners`{.variable}, `phase`{.variable},
        `invocationTargetInShadowTree`{.variable}, and
        `legacyOutputDidListenersThrowFlag`{.variable} if given.

    4.  Set `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①④
        link-type="idl"} attribute value to
        `originalEventType`{.variable}.
:::

::: {.algorithm algorithm="inner invoke"}
To [inner invoke]{#concept-event-listener-inner-invoke .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an `event`{.variable},
`listeners`{.variable}, `phase`{.variable},
`invocationTargetInShadowTree`{.variable}, and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Let `found`{.variable} be false.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑦
    link-type="dfn"} `listener`{.variable} of `listeners`{.variable},
    whose
    [removed](#event-listener-removed){#ref-for-event-listener-removed②
    link-type="dfn"} is false:

    1.  If `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑤
        link-type="idl"} attribute value is not `listener`{.variable}'s
        [type](#event-listener-type){#ref-for-event-listener-type⑨
        link-type="dfn"}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue①
        link-type="dfn"}.

    2.  Set `found`{.variable} to true.

    3.  If `phase`{.variable} is \"`capturing`\" and
        `listener`{.variable}'s
        [capture](#event-listener-capture){#ref-for-event-listener-capture⑦
        link-type="dfn"} is false, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue②
        link-type="dfn"}.

    4.  If `phase`{.variable} is \"`bubbling`\" and
        `listener`{.variable}'s
        [capture](#event-listener-capture){#ref-for-event-listener-capture⑧
        link-type="dfn"} is true, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue③
        link-type="dfn"}.

    5.  If `listener`{.variable}'s
        [once](#event-listener-once){#ref-for-event-listener-once①
        link-type="dfn"} is true, then [remove an event
        listener](#remove-an-event-listener){#ref-for-remove-an-event-listener③
        link-type="dfn"} given `event`{.variable}'s
        [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑧
        link-type="idl"} attribute value and `listener`{.variable}.

    6.  Let `global`{.variable} be `listener`{.variable}
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑧
        link-type="dfn"}'s [associated
        realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm
        link-type="dfn"}'s [global
        object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global
        link-type="dfn"}.

    7.  Let `currentEvent`{.variable} be undefined.

    8.  If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑥
        link-type="idl"} object:

        1.  Set `currentEvent`{.variable} to `global`{.variable}'s
            [current
            event](#window-current-event){#ref-for-window-current-event①
            link-type="dfn"}.

        2.  If `invocationTargetInShadowTree`{.variable} is false, then
            set `global`{.variable}'s [current
            event](#window-current-event){#ref-for-window-current-event②
            link-type="dfn"} to `event`{.variable}.

    9.  If `listener`{.variable}'s
        [passive](#event-listener-passive){#ref-for-event-listener-passive②
        link-type="dfn"} is true, then set `event`{.variable}'s [in
        passive listener
        flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag①
        link-type="dfn"}.

    10. If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑦
        link-type="idl"} object, then [record timing info for event
        listener](https://w3c.github.io/long-animation-frames/#record-timing-info-for-event-listener){#ref-for-record-timing-info-for-event-listener
        link-type="dfn"} given `event`{.variable} and
        `listener`{.variable}.

    11. [Call a user object's
        operation](https://webidl.spec.whatwg.org/#call-a-user-objects-operation){#ref-for-call-a-user-objects-operation
        link-type="dfn"} with `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑨
        link-type="dfn"}, \"`handleEvent`\", « `event`{.variable} », and
        `event`{.variable}'s
        [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑨
        link-type="idl"} attribute value. If this throws an exception
        `exception`{.variable}:

        1.  [Report](https://html.spec.whatwg.org/multipage/webappapis.html#report-an-exception){#ref-for-report-an-exception
            link-type="dfn"} `exception`{.variable} for
            `listener`{.variable}'s
            [callback](#event-listener-callback){#ref-for-event-listener-callback②⓪
            link-type="dfn"}'s corresponding JavaScript object's
            [associated
            realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm①
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global①
            link-type="dfn"}.

        2.  Set `legacyOutputDidListenersThrowFlag`{.variable} if given.

            The `legacyOutputDidListenersThrowFlag`{.variable} is only
            used by Indexed Database API.
            [\[INDEXEDDB\]](#biblio-indexeddb "Indexed Database API"){link-type="biblio"}

    12. Unset `event`{.variable}'s [in passive listener
        flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag②
        link-type="dfn"}.

    13. If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑧
        link-type="idl"} object, then set `global`{.variable}'s [current
        event](#window-current-event){#ref-for-window-current-event③
        link-type="dfn"} to `currentEvent`{.variable}.

    14. If `event`{.variable}'s [stop immediate propagation
        flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag③
        link-type="dfn"} is set, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break①
        link-type="dfn"}.

3.  Return `found`{.variable}.
:::

### [2.10. ]{.secno}[Firing events]{.content}[](#firing-events){.self-link} {#firing-events .heading .settled level="2.10"}

::: {.algorithm algorithm="fire an event"}
To [fire an event]{#concept-event-fire .dfn .dfn-paneled dfn-type="dfn"
export=""} named `e`{.variable} at `target`{.variable}, optionally using
an `eventConstructor`{.variable}, with a description of how IDL
attributes are to be initialized, and a
`legacy target override flag`{.variable}, run these steps:

1.  If `eventConstructor`{.variable} is not given, then let
    `eventConstructor`{.variable} be
    [`Event`{.idl}](#event){#ref-for-event①⑥ link-type="idl"}.

2.  Let `event`{.variable} be the result of [creating an
    event](#concept-event-create){#ref-for-concept-event-create②
    link-type="dfn"} given `eventConstructor`{.variable}, in the
    [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm
    link-type="dfn"} of `target`{.variable}.

3.  Initialize `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑥
    link-type="idl"} attribute to `e`{.variable}.

4.  Initialize any other IDL attributes of `event`{.variable} as
    described in the invocation of this algorithm.

    This also allows for the
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑧
    link-type="idl"} attribute to be set to false.

5.  Return the result of
    [dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②⑤
    link-type="dfn"} `event`{.variable} at `target`{.variable}, with
    `legacy target override flag`{.variable} set if set.
:::

Fire in the context of DOM is short for
[creating](#concept-event-create){#ref-for-concept-event-create③
link-type="dfn"}, initializing, and
[dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②⑥
link-type="dfn"} an [event](#concept-event){#ref-for-concept-event④⑧
link-type="dfn"}. [Fire an
event](#concept-event-fire){#ref-for-concept-event-fire①
link-type="dfn"} makes that process easier to write down.

::: {#firing-events-example .example}
[](#firing-events-example){.self-link}

If the [event](#concept-event){#ref-for-concept-event④⑨ link-type="dfn"}
needs its
[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑦
link-type="idl"} or
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑨
link-type="idl"} attribute initialized, one could write \"[fire an
event](#concept-event-fire){#ref-for-concept-event-fire②
link-type="dfn"} named `submit` at `target`{.variable} with its
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable①⓪
link-type="idl"} attribute initialized to true\".

Or, when a custom constructor is needed, \"[fire an
event](#concept-event-fire){#ref-for-concept-event-fire③
link-type="dfn"} named `click` at `target`{.variable} using
[`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent②
link-type="idl"} with its
[`detail`{.idl}](https://w3c.github.io/uievents/#dom-uievent-detail){#ref-for-dom-uievent-detail
link-type="idl"} attribute initialized to 1\".

Occasionally the return value is important:

1.  Let `doAction`{.variable} be the result of [firing an
    event](#concept-event-fire){#ref-for-concept-event-fire④
    link-type="dfn"} named `like` at `target`{.variable}.

2.  If `doAction`{.variable} is true, then ...
:::

### [2.11. ]{.secno}[]{#action-versus-occurance .bs-old-id}[Action versus occurrence]{.content}[](#action-versus-occurrence){.self-link} {#action-versus-occurrence .heading .settled level="2.11"}

An [event](#concept-event){#ref-for-concept-event⑤⓪ link-type="dfn"}
signifies an occurrence, not an action. Phrased differently, it
represents a notification from an algorithm and can be used to influence
the future course of that algorithm (e.g., through invoking
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑨
link-type="idl"}). [Events](#concept-event){#ref-for-concept-event⑤①
link-type="dfn"} must not be used as actions or initiators that cause
some algorithm to start running. That is not what they are for.

This is called out here specifically because previous iterations of the
DOM had a concept of \"default actions\" associated with
[events](#concept-event){#ref-for-concept-event⑤② link-type="dfn"} that
gave folks all the wrong ideas.
[Events](#concept-event){#ref-for-concept-event⑤③ link-type="dfn"} do
not represent or cause actions, they can only be used to influence an
ongoing one.

