HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.9 Tabular data](tables.html) --- [Table of Contents](./) ---
[4.10.5 The input element →](input.html)

1.  ::: {#toc-semantics}
    1.  [[4.10]{.secno} Forms](forms.html#forms)
        1.  [[4.10.1]{.secno} Introduction](forms.html#introduction-4)
            1.  [[4.10.1.1]{.secno} Writing a form\'s user
                interface](forms.html#writing-a-form's-user-interface)
            2.  [[4.10.1.2]{.secno} Implementing the server-side
                processing for a
                form](forms.html#implementing-the-server-side-processing-for-a-form)
            3.  [[4.10.1.3]{.secno} Configuring a form to communicate
                with a
                server](forms.html#configuring-a-form-to-communicate-with-a-server)
            4.  [[4.10.1.4]{.secno} Client-side form
                validation](forms.html#client-side-form-validation)
            5.  [[4.10.1.5]{.secno} Enabling client-side automatic
                filling of form
                controls](forms.html#enabling-client-side-automatic-filling-of-form-controls)
            6.  [[4.10.1.6]{.secno} Improving the user experience on
                mobile
                devices](forms.html#improving-the-user-experience-on-mobile-devices)
            7.  [[4.10.1.7]{.secno} The difference between the field
                type, the autofill field name, and the input
                modality](forms.html#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality)
            8.  [[4.10.1.8]{.secno} Date, time, and number
                formats](forms.html#input-author-notes)
        2.  [[4.10.2]{.secno} Categories](forms.html#categories)
        3.  [[4.10.3]{.secno} The `form`
            element](forms.html#the-form-element)
        4.  [[4.10.4]{.secno} The `label`
            element](forms.html#the-label-element)
    :::

### [4.10]{.secno} Forms[](#forms){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element#Forms](https://developer.mozilla.org/en-US/docs/Web/HTML/Element#Forms "This page lists all the HTML elements, which are created using tags.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome61+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera52+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)16+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android5+]{.firefox_android .yes}[Safari iOS3.2+]{.safari_ios
.yes}[Chrome Android61+]{.chrome_android .yes}[WebView
Android61+]{.webview_android .yes}[Samsung
Internet8.0+]{.samsunginternet_android .yes}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

#### [4.10.1]{.secno} Introduction[](#introduction-4){.self-link} {#introduction-4}

*This section is non-normative.*

A form is a component of a web page that has form controls, such as
text, buttons, checkboxes, range, or color picker controls. A user can
interact with such a form, providing data that can then be sent to the
server for further processing (e.g. returning the results of a search or
calculation). No client-side scripting is needed in many cases, though
an API is available so that scripts can augment the user experience or
use forms for purposes other than submitting data to a server.

Writing a form consists of several steps, which can be performed in any
order: writing the user interface, implementing the server-side
processing, and configuring the user interface to communicate with the
server.

##### [4.10.1.1]{.secno} Writing a form\'s user interface[](#writing-a-form's-user-interface){.self-link} {#writing-a-form's-user-interface}

*This section is non-normative.*

For the purposes of this brief introduction, we will create a pizza
ordering form.

Any form starts with a
[`form`{#writing-a-form's-user-interface:the-form-element}](#the-form-element)
element, inside which are placed the controls. Most controls are
represented by the
[`input`{#writing-a-form's-user-interface:the-input-element}](input.html#the-input-element)
element, which by default provides a text control. To label a control,
the
[`label`{#writing-a-form's-user-interface:the-label-element}](#the-label-element)
element is used; the label text and the control itself go inside the
[`label`{#writing-a-form's-user-interface:the-label-element-2}](#the-label-element)
element. Each part of a form is considered a
[paragraph](dom.html#paragraph){#writing-a-form's-user-interface:paragraph},
and is typically separated from other parts using
[`p`{#writing-a-form's-user-interface:the-p-element}](grouping-content.html#the-p-element)
elements. Putting this together, here is how one might ask for the
customer\'s name:

``` html
<form>
 <p><label>Customer name: <input></label></p>
</form>
```

To let the user select the size of the pizza, we can use a set of radio
buttons. Radio buttons also use the
[`input`{#writing-a-form's-user-interface:the-input-element-2}](input.html#the-input-element)
element, this time with a
[`type`{#writing-a-form's-user-interface:attr-input-type}](input.html#attr-input-type)
attribute with the value
[`radio`{#writing-a-form's-user-interface:radio-button-state-(type=radio)}](input.html#radio-button-state-(type=radio)).
To make the radio buttons work as a group, they are given a common name
using the
[`name`{#writing-a-form's-user-interface:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
attribute. To group a batch of controls together, such as, in this case,
the radio buttons, one can use the
[`fieldset`{#writing-a-form's-user-interface:the-fieldset-element}](form-elements.html#the-fieldset-element)
element. The title of such a group of controls is given by the first
element in the
[`fieldset`{#writing-a-form's-user-interface:the-fieldset-element-2}](form-elements.html#the-fieldset-element),
which has to be a
[`legend`{#writing-a-form's-user-interface:the-legend-element}](form-elements.html#the-legend-element)
element.

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
</form>
```

Changes from the previous step are highlighted.

To pick toppings, we can use checkboxes. These use the
[`input`{#writing-a-form's-user-interface:the-input-element-3}](input.html#the-input-element)
element with a
[`type`{#writing-a-form's-user-interface:attr-input-type-2}](input.html#attr-input-type)
attribute with the value
[`checkbox`{#writing-a-form's-user-interface:checkbox-state-(type=checkbox)}](input.html#checkbox-state-(type=checkbox)):

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox> Bacon </label></p>
  <p><label> <input type=checkbox> Extra Cheese </label></p>
  <p><label> <input type=checkbox> Onion </label></p>
  <p><label> <input type=checkbox> Mushroom </label></p>
 </fieldset>
</form>
```

The pizzeria for which this form is being written is always making
mistakes, so it needs a way to contact the customer. For this purpose,
we can use form controls specifically for telephone numbers
([`input`{#writing-a-form's-user-interface:the-input-element-4}](input.html#the-input-element)
elements with their
[`type`{#writing-a-form's-user-interface:attr-input-type-3}](input.html#attr-input-type)
attribute set to
[`tel`{#writing-a-form's-user-interface:telephone-state-(type=tel)}](input.html#telephone-state-(type=tel)))
and email addresses
([`input`{#writing-a-form's-user-interface:the-input-element-5}](input.html#the-input-element)
elements with their
[`type`{#writing-a-form's-user-interface:attr-input-type-4}](input.html#attr-input-type)
attribute set to
[`email`{#writing-a-form's-user-interface:email-state-(type=email)}](input.html#email-state-(type=email))):

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <p><label>Telephone: <input type=tel></label></p>
 <p><label>Email address: <input type=email></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox> Bacon </label></p>
  <p><label> <input type=checkbox> Extra Cheese </label></p>
  <p><label> <input type=checkbox> Onion </label></p>
  <p><label> <input type=checkbox> Mushroom </label></p>
 </fieldset>
</form>
```

We can use an
[`input`{#writing-a-form's-user-interface:the-input-element-6}](input.html#the-input-element)
element with its
[`type`{#writing-a-form's-user-interface:attr-input-type-5}](input.html#attr-input-type)
attribute set to
[`time`{#writing-a-form's-user-interface:time-state-(type=time)}](input.html#time-state-(type=time))
to ask for a delivery time. Many of these form controls have attributes
to control exactly what values can be specified; in this case, three
attributes of particular interest are
[`min`{#writing-a-form's-user-interface:attr-input-min}](input.html#attr-input-min),
[`max`{#writing-a-form's-user-interface:attr-input-max}](input.html#attr-input-max),
and
[`step`{#writing-a-form's-user-interface:attr-input-step}](input.html#attr-input-step).
These set the minimum time, the maximum time, and the interval between
allowed values (in seconds). This pizzeria only delivers between 11am
and 9pm, and doesn\'t promise anything better than 15 minute increments,
which we can mark up as follows:

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <p><label>Telephone: <input type=tel></label></p>
 <p><label>Email address: <input type=email></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox> Bacon </label></p>
  <p><label> <input type=checkbox> Extra Cheese </label></p>
  <p><label> <input type=checkbox> Onion </label></p>
  <p><label> <input type=checkbox> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900"></label></p>
</form>
```

The
[`textarea`{#writing-a-form's-user-interface:the-textarea-element}](form-elements.html#the-textarea-element)
element can be used to provide a multiline text control. In this
instance, we are going to use it to provide a space for the customer to
give delivery instructions:

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <p><label>Telephone: <input type=tel></label></p>
 <p><label>Email address: <input type=email></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox> Bacon </label></p>
  <p><label> <input type=checkbox> Extra Cheese </label></p>
  <p><label> <input type=checkbox> Onion </label></p>
  <p><label> <input type=checkbox> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900"></label></p>
 <p><label>Delivery instructions: <textarea></textarea></label></p>
</form>
```

Finally, to make the form submittable we use the
[`button`{#writing-a-form's-user-interface:the-button-element}](form-elements.html#the-button-element)
element:

``` html
<form>
 <p><label>Customer name: <input></label></p>
 <p><label>Telephone: <input type=tel></label></p>
 <p><label>Email address: <input type=email></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size> Small </label></p>
  <p><label> <input type=radio name=size> Medium </label></p>
  <p><label> <input type=radio name=size> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox> Bacon </label></p>
  <p><label> <input type=checkbox> Extra Cheese </label></p>
  <p><label> <input type=checkbox> Onion </label></p>
  <p><label> <input type=checkbox> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900"></label></p>
 <p><label>Delivery instructions: <textarea></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

##### [4.10.1.2]{.secno} Implementing the server-side processing for a form[](#implementing-the-server-side-processing-for-a-form){.self-link}

*This section is non-normative.*

The exact details for writing a server-side processor are out of scope
for this specification. For the purposes of this introduction, we will
assume that the script at `https://pizza.example.com/order.cgi` is
configured to accept submissions using the
[`application/x-www-form-urlencoded`{#implementing-the-server-side-processing-for-a-form:attr-fs-enctype-urlencoded}](form-control-infrastructure.html#attr-fs-enctype-urlencoded)
format, expecting the following parameters sent in an HTTP POST body:

`custname`
:   Customer\'s name

`custtel`
:   Customer\'s telephone number

`custemail`
:   Customer\'s email address

`size`
:   The pizza size, either `small`, `medium`, or `large`

`topping`
:   A topping, specified once for each selected topping, with the
    allowed values being `bacon`, `cheese`, `onion`, and `mushroom`

`delivery`
:   The requested delivery time

`comments`
:   The delivery instructions

##### [4.10.1.3]{.secno} Configuring a form to communicate with a server[](#configuring-a-form-to-communicate-with-a-server){.self-link}

*This section is non-normative.*

Form submissions are exposed to servers in a variety of ways, most
commonly as HTTP GET or POST requests. To specify the exact method used,
the
[`method`{#configuring-a-form-to-communicate-with-a-server:attr-fs-method}](form-control-infrastructure.html#attr-fs-method)
attribute is specified on the
[`form`{#configuring-a-form-to-communicate-with-a-server:the-form-element}](#the-form-element)
element. This doesn\'t specify how the form data is encoded, though; to
specify that, you use the
[`enctype`{#configuring-a-form-to-communicate-with-a-server:attr-fs-enctype}](form-control-infrastructure.html#attr-fs-enctype)
attribute. You also have to specify the
[URL](https://url.spec.whatwg.org/#concept-url){#configuring-a-form-to-communicate-with-a-server:url
x-internal="url"} of the service that will handle the submitted data,
using the
[`action`{#configuring-a-form-to-communicate-with-a-server:attr-fs-action}](form-control-infrastructure.html#attr-fs-action)
attribute.

For each form control you want submitted, you then have to give a name
that will be used to refer to the data in the submission. We already
specified the name for the group of radio buttons; the same attribute
([`name`{#configuring-a-form-to-communicate-with-a-server:attr-fe-name}](form-control-infrastructure.html#attr-fe-name))
also specifies the submission name. Radio buttons can be distinguished
from each other in the submission by giving them different values, using
the
[`value`{#configuring-a-form-to-communicate-with-a-server:attr-input-value}](input.html#attr-input-value)
attribute.

Multiple controls can have the same name; for example, here we give all
the checkboxes the same name, and the server distinguishes which
checkbox was checked by seeing which values are submitted with that name
--- like the radio buttons, they are also given unique values with the
[`value`{#configuring-a-form-to-communicate-with-a-server:attr-input-value-2}](input.html#attr-input-value)
attribute.

Given the settings in the previous section, this all becomes:

``` html
<form method="post"
      enctype="application/x-www-form-urlencoded"
      action="https://pizza.example.com/order.cgi">
 <p><label>Customer name: <input name="custname"></label></p>
 <p><label>Telephone: <input type=tel name="custtel"></label></p>
 <p><label>Email address: <input type=email name="custemail"></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size value="small"> Small </label></p>
  <p><label> <input type=radio name=size value="medium"> Medium </label></p>
  <p><label> <input type=radio name=size value="large"> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox name="topping" value="bacon"> Bacon </label></p>
  <p><label> <input type=checkbox name="topping" value="cheese"> Extra Cheese </label></p>
  <p><label> <input type=checkbox name="topping" value="onion"> Onion </label></p>
  <p><label> <input type=checkbox name="topping" value="mushroom"> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900" name="delivery"></label></p>
 <p><label>Delivery instructions: <textarea name="comments"></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

There is no particular significance to the way some of the attributes
have their values quoted and others don\'t. The HTML syntax allows a
variety of equally valid ways to specify attributes, as discussed [in
the syntax
section](syntax.html#syntax-attributes){#configuring-a-form-to-communicate-with-a-server:syntax-attributes}.

For example, if the customer entered \"Denise Lawrence\" as their name,
\"555-321-8642\" as their telephone number, did not specify an email
address, asked for a medium-sized pizza, selected the Extra Cheese and
Mushroom toppings, entered a delivery time of 7pm, and left the delivery
instructions text control blank, the user agent would submit the
following to the online web service:

``` html
custname=Denise+Lawrence&custtel=555-321-8642&custemail=&size=medium&topping=cheese&topping=mushroom&delivery=19%3A00&comments=
```

##### [4.10.1.4]{.secno} Client-side form validation[](#client-side-form-validation){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Form_validation](https://developer.mozilla.org/en-US/docs/Web/Forms/Form_validation "Client-side form validation sometimes requires JavaScript if you want to customize styling and error messages, but it always requires you to think carefully about the user. Always remember to help your users correct the data they provide. To that end, be sure to:")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}
:::
::::
:::::

*This section is non-normative.*

Forms can be annotated in such a way that the user agent will check the
user\'s input before the form is submitted. The server still has to
verify the input is valid (since hostile users can easily bypass the
form validation), but it allows the user to avoid the wait incurred by
having the server be the sole checker of the user\'s input.

The simplest annotation is the
[`required`{#client-side-form-validation:attr-input-required}](input.html#attr-input-required)
attribute, which can be specified on
[`input`{#client-side-form-validation:the-input-element}](input.html#the-input-element)
elements to indicate that the form is not to be submitted until a value
is given. By adding this attribute to the customer name, pizza size, and
delivery time fields, we allow the user agent to notify the user when
the user submits the form without filling in those fields:

``` html
<form method="post"
      enctype="application/x-www-form-urlencoded"
      action="https://pizza.example.com/order.cgi">
 <p><label>Customer name: <input name="custname" required></label></p>
 <p><label>Telephone: <input type=tel name="custtel"></label></p>
 <p><label>Email address: <input type=email name="custemail"></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size required value="small"> Small </label></p>
  <p><label> <input type=radio name=size required value="medium"> Medium </label></p>
  <p><label> <input type=radio name=size required value="large"> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox name="topping" value="bacon"> Bacon </label></p>
  <p><label> <input type=checkbox name="topping" value="cheese"> Extra Cheese </label></p>
  <p><label> <input type=checkbox name="topping" value="onion"> Onion </label></p>
  <p><label> <input type=checkbox name="topping" value="mushroom"> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900" name="delivery" required></label></p>
 <p><label>Delivery instructions: <textarea name="comments"></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

It is also possible to limit the length of the input, using the
[`maxlength`{#client-side-form-validation:attr-fe-maxlength}](form-control-infrastructure.html#attr-fe-maxlength)
attribute. By adding this to the
[`textarea`{#client-side-form-validation:the-textarea-element}](form-elements.html#the-textarea-element)
element, we can limit users to 1000 characters, preventing them from
writing huge essays to the busy delivery drivers instead of staying
focused and to the point:

``` html
<form method="post"
      enctype="application/x-www-form-urlencoded"
      action="https://pizza.example.com/order.cgi">
 <p><label>Customer name: <input name="custname" required></label></p>
 <p><label>Telephone: <input type=tel name="custtel"></label></p>
 <p><label>Email address: <input type=email name="custemail"></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size required value="small"> Small </label></p>
  <p><label> <input type=radio name=size required value="medium"> Medium </label></p>
  <p><label> <input type=radio name=size required value="large"> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox name="topping" value="bacon"> Bacon </label></p>
  <p><label> <input type=checkbox name="topping" value="cheese"> Extra Cheese </label></p>
  <p><label> <input type=checkbox name="topping" value="onion"> Onion </label></p>
  <p><label> <input type=checkbox name="topping" value="mushroom"> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900" name="delivery" required></label></p>
 <p><label>Delivery instructions: <textarea name="comments" maxlength=1000></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

When a form is submitted,
[`invalid`{#client-side-form-validation:event-invalid}](indices.html#event-invalid)
events are fired at each form control that is invalid. This can be
useful for displaying a summary of the problems with the form, since
typically the browser itself will only report one problem at a time.

##### [4.10.1.5]{.secno} Enabling client-side automatic filling of form controls[](#enabling-client-side-automatic-filling-of-form-controls){.self-link}

*This section is non-normative.*

Some browsers attempt to aid the user by automatically filling form
controls rather than having the user reenter their information each
time. For example, a field asking for the user\'s telephone number can
be automatically filled with the user\'s phone number.

To help the user agent with this, the
[`autocomplete`{#enabling-client-side-automatic-filling-of-form-controls:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute can be used to describe the field\'s purpose. In the case of
this form, we have three fields that can be usefully annotated in this
way: the information about who the pizza is to be delivered to. Adding
this information looks like this:

``` html
<form method="post"
      enctype="application/x-www-form-urlencoded"
      action="https://pizza.example.com/order.cgi">
 <p><label>Customer name: <input name="custname" required autocomplete="shipping name"></label></p>
 <p><label>Telephone: <input type=tel name="custtel" autocomplete="shipping tel"></label></p>
 <p><label>Email address: <input type=email name="custemail" autocomplete="shipping email"></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size required value="small"> Small </label></p>
  <p><label> <input type=radio name=size required value="medium"> Medium </label></p>
  <p><label> <input type=radio name=size required value="large"> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox name="topping" value="bacon"> Bacon </label></p>
  <p><label> <input type=checkbox name="topping" value="cheese"> Extra Cheese </label></p>
  <p><label> <input type=checkbox name="topping" value="onion"> Onion </label></p>
  <p><label> <input type=checkbox name="topping" value="mushroom"> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900" name="delivery" required></label></p>
 <p><label>Delivery instructions: <textarea name="comments" maxlength=1000></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

##### [4.10.1.6]{.secno} Improving the user experience on mobile devices[](#improving-the-user-experience-on-mobile-devices){.self-link}

*This section is non-normative.*

Some devices, in particular those with virtual keyboards can provide the
user with multiple input modalities. For example, when typing in a
credit card number the user may wish to only see keys for digits 0-9,
while when typing in their name they may wish to see a form field that
by default capitalizes each word.

Using the
[`inputmode`{#improving-the-user-experience-on-mobile-devices:attr-inputmode}](interaction.html#attr-inputmode)
attribute we can select appropriate input modalities:

``` html
<form method="post"
      enctype="application/x-www-form-urlencoded"
      action="https://pizza.example.com/order.cgi">
 <p><label>Customer name: <input name="custname" required autocomplete="shipping name"></label></p>
 <p><label>Telephone: <input type=tel name="custtel" autocomplete="shipping tel"></label></p>
 <p><label>Buzzer code: <input name="custbuzz" inputmode="numeric"></label></p>
 <p><label>Email address: <input type=email name="custemail" autocomplete="shipping email"></label></p>
 <fieldset>
  <legend> Pizza Size </legend>
  <p><label> <input type=radio name=size required value="small"> Small </label></p>
  <p><label> <input type=radio name=size required value="medium"> Medium </label></p>
  <p><label> <input type=radio name=size required value="large"> Large </label></p>
 </fieldset>
 <fieldset>
  <legend> Pizza Toppings </legend>
  <p><label> <input type=checkbox name="topping" value="bacon"> Bacon </label></p>
  <p><label> <input type=checkbox name="topping" value="cheese"> Extra Cheese </label></p>
  <p><label> <input type=checkbox name="topping" value="onion"> Onion </label></p>
  <p><label> <input type=checkbox name="topping" value="mushroom"> Mushroom </label></p>
 </fieldset>
 <p><label>Preferred delivery time: <input type=time min="11:00" max="21:00" step="900" name="delivery" required></label></p>
 <p><label>Delivery instructions: <textarea name="comments" maxlength=1000></textarea></label></p>
 <p><button>Submit order</button></p>
</form>
```

##### [4.10.1.7]{.secno} The difference between the field type, the autofill field name, and the input modality[](#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality){.self-link} {#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality}

*This section is non-normative.*

The
[`type`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-input-type}](input.html#attr-input-type),
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
and
[`inputmode`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-inputmode}](interaction.html#attr-inputmode)
attributes can seem confusingly similar. For instance, in all three
cases, the string \"`email`\" is a valid value. This section attempts to
illustrate the difference between the three attributes and provides
advice suggesting how to use them.

The
[`type`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-input-type-2}](input.html#attr-input-type)
attribute on
[`input`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:the-input-element}](input.html#the-input-element)
elements decides what kind of control the user agent will use to expose
the field. Choosing between different values of this attribute is the
same choice as choosing whether to use an
[`input`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:the-input-element-2}](input.html#the-input-element)
element, a
[`textarea`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:the-textarea-element}](form-elements.html#the-textarea-element)
element, a
[`select`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:the-select-element}](form-elements.html#the-select-element)
element, etc.

The
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-2}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute, in contrast, describes what the value that the user will
enter actually represents. Choosing between different values of this
attribute is the same choice as choosing what the label for the element
will be.

First, consider telephone numbers. If a page is asking for a telephone
number from the user, the right form control to use is
[`<input type=tel>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:telephone-state-(type=tel)}](input.html#telephone-state-(type=tel)).
However, which
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-3}](form-control-infrastructure.html#attr-fe-autocomplete)
value to use depends on which phone number the page is asking for,
whether they expect a telephone number in the international format or
just the local format, and so forth.

For example, a page that forms part of a checkout process on an
e-commerce site for a customer buying a gift to be shipped to a friend
might need both the buyer\'s telephone number (in case of payment
issues) and the friend\'s telephone number (in case of delivery issues).
If the site expects international phone numbers (with the country code
prefix), this could thus look like this:

``` html
<p><label>Your phone number: <input type=tel name=custtel autocomplete="billing tel"></label>
<p><label>Recipient's phone number: <input type=tel name=shiptel autocomplete="shipping tel"></label>
<p>Please enter complete phone numbers including the country code prefix, as in "+1 555 123 4567".
```

But if the site only supports British customers and recipients, it might
instead look like this (notice the use of
[`tel-national`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-tel-national}](form-control-infrastructure.html#attr-fe-autocomplete-tel-national)
rather than
[`tel`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-tel}](form-control-infrastructure.html#attr-fe-autocomplete-tel)):

``` html
<p><label>Your phone number: <input type=tel name=custtel autocomplete="billing tel-national"></label>
<p><label>Recipient's phone number: <input type=tel name=shiptel autocomplete="shipping tel-national"></label>
<p>Please enter complete UK phone numbers, as in "(01632) 960 123".
```

Now, consider a person\'s preferred languages. The right
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-4}](form-control-infrastructure.html#attr-fe-autocomplete)
value is
[`language`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-language}](form-control-infrastructure.html#attr-fe-autocomplete-language).
However, there could be a number of different form controls used for the
purpose: a text control
([`<input type=text>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:text-(type=text)-state-and-search-state-(type=search)}](input.html#text-(type=text)-state-and-search-state-(type=search))),
a drop-down list
([`<select>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:the-select-element-2}](form-elements.html#the-select-element)),
radio buttons
([`<input type=radio>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:radio-button-state-(type=radio)}](input.html#radio-button-state-(type=radio))),
etc. It only depends on what kind of interface is desired.

Finally, consider names. If a page just wants one name from the user,
then the relevant control is
[`<input type=text>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:text-(type=text)-state-and-search-state-(type=search)-2}](input.html#text-(type=text)-state-and-search-state-(type=search)).
If the page is asking for the user\'s full name, then the relevant
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-5}](form-control-infrastructure.html#attr-fe-autocomplete)
value is
[`name`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-name}](form-control-infrastructure.html#attr-fe-autocomplete-name).

``` html
<p><label>Japanese name: <input name="j" type="text" autocomplete="section-jp name"></label>
<label>Romanized name: <input name="e" type="text" autocomplete="section-en name"></label>
```

In this example, the
\"[`section-*`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-section}](form-control-infrastructure.html#attr-fe-autocomplete-section)\"
keywords in the
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-6}](form-control-infrastructure.html#attr-fe-autocomplete)
attributes\' values tell the user agent that the two fields expect
*different* names. Without them, the user agent could automatically fill
the second field with the value given in the first field when the user
gave a value to the first field.

The \"`-jp`\" and \"`-en`\" parts of the keywords are opaque to the user
agent; the user agent cannot guess, from those, that the two names are
expected to be in Japanese and English respectively.

Separate from the choices regarding
[`type`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-input-type-3}](input.html#attr-input-type)
and
[`autocomplete`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-fe-autocomplete-7}](form-control-infrastructure.html#attr-fe-autocomplete),
the
[`inputmode`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:attr-inputmode-2}](interaction.html#attr-inputmode)
attribute decides what kind of input modality (e.g., virtual keyboard)
to use, when the control is a text control.

Consider credit card numbers. The appropriate input type is *not*
[`<input type=number>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:number-state-(type=number)}](input.html#number-state-(type=number)),
[as explained below](input.html#when-number-is-not-appropriate); it is
instead
[`<input type=text>`{#the-difference-between-the-field-type,-the-autofill-field-name,-and-the-input-modality:text-(type=text)-state-and-search-state-(type=search)-3}](input.html#text-(type=text)-state-and-search-state-(type=search)).
To encourage the user agent to use a numeric input modality anyway
(e.g., a virtual keyboard displaying only digits), the page would use

``` html
<p><label>Credit card number:
                <input name="cc" type="text" inputmode="numeric" pattern="[0-9]{8,19}" autocomplete="cc-number">
</label></p>
```

##### [4.10.1.8]{.secno} Date, time, and number formats[](#input-author-notes){.self-link} {#input-author-notes}

*This section is non-normative.*

In this pizza delivery example, the times are specified in the format
\"HH:MM\": two digits for the hour, in 24-hour format, and two digits
for the time. (Seconds could also be specified, though they are not
necessary in this example.)

In some locales, however, times are often expressed differently when
presented to users. For example, in the United States, it is still
common to use the 12-hour clock with an am/pm indicator, as in \"2pm\".
In France, it is common to separate the hours from the minutes using an
\"h\" character, as in \"14h00\".

Similar issues exist with dates, with the added complication that even
the order of the components is not always consistent --- for example, in
Cyprus the first of February 2003 would typically be written \"1/2/03\",
while that same date in Japan would typically be written as
\"2003年02月01日\" --- and even with numbers, where locales differ, for
example, in what punctuation is used as the decimal separator and the
thousands separator.

It is therefore important to distinguish the time, date, and number
formats used in HTML and in form submissions, which are always the
formats defined in this specification (and based on the well-established
ISO 8601 standard for computer-readable date and time formats), from the
time, date, and number formats presented to the user by the browser and
accepted as input from the user by the browser.

The format used \"on the wire\", i.e., in HTML markup and in form
submissions, is intended to be computer-readable and consistent
irrespective of the user\'s locale. Dates, for instance, are always
written in the format \"YYYY-MM-DD\", as in \"2003-02-01\". While some
users might see this format, others might see it as \"01.02.2003\" or
\"February 1, 2003\".

The time, date, or number given by the page in the wire format is then
translated to the user\'s preferred presentation (based on user
preferences or on the locale of the page itself), before being displayed
to the user. Similarly, after the user inputs a time, date, or number
using their preferred format, the user agent converts it back to the
wire format before putting it in the DOM or submitting it.

This allows scripts in pages and on servers to process times, dates, and
numbers in a consistent manner without needing to support dozens of
different formats, while still supporting the users\' needs.

See also the [implementation notes](input.html#input-impl-notes)
regarding localization of form controls.

#### [4.10.2]{.secno} Categories[](#categories){.self-link}

Mostly for historical reasons, elements in this section fall into
several overlapping (but subtly different) categories in addition to the
usual ones like [flow
content](dom.html#flow-content-2){#categories:flow-content-2}, [phrasing
content](dom.html#phrasing-content-2){#categories:phrasing-content-2},
and [interactive
content](dom.html#interactive-content-2){#categories:interactive-content-2}.

A number of the elements are [form-associated
elements]{#form-associated-element .dfn}, which means they can have a
[form
owner](form-control-infrastructure.html#form-owner){#categories:form-owner}.

- [`button`{#categories:the-button-element}](form-elements.html#the-button-element)
- [`fieldset`{#categories:the-fieldset-element}](form-elements.html#the-fieldset-element)
- [`input`{#categories:the-input-element}](input.html#the-input-element)
- [`object`{#categories:the-object-element}](iframe-embed-object.html#the-object-element)
- [`output`{#categories:the-output-element}](form-elements.html#the-output-element)
- [`select`{#categories:the-select-element}](form-elements.html#the-select-element)
- [`textarea`{#categories:the-textarea-element}](form-elements.html#the-textarea-element)
- [`img`{#categories:the-img-element}](embedded-content.html#the-img-element)
- [form-associated custom
  elements](custom-elements.html#form-associated-custom-element){#categories:form-associated-custom-element}

The [form-associated
elements](#form-associated-element){#categories:form-associated-element}
fall into several subcategories:

[Listed elements]{#category-listed .dfn}

:   Denotes elements that are listed in the
    [`form`{.variable}`.elements`{#categories:dom-form-elements}](#dom-form-elements)
    and
    [`fieldset`{.variable}`.elements`{#categories:dom-fieldset-elements}](form-elements.html#dom-fieldset-elements)
    APIs. These elements also have a
    [`form`{#categories:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    content attribute, and a matching
    [`form`{#categories:dom-fae-form}](form-control-infrastructure.html#dom-fae-form)
    IDL attribute, that allow authors to specify an explicit [form
    owner](form-control-infrastructure.html#form-owner){#categories:form-owner-2}.

    - [`button`{#categories:the-button-element-2}](form-elements.html#the-button-element)
    - [`fieldset`{#categories:the-fieldset-element-2}](form-elements.html#the-fieldset-element)
    - [`input`{#categories:the-input-element-2}](input.html#the-input-element)
    - [`object`{#categories:the-object-element-2}](iframe-embed-object.html#the-object-element)
    - [`output`{#categories:the-output-element-2}](form-elements.html#the-output-element)
    - [`select`{#categories:the-select-element-2}](form-elements.html#the-select-element)
    - [`textarea`{#categories:the-textarea-element-2}](form-elements.html#the-textarea-element)
    - [form-associated custom
      elements](custom-elements.html#form-associated-custom-element){#categories:form-associated-custom-element-2}

[Submittable elements]{#category-submit .dfn}

:   Denotes elements that can be used for [constructing the entry
    list](form-control-infrastructure.html#constructing-the-form-data-set){#categories:constructing-the-form-data-set}
    when a [`form`{#categories:the-form-element}](#the-form-element)
    element is
    [submitted](form-control-infrastructure.html#concept-form-submit){#categories:concept-form-submit}.

    - [`button`{#categories:the-button-element-3}](form-elements.html#the-button-element)
    - [`input`{#categories:the-input-element-3}](input.html#the-input-element)
    - [`select`{#categories:the-select-element-3}](form-elements.html#the-select-element)
    - [`textarea`{#categories:the-textarea-element-3}](form-elements.html#the-textarea-element)
    - [form-associated custom
      elements](custom-elements.html#form-associated-custom-element){#categories:form-associated-custom-element-3}

    Some [submittable
    elements](#category-submit){#categories:category-submit} can be,
    depending on their attributes, [buttons]{#concept-button .dfn}. The
    prose below defines when an element is a button. Some buttons are
    specifically [submit buttons]{#concept-submit-button .dfn
    lt="submit button" export=""}.

[Resettable elements]{#category-reset .dfn}

:   Denotes elements that can be affected when a
    [`form`{#categories:the-form-element-2}](#the-form-element) element
    is
    [reset](form-control-infrastructure.html#concept-form-reset){#categories:concept-form-reset}.

    - [`input`{#categories:the-input-element-4}](input.html#the-input-element)
    - [`output`{#categories:the-output-element-3}](form-elements.html#the-output-element)
    - [`select`{#categories:the-select-element-4}](form-elements.html#the-select-element)
    - [`textarea`{#categories:the-textarea-element-4}](form-elements.html#the-textarea-element)
    - [form-associated custom
      elements](custom-elements.html#form-associated-custom-element){#categories:form-associated-custom-element-4}

[Autocapitalize-and-autocorrect-inheriting elements]{#category-autocapitalize .dfn}

:   Denotes elements that inherit the
    [`autocapitalize`{#categories:attr-autocapitalize}](interaction.html#attr-autocapitalize)
    and
    [`autocorrect`{#categories:attr-autocorrect}](interaction.html#attr-autocorrect)
    attributes from their [form
    owner](form-control-infrastructure.html#form-owner){#categories:form-owner-3}.

    - [`button`{#categories:the-button-element-4}](form-elements.html#the-button-element)
    - [`fieldset`{#categories:the-fieldset-element-3}](form-elements.html#the-fieldset-element)
    - [`input`{#categories:the-input-element-5}](input.html#the-input-element)
    - [`output`{#categories:the-output-element-4}](form-elements.html#the-output-element)
    - [`select`{#categories:the-select-element-5}](form-elements.html#the-select-element)
    - [`textarea`{#categories:the-textarea-element-5}](form-elements.html#the-textarea-element)

Some elements, not all of them
[form-associated](#form-associated-element){#categories:form-associated-element-2},
are categorized as [labelable elements]{#category-label .dfn}. These are
elements that can be associated with a
[`label`{#categories:the-label-element}](#the-label-element) element.

- [`button`{#categories:the-button-element-5}](form-elements.html#the-button-element)
- [`input`{#categories:the-input-element-6}](input.html#the-input-element)
  (if the
  [`type`{#categories:attr-input-type}](input.html#attr-input-type)
  attribute is *not* in the
  [Hidden](input.html#hidden-state-(type=hidden)){#categories:hidden-state-(type=hidden)}
  state)
- [`meter`{#categories:the-meter-element}](form-elements.html#the-meter-element)
- [`output`{#categories:the-output-element-5}](form-elements.html#the-output-element)
- [`progress`{#categories:the-progress-element}](form-elements.html#the-progress-element)
- [`select`{#categories:the-select-element-6}](form-elements.html#the-select-element)
- [`textarea`{#categories:the-textarea-element-6}](form-elements.html#the-textarea-element)
- [form-associated custom
  elements](custom-elements.html#form-associated-custom-element){#categories:form-associated-custom-element-5}

#### [4.10.3]{.secno} The [`form`]{.dfn dfn-type="element"} element[](#the-form-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/form](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form "The <form> HTML element represents a document section containing interactive controls for submitting information.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement "The HTMLFormElement interface represents a <form> element in the DOM. It allows access to—and, in some cases, modification of—aspects of the form, as well as access to its component elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-form-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-form-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-form-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-form-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-form-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-form-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-form-element:flow-content-2-3},
    but with no
    [`form`{#the-form-element:the-form-element}](#the-form-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-form-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-form-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-form-element:global-attributes}
:   [`accept-charset`{#the-form-element:attr-form-accept-charset}](#attr-form-accept-charset)
    --- Character encodings to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2}
:   [`action`{#the-form-element:attr-fs-action}](form-control-infrastructure.html#attr-fs-action)
    ---
    [URL](https://url.spec.whatwg.org/#concept-url){#the-form-element:url
    x-internal="url"} to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2-2}
:   [`autocomplete`{#the-form-element:attr-form-autocomplete}](#attr-form-autocomplete)
    --- Default setting for autofill feature for controls in the form
:   [`enctype`{#the-form-element:attr-fs-enctype}](form-control-infrastructure.html#attr-fs-enctype)
    --- [Entry
    list](form-control-infrastructure.html#entry-list){#the-form-element:entry-list}
    encoding type to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2-3}
:   [`method`{#the-form-element:attr-fs-method}](form-control-infrastructure.html#attr-fs-method)
    --- Variant to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2-4}
:   [`name`{#the-form-element:attr-form-name}](#attr-form-name) --- Name
    of form to use in the
    [`document.forms`{#the-form-element:dom-document-forms}](dom.html#dom-document-forms)
    API
:   [`novalidate`{#the-form-element:attr-fs-novalidate}](form-control-infrastructure.html#attr-fs-novalidate)
    --- Bypass form control validation for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2-5}
:   [`target`{#the-form-element:attr-fs-target}](form-control-infrastructure.html#attr-fs-target)
    ---
    [Navigable](document-sequences.html#navigable){#the-form-element:navigable}
    for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-form-element:form-submission-2-6}
:   [`rel`{#the-form-element:attr-form-rel}](#attr-form-rel)

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-form-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-form).
:   [For implementers](https://w3c.github.io/html-aam/#el-form).

[DOM interface](dom.html#concept-element-dom){#the-form-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window,
     LegacyOverrideBuiltIns,
     LegacyUnenumerableNamedProperties]
    interface HTMLFormElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString acceptCharset;
      [CEReactions] attribute USVString action;
      [CEReactions] attribute DOMString autocomplete;
      [CEReactions] attribute DOMString enctype;
      [CEReactions] attribute DOMString encoding;
      [CEReactions] attribute DOMString method;
      [CEReactions] attribute DOMString name;
      [CEReactions] attribute boolean noValidate;
      [CEReactions] attribute DOMString target;
      [CEReactions] attribute DOMString rel;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList relList;

      [SameObject] readonly attribute HTMLFormControlsCollection elements;
      readonly attribute unsigned long length;
      getter Element (unsigned long index);
      getter (RadioNodeList or Element) (DOMString name);

      undefined submit();
      undefined requestSubmit(optional HTMLElement? submitter = null);
      [CEReactions] undefined reset();
      boolean checkValidity();
      boolean reportValidity();
    };
    ```

The [`form`{#the-form-element:the-form-element-2}](#the-form-element)
element [represents](dom.html#represents){#the-form-element:represents}
a [hyperlink](links.html#hyperlink){#the-form-element:hyperlink} that
can be manipulated through a collection of [form-associated
elements](#form-associated-element){#the-form-element:form-associated-element},
some of which can represent editable values that can be submitted to a
server for processing.

The [`accept-charset`]{#attr-form-accept-charset .dfn dfn-for="form"
dfn-type="element-attr"} attribute gives the character encodings that
are to be used for the submission. If specified, the value must be an
[ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-form-element:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for \"`UTF-8`\".
[\[ENCODING\]](references.html#refsENCODING)

The [`name`]{#attr-form-name .dfn dfn-for="form"
dfn-type="element-attr"} attribute represents the
[`form`{#the-form-element:the-form-element-3}](#the-form-element)\'s
name within the
[`forms`{#the-form-element:dom-document-forms-2}](dom.html#dom-document-forms)
collection. The value must not be the empty string, and the value must
be unique amongst the
[`form`{#the-form-element:the-form-element-4}](#the-form-element)
elements in the
[`forms`{#the-form-element:dom-document-forms-3}](dom.html#dom-document-forms)
collection that it is in, if any.

The [`autocomplete`]{#attr-form-autocomplete .dfn dfn-for="form"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-form-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`on`]{#attr-form-autocomplete-on .dfn dfn-for="form/autocomplete"
dfn-type="attr-value"}

[On]{#attr-form-autocomplete-on-state .dfn}

Form controls will have their [autofill field
name](form-control-infrastructure.html#autofill-field-name){#the-form-element:autofill-field-name}
set to
\"[`on`{#the-form-element:attr-fe-autocomplete-on}](form-control-infrastructure.html#attr-fe-autocomplete-on)\"
by default.

[`off`]{#attr-form-autocomplete-off .dfn dfn-for="form/autocomplete"
dfn-type="attr-value"}

[Off]{#attr-form-autocomplete-off-state .dfn}

Form controls will have their [autofill field
name](form-control-infrastructure.html#autofill-field-name){#the-form-element:autofill-field-name-2}
set to
\"[`off`{#the-form-element:attr-fe-autocomplete-off}](form-control-infrastructure.html#attr-fe-autocomplete-off)\"
by default.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[On](#attr-form-autocomplete-on-state){#the-form-element:attr-form-autocomplete-on-state}
state.

The
[`action`{#the-form-element:attr-fs-action-2}](form-control-infrastructure.html#attr-fs-action),
[`enctype`{#the-form-element:attr-fs-enctype-2}](form-control-infrastructure.html#attr-fs-enctype),
[`method`{#the-form-element:attr-fs-method-2}](form-control-infrastructure.html#attr-fs-method),
[`novalidate`{#the-form-element:attr-fs-novalidate-2}](form-control-infrastructure.html#attr-fs-novalidate),
and
[`target`{#the-form-element:attr-fs-target-2}](form-control-infrastructure.html#attr-fs-target)
attributes are [attributes for form
submission](form-control-infrastructure.html#attributes-for-form-submission){#the-form-element:attributes-for-form-submission}.

The [`rel`]{#attr-form-rel .dfn dfn-for="form" dfn-type="element-attr"}
attribute on
[`form`{#the-form-element:the-form-element-5}](#the-form-element)
elements controls what kinds of links the elements create. The
attribute\'s value must be a [unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#the-form-element:unordered-set-of-unique-space-separated-tokens}.
The [allowed keywords and their meanings](links.html#linkTypes) are
defined in an earlier section.

[`rel`{#the-form-element:attr-form-rel-2}](#attr-form-rel)\'s [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-form-element:concept-supported-tokens
x-internal="concept-supported-tokens"} are the keywords defined in [HTML
link types](links.html#linkTypes) which are allowed on
[`form`{#the-form-element:the-form-element-6}](#the-form-element)
elements, impact the processing model, and are supported by the user
agent. The possible [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-form-element:concept-supported-tokens-2
x-internal="concept-supported-tokens"} are
[`noreferrer`{#the-form-element:link-type-noreferrer}](links.html#link-type-noreferrer),
[`noopener`{#the-form-element:link-type-noopener}](links.html#link-type-noopener),
and
[`opener`{#the-form-element:link-type-opener}](links.html#link-type-opener).
[`rel`{#the-form-element:attr-form-rel-3}](#attr-form-rel)\'s [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-form-element:concept-supported-tokens-3
x-internal="concept-supported-tokens"} must only include the tokens from
this list that the user agent implements the processing model for.

`form`{.variable}`.`[`elements`](#dom-form-elements){#dom-form-elements-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/elements](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/elements "The HTMLFormElement property elements returns an HTMLFormControlsCollection listing all the form controls contained in the <form> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`HTMLFormControlsCollection`{#the-form-element:htmlformcontrolscollection-2}](common-dom-interfaces.html#htmlformcontrolscollection)
of the form controls in the form (excluding image buttons for historical
reasons).

`form`{.variable}`.`[`length`](#dom-form-length){#dom-form-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/length](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/length "The HTMLFormElement.length read-only property returns the number of controls in the <form> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the number of form controls in the form (excluding image buttons
for historical reasons).

`form`{.variable}`[``index`{.variable}`]`

Returns the `index`{.variable}th element in the form (excluding image
buttons for historical reasons).

`form`{.variable}`[``name`{.variable}`]`

Returns the form control (or, if there are several, a
[`RadioNodeList`{#the-form-element:radionodelist-2}](common-dom-interfaces.html#radionodelist)
of the form controls) in the form with the given
[ID](https://dom.spec.whatwg.org/#concept-id){#the-form-element:concept-id
x-internal="concept-id"} or
[`name`{#the-form-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
(excluding image buttons for historical reasons); or, if there are none,
returns the
[`img`{#the-form-element:the-img-element}](embedded-content.html#the-img-element)
element with the given ID.

Once an element has been referenced using a particular name, that name
will continue being available as a way to reference that element in this
method, even if the element\'s actual
[ID](https://dom.spec.whatwg.org/#concept-id){#the-form-element:concept-id-2
x-internal="concept-id"} or
[`name`{#the-form-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
changes, for as long as the element remains in the
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-form-element:tree
x-internal="tree"}.

If there are multiple matching items, then a
[`RadioNodeList`{#the-form-element:radionodelist-3}](common-dom-interfaces.html#radionodelist)
object containing all those elements is returned.

`form`{.variable}`.`[`submit`](#dom-form-submit){#dom-form-submit-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/submit](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/submit "The HTMLFormElement.submit() method submits a given <form>.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Submits the form, bypassing [interactive constraint
validation](form-control-infrastructure.html#interactively-validate-the-constraints){#the-form-element:interactively-validate-the-constraints}
and without firing a
[`submit`{#the-form-element:event-submit}](indices.html#event-submit)
event.

`form`{.variable}`.`[`requestSubmit`](#dom-form-requestsubmit){#dom-form-requestsubmit-dev}`([ ``submitter`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/requestSubmit](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/requestSubmit "The HTMLFormElement method requestSubmit() requests that the form be submitted using a specific submit button.")

Support in all current engines.

::: support
[Firefox75+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome76+]{.chrome
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

Requests to submit the form. Unlike
[`submit()`{#the-form-element:dom-form-submit-2}](#dom-form-submit),
this method includes [interactive constraint
validation](form-control-infrastructure.html#interactively-validate-the-constraints){#the-form-element:interactively-validate-the-constraints-2}
and firing a
[`submit`{#the-form-element:event-submit-2}](indices.html#event-submit)
event, either of which can cancel submission.

The `submitter`{.variable} argument can be used to point to a specific
[submit
button](#concept-submit-button){#the-form-element:concept-submit-button},
whose
[`formaction`{#the-form-element:attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#the-form-element:attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#the-form-element:attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#the-form-element:attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
and
[`formtarget`{#the-form-element:attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget)
attributes can impact submission. Additionally, the submitter will be
included when [constructing the entry
list](form-control-infrastructure.html#constructing-the-form-data-set){#the-form-element:constructing-the-form-data-set}
for submission; normally, buttons are excluded.

`form`{.variable}`.`[`reset`](#dom-form-reset){#dom-form-reset-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/reset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/reset "The HTMLFormElement.reset() method restores a form element's default values. This method does the same thing as clicking the form's <input type="reset"> control.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Resets the form.

`form`{.variable}`.`[`checkValidity`](#dom-form-checkvalidity){#dom-form-checkvalidity-dev}`()`

Returns true if the form\'s controls are all valid; otherwise, returns
false.

`form`{.variable}`.`[`reportValidity`](#dom-form-reportvalidity){#dom-form-reportvalidity-dev}`()`

Returns true if the form\'s controls are all valid; otherwise, returns
false and informs the user.

The [`autocomplete`]{#dom-form-autocomplete .dfn
dfn-for="HTMLFormElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-form-element:reflect}
the content attribute of the same name, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-form-element:limited-to-only-known-values}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/name](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/name "The HTMLFormElement.name property represents the name of the current <form> element as a string.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`name`]{#dom-form-name .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} and [`rel`]{#dom-form-rel .dfn
dfn-for="HTMLFormElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-form-element:reflect-2}
the content attribute of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/acceptCharset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/acceptCharset "The HTMLFormElement.acceptCharset property represents a list of the supported character encodings for the given <form> element. This list can be comma-separated or space-separated.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`acceptCharset`]{#dom-form-acceptcharset .dfn
dfn-for="HTMLFormElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-form-element:reflect-3}
the
[`accept-charset`{#the-form-element:attr-form-accept-charset-2}](#attr-form-accept-charset)
content attribute.

The [`relList`]{#dom-form-rellist .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-form-element:reflect-4}
the [`rel`{#the-form-element:attr-form-rel-4}](#attr-form-rel) content
attribute.

------------------------------------------------------------------------

The [`elements`]{#dom-form-elements .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must return an
[`HTMLFormControlsCollection`{#the-form-element:htmlformcontrolscollection-3}](common-dom-interfaces.html#htmlformcontrolscollection)
rooted at the
[`form`{#the-form-element:the-form-element-7}](#the-form-element)
element\'s
[root](https://dom.spec.whatwg.org/#concept-tree-root){#the-form-element:root
x-internal="root"}, whose filter matches [listed
elements](#category-listed){#the-form-element:category-listed} whose
[form
owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner}
is the [`form`{#the-form-element:the-form-element-8}](#the-form-element)
element, with the exception of
[`input`{#the-form-element:the-input-element}](input.html#the-input-element)
elements whose
[`type`{#the-form-element:attr-input-type}](input.html#attr-input-type)
attribute is in the [Image
Button](input.html#image-button-state-(type=image)){#the-form-element:image-button-state-(type=image)}
state, which must, for historical reasons, be excluded from this
particular collection.

The [`length`]{#dom-form-length .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must return the number of nodes
[represented](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-form-element:represented-by-the-collection
x-internal="represented-by-the-collection"} by the
[`elements`{#the-form-element:dom-form-elements-2}](#dom-form-elements)
collection.

The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-form-element:supported-property-indices
x-internal="supported-property-indices"} at any instant are the indices
supported by the object returned by the
[`elements`{#the-form-element:dom-form-elements-3}](#dom-form-elements)
attribute at that instant.

To [determine the value of an indexed
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property){#the-form-element:determine-the-value-of-an-indexed-property
x-internal="determine-the-value-of-an-indexed-property"} for a
[`form`{#the-form-element:the-form-element-9}](#the-form-element)
element, the user agent must return the value returned by the
[`item`{#the-form-element:dom-htmlcollection-item}](https://dom.spec.whatwg.org/#dom-htmlcollection-item){x-internal="dom-htmlcollection-item"}
method on the
[`elements`{#the-form-element:dom-form-elements-4}](#dom-form-elements)
collection, when invoked with the given index as its argument.

------------------------------------------------------------------------

Each [`form`{#the-form-element:the-form-element-10}](#the-form-element)
element has a mapping of names to elements called the [past names
map]{#past-names-map .dfn}. It is used to persist names of controls even
when they change names.

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#the-form-element:supported-property-names
x-internal="supported-property-names"} consist of the names obtained
from the following algorithm, in the order obtained from this algorithm:

1.  Let `sourced names`{.variable} be an initially empty ordered list of
    tuples consisting of a string, an element, a source, where the
    source is either *id*, *name*, or *past*, and, if the source is
    *past*, an age.

2.  For each [listed
    element](#category-listed){#the-form-element:category-listed-2}
    `candidate`{.variable} whose [form
    owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-2}
    is the
    [`form`{#the-form-element:the-form-element-11}](#the-form-element)
    element, with the exception of any
    [`input`{#the-form-element:the-input-element-2}](input.html#the-input-element)
    elements whose
    [`type`{#the-form-element:attr-input-type-2}](input.html#attr-input-type)
    attribute is in the [Image
    Button](input.html#image-button-state-(type=image)){#the-form-element:image-button-state-(type=image)-2}
    state:

    1.  If `candidate`{.variable} has an
        [`id`{#the-form-element:the-id-attribute}](dom.html#the-id-attribute)
        attribute, add an entry to `sourced names`{.variable} with that
        [`id`{#the-form-element:the-id-attribute-2}](dom.html#the-id-attribute)
        attribute\'s value as the string, `candidate`{.variable} as the
        element, and *id* as the source.

    2.  If `candidate`{.variable} has a
        [`name`{#the-form-element:attr-fe-name-3}](form-control-infrastructure.html#attr-fe-name)
        attribute, add an entry to `sourced names`{.variable} with that
        [`name`{#the-form-element:attr-fe-name-4}](form-control-infrastructure.html#attr-fe-name)
        attribute\'s value as the string, `candidate`{.variable} as the
        element, and *name* as the source.

3.  For each
    [`img`{#the-form-element:the-img-element-2}](embedded-content.html#the-img-element)
    element `candidate`{.variable} whose [form
    owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-3}
    is the
    [`form`{#the-form-element:the-form-element-12}](#the-form-element)
    element:

    1.  If `candidate`{.variable} has an
        [`id`{#the-form-element:the-id-attribute-3}](dom.html#the-id-attribute)
        attribute, add an entry to `sourced names`{.variable} with that
        [`id`{#the-form-element:the-id-attribute-4}](dom.html#the-id-attribute)
        attribute\'s value as the string, `candidate`{.variable} as the
        element, and *id* as the source.

    2.  If `candidate`{.variable} has a
        [`name`{#the-form-element:attr-img-name}](obsolete.html#attr-img-name)
        attribute, add an entry to `sourced names`{.variable} with that
        [`name`{#the-form-element:attr-img-name-2}](obsolete.html#attr-img-name)
        attribute\'s value as the string, `candidate`{.variable} as the
        element, and *name* as the source.

4.  For each entry `past entry`{.variable} in the [past names
    map](#past-names-map){#the-form-element:past-names-map}, add an
    entry to `sourced names`{.variable} with the
    `past entry`{.variable}\'s name as the string,
    `past entry`{.variable}\'s element as the element, *past* as the
    source, and the length of time `past entry`{.variable} has been in
    the [past names
    map](#past-names-map){#the-form-element:past-names-map-2} as the
    age.

5.  Sort `sourced names`{.variable} by [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-form-element:tree-order
    x-internal="tree-order"} of the element entry of each tuple, sorting
    entries with the same element by putting entries whose source is
    *id* first, then entries whose source is *name*, and finally entries
    whose source is *past*, and sorting entries with the same element
    and source by their age, oldest first.

6.  Remove any entries in `sourced names`{.variable} that have the empty
    string as their name.

7.  Remove any entries in `sourced names`{.variable} that have the same
    name as an earlier entry in the map.

8.  Return the list of names from `sourced names`{.variable},
    maintaining their relative order.

To [determine the value of a named
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-a-named-property){#the-form-element:determine-the-value-of-a-named-property
x-internal="determine-the-value-of-a-named-property"} `name`{.variable}
for a [`form`{#the-form-element:the-form-element-13}](#the-form-element)
element, the user agent must run the following steps:

1.  Let `candidates`{.variable} be a
    [live](infrastructure.html#live){#the-form-element:live}
    [`RadioNodeList`{#the-form-element:radionodelist-4}](common-dom-interfaces.html#radionodelist)
    object containing all the [listed
    elements](#category-listed){#the-form-element:category-listed-3},
    whose [form
    owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-4}
    is the
    [`form`{#the-form-element:the-form-element-14}](#the-form-element)
    element, that have either an
    [`id`{#the-form-element:the-id-attribute-5}](dom.html#the-id-attribute)
    attribute or a
    [`name`{#the-form-element:attr-fe-name-5}](form-control-infrastructure.html#attr-fe-name)
    attribute equal to `name`{.variable}, with the exception of
    [`input`{#the-form-element:the-input-element-3}](input.html#the-input-element)
    elements whose
    [`type`{#the-form-element:attr-input-type-3}](input.html#attr-input-type)
    attribute is in the [Image
    Button](input.html#image-button-state-(type=image)){#the-form-element:image-button-state-(type=image)-3}
    state, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-form-element:tree-order-2
    x-internal="tree-order"}.

2.  If `candidates`{.variable} is empty, let `candidates`{.variable} be
    a [live](infrastructure.html#live){#the-form-element:live-2}
    [`RadioNodeList`{#the-form-element:radionodelist-5}](common-dom-interfaces.html#radionodelist)
    object containing all the
    [`img`{#the-form-element:the-img-element-3}](embedded-content.html#the-img-element)
    elements, whose [form
    owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-5}
    is the
    [`form`{#the-form-element:the-form-element-15}](#the-form-element)
    element, that have either an
    [`id`{#the-form-element:the-id-attribute-6}](dom.html#the-id-attribute)
    attribute or a
    [`name`{#the-form-element:attr-img-name-3}](obsolete.html#attr-img-name)
    attribute equal to `name`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-form-element:tree-order-3
    x-internal="tree-order"}.

3.  If `candidates`{.variable} is empty, `name`{.variable} is the name
    of one of the entries in the
    [`form`{#the-form-element:the-form-element-16}](#the-form-element)
    element\'s [past names
    map](#past-names-map){#the-form-element:past-names-map-3}: return
    the object associated with `name`{.variable} in that map.

4.  If `candidates`{.variable} contains more than one node, return
    `candidates`{.variable}.

5.  Otherwise, `candidates`{.variable} contains exactly one node. Add a
    mapping from `name`{.variable} to the node in
    `candidates`{.variable} in the
    [`form`{#the-form-element:the-form-element-17}](#the-form-element)
    element\'s [past names
    map](#past-names-map){#the-form-element:past-names-map-4}, replacing
    the previous entry with the same name, if any.

6.  Return the node in `candidates`{.variable}.

If an element listed in a
[`form`{#the-form-element:the-form-element-18}](#the-form-element)
element\'s [past names
map](#past-names-map){#the-form-element:past-names-map-5} changes [form
owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-6},
then its entries must be removed from that map.

------------------------------------------------------------------------

The [`submit()`]{#dom-form-submit .dfn dfn-for="HTMLFormElement"
dfn-type="method"} method steps are to
[submit](form-control-infrastructure.html#concept-form-submit){#the-form-element:concept-form-submit}
[this](https://webidl.spec.whatwg.org/#this){#the-form-element:this
x-internal="this"} from
[this](https://webidl.spec.whatwg.org/#this){#the-form-element:this-2
x-internal="this"}, with *[submitted from `submit()`
method](form-control-infrastructure.html#submit-subbmitted-from-method)*
set to true.

The [`requestSubmit(``submitter`{.variable}`)`]{#dom-form-requestsubmit
.dfn dfn-for="HTMLFormElement" dfn-type="method"} method, when invoked,
must run the following steps:

1.  If `submitter`{.variable} is not null, then:

    1.  If `submitter`{.variable} is not a [submit
        button](#concept-submit-button){#the-form-element:concept-submit-button-2},
        then throw a
        [`TypeError`{#the-form-element:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

    2.  If `submitter`{.variable}\'s [form
        owner](form-control-infrastructure.html#form-owner){#the-form-element:form-owner-7}
        is not this
        [`form`{#the-form-element:the-form-element-19}](#the-form-element)
        element, then throw a
        [\"`NotFoundError`\"](https://webidl.spec.whatwg.org/#notfounderror){#the-form-element:notfounderror
        x-internal="notfounderror"}
        [`DOMException`{#the-form-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Otherwise, set `submitter`{.variable} to this
    [`form`{#the-form-element:the-form-element-20}](#the-form-element)
    element.

3.  [Submit](form-control-infrastructure.html#concept-form-submit){#the-form-element:concept-form-submit-2}
    this
    [`form`{#the-form-element:the-form-element-21}](#the-form-element)
    element, from `submitter`{.variable}.

The [`reset()`]{#dom-form-reset .dfn dfn-for="HTMLFormElement"
dfn-type="method"} method, when invoked, must run the following steps:

1.  If the
    [`form`{#the-form-element:the-form-element-22}](#the-form-element)
    element is marked as *[locked for reset](#locked-for-reset)*, then
    return.

2.  Mark the
    [`form`{#the-form-element:the-form-element-23}](#the-form-element)
    element as [locked for reset]{#locked-for-reset .dfn}.

3.  [Reset](form-control-infrastructure.html#concept-form-reset){#the-form-element:concept-form-reset}
    the
    [`form`{#the-form-element:the-form-element-24}](#the-form-element)
    element.

4.  Unmark the
    [`form`{#the-form-element:the-form-element-25}](#the-form-element)
    element as *[locked for reset](#locked-for-reset)*.

If the [`checkValidity()`]{#dom-form-checkvalidity .dfn
dfn-for="HTMLFormElement" dfn-type="method"} method is invoked, the user
agent must [statically validate the
constraints](form-control-infrastructure.html#statically-validate-the-constraints){#the-form-element:statically-validate-the-constraints}
of the
[`form`{#the-form-element:the-form-element-26}](#the-form-element)
element, and return true if the constraint validation returned a
*positive* result, and false if it returned a *negative* result.

If the [`reportValidity()`]{#dom-form-reportvalidity .dfn
dfn-for="HTMLFormElement" dfn-type="method"} method is invoked, the user
agent must [interactively validate the
constraints](form-control-infrastructure.html#interactively-validate-the-constraints){#the-form-element:interactively-validate-the-constraints-3}
of the
[`form`{#the-form-element:the-form-element-27}](#the-form-element)
element, and return true if the constraint validation returned a
*positive* result, and false if it returned a *negative* result.

::: example
This example shows two search forms:

``` html
<form action="https://www.google.com/search" method="get">
 <label>Google: <input type="search" name="q"></label> <input type="submit" value="Search...">
</form>
<form action="https://www.bing.com/search" method="get">
 <label>Bing: <input type="search" name="q"></label> <input type="submit" value="Search...">
</form>
```
:::

#### [4.10.4]{.secno} The [`label`]{.dfn dfn-type="element"} element[](#the-label-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/label](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label "The <label> HTML element represents a caption for an item in a user interface.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLabelElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement "The HTMLLabelElement interface gives access to properties specific to <label> elements. It inherits methods and properties from the base HTMLElement interface.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-label-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-label-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-label-element:phrasing-content-2}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-label-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-label-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-label-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-label-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-label-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-label-element:phrasing-content-2-3},
    but with no descendant [labelable
    elements](#category-label){#the-label-element:category-label} unless
    it is the element\'s [labeled
    control](#labeled-control){#the-label-element:labeled-control}, and
    no descendant
    [`label`{#the-label-element:the-label-element}](#the-label-element)
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-label-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-label-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-label-element:global-attributes}
:   [`for`{#the-label-element:attr-label-for}](#attr-label-for) ---
    Associate the label with form control

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-label-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-label).
:   [For implementers](https://w3c.github.io/html-aam/#el-label).

[DOM interface](dom.html#concept-element-dom){#the-label-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLLabelElement : HTMLElement {
      [HTMLConstructor] constructor();

      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute DOMString htmlFor;
      readonly attribute HTMLElement? control;
    };
    ```

The
[`label`{#the-label-element:the-label-element-2}](#the-label-element)
element [represents](dom.html#represents){#the-label-element:represents}
a caption in a user interface. The caption can be associated with a
specific form control, known as the
[`label`{#the-label-element:the-label-element-3}](#the-label-element)
element\'s [labeled control]{#labeled-control .dfn}, either using the
[`for`{#the-label-element:attr-label-for-2}](#attr-label-for) attribute,
or by putting the form control inside the
[`label`{#the-label-element:the-label-element-4}](#the-label-element)
element itself.

Except where otherwise specified by the following rules, a
[`label`{#the-label-element:the-label-element-5}](#the-label-element)
element has no [labeled
control](#labeled-control){#the-label-element:labeled-control-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/for](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/for "The for attribute is an allowed attribute for <label> and <output>. When used on a <label> element it indicates the form element that this label describes. When used on an <output> element it allows for an explicit relationship between the elements that represent values which are used in the output.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

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

The [`for`]{#attr-label-for .dfn dfn-for="label"
dfn-type="element-attr"} attribute may be specified to indicate a form
control with which the caption is to be associated. If the attribute is
specified, the attribute\'s value must be the
[ID](https://dom.spec.whatwg.org/#concept-id){#the-label-element:concept-id
x-internal="concept-id"} of a [labelable
element](#category-label){#the-label-element:category-label-2} in the
same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-label-element:tree
x-internal="tree"} as the
[`label`{#the-label-element:the-label-element-6}](#the-label-element)
element. If the attribute is specified and there is an element in the
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-label-element:tree-2
x-internal="tree"} whose
[ID](https://dom.spec.whatwg.org/#concept-id){#the-label-element:concept-id-2
x-internal="concept-id"} is equal to the value of the
[`for`{#the-label-element:attr-label-for-3}](#attr-label-for) attribute,
and the first such element in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-label-element:tree-order
x-internal="tree-order"} is a [labelable
element](#category-label){#the-label-element:category-label-3}, then
that element is the
[`label`{#the-label-element:the-label-element-7}](#the-label-element)
element\'s [labeled
control](#labeled-control){#the-label-element:labeled-control-3}.

If the [`for`{#the-label-element:attr-label-for-4}](#attr-label-for)
attribute is not specified, but the
[`label`{#the-label-element:the-label-element-8}](#the-label-element)
element has a [labelable
element](#category-label){#the-label-element:category-label-4}
descendant, then the first such descendant in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-label-element:tree-order-2
x-internal="tree-order"} is the
[`label`{#the-label-element:the-label-element-9}](#the-label-element)
element\'s [labeled
control](#labeled-control){#the-label-element:labeled-control-4}.

The
[`label`{#the-label-element:the-label-element-10}](#the-label-element)
element\'s exact default presentation and behavior, in particular what
its [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-label-element:activation-behaviour
x-internal="activation-behaviour"} might be, if anything, should match
the platform\'s label behavior. The [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-label-element:activation-behaviour-2
x-internal="activation-behaviour"} of a
[`label`{#the-label-element:the-label-element-11}](#the-label-element)
element for events targeted at [interactive
content](dom.html#interactive-content-2){#the-label-element:interactive-content-2-2}
descendants of a
[`label`{#the-label-element:the-label-element-12}](#the-label-element)
element, and any descendants of those [interactive
content](dom.html#interactive-content-2){#the-label-element:interactive-content-2-3}
descendants, must be to do nothing.

[Form-associated custom
elements](custom-elements.html#form-associated-custom-element){#the-label-element:form-associated-custom-element}
are [labelable
elements](#category-label){#the-label-element:category-label-5}, so for
user agents where the
[`label`{#the-label-element:the-label-element-13}](#the-label-element)
element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-label-element:activation-behaviour-3
x-internal="activation-behaviour"} impacts the [labeled
control](#labeled-control){#the-label-element:labeled-control-5}, both
built-in and custom elements will be impacted.

::: example
For example, on platforms where clicking a label activates the form
control, clicking the
[`label`{#the-label-element:the-label-element-14}](#the-label-element)
in the following snippet could trigger the user agent to [fire a `click`
event](webappapis.html#fire-a-click-event){#the-label-element:fire-a-click-event}
at the
[`input`{#the-label-element:the-input-element}](input.html#the-input-element)
element, as if the element itself had been triggered by the user:

``` html
<label><input type=checkbox name=lost> Lost</label>
```

Similarly, assuming `my-checkbox` was declared as a [form-associated
custom
element](custom-elements.html#form-associated-custom-element){#the-label-element:form-associated-custom-element-2}
(like in [this
example](custom-elements.html#custom-elements-face-example)), then the
code

``` html
<label><my-checkbox name=lost></my-checkbox> Lost</label>
```

would have the same behavior, [firing a `click`
event](webappapis.html#fire-a-click-event){#the-label-element:fire-a-click-event-2}
at the `my-checkbox` element.

On other platforms, the behavior in both cases might be just to focus
the control, or to do nothing.
:::

::: example
The following example shows three form controls each with a label, two
of which have small text showing the right format for users to use.

``` html
<p><label>Full name: <input name=fn> <small>Format: First Last</small></label></p>
<p><label>Age: <input name=age type=number min=0></label></p>
<p><label>Post code: <input name=pc> <small>Format: AB12 3CD</small></label></p>
```
:::

`label`{.variable}`.`[`control`](#dom-label-control){#dom-label-control-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLabelElement/control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control "The read-only HTMLLabelElement.control property returns a reference to the control (in the form of an object of type HTMLElement or one of its derivatives) with which the <label> element is associated, or null if the label isn't associated with a control.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the form control that is associated with this element.

`label`{.variable}`.`[`form`](#dom-label-form){#dom-label-form-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLabelElement/form](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/form "The read-only HTMLLabelElement.form property returns an HTMLFormElement object which represents the form of which the label's associated control is a part, or null if there is either no associated control, or if that control isn't in a form.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [form
owner](form-control-infrastructure.html#form-owner){#the-label-element:form-owner}
of the form control that is associated with this element.

Returns null if there isn\'t one.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLabelElement/htmlFor](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/htmlFor "The HTMLLabelElement.htmlFor property reflects the value of the for content property. That means that this script-accessible property is used to set and read the value of the content property for, which is the ID of the label's associated control element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`htmlFor`]{#dom-label-htmlfor .dfn dfn-for="HTMLLabelElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-label-element:reflect}
the [`for`{#the-label-element:attr-label-for-5}](#attr-label-for)
content attribute.

The [`control`]{#dom-label-control .dfn dfn-for="HTMLLabelElement"
dfn-type="attribute"} IDL attribute must return the
[`label`{#the-label-element:the-label-element-15}](#the-label-element)
element\'s [labeled
control](#labeled-control){#the-label-element:labeled-control-6}, if
any, or null if there isn\'t one.

The [`form`]{#dom-label-form .dfn dfn-for="HTMLLabelElement"
dfn-type="attribute"} IDL attribute must run the following steps:

1.  If the
    [`label`{#the-label-element:the-label-element-16}](#the-label-element)
    element has no [labeled
    control](#labeled-control){#the-label-element:labeled-control-7},
    then return null.

2.  If the
    [`label`{#the-label-element:the-label-element-17}](#the-label-element)
    element\'s [labeled
    control](#labeled-control){#the-label-element:labeled-control-8} is
    not a [form-associated
    element](#form-associated-element){#the-label-element:form-associated-element},
    then return null.

3.  Return the
    [`label`{#the-label-element:the-label-element-18}](#the-label-element)
    element\'s [labeled
    control](#labeled-control){#the-label-element:labeled-control-9}\'s
    [form
    owner](form-control-infrastructure.html#form-owner){#the-label-element:form-owner-2}
    (which can still be null).

The [`form`{#the-label-element:dom-label-form-2}](#dom-label-form) IDL
attribute on the
[`label`{#the-label-element:the-label-element-19}](#the-label-element)
element is different from the
[`form`{#the-label-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
IDL attribute on
[listed](#category-listed){#the-label-element:category-listed}
[form-associated
elements](#form-associated-element){#the-label-element:form-associated-element-2},
and the
[`label`{#the-label-element:the-label-element-20}](#the-label-element)
element does not have a
[`form`{#the-label-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
content attribute.

------------------------------------------------------------------------

`control`{.variable}`.`[`labels`](#dom-lfe-labels){#dom-lfe-labels-dev}

::::::::::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLButtonElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/labels "The HTMLButtonElement.labels read-only property returns a NodeList of the <label> elements associated with the <button> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLInputElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/labels "The HTMLInputElement.labels read-only property returns a NodeList of the <label> elements associated with the <input> element, if the element is not hidden. If the element has the type hidden, the property returns null.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLMeterElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMeterElement/labels "The HTMLMeterElement.labels read-only property returns a NodeList of the <label> elements associated with the <meter> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLOutputElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOutputElement/labels "The HTMLOutputElement.labels read-only property returns a NodeList of the <label> elements associated with the <output> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLProgressElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLProgressElement/labels "The HTMLProgressElement.labels read-only property returns a NodeList of the <label> elements associated with the <progress> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLSelectElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/labels "The HTMLSelectElement.labels read-only property returns a NodeList of the <label> elements associated with the <select> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLTextAreaElement/labels](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTextAreaElement/labels "The HTMLTextAreaElement.labels read-only property returns a NodeList of the <label> elements associated with the <textarea> element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::::::::::::

Returns a
[`NodeList`{#the-label-element:nodelist}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
of all the
[`label`{#the-label-element:the-label-element-21}](#the-label-element)
elements that the form control is associated with.

[Labelable
elements](#category-label){#the-label-element:category-label-6} and all
[`input`{#the-label-element:the-input-element-2}](input.html#the-input-element)
elements have a
[live](infrastructure.html#live){#the-label-element:live}
[`NodeList`{#the-label-element:nodelist-2}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
object associated with them that represents the list of
[`label`{#the-label-element:the-label-element-22}](#the-label-element)
elements, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-label-element:tree-order-3
x-internal="tree-order"}, whose [labeled
control](#labeled-control){#the-label-element:labeled-control-10} is the
element in question. The [`labels`]{#dom-lfe-labels .dfn
dfn-for="HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLProgressElement,HTMLMeterElement"
dfn-type="attribute"} IDL attribute of [labelable
elements](#category-label){#the-label-element:category-label-7} that are
not [form-associated custom
elements](custom-elements.html#form-associated-custom-element){#the-label-element:form-associated-custom-element-3},
and the [`labels`{#the-label-element:dom-lfe-labels}](#dom-lfe-labels)
IDL attribute of
[`input`{#the-label-element:the-input-element-3}](input.html#the-input-element)
elements, on getting, must return that
[`NodeList`{#the-label-element:nodelist-3}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
object, and that same value must always be returned, unless this element
is an
[`input`{#the-label-element:the-input-element-4}](input.html#the-input-element)
element whose
[`type`{#the-label-element:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Hidden](input.html#hidden-state-(type=hidden)){#the-label-element:hidden-state-(type=hidden)}
state, in which case it must instead return null.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/labels](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/labels "The labels read-only property of the ElementInternals interface returns the labels associated with the element.")

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

[Form-associated custom
elements](custom-elements.html#form-associated-custom-element){#the-label-element:form-associated-custom-element-4}
don\'t have a
[`labels`{#the-label-element:dom-lfe-labels-2}](#dom-lfe-labels) IDL
attribute. Instead, their
[`ElementInternals`{#the-label-element:elementinternals}](custom-elements.html#elementinternals)
object has a [`labels`]{#dom-elementinternals-labels .dfn
dfn-for="ElementInternals" dfn-type="attribute"} IDL attribute. On
getting, it must throw a
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-label-element:notsupportederror
x-internal="notsupportederror"}
[`DOMException`{#the-label-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the [target
element](custom-elements.html#internals-target){#the-label-element:internals-target}
is not a [form-associated custom
element](custom-elements.html#form-associated-custom-element){#the-label-element:form-associated-custom-element-5}.
Otherwise, it must return that
[`NodeList`{#the-label-element:nodelist-4}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
object, and that same value must always be returned.

::: example
This (non-conforming) example shows what happens to the
[`NodeList`{#the-label-element:nodelist-5}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
and what
[`labels`{#the-label-element:dom-lfe-labels-3}](#dom-lfe-labels) returns
when an
[`input`{#the-label-element:the-input-element-5}](input.html#the-input-element)
element has its
[`type`{#the-label-element:attr-input-type-2}](input.html#attr-input-type)
attribute changed.

``` html
<!doctype html>
<p><label><input></label></p>
<script>
 const input = document.querySelector('input');
 const labels = input.labels;
 console.assert(labels.length === 1);

 input.type = 'hidden';
 console.assert(labels.length === 0); // the input is no longer the label's labeled control
 console.assert(input.labels === null);

 input.type = 'checkbox';
 console.assert(labels.length === 1); // the input is once again the label's labeled control
 console.assert(input.labels === labels); // same value as returned originally
</script>
```
:::

[← 4.9 Tabular data](tables.html) --- [Table of Contents](./) ---
[4.10.5 The input element →](input.html)
