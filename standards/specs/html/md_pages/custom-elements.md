HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.12.5 The canvas element](canvas.html) --- [Table of Contents](./)
--- [4.14 Common idioms without dedicated elements
→](semantics-other.html)

1.  ::: {#toc-semantics}
    1.  [[4.13]{.secno} Custom
        elements](custom-elements.html#custom-elements)
        1.  [[4.13.1]{.secno}
            Introduction](custom-elements.html#custom-elements-intro)
            1.  [[4.13.1.1]{.secno} Creating an autonomous custom
                element](custom-elements.html#custom-elements-autonomous-example)
            2.  [[4.13.1.2]{.secno} Creating a form-associated custom
                element](custom-elements.html#custom-elements-face-example)
            3.  [[4.13.1.3]{.secno} Creating a custom element with
                default accessible roles, states, and
                properties](custom-elements.html#custom-elements-accessibility-example)
            4.  [[4.13.1.4]{.secno} Creating a customized built-in
                element](custom-elements.html#custom-elements-customized-builtin-example)
            5.  [[4.13.1.5]{.secno} Drawbacks of autonomous custom
                elements](custom-elements.html#custom-elements-autonomous-drawbacks)
            6.  [[4.13.1.6]{.secno} Upgrading elements after their
                creation](custom-elements.html#custom-elements-upgrades-examples)
            7.  [[4.13.1.7]{.secno} Exposing custom element
                states](custom-elements.html#exposing-custom-element-states)
        2.  [[4.13.2]{.secno} Requirements for custom element
            constructors and
            reactions](custom-elements.html#custom-element-conformance)
            1.  [[4.13.2.1]{.secno} Preserving custom element state when
                moved](custom-elements.html#preserving-custom-element-state-when-moved)
        3.  [[4.13.3]{.secno} Core
            concepts](custom-elements.html#custom-elements-core-concepts)
        4.  [[4.13.4]{.secno} The `CustomElementRegistry`
            interface](custom-elements.html#custom-elements-api)
        5.  [[4.13.5]{.secno} Upgrades](custom-elements.html#upgrades)
        6.  [[4.13.6]{.secno} Custom element
            reactions](custom-elements.html#custom-element-reactions)
        7.  [[4.13.7]{.secno} Element
            internals](custom-elements.html#element-internals)
            1.  [[4.13.7.1]{.secno} The `ElementInternals`
                interface](custom-elements.html#the-elementinternals-interface)
            2.  [[4.13.7.2]{.secno} Shadow root
                access](custom-elements.html#shadow-root-access)
            3.  [[4.13.7.3]{.secno} Form-associated custom
                elements](custom-elements.html#form-associated-custom-elements)
            4.  [[4.13.7.4]{.secno} Accessibility
                semantics](custom-elements.html#accessibility-semantics)
            5.  [[4.13.7.5]{.secno} Custom state
                pseudo-class](custom-elements.html#custom-state-pseudo-class)
    :::

### [4.13]{.secno} Custom elements[](#custom-elements){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Using_custom_elements](https://developer.mozilla.org/en-US/docs/Web/Web_Components/Using_custom_elements "One of the key features of the Web Components standard is the ability to create custom elements that encapsulate your functionality on an HTML page, rather than having to make do with a long, nested batch of elements that together provide a custom page feature. This article introduces the use of the Custom Elements API.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

#### [4.13.1]{.secno} Introduction[](#custom-elements-intro){.self-link} {#custom-elements-intro}

*This section is non-normative.*

[Custom
elements](#custom-element){#custom-elements-intro:custom-element}
provide a way for authors to build their own fully-featured DOM
elements. Although authors could always use non-standard elements in
their documents, with application-specific behavior added after the fact
by scripting or similar, such elements have historically been
non-conforming and not very functional. By
[defining](#element-definition){#custom-elements-intro:element-definition}
a custom element, authors can inform the parser how to properly
construct an element and how elements of that class should react to
changes.

Custom elements are part of a larger effort to \"rationalise the
platform\", by explaining existing platform features (like the elements
of HTML) in terms of lower-level author-exposed extensibility points
(like custom element definition). Although today there are many
limitations on the capabilities of custom elements---both functionally
and semantically---that prevent them from fully explaining the behaviors
of HTML\'s existing elements, we hope to shrink this gap over time.

##### [4.13.1.1]{.secno} Creating an autonomous custom element[](#custom-elements-autonomous-example){.self-link} {#custom-elements-autonomous-example}

*This section is non-normative.*

For the purposes of illustrating how to create an [autonomous custom
element](#autonomous-custom-element){#custom-elements-autonomous-example:autonomous-custom-element},
let\'s define a custom element that encapsulates rendering a small icon
for a country flag. Our goal is to be able to use it like so:

``` html
<flag-icon country="nl"></flag-icon>
```

To do this, we first declare a class for the custom element, extending
[`HTMLElement`{#custom-elements-autonomous-example:htmlelement}](dom.html#htmlelement):

``` js
class FlagIcon extends HTMLElement {
  constructor() {
    super();
    this._countryCode = null;
  }

  static observedAttributes = ["country"];

  attributeChangedCallback(name, oldValue, newValue) {
    // name will always be "country" due to observedAttributes
    this._countryCode = newValue;
    this._updateRendering();
  }
  connectedCallback() {
    this._updateRendering();
  }

  get country() {
    return this._countryCode;
  }
  set country(v) {
    this.setAttribute("country", v);
  }

  _updateRendering() {
    // Left as an exercise for the reader. But, you'll probably want to
    // check this.ownerDocument.defaultView to see if we've been
    // inserted into a document with a browsing context, and avoid
    // doing any work if not.
  }
}
```

We then need to use this class to define the element:

``` js
customElements.define("flag-icon", FlagIcon);
```

At this point, our above code will work! The parser, whenever it sees
the `flag-icon` tag, will construct a new instance of our `FlagIcon`
class, and tell our code about its new `country` attribute, which we
then use to set the element\'s internal state and update its rendering
(when appropriate).

You can also create `flag-icon` elements using the DOM API:

``` js
const flagIcon = document.createElement("flag-icon")
flagIcon.country = "jp"
document.body.appendChild(flagIcon)
```

Finally, we can also use the [custom element
constructor](#custom-element-constructor){#custom-elements-autonomous-example:custom-element-constructor}
itself. That is, the above code is equivalent to:

``` js
const flagIcon = new FlagIcon()
flagIcon.country = "jp"
document.body.appendChild(flagIcon)
```

##### [4.13.1.2]{.secno} Creating a form-associated custom element[](#custom-elements-face-example){.self-link} {#custom-elements-face-example}

*This section is non-normative.*

Adding a static `formAssociated` property, with a true value, makes an
[autonomous custom
element](#autonomous-custom-element){#custom-elements-face-example:autonomous-custom-element}
a [form-associated custom
element](#form-associated-custom-element){#custom-elements-face-example:form-associated-custom-element}.
The
[`ElementInternals`{#custom-elements-face-example:elementinternals}](#elementinternals)
interface helps you to implement functions and properties common to form
control elements.

``` js
class MyCheckbox extends HTMLElement {
  static formAssociated = true;
  static observedAttributes = ['checked'];

  constructor() {
    super();
    this._internals = this.attachInternals();
    this.addEventListener('click', this._onClick.bind(this));
  }

  get form() { return this._internals.form; }
  get name() { return this.getAttribute('name'); }
  get type() { return this.localName; }

  get checked() { return this.hasAttribute('checked'); }
  set checked(flag) { this.toggleAttribute('checked', Boolean(flag)); }

  attributeChangedCallback(name, oldValue, newValue) {
    // name will always be "checked" due to observedAttributes
    this._internals.setFormValue(this.checked ? 'on' : null);
  }

  _onClick(event) {
    this.checked = !this.checked;
  }
}
customElements.define('my-checkbox', MyCheckbox);
```

You can use the custom element `my-checkbox` like a built-in
form-associated element. For example, putting it in
[`form`{#custom-elements-face-example:the-form-element}](forms.html#the-form-element)
or
[`label`{#custom-elements-face-example:the-label-element}](forms.html#the-label-element)
associates the `my-checkbox` element with them, and submitting the
[`form`{#custom-elements-face-example:the-form-element-2}](forms.html#the-form-element)
will send data provided by `my-checkbox` implementation.

``` html
<form action="..." method="...">
  <label><my-checkbox name="agreed"></my-checkbox> I read the agreement.</label>
  <input type="submit">
</form>
```

##### [4.13.1.3]{.secno} Creating a custom element with default accessible roles, states, and properties[](#custom-elements-accessibility-example){.self-link} {#custom-elements-accessibility-example}

*This section is non-normative.*

By using the appropriate properties of
[`ElementInternals`{#custom-elements-accessibility-example:elementinternals}](#elementinternals),
your custom element can have default accessibility semantics. The
following code expands our form-associated checkbox from the previous
section to properly set its default role and checkedness, as viewed by
accessibility technology:

``` js
class MyCheckbox extends HTMLElement {
  static formAssociated = true;
  static observedAttributes = ['checked'];

  constructor() {
    super();
    this._internals = this.attachInternals();
    this.addEventListener('click', this._onClick.bind(this));

    this._internals.role = 'checkbox';
    this._internals.ariaChecked = 'false';
  }

  get form() { return this._internals.form; }
  get name() { return this.getAttribute('name'); }
  get type() { return this.localName; }

  get checked() { return this.hasAttribute('checked'); }
  set checked(flag) { this.toggleAttribute('checked', Boolean(flag)); }

  attributeChangedCallback(name, oldValue, newValue) {
    // name will always be "checked" due to observedAttributes
    this._internals.setFormValue(this.checked ? 'on' : null);
    this._internals.ariaChecked = this.checked;
  }

  _onClick(event) {
    this.checked = !this.checked;
  }
}
customElements.define('my-checkbox', MyCheckbox);
```

Note that, like for built-in elements, these are only defaults, and can
be overridden by the page author using the
[`role`{#custom-elements-accessibility-example:attr-aria-role}](infrastructure.html#attr-aria-role)
and
[`aria-*`{#custom-elements-accessibility-example:attr-aria-*}](infrastructure.html#attr-aria-*)
attributes:

``` bad
<!-- This markup is non-conforming -->
<input type="checkbox" checked role="button" aria-checked="false">
```

``` bad
<!-- This markup is probably not what the custom element author intended -->
<my-checkbox role="button" checked aria-checked="false">
```

Custom element authors are encouraged to state what aspects of their
accessibility semantics are strong native semantics, i.e., should not be
overridden by users of the custom element. In our example, the author of
the `my-checkbox` element would state that its
[role](https://w3c.github.io/aria/#dfn-role){#custom-elements-accessibility-example:role
x-internal="role"} and
[`aria-checked`{#custom-elements-accessibility-example:attr-aria-checked}](https://w3c.github.io/aria/#aria-checked){x-internal="attr-aria-checked"}
values are strong native semantics, thus discouraging code such as the
above.

##### [4.13.1.4]{.secno} Creating a customized built-in element[](#custom-elements-customized-builtin-example){.self-link} {#custom-elements-customized-builtin-example}

*This section is non-normative.*

[Customized built-in
elements](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element}
are a distinct kind of [custom
element](#custom-element){#custom-elements-customized-builtin-example:custom-element},
which are defined slightly differently and used very differently
compared to [autonomous custom
elements](#autonomous-custom-element){#custom-elements-customized-builtin-example:autonomous-custom-element}.
They exist to allow reuse of behaviors from the existing elements of
HTML, by extending those elements with new custom functionality. This is
important since many of the existing behaviors of HTML elements can
unfortunately not be duplicated by using purely [autonomous custom
elements](#autonomous-custom-element){#custom-elements-customized-builtin-example:autonomous-custom-element-2}.
Instead, [customized built-in
elements](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-2}
allow the installation of custom construction behavior, lifecycle hooks,
and prototype chain onto existing elements, essentially \"mixing in\"
these capabilities on top of the already-existing element.

[Customized built-in
elements](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-3}
require a distinct syntax from [autonomous custom
elements](#autonomous-custom-element){#custom-elements-customized-builtin-example:autonomous-custom-element-3}
because user agents and other software key off an element\'s local name
in order to identify the element\'s semantics and behavior. That is, the
concept of [customized built-in
elements](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-4}
building on top of existing behavior depends crucially on the extended
elements retaining their original local name.

In this example, we\'ll be creating a [customized built-in
element](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-5}
named `plastic-button`, which behaves like a normal button but gets
fancy animation effects added whenever you click on it. We start by
defining a class, just like before, although this time we extend
[`HTMLButtonElement`{#custom-elements-customized-builtin-example:htmlbuttonelement}](form-elements.html#htmlbuttonelement)
instead of
[`HTMLElement`{#custom-elements-customized-builtin-example:htmlelement}](dom.html#htmlelement):

``` js
class PlasticButton extends HTMLButtonElement {
  constructor() {
    super();

    this.addEventListener("click", () => {
      // Draw some fancy animation effects!
    });
  }
}
```

When defining our custom element, we have to also specify the
[`extends`{#custom-elements-customized-builtin-example:dom-elementdefinitionoptions-extends}](#dom-elementdefinitionoptions-extends)
option:

``` js
customElements.define("plastic-button", PlasticButton, { extends: "button" });
```

In general, the name of the element being extended cannot be determined
simply by looking at what element interface it extends, as many elements
share the same interface (such as
[`q`{#custom-elements-customized-builtin-example:the-q-element}](text-level-semantics.html#the-q-element)
and
[`blockquote`{#custom-elements-customized-builtin-example:the-blockquote-element}](grouping-content.html#the-blockquote-element)
both sharing
[`HTMLQuoteElement`{#custom-elements-customized-builtin-example:htmlquoteelement}](grouping-content.html#htmlquoteelement)).

To construct our [customized built-in
element](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-6}
from parsed HTML source text, we use the
[`is`{#custom-elements-customized-builtin-example:attr-is}](#attr-is)
attribute on a
[`button`{#custom-elements-customized-builtin-example:the-button-element}](form-elements.html#the-button-element)
element:

``` html
<button is="plastic-button">Click Me!</button>
```

Trying to use a [customized built-in
element](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-7}
as an [autonomous custom
element](#autonomous-custom-element){#custom-elements-customized-builtin-example:autonomous-custom-element-4}
will *not* work; that is, `<plastic-button>Click me?</plastic-button>`
will simply create an
[`HTMLElement`{#custom-elements-customized-builtin-example:htmlelement-2}](dom.html#htmlelement)
with no special behavior.

If you need to create a customized built-in element programmatically,
you can use the following form of
[`createElement()`{#custom-elements-customized-builtin-example:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}:

``` js
const plasticButton = document.createElement("button", { is: "plastic-button" });
plasticButton.textContent = "Click me!";
```

And as before, the constructor will also work:

``` js
const plasticButton2 = new PlasticButton();
console.log(plasticButton2.localName);  // will output "button"
console.assert(plasticButton2 instanceof PlasticButton);
console.assert(plasticButton2 instanceof HTMLButtonElement);
```

Note that when creating a customized built-in element programmatically,
the
[`is`{#custom-elements-customized-builtin-example:attr-is-2}](#attr-is)
attribute will not be present in the DOM, since it was not explicitly
set. However, [it will be added to the output when
serializing](parsing.html#attr-is-during-serialization):

``` js
console.assert(!plasticButton.hasAttribute("is"));
console.log(plasticButton.outerHTML); // will output '<button is="plastic-button"></button>'
```

Regardless of how it is created, all of the ways in which
[`button`{#custom-elements-customized-builtin-example:the-button-element-2}](form-elements.html#the-button-element)
is special apply to such \"plastic buttons\" as well: their focus
behavior, ability to participate in [form
submission](form-control-infrastructure.html#concept-form-submit){#custom-elements-customized-builtin-example:concept-form-submit},
the
[`disabled`{#custom-elements-customized-builtin-example:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
attribute, and so on.

[Customized built-in
elements](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-8}
are designed to allow extension of existing HTML elements that have
useful user-agent supplied behavior or APIs. As such, they can only
extend existing HTML elements defined in this specification, and cannot
extend legacy elements such as
[`bgsound`{#custom-elements-customized-builtin-example:bgsound}](obsolete.html#bgsound),
[`blink`{#custom-elements-customized-builtin-example:blink}](obsolete.html#blink),
[`isindex`{#custom-elements-customized-builtin-example:isindex}](obsolete.html#isindex),
[`keygen`{#custom-elements-customized-builtin-example:keygen}](obsolete.html#keygen),
[`multicol`{#custom-elements-customized-builtin-example:multicol}](obsolete.html#multicol),
[`nextid`{#custom-elements-customized-builtin-example:nextid}](obsolete.html#nextid),
or
[`spacer`{#custom-elements-customized-builtin-example:spacer}](obsolete.html#spacer)
that have been defined to use
[`HTMLUnknownElement`{#custom-elements-customized-builtin-example:htmlunknownelement}](dom.html#htmlunknownelement)
as their [element
interface](https://dom.spec.whatwg.org/#concept-element-interface){#custom-elements-customized-builtin-example:element-interface
x-internal="element-interface"}.

One reason for this requirement is future-compatibility: if a
[customized built-in
element](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-9}
was defined that extended a currently-unknown element, for example
`combobox`, this would prevent this specification from defining a
`combobox` element in the future, as consumers of the derived
[customized built-in
element](#customized-built-in-element){#custom-elements-customized-builtin-example:customized-built-in-element-10}
would have come to depend on their base element having no interesting
user-agent-supplied behavior.

##### [4.13.1.5]{.secno} Drawbacks of autonomous custom elements[](#custom-elements-autonomous-drawbacks){.self-link} {#custom-elements-autonomous-drawbacks}

*This section is non-normative.*

As specified below, and alluded to above, simply defining and using an
element called `taco-button` does not mean that such elements
[represent](dom.html#represents){#custom-elements-autonomous-drawbacks:represents}
buttons. That is, tools such as web browsers, search engines, or
accessibility technology will not automatically treat the resulting
element as a button just based on its defined name.

To convey the desired button semantics to a variety of users, while
still using an [autonomous custom
element](#autonomous-custom-element){#custom-elements-autonomous-drawbacks:autonomous-custom-element},
a number of techniques would need to be employed:

- The addition of the
  [`tabindex`{#custom-elements-autonomous-drawbacks:attr-tabindex}](interaction.html#attr-tabindex)
  attribute would make the `taco-button`
  [focusable](interaction.html#focusable){#custom-elements-autonomous-drawbacks:focusable}.
  Note that if the `taco-button` were to become logically disabled, the
  [`tabindex`{#custom-elements-autonomous-drawbacks:attr-tabindex-2}](interaction.html#attr-tabindex)
  attribute would need to be removed.

- The addition of an ARIA role and various ARIA states and properties
  helps convey semantics to accessibility technology. For example,
  setting the
  [role](https://w3c.github.io/aria/#dfn-role){#custom-elements-autonomous-drawbacks:role
  x-internal="role"} to
  \"[`button`{#custom-elements-autonomous-drawbacks:attr-aria-role-button}](https://w3c.github.io/aria/#button){x-internal="attr-aria-role-button"}\"
  will convey the semantics that this is a button, enabling users to
  successfully interact with the control using usual button-like
  interactions in their accessibility technology. Setting the
  [`aria-label`{#custom-elements-autonomous-drawbacks:attr-aria-label}](https://w3c.github.io/aria/#aria-label){x-internal="attr-aria-label"}
  property is necessary to give the button an [accessible
  name](https://w3c.github.io/aria/#dfn-accessible-name){#custom-elements-autonomous-drawbacks:concept-accessible-name
  x-internal="concept-accessible-name"}, instead of having accessibility
  technology traverse its child text nodes and announce them. And
  setting the
  [`aria-disabled`{#custom-elements-autonomous-drawbacks:attr-aria-disabled}](https://w3c.github.io/aria/#aria-disabled){x-internal="attr-aria-disabled"}
  state to \"`true`\" when the button is logically disabled conveys to
  accessibility technology the button\'s disabled state.

- The addition of event handlers to handle commonly-expected button
  behaviors helps convey the semantics of the button to web browser
  users. In this case, the most relevant event handler would be one that
  proxies appropriate
  [`keydown`{#custom-elements-autonomous-drawbacks:event-keydown}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
  events to become
  [`click`{#custom-elements-autonomous-drawbacks:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
  events, so that you can activate the button both with keyboard and by
  clicking.

- In addition to any default visual styling provided for `taco-button`
  elements, the visual styling will also need to be updated to reflect
  changes in logical state, such as becoming disabled; that is, whatever
  style sheet has rules for `taco-button` will also need to have rules
  for `taco-button[disabled]`.

With these points in mind, a full-featured `taco-button` that took on
the responsibility of conveying button semantics (including the ability
to be disabled) might look something like this:

``` {#custom-elements-autonomous-drawbacks-example}
class TacoButton extends HTMLElement {
  static observedAttributes = ["disabled"];

  constructor() {
    super();
    this._internals = this.attachInternals();
    this._internals.role = "button";

    this.addEventListener("keydown", e => {
      if (e.code === "Enter" || e.code === "Space") {
        this.dispatchEvent(new PointerEvent("click", {
          bubbles: true,
          cancelable: true
        }));
      }
    });

    this.addEventListener("click", e => {
      if (this.disabled) {
        e.preventDefault();
        e.stopImmediatePropagation();
      }
    });

    this._observer = new MutationObserver(() => {
      this._internals.ariaLabel = this.textContent;
    });
  }

  connectedCallback() {
    this.setAttribute("tabindex", "0");

    this._observer.observe(this, {
      childList: true,
      characterData: true,
      subtree: true
    });
  }

  disconnectedCallback() {
    this._observer.disconnect();
  }

  get disabled() {
    return this.hasAttribute("disabled");
  }
  set disabled(flag) {
    this.toggleAttribute("disabled", Boolean(flag));
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // name will always be "disabled" due to observedAttributes
    if (this.disabled) {
      this.removeAttribute("tabindex");
      this._internals.ariaDisabled = "true";
    } else {
      this.setAttribute("tabindex", "0");
      this._internals.ariaDisabled = "false";
    }
  }
}
```

Even with this rather-complicated element definition, the element is not
a pleasure to use for consumers: it will be continually \"sprouting\"
[`tabindex`{#custom-elements-autonomous-drawbacks:attr-tabindex-3}](interaction.html#attr-tabindex)
attributes of its own volition, and its choice of `tabindex="0"`
focusability behavior may not match the
[`button`{#custom-elements-autonomous-drawbacks:the-button-element}](form-elements.html#the-button-element)
behavior on the current platform. This is because as of now there is no
way to specify default focus behavior for custom elements, forcing the
use of the
[`tabindex`{#custom-elements-autonomous-drawbacks:attr-tabindex-4}](interaction.html#attr-tabindex)
attribute to do so (even though it is usually reserved for allowing the
consumer to override default behavior).

In contrast, a simple [customized built-in
element](#customized-built-in-element){#custom-elements-autonomous-drawbacks:customized-built-in-element},
as shown in the previous section, would automatically inherit the
semantics and behavior of the
[`button`{#custom-elements-autonomous-drawbacks:the-button-element-2}](form-elements.html#the-button-element)
element, with no need to implement these behaviors manually. In general,
for any elements with nontrivial behavior and semantics that build on
top of existing elements of HTML, [customized built-in
elements](#customized-built-in-element){#custom-elements-autonomous-drawbacks:customized-built-in-element-2}
will be easier to develop, maintain, and consume.

##### [4.13.1.6]{.secno} Upgrading elements after their creation[](#custom-elements-upgrades-examples){.self-link} {#custom-elements-upgrades-examples}

*This section is non-normative.*

Because [element
definition](#element-definition){#custom-elements-upgrades-examples:element-definition}
can occur at any time, a non-custom element could be
[created](https://dom.spec.whatwg.org/#concept-create-element){#custom-elements-upgrades-examples:create-an-element
x-internal="create-an-element"}, and then later become a [custom
element](#custom-element){#custom-elements-upgrades-examples:custom-element}
after an appropriate
[definition](#custom-element-definition){#custom-elements-upgrades-examples:custom-element-definition}
is registered. We call this process \"upgrading\" the element, from a
normal element into a custom element.

[Upgrades](#upgrades){#custom-elements-upgrades-examples:upgrades}
enable scenarios where it may be preferable for [custom element
definitions](#custom-element-definition){#custom-elements-upgrades-examples:custom-element-definition-2}
to be registered after relevant elements have been initially created,
such as by the parser. They allow progressive enhancement of the content
in the custom element. For example, in the following HTML document the
element definition for `img-viewer` is loaded asynchronously:

``` html
<!DOCTYPE html>
<html lang="en">
<title>Image viewer example</title>

<img-viewer filter="Kelvin">
  <img src="images/tree.jpg" alt="A beautiful tree towering over an empty savannah">
</img-viewer>

<script src="js/elements/img-viewer.js" async></script>
```

The definition for the `img-viewer` element here is loaded using a
[`script`{#custom-elements-upgrades-examples:the-script-element}](scripting.html#the-script-element)
element marked with the
[`async`{#custom-elements-upgrades-examples:attr-script-async}](scripting.html#attr-script-async)
attribute, placed after the `<img-viewer>` tag in the markup. While the
script is loading, the `img-viewer` element will be treated as an
undefined element, similar to a
[`span`{#custom-elements-upgrades-examples:the-span-element}](text-level-semantics.html#the-span-element).
Once the script loads, it will define the `img-viewer` element, and the
existing `img-viewer` element on the page will be upgraded, applying the
custom element\'s definition (which presumably includes applying an
image filter identified by the string \"Kelvin\", enhancing the image\'s
visual appearance).

------------------------------------------------------------------------

Note that
[upgrades](#upgrades){#custom-elements-upgrades-examples:upgrades-2}
only apply to elements in the document tree. (Formally, elements that
are
[connected](https://dom.spec.whatwg.org/#connected){#custom-elements-upgrades-examples:connected
x-internal="connected"}.) An element that is not inserted into a
document will stay un-upgraded. An example illustrates this point:

``` html
<!DOCTYPE html>
<html lang="en">
<title>Upgrade edge-cases example</title>

<example-element></example-element>

<script>
  "use strict";

  const inDocument = document.querySelector("example-element");
  const outOfDocument = document.createElement("example-element");

  // Before the element definition, both are HTMLElement:
  console.assert(inDocument instanceof HTMLElement);
  console.assert(outOfDocument instanceof HTMLElement);

  class ExampleElement extends HTMLElement {}
  customElements.define("example-element", ExampleElement);

  // After element definition, the in-document element was upgraded:
  console.assert(inDocument instanceof ExampleElement);
  console.assert(!(outOfDocument instanceof ExampleElement));

  document.body.appendChild(outOfDocument);

  // Now that we've moved the element into the document, it too was upgraded:
  console.assert(outOfDocument instanceof ExampleElement);
</script>
```

##### [4.13.1.7]{.secno} Exposing custom element states[](#exposing-custom-element-states){.self-link}

Built-in elements provided by user agents have certain states that can
change over time depending on user interaction and other factors, and
are exposed to web authors through
[pseudo-classes](https://drafts.csswg.org/selectors/#pseudo-class){#exposing-custom-element-states:pseudo-class
x-internal="pseudo-class"}. For example, some form controls have the
\"invalid\" state, which is exposed through the
[`:invalid`{#exposing-custom-element-states:selector-invalid}](semantics-other.html#selector-invalid)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#exposing-custom-element-states:pseudo-class-2
x-internal="pseudo-class"}.

Like built-in elements, [custom
elements](#custom-element){#exposing-custom-element-states:custom-element}
can have various states to be in too, and [custom
element](#custom-element){#exposing-custom-element-states:custom-element-2}
authors want to expose these states in a similar fashion as the built-in
elements.

This is done via the
[`:state()`{#exposing-custom-element-states:selector-custom}](semantics-other.html#selector-custom)
pseudo-class. A custom element author can use the
[`states`{#exposing-custom-element-states:dom-elementinternals-states}](#dom-elementinternals-states)
property of
[`ElementInternals`{#exposing-custom-element-states:elementinternals}](#elementinternals)
to add and remove such custom states, which are then exposed as
arguments to the
[`:state()`{#exposing-custom-element-states:selector-custom-2}](semantics-other.html#selector-custom)
pseudo-class.

::: example
The following shows how
[`:state()`{#exposing-custom-element-states:selector-custom-3}](semantics-other.html#selector-custom)
can be used to style a custom checkbox element. Assume that
`LabeledCheckbox` doesn\'t expose its \"checked\" state via a content
attribute.

``` html
<script>
class LabeledCheckbox extends HTMLElement {
  constructor() {
    super();
    this._internals = this.attachInternals();
    this.addEventListener('click', this._onClick.bind(this));

    const shadowRoot = this.attachShadow({mode: 'closed'});
    shadowRoot.innerHTML =
      `<style>
       :host::before {
         content: '[ ]';
         white-space: pre;
         font-family: monospace;
       }
       :host(:state(checked))::before { content: '[x]' }
       </style>
       <slot>Label</slot>`;
  }

  get checked() { return this._internals.states.has('checked'); }

  set checked(flag) {
    if (flag)
      this._internals.states.add('checked');
    else
      this._internals.states.delete('checked');
  }

  _onClick(event) {
    this.checked = !this.checked;
  }
}

customElements.define('labeled-checkbox', LabeledCheckbox);
</script>

<style>
labeled-checkbox { border: dashed red; }
labeled-checkbox:state(checked) { border: solid; }
</style>

<labeled-checkbox>You need to check this</labeled-checkbox>
```
:::

::: example
Custom pseudo-classes can even target shadow parts. An extension of the
above example shows this:

``` html
<script>
class QuestionBox extends HTMLElement {
  constructor() {
    super();
    const shadowRoot = this.attachShadow({mode: 'closed'});
    shadowRoot.innerHTML =
      `<div><slot>Question</slot></div>
       <labeled-checkbox part='checkbox'>Yes</labeled-checkbox>`;
  }
}
customElements.define('question-box', QuestionBox);
</script>

<style>
question-box::part(checkbox) { color: red; }
question-box::part(checkbox):state(checked) { color: green; }
</style>

<question-box>Continue?</question-box>
```
:::

#### [4.13.2]{.secno} Requirements for custom element constructors and reactions[](#custom-element-conformance){.self-link} {#custom-element-conformance}

When authoring [custom element
constructors](#custom-element-constructor){#custom-element-conformance:custom-element-constructor},
authors are bound by the following conformance requirements:

- A parameter-less call to `super()` must be the first statement in the
  constructor body, to establish the correct prototype chain and
  **this** value before any further code is run.

- A `return` statement must not appear anywhere inside the constructor
  body, unless it is a simple early-return (`return` or `return this`).

- The constructor must not use the
  [`document.write()`{#custom-element-conformance:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
  or
  [`document.open()`{#custom-element-conformance:dom-document-open}](dynamic-markup-insertion.html#dom-document-open)
  methods.

- The element\'s attributes and children must not be inspected, as in
  the non-[upgrade](#upgrades){#custom-element-conformance:upgrades}
  case none will be present, and relying on upgrades makes the element
  less usable.

- The element must not gain any attributes or children, as this violates
  the expectations of consumers who use the
  [`createElement`{#custom-element-conformance:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}
  or
  [`createElementNS`{#custom-element-conformance:dom-document-createelementns}](https://dom.spec.whatwg.org/#dom-document-createelementns){x-internal="dom-document-createelementns"}
  methods.

- In general, work should be deferred to `connectedCallback` as much as
  possible---especially work involving fetching resources or rendering.
  However, note that `connectedCallback` can be called more than once,
  so any initialization work that is truly one-time will need a guard to
  prevent it from running twice.

- In general, the constructor should be used to set up initial state and
  default values, and to set up event listeners and possibly a [shadow
  root](https://dom.spec.whatwg.org/#concept-shadow-root){#custom-element-conformance:shadow-root
  x-internal="shadow-root"}.

Several of these requirements are checked during [element
creation](https://dom.spec.whatwg.org/#concept-create-element){#custom-element-conformance:create-an-element
x-internal="create-an-element"}, either directly or indirectly, and
failing to follow them will result in a custom element that cannot be
instantiated by the parser or DOM APIs. This is true even if the work is
done inside a constructor-initiated
[microtask](webappapis.html#microtask){#custom-element-conformance:microtask},
as a [microtask
checkpoint](webappapis.html#perform-a-microtask-checkpoint){#custom-element-conformance:perform-a-microtask-checkpoint}
can occur immediately after construction.

When authoring [custom element
reactions](#concept-custom-element-reaction){#custom-element-conformance:concept-custom-element-reaction},
authors should avoid manipulating the node tree as this can lead to
unexpected results.

::: example
An element\'s `connectedCallback` can be queued before the element is
disconnected, but as the callback queue is still processed, it results
in a `connectedCallback` for an element that is no longer connected:

``` js
class CParent extends HTMLElement {
  connectedCallback() {
    this.firstChild.remove();
  }
}
customElements.define("c-parent", CParent);

class CChild extends HTMLElement {
  connectedCallback() {
    console.log("CChild connectedCallback: isConnected =", this.isConnected);
  }
}
customElements.define("c-child", CChild);

const parent = new CParent(),
      child = new CChild();
parent.append(child);
document.body.append(parent);

// Logs:
// CChild connectedCallback: isConnected = false
```
:::

##### [4.13.2.1]{.secno} Preserving custom element state when moved[](#preserving-custom-element-state-when-moved){.self-link}

*This section is non-normative.*

When manipulating the DOM tree, an element can be
[moved](https://dom.spec.whatwg.org/#concept-node-move-ext){#preserving-custom-element-state-when-moved:concept-node-move-ext
x-internal="concept-node-move-ext"} in the tree while connected. This
applies to custom elements as well. By default, the
\"`disconnectedCallback`\" and \"`connectedCallback`\" would be called
on the element, one after the other. This is done to maintain
compatibility with existing custom elements that predate the
[`moveBefore()`{#preserving-custom-element-state-when-moved:dom-parentnode-movebefore}](https://dom.spec.whatwg.org/#dom-parentnode-movebefore){x-internal="dom-parentnode-movebefore"}
method. This means that by default, custom elements reset their state as
if they were removed and re-inserted. In the example
[above](#custom-elements-autonomous-drawbacks-example), the impact would
be that the observer would be disconnected and re-connected, and the tab
index would be reset.

To opt in to a state-preserving behavior while
[moving](https://dom.spec.whatwg.org/#concept-node-move-ext){#preserving-custom-element-state-when-moved:concept-node-move-ext-2
x-internal="concept-node-move-ext"}, the author can implement a
\"`connectedMoveCallback`\". The existence of this callback, even if
empty, would supercede the default behavior of calling
\"`disconnectedCallback`\" and \"`connectedCallback`\".
\"`connectedMoveCallback`\" can also be an appropriate place to execute
logic that depends on the element\'s ancestors. For example:

``` js
class TacoButton extends HTMLElement {
  static observedAttributes = ["disabled"];

  constructor() {
    super();
    this._internals = this.attachInternals();
    this._internals.role = "button";

    this._observer = new MutationObserver(() => {
      this._internals.ariaLabel = this.textContent;
    });
  }

  _notifyMain() {
    if (this.parentElement.tagName === "MAIN") {
      // Execute logic that depends on ancestors.
    }
  }

  connectedCallback() {
    this.setAttribute("tabindex", "0");

    this._observer.observe(this, {
      childList: true,
      characterData: true,
      subtree: true
    });

    this._notifyMain();
  }

  disconnectedCallback() {
    this._observer.disconnect();
  }

  // Implementing this function would avoid resetting the tab index or re-registering the
  // mutation observer when this element is moved inside the DOM without being disconnected.
  connectedMoveCallback() {
    // The parent can change during a state-preserving move.
    this._notifyMain();
  }
}
```

#### [4.13.3]{.secno} Core concepts[](#custom-elements-core-concepts){.self-link} {#custom-elements-core-concepts}

A [custom element]{#custom-element .dfn export=""} is an element that is
[custom](https://dom.spec.whatwg.org/#concept-element-custom){#custom-elements-core-concepts:concept-element-custom
x-internal="concept-element-custom"}. Informally, this means that its
constructor and prototype are defined by the author, instead of by the
user agent. This author-supplied constructor function is called the
[custom element constructor]{#custom-element-constructor .dfn
export=""}.

Two distinct types of [custom
elements](#custom-element){#custom-elements-core-concepts:custom-element}
can be defined:

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Global_attributes/is](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/is "The is global attribute allows you to specify that a standard HTML element should behave like a defined custom built-in element (see Using custom elements for more details).")

::: support
[Firefox63+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome67+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

1.  An [autonomous custom element]{#autonomous-custom-element .dfn
    export=""}, which is defined with no
    [`extends`{#custom-elements-core-concepts:dom-elementdefinitionoptions-extends}](#dom-elementdefinitionoptions-extends)
    option. These types of custom elements have a local name equal to
    their [defined
    name](#concept-custom-element-definition-name){#custom-elements-core-concepts:concept-custom-element-definition-name}.

2.  A [customized built-in element]{#customized-built-in-element .dfn
    export=""}, which is defined with an
    [`extends`{#custom-elements-core-concepts:dom-elementdefinitionoptions-extends-2}](#dom-elementdefinitionoptions-extends)
    option. These types of custom elements have a local name equal to
    the value passed in their
    [`extends`{#custom-elements-core-concepts:dom-elementdefinitionoptions-extends-3}](#dom-elementdefinitionoptions-extends)
    option, and their [defined
    name](#concept-custom-element-definition-name){#custom-elements-core-concepts:concept-custom-element-definition-name-2}
    is used as the value of the [`is`]{#attr-is .dfn
    dfn-for="html-global" dfn-type="element-attr"} attribute, which
    therefore must be a [valid custom element
    name](#valid-custom-element-name){#custom-elements-core-concepts:valid-custom-element-name}.

After a [custom
element](#custom-element){#custom-elements-core-concepts:custom-element-2}
is
[created](https://dom.spec.whatwg.org/#concept-create-element){#custom-elements-core-concepts:create-an-element
x-internal="create-an-element"}, changing the value of the
[`is`{#custom-elements-core-concepts:attr-is}](#attr-is) attribute does
not change the element\'s behavior, as it is saved on the element as its
[`is`
value](https://dom.spec.whatwg.org/#concept-element-is-value){#custom-elements-core-concepts:concept-element-is-value
x-internal="concept-element-is-value"}.

[Autonomous custom
elements](#autonomous-custom-element){#custom-elements-core-concepts:autonomous-custom-element}
have the following element definition:

[Categories](dom.html#concept-element-categories){#custom-elements-core-concepts:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#custom-elements-core-concepts:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#custom-elements-core-concepts:phrasing-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#custom-elements-core-concepts:palpable-content-2}.
:   For [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element}:
    [Listed](forms.html#category-listed){#custom-elements-core-concepts:category-listed},
    [labelable](forms.html#category-label){#custom-elements-core-concepts:category-label},
    [submittable](forms.html#category-submit){#custom-elements-core-concepts:category-submit},
    and
    [resettable](forms.html#category-reset){#custom-elements-core-concepts:category-reset}
    [form-associated
    element](forms.html#form-associated-element){#custom-elements-core-concepts:form-associated-element}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#custom-elements-core-concepts:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#custom-elements-core-concepts:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#custom-elements-core-concepts:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#custom-elements-core-concepts:transparent}.

[Content attributes](dom.html#concept-element-attributes){#custom-elements-core-concepts:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#custom-elements-core-concepts:global-attributes},
    except the
    [`is`{#custom-elements-core-concepts:attr-is-2}](#attr-is) attribute
:   [`form`{#custom-elements-core-concepts:attr-fae-form}](form-control-infrastructure.html#attr-fae-form),
    for [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-2}
    --- Associates the element with a
    [`form`{#custom-elements-core-concepts:the-form-element}](forms.html#the-form-element)
    element
:   [`disabled`{#custom-elements-core-concepts:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled),
    for [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-3}
    --- Whether the form control is disabled
:   [`readonly`{#custom-elements-core-concepts:attr-face-readonly}](#attr-face-readonly),
    for [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-4}
    --- Affects
    [`willValidate`{#custom-elements-core-concepts:dom-elementinternals-willvalidate}](form-control-infrastructure.html#dom-elementinternals-willvalidate),
    plus any behavior added by the custom element author
:   [`name`{#custom-elements-core-concepts:attr-fe-name}](form-control-infrastructure.html#attr-fe-name),
    for [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-5}
    --- Name of the element to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#custom-elements-core-concepts:form-submission-2}
    and in the
    [`form.elements`{#custom-elements-core-concepts:dom-form-elements}](forms.html#dom-form-elements)
    API
:   Any other attribute that has no namespace (see prose).

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#custom-elements-core-concepts:concept-element-accessibility-considerations}:
:   For [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-6}:
    [for
    authors](https://w3c.github.io/html-aria/#el-form-associated-custom-element);
    [for
    implementers](https://w3c.github.io/html-aam/#el-form-associated-custom-element).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-autonomous-custom-element);
    [for
    implementers](https://w3c.github.io/html-aam/#el-autonomous-custom-element).

[DOM interface](dom.html#concept-element-dom){#custom-elements-core-concepts:concept-element-dom}:
:   Supplied by the element\'s author (inherits from
    [`HTMLElement`{#custom-elements-core-concepts:htmlelement}](dom.html#htmlelement))

An [autonomous custom
element](#autonomous-custom-element){#custom-elements-core-concepts:autonomous-custom-element-2}
does not have any special meaning: it
[represents](dom.html#represents){#custom-elements-core-concepts:represents}
its children. A [customized built-in
element](#customized-built-in-element){#custom-elements-core-concepts:customized-built-in-element}
inherits the semantics of the element that it extends.

Any namespace-less attribute that is relevant to the element\'s
functioning, as determined by the element\'s author, may be specified on
an [autonomous custom
element](#autonomous-custom-element){#custom-elements-core-concepts:autonomous-custom-element-3},
so long as the attribute name is
[XML-compatible](infrastructure.html#xml-compatible){#custom-elements-core-concepts:xml-compatible}
and contains no [ASCII upper
alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#custom-elements-core-concepts:uppercase-ascii-letters
x-internal="uppercase-ascii-letters"}. The exception is the
[`is`{#custom-elements-core-concepts:attr-is-3}](#attr-is) attribute,
which must not be specified on an [autonomous custom
element](#autonomous-custom-element){#custom-elements-core-concepts:autonomous-custom-element-4}
(and which will have no effect if it is).

[Customized built-in
elements](#customized-built-in-element){#custom-elements-core-concepts:customized-built-in-element-2}
follow the normal requirements for attributes, based on the elements
they extend. To add custom attribute-based behavior, use
[`data-*`{#custom-elements-core-concepts:attr-data-*}](dom.html#attr-data-*)
attributes.

------------------------------------------------------------------------

An [autonomous custom
element](#autonomous-custom-element){#custom-elements-core-concepts:autonomous-custom-element-5}
is called a [form-associated custom
element]{#form-associated-custom-element .dfn export=""} if the element
is associated with a [custom element
definition](#custom-element-definition){#custom-elements-core-concepts:custom-element-definition}
whose
[form-associated](#concept-custom-element-definition-form-associated){#custom-elements-core-concepts:concept-custom-element-definition-form-associated}
field is set to true.

The
[`name`{#custom-elements-core-concepts:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the [form-associated custom
element](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-7}\'s
name. The
[`disabled`{#custom-elements-core-concepts:attr-fe-disabled-2}](form-control-infrastructure.html#attr-fe-disabled)
attribute is used to make the [form-associated custom
element](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-8}
non-interactive and to prevent its [submission
value](#face-submission-value){#custom-elements-core-concepts:face-submission-value}
from being submitted. The
[`form`{#custom-elements-core-concepts:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the [form-associated custom
element](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-9}
with its [form
owner](form-control-infrastructure.html#form-owner){#custom-elements-core-concepts:form-owner}.

The [`readonly`]{#attr-face-readonly .dfn
dfn-for="form-associated custom elements" dfn-type="element-attr"}
attribute of [form-associated custom
elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-10}
specifies that the element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#custom-elements-core-concepts:barred-from-constraint-validation}.
User agents don\'t provide any other behavior for the attribute, but
custom element authors should, where possible, use its presence to make
their control non-editable in some appropriate fashion, similar to the
behavior for the
[readonly](input.html#attr-input-readonly){#custom-elements-core-concepts:attr-input-readonly}
attribute on built-in form controls.

**Constraint validation**: If the
[`readonly`{#custom-elements-core-concepts:attr-face-readonly-2}](#attr-face-readonly)
attribute is specified on a [form-associated custom
element](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-11},
the element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#custom-elements-core-concepts:barred-from-constraint-validation-2}.

The [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#custom-elements-core-concepts:concept-form-reset-control}
for [form-associated custom
elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-12}
is to [enqueue a custom element callback
reaction](#enqueue-a-custom-element-callback-reaction){#custom-elements-core-concepts:enqueue-a-custom-element-callback-reaction}
with the element, callback name \"`formResetCallback`\", and « ».

------------------------------------------------------------------------

A [valid custom element name]{#valid-custom-element-name .dfn export=""}
is a sequence of characters `name`{.variable} that meets all of the
following requirements:

- `name`{.variable} must match the
  [`PotentialCustomElementName`{#custom-elements-core-concepts:prod-potentialcustomelementname}](#prod-potentialcustomelementname)
  production:

  [`PotentialCustomElementName`]{#prod-potentialcustomelementname .dfn}` ::=`
  :   `[a-z] (`[`PCENChar`](#prod-pcenchar){#custom-elements-core-concepts:prod-pcenchar}`)* '-' (`[`PCENChar`](#prod-pcenchar){#custom-elements-core-concepts:prod-pcenchar-2}`)*`

  [`PCENChar`]{#prod-pcenchar .dfn}` ::=`
  :   `"-" | "." | [0-9] | "_" | [a-z] | #xB7 | [#xC0-#xD6] | [#xD8-#xF6] | [#xF8-#x37D] | [#x37F-#x1FFF] | [#x200C-#x200D] | [#x203F-#x2040] | [#x2070-#x218F] | [#x2C00-#x2FEF] | [#x3001-#xD7FF] | [#xF900-#xFDCF] | [#xFDF0-#xFFFD] | [#x10000-#xEFFFF]`

  This uses the [EBNF notation](https://www.w3.org/TR/xml/#sec-notation)
  from the XML specification. [\[XML\]](references.html#refsXML)

- `name`{.variable} must not be any of the following:

  - `annotation-xml`
  - `color-profile`
  - `font-face`
  - `font-face-src`
  - `font-face-uri`
  - `font-face-format`
  - `font-face-name`
  - `missing-glyph`

  The list of names above is the summary of all hyphen-containing
  element names from the [applicable
  specifications](infrastructure.html#other-applicable-specifications){#custom-elements-core-concepts:other-applicable-specifications},
  namely SVG 2 and MathML. [\[SVG\]](references.html#refsSVG)
  [\[MATHML\]](references.html#refsMATHML)

::: note
These requirements ensure a number of goals for [valid custom element
names](#valid-custom-element-name){#custom-elements-core-concepts:valid-custom-element-name-2}:

- They start with an [ASCII lower
  alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#custom-elements-core-concepts:lowercase-ascii-letters
  x-internal="lowercase-ascii-letters"}, ensuring that the HTML parser
  will treat them as tags instead of as text.

- They do not contain any [ASCII upper
  alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#custom-elements-core-concepts:uppercase-ascii-letters-2
  x-internal="uppercase-ascii-letters"}, ensuring that the user agent
  can always treat HTML elements ASCII-case-insensitively.

- They contain a hyphen, used for namespacing and to ensure forward
  compatibility (since no elements will be added to HTML, SVG, or MathML
  with hyphen-containing local names in the future).

- They can always be created with
  [`createElement()`{#custom-elements-core-concepts:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}
  and
  [`createElementNS()`{#custom-elements-core-concepts:dom-document-createelementns}](https://dom.spec.whatwg.org/#dom-document-createelementns){x-internal="dom-document-createelementns"},
  which have restrictions that go beyond the parser\'s.

Apart from these restrictions, a large variety of names is allowed, to
give maximum flexibility for use cases like `<math-α>` or
`<emotion-😍>`.
:::

A [custom element definition]{#custom-element-definition .dfn export=""}
describes a [custom
element](#custom-element){#custom-elements-core-concepts:custom-element-3}
and consists of:

A [name]{#concept-custom-element-definition-name .dfn dfn-for="custom element definition" export=""}
:   A [valid custom element
    name](#valid-custom-element-name){#custom-elements-core-concepts:valid-custom-element-name-3}

A [local name]{#concept-custom-element-definition-local-name .dfn dfn-for="custom element definition" export=""}
:   A local name

A [constructor]{#concept-custom-element-definition-constructor .dfn dfn-for="custom element definition" export=""}
:   A Web IDL
    [`CustomElementConstructor`{#custom-elements-core-concepts:customelementconstructor}](#customelementconstructor)
    callback function type value wrapping the [custom element
    constructor](#custom-element-constructor){#custom-elements-core-concepts:custom-element-constructor}

A list of [observed attributes]{#concept-custom-element-definition-observed-attributes .dfn}
:   A `sequence<DOMString>`

A collection of [lifecycle callbacks]{#concept-custom-element-definition-lifecycle-callbacks .dfn}
:   A map, whose keys are the strings \"`connectedCallback`\",
    \"`disconnectedCallback`\", \"`adoptedCallback`\",
    \"`connectedMoveCallback`\", \"`attributeChangedCallback`\",
    \"`formAssociatedCallback`\", \"`formDisabledCallback`\",
    \"`formResetCallback`\", and \"`formStateRestoreCallback`\". The
    corresponding values are either a Web IDL
    [`Function`{#custom-elements-core-concepts:idl-function}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
    callback function type value, or null. By default the value of each
    entry is null.

A [construction stack]{#concept-custom-element-definition-construction-stack .dfn}
:   A list, initially empty, that is manipulated by the [upgrade an
    element](#concept-upgrade-an-element){#custom-elements-core-concepts:concept-upgrade-an-element}
    algorithm and the [HTML element
    constructors](dom.html#html-element-constructors). Each entry in the
    list will be either an element or an [*already constructed*
    marker]{#concept-already-constructed-marker .dfn}.

A [form-associated]{#concept-custom-element-definition-form-associated .dfn} boolean
:   If this is true, user agent treats elements associated to this
    [custom element
    definition](#custom-element-definition){#custom-elements-core-concepts:custom-element-definition-2}
    as [form-associated custom
    elements](#form-associated-custom-element){#custom-elements-core-concepts:form-associated-custom-element-13}.

A [disable internals]{#concept-custom-element-definition-disable-internals .dfn} boolean
:   Controls
    [`attachInternals()`{#custom-elements-core-concepts:dom-attachinternals}](#dom-attachinternals).

A [disable shadow]{#concept-custom-element-definition-disable-shadow .dfn dfn-for="custom element definition" export=""} boolean
:   Controls
    [`attachShadow()`{#custom-elements-core-concepts:dom-element-attachshadow}](https://dom.spec.whatwg.org/#dom-element-attachshadow){x-internal="dom-element-attachshadow"}.

To [look up a custom element
definition]{#look-up-a-custom-element-definition .dfn export=""}, given
null or a
[`CustomElementRegistry`{#custom-elements-core-concepts:customelementregistry}](#customelementregistry)
object `registry`{.variable}, string-or-null `namespace`{.variable},
string `localName`{.variable}, and string-or-null `is`{.variable},
perform the following steps. They will return either a [custom element
definition](#custom-element-definition){#custom-elements-core-concepts:custom-element-definition-3}
or null:

1.  If `registry`{.variable} is null, then return null.

2.  If `namespace`{.variable} is not the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#custom-elements-core-concepts:html-namespace-2
    x-internal="html-namespace-2"}, then return null.

3.  If `registry`{.variable}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-core-concepts:custom-element-definition-set}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-core-concepts:list-contains
    x-internal="list-contains"} an item with
    [name](#concept-custom-element-definition-name){#custom-elements-core-concepts:concept-custom-element-definition-name-3}
    and [local
    name](#concept-custom-element-definition-local-name){#custom-elements-core-concepts:concept-custom-element-definition-local-name}
    both equal to `localName`{.variable}, then return that item.

4.  If `registry`{.variable}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-core-concepts:custom-element-definition-set-2}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-core-concepts:list-contains-2
    x-internal="list-contains"} an item with
    [name](#concept-custom-element-definition-name){#custom-elements-core-concepts:concept-custom-element-definition-name-4}
    equal to `is`{.variable} and [local
    name](#concept-custom-element-definition-local-name){#custom-elements-core-concepts:concept-custom-element-definition-local-name-2}
    equal to `localName`{.variable}, then return that item.

5.  Return null.

#### [4.13.4]{.secno} The [`CustomElementRegistry`{#custom-elements-api:customelementregistry}](#customelementregistry) interface[](#custom-elements-api){.self-link} {#custom-elements-api}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[CustomElementRegistry](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry "The CustomElementRegistry interface provides methods for registering custom elements and querying registered elements. To get an instance of it, use the window.customElements property.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Each [similar-origin window
agent](webappapis.html#similar-origin-window-agent){#custom-elements-api:similar-origin-window-agent}
has an associated [active custom element constructor
map]{#active-custom-element-constructor-map .dfn dfn-for="similar-origin
  window agent" export=""}, which is a
[map](https://infra.spec.whatwg.org/#ordered-map){#custom-elements-api:ordered-map
x-internal="ordered-map"} of constructors to
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-2}](#customelementregistry)
objects.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/customElements](https://developer.mozilla.org/en-US/docs/Web/API/Window/customElements "The customElements read-only property of the Window interface returns a reference to the CustomElementRegistry object, which can be used to register new custom elements and get information about previously registered custom elements.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`Window`{#custom-elements-api:window}](nav-history-apis.html#window)
[`customElements`]{#dom-window-customelements .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#custom-elements-api:assert
    x-internal="assert"}:
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#custom-elements-api:concept-document-window}\'s
    [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#custom-elements-api:document-custom-element-registry
    x-internal="document-custom-element-registry"} is a
    [`CustomElementRegistry`{#custom-elements-api:customelementregistry-3}](#customelementregistry)
    object.

    A
    [`Window`{#custom-elements-api:window-2}](nav-history-apis.html#window)\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#custom-elements-api:concept-document-window-2}
    is always created with a new
    [`CustomElementRegistry`{#custom-elements-api:customelementregistry-4}](#customelementregistry)
    object.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-2
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#custom-elements-api:concept-document-window-3}\'s
    [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#custom-elements-api:document-custom-element-registry-2
    x-internal="document-custom-element-registry"}.

``` idl
[Exposed=Window]
interface CustomElementRegistry {
  constructor();

  [CEReactions] undefined define(DOMString name, CustomElementConstructor constructor, optional ElementDefinitionOptions options = {});
  (CustomElementConstructor or undefined) get(DOMString name);
  DOMString? getName(CustomElementConstructor constructor);
  Promise<CustomElementConstructor> whenDefined(DOMString name);
  [CEReactions] undefined upgrade(Node root);
  undefined initialize(Node root);
};

callback CustomElementConstructor = HTMLElement ();

dictionary ElementDefinitionOptions {
  DOMString extends;
};
```

Every
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-5}](#customelementregistry)
has an [is scoped]{#is-scoped .dfn dfn-for="CustomElementRegistry"
export=""}, a boolean, initially false.

Every
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-6}](#customelementregistry)
has a [scoped document set]{#scoped-document-set .dfn
dfn-for="CustomElementRegistry" export=""}, a
[set](https://infra.spec.whatwg.org/#ordered-set){#custom-elements-api:set
x-internal="set"} of
[`Document`{#custom-elements-api:document}](dom.html#document) objects,
initially « ».

Every
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-7}](#customelementregistry)
has a [custom element definition set]{#custom-element-definition-set
.dfn}, a
[set](https://infra.spec.whatwg.org/#ordered-set){#custom-elements-api:set-2
x-internal="set"} of [custom element
definitions](#custom-element-definition){#custom-elements-api:custom-element-definition},
initially « ». Lookup of items in this
[set](https://infra.spec.whatwg.org/#ordered-set){#custom-elements-api:set-3
x-internal="set"} uses their
[name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name},
[local
name](#concept-custom-element-definition-local-name){#custom-elements-api:concept-custom-element-definition-local-name},
or
[constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor}.

Every
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-8}](#customelementregistry)
also has an [element definition is
running]{#element-definition-is-running .dfn} boolean which is used to
prevent reentrant invocations of [element
definition](#element-definition){#custom-elements-api:element-definition}.
It is initially false.

Every
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-9}](#customelementregistry)
also has a [when-defined promise map]{#when-defined-promise-map .dfn}, a
[map](https://infra.spec.whatwg.org/#ordered-map){#custom-elements-api:ordered-map-2
x-internal="ordered-map"} of [valid custom element
names](#valid-custom-element-name){#custom-elements-api:valid-custom-element-name}
to promises. It is used to implement the
[`whenDefined()`{#custom-elements-api:dom-customelementregistry-whendefined-2}](#dom-customelementregistry-whendefined)
method.

To [look up a custom element
registry]{#look-up-a-custom-element-registry .dfn export=""}, given a
[`Node`{#custom-elements-api:node-3}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
object `node`{.variable}:

1.  If `node`{.variable} is an
    [`Element`{#custom-elements-api:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    object, then return `node`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#custom-elements-api:element-custom-element-registry
    x-internal="element-custom-element-registry"}.

2.  If `node`{.variable} is a
    [`ShadowRoot`{#custom-elements-api:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}
    object, then return `node`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#shadowroot-custom-element-registry){#custom-elements-api:shadow-root-custom-element-registry
    x-internal="shadow-root-custom-element-registry"}.

3.  If `node`{.variable} is a
    [`Document`{#custom-elements-api:document-2}](dom.html#document)
    object, then return `node`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#custom-elements-api:document-custom-element-registry-3
    x-internal="document-custom-element-registry"}.

4.  Return null.

`registry`{.variable}` = ``window`{.variable}`.`[`customElements`](#dom-window-customelements){#dom-window-customelements-dev}

Returns the global\'s associated
[`Document`{#custom-elements-api:document-3}](dom.html#document)\'s
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-10}](#customelementregistry)
object.

`registry`{.variable}` = new `[`CustomElementRegistry`](#dom-customelementregistry){#dom-customelementregistry-dev}`()`

Constructs a new
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-11}](#customelementregistry)
object, for scoped usage.

`registry`{.variable}`.`[`define`](#dom-customelementregistry-define){#dom-customelementregistry-define-dev}`(``name`{.variable}`, ``constructor`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CustomElementRegistry/define](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry/define "The define() method of the CustomElementRegistry interface defines a new custom element.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Defines a new [custom
element](#custom-element){#custom-elements-api:custom-element}, mapping
the given name to the given constructor as an [autonomous custom
element](#autonomous-custom-element){#custom-elements-api:autonomous-custom-element}.

`registry`{.variable}`.`[`define`](#dom-customelementregistry-define){#custom-elements-api:dom-customelementregistry-define-2}`(``name`{.variable}`, ``constructor`{.variable}`, { extends: ``baseLocalName`{.variable}` })`

Defines a new [custom
element](#custom-element){#custom-elements-api:custom-element-2},
mapping the given name to the given constructor as a [customized
built-in
element](#customized-built-in-element){#custom-elements-api:customized-built-in-element}
for the [element
type](infrastructure.html#element-type){#custom-elements-api:element-type}
identified by the supplied `baseLocalName`{.variable}. A
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror
x-internal="notsupportederror"}
[`DOMException`{#custom-elements-api:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
will be thrown upon trying to extend a [custom
element](#custom-element){#custom-elements-api:custom-element-3} or an
unknown element, or when `registry`{.variable} is not a global
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-12}](#customelementregistry)
object.

`registry`{.variable}`.`[`get`](#dom-customelementregistry-get){#dom-customelementregistry-get-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CustomElementRegistry/get](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry/get "The get() method of the CustomElementRegistry interface returns the constructor for a previously-defined custom element.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Retrieves the [custom element
constructor](#custom-element-constructor){#custom-elements-api:custom-element-constructor}
defined for the given
[name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-2}.
Returns undefined if there is no [custom element
definition](#custom-element-definition){#custom-elements-api:custom-element-definition-2}
with the given
[name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-3}.

`registry`{.variable}`.`[`getName`](#dom-customelementregistry-getname){#dom-customelementregistry-getname-dev}`(``constructor`{.variable}`)`

Retrieves the given name for a [custom
element](#custom-element){#custom-elements-api:custom-element-4} defined
for the given
[constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-2}.
Returns null if there is no [custom element
definition](#custom-element-definition){#custom-elements-api:custom-element-definition-3}
with the given
[constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-3}.

`registry`{.variable}`.`[`whenDefined`](#dom-customelementregistry-whendefined){#dom-customelementregistry-whendefined-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CustomElementRegistry/whenDefined](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry/whenDefined "The whenDefined() method of the CustomElementRegistry interface returns a Promise that resolves when the named element is defined.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome54+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns a promise that will be fulfilled with the [custom
element](#custom-element){#custom-elements-api:custom-element-5}\'s
constructor when a [custom
element](#custom-element){#custom-elements-api:custom-element-6} becomes
defined with the given name. (If such a [custom
element](#custom-element){#custom-elements-api:custom-element-7} is
already defined, the returned promise will be immediately fulfilled.)
Returns a promise rejected with a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-elements-api:syntaxerror
x-internal="syntaxerror"}
[`DOMException`{#custom-elements-api:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if not given a [valid custom element
name](#valid-custom-element-name){#custom-elements-api:valid-custom-element-name-2}.

`registry`{.variable}`.`[`upgrade`](#dom-customelementregistry-upgrade){#dom-customelementregistry-upgrade-dev}`(``root`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CustomElementRegistry/upgrade](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry/upgrade "The upgrade() method of the CustomElementRegistry interface upgrades all shadow-containing custom elements in a Node subtree, even before they are connected to the main document.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome68+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Tries to
upgrade](#concept-try-upgrade){#custom-elements-api:concept-try-upgrade}
all [shadow-including inclusive
descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#custom-elements-api:shadow-including-inclusive-descendant
x-internal="shadow-including-inclusive-descendant"} elements of
`root`{.variable}, even if they are not
[connected](https://dom.spec.whatwg.org/#connected){#custom-elements-api:connected
x-internal="connected"}.

`registry`{.variable}`.`[`initialize`](#dom-customelementregistry-initialize){#dom-customelementregistry-initialize-dev}`(``root`{.variable}`)`

Each [inclusive
descendant](https://dom.spec.whatwg.org/#concept-tree-inclusive-descendant){#custom-elements-api:inclusive-descendant
x-internal="inclusive-descendant"} of `root`{.variable} with a null
registry will have it updated to this
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-13}](#customelementregistry)
object.

The [`new CustomElementRegistry()`]{#dom-customelementregistry .dfn}
constructor steps are to set
[this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-3
x-internal="this"}\'s [is
scoped](#is-scoped){#custom-elements-api:is-scoped} to true.

[Element definition]{#element-definition .dfn} is a process of adding a
[custom element
definition](#custom-element-definition){#custom-elements-api:custom-element-definition-4}
to the
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-14}](#customelementregistry).
This is accomplished by the
[`define()`{#custom-elements-api:dom-customelementregistry-define-3}](#dom-customelementregistry-define)
method. The
[`define(``name`{.variable}`, ``constructor`{.variable}`, ``options`{.variable}`)`]{#dom-customelementregistry-define
.dfn dfn-for="CustomElementRegistry" dfn-type="method"} method steps
are:

1.  If
    [IsConstructor](https://tc39.es/ecma262/#sec-isconstructor){#custom-elements-api:isconstructor
    x-internal="isconstructor"}(`constructor`{.variable}) is false, then
    throw a
    [`TypeError`{#custom-elements-api:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

2.  If `name`{.variable} is not a [valid custom element
    name](#valid-custom-element-name){#custom-elements-api:valid-custom-element-name-3},
    then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-elements-api:syntaxerror-2
    x-internal="syntaxerror"}
    [`DOMException`{#custom-elements-api:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-4
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains
    x-internal="list-contains"} an item with
    [name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-4}
    `name`{.variable}, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-2
    x-internal="notsupportederror"}
    [`DOMException`{#custom-elements-api:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-5
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set-2}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-2
    x-internal="list-contains"} an item with
    [constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-4}
    `constructor`{.variable}, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-3
    x-internal="notsupportederror"}
    [`DOMException`{#custom-elements-api:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  Let `localName`{.variable} be `name`{.variable}.

6.  Let `extends`{.variable} be
    `options`{.variable}\[\"[`extends`{#custom-elements-api:dom-elementdefinitionoptions-extends}](#dom-elementdefinitionoptions-extends)\"\]
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#custom-elements-api:map-exists
    x-internal="map-exists"}; otherwise null.

7.  If `extends`{.variable} is not null:

    1.  If
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-6
        x-internal="this"}\'s [is
        scoped](#is-scoped){#custom-elements-api:is-scoped-2} is true,
        then throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-4
        x-internal="notsupportederror"}
        [`DOMException`{#custom-elements-api:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  If `extends`{.variable} is a [valid custom element
        name](#valid-custom-element-name){#custom-elements-api:valid-custom-element-name-4},
        then throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-5
        x-internal="notsupportederror"}
        [`DOMException`{#custom-elements-api:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  If the [element
        interface](https://dom.spec.whatwg.org/#concept-element-interface){#custom-elements-api:element-interface
        x-internal="element-interface"} for `extends`{.variable} and the
        [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#custom-elements-api:html-namespace-2
        x-internal="html-namespace-2"} is
        [`HTMLUnknownElement`{#custom-elements-api:htmlunknownelement}](dom.html#htmlunknownelement)
        (e.g., if `extends`{.variable} does not indicate an element
        definition in this specification), then throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-6
        x-internal="notsupportederror"}
        [`DOMException`{#custom-elements-api:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    4.  Set `localName`{.variable} to `extends`{.variable}.

8.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-7
    x-internal="this"}\'s [element definition is
    running](#element-definition-is-running){#custom-elements-api:element-definition-is-running}
    is true, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#custom-elements-api:notsupportederror-7
    x-internal="notsupportederror"}
    [`DOMException`{#custom-elements-api:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

9.  Set
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-8
    x-internal="this"}\'s [element definition is
    running](#element-definition-is-running){#custom-elements-api:element-definition-is-running-2}
    to true.

10. Let `formAssociated`{.variable} be false.

11. Let `disableInternals`{.variable} be false.

12. Let `disableShadow`{.variable} be false.

13. Let `observedAttributes`{.variable} be an empty
    `sequence<DOMString>`.

14. Run the following steps while catching any exceptions:

    1.  Let `prototype`{.variable} be ?
        [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get
        x-internal="js-get"}(`constructor`{.variable}, \"prototype\").

    2.  If `prototype`{.variable} [is not an
        Object](https://tc39.es/ecma262/#sec-object-type){#custom-elements-api:js-object
        x-internal="js-object"}, then throw a
        [`TypeError`{#custom-elements-api:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
        exception.

    3.  Let `lifecycleCallbacks`{.variable} be the [ordered
        map](https://infra.spec.whatwg.org/#ordered-map){#custom-elements-api:ordered-map-3
        x-internal="ordered-map"} «\[ \"`connectedCallback`\" → null,
        \"`disconnectedCallback`\" → null, \"`adoptedCallback`\" → null,
        \"`connectedMoveCallback`\" → null,
        \"`attributeChangedCallback`\" → null \]».

    4.  For each `callbackName`{.variable} of [the
        keys](https://infra.spec.whatwg.org/#map-getting-the-keys){#custom-elements-api:map-get-the-keys
        x-internal="map-get-the-keys"} of
        `lifecycleCallbacks`{.variable}:

        1.  Let `callbackValue`{.variable} be ?
            [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get-2
            x-internal="js-get"}(`prototype`{.variable},
            `callbackName`{.variable}).

        2.  If `callbackValue`{.variable} is not undefined, then
            [set](https://infra.spec.whatwg.org/#map-set){#custom-elements-api:map-set
            x-internal="map-set"}
            `lifecycleCallbacks`{.variable}\[`callbackName`{.variable}\]
            to the result of
            [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#custom-elements-api:concept-idl-convert
            x-internal="concept-idl-convert"} `callbackValue`{.variable}
            to the Web IDL
            [`Function`{#custom-elements-api:idl-function}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
            callback type.

    5.  If
        `lifecycleCallbacks`{.variable}\[\"`attributeChangedCallback`\"\]
        is not null:

        1.  Let `observedAttributesIterable`{.variable} be ?
            [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get-3
            x-internal="js-get"}(`constructor`{.variable},
            \"observedAttributes\").

        2.  If `observedAttributesIterable`{.variable} is not undefined,
            then set `observedAttributes`{.variable} to the result of
            [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#custom-elements-api:concept-idl-convert-2
            x-internal="concept-idl-convert"}
            `observedAttributesIterable`{.variable} to a
            `sequence<DOMString>`. Rethrow any exceptions from the
            conversion.

    6.  Let `disabledFeatures`{.variable} be an empty
        `sequence<DOMString>`.

    7.  Let `disabledFeaturesIterable`{.variable} be ?
        [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get-4
        x-internal="js-get"}(`constructor`{.variable},
        \"disabledFeatures\").

    8.  If `disabledFeaturesIterable`{.variable} is not undefined, then
        set `disabledFeatures`{.variable} to the result of
        [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#custom-elements-api:concept-idl-convert-3
        x-internal="concept-idl-convert"}
        `disabledFeaturesIterable`{.variable} to a
        `sequence<DOMString>`. Rethrow any exceptions from the
        conversion.

    9.  If `disabledFeatures`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-3
        x-internal="list-contains"} \"`internals`\", then set
        `disableInternals`{.variable} to true.

    10. If `disabledFeatures`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-4
        x-internal="list-contains"} \"`shadow`\", then set
        `disableShadow`{.variable} to true.

    11. Let `formAssociatedValue`{.variable} be ?
        [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get-5
        x-internal="js-get"}( `constructor`{.variable},
        \"formAssociated\").

    12. Set `formAssociated`{.variable} to the result of
        [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#custom-elements-api:concept-idl-convert-4
        x-internal="concept-idl-convert"}
        `formAssociatedValue`{.variable} to a `boolean`.

    13. If `formAssociated`{.variable} is true, then for each
        `callbackName`{.variable} of « \"`formAssociatedCallback`\",
        \"`formResetCallback`\", \"`formDisabledCallback`\",
        \"`formStateRestoreCallback`\" »:

        1.  Let `callbackValue`{.variable} be ?
            [Get](https://tc39.es/ecma262/#sec-get-o-p){#custom-elements-api:js-get-6
            x-internal="js-get"}(`prototype`{.variable},
            `callbackName`{.variable}).

        2.  If `callbackValue`{.variable} is not undefined, then
            [set](https://infra.spec.whatwg.org/#map-set){#custom-elements-api:map-set-2
            x-internal="map-set"}
            `lifecycleCallbacks`{.variable}\[`callbackName`{.variable}\]
            to the result of
            [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#custom-elements-api:concept-idl-convert-5
            x-internal="concept-idl-convert"} `callbackValue`{.variable}
            to the Web IDL
            [`Function`{#custom-elements-api:idl-function-2}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
            callback type.

    Then, regardless of whether the above steps threw an exception or
    not: set
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-9
    x-internal="this"}\'s [element definition is
    running](#element-definition-is-running){#custom-elements-api:element-definition-is-running-3}
    to false.

    Finally, if the steps threw an exception, rethrow that exception.

15. Let `definition`{.variable} be a new [custom element
    definition](#custom-element-definition){#custom-elements-api:custom-element-definition-5}
    with
    [name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-5}
    `name`{.variable}, [local
    name](#concept-custom-element-definition-local-name){#custom-elements-api:concept-custom-element-definition-local-name-2}
    `localName`{.variable},
    [constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-5}
    `constructor`{.variable}, [observed
    attributes](#concept-custom-element-definition-observed-attributes){#custom-elements-api:concept-custom-element-definition-observed-attributes}
    `observedAttributes`{.variable}, [lifecycle
    callbacks](#concept-custom-element-definition-lifecycle-callbacks){#custom-elements-api:concept-custom-element-definition-lifecycle-callbacks}
    `lifecycleCallbacks`{.variable},
    [form-associated](#concept-custom-element-definition-form-associated){#custom-elements-api:concept-custom-element-definition-form-associated}
    `formAssociated`{.variable}, [disable
    internals](#concept-custom-element-definition-disable-internals){#custom-elements-api:concept-custom-element-definition-disable-internals}
    `disableInternals`{.variable}, and [disable
    shadow](#concept-custom-element-definition-disable-shadow){#custom-elements-api:concept-custom-element-definition-disable-shadow}
    `disableShadow`{.variable}.

16. [Append](https://infra.spec.whatwg.org/#set-append){#custom-elements-api:set-append
    x-internal="set-append"} `definition`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-10
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set-3}.

17. If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-11
    x-internal="this"}\'s [is
    scoped](#is-scoped){#custom-elements-api:is-scoped-3} is true, then
    for each `document`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-12
    x-internal="this"}\'s [scoped document
    set](#scoped-document-set){#custom-elements-api:scoped-document-set}:
    [upgrade particular elements within a
    document](#upgrade-particular-elements-within-a-document){#custom-elements-api:upgrade-particular-elements-within-a-document}
    given
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-13
    x-internal="this"}, `document`{.variable}, `definition`{.variable},
    and `localName`{.variable}.

18. Otherwise, [upgrade particular elements within a
    document](#upgrade-particular-elements-within-a-document){#custom-elements-api:upgrade-particular-elements-within-a-document-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-14
    x-internal="this"},
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-15
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#custom-elements-api:concept-relevant-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#custom-elements-api:concept-document-window-4},
    `definition`{.variable}, `localName`{.variable}, and
    `name`{.variable}.

19. If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-16
    x-internal="this"}\'s [when-defined promise
    map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map}\[`name`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#custom-elements-api:map-exists-2
    x-internal="map-exists"}:

    1.  Resolve
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-17
        x-internal="this"}\'s [when-defined promise
        map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map-2}\[`name`{.variable}\]
        with `constructor`{.variable}.

    2.  [Remove](https://infra.spec.whatwg.org/#map-remove){#custom-elements-api:map-remove
        x-internal="map-remove"}
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-18
        x-internal="this"}\'s [when-defined promise
        map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map-3}\[`name`{.variable}\].

To [upgrade particular elements within a
document]{#upgrade-particular-elements-within-a-document .dfn} given a
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-15}](#customelementregistry)
object `registry`{.variable}, a
[`Document`{#custom-elements-api:document-4}](dom.html#document) object
`document`{.variable}, a [custom element
definition](#custom-element-definition){#custom-elements-api:custom-element-definition-6}
`definition`{.variable}, a string `localName`{.variable}, and optionally
a string `name`{.variable} (default `localName`{.variable}):

1.  Let `upgradeCandidates`{.variable} be all elements that are
    [shadow-including
    descendants](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#custom-elements-api:shadow-including-descendant
    x-internal="shadow-including-descendant"} of `document`{.variable},
    whose [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#custom-elements-api:element-custom-element-registry-2
    x-internal="element-custom-element-registry"} is
    `registry`{.variable}, whose namespace is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#custom-elements-api:html-namespace-2-2
    x-internal="html-namespace-2"}, and whose local name is
    `localName`{.variable}, in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#custom-elements-api:shadow-including-tree-order
    x-internal="shadow-including-tree-order"}. Additionally, if
    `name`{.variable} is not `localName`{.variable}, only include
    elements whose [`is`
    value](https://dom.spec.whatwg.org/#concept-element-is-value){#custom-elements-api:concept-element-is-value
    x-internal="concept-element-is-value"} is equal to
    `name`{.variable}.

2.  For each element `element`{.variable} of
    `upgradeCandidates`{.variable}: [enqueue a custom element upgrade
    reaction](#enqueue-a-custom-element-upgrade-reaction){#custom-elements-api:enqueue-a-custom-element-upgrade-reaction}
    given `element`{.variable} and `definition`{.variable}.

The [`get(``name`{.variable}`)`]{#dom-customelementregistry-get .dfn
dfn-for="CustomElementRegistry" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-19
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set-4}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-5
    x-internal="list-contains"} an item with
    [name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-6}
    `name`{.variable}, then return that item\'s
    [constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-6}.

2.  Return undefined.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[CustomElementRegistry/getName](https://developer.mozilla.org/en-US/docs/Web/API/CustomElementRegistry/getName "The getName() method of the CustomElementRegistry interface returns the name for a previously-defined custom element.")

::: support
[Firefox116+]{.firefox .yes}[Safari[🔰
preview+]{title="Partial implementation."}]{.safari
.yes}[Chrome117+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge117+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`getName(``constructor`{.variable}`)`]{#dom-customelementregistry-getname
.dfn dfn-for="CustomElementRegistry" dfn-type="method"} method steps
are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-20
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set-5}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-6
    x-internal="list-contains"} an item with
    [constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-7}
    `constructor`{.variable}, then return that item\'s
    [name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-7}.

2.  Return null.

The
[`whenDefined(``name`{.variable}`)`]{#dom-customelementregistry-whendefined
.dfn dfn-for="CustomElementRegistry" dfn-type="method"} method steps
are:

1.  If `name`{.variable} is not a [valid custom element
    name](#valid-custom-element-name){#custom-elements-api:valid-custom-element-name-5},
    then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#custom-elements-api:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-elements-api:syntaxerror-3
    x-internal="syntaxerror"}
    [`DOMException`{#custom-elements-api:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-21
    x-internal="this"}\'s [custom element definition
    set](#custom-element-definition-set){#custom-elements-api:custom-element-definition-set-6}
    [contains](https://infra.spec.whatwg.org/#list-contain){#custom-elements-api:list-contains-7
    x-internal="list-contains"} an item with
    [name](#concept-custom-element-definition-name){#custom-elements-api:concept-custom-element-definition-name-8}
    `name`{.variable}, then return [a promise resolved
    with](https://webidl.spec.whatwg.org/#a-promise-resolved-with){#custom-elements-api:a-promise-resolved-with
    x-internal="a-promise-resolved-with"} that item\'s
    [constructor](#concept-custom-element-definition-constructor){#custom-elements-api:concept-custom-element-definition-constructor-8}.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-22
    x-internal="this"}\'s [when-defined promise
    map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map-4}\[`name`{.variable}\]
    does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#custom-elements-api:map-exists-3
    x-internal="map-exists"}, then
    [set](https://infra.spec.whatwg.org/#map-set){#custom-elements-api:map-set-3
    x-internal="map-set"}
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-23
    x-internal="this"}\'s [when-defined promise
    map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map-5}\[`name`{.variable}\]
    to a new promise.

4.  Return
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-24
    x-internal="this"}\'s [when-defined promise
    map](#when-defined-promise-map){#custom-elements-api:when-defined-promise-map-6}\[`name`{.variable}\].

::: example
The
[`whenDefined()`{#custom-elements-api:dom-customelementregistry-whendefined-3}](#dom-customelementregistry-whendefined)
method can be used to avoid performing an action until all appropriate
[custom
elements](#custom-element){#custom-elements-api:custom-element-8} are
[defined](https://dom.spec.whatwg.org/#concept-element-defined){#custom-elements-api:concept-element-defined
x-internal="concept-element-defined"}. In this example, we combine it
with the
[`:defined`{#custom-elements-api:selector-defined}](semantics-other.html#selector-defined)
pseudo-class to hide a dynamically-loaded article\'s contents until
we\'re sure that all of the [autonomous custom
elements](#autonomous-custom-element){#custom-elements-api:autonomous-custom-element-2}
it uses are defined.

``` js
articleContainer.hidden = true;

fetch(articleURL)
  .then(response => response.text())
  .then(text => {
    articleContainer.innerHTML = text;

    return Promise.all(
      [...articleContainer.querySelectorAll(":not(:defined)")]
        .map(el => customElements.whenDefined(el.localName))
    );
  })
  .then(() => {
    articleContainer.hidden = false;
  });
```
:::

The [`upgrade(``root`{.variable}`)`]{#dom-customelementregistry-upgrade
.dfn dfn-for="CustomElementRegistry" dfn-type="method"} method steps
are:

1.  Let `candidates`{.variable} be a
    [list](https://infra.spec.whatwg.org/#list){#custom-elements-api:list
    x-internal="list"} of all of `root`{.variable}\'s [shadow-including
    inclusive
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#custom-elements-api:shadow-including-inclusive-descendant-2
    x-internal="shadow-including-inclusive-descendant"} elements, in
    [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#custom-elements-api:shadow-including-tree-order-2
    x-internal="shadow-including-tree-order"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#custom-elements-api:list-iterate
    x-internal="list-iterate"} `candidate`{.variable} of
    `candidates`{.variable}, [try to
    upgrade](#concept-try-upgrade){#custom-elements-api:concept-try-upgrade-2}
    `candidate`{.variable}.

::: example
The
[`upgrade()`{#custom-elements-api:dom-customelementregistry-upgrade-2}](#dom-customelementregistry-upgrade)
method allows upgrading of elements at will. Normally elements are
automatically upgraded when they become
[connected](https://dom.spec.whatwg.org/#connected){#custom-elements-api:connected-2
x-internal="connected"}, but this method can be used if you need to
upgrade before you\'re ready to connect the element.

``` js
const el = document.createElement("spider-man");

class SpiderMan extends HTMLElement {}
customElements.define("spider-man", SpiderMan);

console.assert(!(el instanceof SpiderMan)); // not yet upgraded

customElements.upgrade(el);
console.assert(el instanceof SpiderMan);    // upgraded!
```
:::

The
[`initialize(``root`{.variable}`)`]{#dom-customelementregistry-initialize
.dfn dfn-for="CustomElementRegistry" dfn-type="method"} method steps
are:

1.  If `root`{.variable} is a
    [`Document`{#custom-elements-api:document-5}](dom.html#document)
    node whose [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#custom-elements-api:document-custom-element-registry-4
    x-internal="document-custom-element-registry"} is null, then set
    `root`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#custom-elements-api:document-custom-element-registry-5
    x-internal="document-custom-element-registry"} to
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-25
    x-internal="this"}.

2.  Otherwise, if `root`{.variable} is a
    [`ShadowRoot`{#custom-elements-api:shadowroot-2}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}
    node whose [custom element
    registry](https://dom.spec.whatwg.org/#shadowroot-custom-element-registry){#custom-elements-api:shadow-root-custom-element-registry-2
    x-internal="shadow-root-custom-element-registry"} is null, then set
    `root`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#shadowroot-custom-element-registry){#custom-elements-api:shadow-root-custom-element-registry-3
    x-internal="shadow-root-custom-element-registry"} to
    [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-26
    x-internal="this"}.

3.  For each [inclusive
    descendant](https://dom.spec.whatwg.org/#concept-tree-inclusive-descendant){#custom-elements-api:inclusive-descendant-2
    x-internal="inclusive-descendant"} `inclusiveDescendant`{.variable}
    of `root`{.variable}: if `inclusiveDescendant`{.variable} is an
    [`Element`{#custom-elements-api:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    node whose [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#custom-elements-api:element-custom-element-registry-3
    x-internal="element-custom-element-registry"} is null:

    1.  Set `inclusiveDescendant`{.variable}\'s [custom element
        registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#custom-elements-api:element-custom-element-registry-4
        x-internal="element-custom-element-registry"} to
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-27
        x-internal="this"}.

    2.  If
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-28
        x-internal="this"}\'s [is
        scoped](#is-scoped){#custom-elements-api:is-scoped-4} is true,
        then
        [append](https://infra.spec.whatwg.org/#set-append){#custom-elements-api:set-append-2
        x-internal="set-append"} `inclusiveDescendant`{.variable}\'s
        [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#custom-elements-api:node-document
        x-internal="node-document"} to
        [this](https://webidl.spec.whatwg.org/#this){#custom-elements-api:this-29
        x-internal="this"}\'s [scoped document
        set](#scoped-document-set){#custom-elements-api:scoped-document-set-2}.

Once the custom element registry of a node is initialized to a
[`CustomElementRegistry`{#custom-elements-api:customelementregistry-16}](#customelementregistry)
object, it intentionally cannot be changed any further. This simplifies
reasoning about code and allows implementations to optimize.

#### [4.13.5]{.secno} [Upgrades]{.dfn}[](#upgrades){.self-link}

To [upgrade an element]{#concept-upgrade-an-element .dfn export=""},
given as input a [custom element
definition](#custom-element-definition){#upgrades:custom-element-definition}
`definition`{.variable} and an element `element`{.variable}, run the
following steps:

1.  :::: {#concept-upgrade-an-element-early-exit}
    If `element`{.variable}\'s [custom element
    state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#upgrades:custom-element-state
    x-internal="custom-element-state"} is not \"`undefined`\" or
    \"`uncustomized`\", then return.

    ::: example
    One scenario where this can occur due to reentrant invocation of
    this algorithm, as in the following example:

    ``` html
    <!DOCTYPE html>
    <x-foo id="a"></x-foo>
    <x-foo id="b"></x-foo>

    <script>
    // Defining enqueues upgrade reactions for both "a" and "b"
    customElements.define("x-foo", class extends HTMLElement {
      constructor() {
        super();

        const b = document.querySelector("#b");
        b.remove();

        // While this constructor is running for "a", "b" is still
        // undefined, and so inserting it into the document will enqueue a
        // second upgrade reaction for "b" in addition to the one enqueued
        // by defining x-foo.
        document.body.appendChild(b);
      }
    })
    </script>
    ```

    This step will thus bail out the algorithm early when [upgrade an
    element](#concept-upgrade-an-element){#upgrades:concept-upgrade-an-element}
    is invoked with \"`b`\" a second time.
    :::
    ::::

2.  Set `element`{.variable}\'s [custom element
    definition](https://dom.spec.whatwg.org/#concept-element-custom-element-definition){#upgrades:concept-element-custom-element-definition
    x-internal="concept-element-custom-element-definition"} to
    `definition`{.variable}.

3.  Set `element`{.variable}\'s [custom element
    state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#upgrades:custom-element-state-2
    x-internal="custom-element-state"} to \"`failed`\".

    It will be set to \"`custom`\" [after the upgrade
    succeeds](#concept-upgrade-an-element-set-state-to-custom). For now,
    we set it to \"`failed`\" so that any reentrant invocations will hit
    [the above early-exit step](#concept-upgrade-an-element-early-exit).

4.  For each `attribute`{.variable} in `element`{.variable}\'s
    [attribute
    list](https://dom.spec.whatwg.org/#concept-element-attribute){#upgrades:attribute-list
    x-internal="attribute-list"}, in order, [enqueue a custom element
    callback
    reaction](#enqueue-a-custom-element-callback-reaction){#upgrades:enqueue-a-custom-element-callback-reaction}
    with `element`{.variable}, callback name
    \"`attributeChangedCallback`\", and « `attribute`{.variable}\'s
    local name, null, `attribute`{.variable}\'s value,
    `attribute`{.variable}\'s namespace ».

5.  If `element`{.variable} is
    [connected](https://dom.spec.whatwg.org/#connected){#upgrades:connected
    x-internal="connected"}, then [enqueue a custom element callback
    reaction](#enqueue-a-custom-element-callback-reaction){#upgrades:enqueue-a-custom-element-callback-reaction-2}
    with `element`{.variable}, callback name \"`connectedCallback`\",
    and « ».

6.  Add `element`{.variable} to the end of `definition`{.variable}\'s
    [construction
    stack](#concept-custom-element-definition-construction-stack){#upgrades:concept-custom-element-definition-construction-stack}.

7.  Let `C`{.variable} be `definition`{.variable}\'s
    [constructor](#concept-custom-element-definition-constructor){#upgrades:concept-custom-element-definition-constructor}.

8.  [Set](https://infra.spec.whatwg.org/#map-set){#upgrades:map-set
    x-internal="map-set"} the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#upgrades:surrounding-agent
    x-internal="surrounding-agent"}\'s [active custom element
    constructor
    map](#active-custom-element-constructor-map){#upgrades:active-custom-element-constructor-map}\[`C`{.variable}\]
    to `element`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#upgrades:element-custom-element-registry
    x-internal="element-custom-element-registry"}.

9.  Run the following steps while catching any exceptions:

    1.  If `definition`{.variable}\'s [disable
        shadow](#concept-custom-element-definition-disable-shadow){#upgrades:concept-custom-element-definition-disable-shadow}
        is true and `element`{.variable}\'s [shadow
        root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#upgrades:concept-element-shadow-root
        x-internal="concept-element-shadow-root"} is non-null, then
        throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#upgrades:notsupportederror
        x-internal="notsupportederror"}
        [`DOMException`{#upgrades:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        This is needed as
        [`attachShadow()`{#upgrades:dom-element-attachshadow}](https://dom.spec.whatwg.org/#dom-element-attachshadow){x-internal="dom-element-attachshadow"}
        does not use [look up a custom element
        definition](#look-up-a-custom-element-definition){#upgrades:look-up-a-custom-element-definition}
        while
        [`attachInternals()`{#upgrades:dom-attachinternals}](#dom-attachinternals)
        does.

    2.  Set `element`{.variable}\'s [custom element
        state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#upgrades:custom-element-state-3
        x-internal="custom-element-state"} to \"`precustomized`\".

    3.  Let `constructResult`{.variable} be the result of
        [constructing](https://webidl.spec.whatwg.org/#construct-a-callback-function){#upgrades:es-constructing-callback-functions
        x-internal="es-constructing-callback-functions"} `C`{.variable},
        with no arguments.

        If `C`{.variable}
        [non-conformantly](#custom-element-conformance) uses an API
        decorated with the
        [`[CEReactions]`{#upgrades:cereactions}](#cereactions) extended
        attribute, then the reactions enqueued at the beginning of this
        algorithm will execute during this step, before `C`{.variable}
        finishes and control returns to this algorithm. Otherwise, they
        will execute after `C`{.variable} and the rest of the upgrade
        process finishes.

    4.  If
        [SameValue](https://tc39.es/ecma262/#sec-samevalue){#upgrades:samevalue
        x-internal="samevalue"}(`constructResult`{.variable},
        `element`{.variable}) is false, then throw a
        [`TypeError`{#upgrades:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

        This can occur if `C`{.variable} constructs another instance of
        the same custom element before calling `super()`, or if
        `C`{.variable} uses JavaScript\'s `return`-override feature to
        return an arbitrary
        [`HTMLElement`{#upgrades:htmlelement}](dom.html#htmlelement)
        object from the constructor.

    Then, perform the following steps, regardless of whether the above
    steps threw an exception or not:

    1.  [Remove](https://infra.spec.whatwg.org/#map-remove){#upgrades:map-remove
        x-internal="map-remove"} the [surrounding
        agent](https://tc39.es/ecma262/#surrounding-agent){#upgrades:surrounding-agent-2
        x-internal="surrounding-agent"}\'s [active custom element
        constructor
        map](#active-custom-element-constructor-map){#upgrades:active-custom-element-constructor-map-2}\[`C`{.variable}\].

        This is a no-op if `C`{.variable} immediately calls `super()` as
        it ought to do.

    2.  Remove the last entry from the end of `definition`{.variable}\'s
        [construction
        stack](#concept-custom-element-definition-construction-stack){#upgrades:concept-custom-element-definition-construction-stack-2}.

        ::: note
        Assuming `C`{.variable} calls `super()` (as it will if it is
        [conformant](#custom-element-conformance)), and that the call
        succeeds, this will be the [*already constructed*
        marker](#concept-already-constructed-marker){#upgrades:concept-already-constructed-marker}
        that replaced the `element`{.variable} we pushed at the
        beginning of this algorithm. (The [HTML element
        constructor](dom.html#html-element-constructors) carries out
        this replacement.)

        If `C`{.variable} does not call `super()` (i.e. it is not
        [conformant](#custom-element-conformance)), or if any step in
        the [HTML element
        constructor](dom.html#html-element-constructors) throws, then
        this entry will still be `element`{.variable}.
        :::

    Finally, if the above steps threw an exception, then:

    1.  Set `element`{.variable}\'s [custom element
        definition](https://dom.spec.whatwg.org/#concept-element-custom-element-definition){#upgrades:concept-element-custom-element-definition-2
        x-internal="concept-element-custom-element-definition"} to null.

    2.  Empty `element`{.variable}\'s [custom element reaction
        queue](#custom-element-reaction-queue){#upgrades:custom-element-reaction-queue}.

    3.  Rethrow the exception (thus terminating this algorithm).

    If the above steps threw an exception, then `element`{.variable}\'s
    [custom element
    state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#upgrades:custom-element-state-4
    x-internal="custom-element-state"} will remain \"`failed`\" or
    \"`precustomized`\".

10. If `element`{.variable} is a [form-associated custom
    element](#form-associated-custom-element){#upgrades:form-associated-custom-element},
    then:

    1.  [Reset the form
        owner](form-control-infrastructure.html#reset-the-form-owner){#upgrades:reset-the-form-owner}
        of `element`{.variable}. If `element`{.variable} is associated
        with a
        [`form`{#upgrades:the-form-element}](forms.html#the-form-element)
        element, then [enqueue a custom element callback
        reaction](#enqueue-a-custom-element-callback-reaction){#upgrades:enqueue-a-custom-element-callback-reaction-3}
        with `element`{.variable}, callback name
        \"`formAssociatedCallback`\", and « the associated
        [`form`{#upgrades:the-form-element-2}](forms.html#the-form-element)
        ».

    2.  If `element`{.variable} is
        [disabled](form-control-infrastructure.html#concept-fe-disabled){#upgrades:concept-fe-disabled},
        then [enqueue a custom element callback
        reaction](#enqueue-a-custom-element-callback-reaction){#upgrades:enqueue-a-custom-element-callback-reaction-4}
        with `element`{.variable}, callback name
        \"`formDisabledCallback`\", and « true ».

11. ::: {#concept-upgrade-an-element-set-state-to-custom}
    Set `element`{.variable}\'s [custom element
    state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#upgrades:custom-element-state-5
    x-internal="custom-element-state"} to \"`custom`\".
    :::

To [try to upgrade an element]{#concept-try-upgrade .dfn export=""}
given an element `element`{.variable}:

1.  Let `definition`{.variable} be the result of [looking up a custom
    element
    definition](#look-up-a-custom-element-definition){#upgrades:look-up-a-custom-element-definition-2}
    given `element`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#upgrades:element-custom-element-registry-2
    x-internal="element-custom-element-registry"},
    `element`{.variable}\'s
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#upgrades:concept-element-namespace
    x-internal="concept-element-namespace"}, `element`{.variable}\'s
    [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#upgrades:concept-element-local-name
    x-internal="concept-element-local-name"}, and
    `element`{.variable}\'s [`is`
    value](https://dom.spec.whatwg.org/#concept-element-is-value){#upgrades:concept-element-is-value
    x-internal="concept-element-is-value"}.

2.  If `definition`{.variable} is not null, then [enqueue a custom
    element upgrade
    reaction](#enqueue-a-custom-element-upgrade-reaction){#upgrades:enqueue-a-custom-element-upgrade-reaction}
    given `element`{.variable} and `definition`{.variable}.

#### [4.13.6]{.secno} Custom element reactions[](#custom-element-reactions){.self-link}

A [custom
element](#custom-element){#custom-element-reactions:custom-element}
possesses the ability to respond to certain occurrences by running
author code:

- When [upgraded](#upgrades){#custom-element-reactions:upgrades}, its
  [constructor](#custom-element-constructor){#custom-element-reactions:custom-element-constructor}
  is run, with no arguments.

- When it [becomes
  connected](infrastructure.html#becomes-connected){#custom-element-reactions:becomes-connected},
  its `connectedCallback` is called, with no arguments.

- When it [becomes
  disconnected](infrastructure.html#becomes-disconnected){#custom-element-reactions:becomes-disconnected},
  its `disconnectedCallback` is called, with no arguments.

- When it is
  [moved](https://dom.spec.whatwg.org/#concept-node-move-ext){#custom-element-reactions:concept-node-move-ext
  x-internal="concept-node-move-ext"}, its `connectedMoveCallback` is
  called, with no arguments.

- When it is
  [adopted](https://dom.spec.whatwg.org/#concept-node-adopt){#custom-element-reactions:concept-node-adopt
  x-internal="concept-node-adopt"} into a new document, its
  `adoptedCallback` is called, given the old document and new document
  as arguments.

- When any of its attributes are
  [changed](https://dom.spec.whatwg.org/#concept-element-attributes-change){#custom-element-reactions:concept-element-attributes-change
  x-internal="concept-element-attributes-change"},
  [appended](https://dom.spec.whatwg.org/#concept-element-attributes-append){#custom-element-reactions:concept-element-attributes-append
  x-internal="concept-element-attributes-append"},
  [removed](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#custom-element-reactions:concept-element-attributes-remove
  x-internal="concept-element-attributes-remove"}, or
  [replaced](https://dom.spec.whatwg.org/#concept-element-attributes-replace){#custom-element-reactions:concept-element-attributes-replace
  x-internal="concept-element-attributes-replace"}, its
  `attributeChangedCallback` is called, given the attribute\'s local
  name, old value, new value, and namespace as arguments. (An
  attribute\'s old or new value is considered to be null when the
  attribute is added or removed, respectively.)

- When the user agent [resets the form
  owner](form-control-infrastructure.html#reset-the-form-owner){#custom-element-reactions:reset-the-form-owner}
  of a [form-associated custom
  element](#form-associated-custom-element){#custom-element-reactions:form-associated-custom-element}
  and doing so changes the form owner, its `formAssociatedCallback` is
  called, given the new form owner (or null if no owner) as an argument.

- When the form owner of a [form-associated custom
  element](#form-associated-custom-element){#custom-element-reactions:form-associated-custom-element-2}
  is
  [reset](form-control-infrastructure.html#concept-form-reset){#custom-element-reactions:concept-form-reset},
  its `formResetCallback` is called.

- When the
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#custom-element-reactions:concept-fe-disabled}
  state of a [form-associated custom
  element](#form-associated-custom-element){#custom-element-reactions:form-associated-custom-element-3}
  is changed, its `formDisabledCallback` is called, given the new state
  as an argument.

- When user agent updates a [form-associated custom
  element](#form-associated-custom-element){#custom-element-reactions:form-associated-custom-element-4}\'s
  value on behalf of a user or [as part of
  navigation](browsing-the-web.html#restore-persisted-state){#custom-element-reactions:restore-persisted-state},
  its `formStateRestoreCallback` is called, given the new state and a
  string indicating a reason, \"`autocomplete`\" or \"`restore`\", as
  arguments.

We call these reactions collectively [custom element
reactions]{#concept-custom-element-reaction .dfn}.

The way in which [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction}
are invoked is done with special care, to avoid running author code
during the middle of delicate operations. Effectively, they are delayed
until \"just before returning to user script\". This means that for most
purposes they appear to execute synchronously, but in the case of
complicated composite operations (like
[cloning](https://dom.spec.whatwg.org/#concept-node-clone){#custom-element-reactions:concept-node-clone
x-internal="concept-node-clone"}, or
[range](https://dom.spec.whatwg.org/#concept-range){#custom-element-reactions:concept-range
x-internal="concept-range"} manipulation), they will instead be delayed
until after all the relevant user agent processing steps have completed,
and then run together as a batch.

Additionally, the precise ordering of these reactions is managed via a
somewhat-complicated stack-of-queues system, described below. The
intention behind this system is to guarantee that [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-2}
always are invoked in the same order as their triggering actions, at
least within the local context of a single [custom
element](#custom-element){#custom-element-reactions:custom-element-2}.
(Because [custom element
reaction](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-3}
code can perform its own mutations, it is not possible to give a global
ordering guarantee across multiple elements.)

------------------------------------------------------------------------

Each [similar-origin window
agent](webappapis.html#similar-origin-window-agent){#custom-element-reactions:similar-origin-window-agent}
has a [custom element reactions stack]{#custom-element-reactions-stack
.dfn}, which is initially empty. A [similar-origin window
agent](webappapis.html#similar-origin-window-agent){#custom-element-reactions:similar-origin-window-agent-2}\'s
[current element queue]{#current-element-queue .dfn} is the [element
queue](#element-queue){#custom-element-reactions:element-queue} at the
top of its [custom element reactions
stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack}.
Each item in the stack is an [element queue]{#element-queue .dfn}, which
is initially empty as well. Each item in an [element
queue](#element-queue){#custom-element-reactions:element-queue-2} is an
element. (The elements are not necessarily
[custom](https://dom.spec.whatwg.org/#concept-element-custom){#custom-element-reactions:concept-element-custom
x-internal="concept-element-custom"} yet, since this queue is used for
[upgrades](#upgrades){#custom-element-reactions:upgrades-2} as well.)

Each [custom element reactions
stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack-2}
has an associated [backup element queue]{#backup-element-queue .dfn},
which is an initially-empty [element
queue](#element-queue){#custom-element-reactions:element-queue-3}.
Elements are pushed onto the [backup element
queue](#backup-element-queue){#custom-element-reactions:backup-element-queue}
during operations that affect the DOM without going through an API
decorated with
[`[CEReactions]`{#custom-element-reactions:cereactions}](#cereactions),
or through the parser\'s [create an element for the
token](parsing.html#create-an-element-for-the-token){#custom-element-reactions:create-an-element-for-the-token}
algorithm. An example of this is a user-initiated editing operation
which modifies the descendants or attributes of an
[editable](https://w3c.github.io/editing/docs/execCommand/#editable){#custom-element-reactions:editable
x-internal="editable"} element. To prevent reentrancy when processing
the [backup element
queue](#backup-element-queue){#custom-element-reactions:backup-element-queue-2},
each [custom element reactions
stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack-3}
also has a [processing the backup element
queue]{#processing-the-backup-element-queue .dfn} flag, initially unset.

All elements have an associated [custom element reaction
queue]{#custom-element-reaction-queue .dfn}, initially empty. Each item
in the [custom element reaction
queue](#custom-element-reaction-queue){#custom-element-reactions:custom-element-reaction-queue}
is of one of two types:

- An [upgrade reaction]{#upgrade-reaction .dfn}, which will
  [upgrade](#upgrades){#custom-element-reactions:upgrades-3} the custom
  element and contains a [custom element
  definition](#custom-element-definition){#custom-element-reactions:custom-element-definition};
  or

- A [callback reaction]{#callback-reaction .dfn}, which will call a
  lifecycle callback, and contains a callback function as well as a list
  of arguments.

This is all summarized in the following schematic diagram:

![A custom element reactions stack consists of a stack of element
queues. Zooming in on a particular queue, we see that it contains a
number of elements (in our example, \<x-a\>, then \<x-b\>, then
\<x-c\>). Any particular element in the queue then has a custom element
reaction queue. Zooming in on the custom element reaction queue, we see
that it contains a variety of queued-up reactions (in our example,
upgrade, then attribute changed, then another attribute changed, then
connected).](/images/custom-element-reactions.svg){.darkmode-aware
style="width: 80%; max-width: 580px;"}

To [enqueue an element on the appropriate element
queue]{#enqueue-an-element-on-the-appropriate-element-queue .dfn}, given
an element `element`{.variable}, run the following steps:

1.  Let `reactionsStack`{.variable} be `element`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#custom-element-reactions:relevant-agent}\'s
    [custom element reactions
    stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack-4}.

2.  If `reactionsStack`{.variable} is empty, then:

    1.  Add `element`{.variable} to `reactionsStack`{.variable}\'s
        [backup element
        queue](#backup-element-queue){#custom-element-reactions:backup-element-queue-3}.

    2.  If `reactionsStack`{.variable}\'s [processing the backup element
        queue](#processing-the-backup-element-queue){#custom-element-reactions:processing-the-backup-element-queue}
        flag is set, then return.

    3.  Set `reactionsStack`{.variable}\'s [processing the backup
        element
        queue](#processing-the-backup-element-queue){#custom-element-reactions:processing-the-backup-element-queue-2}
        flag.

    4.  [Queue a
        microtask](webappapis.html#queue-a-microtask){#custom-element-reactions:queue-a-microtask}
        to perform the following steps:

        1.  [Invoke custom element
            reactions](#invoke-custom-element-reactions){#custom-element-reactions:invoke-custom-element-reactions}
            in `reactionsStack`{.variable}\'s [backup element
            queue](#backup-element-queue){#custom-element-reactions:backup-element-queue-4}.

        2.  Unset `reactionsStack`{.variable}\'s [processing the backup
            element
            queue](#processing-the-backup-element-queue){#custom-element-reactions:processing-the-backup-element-queue-3}
            flag.

3.  Otherwise, add `element`{.variable} to `element`{.variable}\'s
    [relevant
    agent](webappapis.html#relevant-agent){#custom-element-reactions:relevant-agent-2}\'s
    [current element
    queue](#current-element-queue){#custom-element-reactions:current-element-queue}.

To [enqueue a custom element callback
reaction]{#enqueue-a-custom-element-callback-reaction .dfn export=""},
given a [custom
element](#custom-element){#custom-element-reactions:custom-element-3}
`element`{.variable}, a callback name `callbackName`{.variable}, and a
list of arguments `args`{.variable}, run the following steps:

1.  Let `definition`{.variable} be `element`{.variable}\'s [custom
    element
    definition](https://dom.spec.whatwg.org/#concept-element-custom-element-definition){#custom-element-reactions:concept-element-custom-element-definition
    x-internal="concept-element-custom-element-definition"}.

2.  Let `callback`{.variable} be the value of the entry in
    `definition`{.variable}\'s [lifecycle
    callbacks](#concept-custom-element-definition-lifecycle-callbacks){#custom-element-reactions:concept-custom-element-definition-lifecycle-callbacks}
    with key `callbackName`{.variable}.

3.  If `callbackName`{.variable} is \"`connectedMoveCallback`\" and
    `callback`{.variable} is null:

    1.  Let `disconnectedCallback`{.variable} be the value of the entry
        in `definition`{.variable}\'s [lifecycle
        callbacks](#concept-custom-element-definition-lifecycle-callbacks){#custom-element-reactions:concept-custom-element-definition-lifecycle-callbacks-2}
        with key \"`disconnectedCallback`\".

    2.  Let `connectedCallback`{.variable} be the value of the entry in
        `definition`{.variable}\'s [lifecycle
        callbacks](#concept-custom-element-definition-lifecycle-callbacks){#custom-element-reactions:concept-custom-element-definition-lifecycle-callbacks-3}
        with key \"`connectedCallback`\".

    3.  If `connectedCallback`{.variable} and
        `disconnectedCallback`{.variable} are null, then return.

    4.  Set `callback`{.variable} to the following steps:

        1.  If `disconnectedCallback`{.variable} is not null, then call
            `disconnectedCallback`{.variable} with no arguments.

        2.  If `connectedCallback`{.variable} is not null, then call
            `connectedCallback`{.variable} with no arguments.

4.  If `callback`{.variable} is null, then return.

5.  If `callbackName`{.variable} is \"`attributeChangedCallback`\":

    1.  Let `attributeName`{.variable} be the first element of
        `args`{.variable}.

    2.  If `definition`{.variable}\'s [observed
        attributes](#concept-custom-element-definition-observed-attributes){#custom-element-reactions:concept-custom-element-definition-observed-attributes}
        does not contain `attributeName`{.variable}, then return.

6.  Add a new [callback
    reaction](#callback-reaction){#custom-element-reactions:callback-reaction}
    to `element`{.variable}\'s [custom element reaction
    queue](#custom-element-reaction-queue){#custom-element-reactions:custom-element-reaction-queue-2},
    with callback function `callback`{.variable} and arguments
    `args`{.variable}.

7.  [Enqueue an element on the appropriate element
    queue](#enqueue-an-element-on-the-appropriate-element-queue){#custom-element-reactions:enqueue-an-element-on-the-appropriate-element-queue}
    given `element`{.variable}.

To [enqueue a custom element upgrade
reaction]{#enqueue-a-custom-element-upgrade-reaction .dfn export=""},
given an element `element`{.variable} and [custom element
definition](#custom-element-definition){#custom-element-reactions:custom-element-definition-2}
`definition`{.variable}, run the following steps:

1.  Add a new [upgrade
    reaction](#upgrade-reaction){#custom-element-reactions:upgrade-reaction}
    to `element`{.variable}\'s [custom element reaction
    queue](#custom-element-reaction-queue){#custom-element-reactions:custom-element-reaction-queue-3},
    with [custom element
    definition](#custom-element-definition){#custom-element-reactions:custom-element-definition-3}
    `definition`{.variable}.

2.  [Enqueue an element on the appropriate element
    queue](#enqueue-an-element-on-the-appropriate-element-queue){#custom-element-reactions:enqueue-an-element-on-the-appropriate-element-queue-2}
    given `element`{.variable}.

To [invoke custom element reactions]{#invoke-custom-element-reactions
.dfn} in an [element
queue](#element-queue){#custom-element-reactions:element-queue-4}
`queue`{.variable}, run the following steps:

1.  While `queue`{.variable} is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#custom-element-reactions:list-is-empty
    x-internal="list-is-empty"}:

    1.  Let `element`{.variable} be the result of
        [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#custom-element-reactions:dequeue
        x-internal="dequeue"} from `queue`{.variable}.

    2.  Let `reactions`{.variable} be `element`{.variable}\'s [custom
        element reaction
        queue](#custom-element-reaction-queue){#custom-element-reactions:custom-element-reaction-queue-4}.

    3.  Repeat until `reactions`{.variable} is empty:

        1.  Remove the first element of `reactions`{.variable}, and let
            `reaction`{.variable} be that element. Switch on
            `reaction`{.variable}\'s type:

            [upgrade reaction](#upgrade-reaction){#custom-element-reactions:upgrade-reaction-2}

            :   [Upgrade](#concept-upgrade-an-element){#custom-element-reactions:concept-upgrade-an-element}
                `element`{.variable} using `reaction`{.variable}\'s
                [custom element
                definition](#custom-element-definition){#custom-element-reactions:custom-element-definition-4}.

                If this throws an exception, catch it, and
                [report](webappapis.html#report-an-exception){#custom-element-reactions:report-an-exception}
                it for `reaction`{.variable}\'s [custom element
                definition](#custom-element-definition){#custom-element-reactions:custom-element-definition-5}\'s
                [constructor](#concept-custom-element-definition-constructor){#custom-element-reactions:concept-custom-element-definition-constructor}\'s
                corresponding JavaScript object\'s [associated
                realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#custom-element-reactions:associated-realm
                x-internal="associated-realm"}\'s [global
                object](webappapis.html#concept-realm-global){#custom-element-reactions:concept-realm-global}.

            [callback reaction](#callback-reaction){#custom-element-reactions:callback-reaction-2}

            :   [Invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#custom-element-reactions:es-invoking-callback-functions
                x-internal="es-invoking-callback-functions"}
                `reaction`{.variable}\'s callback function with
                `reaction`{.variable}\'s arguments and \"`report`\", and
                *[callback this
                value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){x-internal="dfn-callback-this-value"}*
                set to `element`{.variable}.

------------------------------------------------------------------------

To ensure [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-4}
are triggered appropriately, we introduce the
[`[CEReactions]`]{#cereactions .dfn dfn-type="extended-attribute"
lt="CEReactions"} IDL [extended
attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute){#custom-element-reactions:extended-attribute
x-internal="extended-attribute"}. It indicates that the relevant
algorithm is to be supplemented with additional steps in order to
appropriately track and invoke [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-5}.

The
[`[CEReactions]`{#custom-element-reactions:cereactions-2}](#cereactions)
extended attribute must take no arguments, and must not appear on
anything other than an operation, attribute, setter, or deleter.
Additionally, it must not appear on readonly attributes.

Operations, attributes, setters, or deleters annotated with the
[`[CEReactions]`{#custom-element-reactions:cereactions-3}](#cereactions)
extended attribute must run the following steps in place of the ones
specified in their description:

1.  [Push](https://infra.spec.whatwg.org/#stack-push){#custom-element-reactions:stack-push
    x-internal="stack-push"} a new [element
    queue](#element-queue){#custom-element-reactions:element-queue-5}
    onto this object\'s [relevant
    agent](webappapis.html#relevant-agent){#custom-element-reactions:relevant-agent-3}\'s
    [custom element reactions
    stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack-5}.

2.  Run the originally-specified steps for this construct, catching any
    exceptions. If the steps return a value, let `value`{.variable} be
    the returned value. If they throw an exception, let
    `exception`{.variable} be the thrown exception.

3.  Let `queue`{.variable} be the result of
    [popping](https://infra.spec.whatwg.org/#stack-pop){#custom-element-reactions:stack-pop
    x-internal="stack-pop"} from this object\'s [relevant
    agent](webappapis.html#relevant-agent){#custom-element-reactions:relevant-agent-4}\'s
    [custom element reactions
    stack](#custom-element-reactions-stack){#custom-element-reactions:custom-element-reactions-stack-6}.

4.  [Invoke custom element
    reactions](#invoke-custom-element-reactions){#custom-element-reactions:invoke-custom-element-reactions-2}
    in `queue`{.variable}.

5.  If an exception `exception`{.variable} was thrown by the original
    steps, rethrow `exception`{.variable}.

6.  If a value `value`{.variable} was returned from the original steps,
    return `value`{.variable}.

::: note
The intent behind this extended attribute is somewhat subtle. One way of
accomplishing its goals would be to say that every operation, attribute,
setter, and deleter on the platform must have these steps inserted, and
to allow implementers to optimize away unnecessary cases (where no DOM
mutation is possible that could cause [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-6}
to occur).

However, in practice this imprecision could lead to non-interoperable
implementations of [custom element
reactions](#concept-custom-element-reaction){#custom-element-reactions:concept-custom-element-reaction-7},
as some implementations might forget to invoke these steps in some
cases. Instead, we settled on the approach of explicitly annotating all
relevant IDL constructs, as a way of ensuring interoperable behavior and
helping implementations easily pinpoint all cases where these steps are
necessary.
:::

Any nonstandard APIs introduced by the user agent that could modify the
DOM in such a way as to cause [enqueuing a custom element callback
reaction](#enqueue-a-custom-element-callback-reaction){#custom-element-reactions:enqueue-a-custom-element-callback-reaction}
or [enqueuing a custom element upgrade
reaction](#enqueue-a-custom-element-upgrade-reaction){#custom-element-reactions:enqueue-a-custom-element-upgrade-reaction},
for example by modifying any attributes or child elements, must also be
decorated with the
[`[CEReactions]`{#custom-element-reactions:cereactions-4}](#cereactions)
attribute.

::: note
As of the time of this writing, the following nonstandard or
not-yet-standardized APIs are known to fall into this category:

- [`HTMLInputElement`{#custom-element-reactions:htmlinputelement}](input.html#htmlinputelement)\'s
  `webkitdirectory` and `incremental` IDL attributes

- [`HTMLLinkElement`{#custom-element-reactions:htmllinkelement}](semantics.html#htmllinkelement)\'s
  `scope` IDL attribute
:::

#### [4.13.7]{.secno} Element internals[](#element-internals){.self-link}

Certain capabilities are meant to be available to a custom element
author, but not to a custom element consumer. These are provided by the
[`element.attachInternals()`{#element-internals:dom-attachinternals}](#dom-attachinternals)
method, which returns an instance of
[`ElementInternals`{#element-internals:elementinternals}](#elementinternals).
The properties and methods of
[`ElementInternals`{#element-internals:elementinternals-2}](#elementinternals)
allow control over internal features which the user agent provides to
all elements.

`element`{.variable}`.`[`attachInternals()`](#dom-attachinternals){#dom-attachinternals-dev}

:   Returns an
    [`ElementInternals`{#element-internals:elementinternals-3}](#elementinternals)
    object targeting the [custom
    element](#custom-element){#element-internals:custom-element}
    `element`{.variable}. Throws an exception if `element`{.variable} is
    not a [custom
    element](#custom-element){#element-internals:custom-element-2}, if
    the \"`internals`\" feature was disabled as part of the element
    definition, or if it is called twice on the same element.

Each
[`HTMLElement`{#element-internals:htmlelement}](dom.html#htmlelement)
has an [attached internals]{#attached-internals .dfn} (null or an
[`ElementInternals`{#element-internals:elementinternals-4}](#elementinternals)
object), initially null.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/attachInternals](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/attachInternals "The HTMLElement.attachInternals() method returns an ElementInternals object. This method allows a custom element to participate in HTML forms. The ElementInternals interface provides utilities for working with these elements in the same way you would work with any standard HTML form element, and also exposes the Accessibility Object Model to the element.")

Support in all current engines.

::: support
[Firefox93+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`attachInternals()`]{#dom-attachinternals .dfn
dfn-for="HTMLElement" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this
    x-internal="this"}\'s [`is`
    value](https://dom.spec.whatwg.org/#concept-element-is-value){#element-internals:concept-element-is-value
    x-internal="concept-element-is-value"} is not null, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#element-internals:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#element-internals:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `definition`{.variable} be the result of [looking up a custom
    element
    definition](#look-up-a-custom-element-definition){#element-internals:look-up-a-custom-element-definition}
    given
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-2
    x-internal="this"}\'s [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#element-internals:element-custom-element-registry
    x-internal="element-custom-element-registry"},
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-3
    x-internal="this"}\'s
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#element-internals:concept-element-namespace
    x-internal="concept-element-namespace"},
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-4
    x-internal="this"}\'s [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#element-internals:concept-element-local-name
    x-internal="concept-element-local-name"}, and null.

3.  If `definition`{.variable} is null, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#element-internals:notsupportederror-2
    x-internal="notsupportederror"}
    [`DOMException`{#element-internals:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  If `definition`{.variable}\'s [disable
    internals](#concept-custom-element-definition-disable-internals){#element-internals:concept-custom-element-definition-disable-internals}
    is true, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#element-internals:notsupportederror-3
    x-internal="notsupportederror"}
    [`DOMException`{#element-internals:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-5
    x-internal="this"}\'s [attached
    internals](#attached-internals){#element-internals:attached-internals}
    is non-null, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#element-internals:notsupportederror-4
    x-internal="notsupportederror"}
    [`DOMException`{#element-internals:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  If
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-6
    x-internal="this"}\'s [custom element
    state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#element-internals:custom-element-state
    x-internal="custom-element-state"} is not \"`precustomized`\" or
    \"`custom`\", then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#element-internals:notsupportederror-5
    x-internal="notsupportederror"}
    [`DOMException`{#element-internals:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

7.  Set
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-7
    x-internal="this"}\'s [attached
    internals](#attached-internals){#element-internals:attached-internals-2}
    to a new
    [`ElementInternals`{#element-internals:elementinternals-5}](#elementinternals)
    instance whose [target
    element](#internals-target){#element-internals:internals-target} is
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-8
    x-internal="this"}.

8.  Return
    [this](https://webidl.spec.whatwg.org/#this){#element-internals:this-9
    x-internal="this"}\'s [attached
    internals](#attached-internals){#element-internals:attached-internals-3}.

##### [4.13.7.1]{.secno} The [`ElementInternals`{#the-elementinternals-interface:elementinternals}](#elementinternals) interface[](#the-elementinternals-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[ElementInternals](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals "The ElementInternals interface of the Document Object Model gives web developers a way to allow custom elements to fully participate in HTML forms. It provides utilities for working with these elements in the same way you would work with any standard HTML form element, and also exposes the Accessibility Object Model to the element.")

Support in all current engines.

::: support
[Firefox93+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The IDL for the
[`ElementInternals`{#the-elementinternals-interface:elementinternals-2}](#elementinternals)
interface is as follows, with the various operations and attributes
defined in the following sections:

``` idl
[Exposed=Window]
interface ElementInternals {
  // Shadow root access
  readonly attribute ShadowRoot? shadowRoot;

  // Form-associated custom elements
  undefined setFormValue((File or USVString or FormData)? value,
                         optional (File or USVString or FormData)? state);

  readonly attribute HTMLFormElement? form;

  undefined setValidity(optional ValidityStateFlags flags = {},
                        optional DOMString message,
                        optional HTMLElement anchor);
  readonly attribute boolean willValidate;
  readonly attribute ValidityState validity;
  readonly attribute DOMString validationMessage;
  boolean checkValidity();
  boolean reportValidity();

  readonly attribute NodeList labels;

  // Custom state pseudo-class
  [SameObject] readonly attribute CustomStateSet states;
};

// Accessibility semantics
ElementInternals includes ARIAMixin;

dictionary ValidityStateFlags {
  boolean valueMissing = false;
  boolean typeMismatch = false;
  boolean patternMismatch = false;
  boolean tooLong = false;
  boolean tooShort = false;
  boolean rangeUnderflow = false;
  boolean rangeOverflow = false;
  boolean stepMismatch = false;
  boolean badInput = false;
  boolean customError = false;
};
```

Each
[`ElementInternals`{#the-elementinternals-interface:elementinternals-4}](#elementinternals)
has a [target element]{#internals-target .dfn}, which is a [custom
element](#custom-element){#the-elementinternals-interface:custom-element}.

##### [4.13.7.2]{.secno} Shadow root access[](#shadow-root-access){.self-link}

`internals`{.variable}`.`[`shadowRoot`](#dom-elementinternals-shadowroot){#shadow-root-access:dom-elementinternals-shadowroot}

:   Returns the
    [`ShadowRoot`{#shadow-root-access:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}
    for `internals`{.variable}\'s [target
    element](#internals-target){#shadow-root-access:internals-target},
    if the [target
    element](#internals-target){#shadow-root-access:internals-target-2}
    is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#shadow-root-access:shadow-host
    x-internal="shadow-host"}, or null otherwise.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/shadowRoot](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/shadowRoot "The shadowRoot read-only property of the ElementInternals interface returns the ShadowRoot for this element.")

Support in all current engines.

::: support
[Firefox93+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome88+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge88+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`shadowRoot`]{#dom-elementinternals-shadowroot .dfn
dfn-for="ElementInternals" dfn-type="attribute"} getter steps are:

1.  Let `target`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#shadow-root-access:this
    x-internal="this"}\'s [target
    element](#internals-target){#shadow-root-access:internals-target-3}.

2.  If `target`{.variable} is not a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#shadow-root-access:shadow-host-2
    x-internal="shadow-host"}, then return null.

3.  Let `shadow`{.variable} be `target`{.variable}\'s [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#shadow-root-access:concept-element-shadow-root
    x-internal="concept-element-shadow-root"}.

4.  If `shadow`{.variable}\'s [available to element
    internals](https://dom.spec.whatwg.org/#shadowroot-available-to-element-internals){#shadow-root-access:available-to-element-internals
    x-internal="available-to-element-internals"} is false, then return
    null.

5.  Return `shadow`{.variable}.

##### [4.13.7.3]{.secno} Form-associated custom elements[](#form-associated-custom-elements){.self-link}

`internals`{.variable}`.`[`setFormValue`](#dom-elementinternals-setformvalue){#form-associated-custom-elements:dom-elementinternals-setformvalue}`(``value`{.variable}`)`

:   Sets both the
    [state](#face-state){#form-associated-custom-elements:face-state}
    and [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value}
    of `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target}
    to `value`{.variable}.

    If `value`{.variable} is null, the element won\'t participate in
    form submission.

`internals`{.variable}`.`[`setFormValue`](#dom-elementinternals-setformvalue){#form-associated-custom-elements:dom-elementinternals-setformvalue-2}`(``value`{.variable}`, ``state`{.variable}`)`

:   Sets the [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-2}
    of `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-2}
    to `value`{.variable}, and its
    [state](#face-state){#form-associated-custom-elements:face-state-2}
    to `state`{.variable}.

    If `value`{.variable} is null, the element won\'t participate in
    form submission.

`internals`{.variable}`.`[`form`](form-control-infrastructure.html#dom-elementinternals-form){#form-associated-custom-elements:dom-elementinternals-form}
:   Returns the [form
    owner](form-control-infrastructure.html#form-owner){#form-associated-custom-elements:form-owner}
    of `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-3}.

`internals`{.variable}`.`[`setValidity`](#dom-elementinternals-setvalidity){#form-associated-custom-elements:dom-elementinternals-setvalidity}`(``flags`{.variable}`, ``message`{.variable}` [, ``anchor`{.variable}` ])`
:   Marks `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-4}
    as suffering from the constraints indicated by the
    `flags`{.variable} argument, and sets the element\'s validation
    message to `message`{.variable}. If `anchor`{.variable} is
    specified, the user agent might use it to indicate problems with the
    constraints of `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-5}
    when the [form
    owner](form-control-infrastructure.html#form-owner){#form-associated-custom-elements:form-owner-2}
    is validated interactively or
    [`reportValidity()`{#form-associated-custom-elements:dom-elementinternals-reportvalidity}](form-control-infrastructure.html#dom-elementinternals-reportvalidity)
    is called.

`internals`{.variable}`.`[`setValidity`](#dom-elementinternals-setvalidity){#form-associated-custom-elements:dom-elementinternals-setvalidity-2}`({})`
:   Marks `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-6}
    as [satisfying its
    constraints](form-control-infrastructure.html#concept-fv-valid){#form-associated-custom-elements:concept-fv-valid}.

`internals`{.variable}`.`[`willValidate`](form-control-infrastructure.html#dom-elementinternals-willvalidate){#form-associated-custom-elements:dom-elementinternals-willvalidate}
:   Returns true if `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-7}
    will be validated when the form is submitted; false otherwise.

`internals`{.variable}`.`[`validity`](form-control-infrastructure.html#dom-elementinternals-validity){#form-associated-custom-elements:dom-elementinternals-validity}
:   Returns the
    [`ValidityState`{#form-associated-custom-elements:validitystate}](form-control-infrastructure.html#validitystate)
    object for `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-8}.

`internals`{.variable}`.`[`validationMessage`](#dom-elementinternals-validationmessage){#form-associated-custom-elements:dom-elementinternals-validationmessage}
:   Returns the error message that would be shown to the user if
    `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-9}
    was to be checked for validity.

`valid`{.variable}` = ``internals`{.variable}`.`[`checkValidity()`](form-control-infrastructure.html#dom-elementinternals-checkvalidity){#form-associated-custom-elements:dom-elementinternals-checkvalidity}
:   Returns true if `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-10}
    has no validity problems; false otherwise. Fires an
    [`invalid`{#form-associated-custom-elements:event-invalid}](indices.html#event-invalid)
    event at the element in the latter case.

`valid`{.variable}` = ``internals`{.variable}`.`[`reportValidity()`](form-control-infrastructure.html#dom-elementinternals-reportvalidity){#form-associated-custom-elements:dom-elementinternals-reportvalidity-2}
:   Returns true if `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-11}
    has no validity problems; otherwise, returns false, fires an
    [`invalid`{#form-associated-custom-elements:event-invalid-2}](indices.html#event-invalid)
    event at the element, and (if the event isn\'t canceled) reports the
    problem to the user.

`internals`{.variable}`.`[`labels`](forms.html#dom-elementinternals-labels){#form-associated-custom-elements:dom-elementinternals-labels}

:   Returns a
    [`NodeList`{#form-associated-custom-elements:nodelist}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
    of all the
    [`label`{#form-associated-custom-elements:the-label-element}](forms.html#the-label-element)
    elements that `internals`{.variable}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-12}
    is associated with.

Each [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element}
has [submission value]{#face-submission-value .dfn}. It is used to
provide one or more
[entries](form-control-infrastructure.html#form-entry){#form-associated-custom-elements:form-entry}
on form submission. The initial value of [submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-3}
is null, and [submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-4}
can be null, a string, a
[`File`{#form-associated-custom-elements:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"},
or a
[list](https://infra.spec.whatwg.org/#list){#form-associated-custom-elements:list
x-internal="list"} of
[entries](form-control-infrastructure.html#form-entry){#form-associated-custom-elements:form-entry-2}.

Each [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-2}
has [state]{#face-state .dfn}. It is information with which the user
agent can restore a user\'s input for the element. The initial value of
[state](#face-state){#form-associated-custom-elements:face-state-3} is
null, and
[state](#face-state){#form-associated-custom-elements:face-state-4} can
be null, a string, a
[`File`{#form-associated-custom-elements:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"},
or a
[list](https://infra.spec.whatwg.org/#list){#form-associated-custom-elements:list-2
x-internal="list"} of
[entries](form-control-infrastructure.html#form-entry){#form-associated-custom-elements:form-entry-3}.

The
[`setFormValue()`{#form-associated-custom-elements:dom-elementinternals-setformvalue-3}](#dom-elementinternals-setformvalue)
method is used by the custom element author to set the element\'s
[submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-5}
and [state](#face-state){#form-associated-custom-elements:face-state-5},
thus communicating these to the user agent.

When the user agent believes it is a good idea to restore a
[form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-3}\'s
[state](#face-state){#form-associated-custom-elements:face-state-6}, for
example [after
navigation](browsing-the-web.html#restore-persisted-state){#form-associated-custom-elements:restore-persisted-state}
or restarting the user agent, they may [enqueue a custom element
callback
reaction](#enqueue-a-custom-element-callback-reaction){#form-associated-custom-elements:enqueue-a-custom-element-callback-reaction}
with that element, callback name \"`formStateRestoreCallback`\", and «
the state to be restored, \"`restore`\" ».

If the user agent has a form-filling assist feature, then when the
feature is invoked, it may [enqueue a custom element callback
reaction](#enqueue-a-custom-element-callback-reaction){#form-associated-custom-elements:enqueue-a-custom-element-callback-reaction-2}
with a [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-4},
callback name \"`formStateRestoreCallback`\", and « the state value
determined by history of state value and some heuristics,
\"`autocomplete`\" ».

In general, the
[state](#face-state){#form-associated-custom-elements:face-state-7} is
information specified by a user, and the [submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-6}
is a value after canonicalization or sanitization, suitable for
submission to the server. The following examples makes this concrete:

Suppose that we have a [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-5}
which asks a user to specify a date. The user specifies
[\"3/15/2019\"]{.kbd}, but the control wishes to submit `"2019-03-15"`
to the server. [\"3/15/2019\"]{.kbd} would be a
[state](#face-state){#form-associated-custom-elements:face-state-8} of
the element, and `"2019-03-15"` would be a [submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-7}.

Suppose you develop a custom element emulating a the behavior of the
existing
[checkbox](input.html#checkbox-state-(type=checkbox)){#form-associated-custom-elements:checkbox-state-(type=checkbox)}
[`input`{#form-associated-custom-elements:the-input-element}](input.html#the-input-element)
type. Its [submission
value](#face-submission-value){#form-associated-custom-elements:face-submission-value-8}
would be the value of its `value` content attribute, or the string
`"on"`. Its
[state](#face-state){#form-associated-custom-elements:face-state-9}
would be one of `"checked"`, `"unchecked"`, `"checked/indeterminate"`,
or `"unchecked/indeterminate"`.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/setFormValue](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/setFormValue "The setFormValue() method of the ElementInternals interface sets the element's submission value and state, communicating these to the user agent.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`setFormValue(``value`{.variable}`, ``state`{.variable}`)`]{#dom-elementinternals-setformvalue
.dfn dfn-for="ElementInternals" dfn-type="method"} method steps are:

1.  Let `element`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#form-associated-custom-elements:this
    x-internal="this"}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-13}.

2.  If `element`{.variable} is not a [form-associated custom
    element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-6},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#form-associated-custom-elements:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#form-associated-custom-elements:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Set `element`{.variable}\'s [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-9}
    to `value`{.variable} if `value`{.variable} is not a
    [`FormData`{#form-associated-custom-elements:formdata}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    object, or to a
    [clone](https://infra.spec.whatwg.org/#list-clone){#form-associated-custom-elements:list-clone
    x-internal="list-clone"} of `value`{.variable}\'s [entry
    list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#form-associated-custom-elements:formdata-entry-list
    x-internal="formdata-entry-list"} otherwise.

4.  If the `state`{.variable} argument of the function is omitted, set
    `element`{.variable}\'s
    [state](#face-state){#form-associated-custom-elements:face-state-10}
    to its [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-10}.

5.  Otherwise, if `state`{.variable} is a
    [`FormData`{#form-associated-custom-elements:formdata-2}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    object, set `element`{.variable}\'s
    [state](#face-state){#form-associated-custom-elements:face-state-11}
    to a
    [clone](https://infra.spec.whatwg.org/#list-clone){#form-associated-custom-elements:list-clone-2
    x-internal="list-clone"} of `state`{.variable}\'s [entry
    list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#form-associated-custom-elements:formdata-entry-list-2
    x-internal="formdata-entry-list"}.

6.  Otherwise, set `element`{.variable}\'s
    [state](#face-state){#form-associated-custom-elements:face-state-12}
    to `state`{.variable}.

------------------------------------------------------------------------

Each [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-7}
has validity flags named `valueMissing`, `typeMismatch`,
`patternMismatch`, `tooLong`, `tooShort`, `rangeUnderflow`,
`rangeOverflow`, `stepMismatch`, and `customError`. They are false
initially.

Each [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-8}
has a [validation message]{#face-validation-message .dfn} string. It is
the empty string initially.

Each [form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-9}
has a [validation anchor]{#face-validation-anchor .dfn} element. It is
null initially.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/setValidity](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/setValidity "The setValidity() method of the ElementInternals interface sets the validity of the element.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`setValidity(``flags`{.variable}`, ``message`{.variable}`, ``anchor`{.variable}`)`]{#dom-elementinternals-setvalidity
.dfn dfn-for="ElementInternals" dfn-type="method"} method steps are:

1.  Let `element`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#form-associated-custom-elements:this-2
    x-internal="this"}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-14}.

2.  If `element`{.variable} is not a [form-associated custom
    element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-10},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#form-associated-custom-elements:notsupportederror-2
    x-internal="notsupportederror"}
    [`DOMException`{#form-associated-custom-elements:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `flags`{.variable} contains one or more true values and
    `message`{.variable} is not given or is the empty string, then throw
    a
    [`TypeError`{#form-associated-custom-elements:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

4.  For each entry `flag`{.variable} → `value`{.variable} of
    `flags`{.variable}, set `element`{.variable}\'s validity flag with
    the name `flag`{.variable} to `value`{.variable}.

5.  Set `element`{.variable}\'s [validation
    message](#face-validation-message){#form-associated-custom-elements:face-validation-message}
    to the empty string if `message`{.variable} is not given or all of
    `element`{.variable}\'s validity flags are false, or to
    `message`{.variable} otherwise.

6.  If `element`{.variable}\'s `customError` validity flag is true, then
    set `element`{.variable}\'s [custom validity error
    message](form-control-infrastructure.html#custom-validity-error-message){#form-associated-custom-elements:custom-validity-error-message}
    to `element`{.variable}\'s [validation
    message](#face-validation-message){#form-associated-custom-elements:face-validation-message-2}.
    Otherwise, set `element`{.variable}\'s [custom validity error
    message](form-control-infrastructure.html#custom-validity-error-message){#form-associated-custom-elements:custom-validity-error-message-2}
    to the empty string.

7.  Set `element`{.variable}\'s [validation
    anchor](#face-validation-anchor){#form-associated-custom-elements:face-validation-anchor}
    to null if `anchor`{.variable} is not given. Otherwise, if
    `anchor`{.variable} is not a [shadow-including
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#form-associated-custom-elements:shadow-including-descendant
    x-internal="shadow-including-descendant"} of `element`{.variable},
    then throw a
    [\"`NotFoundError`\"](https://webidl.spec.whatwg.org/#notfounderror){#form-associated-custom-elements:notfounderror
    x-internal="notfounderror"}
    [`DOMException`{#form-associated-custom-elements:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
    Otherwise, set `element`{.variable}\'s [validation
    anchor](#face-validation-anchor){#form-associated-custom-elements:face-validation-anchor-2}
    to `anchor`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/validationMessage](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/validationMessage "The validationMessage read-only property of the ElementInternals interface returns the validation message for the element.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`validationMessage`]{#dom-elementinternals-validationmessage .dfn
dfn-for="ElementInternals" dfn-type="attribute"} getter steps are:

1.  Let `element`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#form-associated-custom-elements:this-3
    x-internal="this"}\'s [target
    element](#internals-target){#form-associated-custom-elements:internals-target-15}.

2.  If `element`{.variable} is not a [form-associated custom
    element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-11},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#form-associated-custom-elements:notsupportederror-3
    x-internal="notsupportederror"}
    [`DOMException`{#form-associated-custom-elements:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Return `element`{.variable}\'s [validation
    message](#face-validation-message){#form-associated-custom-elements:face-validation-message-3}.

The [entry construction algorithm]{#face-entry-construction .dfn} for a
[form-associated custom
element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-12},
given an element `element`{.variable} and an [entry
list](form-control-infrastructure.html#entry-list){#form-associated-custom-elements:entry-list}
`entry list`{.variable}, consists of the following steps:

1.  If `element`{.variable}\'s [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-11}
    is a
    [list](https://infra.spec.whatwg.org/#list){#form-associated-custom-elements:list-3
    x-internal="list"} of
    [entries](form-control-infrastructure.html#form-entry){#form-associated-custom-elements:form-entry-4},
    then
    [append](https://infra.spec.whatwg.org/#list-append){#form-associated-custom-elements:list-append
    x-internal="list-append"} each item of `element`{.variable}\'s
    [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-12}
    to `entry list`{.variable}, and return.

    In this case, user agent does not refer to the
    [`name`{#form-associated-custom-elements:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    content attribute value. An implementation of [form-associated
    custom
    element](#form-associated-custom-element){#form-associated-custom-elements:form-associated-custom-element-13}
    is responsible to decide names of
    [entries](form-control-infrastructure.html#form-entry){#form-associated-custom-elements:form-entry-5}.
    They can be the
    [`name`{#form-associated-custom-elements:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
    content attribute value, they can be strings based on the
    [`name`{#form-associated-custom-elements:attr-fe-name-3}](form-control-infrastructure.html#attr-fe-name)
    content attribute value, or they can be unrelated to the
    [`name`{#form-associated-custom-elements:attr-fe-name-4}](form-control-infrastructure.html#attr-fe-name)
    content attribute.

2.  If the element does not have a
    [`name`{#form-associated-custom-elements:attr-fe-name-5}](form-control-infrastructure.html#attr-fe-name)
    attribute specified, or its
    [`name`{#form-associated-custom-elements:attr-fe-name-6}](form-control-infrastructure.html#attr-fe-name)
    attribute\'s value is the empty string, then return.

3.  If the element\'s [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-13}
    is not null, [create an
    entry](form-control-infrastructure.html#create-an-entry){#form-associated-custom-elements:create-an-entry}
    with the
    [`name`{#form-associated-custom-elements:attr-fe-name-7}](form-control-infrastructure.html#attr-fe-name)
    attribute value and the [submission
    value](#face-submission-value){#form-associated-custom-elements:face-submission-value-14},
    and
    [append](https://infra.spec.whatwg.org/#list-append){#form-associated-custom-elements:list-append-2
    x-internal="list-append"} it to `entry list`{.variable}.

##### [4.13.7.4]{.secno} Accessibility semantics[](#accessibility-semantics){.self-link}

`internals`{.variable}`.`[`role`](https://w3c.github.io/aria/#idl-def-ariamixin-role){#accessibility-semantics:dom-ariamixin-role x-internal="dom-ariamixin-role"}` [ = ``value`{.variable}` ]`
:   Sets or retrieves the default ARIA role for
    `internals`{.variable}\'s [target
    element](#internals-target){#accessibility-semantics:internals-target},
    which will be used unless the page author overrides it using the
    [`role`{#accessibility-semantics:attr-aria-role}](infrastructure.html#attr-aria-role)
    attribute.

`internals`{.variable}`.`[`aria*`](https://w3c.github.io/aria/#idl-def-ariamixin-ariaactivedescendantelement){#accessibility-semantics:dom-ariamixin-aria* x-internal="dom-ariamixin-aria*"}` [ = ``value`{.variable}` ]`

:   Sets or retrieves various default ARIA states or property values for
    `internals`{.variable}\'s [target
    element](#internals-target){#accessibility-semantics:internals-target-2},
    which will be used unless the page author overrides them using the
    [`aria-*`{#accessibility-semantics:attr-aria-*}](infrastructure.html#attr-aria-*)
    attributes.

Each [custom
element](#custom-element){#accessibility-semantics:custom-element} has
an [internal content attribute map]{#internal-content-attribute-map
.dfn}, which is a
[map](https://infra.spec.whatwg.org/#ordered-map){#accessibility-semantics:ordered-map
x-internal="ordered-map"}, initially empty. See the [Requirements
related to ARIA and to platform accessibility APIs](dom.html#wai-aria)
section for information on how this impacts platform accessibility APIs.

##### [4.13.7.5]{.secno} Custom state pseudo-class[](#custom-state-pseudo-class){.self-link}

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states}`.add(``value`{.variable}`)`

:   Adds the string `value`{.variable} to the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set} to be
    exposed as a pseudo-class.

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-2}`.has(``value`{.variable}`)`

:   Returns true if `value`{.variable} is in the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-2},
    otherwise false.

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-3}`.delete(``value`{.variable}`)`

:   If the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-3} has
    `value`{.variable}, then it will be removed and true will be
    returned. Otherwise, false will be returned.

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-4}`.clear()`

:   Removes all values from the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-4}.

`for (const ``stateName`{.variable}` of ``internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-5}`)`\
`for (const ``stateName`{.variable}` of ``internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-6}`.entries())`\
`for (const ``stateName`{.variable}` of ``internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-7}`.keys())`\
`for (const ``stateName`{.variable}` of ``internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-8}`.values())`

:   Iterates over all values in the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-5}.

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-9}`.forEach(``callback`{.variable}`)`

:   Iterates over all values in the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-6} by
    calling `callback`{.variable} once for each value.

`internals`{.variable}`.`[`states`](#dom-elementinternals-states){#custom-state-pseudo-class:dom-elementinternals-states-10}`.size`

:   Returns the number of values in the element\'s [states
    set](#states-set){#custom-state-pseudo-class:states-set-7}.

Each [custom
element](#custom-element){#custom-state-pseudo-class:custom-element} has
a [states set]{#states-set .dfn}, which is a
[`CustomStateSet`{#custom-state-pseudo-class:customstateset}](#customstateset),
initially empty.

``` idl
[Exposed=Window]
interface CustomStateSet {
  setlike<DOMString>;
};
```

The [`states`]{#dom-elementinternals-states .dfn dfn-for="HTMLElement"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#custom-state-pseudo-class:this
x-internal="this"}\'s [target
element](#internals-target){#custom-state-pseudo-class:internals-target}\'s
[states set](#states-set){#custom-state-pseudo-class:states-set-8}.

::: example
The [states set](#states-set){#custom-state-pseudo-class:states-set-9}
can expose boolean states represented by existence/non-existence of
string values. If an author wants to expose a state which can have three
values, it can be converted to three exclusive boolean states. For
example, a state called `readyState` with `"loading"`, `"interactive"`,
and `"complete"` values can be mapped to three exclusive boolean states,
`"loading"`, `"interactive"`, and `"complete"`:

``` js
// Change the readyState from anything to "complete".
this._readyState = "complete";
this._internals.states.delete("loading");
this._internals.states.delete("interactive");
this._internals.states.add("complete");
```
:::

[← 4.12.5 The canvas element](canvas.html) --- [Table of Contents](./)
--- [4.14 Common idioms without dedicated elements
→](semantics-other.html)
