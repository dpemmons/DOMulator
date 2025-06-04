HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8 Embedded content](embedded-content.html) --- [Table of
Contents](./) --- [4.8.5 The iframe element →](iframe-embed-object.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.8.4]{.secno} Images](images.html#images)
            1.  [[4.8.4.1]{.secno}
                Introduction](images.html#introduction-3)
                1.  [[4.8.4.1.1]{.secno} Adaptive
                    images](images.html#adaptive-images)
            2.  [[4.8.4.2]{.secno} Attributes common to `source`, `img`,
                and `link`
                elements](images.html#attributes-common-to-source-and-img-elements)
                1.  [[4.8.4.2.1]{.secno} Srcset
                    attributes](images.html#srcset-attributes)
                2.  [[4.8.4.2.2]{.secno} Sizes
                    attributes](images.html#sizes-attributes)
            3.  [[4.8.4.3]{.secno} Processing
                model](images.html#images-processing-model)
                1.  [[4.8.4.3.1]{.secno} When to obtain
                    images](images.html#when-to-obtain-images)
                2.  [[4.8.4.3.2]{.secno} Reacting to DOM
                    mutations](images.html#reacting-to-dom-mutations)
                3.  [[4.8.4.3.3]{.secno} The list of available
                    images](images.html#the-list-of-available-images)
                4.  [[4.8.4.3.4]{.secno} Decoding
                    images](images.html#decoding-images)
                5.  [[4.8.4.3.5]{.secno} Updating the image
                    data](images.html#updating-the-image-data)
                6.  [[4.8.4.3.6]{.secno} Preparing an image for
                    presentation](images.html#preparing-an-image-for-presentation)
                7.  [[4.8.4.3.7]{.secno} Selecting an image
                    source](images.html#selecting-an-image-source)
                8.  [[4.8.4.3.8]{.secno} Creating a source set from
                    attributes](images.html#creating-a-source-set-from-attributes)
                9.  [[4.8.4.3.9]{.secno} Updating the source
                    set](images.html#updating-the-source-set)
                10. [[4.8.4.3.10]{.secno} Parsing a srcset
                    attribute](images.html#parsing-a-srcset-attribute)
                11. [[4.8.4.3.11]{.secno} Parsing a sizes
                    attribute](images.html#parsing-a-sizes-attribute)
                12. [[4.8.4.3.12]{.secno} Normalizing the source
                    densities](images.html#normalizing-the-source-densities)
                13. [[4.8.4.3.13]{.secno} Reacting to environment
                    changes](images.html#reacting-to-environment-changes)
            4.  [[4.8.4.4]{.secno} Requirements for providing text to
                act as an alternative for images](images.html#alt)
                1.  [[4.8.4.4.1]{.secno} General
                    guidelines](images.html#general-guidelines)
                2.  [[4.8.4.4.2]{.secno} A link or button containing
                    nothing but the
                    image](images.html#a-link-or-button-containing-nothing-but-the-image)
                3.  [[4.8.4.4.3]{.secno} A phrase or paragraph with an
                    alternative graphical representation: charts,
                    diagrams, graphs, maps,
                    illustrations](images.html#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations)
                4.  [[4.8.4.4.4]{.secno} A short phrase or label with an
                    alternative graphical representation: icons,
                    logos](images.html#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos)
                5.  [[4.8.4.4.5]{.secno} Text that has been rendered to
                    a graphic for typographical
                    effect](images.html#text-that-has-been-rendered-to-a-graphic-for-typographical-effect)
                6.  [[4.8.4.4.6]{.secno} A graphical representation of
                    some of the surrounding
                    text](images.html#a-graphical-representation-of-some-of-the-surrounding-text)
                7.  [[4.8.4.4.7]{.secno} Ancillary
                    images](images.html#ancillary-images)
                8.  [[4.8.4.4.8]{.secno} A purely decorative image that
                    doesn\'t add any
                    information](images.html#a-purely-decorative-image-that-doesn't-add-any-information)
                9.  [[4.8.4.4.9]{.secno} A group of images that form a
                    single larger picture with no
                    links](images.html#a-group-of-images-that-form-a-single-larger-picture-with-no-links)
                10. [[4.8.4.4.10]{.secno} A group of images that form a
                    single larger picture with
                    links](images.html#a-group-of-images-that-form-a-single-larger-picture-with-links)
                11. [[4.8.4.4.11]{.secno} A key part of the
                    content](images.html#a-key-part-of-the-content)
                12. [[4.8.4.4.12]{.secno} An image not intended for the
                    user](images.html#an-image-not-intended-for-the-user)
                13. [[4.8.4.4.13]{.secno} An image in an email or
                    private document intended for a specific person who
                    is known to be able to view
                    images](images.html#an-image-in-an-e-mail-or-private-document-intended-for-a-specific-person-who-is-known-to-be-able-to-view-images)
                14. [[4.8.4.4.14]{.secno} Guidance for markup
                    generators](images.html#guidance-for-markup-generators)
                15. [[4.8.4.4.15]{.secno} Guidance for conformance
                    checkers](images.html#guidance-for-conformance-checkers)
    :::

#### [4.8.4]{.secno} Images[](#images){.self-link}

##### [4.8.4.1]{.secno} Introduction[](#introduction-3){.self-link} {#introduction-3}

*This section is non-normative.*

To embed an image in HTML, when there is only a single image resource,
use the
[`img`{#introduction-3:the-img-element}](embedded-content.html#the-img-element)
element and its
[`src`{#introduction-3:attr-img-src}](embedded-content.html#attr-img-src)
attribute.

::: example
``` html
<h2>From today's featured article</h2>
<img src="/uploads/100-marie-lloyd.jpg" alt="" width="100" height="150">
<p><b><a href="/wiki/Marie_Lloyd">Marie Lloyd</a></b> (1870–1922)
was an English <a href="/wiki/Music_hall">music hall</a> singer, ...
```
:::

However, there are a number of situations for which the author might
wish to use multiple image resources that the user agent can choose
from:

- Different users might have different environmental characteristics:

  - The users\' physical screen size might be different from one
    another.

    ::: example
    A mobile phone\'s screen might be 4 inches diagonally, while a
    laptop\'s screen might be 14 inches diagonally.

    ![](data:image/svg+xml;base64,PHN2ZyBmb250LXNpemU9IjIuNSIgcm9sZT0iaW1nIiB2aWV3Ym94PSIwIDAgODAgMzIiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBoZWlnaHQ9IjIwMCIgYXJpYS1sYWJlbD0iVGhlIHBob25lJiMzOTtzIHNjcmVlbiBpcyBtdWNoIHNtYWxsZXIgY29tcGFyZWQgdG8gdGhlIGxhcHRvcCYjMzk7cyBzY3JlZW4uIj4KICAgICAgICAKICAgICAgICA8cmVjdCB3aWR0aD0iNiIgZmlsbD0id2hpdGUiIHN0cm9rZT0iYmxhY2siIHg9IjYiIGhlaWdodD0iMTIiIHk9IjE4IiByeD0iMSIgLz4KICAgICAgICA8dGV4dCB0cmFuc2Zvcm09InRyYW5zbGF0ZSg5IDI0KQogICAgICAgIHJvdGF0ZSgtNjApIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIj404oCzPC90ZXh0PgogICAgICAgIAogICAgICAgIDxyZWN0IHdpZHRoPSI0MCIgc3Ryb2tlLXdpZHRoPSIyIiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iMzAiIGhlaWdodD0iMjYiIHk9IjIiIHJ4PSIyIiAvPgogICAgICAgIDxsaW5lIHkyPSIzMCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2Utd2lkdGg9IjIiIHN0cm9rZT0iYmxhY2siIHgxPSIyNiIgeTE9IjMwIiB4Mj0iNzQiPjwvbGluZT4KICAgICAgICA8dGV4dCB0cmFuc2Zvcm09InRyYW5zbGF0ZSg1MCAxNSkKICAgICAgICByb3RhdGUoLTMwKSIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZG9taW5hbnQtYmFzZWxpbmU9Im1pZGRsZSI+MTTigLM8L3RleHQ+CiAgICAgICA8L3N2Zz4=)
    :::

    This is only relevant when an image\'s rendered size depends on the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport
    x-internal="viewport"} size.

  - The users\' screen pixel density might be different from one
    another.

    ::: example
    A mobile phone\'s screen might have three times as many physical
    pixels per inch compared to another mobile phone\'s screen,
    regardless of their physical screen size.

    ![](data:image/svg+xml;base64,PHN2ZyBmb250LXNpemU9IjIuNSIgcm9sZT0iaW1nIiB2aWV3Ym94PSIwIDAgNTYgMjciIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBoZWlnaHQ9IjE3MCIgYXJpYS1sYWJlbD0iT25lIHBob25lIGhhcyBiaWcgcGl4ZWxzLCB0aGUgb3RoZXIgaGFzIHNtYWxsIHBpeGVscy4iPgogICAgICAgIDxkZWZzPgogICAgICAgICA8cGF0dGVybiBpZD0iaW1nLWludHJvLXBpeGVsIiB3aWR0aD0iMyIgcGF0dGVybnVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiB4PSIwIiB5PSIwIiBoZWlnaHQ9IjMiPgogICAgICAgICAgPHJlY3QgZmlsbD0iYmxhY2siIHdpZHRoPSIzIiB4PSIwIiBoZWlnaHQ9IjMiIHk9IjAiIC8+CiAgICAgICAgICA8bGluZSB5Mj0iMi41IiBzdHJva2U9InJlZCIgeDE9IjAuNSIgeDI9IjAuNSIgeTE9IjAuNSI+PC9saW5lPgogICAgICAgICAgPGxpbmUgeTI9IjIuNSIgc3Ryb2tlPSJsaW1lIiB4MT0iMS41IiB4Mj0iMS41IiB5MT0iMC41Ij48L2xpbmU+CiAgICAgICAgICA8bGluZSB5Mj0iMi41IiBzdHJva2U9ImJsdWUiIHgxPSIyLjUiIHgyPSIyLjUiIHkxPSIwLjUiPjwvbGluZT4KICAgICAgICAgPC9wYXR0ZXJuPgogICAgICAgIDwvZGVmcz4KICAgICAgICAKICAgICAgICA8cmVjdCB3aWR0aD0iNiIgZmlsbD0id2hpdGUiIHN0cm9rZT0iYmxhY2siIHg9IjYiIGhlaWdodD0iMTIiIHk9IjgiIHJ4PSIxIiAvPgogICAgICAgIDxyZWN0IHdpZHRoPSI2IiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iMzYiIGhlaWdodD0iMTIiIHk9IjgiIHJ4PSIxIiAvPgogICAgICAgIAogICAgICAgIDxsaW5lIHkyPSIyNCIgc3Ryb2tlLXdpZHRoPSIzIiBzdHJva2U9ImJyb3duIiB4MT0iMTUiIHkxPSIxNiIgeDI9IjIyIj48L2xpbmU+CiAgICAgICAgPGNpcmNsZSBjeD0iMTAiIGN5PSIxMCIgZmlsbD0idXJsKCNpbWctaW50cm8tcGl4ZWwpIiBzdHJva2U9ImJsYWNrIiByPSI4Ij48L2NpcmNsZT4KICAgICAgICA8bGluZSB5Mj0iMjQiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlPSJicm93biIgeDE9IjQ1IiB5MT0iMTYiIHgyPSI1MiI+PC9saW5lPgogICAgICAgIDxjaXJjbGUgY3g9IjEyMCIgdHJhbnNmb3JtPSJzY2FsZSgwLjMzMzMzMykiIGN5PSIzMCIgcj0iMjQiIHN0cm9rZS13aWR0aD0iMyIgZmlsbD0idXJsKCNpbWctaW50cm8tcGl4ZWwpIiBzdHJva2U9ImJsYWNrIj48L2NpcmNsZT4KICAgICAgICAKICAgICAgICA8dGV4dCB4PSIyMCIgeT0iMTAiIGRvbWluYW50LWJhc2VsaW5lPSJtaWRkbGUiPjF4PC90ZXh0PgogICAgICAgIDx0ZXh0IHg9IjUwIiB5PSIxMCIgZG9taW5hbnQtYmFzZWxpbmU9Im1pZGRsZSI+M3g8L3RleHQ+CiAgICAgICA8L3N2Zz4=)
    :::

  - The users\' zoom level might be different from one another, or might
    change for a single user over time.

    A user might zoom in to a particular image to be able to get a more
    detailed look.

    The zoom level and the screen pixel density (the previous point) can
    both affect the number of physical screen pixels per [CSS
    pixel](https://drafts.csswg.org/css-values/#px){#introduction-3:'px'
    x-internal="'px'"}. This ratio is usually referred to as
    [device-pixel-ratio]{#device-pixel-ratio .dfn}.

  - The users\' screen orientation might be different from one another,
    or might change for a single user over time.

    ::: example
    A tablet can be held upright or rotated 90 degrees, so that the
    screen is either \"portrait\" or \"landscape\".

    ![](data:image/svg+xml;base64,PHN2ZyBmb250LXNpemU9IjIuNSIgcm9sZT0iaW1nIiB2aWV3Ym94PSIwIDAgNjAgMzIiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBoZWlnaHQ9IjIwMCIgYXJpYS1sYWJlbD0iVGhlIHRhYmxldCBoYXMgdHdvIG9yaWVudGF0aW9ucy4iPgogICAgICAgIAogICAgICAgIDxyZWN0IHdpZHRoPSIxNCIgZmlsbD0id2hpdGUiIHN0cm9rZT0iYmxhY2siIHg9IjYiIGhlaWdodD0iMjAiIHk9IjUiIHJ4PSIxIiAvPgogICAgICAgIDxsaW5lIHkyPSIyNCIgc3Ryb2tlLXdpZHRoPSIxLjEiIHN0cm9rZT0iYmxhY2siIHgxPSI2IiB5MT0iMjQiIHgyPSIyMCI+PC9saW5lPgogICAgICAgIDx0ZXh0IHRleHQtYW5jaG9yPSJtaWRkbGUiIHg9IjEzIiB5PSIxNC41IiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIj5Qb3J0cmFpdDwvdGV4dD4KICAgICAgICA8cmVjdCB3aWR0aD0iMjAiIGZpbGw9IndoaXRlIiBzdHJva2U9ImJsYWNrIiB4PSIzMCIgaGVpZ2h0PSIxNCIgeT0iMTEiIHJ4PSIxIiAvPgogICAgICAgIDxsaW5lIHkyPSIxMSIgc3Ryb2tlLXdpZHRoPSIxLjEiIHN0cm9rZT0iYmxhY2siIHgxPSIzMSIgeTE9IjI1IiB4Mj0iMzEiPjwvbGluZT4KICAgICAgICA8dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSI0MC41IiB5PSIxOCIgZG9taW5hbnQtYmFzZWxpbmU9Im1pZGRsZSI+TGFuZHNjYXBlPC90ZXh0PgogICAgICAgPC9zdmc+)
    :::

  - The users\' network speed, network latency and bandwidth cost might
    be different from one another, or might change for a single user
    over time.

    A user might be on a fast, low-latency and constant-cost connection
    while at work, on a slow, low-latency and constant-cost connection
    while at home, and on a variable-speed, high-latency and
    variable-cost connection anywhere else.

- Authors might want to show the same image content but with different
  rendered size depending on, usually, the width of the
  [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-2
  x-internal="viewport"}. This is usually referred to as [viewport-based
  selection]{#viewport-based-selection .dfn}.

  ::: example
  A web page might have a banner at the top that always spans the entire
  [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-3
  x-internal="viewport"} width. In this case, the rendered size of the
  image depends on the physical size of the screen (assuming a maximised
  browser window).

  ![](data:image/svg+xml;base64,PHN2ZyB2aWV3Ym94PSIwIDAgNTIgMjUuNiIgaGVpZ2h0PSIxNjAiIGFyaWEtbGFiZWw9IlRoZSB1cHJpZ2h0LWhlbGQgcGhvbmUgc2hvd3MgYQogICAgIHNtYWxsIHdvbGYgYXQgdGhlIHRvcCwgYW5kIHRoZSB0YWJsZXQgc2hvd3MgdGhlIHNhbWUgaW1hZ2UgYnV0IGl0IGlzIGJpZ2dlci4iIHJvbGU9ImltZyI+CiAgICAgIAoKICAgICAgCiAgICAgIDxyZWN0IHdpZHRoPSI2IiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iNiIgaGVpZ2h0PSIxMiIgeT0iOCIgcng9IjEiIC8+CiAgICAgIDxyZWN0IGZpbGw9IiM3MTY5NjYiIHg9IjciIHdpZHRoPSI0IiBoZWlnaHQ9IjIiIHk9IjkiIC8+CiAgICAgIDxpbWFnZSB4bGluazpocmVmPSIvaW1hZ2VzL3dvbGYuanBnIiB3aWR0aD0iNCIgeD0iNyIgeT0iOSIgaGVpZ2h0PSIyIj48L2ltYWdlPgoKICAgICAgCiAgICAgIDxyZWN0IHdpZHRoPSIyMCIgZmlsbD0id2hpdGUiIHN0cm9rZT0iYmxhY2siIHg9IjI1IiBoZWlnaHQ9IjE0IiB5PSI2IiByeD0iMSIgLz4KICAgICAgPGxpbmUgeTI9IjYiIHN0cm9rZS13aWR0aD0iMS4xIiBzdHJva2U9ImJsYWNrIiB4MT0iMjYiIHkxPSIyMCIgeDI9IjI2Ij48L2xpbmU+CiAgICAgIDxyZWN0IGZpbGw9IiM3MTY5NjYiIHg9IjI3IiB3aWR0aD0iMTciIGhlaWdodD0iOC41IiB5PSI3IiAvPgogICAgICA8aW1hZ2UgeGxpbms6aHJlZj0iL2ltYWdlcy93b2xmLmpwZyIgd2lkdGg9IjE3IiB4PSIyNyIgeT0iNyIgaGVpZ2h0PSI4LjUiPjwvaW1hZ2U+CgogICAgIDwvc3ZnPg==)
  :::

  ::: example
  Another web page might have images in columns, with a single column
  for screens with a small physical size, two columns for screens with
  medium physical size, and three columns for screens with big physical
  size, with the images varying in rendered size in each case to fill up
  the
  [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-4
  x-internal="viewport"}. In this case, the rendered size of an image
  might be *bigger* in the one-column layout compared to the two-column
  layout, despite the screen being smaller.

  ![](data:image/svg+xml;base64,PHN2ZyBmb250LXNpemU9IjIuNSIgcm9sZT0iaW1nIiB2aWV3Ym94PSIwIDAgMTE1IDMyIiBmb250LWZhbWlseT0ic2Fucy1zZXJpZiIgaGVpZ2h0PSIyMDAiIGFyaWEtbGFiZWw9IlRoZSByb3RhdGVkIHBob25lIHNob3dzIGEgdG9wIHBhcnQgb2YgYW4gaW1hZ2Ugb2YgYSBrZXR0bGViZWxsIHN3aW5nOyB0aGUKICAgICB1cHJpZ2h0LWhlbGQgdGFibGV0IHNob3dzIGEgYml0IHNtYWxsZXIgaW1hZ2VzIGluIHR3byBjb2x1bW5zOyB0aGUgbGFwdG9wIHNob3dzIGltYWdlcyBpbiB0aHJlZQogICAgIGNvbHVtbnMuIj4KICAgICAgCgogICAgICAKICAgICAgPHJlY3Qgd2lkdGg9IjEyIiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iNiIgaGVpZ2h0PSI2IiB5PSIyNCIgcng9IjEiIC8+CiAgICAgIDxyZWN0IGZpbGw9IiM5ODdiNWEiIHdpZHRoPSIxMCIgeD0iNyIgaGVpZ2h0PSI0LjUiIHk9IjI1IiAvPgogICAgICA8aW1hZ2UgeGxpbms6aHJlZj0iL2ltYWdlcy9rZXR0bGViZWxsLmpwZyIgcHJlc2VydmVhc3BlY3RyYXRpbz0ieE1pbllNaW4gc2xpY2UiIHdpZHRoPSIxMCIgeD0iNyIgeT0iMjUiIGhlaWdodD0iNC41Ij48L2ltYWdlPgogICAgICA8dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSIxMiIgeT0iMjEiPk5hcnJvdywgMSBjb2x1bW48L3RleHQ+CgogICAgICAKICAgICAgPHJlY3Qgd2lkdGg9IjE0IiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iMzIiIGhlaWdodD0iMjAiIHk9IjEwIiByeD0iMSIgLz4KICAgICAgPGxpbmUgeTI9IjI5IiBzdHJva2Utd2lkdGg9IjEuMSIgc3Ryb2tlPSJibGFjayIgeDE9IjMyIiB5MT0iMjkiIHgyPSI0NiI+PC9saW5lPgogICAgICA8cmVjdCBmaWxsPSIjOTg3YjVhIiB3aWR0aD0iNS41IiB4PSIzMyIgaGVpZ2h0PSI1LjUiIHk9IjExIiAvPgogICAgICA8aW1hZ2UgeGxpbms6aHJlZj0iL2ltYWdlcy9rZXR0bGViZWxsLmpwZyIgd2lkdGg9IjUuNSIgeD0iMzMiIHk9IjExIiBoZWlnaHQ9IjUuNSI+PC9pbWFnZT4KICAgICAgPHJlY3QgZmlsbD0iYnVybHl3b29kIiB3aWR0aD0iNS41IiB4PSIzOS41IiBoZWlnaHQ9IjUuNSIgeT0iMTEiIC8+CiAgICAgIDxyZWN0IGZpbGw9InNpbHZlciIgd2lkdGg9IjUuNSIgeD0iMzMiIGhlaWdodD0iNS41IiB5PSIxNy41IiAvPgogICAgICA8dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiB4PSIzOSIgeT0iNyI+TWVkaXVtLCAyIGNvbHVtbnM8L3RleHQ+CgogICAgICAKICAgICAgPHJlY3Qgd2lkdGg9IjQwIiBzdHJva2Utd2lkdGg9IjIiIGZpbGw9IndoaXRlIiBzdHJva2U9ImJsYWNrIiB4PSI2NSIgaGVpZ2h0PSIyNiIgeT0iMiIgcng9IjIiIC8+CiAgICAgIDxsaW5lIHkyPSIzMCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2Utd2lkdGg9IjIiIHN0cm9rZT0iYmxhY2siIHgxPSI2MSIgeTE9IjMwIiB4Mj0iMTA5Ij48L2xpbmU+CiAgICAgIDxyZWN0IGZpbGw9IiM5ODdiNWEiIHdpZHRoPSIxMCIgeD0iNjciIGhlaWdodD0iMTAiIHk9IjQiIC8+CiAgICAgIDxpbWFnZSB4bGluazpocmVmPSIvaW1hZ2VzL2tldHRsZWJlbGwuanBnIiB3aWR0aD0iMTAiIHg9IjY3IiB5PSI0IiBoZWlnaHQ9IjEwIj48L2ltYWdlPgogICAgICA8cmVjdCBmaWxsPSJidXJseXdvb2QiIHdpZHRoPSIxMCIgeD0iODAiIGhlaWdodD0iMTAiIHk9IjQiIC8+CiAgICAgIDxyZWN0IGZpbGw9InNpbHZlciIgd2lkdGg9IjEwIiB4PSI5MyIgaGVpZ2h0PSIxMCIgeT0iNCIgLz4KICAgICAgPHRleHQgdGV4dC1hbmNob3I9Im1pZGRsZSIgeD0iODUiIHk9IjI1Ij5XaWRlLCAzIGNvbHVtbnM8L3RleHQ+CiAgICAgPC9zdmc+)
  :::

- Authors might want to show different image content depending on the
  rendered size of the image. This is usually referred to as [art
  direction]{#art-direction .dfn}.

  ::: example
  When a web page is viewed on a screen with a large physical size
  (assuming a maximised browser window), the author might wish to
  include some less relevant parts surrounding the critical part of the
  image. When the same web page is viewed on a screen with a small
  physical size, the author might wish to show only the critical part of
  the image.

  ![](data:image/svg+xml;base64,PHN2ZyB2aWV3Ym94PSIwIDAgNTIgMjUuNiIgaGVpZ2h0PSIxNjAiIGFyaWEtbGFiZWw9IlRoZSB1cHJpZ2h0LWhlbGQgcGhvbmUgc2hvd3MgYQogICAgIGNyb3BwZWQgaW1hZ2Ugb2YgYSB3b2xmOyB0aGUgcm90YXRlZCB0YWJsZXQgc2hvd3MgdGhlIHVuY3JvcHBlZCBpbWFnZS4iIHJvbGU9ImltZyI+CiAgICAgIAoKICAgICAgCiAgICAgIDxyZWN0IHdpZHRoPSI2IiBmaWxsPSJ3aGl0ZSIgc3Ryb2tlPSJibGFjayIgeD0iNiIgaGVpZ2h0PSIxMiIgeT0iOCIgcng9IjEiIC8+CiAgICAgIDxyZWN0IGZpbGw9IiM3MTY5NjYiIHg9IjciIHdpZHRoPSI0IiBoZWlnaHQ9IjYiIHk9IjkiIC8+CiAgICAgIDxpbWFnZSB4bGluazpocmVmPSIvaW1hZ2VzL3dvbGYuanBnIiBwcmVzZXJ2ZWFzcGVjdHJhdGlvPSJ4TWlkWU1pZCBzbGljZSIgd2lkdGg9IjQiIHg9IjciIHk9IjkiIGhlaWdodD0iNiI+PC9pbWFnZT4KCiAgICAgIAogICAgICA8cmVjdCB3aWR0aD0iMjAiIGZpbGw9IndoaXRlIiBzdHJva2U9ImJsYWNrIiB4PSIyNSIgaGVpZ2h0PSIxNCIgeT0iNiIgcng9IjEiIC8+CiAgICAgIDxsaW5lIHkyPSI2IiBzdHJva2Utd2lkdGg9IjEuMSIgc3Ryb2tlPSJibGFjayIgeDE9IjI2IiB5MT0iMjAiIHgyPSIyNiI+PC9saW5lPgogICAgICA8cmVjdCBmaWxsPSIjNzE2OTY2IiB4PSIyNyIgd2lkdGg9IjE3IiBoZWlnaHQ9IjguNSIgeT0iNyIgLz4KICAgICAgPGltYWdlIHhsaW5rOmhyZWY9Ii9pbWFnZXMvd29sZi5qcGciIHdpZHRoPSIxNyIgeD0iMjciIHk9IjciIGhlaWdodD0iOC41Ij48L2ltYWdlPgoKICAgICA8L3N2Zz4=)
  :::

- Authors might want to show the same image content but using different
  image formats, depending on which image formats the user agent
  supports. This is usually referred to as [image format-based
  selection]{#image-format-based-selection .dfn}.

  A web page might have some images in the JPEG, WebP and JPEG XR image
  formats, with the latter two having better compression abilities
  compared to JPEG. Since different user agents can support different
  image formats, with some formats offering better compression ratios,
  the author would like to serve the better formats to user agents that
  support them, while providing JPEG fallback for user agents that
  don\'t.

The above situations are not mutually exclusive. For example, it is
reasonable to combine different resources for different
[device-pixel-ratio](#device-pixel-ratio){#introduction-3:device-pixel-ratio}
with different resources for [art
direction](#art-direction){#introduction-3:art-direction}.

While it is possible to solve these problems using scripting, doing so
introduces some other problems:

- Some user agents aggressively download images specified in the HTML
  markup, before scripts have had a chance to run, so that web pages
  complete loading sooner. If a script changes which image to download,
  the user agent will potentially start two separate downloads, which
  can instead cause worse page loading performance.

- If the author avoids specifying any image in the HTML markup and
  instead instantiates a single download from script, that avoids the
  double download problem above but then no image will be downloaded at
  all for users with scripting disabled and the aggressive image
  downloading optimization will also be disabled.

With this in mind, this specification introduces a number of features to
address the above problems in a declarative manner.

[Device-pixel-ratio](#device-pixel-ratio){#introduction-3:device-pixel-ratio-2}-based selection when the rendered size of the image is fixed

:   The
    [`src`{#introduction-3:attr-img-src-2}](embedded-content.html#attr-img-src)
    and
    [`srcset`{#introduction-3:attr-img-srcset}](embedded-content.html#attr-img-srcset)
    attributes on the
    [`img`{#introduction-3:the-img-element-2}](embedded-content.html#the-img-element)
    element can be used, using the `x` descriptor, to provide multiple
    images that only vary in their size (the smaller image is a
    scaled-down version of the bigger image).

    The `x` descriptor is not appropriate when the rendered size of the
    image depends on the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-5
    x-internal="viewport"} width ([viewport-based
    selection](#viewport-based-selection){#introduction-3:viewport-based-selection}),
    but can be used together with [art
    direction](#art-direction){#introduction-3:art-direction-2}.

    ::: example
    ``` html
    <h2>From today's featured article</h2>
    <img src="/uploads/100-marie-lloyd.jpg"
         srcset="/uploads/150-marie-lloyd.jpg 1.5x, /uploads/200-marie-lloyd.jpg 2x"
         alt="" width="100" height="150">
    <p><b><a href="/wiki/Marie_Lloyd">Marie Lloyd</a></b> (1870–1922)
    was an English <a href="/wiki/Music_hall">music hall</a> singer, ...
    ```

    The user agent can choose any of the given resources depending on
    the user\'s screen\'s pixel density, zoom level, and possibly other
    factors such as the user\'s network conditions.

    For backwards compatibility with older user agents that don\'t yet
    understand the
    [`srcset`{#introduction-3:attr-img-srcset-2}](embedded-content.html#attr-img-srcset)
    attribute, one of the URLs is specified in the
    [`img`{#introduction-3:the-img-element-3}](embedded-content.html#the-img-element)
    element\'s
    [`src`{#introduction-3:attr-img-src-3}](embedded-content.html#attr-img-src)
    attribute. This will result in something useful (though perhaps
    lower-resolution than the user would like) being displayed even in
    older user agents. For new user agents, the
    [`src`{#introduction-3:attr-img-src-4}](embedded-content.html#attr-img-src)
    attribute participates in the resource selection, as if it was
    specified in
    [`srcset`{#introduction-3:attr-img-srcset-3}](embedded-content.html#attr-img-srcset)
    with a `1x` descriptor.

    The image\'s rendered size is given in the
    [`width`{#introduction-3:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    and
    [`height`{#introduction-3:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    attributes, which allows the user agent to allocate space for the
    image before it is downloaded.
    :::

[Viewport-based selection](#viewport-based-selection){#introduction-3:viewport-based-selection-2}

:   The
    [`srcset`{#introduction-3:attr-img-srcset-4}](embedded-content.html#attr-img-srcset)
    and
    [`sizes`{#introduction-3:attr-img-sizes}](embedded-content.html#attr-img-sizes)
    attributes can be used, using the `w` descriptor, to provide
    multiple images that only vary in their size (the smaller image is a
    scaled-down version of the bigger image).

    ::: example
    In this example, a banner image takes up the entire
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-6
    x-internal="viewport"} width (using appropriate CSS).

    ``` html
    <h1><img sizes="100vw" srcset="wolf-400.jpg 400w, wolf-800.jpg 800w, wolf-1600.jpg 1600w"
         src="wolf-400.jpg" alt="The rad wolf"></h1>
    ```

    The user agent will calculate the effective pixel density of each
    image from the specified `w` descriptors and the specified rendered
    size in the
    [`sizes`{#introduction-3:attr-img-sizes-2}](embedded-content.html#attr-img-sizes)
    attribute. It can then choose any of the given resources depending
    on the user\'s screen\'s pixel density, zoom level, and possibly
    other factors such as the user\'s network conditions.

    If the user\'s screen is 320 [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#introduction-3:'px'-2
    x-internal="'px'"} wide, this is equivalent to specifying
    `wolf-400.jpg 1.25x, wolf-800.jpg 2.5x, wolf-1600.jpg 5x`. On the
    other hand, if the user\'s screen is 1200 [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#introduction-3:'px'-3
    x-internal="'px'"} wide, this is equivalent to specifying
    `wolf-400.jpg 0.33x, wolf-800.jpg 0.67x, wolf-1600.jpg 1.33x`. By
    using the `w` descriptors and the
    [`sizes`{#introduction-3:attr-img-sizes-3}](embedded-content.html#attr-img-sizes)
    attribute, the user agent can choose the correct image source to
    download regardless of how large the user\'s device is.

    For backwards compatibility, one of the URLs is specified in the
    [`img`{#introduction-3:the-img-element-4}](embedded-content.html#the-img-element)
    element\'s
    [`src`{#introduction-3:attr-img-src-5}](embedded-content.html#attr-img-src)
    attribute. In new user agents, the
    [`src`{#introduction-3:attr-img-src-6}](embedded-content.html#attr-img-src)
    attribute is ignored when the
    [`srcset`{#introduction-3:attr-img-srcset-5}](embedded-content.html#attr-img-srcset)
    attribute uses `w` descriptors.
    :::

    ::: example
    In this example, the web page has three layouts depending on the
    width of the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-7
    x-internal="viewport"}. The narrow layout has one column of images
    (the width of each image is about 100%), the middle layout has two
    columns of images (the width of each image is about 50%), and the
    widest layout has three columns of images, and some page margin (the
    width of each image is about 33%). It breaks between these layouts
    when the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-8
    x-internal="viewport"} is `30em` wide and `50em` wide, respectively.

    ``` html
    <img sizes="(max-width: 30em) 100vw, (max-width: 50em) 50vw, calc(33vw - 100px)"
         srcset="swing-200.jpg 200w, swing-400.jpg 400w, swing-800.jpg 800w, swing-1600.jpg 1600w"
         src="swing-400.jpg" alt="Kettlebell Swing">
    ```

    The
    [`sizes`{#introduction-3:attr-img-sizes-4}](embedded-content.html#attr-img-sizes)
    attribute sets up the layout breakpoints at `30em` and `50em`, and
    declares the image sizes between these breakpoints to be `100vw`,
    `50vw`, or `calc(33vw - 100px)`. These sizes do not necessarily have
    to match up exactly with the actual image width as specified in the
    CSS.

    The user agent will pick a width from the
    [`sizes`{#introduction-3:attr-img-sizes-5}](embedded-content.html#attr-img-sizes)
    attribute, using the first item with a
    [\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition){#introduction-3:media-condition
    x-internal="media-condition"} (the part in parentheses) that
    evaluates to true, or using the last item (`calc(33vw - 100px)`) if
    they all evaluate to false.

    For example, if the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-9
    x-internal="viewport"} width is `29em`, then `(max-width: 30em)`
    evaluates to true and `100vw` is used, so the image size, for the
    purpose of resource selection, is `29em`. If the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-10
    x-internal="viewport"} width is instead `32em`, then
    `(max-width: 30em)` evaluates to false, but `(max-width: 50em)`
    evaluates to true and `50vw` is used, so the image size, for the
    purpose of resource selection, is `16em` (half the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-11
    x-internal="viewport"} width). Notice that the slightly wider
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-12
    x-internal="viewport"} results in a smaller image because of the
    different layout.

    The user agent can then calculate the effective pixel density and
    choose an appropriate resource similarly to the previous example.
    :::

    ::: example
    This example is the same as the previous example, but the image is
    [lazy-loaded](urls-and-fetching.html#attr-loading-lazy-state){#introduction-3:attr-loading-lazy-state}.
    In this case, the
    [`sizes`{#introduction-3:attr-img-sizes-6}](embedded-content.html#attr-img-sizes)
    attribute can use the
    [`auto`{#introduction-3:valdef-sizes-auto}](#valdef-sizes-auto)
    keyword, and the user agent will use the
    [`width`{#introduction-3:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
    attribute (or the width specified in CSS) for the [source
    size](#source-size-2){#introduction-3:source-size-2}.

    ``` html
    <img loading="lazy" width="200" height="200" sizes="auto"
         srcset="swing-200.jpg 200w, swing-400.jpg 400w, swing-800.jpg 800w, swing-1600.jpg 1600w"
         src="swing-400.jpg" alt="Kettlebell Swing">
    ```

    For better backwards-compatibility with legacy user agents that
    don\'t support the
    [`auto`{#introduction-3:valdef-sizes-auto-2}](#valdef-sizes-auto)
    keyword, fallback sizes can be specified if desired.

    ``` html
    <img loading="lazy" width="200" height="200"
         sizes="auto, (max-width: 30em) 100vw, (max-width: 50em) 50vw, calc(33vw - 100px)"
         srcset="swing-200.jpg 200w, swing-400.jpg 400w, swing-800.jpg 800w, swing-1600.jpg 1600w"
         src="swing-400.jpg" alt="Kettlebell Swing">
    ```
    :::

[Art direction](#art-direction){#introduction-3:art-direction-3}-based selection

:   The
    [`picture`{#introduction-3:the-picture-element}](embedded-content.html#the-picture-element)
    element and the
    [`source`{#introduction-3:the-source-element}](embedded-content.html#the-source-element)
    element, together with the
    [`media`{#introduction-3:attr-source-media}](embedded-content.html#attr-source-media)
    attribute, can be used to provide multiple images that vary the
    image content (for instance the smaller image might be a cropped
    version of the bigger image).

    ::: example
    ``` html
    <picture>
      <source media="(min-width: 45em)" srcset="large.jpg">
      <source media="(min-width: 32em)" srcset="med.jpg">
      <img src="small.jpg" alt="The wolf runs through the snow.">
    </picture>
    ```

    The user agent will choose the first
    [`source`{#introduction-3:the-source-element-2}](embedded-content.html#the-source-element)
    element for which the media query in the
    [`media`{#introduction-3:attr-source-media-2}](embedded-content.html#attr-source-media)
    attribute matches, and then choose an appropriate URL from its
    [`srcset`{#introduction-3:attr-source-srcset}](embedded-content.html#attr-source-srcset)
    attribute.

    The rendered size of the image varies depending on which resource is
    chosen. To specify dimensions that the user agent can use before
    having downloaded the image, CSS can be used.

    ``` css
    img { width: 300px; height: 300px }
    @media (min-width: 32em) { img { width: 500px; height:300px } }
    @media (min-width: 45em) { img { width: 700px; height:400px } }
    ```
    :::

    ::: example
    This example combines [art
    direction](#art-direction){#introduction-3:art-direction-4}- and
    [device-pixel-ratio](#device-pixel-ratio){#introduction-3:device-pixel-ratio-3}-based
    selection. A banner that takes half the
    [viewport](https://drafts.csswg.org/css2/#viewport){#introduction-3:viewport-13
    x-internal="viewport"} is provided in two versions, one for wide
    screens and one for narrow screens.

    ``` html
    <h1>
     <picture>
      <source media="(max-width: 500px)" srcset="banner-phone.jpeg, banner-phone-HD.jpeg 2x">
      <img src="banner.jpeg" srcset="banner-HD.jpeg 2x" alt="The Breakfast Combo">
     </picture>
    </h1>
    ```
    :::

[Image format-based selection](#image-format-based-selection){#introduction-3:image-format-based-selection}

:   The
    [`type`{#introduction-3:attr-source-type}](embedded-content.html#attr-source-type)
    attribute on the
    [`source`{#introduction-3:the-source-element-3}](embedded-content.html#the-source-element)
    element can be used to provide multiple images in different formats.

    ::: example
    ``` html
    <h2>From today's featured article</h2>
    <picture>
     <source srcset="/uploads/100-marie-lloyd.webp" type="image/webp">
     <source srcset="/uploads/100-marie-lloyd.jxr" type="image/vnd.ms-photo">
     <img src="/uploads/100-marie-lloyd.jpg" alt="" width="100" height="150">
    </picture>
    <p><b><a href="/wiki/Marie_Lloyd">Marie Lloyd</a></b> (1870–1922)
    was an English <a href="/wiki/Music_hall">music hall</a> singer, ...
    ```

    In this example, the user agent will choose the first source that
    has a
    [`type`{#introduction-3:attr-source-type-2}](embedded-content.html#attr-source-type)
    attribute with a supported MIME type. If the user agent supports
    WebP images, the first
    [`source`{#introduction-3:the-source-element-4}](embedded-content.html#the-source-element)
    element will be chosen. If not, but the user agent does support JPEG
    XR images, the second
    [`source`{#introduction-3:the-source-element-5}](embedded-content.html#the-source-element)
    element will be chosen. If neither of those formats are supported,
    the
    [`img`{#introduction-3:the-img-element-5}](embedded-content.html#the-img-element)
    element will be chosen.
    :::

###### [4.8.4.1.1]{.secno} Adaptive images[](#adaptive-images){.self-link}

*This section is non-normative.*

CSS and media queries can be used to construct graphical page layouts
that adapt dynamically to the user\'s environment, in particular to
different
[viewport](https://drafts.csswg.org/css2/#viewport){#adaptive-images:viewport
x-internal="viewport"} dimensions and pixel densities. For content,
however, CSS does not help; instead, we have the
[`img`{#adaptive-images:the-img-element}](embedded-content.html#the-img-element)
element\'s
[`srcset`{#adaptive-images:attr-img-srcset}](embedded-content.html#attr-img-srcset)
attribute and the
[`picture`{#adaptive-images:the-picture-element}](embedded-content.html#the-picture-element)
element. This section walks through a sample case showing how to use
these features.

Consider a situation where on wide screens (wider than 600 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#adaptive-images:'px'
x-internal="'px'"}) a 300×150 image named `a-rectangle.png` is to be
used, but on smaller screens (600 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#adaptive-images:'px'-2
x-internal="'px'"} and less), a smaller 100×100 image called
`a-square.png` is to be used. The markup for this would look like this:

``` html
<figure>
 <picture>
  <source srcset="a-square.png" media="(max-width: 600px)">
  <img src="a-rectangle.png" alt="Barney Frank wears a suit and glasses.">
 </picture>
 <figcaption>Barney Frank, 2011</figcaption>
</figure>
```

For details on what to put in the
[`alt`{#adaptive-images:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute, see the [Requirements for providing text to act as an
alternative for images](#alt) section.

The problem with this is that the user agent does not necessarily know
what dimensions to use for the image when the image is loading. To avoid
the layout having to be reflowed multiple times as the page is loading,
CSS and CSS media queries can be used to provide the dimensions:

``` html
<style>
 #a { width: 300px; height: 150px; }
 @media (max-width: 600px) { #a { width: 100px; height: 100px; } }
</style>
<figure>
 <picture>
  <source srcset="a-square.png" media="(max-width: 600px)">
  <img src="a-rectangle.png" alt="Barney Frank wears a suit and glasses." id="a">
 </picture>
 <figcaption>Barney Frank, 2011</figcaption>
</figure>
```

Alternatively, the
[`width`{#adaptive-images:attr-dim-width}](embedded-content-other.html#attr-dim-width)
and
[`height`{#adaptive-images:attr-dim-height}](embedded-content-other.html#attr-dim-height)
attributes can be used on the
[`source`{#adaptive-images:the-source-element}](embedded-content.html#the-source-element)
and
[`img`{#adaptive-images:the-img-element-2}](embedded-content.html#the-img-element)
elements to provide the width and height:

``` html
<figure>
 <picture>
  <source srcset="a-square.png" media="(max-width: 600px)" width="100" height="100">
  <img src="a-rectangle.png" width="300" height="150"
  alt="Barney Frank wears a suit and glasses.">
 </picture>
 <figcaption>Barney Frank, 2011</figcaption>
</figure>
```

------------------------------------------------------------------------

The
[`img`{#adaptive-images:the-img-element-3}](embedded-content.html#the-img-element)
element is used with the
[`src`{#adaptive-images:attr-img-src}](embedded-content.html#attr-img-src)
attribute, which gives the URL of the image to use for legacy user
agents that do not support the
[`picture`{#adaptive-images:the-picture-element-2}](embedded-content.html#the-picture-element)
element. This leads to a question of which image to provide in the
[`src`{#adaptive-images:attr-img-src-2}](embedded-content.html#attr-img-src)
attribute.

If the author wants the biggest image in legacy user agents, the markup
could be as follows:

``` html
<picture>
 <source srcset="pear-mobile.jpeg" media="(max-width: 720px)">
 <source srcset="pear-tablet.jpeg" media="(max-width: 1280px)">
 <img src="pear-desktop.jpeg" alt="The pear is juicy.">
</picture>
```

However, if legacy mobile user agents are more important, one can list
all three images in the
[`source`{#adaptive-images:the-source-element-2}](embedded-content.html#the-source-element)
elements, overriding the
[`src`{#adaptive-images:attr-img-src-3}](embedded-content.html#attr-img-src)
attribute entirely.

``` html
<picture>
 <source srcset="pear-mobile.jpeg" media="(max-width: 720px)">
 <source srcset="pear-tablet.jpeg" media="(max-width: 1280px)">
 <source srcset="pear-desktop.jpeg">
 <img src="pear-mobile.jpeg" alt="The pear is juicy.">
</picture>
```

Since at this point the
[`src`{#adaptive-images:attr-img-src-4}](embedded-content.html#attr-img-src)
attribute is actually being ignored entirely by
[`picture`{#adaptive-images:the-picture-element-3}](embedded-content.html#the-picture-element)-supporting
user agents, the
[`src`{#adaptive-images:attr-img-src-5}](embedded-content.html#attr-img-src)
attribute can default to any image, including one that is neither the
smallest nor biggest:

``` html
<picture>
 <source srcset="pear-mobile.jpeg" media="(max-width: 720px)">
 <source srcset="pear-tablet.jpeg" media="(max-width: 1280px)">
 <source srcset="pear-desktop.jpeg">
 <img src="pear-tablet.jpeg" alt="The pear is juicy.">
</picture>
```

------------------------------------------------------------------------

Above the `max-width` media feature is used, giving the maximum
([viewport](https://drafts.csswg.org/css2/#viewport){#adaptive-images:viewport-2
x-internal="viewport"}) dimensions that an image is intended for. It is
also possible to use `min-width` instead.

``` html
<picture>
 <source srcset="pear-desktop.jpeg" media="(min-width: 1281px)">
 <source srcset="pear-tablet.jpeg" media="(min-width: 721px)">
 <img src="pear-mobile.jpeg" alt="The pear is juicy.">
</picture>
```

##### [4.8.4.2]{.secno} Attributes common to [`source`{#attributes-common-to-source-and-img-elements:the-source-element}](embedded-content.html#the-source-element), [`img`{#attributes-common-to-source-and-img-elements:the-img-element}](embedded-content.html#the-img-element), and [`link`{#attributes-common-to-source-and-img-elements:the-link-element}](semantics.html#the-link-element) elements[](#attributes-common-to-source-and-img-elements){.self-link} {#attributes-common-to-source-and-img-elements}

###### [4.8.4.2.1]{.secno} Srcset attributes[](#srcset-attributes){.self-link}

A [srcset attribute]{#srcset-attribute .dfn} is an attribute with
requirements defined in this section.

If present, its value must consist of one or more [image candidate
strings](#image-candidate-string){#srcset-attributes:image-candidate-string},
each separated from the next by a U+002C COMMA character (,). If an
[image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-2}
contains no descriptors and no [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters
x-internal="space-characters"} after the URL, the following [image
candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-3},
if there is one, must begin with one or more [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-2
x-internal="space-characters"}.

An [image candidate string]{#image-candidate-string .dfn} consists of
the following components, in order, with the further restrictions
described below this list:

1.  Zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-3
    x-internal="space-characters"}.

2.  A [valid non-empty
    URL](urls-and-fetching.html#valid-non-empty-url){#srcset-attributes:valid-non-empty-url}
    that does not start or end with a U+002C COMMA character (,),
    referencing a non-interactive, optionally animated, image resource
    that is neither paged nor scripted.

3.  Zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-4
    x-internal="space-characters"}.

4.  Zero or one of the following:

    - A [width descriptor]{#width-descriptor .dfn}, consisting of:
      [ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-5
      x-internal="space-characters"}, a [valid non-negative
      integer](common-microsyntaxes.html#valid-non-negative-integer){#srcset-attributes:valid-non-negative-integer}
      giving a number greater than zero representing the [width
      descriptor value]{#width-descriptor-value .dfn}, and a U+0077
      LATIN SMALL LETTER W character.

    - A [pixel density descriptor]{#pixel-density-descriptor .dfn},
      consisting of: [ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-6
      x-internal="space-characters"}, a [valid floating-point
      number](common-microsyntaxes.html#valid-floating-point-number){#srcset-attributes:valid-floating-point-number}
      giving a number greater than zero representing the [pixel density
      descriptor value]{#pixel-density-descriptor-value .dfn}, and a
      U+0078 LATIN SMALL LETTER X character.

5.  Zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#srcset-attributes:space-characters-7
    x-internal="space-characters"}.

There must not be an [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-4}
for an element that has the same [width descriptor
value](#width-descriptor-value){#srcset-attributes:width-descriptor-value}
as another [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-5}\'s
[width descriptor
value](#width-descriptor-value){#srcset-attributes:width-descriptor-value-2}
for the same element.

There must not be an [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-6}
for an element that has the same [pixel density descriptor
value](#pixel-density-descriptor-value){#srcset-attributes:pixel-density-descriptor-value}
as another [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-7}\'s
[pixel density descriptor
value](#pixel-density-descriptor-value){#srcset-attributes:pixel-density-descriptor-value-2}
for the same element. For the purpose of this requirement, an [image
candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-8}
with no descriptors is equivalent to an [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-9}
with a `1x` descriptor.

If an [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-10}
for an element has the [width
descriptor](#width-descriptor){#srcset-attributes:width-descriptor}
specified, all other [image candidate
strings](#image-candidate-string){#srcset-attributes:image-candidate-string-11}
for that element must also have the [width
descriptor](#width-descriptor){#srcset-attributes:width-descriptor-2}
specified.

The specified width in an [image candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-12}\'s
[width
descriptor](#width-descriptor){#srcset-attributes:width-descriptor-3}
must match the [natural
width](https://drafts.csswg.org/css-images/#natural-width){#srcset-attributes:natural-width
x-internal="natural-width"} in the resource given by the [image
candidate
string](#image-candidate-string){#srcset-attributes:image-candidate-string-13}\'s
URL, if it has a [natural
width](https://drafts.csswg.org/css-images/#natural-width){#srcset-attributes:natural-width-2
x-internal="natural-width"}.

If an element has a [sizes
attribute](#sizes-attribute){#srcset-attributes:sizes-attribute}
present, all [image candidate
strings](#image-candidate-string){#srcset-attributes:image-candidate-string-14}
for that element must have the [width
descriptor](#width-descriptor){#srcset-attributes:width-descriptor-4}
specified.

###### [4.8.4.2.2]{.secno} Sizes attributes[](#sizes-attributes){.self-link}

A [sizes attribute]{#sizes-attribute .dfn} is an attribute with
requirements defined in this section.

If present, the value must be a [valid source size
list](#valid-source-size-list){#sizes-attributes:valid-source-size-list}.

A [valid source size list]{#valid-source-size-list .dfn} is a string
that matches the following grammar:
[\[CSSVALUES\]](references.html#refsCSSVALUES)
[\[MQ\]](references.html#refsMQ)

``` html
<source-size-list> = <source-size>#? , <source-size-value>
<source-size> = <media-condition> <source-size-value> | auto
<source-size-value> = <length> | auto
```

A
[\<source-size-value\>](#source-size-value){#sizes-attributes:source-size-value-3}
that is a
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#sizes-attributes:length-2-2
x-internal="length-2"} must not be negative, and must not use CSS
functions other than the [math
functions](https://drafts.csswg.org/css-values/#math-function){#sizes-attributes:math-functions
x-internal="math-functions"}.

The keyword [`auto`]{#valdef-sizes-auto .dfn} is a width that is
computed in [parse a sizes
attribute](#parse-a-sizes-attribute){#sizes-attributes:parse-a-sizes-attribute}.
If present, it must be the first entry and the entire
[\<source-size-list\>](#source-size-list){#sizes-attributes:source-size-list}
value must either be the string \"`auto`\" ([ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#sizes-attributes:ascii-case-insensitive
x-internal="ascii-case-insensitive"}) or start with the string
\"`auto,`\" ([ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#sizes-attributes:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"}).

If the
[`img`{#sizes-attributes:the-img-element}](embedded-content.html#the-img-element)
element that initiated the image loading (with the [update the image
data](#update-the-image-data){#sizes-attributes:update-the-image-data}
or [react to environment
changes](#img-environment-changes){#sizes-attributes:img-environment-changes}
algorithms) [allows
auto-sizes](embedded-content.html#allows-auto-sizes){#sizes-attributes:allows-auto-sizes}
and is [being
rendered](rendering.html#being-rendered){#sizes-attributes:being-rendered},
then [`auto`{#sizes-attributes:valdef-sizes-auto-3}](#valdef-sizes-auto)
is the [concrete object
size](https://drafts.csswg.org/css-images/#concrete-object-size){#sizes-attributes:concrete-object-size
x-internal="concrete-object-size"} width. Otherwise, the
[`auto`{#sizes-attributes:valdef-sizes-auto-4}](#valdef-sizes-auto)
value is ignored and the next [source
size](#source-size-2){#sizes-attributes:source-size-2} is used instead,
if any.

The [`auto`{#sizes-attributes:valdef-sizes-auto-5}](#valdef-sizes-auto)
keyword may be specified in the
[`sizes`{#sizes-attributes:attr-source-sizes}](embedded-content.html#attr-source-sizes)
attribute of
[`source`{#sizes-attributes:the-source-element}](embedded-content.html#the-source-element)
elements and
[`sizes`{#sizes-attributes:attr-img-sizes}](embedded-content.html#attr-img-sizes)
attribute of
[`img`{#sizes-attributes:the-img-element-2}](embedded-content.html#the-img-element)
elements, if the following conditions are met. Otherwise,
[`auto`{#sizes-attributes:valdef-sizes-auto-6}](#valdef-sizes-auto) must
not be specified.

- The element is a
  [`source`{#sizes-attributes:the-source-element-2}](embedded-content.html#the-source-element)
  element with a following sibling
  [`img`{#sizes-attributes:the-img-element-3}](embedded-content.html#the-img-element)
  element.

- The element is an
  [`img`{#sizes-attributes:the-img-element-4}](embedded-content.html#the-img-element)
  element.

- The
  [`img`{#sizes-attributes:the-img-element-5}](embedded-content.html#the-img-element)
  element referenced in either condition above [allows
  auto-sizes](embedded-content.html#allows-auto-sizes){#sizes-attributes:allows-auto-sizes-2}.

In addition, it is strongly encouraged to specify dimensions using the
[`width`{#sizes-attributes:attr-dim-width}](embedded-content-other.html#attr-dim-width)
and
[`height`{#sizes-attributes:attr-dim-height}](embedded-content-other.html#attr-dim-height)
attributes or with CSS. Without specified dimensions, the image will
likely render with 300x150 dimensions because `sizes="auto"` implies
`contain-intrinsic-size: 300px 150px` in [the Rendering
section](rendering.html#img-contain-size).

The
[\<source-size-value\>](#source-size-value){#sizes-attributes:source-size-value-4}
gives the intended layout width of the image. The author can specify
different widths for different environments with
[\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition){#sizes-attributes:media-condition-2
x-internal="media-condition"}s.

Percentages are not allowed in a
[\<source-size-value\>](#source-size-value){#sizes-attributes:source-size-value-5},
to avoid confusion about what it would be relative to. The
[\'vw\'](https://drafts.csswg.org/css-values/#vw){#sizes-attributes:'vw'
x-internal="'vw'"} unit can be used for sizes relative to the
[viewport](https://drafts.csswg.org/css2/#viewport){#sizes-attributes:viewport
x-internal="viewport"} width.

##### [4.8.4.3]{.secno} Processing model[](#images-processing-model){.self-link} {#images-processing-model}

------------------------------------------------------------------------

An
[`img`{#images-processing-model:the-img-element}](embedded-content.html#the-img-element)
element has a [current request]{#current-request .dfn} and a [pending
request]{#pending-request .dfn}. The [current
request](#current-request){#images-processing-model:current-request} is
initially set to a new [image
request](#image-request){#images-processing-model:image-request}. The
[pending
request](#pending-request){#images-processing-model:pending-request} is
initially set to null.

An [image request]{#image-request .dfn export=""} has a
[state]{#img-req-state .dfn dfn-for="image request" export=""}, [current
URL]{#img-req-url .dfn dfn-for="image request" export=""}, and [image
data]{#img-req-data .dfn dfn-for="image request" export=""}.

An [image
request](#image-request){#images-processing-model:image-request-2}\'s
[state](#img-req-state){#images-processing-model:img-req-state} is one
of the following:

[Unavailable]{#img-none .dfn}
:   The user agent hasn\'t obtained any image data, or has obtained some
    or all of the image data but hasn\'t yet decoded enough of the image
    to get the image dimensions.

[Partially available]{#img-inc .dfn}
:   The user agent has obtained some of the image data and at least the
    image dimensions are available.

[Completely available]{#img-all .dfn}
:   The user agent has obtained all of the image data and at least the
    image dimensions are available.

[Broken]{#img-error .dfn}
:   The user agent has obtained all of the image data that it can, but
    it cannot even decode the image enough to get the image dimensions
    (e.g. the image is corrupted, or the format is not supported, or no
    data could be obtained).

An [image
request](#image-request){#images-processing-model:image-request-3}\'s
[current URL](#img-req-url){#images-processing-model:img-req-url} is
initially the empty string.

An [image
request](#image-request){#images-processing-model:image-request-4}\'s
[image data](#img-req-data){#images-processing-model:img-req-data} is
the decoded image data.

When an [image
request](#image-request){#images-processing-model:image-request-5}\'s
[state](#img-req-state){#images-processing-model:img-req-state-2} is
either [partially available](#img-inc){#images-processing-model:img-inc}
or [completely available](#img-all){#images-processing-model:img-all},
the [image
request](#image-request){#images-processing-model:image-request-6} is
said to be [available]{#img-available .dfn}.

When an
[`img`{#images-processing-model:the-img-element-2}](embedded-content.html#the-img-element)
element\'s [current
request](#current-request){#images-processing-model:current-request-2}\'s
[state](#img-req-state){#images-processing-model:img-req-state-3} is
[completely available](#img-all){#images-processing-model:img-all-2} and
the user agent can decode the media data without errors, then the
[`img`{#images-processing-model:the-img-element-3}](embedded-content.html#the-img-element)
element is said to be [fully decodable]{#img-good .dfn}.

An [image
request](#image-request){#images-processing-model:image-request-7}\'s
[state](#img-req-state){#images-processing-model:img-req-state-4} is
initially [unavailable](#img-none){#images-processing-model:img-none}.

When an
[`img`{#images-processing-model:the-img-element-4}](embedded-content.html#the-img-element)
element\'s [current
request](#current-request){#images-processing-model:current-request-3}
is [available](#img-available){#images-processing-model:img-available},
the
[`img`{#images-processing-model:the-img-element-5}](embedded-content.html#the-img-element)
element provides a [paint
source](https://drafts.csswg.org/css-images-4/#paint-source){#images-processing-model:paint-source
x-internal="paint-source"} whose width is the image\'s
[density-corrected natural
width](#density-corrected-intrinsic-width-and-height){#images-processing-model:density-corrected-intrinsic-width-and-height}
(if any), whose height is the image\'s [density-corrected natural
height](#density-corrected-intrinsic-width-and-height){#images-processing-model:density-corrected-intrinsic-width-and-height-2}
(if any), and whose appearance is the natural appearance of the image.

------------------------------------------------------------------------

An
[`img`{#images-processing-model:the-img-element-6}](embedded-content.html#the-img-element)
element is said to [use `srcset` or `picture`]{#use-srcset-or-picture
.dfn} if it has a
[`srcset`{#images-processing-model:attr-img-srcset}](embedded-content.html#attr-img-srcset)
attribute specified or if it has a parent that is a
[`picture`{#images-processing-model:the-picture-element}](embedded-content.html#the-picture-element)
element.

------------------------------------------------------------------------

Each
[`img`{#images-processing-model:the-img-element-7}](embedded-content.html#the-img-element)
element has a [last selected source]{#last-selected-source .dfn}, which
must initially be null.

Each [image
request](#image-request){#images-processing-model:image-request-8} has a
[current pixel density]{#current-pixel-density .dfn}, which must
initially be 1.

Each [image
request](#image-request){#images-processing-model:image-request-9} has
[preferred density-corrected
dimensions]{#preferred-density-corrected-dimensions .dfn}, which is
either a struct consisting of a width and a height or is null. It must
initially be null.

To determine the [density-corrected natural width and
height]{#density-corrected-intrinsic-width-and-height .dfn} of an
[`img`{#images-processing-model:the-img-element-8}](embedded-content.html#the-img-element)
element `img`{.variable}:

1.  Let `dim`{.variable} be `img`{.variable}\'s [current
    request](#current-request){#images-processing-model:current-request-4}\'s
    [preferred density-corrected
    dimensions](#preferred-density-corrected-dimensions){#images-processing-model:preferred-density-corrected-dimensions}.

    The [preferred density-corrected
    dimensions](#preferred-density-corrected-dimensions){#images-processing-model:preferred-density-corrected-dimensions-2}
    are set in the [prepare an image for
    presentation](#prepare-an-image-for-presentation){#images-processing-model:prepare-an-image-for-presentation}
    algorithm based on meta information in the image.

2.  If `dim`{.variable} is null, set `dim`{.variable} to
    `img`{.variable}\'s [natural
    dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#images-processing-model:natural-dimensions
    x-internal="natural-dimensions"}.

3.  Set `dim`{.variable}\'s width to `dim`{.variable}\'s width divided
    by `img`{.variable}\'s [current
    request](#current-request){#images-processing-model:current-request-5}\'s
    [current pixel
    density](#current-pixel-density){#images-processing-model:current-pixel-density}.

4.  Set `dim`{.variable}\'s height to `dim`{.variable}\'s height divided
    by `img`{.variable}\'s [current
    request](#current-request){#images-processing-model:current-request-6}\'s
    [current pixel
    density](#current-pixel-density){#images-processing-model:current-pixel-density-2}.

5.  Return `dim`{.variable}.

For example, if the [current pixel
density](#current-pixel-density){#images-processing-model:current-pixel-density-3}
is 3.125, that means that there are 300 device pixels per [CSS
inch](https://drafts.csswg.org/css-values/#in){#images-processing-model:'in'
x-internal="'in'"}, and thus if the image data is 300x600, it has
[density-corrected natural width and
height](#density-corrected-intrinsic-width-and-height){#images-processing-model:density-corrected-intrinsic-width-and-height-3}
of 96 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#images-processing-model:'px'
x-internal="'px'"} by 192 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#images-processing-model:'px'-2
x-internal="'px'"}.

All
[`img`{#images-processing-model:the-img-element-9}](embedded-content.html#the-img-element)
and
[`link`{#images-processing-model:the-link-element}](semantics.html#the-link-element)
elements are associated with a [source
set](#source-set){#images-processing-model:source-set}.

A [source set]{#source-set .dfn} is an ordered set of zero or more
[image sources](#image-source){#images-processing-model:image-source}
and a [source
size](#source-size-2){#images-processing-model:source-size-2}.

An [image source]{#image-source .dfn} is a
[URL](https://url.spec.whatwg.org/#concept-url){#images-processing-model:url
x-internal="url"}, and optionally either a [pixel density
descriptor](#pixel-density-descriptor){#images-processing-model:pixel-density-descriptor},
or a [width
descriptor](#width-descriptor){#images-processing-model:width-descriptor}.

A [source size]{#source-size-2 .dfn} is a
[\<source-size-value\>](#source-size-value){#images-processing-model:source-size-value}.
When a [source
size](#source-size-2){#images-processing-model:source-size-2-2} has a
unit relative to the
[viewport](https://drafts.csswg.org/css2/#viewport){#images-processing-model:viewport
x-internal="viewport"}, it must be interpreted relative to the
[`img`{#images-processing-model:the-img-element-10}](embedded-content.html#the-img-element)
element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#images-processing-model:node-document
x-internal="node-document"}\'s
[viewport](https://drafts.csswg.org/css2/#viewport){#images-processing-model:viewport-2
x-internal="viewport"}. Other units must be interpreted the same as in
Media Queries. [\[MQ\]](references.html#refsMQ)

------------------------------------------------------------------------

A [parse error]{#concept-microsyntax-parse-error .dfn} for algorithms in
this section indicates a non-fatal mismatch between input and
requirements. User agents are encouraged to expose [parse
error](#concept-microsyntax-parse-error){#images-processing-model:concept-microsyntax-parse-error}s
somehow.

------------------------------------------------------------------------

Whether the image is fetched successfully or not (e.g. whether the
response status was an [ok
status](https://fetch.spec.whatwg.org/#ok-status){#images-processing-model:ok-status
x-internal="ok-status"}) must be ignored when determining the image\'s
type and whether it is a valid image.

This allows servers to return images with error responses, and have them
displayed.

The user agent should apply the [image sniffing
rules](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#images-processing-model:content-type-sniffing:-image
x-internal="content-type-sniffing:-image"} to determine the type of the
image, with the image\'s [associated Content-Type
headers](urls-and-fetching.html#content-type){#images-processing-model:content-type}
giving the `official type`{.variable}. If these rules are not applied,
then the type of the image must be the type given by the image\'s
[associated Content-Type
headers](urls-and-fetching.html#content-type){#images-processing-model:content-type-2}.

User agents must not support non-image resources with the
[`img`{#images-processing-model:the-img-element-11}](embedded-content.html#the-img-element)
element (e.g. XML files whose [document
element](https://dom.spec.whatwg.org/#document-element){#images-processing-model:document-element
x-internal="document-element"} is an HTML element). User agents must not
run executable code (e.g. scripts) embedded in the image resource. User
agents must only display the first page of a multipage resource (e.g. a
PDF file). User agents must not allow the resource to act in an
interactive fashion, but should honour any animation in the resource.

This specification does not specify which image types are to be
supported.

###### [4.8.4.3.1]{.secno} When to obtain images[](#when-to-obtain-images){.self-link}

By default, images are obtained immediately. User agents may provide
users with the option to instead obtain them on-demand. (The on-demand
option might be used by bandwidth-constrained users, for example.)

When obtaining images immediately, the user agent must synchronously
[update the image
data](#update-the-image-data){#when-to-obtain-images:update-the-image-data}
of the
[`img`{#when-to-obtain-images:the-img-element}](embedded-content.html#the-img-element)
element, with the *restart animation* flag set if so stated, whenever
that element is created or has experienced [relevant
mutations](#relevant-mutations){#when-to-obtain-images:relevant-mutations}.

When obtaining images on demand, the user agent must [update the image
data](#update-the-image-data){#when-to-obtain-images:update-the-image-data-2}
of an
[`img`{#when-to-obtain-images:the-img-element-2}](embedded-content.html#the-img-element)
element whenever it needs the image data (i.e., on demand), but only if
the
[`img`{#when-to-obtain-images:the-img-element-3}](embedded-content.html#the-img-element)
element\'s [current
request](#current-request){#when-to-obtain-images:current-request}\'s
[state](#img-req-state){#when-to-obtain-images:img-req-state} is
[unavailable](#img-none){#when-to-obtain-images:img-none}. When an
[`img`{#when-to-obtain-images:the-img-element-4}](embedded-content.html#the-img-element)
element has experienced [relevant
mutations](#relevant-mutations){#when-to-obtain-images:relevant-mutations-2},
if the user agent only obtains images on demand, the
[`img`{#when-to-obtain-images:the-img-element-5}](embedded-content.html#the-img-element)
element\'s [current
request](#current-request){#when-to-obtain-images:current-request-2}\'s
[state](#img-req-state){#when-to-obtain-images:img-req-state-2} must
return to [unavailable](#img-none){#when-to-obtain-images:img-none-2}.

###### [4.8.4.3.2]{.secno} Reacting to DOM mutations[](#reacting-to-dom-mutations){.self-link}

The [relevant mutations]{#relevant-mutations .dfn} for an
[`img`{#reacting-to-dom-mutations:the-img-element}](embedded-content.html#the-img-element)
element are as follows:

- The element\'s
  [`src`{#reacting-to-dom-mutations:attr-img-src}](embedded-content.html#attr-img-src),
  [`srcset`{#reacting-to-dom-mutations:attr-img-srcset}](embedded-content.html#attr-img-srcset),
  [`width`{#reacting-to-dom-mutations:attr-dim-width}](embedded-content-other.html#attr-dim-width),
  or
  [`sizes`{#reacting-to-dom-mutations:attr-img-sizes}](embedded-content.html#attr-img-sizes)
  attributes are set, changed, or removed.

- The element\'s
  [`src`{#reacting-to-dom-mutations:attr-img-src-2}](embedded-content.html#attr-img-src)
  attribute is set to the same value as the previous value. This must
  set the *restart animation* flag for the [update the image
  data](#update-the-image-data){#reacting-to-dom-mutations:update-the-image-data}
  algorithm.

- The element\'s
  [`crossorigin`{#reacting-to-dom-mutations:attr-img-crossorigin}](embedded-content.html#attr-img-crossorigin)
  attribute\'s state is changed.

- The element\'s
  [`referrerpolicy`{#reacting-to-dom-mutations:attr-img-referrerpolicy}](embedded-content.html#attr-img-referrerpolicy)
  attribute\'s state is changed.

- The
  [`img`{#reacting-to-dom-mutations:the-img-element-2}](embedded-content.html#the-img-element)
  or
  [`source`{#reacting-to-dom-mutations:the-source-element}](embedded-content.html#the-source-element)
  [HTML element insertion
  steps](infrastructure.html#html-element-insertion-steps){#reacting-to-dom-mutations:html-element-insertion-steps},
  [HTML element removing
  steps](infrastructure.html#html-element-removing-steps){#reacting-to-dom-mutations:html-element-removing-steps},
  and [HTML element moving
  steps](infrastructure.html#html-element-moving-steps){#reacting-to-dom-mutations:html-element-moving-steps}
  count the mutation as a [relevant
  mutation](#relevant-mutations){#reacting-to-dom-mutations:relevant-mutations}.

- The element\'s parent is a
  [`picture`{#reacting-to-dom-mutations:the-picture-element}](embedded-content.html#the-picture-element)
  element and a
  [`source`{#reacting-to-dom-mutations:the-source-element-2}](embedded-content.html#the-source-element)
  element that is a previous sibling has its
  [`srcset`{#reacting-to-dom-mutations:attr-source-srcset}](embedded-content.html#attr-source-srcset),
  [`sizes`{#reacting-to-dom-mutations:attr-source-sizes}](embedded-content.html#attr-source-sizes),
  [`media`{#reacting-to-dom-mutations:attr-source-media}](embedded-content.html#attr-source-media),
  [`type`{#reacting-to-dom-mutations:attr-source-type}](embedded-content.html#attr-source-type),
  [`width`{#reacting-to-dom-mutations:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
  or
  [`height`{#reacting-to-dom-mutations:attr-dim-height}](embedded-content-other.html#attr-dim-height)
  attributes set, changed, or removed.

- The element\'s [adopting
  steps](https://dom.spec.whatwg.org/#concept-node-adopt-ext){#reacting-to-dom-mutations:concept-node-adopt-ext
  x-internal="concept-node-adopt-ext"} are run.

- If the element [allows
  auto-sizes](embedded-content.html#allows-auto-sizes){#reacting-to-dom-mutations:allows-auto-sizes}:
  the element starts or stops [being
  rendered](rendering.html#being-rendered){#reacting-to-dom-mutations:being-rendered},
  or its [concrete object
  size](https://drafts.csswg.org/css-images/#concrete-object-size){#reacting-to-dom-mutations:concrete-object-size
  x-internal="concrete-object-size"} width changes. This must set the
  *maybe omit events* flag for the [update the image
  data](#update-the-image-data){#reacting-to-dom-mutations:update-the-image-data-2}
  algorithm.

###### [4.8.4.3.3]{.secno} The list of available images[](#the-list-of-available-images){.self-link}

Each
[`Document`{#the-list-of-available-images:document}](dom.html#document)
object must have a [list of available images]{#list-of-available-images
.dfn}. Each image in this list is identified by a tuple consisting of an
[absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#the-list-of-available-images:absolute-url
x-internal="absolute-url"}, a [CORS settings
attribute](urls-and-fetching.html#cors-settings-attribute){#the-list-of-available-images:cors-settings-attribute}
mode, and, if the mode is not [No
CORS](urls-and-fetching.html#attr-crossorigin-none){#the-list-of-available-images:attr-crossorigin-none},
an
[origin](browsers.html#concept-origin){#the-list-of-available-images:concept-origin}.
Each image furthermore has an [ignore higher-layer
caching]{#ignore-higher-layer-caching .dfn} flag. User agents may copy
entries from one
[`Document`{#the-list-of-available-images:document-2}](dom.html#document)
object\'s [list of available
images](#list-of-available-images){#the-list-of-available-images:list-of-available-images}
to another at any time (e.g. when the
[`Document`{#the-list-of-available-images:document-3}](dom.html#document)
is created, user agents can add to it all the images that are loaded in
other
[`Document`{#the-list-of-available-images:document-4}](dom.html#document)s),
but must not change the keys of entries copied in this way when doing
so, and must unset the [ignore higher-layer
caching](#ignore-higher-layer-caching){#the-list-of-available-images:ignore-higher-layer-caching}
flag for the copied entry. User agents may also remove images from such
lists at any time (e.g. to save memory). User agents must remove entries
in the [list of available
images](#list-of-available-images){#the-list-of-available-images:list-of-available-images-2}
as appropriate given higher-layer caching semantics for the resource
(e.g. the HTTP
\`[`Cache-Control`{#the-list-of-available-images:http-cache-control}](https://httpwg.org/specs/rfc7234.html#header.cache-control){x-internal="http-cache-control"}\`
response header) when the [ignore higher-layer
caching](#ignore-higher-layer-caching){#the-list-of-available-images:ignore-higher-layer-caching-2}
flag is unset.

The [list of available
images](#list-of-available-images){#the-list-of-available-images:list-of-available-images-3}
is intended to enable synchronous switching when changing the
[`src`{#the-list-of-available-images:attr-img-src}](embedded-content.html#attr-img-src)
attribute to a URL that has previously been loaded, and to avoid
re-downloading images in the same document even when they don\'t allow
caching per HTTP. It is not used to avoid re-downloading the same image
while the previous image is still loading.

The user agent can also store the image data separately from the [list
of available
images](#list-of-available-images){#the-list-of-available-images:list-of-available-images-4}.

For example, if a resource has the HTTP response header
\``Cache-Control: must-revalidate`\`, and its [ignore higher-layer
caching](#ignore-higher-layer-caching){#the-list-of-available-images:ignore-higher-layer-caching-3}
flag is unset, the user agent would remove it from the [list of
available
images](#list-of-available-images){#the-list-of-available-images:list-of-available-images-5}
but could keep the image data separately, and use that if the server
responds with a `304 Not Modified` status.

###### [4.8.4.3.4]{.secno} Decoding images[](#decoding-images){.self-link}

Image data is usually encoded in order to reduce file size. This means
that in order for the user agent to present the image to the screen, the
data needs to be decoded. [Decoding]{#img-decoding-process .dfn} is the
process which converts an image\'s media data into a bitmap form,
suitable for presentation to the screen. Note that this process can be
slow relative to other processes involved in presenting content. Thus,
the user agent can choose when to perform decoding, in order to create
the best user experience.

Image decoding is said to be synchronous if it prevents presentation of
other content until it is finished. Typically, this has an effect of
atomically presenting the image and any other content at the same time.
However, this presentation is delayed by the amount of time it takes to
perform the decode.

Image decoding is said to be asynchronous if it does not prevent
presentation of other content. This has an effect of presenting
non-image content faster. However, the image content is missing on
screen until the decode finishes. Once the decode is finished, the
screen is updated with the image.

In both synchronous and asynchronous decoding modes, the final content
is presented to screen after the same amount of time has elapsed. The
main difference is whether the user agent presents non-image content
ahead of presenting the final content.

In order to aid the user agent in deciding whether to perform
synchronous or asynchronous decode, the
[`decoding`{#decoding-images:attr-img-decoding}](embedded-content.html#attr-img-decoding)
attribute can be set on
[`img`{#decoding-images:the-img-element}](embedded-content.html#the-img-element)
elements. The possible values of the
[`decoding`{#decoding-images:attr-img-decoding-2}](embedded-content.html#attr-img-decoding)
attribute are the following [image decoding hint]{#image-decoding-hint
.dfn} keywords:

Keyword

State

Description

[`sync`]{#attr-img-decoding-sync .dfn dfn-for="img/decoding"
dfn-type="attr-value"}

[Sync]{#attr-img-decoding-sync-state .dfn}

Indicates a preference to
[decode](#img-decoding-process){#decoding-images:img-decoding-process}
this image synchronously for atomic presentation with other content.

[`async`]{#attr-img-decoding-async .dfn dfn-for="img/decoding"
dfn-type="attr-value"}

[Async]{#attr-img-decoding-async-state .dfn}

Indicates a preference to
[decode](#img-decoding-process){#decoding-images:img-decoding-process-2}
this image asynchronously to avoid delaying presentation of other
content.

[`auto`]{#attr-img-decoding-auto .dfn dfn-for="img/decoding"
dfn-type="attr-value"}

[Auto]{#attr-img-decoding-auto-state .dfn}

Indicates no preference in decoding mode (the default).

When
[decoding](#img-decoding-process){#decoding-images:img-decoding-process-3}
an image, the user agent should respect the preference indicated by the
[`decoding`{#decoding-images:attr-img-decoding-3}](embedded-content.html#attr-img-decoding)
attribute\'s state. If the state indicated is
[Auto](#attr-img-decoding-auto-state){#decoding-images:attr-img-decoding-auto-state},
then the user agent is free to choose any decoding behavior.

It is also possible to control the decoding behavior using the
[`decode()`{#decoding-images:dom-img-decode}](embedded-content.html#dom-img-decode)
method. Since the
[`decode()`{#decoding-images:dom-img-decode-2}](embedded-content.html#dom-img-decode)
method performs
[decoding](#img-decoding-process){#decoding-images:img-decoding-process-4}
independently from the process responsible for presenting content to
screen, it is unaffected by the
[`decoding`{#decoding-images:attr-img-decoding-4}](embedded-content.html#attr-img-decoding)
attribute.

###### [4.8.4.3.5]{.secno} Updating the image data[](#updating-the-image-data){.self-link}

This algorithm cannot be called from steps running [in
parallel](infrastructure.html#in-parallel){#updating-the-image-data:in-parallel}.
If a user agent needs to call this algorithm from steps running [in
parallel](infrastructure.html#in-parallel){#updating-the-image-data:in-parallel-2},
it needs to
[queue](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task}
a task to do so.

When the user agent is to [update the image data]{#update-the-image-data
.dfn} of an
[`img`{#updating-the-image-data:the-img-element}](embedded-content.html#the-img-element)
element, optionally with the *restart animations* flag set, optionally
with the *maybe omit events* flag set, it must run the following steps:

1.  If the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#updating-the-image-data:fully-active},
    then:

    1.  Continue running this algorithm [in
        parallel](infrastructure.html#in-parallel){#updating-the-image-data:in-parallel-3}.

    2.  Wait until the element\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-2
        x-internal="node-document"} is [fully
        active](document-sequences.html#fully-active){#updating-the-image-data:fully-active-2}.

    3.  If another instance of this algorithm for this
        [`img`{#updating-the-image-data:the-img-element-2}](embedded-content.html#the-img-element)
        element was started after this instance (even if it aborted and
        is no longer running), then return.

    4.  [Queue a
        microtask](webappapis.html#queue-a-microtask){#updating-the-image-data:queue-a-microtask}
        to continue this algorithm.

2.  If the user agent cannot support images, or its support for images
    has been disabled, then [abort the image
    request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request}
    for the [current
    request](#current-request){#updating-the-image-data:current-request}
    and the [pending
    request](#pending-request){#updating-the-image-data:pending-request},
    set the [current
    request](#current-request){#updating-the-image-data:current-request-2}\'s
    [state](#img-req-state){#updating-the-image-data:img-req-state} to
    [unavailable](#img-none){#updating-the-image-data:img-none}, set the
    [pending
    request](#pending-request){#updating-the-image-data:pending-request-2}
    to null, and return.

3.  Let `previousURL`{.variable} be the [current
    request](#current-request){#updating-the-image-data:current-request-3}\'s
    [current URL](#img-req-url){#updating-the-image-data:img-req-url}.

4.  Let `selected source`{.variable} be null and
    `selected pixel density`{.variable} be undefined.

5.  If the element does not [use `srcset` or
    `picture`](#use-srcset-or-picture){#updating-the-image-data:use-srcset-or-picture}
    and it has a
    [`src`{#updating-the-image-data:attr-img-src}](embedded-content.html#attr-img-src)
    attribute specified whose value is not the empty string, then set
    `selected source`{.variable} to the value of the element\'s
    [`src`{#updating-the-image-data:attr-img-src-2}](embedded-content.html#attr-img-src)
    attribute and set `selected pixel density`{.variable} to 1.0.

6.  Set the element\'s [last selected
    source](#last-selected-source){#updating-the-image-data:last-selected-source}
    to `selected source`{.variable}.

7.  If `selected source`{.variable} is not null, then:

    1.  Let `urlString`{.variable} be the result of
        [encoding-parsing-and-serializing a
        URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#updating-the-image-data:encoding-parsing-and-serializing-a-url}
        given `selected source`{.variable}, relative to the element\'s
        [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-3
        x-internal="node-document"}.

    2.  If `urlString`{.variable} is failure, then abort this inner set
        of steps.

    3.  Let `key`{.variable} be a tuple consisting of
        `urlString`{.variable}, the
        [`img`{#updating-the-image-data:the-img-element-3}](embedded-content.html#the-img-element)
        element\'s
        [`crossorigin`{#updating-the-image-data:attr-img-crossorigin}](embedded-content.html#attr-img-crossorigin)
        attribute\'s mode, and, if that mode is not [No
        CORS](urls-and-fetching.html#attr-crossorigin-none){#updating-the-image-data:attr-crossorigin-none},
        the [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-4
        x-internal="node-document"}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#updating-the-image-data:concept-document-origin
        x-internal="concept-document-origin"}.

    4.  If the [list of available
        images](#list-of-available-images){#updating-the-image-data:list-of-available-images}
        contains an entry for `key`{.variable}, then:

        1.  Set the [ignore higher-layer
            caching](#ignore-higher-layer-caching){#updating-the-image-data:ignore-higher-layer-caching}
            flag for that entry.

        2.  [Abort the image
            request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-2}
            for the [current
            request](#current-request){#updating-the-image-data:current-request-4}
            and the [pending
            request](#pending-request){#updating-the-image-data:pending-request-3}.

        3.  Set the [pending
            request](#pending-request){#updating-the-image-data:pending-request-4}
            to null.

        4.  Set the [current
            request](#current-request){#updating-the-image-data:current-request-5}
            to a new [image
            request](#image-request){#updating-the-image-data:image-request}
            whose [image
            data](#img-req-data){#updating-the-image-data:img-req-data}
            is that of the entry and whose
            [state](#img-req-state){#updating-the-image-data:img-req-state-2}
            is [completely
            available](#img-all){#updating-the-image-data:img-all}.

        5.  [Prepare the current request for
            presentation](#prepare-an-image-for-presentation){#updating-the-image-data:prepare-an-image-for-presentation}
            given the
            [`img`{#updating-the-image-data:the-img-element-4}](embedded-content.html#the-img-element)
            element.

        6.  Set the [current
            request](#current-request){#updating-the-image-data:current-request-6}\'s
            [current pixel
            density](#current-pixel-density){#updating-the-image-data:current-pixel-density}
            to `selected pixel density`{.variable}.

        7.  [Queue an element
            task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source}
            given the
            [`img`{#updating-the-image-data:the-img-element-5}](embedded-content.html#the-img-element)
            element and the following steps:

            1.  If *restart animation* is set, then [restart the
                animation](rendering.html#restart-the-animation){#updating-the-image-data:restart-the-animation}.

            2.  Set the [current
                request](#current-request){#updating-the-image-data:current-request-7}\'s
                [current
                URL](#img-req-url){#updating-the-image-data:img-req-url-2}
                to `urlString`{.variable}.

            3.  If *maybe omit events* is not set or
                `previousURL`{.variable} is not equal to
                `urlString`{.variable}, then [fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire
                x-internal="concept-event-fire"} named
                [`load`{#updating-the-image-data:event-load}](indices.html#event-load)
                at the
                [`img`{#updating-the-image-data:the-img-element-6}](embedded-content.html#the-img-element)
                element.

        8.  Abort the [update the image
            data](#update-the-image-data){#updating-the-image-data:update-the-image-data}
            algorithm.

8.  [Queue a
    microtask](webappapis.html#queue-a-microtask){#updating-the-image-data:queue-a-microtask-2}
    to perform the rest of this algorithm, allowing the
    [task](webappapis.html#concept-task){#updating-the-image-data:concept-task}
    that invoked this algorithm to continue.

9.  If another instance of this algorithm for this
    [`img`{#updating-the-image-data:the-img-element-7}](embedded-content.html#the-img-element)
    element was started after this instance (even if it aborted and is
    no longer running), then return.

    Only the last instance takes effect, to avoid multiple requests
    when, for example, the
    [`src`{#updating-the-image-data:attr-img-src-3}](embedded-content.html#attr-img-src),
    [`srcset`{#updating-the-image-data:attr-img-srcset}](embedded-content.html#attr-img-srcset),
    and
    [`crossorigin`{#updating-the-image-data:attr-img-crossorigin-2}](embedded-content.html#attr-img-crossorigin)
    attributes are all set in succession.

10. Let `selected source`{.variable} and
    `selected pixel density`{.variable} be the URL and pixel density
    that results from [selecting an image
    source](#select-an-image-source){#updating-the-image-data:select-an-image-source},
    respectively.

11. If `selected source`{.variable} is null, then:

    1.  Set the [current
        request](#current-request){#updating-the-image-data:current-request-8}\'s
        [state](#img-req-state){#updating-the-image-data:img-req-state-3}
        to [broken](#img-error){#updating-the-image-data:img-error},
        [abort the image
        request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-3}
        for the [current
        request](#current-request){#updating-the-image-data:current-request-9}
        and the [pending
        request](#pending-request){#updating-the-image-data:pending-request-5},
        and set the [pending
        request](#pending-request){#updating-the-image-data:pending-request-6}
        to null.

    2.  [Queue an element
        task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task-2}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source-2}
        given the
        [`img`{#updating-the-image-data:the-img-element-8}](embedded-content.html#the-img-element)
        element and the following steps:

        1.  Change the [current
            request](#current-request){#updating-the-image-data:current-request-10}\'s
            [current
            URL](#img-req-url){#updating-the-image-data:img-req-url-3}
            to the empty string.

        2.  If all of the following are true:

            - the element has a
              [`src`{#updating-the-image-data:attr-img-src-4}](embedded-content.html#attr-img-src)
              attribute or it [uses `srcset` or
              `picture`](#use-srcset-or-picture){#updating-the-image-data:use-srcset-or-picture-2};
              and

            - *maybe omit events* is not set or `previousURL`{.variable}
              is not the empty string,

            then [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`error`{#updating-the-image-data:event-error}](indices.html#event-error)
            at the
            [`img`{#updating-the-image-data:the-img-element-9}](embedded-content.html#the-img-element)
            element.

    3.  Return.

12. Let `urlString`{.variable} be the result of
    [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#updating-the-image-data:encoding-parsing-and-serializing-a-url-2}
    given `selected source`{.variable}, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-5
    x-internal="node-document"}.

13. If `urlString`{.variable} is failure, then:

    1.  [Abort the image
        request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-4}
        for the [current
        request](#current-request){#updating-the-image-data:current-request-11}
        and the [pending
        request](#pending-request){#updating-the-image-data:pending-request-7}.

    2.  Set the [current
        request](#current-request){#updating-the-image-data:current-request-12}\'s
        [state](#img-req-state){#updating-the-image-data:img-req-state-4}
        to [broken](#img-error){#updating-the-image-data:img-error-2}.

    3.  Set the [pending
        request](#pending-request){#updating-the-image-data:pending-request-8}
        to null.

    4.  [Queue an element
        task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task-3}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source-3}
        given the
        [`img`{#updating-the-image-data:the-img-element-10}](embedded-content.html#the-img-element)
        element and the following steps:

        1.  Change the [current
            request](#current-request){#updating-the-image-data:current-request-13}\'s
            [current
            URL](#img-req-url){#updating-the-image-data:img-req-url-4}
            to `selected source`{.variable}.

        2.  If *maybe omit events* is not set or
            `previousURL`{.variable} is not equal to
            `selected source`{.variable}, then [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-3
            x-internal="concept-event-fire"} named
            [`error`{#updating-the-image-data:event-error-2}](indices.html#event-error)
            at the
            [`img`{#updating-the-image-data:the-img-element-11}](embedded-content.html#the-img-element)
            element.

    5.  Return.

14. If the [pending
    request](#pending-request){#updating-the-image-data:pending-request-9}
    is not null and `urlString`{.variable} is the same as the [pending
    request](#pending-request){#updating-the-image-data:pending-request-10}\'s
    [current URL](#img-req-url){#updating-the-image-data:img-req-url-5},
    then return.

15. If `urlString`{.variable} is the same as the [current
    request](#current-request){#updating-the-image-data:current-request-14}\'s
    [current URL](#img-req-url){#updating-the-image-data:img-req-url-6}
    and the [current
    request](#current-request){#updating-the-image-data:current-request-15}\'s
    [state](#img-req-state){#updating-the-image-data:img-req-state-5} is
    [partially available](#img-inc){#updating-the-image-data:img-inc},
    then [abort the image
    request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-5}
    for the [pending
    request](#pending-request){#updating-the-image-data:pending-request-11},
    [queue an element
    task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task-4}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source-4}
    given the
    [`img`{#updating-the-image-data:the-img-element-12}](embedded-content.html#the-img-element)
    element to [restart the
    animation](rendering.html#restart-the-animation){#updating-the-image-data:restart-the-animation-2}
    if *restart animation* is set, and return.

16. [Abort the image
    request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-6}
    for the [pending
    request](#pending-request){#updating-the-image-data:pending-request-12}.

17. Set `image request`{.variable} to a new [image
    request](#image-request){#updating-the-image-data:image-request-2}
    whose [current
    URL](#img-req-url){#updating-the-image-data:img-req-url-7} is
    `urlString`{.variable}.

18. If the [current
    request](#current-request){#updating-the-image-data:current-request-16}\'s
    [state](#img-req-state){#updating-the-image-data:img-req-state-6} is
    [unavailable](#img-none){#updating-the-image-data:img-none-2} or
    [broken](#img-error){#updating-the-image-data:img-error-3}, then set
    the [current
    request](#current-request){#updating-the-image-data:current-request-17}
    to `image request`{.variable}. Otherwise, set the [pending
    request](#pending-request){#updating-the-image-data:pending-request-13}
    to `image request`{.variable}.

19. Let `request`{.variable} be the result of [creating a potential-CORS
    request](urls-and-fetching.html#create-a-potential-cors-request){#updating-the-image-data:create-a-potential-cors-request}
    given `urlString`{.variable}, \"`image`\", and the current state of
    the element\'s
    [`crossorigin`{#updating-the-image-data:attr-img-crossorigin-3}](embedded-content.html#attr-img-crossorigin)
    content attribute.

20. Set `request`{.variable}\'s
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#updating-the-image-data:concept-request-client
    x-internal="concept-request-client"} to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-6
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#updating-the-image-data:relevant-settings-object}.

21. If the element [uses `srcset` or
    `picture`](#use-srcset-or-picture){#updating-the-image-data:use-srcset-or-picture-3},
    set `request`{.variable}\'s
    [initiator](https://fetch.spec.whatwg.org/#concept-request-initiator){#updating-the-image-data:concept-request-initiator
    x-internal="concept-request-initiator"} to \"`imageset`\".

22. Set `request`{.variable}\'s [referrer
    policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#updating-the-image-data:concept-request-referrer-policy
    x-internal="concept-request-referrer-policy"} to the current state
    of the element\'s
    [`referrerpolicy`{#updating-the-image-data:attr-img-referrerpolicy}](embedded-content.html#attr-img-referrerpolicy)
    attribute.

23. Set `request`{.variable}\'s
    [priority](https://fetch.spec.whatwg.org/#request-priority){#updating-the-image-data:concept-request-priority
    x-internal="concept-request-priority"} to the current state of the
    element\'s
    [`fetchpriority`{#updating-the-image-data:attr-img-fetchpriority}](embedded-content.html#attr-img-fetchpriority)
    attribute.

24. Let `delay load event`{.variable} be true if the
    [`img`{#updating-the-image-data:the-img-element-13}](embedded-content.html#the-img-element)\'s
    [lazy loading
    attribute](urls-and-fetching.html#lazy-loading-attribute){#updating-the-image-data:lazy-loading-attribute}
    is in the
    [Eager](urls-and-fetching.html#attr-loading-eager-state){#updating-the-image-data:attr-loading-eager-state}
    state, or if [scripting is
    disabled](webappapis.html#concept-n-noscript){#updating-the-image-data:concept-n-noscript}
    for the
    [`img`{#updating-the-image-data:the-img-element-14}](embedded-content.html#the-img-element),
    and false otherwise.

25. If the [will lazy load element
    steps](urls-and-fetching.html#will-lazy-load-element-steps){#updating-the-image-data:will-lazy-load-element-steps}
    given the
    [`img`{#updating-the-image-data:the-img-element-15}](embedded-content.html#the-img-element)
    return true, then:

    1.  Set the
        [`img`{#updating-the-image-data:the-img-element-16}](embedded-content.html#the-img-element)\'s
        [lazy load resumption
        steps](urls-and-fetching.html#lazy-load-resumption-steps){#updating-the-image-data:lazy-load-resumption-steps}
        to the rest of this algorithm starting with the step labeled
        *fetch the image*.

    2.  [Start intersection-observing a lazy loading
        element](urls-and-fetching.html#start-intersection-observing-a-lazy-loading-element){#updating-the-image-data:start-intersection-observing-a-lazy-loading-element}
        for the
        [`img`{#updating-the-image-data:the-img-element-17}](embedded-content.html#the-img-element)
        element.

    3.  Return.

26. *Fetch the image*:
    [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#updating-the-image-data:concept-fetch
    x-internal="concept-fetch"} `request`{.variable}. Return from this
    algorithm, and run the remaining steps as part of the fetch\'s
    *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
    for the
    [response](https://fetch.spec.whatwg.org/#concept-response){#updating-the-image-data:concept-response
    x-internal="concept-response"} `response`{.variable}.

    The resource obtained in this fashion, if any, is
    `image request`{.variable}\'s [image
    data](#img-req-data){#updating-the-image-data:img-req-data-2}. It
    can be either
    [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#updating-the-image-data:cors-same-origin}
    or
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#updating-the-image-data:cors-cross-origin};
    this affects the image\'s interaction with other APIs (e.g., when
    used on a
    [`canvas`{#updating-the-image-data:the-canvas-element}](canvas.html#the-canvas-element)).

    When `delay load event`{.variable} is true, fetching the image must
    [delay the load
    event](parsing.html#delay-the-load-event){#updating-the-image-data:delay-the-load-event}
    of the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-7
    x-internal="node-document"} until the
    [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-2}
    that is
    [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-2}
    by the [networking task
    source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source}
    once the resource has been fetched ([defined below](#img-load)) has
    been run.

    This, unfortunately, can be used to perform a rudimentary port scan
    of the user\'s local network (especially in conjunction with
    scripting, though scripting isn\'t actually necessary to carry out
    such an attack). User agents may implement
    [cross-origin](browsers.html#concept-origin){#updating-the-image-data:concept-origin}
    access control policies that are stricter than those described above
    to mitigate this attack, but unfortunately such policies are
    typically not compatible with existing web content.

27. As soon as possible, jump to the first applicable entry from the
    following list:

    If the resource type is [`multipart/x-mixed-replace`{#updating-the-image-data:multipart/x-mixed-replace}](iana.html#multipart/x-mixed-replace)

    :   The next
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-3}
        that is
        [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-3}
        by the [networking task
        source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source-2}
        while the image is being fetched must run the following steps:

        1.  If `image request`{.variable} is the [pending
            request](#pending-request){#updating-the-image-data:pending-request-14}
            and at least one body part has been completely decoded,
            [abort the image
            request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-7}
            for the [current
            request](#current-request){#updating-the-image-data:current-request-18},
            and [upgrade the pending request to the current
            request](#upgrade-the-pending-request-to-the-current-request){#updating-the-image-data:upgrade-the-pending-request-to-the-current-request}.

        2.  Otherwise, if `image request`{.variable} is the [pending
            request](#pending-request){#updating-the-image-data:pending-request-15}
            and the user agent is able to determine that
            `image request`{.variable}\'s image is corrupted in some
            fatal way such that the image dimensions cannot be obtained,
            [abort the image
            request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-8}
            for the [current
            request](#current-request){#updating-the-image-data:current-request-19},
            [upgrade the pending request to the current
            request](#upgrade-the-pending-request-to-the-current-request){#updating-the-image-data:upgrade-the-pending-request-to-the-current-request-2},
            and set the [current
            request](#current-request){#updating-the-image-data:current-request-20}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-7}
            to
            [broken](#img-error){#updating-the-image-data:img-error-4}.

        3.  Otherwise, if `image request`{.variable} is the [current
            request](#current-request){#updating-the-image-data:current-request-21},
            its
            [state](#img-req-state){#updating-the-image-data:img-req-state-8}
            is
            [unavailable](#img-none){#updating-the-image-data:img-none-3},
            and the user agent is able to determine
            `image request`{.variable}\'s image\'s width and height, set
            the [current
            request](#current-request){#updating-the-image-data:current-request-22}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-9}
            to [partially
            available](#img-inc){#updating-the-image-data:img-inc-2}.

        4.  Otherwise, if `image request`{.variable} is the [current
            request](#current-request){#updating-the-image-data:current-request-23},
            its
            [state](#img-req-state){#updating-the-image-data:img-req-state-10}
            is
            [unavailable](#img-none){#updating-the-image-data:img-none-4},
            and the user agent is able to determine that
            `image request`{.variable}\'s image is corrupted in some
            fatal way such that the image dimensions cannot be obtained,
            set the [current
            request](#current-request){#updating-the-image-data:current-request-24}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-11}
            to
            [broken](#img-error){#updating-the-image-data:img-error-5}.

        Each
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-4}
        that is
        [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-4}
        by the [networking task
        source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source-3}
        while the image is being fetched must update the presentation of
        the image, but as each new body part comes in, if the user agent
        is able to determine the image\'s width and height, it must
        [prepare the `img` element\'s current request for
        presentation](#prepare-an-image-for-presentation){#updating-the-image-data:prepare-an-image-for-presentation-2}
        given the
        [`img`{#updating-the-image-data:the-img-element-18}](embedded-content.html#the-img-element)
        element and replace the previous image. Once one body part has
        been completely decoded, perform the following steps:

        1.  Set the
            [`img`{#updating-the-image-data:the-img-element-19}](embedded-content.html#the-img-element)
            element\'s [current
            request](#current-request){#updating-the-image-data:current-request-25}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-12}
            to [completely
            available](#img-all){#updating-the-image-data:img-all-2}.

        2.  If *maybe omit events* is not set or
            `previousURL`{.variable} is not equal to
            `urlString`{.variable}, then [queue an element
            task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task-5}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source-5}
            given the
            [`img`{#updating-the-image-data:the-img-element-20}](embedded-content.html#the-img-element)
            element to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-4
            x-internal="concept-event-fire"} named
            [`load`{#updating-the-image-data:event-load-2}](indices.html#event-load)
            at the
            [`img`{#updating-the-image-data:the-img-element-21}](embedded-content.html#the-img-element)
            element.

    If the resource type and data corresponds to a supported image format, [as described below](#img-determine-type)

    :   The next
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-5}
        that is
        [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-5}
        by the [networking task
        source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source-4}
        while the image is being fetched must run the following steps:

        1.  If the user agent is able to determine
            `image request`{.variable}\'s image\'s width and height, and
            `image request`{.variable} is the [pending
            request](#pending-request){#updating-the-image-data:pending-request-16},
            set `image request`{.variable}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-13}
            to [partially
            available](#img-inc){#updating-the-image-data:img-inc-3}.

        2.  Otherwise, if the user agent is able to determine
            `image request`{.variable}\'s image\'s width and height, and
            `image request`{.variable} is the [current
            request](#current-request){#updating-the-image-data:current-request-26},
            [prepare `image request`{.variable} for
            presentation](#prepare-an-image-for-presentation){#updating-the-image-data:prepare-an-image-for-presentation-3}
            given the
            [`img`{#updating-the-image-data:the-img-element-22}](embedded-content.html#the-img-element)
            element and set `image request`{.variable}\'s
            [state](#img-req-state){#updating-the-image-data:img-req-state-14}
            to [partially
            available](#img-inc){#updating-the-image-data:img-inc-4}.

        3.  Otherwise, if the user agent is able to determine that
            `image request`{.variable}\'s image is corrupted in some
            fatal way such that the image dimensions cannot be obtained,
            and `image request`{.variable} is the [pending
            request](#pending-request){#updating-the-image-data:pending-request-17}:

            1.  [Abort the image
                request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-9}
                for the [current
                request](#current-request){#updating-the-image-data:current-request-27}
                and the [pending
                request](#pending-request){#updating-the-image-data:pending-request-18}.

            2.  [Upgrade the pending request to the current
                request](#upgrade-the-pending-request-to-the-current-request){#updating-the-image-data:upgrade-the-pending-request-to-the-current-request-3}.

            3.  Set the [current
                request](#current-request){#updating-the-image-data:current-request-28}\'s
                [state](#img-req-state){#updating-the-image-data:img-req-state-15}
                to
                [broken](#img-error){#updating-the-image-data:img-error-6}.

            4.  [Fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-5
                x-internal="concept-event-fire"} named
                [`error`{#updating-the-image-data:event-error-3}](indices.html#event-error)
                at the
                [`img`{#updating-the-image-data:the-img-element-23}](embedded-content.html#the-img-element)
                element.

        4.  Otherwise, if the user agent is able to determine that
            `image request`{.variable}\'s image is corrupted in some
            fatal way such that the image dimensions cannot be obtained,
            and `image request`{.variable} is the [current
            request](#current-request){#updating-the-image-data:current-request-29}:

            1.  [Abort the image
                request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-10}
                for `image request`{.variable}.

            2.  If *maybe omit events* is not set or
                `previousURL`{.variable} is not equal to
                `urlString`{.variable}, then [fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-6
                x-internal="concept-event-fire"} named
                [`error`{#updating-the-image-data:event-error-4}](indices.html#event-error)
                at the
                [`img`{#updating-the-image-data:the-img-element-24}](embedded-content.html#the-img-element)
                element.

        That
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-6},
        and each subsequent
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-7},
        that is
        [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-6}
        by the [networking task
        source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source-5}
        while the image is being fetched, if `image request`{.variable}
        is the [current
        request](#current-request){#updating-the-image-data:current-request-30},
        must update the presentation of the image appropriately (e.g.,
        if the image is a progressive JPEG, each packet can improve the
        resolution of the image).

        Furthermore, the last
        [task](webappapis.html#concept-task){#updating-the-image-data:concept-task-8}
        that is
        [queued](webappapis.html#queue-a-task){#updating-the-image-data:queue-a-task-7}
        by the [networking task
        source](webappapis.html#networking-task-source){#updating-the-image-data:networking-task-source-6}
        once the resource has been fetched must additionally run these
        steps:

        1.  If `image request`{.variable} is the [pending
            request](#pending-request){#updating-the-image-data:pending-request-19},
            [abort the image
            request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-11}
            for the [current
            request](#current-request){#updating-the-image-data:current-request-31},
            [upgrade the pending request to the current
            request](#upgrade-the-pending-request-to-the-current-request){#updating-the-image-data:upgrade-the-pending-request-to-the-current-request-4},
            and [prepare `image request`{.variable} for
            presentation](#prepare-an-image-for-presentation){#updating-the-image-data:prepare-an-image-for-presentation-4}
            given the
            [`img`{#updating-the-image-data:the-img-element-25}](embedded-content.html#the-img-element)
            element.

        2.  Set `image request`{.variable} to the [completely
            available](#img-all){#updating-the-image-data:img-all-3}
            state.

        3.  Add the image to the [list of available
            images](#list-of-available-images){#updating-the-image-data:list-of-available-images-2}
            using the key `key`{.variable}, with the [ignore
            higher-layer
            caching](#ignore-higher-layer-caching){#updating-the-image-data:ignore-higher-layer-caching-2}
            flag set.

        4.  If *maybe omit events* is not set or
            `previousURL`{.variable} is not equal to
            `urlString`{.variable}, then [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-7
            x-internal="concept-event-fire"} named
            [`load`{#updating-the-image-data:event-load-3}](indices.html#event-load)
            at the
            [`img`{#updating-the-image-data:the-img-element-26}](embedded-content.html#the-img-element)
            element.

    Otherwise

    :   The image data is not in a supported file format; the user agent
        must set `image request`{.variable}\'s
        [state](#img-req-state){#updating-the-image-data:img-req-state-16}
        to [broken](#img-error){#updating-the-image-data:img-error-7},
        [abort the image
        request](#abort-the-image-request){#updating-the-image-data:abort-the-image-request-12}
        for the [current
        request](#current-request){#updating-the-image-data:current-request-32}
        and the [pending
        request](#pending-request){#updating-the-image-data:pending-request-20},
        [upgrade the pending request to the current
        request](#upgrade-the-pending-request-to-the-current-request){#updating-the-image-data:upgrade-the-pending-request-to-the-current-request-5}
        if `image request`{.variable} is the [pending
        request](#pending-request){#updating-the-image-data:pending-request-21},
        and then, if *maybe omit events* is not set or
        `previousURL`{.variable} is not equal to `urlString`{.variable},
        [queue an element
        task](webappapis.html#queue-an-element-task){#updating-the-image-data:queue-an-element-task-6}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#updating-the-image-data:dom-manipulation-task-source-6}
        given the
        [`img`{#updating-the-image-data:the-img-element-27}](embedded-content.html#the-img-element)
        element to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#updating-the-image-data:concept-event-fire-8
        x-internal="concept-event-fire"} named
        [`error`{#updating-the-image-data:event-error-5}](indices.html#event-error)
        at the
        [`img`{#updating-the-image-data:the-img-element-28}](embedded-content.html#the-img-element)
        element.

While a user agent is running the above algorithm for an element
`x`{.variable}, there must be a strong reference from the element\'s
[node
document](https://dom.spec.whatwg.org/#concept-node-document){#updating-the-image-data:node-document-8
x-internal="node-document"} to the element `x`{.variable}, even if that
element is not
[connected](https://dom.spec.whatwg.org/#connected){#updating-the-image-data:connected
x-internal="connected"}.

To [abort the image request]{#abort-the-image-request .dfn} for an
[image
request](#image-request){#updating-the-image-data:image-request-3} or
null `image request`{.variable} means to run the following steps:

1.  If `image request`{.variable} is null, then return.

2.  Forget `image request`{.variable}\'s [image
    data](#img-req-data){#updating-the-image-data:img-req-data-3}, if
    any.

3.  Abort any instance of the
    [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#updating-the-image-data:concept-fetch-2
    x-internal="concept-fetch"} algorithm for
    `image request`{.variable}, discarding any pending tasks generated
    by that algorithm.

To [upgrade the pending request to the current
request]{#upgrade-the-pending-request-to-the-current-request .dfn} for
an
[`img`{#updating-the-image-data:the-img-element-29}](embedded-content.html#the-img-element)
element means to run the following steps:

1.  Set the
    [`img`{#updating-the-image-data:the-img-element-30}](embedded-content.html#the-img-element)
    element\'s [current
    request](#current-request){#updating-the-image-data:current-request-33}
    to the [pending
    request](#pending-request){#updating-the-image-data:pending-request-22}.

2.  Set the
    [`img`{#updating-the-image-data:the-img-element-31}](embedded-content.html#the-img-element)
    element\'s [pending
    request](#pending-request){#updating-the-image-data:pending-request-23}
    to null.

###### [4.8.4.3.6]{.secno} Preparing an image for presentation[](#preparing-an-image-for-presentation){.self-link}

To [prepare an image for
presentation]{#prepare-an-image-for-presentation .dfn} for an [image
request](#image-request){#preparing-an-image-for-presentation:image-request}
`req`{.variable} given image element `img`{.variable}:

1.  Let `exifTagMap`{.variable} be the EXIF tags obtained from
    `req`{.variable}\'s [image
    data](#img-req-data){#preparing-an-image-for-presentation:img-req-data},
    as defined by the relevant codec.
    [\[EXIF\]](references.html#refsEXIF)

2.  Let `physicalWidth`{.variable} and `physicalHeight`{.variable} be
    the width and height obtained from `req`{.variable}\'s [image
    data](#img-req-data){#preparing-an-image-for-presentation:img-req-data-2},
    as defined by the relevant codec.

3.  Let `dimX`{.variable} be the value of `exifTagMap`{.variable}\'s tag
    `0xA002` (`PixelXDimension`).

4.  Let `dimY`{.variable} be the value of `exifTagMap`{.variable}\'s tag
    `0xA003` (`PixelYDimension`).

5.  Let `resX`{.variable} be the value of `exifTagMap`{.variable}\'s tag
    `0x011A` (`XResolution`).

6.  Let `resY`{.variable} be the value of `exifTagMap`{.variable}\'s tag
    `0x011B` (`YResolution`).

7.  Let `resUnit`{.variable} be the value of `exifTagMap`{.variable}\'s
    tag `0x0128` (`ResolutionUnit`).

8.  If either `dimX`{.variable} or `dimY`{.variable} is not a positive
    integer, then return.

9.  If either `resX`{.variable} or `resY`{.variable} is not a positive
    floating-point number, then return.

10. If `resUnit`{.variable} is not equal to `2` (`Inch`), then return.

11. Let `widthFromDensity`{.variable} be the value of
    `physicalWidth`{.variable}, multiplied by 72 and divided by
    `resX`{.variable}.

12. Let `heightFromDensity`{.variable} be the value of
    `physicalHeight`{.variable}, multiplied by 72 and divided by
    `resY`{.variable}.

13. If `widthFromDensity`{.variable} is not equal to `dimX`{.variable}
    or `heightFromDensity`{.variable} is not equal to `dimY`{.variable},
    then return.

14. If `req`{.variable}\'s [image
    data](#img-req-data){#preparing-an-image-for-presentation:img-req-data-3}
    is
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#preparing-an-image-for-presentation:cors-cross-origin},
    then set `img`{.variable}\'s [natural
    dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#preparing-an-image-for-presentation:natural-dimensions
    x-internal="natural-dimensions"} to `dimX`{.variable} and
    `dimY`{.variable}, scale `img`{.variable}\'s pixel data accordingly,
    and return.

15. Set `req`{.variable}\'s [preferred density-corrected
    dimensions](#preferred-density-corrected-dimensions){#preparing-an-image-for-presentation:preferred-density-corrected-dimensions}
    to a struct with its width set to `dimX`{.variable} and its height
    set to `dimY`{.variable}.

16. Update `req`{.variable}\'s
    [`img`{#preparing-an-image-for-presentation:the-img-element}](embedded-content.html#the-img-element)
    element\'s presentation appropriately.

Resolution in EXIF is equivalent to [CSS points per
inch](https://drafts.csswg.org/css-values/#pt){#preparing-an-image-for-presentation:'pt'
x-internal="'pt'"}, therefore 72 is the base for computing size from
resolution.

It is not yet specified what would be the case if EXIF arrives after the
image is already presented. See [issue
#4929](https://github.com/w3c/csswg-drafts/issues/4929).

###### [4.8.4.3.7]{.secno} Selecting an image source[](#selecting-an-image-source){.self-link}

To [select an image source]{#select-an-image-source .dfn} given an
[`img`{#selecting-an-image-source:the-img-element}](embedded-content.html#the-img-element)
element `el`{.variable}:

1.  [Update the source
    set](#update-the-source-set){#selecting-an-image-source:update-the-source-set}
    for `el`{.variable}.

2.  If `el`{.variable}\'s [source
    set](#source-set){#selecting-an-image-source:source-set} is empty,
    return null as the URL and undefined as the pixel density.

3.  Return the result of [selecting an
    image](#select-an-image-source-from-a-source-set){#selecting-an-image-source:select-an-image-source-from-a-source-set}
    from `el`{.variable}\'s [source
    set](#source-set){#selecting-an-image-source:source-set-2}.

To [select an image source from a source
set]{#select-an-image-source-from-a-source-set .dfn} given a [source
set](#source-set){#selecting-an-image-source:source-set-3}
`sourceSet`{.variable}:

1.  If an entry `b`{.variable} in `sourceSet`{.variable} has the same
    associated [pixel density
    descriptor](#pixel-density-descriptor){#selecting-an-image-source:pixel-density-descriptor}
    as an earlier entry `a`{.variable} in `sourceSet`{.variable}, then
    remove entry `b`{.variable}. Repeat this step until none of the
    entries in `sourceSet`{.variable} have the same associated [pixel
    density
    descriptor](#pixel-density-descriptor){#selecting-an-image-source:pixel-density-descriptor-2}
    as an earlier entry.

2.  In an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#selecting-an-image-source:implementation-defined
    x-internal="implementation-defined"} manner, choose one [image
    source](#image-source){#selecting-an-image-source:image-source} from
    `sourceSet`{.variable}. Let this be `selectedSource`{.variable}.

3.  Return `selectedSource`{.variable} and its associated pixel density.

###### [4.8.4.3.8]{.secno} Creating a source set from attributes[](#creating-a-source-set-from-attributes){.self-link}

When asked to [create a source set]{#create-a-source-set .dfn} given a
string `default source`{.variable}, a string `srcset`{.variable}, a
string `sizes`{.variable}, and an element or null `img`{.variable}:

1.  Let `source set`{.variable} be an empty [source
    set](#source-set){#creating-a-source-set-from-attributes:source-set}.

2.  If `srcset`{.variable} is not an empty string, then set
    `source set`{.variable} to the result of
    [parsing](#parse-a-srcset-attribute){#creating-a-source-set-from-attributes:parse-a-srcset-attribute}
    `srcset`{.variable}.

3.  Set `source set`{.variable}\'s [source
    size](#source-size-2){#creating-a-source-set-from-attributes:source-size-2}
    to the result of
    [parsing](#parse-a-sizes-attribute){#creating-a-source-set-from-attributes:parse-a-sizes-attribute}
    `sizes`{.variable} with `img`{.variable}.

4.  If `default source`{.variable} is not the empty string and
    `source set`{.variable} does not contain an [image
    source](#image-source){#creating-a-source-set-from-attributes:image-source}
    with a [pixel density
    descriptor](#pixel-density-descriptor){#creating-a-source-set-from-attributes:pixel-density-descriptor}
    value of 1, and no [image
    source](#image-source){#creating-a-source-set-from-attributes:image-source-2}
    with a [width
    descriptor](#width-descriptor){#creating-a-source-set-from-attributes:width-descriptor},
    append `default source`{.variable} to `source set`{.variable}.

5.  [Normalize the source
    densities](#normalise-the-source-densities){#creating-a-source-set-from-attributes:normalise-the-source-densities}
    of `source set`{.variable}.

6.  Return `source set`{.variable}.

###### [4.8.4.3.9]{.secno} Updating the source set[](#updating-the-source-set){.self-link}

When asked to [update the source set]{#update-the-source-set .dfn} for a
given
[`img`{#updating-the-source-set:the-img-element}](embedded-content.html#the-img-element)
or
[`link`{#updating-the-source-set:the-link-element}](semantics.html#the-link-element)
element `el`{.variable}, user agents must do the following:

1.  Set `el`{.variable}\'s [source
    set](#source-set){#updating-the-source-set:source-set} to an empty
    [source set](#source-set){#updating-the-source-set:source-set-2}.

2.  Let `elements`{.variable} be « `el`{.variable} ».

3.  If `el`{.variable} is an
    [`img`{#updating-the-source-set:the-img-element-2}](embedded-content.html#the-img-element)
    element whose parent node is a
    [`picture`{#updating-the-source-set:the-picture-element}](embedded-content.html#the-picture-element)
    element, then
    [replace](https://infra.spec.whatwg.org/#list-replace){#updating-the-source-set:list-replace
    x-internal="list-replace"} the contents of `elements`{.variable}
    with `el`{.variable}\'s parent node\'s child elements, retaining
    relative order.

4.  Let `img`{.variable} be `el`{.variable} if `el`{.variable} is an
    [`img`{#updating-the-source-set:the-img-element-3}](embedded-content.html#the-img-element)
    element, otherwise null.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#updating-the-source-set:list-iterate
    x-internal="list-iterate"} `child`{.variable} in
    `elements`{.variable}:

    1.  If `child`{.variable} is `el`{.variable}:

        1.  Let `default source`{.variable} be the empty string.

        2.  Let `srcset`{.variable} be the empty string.

        3.  Let `sizes`{.variable} be the empty string.

        4.  If `el`{.variable} is an
            [`img`{#updating-the-source-set:the-img-element-4}](embedded-content.html#the-img-element)
            element that has a
            [`srcset`{#updating-the-source-set:attr-img-srcset}](embedded-content.html#attr-img-srcset)
            attribute, then set `srcset`{.variable} to that attribute\'s
            value.

        5.  Otherwise, if `el`{.variable} is a
            [`link`{#updating-the-source-set:the-link-element-2}](semantics.html#the-link-element)
            element that has an
            [`imagesrcset`{#updating-the-source-set:attr-link-imagesrcset}](semantics.html#attr-link-imagesrcset)
            attribute, then set `srcset`{.variable} to that attribute\'s
            value.

        6.  If `el`{.variable} is an
            [`img`{#updating-the-source-set:the-img-element-5}](embedded-content.html#the-img-element)
            element that has a
            [`sizes`{#updating-the-source-set:attr-img-sizes}](embedded-content.html#attr-img-sizes)
            attribute, then set `sizes`{.variable} to that attribute\'s
            value.

        7.  Otherwise, if `el`{.variable} is a
            [`link`{#updating-the-source-set:the-link-element-3}](semantics.html#the-link-element)
            element that has an
            [`imagesizes`{#updating-the-source-set:attr-link-imagesizes}](semantics.html#attr-link-imagesizes)
            attribute, then set `sizes`{.variable} to that attribute\'s
            value.

        8.  If `el`{.variable} is an
            [`img`{#updating-the-source-set:the-img-element-6}](embedded-content.html#the-img-element)
            element that has a
            [`src`{#updating-the-source-set:attr-img-src}](embedded-content.html#attr-img-src)
            attribute, then set `default source`{.variable} to that
            attribute\'s value.

        9.  Otherwise, if `el`{.variable} is a
            [`link`{#updating-the-source-set:the-link-element-4}](semantics.html#the-link-element)
            element that has an
            [`href`{#updating-the-source-set:attr-link-href}](semantics.html#attr-link-href)
            attribute, then set `default source`{.variable} to that
            attribute\'s value.

        10. Set `el`{.variable}\'s [source
            set](#source-set){#updating-the-source-set:source-set-3} to
            the result of [creating a source
            set](#create-a-source-set){#updating-the-source-set:create-a-source-set}
            given `default source`{.variable}, `srcset`{.variable},
            `sizes`{.variable}, and `img`{.variable}.

        11. Return.

            If `el`{.variable} is a
            [`link`{#updating-the-source-set:the-link-element-5}](semantics.html#the-link-element)
            element, then `elements`{.variable} contains only
            `el`{.variable}, so this step will be reached immediately
            and the rest of the algorithm will not run.

    2.  If `child`{.variable} is not a
        [`source`{#updating-the-source-set:the-source-element}](embedded-content.html#the-source-element)
        element, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-source-set:continue
        x-internal="continue"}.

    3.  If `child`{.variable} does not have a
        [`srcset`{#updating-the-source-set:attr-source-srcset}](embedded-content.html#attr-source-srcset)
        attribute,
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-source-set:continue-2
        x-internal="continue"} to the next child.

    4.  [Parse `child`{.variable}\'s srcset
        attribute](#parse-a-srcset-attribute){#updating-the-source-set:parse-a-srcset-attribute}
        and let the returned [source
        set](#source-set){#updating-the-source-set:source-set-4} be
        `source set`{.variable}.

    5.  If `source set`{.variable} has zero [image
        sources](#image-source){#updating-the-source-set:image-source},
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-source-set:continue-3
        x-internal="continue"} to the next child.

    6.  If `child`{.variable} has a
        [`media`{#updating-the-source-set:attr-source-media}](embedded-content.html#attr-source-media)
        attribute, and its value does not [match the
        environment](common-microsyntaxes.html#matches-the-environment){#updating-the-source-set:matches-the-environment},
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-source-set:continue-4
        x-internal="continue"} to the next child.

    7.  [Parse `child`{.variable}\'s sizes
        attribute](#parse-a-sizes-attribute){#updating-the-source-set:parse-a-sizes-attribute}
        with `img`{.variable}, and let `source set`{.variable}\'s
        [source
        size](#source-size-2){#updating-the-source-set:source-size-2} be
        the returned value.

    8.  If `child`{.variable} has a
        [`type`{#updating-the-source-set:attr-source-type}](embedded-content.html#attr-source-type)
        attribute, and its value is an unknown or unsupported [MIME
        type](https://mimesniff.spec.whatwg.org/#mime-type){#updating-the-source-set:mime-type
        x-internal="mime-type"},
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#updating-the-source-set:continue-5
        x-internal="continue"} to the next child.

    9.  If `child`{.variable} has
        [`width`{#updating-the-source-set:attr-dim-width}](embedded-content-other.html#attr-dim-width)
        or
        [`height`{#updating-the-source-set:attr-dim-height}](embedded-content-other.html#attr-dim-height)
        attributes, set `el`{.variable}\'s [dimension attribute
        source](embedded-content.html#concept-img-dimension-attribute-source){#updating-the-source-set:concept-img-dimension-attribute-source}
        to `child`{.variable}. Otherwise, set `el`{.variable}\'s
        [dimension attribute
        source](embedded-content.html#concept-img-dimension-attribute-source){#updating-the-source-set:concept-img-dimension-attribute-source-2}
        to `el`{.variable}.

    10. [Normalize the source
        densities](#normalise-the-source-densities){#updating-the-source-set:normalise-the-source-densities}
        of `source set`{.variable}.

    11. Set `el`{.variable}\'s [source
        set](#source-set){#updating-the-source-set:source-set-5} to
        `source set`{.variable}.

    12. Return.

Each
[`img`{#updating-the-source-set:the-img-element-7}](embedded-content.html#the-img-element)
element independently considers its previous sibling
[`source`{#updating-the-source-set:the-source-element-2}](embedded-content.html#the-source-element)
elements plus the
[`img`{#updating-the-source-set:the-img-element-8}](embedded-content.html#the-img-element)
element itself for selecting an [image
source](#image-source){#updating-the-source-set:image-source-2},
ignoring any other (invalid) elements, including other
[`img`{#updating-the-source-set:the-img-element-9}](embedded-content.html#the-img-element)
elements in the same
[`picture`{#updating-the-source-set:the-picture-element-2}](embedded-content.html#the-picture-element)
element, or
[`source`{#updating-the-source-set:the-source-element-3}](embedded-content.html#the-source-element)
elements that are following siblings of the relevant
[`img`{#updating-the-source-set:the-img-element-10}](embedded-content.html#the-img-element)
element.

###### [4.8.4.3.10]{.secno} Parsing a srcset attribute[](#parsing-a-srcset-attribute){.self-link}

When asked to [parse a srcset attribute]{#parse-a-srcset-attribute .dfn}
from an element, parse the value of the element\'s [srcset
attribute](#srcset-attribute){#parsing-a-srcset-attribute:srcset-attribute}
as follows:

1.  Let `input`{.variable} be the value passed to this algorithm.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Let `candidates`{.variable} be an initially empty [source
    set](#source-set){#parsing-a-srcset-attribute:source-set}.

4.  *Splitting loop*: [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#parsing-a-srcset-attribute:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parsing-a-srcset-attribute:space-characters
    x-internal="space-characters"} or U+002C COMMA characters from
    `input`{.variable} given `position`{.variable}. If any U+002C COMMA
    characters were collected, that is a [parse
    error](#concept-microsyntax-parse-error){#parsing-a-srcset-attribute:concept-microsyntax-parse-error}.

5.  If `position`{.variable} is past the end of `input`{.variable},
    return `candidates`{.variable}.

6.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#parsing-a-srcset-attribute:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are not [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parsing-a-srcset-attribute:space-characters-2
    x-internal="space-characters"} from `input`{.variable} given
    `position`{.variable}, and let that be `url`{.variable}.

7.  Let `descriptors`{.variable} be a new empty list.

8.  If `url`{.variable} ends with U+002C (,), then:

    1.  Remove all trailing U+002C COMMA characters from
        `url`{.variable}. If this removed more than one character, that
        is a [parse
        error](#concept-microsyntax-parse-error){#parsing-a-srcset-attribute:concept-microsyntax-parse-error-2}.

    Otherwise:

    1.  *Descriptor tokenizer*: [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#parsing-a-srcset-attribute:skip-ascii-whitespace
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

    2.  Let `current descriptor`{.variable} be the empty string.

    3.  Let `state`{.variable} be *in descriptor*.

    4.  Let `c`{.variable} be the character at `position`{.variable}. Do
        the following depending on the value of `state`{.variable}. For
        the purpose of this step, \"EOF\" is a special character
        representing that `position`{.variable} is past the end of
        `input`{.variable}.

        *In descriptor*

        :   Do the following, depending on the value of `c`{.variable}:

            [ASCII whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parsing-a-srcset-attribute:space-characters-3 x-internal="space-characters"}
            :   If `current descriptor`{.variable} is not empty, append
                `current descriptor`{.variable} to
                `descriptors`{.variable} and let
                `current descriptor`{.variable} be the empty string. Set
                `state`{.variable} to *after descriptor*.

            U+002C COMMA (,)
            :   Advance `position`{.variable} to the next character in
                `input`{.variable}. If `current descriptor`{.variable}
                is not empty, append `current descriptor`{.variable} to
                `descriptors`{.variable}. Jump to the step labeled
                *descriptor parser*.

            U+0028 LEFT PARENTHESIS (()
            :   Append `c`{.variable} to
                `current descriptor`{.variable}. Set `state`{.variable}
                to *in parens*.

            EOF
            :   If `current descriptor`{.variable} is not empty, append
                `current descriptor`{.variable} to
                `descriptors`{.variable}. Jump to the step labeled
                *descriptor parser*.

            Anything else

            :   Append `c`{.variable} to
                `current descriptor`{.variable}.

        *In parens*

        :   Do the following, depending on the value of `c`{.variable}:

            U+0029 RIGHT PARENTHESIS ())
            :   Append `c`{.variable} to
                `current descriptor`{.variable}. Set `state`{.variable}
                to *in descriptor*.

            EOF
            :   Append `current descriptor`{.variable} to
                `descriptors`{.variable}. Jump to the step labeled
                *descriptor parser*.

            Anything else

            :   Append `c`{.variable} to
                `current descriptor`{.variable}.

        *After descriptor*

        :   Do the following, depending on the value of `c`{.variable}:

            [ASCII whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parsing-a-srcset-attribute:space-characters-4 x-internal="space-characters"}
            :   Stay in this state.

            EOF
            :   Jump to the step labeled *descriptor parser*.

            Anything else

            :   Set `state`{.variable} to *in descriptor*. Set
                `position`{.variable} to the *previous* character in
                `input`{.variable}.

        Advance `position`{.variable} to the next character in
        `input`{.variable}. Repeat this step.

        In order to be compatible with future additions, this algorithm
        supports multiple descriptors and descriptors with parens.

9.  *Descriptor parser*: Let `error`{.variable} be *no*.

10. Let `width`{.variable} be *absent*.

11. Let `density`{.variable} be *absent*.

12. Let `future-compat-h`{.variable} be *absent*.

13. For each descriptor in `descriptors`{.variable}, run the appropriate
    set of steps from the following list:

    If the descriptor consists of a [valid non-negative integer](common-microsyntaxes.html#valid-non-negative-integer){#parsing-a-srcset-attribute:valid-non-negative-integer} followed by a U+0077 LATIN SMALL LETTER W character

    :   1.  If the user agent does not support the
            [`sizes`{#parsing-a-srcset-attribute:attr-img-sizes}](embedded-content.html#attr-img-sizes)
            attribute, let `error`{.variable} be *yes*.

            A conforming user agent will support the
            [`sizes`{#parsing-a-srcset-attribute:attr-img-sizes-2}](embedded-content.html#attr-img-sizes)
            attribute. However, user agents typically implement and ship
            features in an incremental manner in practice.

        2.  If `width`{.variable} and `density`{.variable} are not both
            *absent*, then let `error`{.variable} be *yes*.

        3.  Apply the [rules for parsing non-negative
            integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#parsing-a-srcset-attribute:rules-for-parsing-non-negative-integers}
            to the descriptor. If the result is 0, let
            `error`{.variable} be *yes*. Otherwise, let
            `width`{.variable} be the result.

    If the descriptor consists of a [valid floating-point number](common-microsyntaxes.html#valid-floating-point-number){#parsing-a-srcset-attribute:valid-floating-point-number} followed by a U+0078 LATIN SMALL LETTER X character

    :   1.  If `width`{.variable}, `density`{.variable} and
            `future-compat-h`{.variable} are not all *absent*, then let
            `error`{.variable} be *yes*.

        2.  Apply the [rules for parsing floating-point number
            values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#parsing-a-srcset-attribute:rules-for-parsing-floating-point-number-values}
            to the descriptor. If the result is less than 0, let
            `error`{.variable} be *yes*. Otherwise, let
            `density`{.variable} be the result.

            If `density`{.variable} is 0, the [natural
            dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#parsing-a-srcset-attribute:natural-dimensions
            x-internal="natural-dimensions"} will be infinite. User
            agents are [expected to have
            limits](https://infra.spec.whatwg.org/#algorithm-limits) in
            how big images can be rendered.

    If the descriptor consists of a [valid non-negative integer](common-microsyntaxes.html#valid-non-negative-integer){#parsing-a-srcset-attribute:valid-non-negative-integer-2} followed by a U+0068 LATIN SMALL LETTER H character

    :   This is a [parse
        error](#concept-microsyntax-parse-error){#parsing-a-srcset-attribute:concept-microsyntax-parse-error-3}.

        1.  If `future-compat-h`{.variable} and `density`{.variable} are
            not both *absent*, then let `error`{.variable} be *yes*.

        2.  Apply the [rules for parsing non-negative
            integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#parsing-a-srcset-attribute:rules-for-parsing-non-negative-integers-2}
            to the descriptor. If the result is 0, let
            `error`{.variable} be *yes*. Otherwise, let
            `future-compat-h`{.variable} be the result.

    Anything else

    :   Let `error`{.variable} be *yes*.

14. If `future-compat-h`{.variable} is not *absent* and
    `width`{.variable} is *absent*, let `error`{.variable} be *yes*.

15. If `error`{.variable} is still *no*, then append a new [image
    source](#image-source){#parsing-a-srcset-attribute:image-source} to
    `candidates`{.variable} whose URL is `url`{.variable}, associated
    with a width `width`{.variable} if not *absent* and a pixel density
    `density`{.variable} if not *absent*. Otherwise, there is a [parse
    error](#concept-microsyntax-parse-error){#parsing-a-srcset-attribute:concept-microsyntax-parse-error-4}.

16. Return to the step labeled *splitting loop*.

###### [4.8.4.3.11]{.secno} Parsing a sizes attribute[](#parsing-a-sizes-attribute){.self-link}

When asked to [parse a sizes attribute]{#parse-a-sizes-attribute .dfn}
from an element `element`{.variable}, with an
[`img`{#parsing-a-sizes-attribute:the-img-element}](embedded-content.html#the-img-element)
element or null `img`{.variable}:

1.  Let `unparsed sizes list`{.variable} be the result of [parsing a
    comma-separated list of component
    values](https://drafts.csswg.org/css-syntax/#parse-a-comma-separated-list-of-component-values){#parsing-a-sizes-attribute:parse-a-comma-separated-list-of-component-values
    x-internal="parse-a-comma-separated-list-of-component-values"} from
    the value of `element`{.variable}\'s [sizes
    attribute](#sizes-attribute){#parsing-a-sizes-attribute:sizes-attribute}
    (or the empty string, if the attribute is absent).
    [\[CSSSYNTAX\]](references.html#refsCSSSYNTAX)

2.  Let `size`{.variable} be null.

3.  For each `unparsed size`{.variable} in
    `unparsed sizes list`{.variable}:

    1.  Remove all consecutive
        [\<whitespace-token\>](https://drafts.csswg.org/css-syntax/#typedef-whitespace-token){#parsing-a-sizes-attribute:whitespace-token
        x-internal="whitespace-token"}s from the end of
        `unparsed size`{.variable}. If `unparsed size`{.variable} is now
        empty, then that is a [parse
        error](#concept-microsyntax-parse-error){#parsing-a-sizes-attribute:concept-microsyntax-parse-error};
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#parsing-a-sizes-attribute:continue
        x-internal="continue"}.

    2.  If the last [component
        value](https://drafts.csswg.org/css-syntax/#component-value){#parsing-a-sizes-attribute:component-value
        x-internal="component-value"} in `unparsed size`{.variable} is a
        valid non-negative
        [\<source-size-value\>](#source-size-value){#parsing-a-sizes-attribute:source-size-value},
        then set `size`{.variable} to its value and remove the
        [component
        value](https://drafts.csswg.org/css-syntax/#component-value){#parsing-a-sizes-attribute:component-value-2
        x-internal="component-value"} from `unparsed size`{.variable}.
        Any CSS function other than the [math
        functions](https://drafts.csswg.org/css-values/#math-function){#parsing-a-sizes-attribute:math-functions
        x-internal="math-functions"} is invalid. Otherwise, there is a
        [parse
        error](#concept-microsyntax-parse-error){#parsing-a-sizes-attribute:concept-microsyntax-parse-error-2};
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#parsing-a-sizes-attribute:continue-2
        x-internal="continue"}.

    3.  If `size`{.variable} is
        [`auto`{#parsing-a-sizes-attribute:valdef-sizes-auto}](#valdef-sizes-auto),
        and `img`{.variable} is not null, and `img`{.variable} is [being
        rendered](rendering.html#being-rendered){#parsing-a-sizes-attribute:being-rendered},
        and `img`{.variable} [allows
        auto-sizes](embedded-content.html#allows-auto-sizes){#parsing-a-sizes-attribute:allows-auto-sizes},
        then set `size`{.variable} to the [concrete object
        size](https://drafts.csswg.org/css-images/#concrete-object-size){#parsing-a-sizes-attribute:concrete-object-size
        x-internal="concrete-object-size"} width of `img`{.variable}, in
        [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#parsing-a-sizes-attribute:'px'
        x-internal="'px'"}.

        If `size`{.variable} is still
        [`auto`{#parsing-a-sizes-attribute:valdef-sizes-auto-2}](#valdef-sizes-auto),
        then it will be ignored.

    4.  Remove all consecutive
        [\<whitespace-token\>](https://drafts.csswg.org/css-syntax/#typedef-whitespace-token){#parsing-a-sizes-attribute:whitespace-token-2
        x-internal="whitespace-token"}s from the end of
        `unparsed size`{.variable}. If `unparsed size`{.variable} is now
        empty:

        1.  If this was not the last item in
            `unparsed sizes list`{.variable}, that is a [parse
            error](#concept-microsyntax-parse-error){#parsing-a-sizes-attribute:concept-microsyntax-parse-error-3}.

        2.  If `size`{.variable} is not
            [`auto`{#parsing-a-sizes-attribute:valdef-sizes-auto-3}](#valdef-sizes-auto),
            then return `size`{.variable}. Otherwise, continue.

    5.  Parse the remaining [component
        values](https://drafts.csswg.org/css-syntax/#component-value){#parsing-a-sizes-attribute:component-value-3
        x-internal="component-value"} in `unparsed size`{.variable} as a
        [\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition){#parsing-a-sizes-attribute:media-condition
        x-internal="media-condition"}. If it does not parse correctly,
        or it does parse correctly but the
        [\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition){#parsing-a-sizes-attribute:media-condition-2
        x-internal="media-condition"} evaluates to false,
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#parsing-a-sizes-attribute:continue-3
        x-internal="continue"}. [\[MQ\]](references.html#refsMQ)

    6.  If `size`{.variable} is not
        [`auto`{#parsing-a-sizes-attribute:valdef-sizes-auto-4}](#valdef-sizes-auto),
        then return `size`{.variable}. Otherwise, continue.

4.  Return `100vw`.

It is invalid to use a bare
[\<source-size-value\>](#source-size-value){#parsing-a-sizes-attribute:source-size-value-2}
that is a
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#parsing-a-sizes-attribute:length-2
x-internal="length-2"} (without an accompanying
[\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition){#parsing-a-sizes-attribute:media-condition-3
x-internal="media-condition"}) as an entry in the
[\<source-size-list\>](#source-size-list){#parsing-a-sizes-attribute:source-size-list}
that is not the last entry. However, the parsing algorithm allows it at
any point in the
[\<source-size-list\>](#source-size-list){#parsing-a-sizes-attribute:source-size-list-2},
and will accept it immediately as the size if the preceding entries in
the list weren\'t used. This is to enable future extensions, and protect
against simple author errors such as a final trailing comma. A bare
[`auto`{#parsing-a-sizes-attribute:valdef-sizes-auto-5}](#valdef-sizes-auto)
keyword is allowed to have other entries following it to provide a
fallback for legacy user agents.

###### [4.8.4.3.12]{.secno} Normalizing the source densities[](#normalizing-the-source-densities){.self-link}

An [image
source](#image-source){#normalizing-the-source-densities:image-source}
can have a [pixel density
descriptor](#pixel-density-descriptor){#normalizing-the-source-densities:pixel-density-descriptor},
a [width
descriptor](#width-descriptor){#normalizing-the-source-densities:width-descriptor},
or no descriptor at all accompanying its URL. Normalizing a [source
set](#source-set){#normalizing-the-source-densities:source-set} gives
every [image
source](#image-source){#normalizing-the-source-densities:image-source-2}
a [pixel density
descriptor](#pixel-density-descriptor){#normalizing-the-source-densities:pixel-density-descriptor-2}.

When asked to [normalize the source
densities]{#normalise-the-source-densities .dfn} of a [source
set](#source-set){#normalizing-the-source-densities:source-set-2}
`source set`{.variable}, the user agent must do the following:

1.  Let `source size`{.variable} be `source set`{.variable}\'s [source
    size](#source-size-2){#normalizing-the-source-densities:source-size-2}.

2.  For each [image
    source](#image-source){#normalizing-the-source-densities:image-source-3}
    in `source set`{.variable}:

    1.  If the [image
        source](#image-source){#normalizing-the-source-densities:image-source-4}
        has a [pixel density
        descriptor](#pixel-density-descriptor){#normalizing-the-source-densities:pixel-density-descriptor-3},
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#normalizing-the-source-densities:continue
        x-internal="continue"} to the next [image
        source](#image-source){#normalizing-the-source-densities:image-source-5}.

    2.  Otherwise, if the [image
        source](#image-source){#normalizing-the-source-densities:image-source-6}
        has a [width
        descriptor](#width-descriptor){#normalizing-the-source-densities:width-descriptor-2},
        replace the [width
        descriptor](#width-descriptor){#normalizing-the-source-densities:width-descriptor-3}
        with a [pixel density
        descriptor](#pixel-density-descriptor){#normalizing-the-source-densities:pixel-density-descriptor-4}
        with a
        [value](#pixel-density-descriptor-value){#normalizing-the-source-densities:pixel-density-descriptor-value}
        of the [width descriptor
        value](#width-descriptor-value){#normalizing-the-source-densities:width-descriptor-value}
        divided by `source size`{.variable} and a unit of `x`.

        If the [source
        size](#source-size-2){#normalizing-the-source-densities:source-size-2-2}
        is 0, then the density would be infinity, which results in the
        [natural
        dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#normalizing-the-source-densities:natural-dimensions
        x-internal="natural-dimensions"} being 0 by 0.

    3.  Otherwise, give the [image
        source](#image-source){#normalizing-the-source-densities:image-source-7}
        a [pixel density
        descriptor](#pixel-density-descriptor){#normalizing-the-source-densities:pixel-density-descriptor-5}
        of `1x`.

###### [4.8.4.3.13]{.secno} Reacting to environment changes[](#reacting-to-environment-changes){.self-link}

The user agent may at any time run the following algorithm to update an
[`img`{#reacting-to-environment-changes:the-img-element}](embedded-content.html#the-img-element)
element\'s image in order to [react to changes in the
environment]{#img-environment-changes .dfn}. (User agents are *not
required* to ever run this algorithm; for example, if the user is not
looking at the page any more, the user agent might want to wait until
the user has returned to the page before determining which image to use,
in case the environment changes again in the meantime.)

User agents are encouraged to run this algorithm in particular when the
user changes the
[viewport](https://drafts.csswg.org/css2/#viewport){#reacting-to-environment-changes:viewport
x-internal="viewport"}\'s size (e.g. by resizing the window or changing
the page zoom), and when an
[`img`{#reacting-to-environment-changes:the-img-element-2}](embedded-content.html#the-img-element)
element is [inserted into a
document](infrastructure.html#insert-an-element-into-a-document){#reacting-to-environment-changes:insert-an-element-into-a-document},
so that the [density-corrected natural width and
height](#density-corrected-intrinsic-width-and-height){#reacting-to-environment-changes:density-corrected-intrinsic-width-and-height}
match the new
[viewport](https://drafts.csswg.org/css2/#viewport){#reacting-to-environment-changes:viewport-2
x-internal="viewport"}, and so that the correct image is chosen when
[art
direction](#art-direction){#reacting-to-environment-changes:art-direction}
is involved.

1.  [Await a stable
    state](webappapis.html#await-a-stable-state){#reacting-to-environment-changes:await-a-stable-state}.
    The [synchronous
    section](webappapis.html#synchronous-section){#reacting-to-environment-changes:synchronous-section}
    consists of all the remaining steps of this algorithm until the
    algorithm says the [synchronous
    section](webappapis.html#synchronous-section){#reacting-to-environment-changes:synchronous-section-2}
    has ended. (Steps in [synchronous
    sections](webappapis.html#synchronous-section){#reacting-to-environment-changes:synchronous-section-3}
    are marked with ⌛.)

2.  ⌛ If the
    [`img`{#reacting-to-environment-changes:the-img-element-3}](embedded-content.html#the-img-element)
    element does not [use `srcset` or
    `picture`](#use-srcset-or-picture){#reacting-to-environment-changes:use-srcset-or-picture},
    its [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#reacting-to-environment-changes:node-document
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#reacting-to-environment-changes:fully-active},
    has image data whose resource type is
    [`multipart/x-mixed-replace`{#reacting-to-environment-changes:multipart/x-mixed-replace}](iana.html#multipart/x-mixed-replace),
    or the [pending
    request](#pending-request){#reacting-to-environment-changes:pending-request}
    is not null, then return.

3.  ⌛ Let `selected source`{.variable} and
    `selected pixel density`{.variable} be the URL and pixel density
    that results from [selecting an image
    source](#select-an-image-source){#reacting-to-environment-changes:select-an-image-source},
    respectively.

4.  ⌛ If `selected source`{.variable} is null, then return.

5.  ⌛ If `selected source`{.variable} and
    `selected pixel density`{.variable} are the same as the element\'s
    [last selected
    source](#last-selected-source){#reacting-to-environment-changes:last-selected-source}
    and [current pixel
    density](#current-pixel-density){#reacting-to-environment-changes:current-pixel-density},
    then return.

6.  ⌛ Let `urlString`{.variable} be the result of
    [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#reacting-to-environment-changes:encoding-parsing-and-serializing-a-url}
    given `selected source`{.variable}, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#reacting-to-environment-changes:node-document-2
    x-internal="node-document"}.

7.  ⌛ If `urlString`{.variable} is failure, then return.

8.  ⌛ Let `corsAttributeState`{.variable} be the state of the
    element\'s
    [`crossorigin`{#reacting-to-environment-changes:attr-img-crossorigin}](embedded-content.html#attr-img-crossorigin)
    content attribute.

9.  ⌛ Let `origin`{.variable} be the
    [`img`{#reacting-to-environment-changes:the-img-element-4}](embedded-content.html#the-img-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#reacting-to-environment-changes:node-document-3
    x-internal="node-document"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#reacting-to-environment-changes:concept-document-origin
    x-internal="concept-document-origin"}.

10. ⌛ Let `client`{.variable} be the
    [`img`{#reacting-to-environment-changes:the-img-element-5}](embedded-content.html#the-img-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#reacting-to-environment-changes:node-document-4
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#reacting-to-environment-changes:relevant-settings-object}.

11. ⌛ Let `key`{.variable} be a tuple consisting of
    `urlString`{.variable}, `corsAttributeState`{.variable}, and, if
    `corsAttributeState`{.variable} is not [No
    CORS](urls-and-fetching.html#attr-crossorigin-none){#reacting-to-environment-changes:attr-crossorigin-none},
    `origin`{.variable}.

12. ⌛ Let `image request`{.variable} be a new [image
    request](#image-request){#reacting-to-environment-changes:image-request}
    whose [current
    URL](#img-req-url){#reacting-to-environment-changes:img-req-url} is
    `urlString`{.variable}.

13. ⌛ Let the element\'s [pending
    request](#pending-request){#reacting-to-environment-changes:pending-request-2}
    be `image request`{.variable}.

14. End the [synchronous
    section](webappapis.html#synchronous-section){#reacting-to-environment-changes:synchronous-section-4},
    continuing the remaining steps [in
    parallel](infrastructure.html#in-parallel){#reacting-to-environment-changes:in-parallel}.

15. If the [list of available
    images](#list-of-available-images){#reacting-to-environment-changes:list-of-available-images}
    contains an entry for `key`{.variable}, then set
    `image request`{.variable}\'s [image
    data](#img-req-data){#reacting-to-environment-changes:img-req-data}
    to that of the entry. Continue to the next step.

    Otherwise:

    1.  Let `request`{.variable} be the result of [creating a
        potential-CORS
        request](urls-and-fetching.html#create-a-potential-cors-request){#reacting-to-environment-changes:create-a-potential-cors-request}
        given `urlString`{.variable}, \"`image`\", and
        `corsAttributeState`{.variable}.

    2.  Set `request`{.variable}\'s
        [client](https://fetch.spec.whatwg.org/#concept-request-client){#reacting-to-environment-changes:concept-request-client
        x-internal="concept-request-client"} to `client`{.variable},
        [initiator](https://fetch.spec.whatwg.org/#concept-request-initiator){#reacting-to-environment-changes:concept-request-initiator
        x-internal="concept-request-initiator"} to \"`imageset`\", and
        set `request`{.variable}\'s [synchronous
        flag](https://fetch.spec.whatwg.org/#synchronous-flag){#reacting-to-environment-changes:synchronous-flag
        x-internal="synchronous-flag"}.

    3.  Set `request`{.variable}\'s [referrer
        policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#reacting-to-environment-changes:concept-request-referrer-policy
        x-internal="concept-request-referrer-policy"} to the current
        state of the element\'s
        [`referrerpolicy`{#reacting-to-environment-changes:attr-img-referrerpolicy}](embedded-content.html#attr-img-referrerpolicy)
        attribute.

    4.  Set `request`{.variable}\'s
        [priority](https://fetch.spec.whatwg.org/#request-priority){#reacting-to-environment-changes:concept-request-priority
        x-internal="concept-request-priority"} to the current state of
        the element\'s
        [`fetchpriority`{#reacting-to-environment-changes:attr-img-fetchpriority}](embedded-content.html#attr-img-fetchpriority)
        attribute.

    5.  Let `response`{.variable} be the result of
        [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#reacting-to-environment-changes:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}.

    6.  If `response`{.variable}\'s [unsafe
        response](urls-and-fetching.html#unsafe-response){#reacting-to-environment-changes:unsafe-response}
        is a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#reacting-to-environment-changes:network-error
        x-internal="network-error"} or if the image format is
        unsupported (as determined by applying the [image sniffing
        rules](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#reacting-to-environment-changes:content-type-sniffing:-image
        x-internal="content-type-sniffing:-image"}, again as mentioned
        earlier), or if the user agent is able to determine that
        `image request`{.variable}\'s image is corrupted in some fatal
        way such that the image dimensions cannot be obtained, or if the
        resource type is
        [`multipart/x-mixed-replace`{#reacting-to-environment-changes:multipart/x-mixed-replace-2}](iana.html#multipart/x-mixed-replace),
        then let the [pending
        request](#pending-request){#reacting-to-environment-changes:pending-request-3}
        be null and abort these steps.

    7.  Otherwise, `response`{.variable}\'s [unsafe
        response](urls-and-fetching.html#unsafe-response){#reacting-to-environment-changes:unsafe-response-2}
        is `image request`{.variable}\'s [image
        data](#img-req-data){#reacting-to-environment-changes:img-req-data-2}.
        It can be either
        [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#reacting-to-environment-changes:cors-same-origin}
        or
        [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#reacting-to-environment-changes:cors-cross-origin};
        this affects the image\'s interaction with other APIs (e.g.,
        when used on a
        [`canvas`{#reacting-to-environment-changes:the-canvas-element}](canvas.html#the-canvas-element)).

16. [Queue an element
    task](webappapis.html#queue-an-element-task){#reacting-to-environment-changes:queue-an-element-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#reacting-to-environment-changes:dom-manipulation-task-source}
    given the
    [`img`{#reacting-to-environment-changes:the-img-element-6}](embedded-content.html#the-img-element)
    element and the following steps:

    1.  If the
        [`img`{#reacting-to-environment-changes:the-img-element-7}](embedded-content.html#the-img-element)
        element has experienced [relevant
        mutations](#relevant-mutations){#reacting-to-environment-changes:relevant-mutations}
        since this algorithm started, then let the [pending
        request](#pending-request){#reacting-to-environment-changes:pending-request-4}
        be null and abort these steps.

    2.  Let the
        [`img`{#reacting-to-environment-changes:the-img-element-8}](embedded-content.html#the-img-element)
        element\'s [last selected
        source](#last-selected-source){#reacting-to-environment-changes:last-selected-source-2}
        be `selected source`{.variable} and the
        [`img`{#reacting-to-environment-changes:the-img-element-9}](embedded-content.html#the-img-element)
        element\'s [current pixel
        density](#current-pixel-density){#reacting-to-environment-changes:current-pixel-density-2}
        be `selected pixel density`{.variable}.

    3.  Set the `image request`{.variable}\'s
        [state](#img-req-state){#reacting-to-environment-changes:img-req-state}
        to [completely
        available](#img-all){#reacting-to-environment-changes:img-all}.

    4.  Add the image to the [list of available
        images](#list-of-available-images){#reacting-to-environment-changes:list-of-available-images-2}
        using the key `key`{.variable}, with the [ignore higher-layer
        caching](#ignore-higher-layer-caching){#reacting-to-environment-changes:ignore-higher-layer-caching}
        flag set.

    5.  [Upgrade the pending request to the current
        request](#upgrade-the-pending-request-to-the-current-request){#reacting-to-environment-changes:upgrade-the-pending-request-to-the-current-request}.

    6.  [Prepare `image request`{.variable} for
        presentation](#prepare-an-image-for-presentation){#reacting-to-environment-changes:prepare-an-image-for-presentation}
        given the
        [`img`{#reacting-to-environment-changes:the-img-element-10}](embedded-content.html#the-img-element)
        element.

    7.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#reacting-to-environment-changes:concept-event-fire
        x-internal="concept-event-fire"} named
        [`load`{#reacting-to-environment-changes:event-load}](indices.html#event-load)
        at the
        [`img`{#reacting-to-environment-changes:the-img-element-11}](embedded-content.html#the-img-element)
        element.

##### [4.8.4.4]{.secno} Requirements for providing text to act as an alternative for images[](#alt){.self-link} {#alt}

###### [4.8.4.4.1]{.secno} General guidelines[](#general-guidelines){.self-link}

Except where otherwise specified, the
[`alt`{#general-guidelines:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be specified and its value must not be empty; the value
must be an appropriate replacement for the image. The specific
requirements for the
[`alt`{#general-guidelines:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute depend on what the image is intended to represent, as
described in the following sections.

The most general rule to consider when writing alternative text is the
following: **the intent is that replacing every image with the text of
its
[`alt`{#general-guidelines:attr-img-alt-3}](embedded-content.html#attr-img-alt)
attribute does not change the meaning of the page**.

So, in general, alternative text can be written by considering what one
would have written had one not been able to include the image.

A corollary to this is that the
[`alt`{#general-guidelines:attr-img-alt-4}](embedded-content.html#attr-img-alt)
attribute\'s value should never contain text that could be considered
the image\'s *caption*, *title*, or *legend*. It is supposed to contain
replacement text that could be used by users *instead* of the image; it
is not meant to supplement the image. The
[`title`{#general-guidelines:attr-title}](dom.html#attr-title) attribute
can be used for supplemental information.

Another corollary is that the
[`alt`{#general-guidelines:attr-img-alt-5}](embedded-content.html#attr-img-alt)
attribute\'s value should not repeat information that is already
provided in the prose next to the image.

One way to think of alternative text is to think about how you would
read the page containing the image to someone over the phone, without
mentioning that there is an image present. Whatever you say instead of
the image is typically a good start for writing the alternative text.

###### [4.8.4.4.2]{.secno} A link or button containing nothing but the image[](#a-link-or-button-containing-nothing-but-the-image){.self-link}

When an
[`a`{#a-link-or-button-containing-nothing-but-the-image:the-a-element}](text-level-semantics.html#the-a-element)
element that creates a
[hyperlink](links.html#hyperlink){#a-link-or-button-containing-nothing-but-the-image:hyperlink},
or a
[`button`{#a-link-or-button-containing-nothing-but-the-image:the-button-element}](form-elements.html#the-button-element)
element, has no textual content but contains one or more images, the
[`alt`{#a-link-or-button-containing-nothing-but-the-image:attr-img-alt}](embedded-content.html#attr-img-alt)
attributes must contain text that together convey the purpose of the
link or button.

::: example
In this example, a user is asked to pick their preferred color from a
list of three. Each color is given by an image, but for users who have
configured their user agent not to display images, the color names are
used instead:

``` html
<h1>Pick your color</h1>
<ul>
 <li><a href="green.html"><img src="green.jpeg" alt="Green"></a></li>
 <li><a href="blue.html"><img src="blue.jpeg" alt="Blue"></a></li>
 <li><a href="red.html"><img src="red.jpeg" alt="Red"></a></li>
</ul>
```
:::

::: example
In this example, each button has a set of images to indicate the kind of
color output desired by the user. The first image is used in each case
to give the alternative text.

``` html
<button name="rgb"><img src="red" alt="RGB"><img src="green" alt=""><img src="blue" alt=""></button>
<button name="cmyk"><img src="cyan" alt="CMYK"><img src="magenta" alt=""><img src="yellow" alt=""><img src="black" alt=""></button>
```

Since each image represents one part of the text, it could also be
written like this:

``` html
<button name="rgb"><img src="red" alt="R"><img src="green" alt="G"><img src="blue" alt="B"></button>
<button name="cmyk"><img src="cyan" alt="C"><img src="magenta" alt="M"><img src="yellow" alt="Y"><img src="black" alt="K"></button>
```

However, with other alternative text, this might not work, and putting
all the alternative text into one image in each case might make more
sense:

``` html
<button name="rgb"><img src="red" alt="sRGB profile"><img src="green" alt=""><img src="blue" alt=""></button>
<button name="cmyk"><img src="cyan" alt="CMYK profile"><img src="magenta" alt=""><img src="yellow" alt=""><img src="black" alt=""></button>
```
:::

###### [4.8.4.4.3]{.secno} A phrase or paragraph with an alternative graphical representation: charts, diagrams, graphs, maps, illustrations[](#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations){.self-link} {#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations}

Sometimes something can be more clearly stated in graphical form, for
example as a flowchart, a diagram, a graph, or a simple map showing
directions. In such cases, an image can be given using the
[`img`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:the-img-element}](embedded-content.html#the-img-element)
element, but the lesser textual version must still be given, so that
users who are unable to view the image (e.g. because they have a very
slow connection, or because they are using a text-only browser, or
because they are listening to the page being read out by a hands-free
automobile voice web browser, or simply because they are blind) are
still able to understand the message being conveyed.

The text must be given in the
[`alt`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute, and must convey the same message as the image specified in
the
[`src`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:attr-img-src}](embedded-content.html#attr-img-src)
attribute.

It is important to realize that the alternative text is a *replacement*
for the image, not a description of the image.

::: example
In the following example we have [a
flowchart](/images/parsing-model-overview.svg) in image form, with text
in the
[`alt`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute rephrasing the flowchart in prose form:

``` html
<p>In the common case, the data handled by the tokenization stage
comes from the network, but it can also come from script.</p>
<p><img src="images/parsing-model-overview.svg" alt="The Network
passes data to the Input Stream Preprocessor, which passes it to the
Tokenizer, which passes it to the Tree Construction stage. From there,
data goes to both the DOM and to Script Execution. Script Execution is
linked to the DOM, and, using document.write(), passes data to the
Tokenizer."></p>
```
:::

::: example
Here\'s another example, showing a good solution and a bad solution to
the problem of including an image in a description.

First, here\'s the good solution. This sample shows how the alternative
text should just be what you would have put in the prose if the image
had never existed.

``` html
<!-- This is the correct way to do things. -->
<p>
 You are standing in an open field west of a house.
 <img src="house.jpeg" alt="The house is white, with a boarded front door.">
 There is a small mailbox here.
</p>
```

Second, here\'s the bad solution. In this incorrect way of doing things,
the alternative text is simply a description of the image, instead of a
textual replacement for the image. It\'s bad because when the image
isn\'t shown, the text doesn\'t flow as well as in the first example.

``` bad
<!-- This is the wrong way to do things. -->
<p>
 You are standing in an open field west of a house.
 <img src="house.jpeg" alt="A white house, with a boarded front door.">
 There is a small mailbox here.
</p>
```

Text such as \"Photo of white house with boarded door\" would be equally
bad alternative text (though it could be suitable for the
[`title`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:attr-title}](dom.html#attr-title)
attribute or in the
[`figcaption`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:the-figcaption-element}](grouping-content.html#the-figcaption-element)
element of a
[`figure`{#a-phrase-or-paragraph-with-an-alternative-graphical-representation:-charts,-diagrams,-graphs,-maps,-illustrations:the-figure-element}](grouping-content.html#the-figure-element)
with this image).
:::

###### [4.8.4.4.4]{.secno} A short phrase or label with an alternative graphical representation: icons, logos[](#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos){.self-link} {#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos}

A document can contain information in iconic form. The icon is intended
to help users of visual browsers to recognize features at a glance.

In some cases, the icon is supplemental to a text label conveying the
same meaning. In those cases, the
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be present but must be empty.

::: example
Here the icons are next to text that conveys the same meaning, so they
have an empty
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute:

``` html
<nav>
 <p><a href="/help/"><img src="/icons/help.png" alt=""> Help</a></p>
 <p><a href="/configure/"><img src="/icons/configuration.png" alt="">
 Configuration Tools</a></p>
</nav>
```
:::

In other cases, the icon has no text next to it describing what it
means; the icon is supposed to be self-explanatory. In those cases, an
equivalent textual label must be given in the
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt-3}](embedded-content.html#attr-img-alt)
attribute.

::: example
Here, posts on a news site are labeled with an icon indicating their
topic.

``` html
<body>
 <article>
  <header>
   <h1>Ratatouille wins <i>Best Movie of the Year</i> award</h1>
   <p><img src="movies.png" alt="Movies"></p>
  </header>
  <p>Pixar has won yet another <i>Best Movie of the Year</i> award,
  making this its 8th win in the last 12 years.</p>
 </article>
 <article>
  <header>
   <h1>Latest TWiT episode is online</h1>
   <p><img src="podcasts.png" alt="Podcasts"></p>
  </header>
  <p>The latest TWiT episode has been posted, in which we hear
  several tech news stories as well as learning much more about the
  iPhone. This week, the panelists compare how reflective their
  iPhones' Apple logos are.</p>
 </article>
</body>
```
:::

Many pages include logos, insignia, flags, or emblems, which stand for a
particular entity such as a company, organization, project, band,
software package, country, or some such.

If the logo is being used to represent the entity, e.g. as a page
heading, the
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt-4}](embedded-content.html#attr-img-alt)
attribute must contain the name of the entity being represented by the
logo. The
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt-5}](embedded-content.html#attr-img-alt)
attribute must *not* contain text like the word \"logo\", as it is not
the fact that it is a logo that is being conveyed, it\'s the entity
itself.

If the logo is being used next to the name of the entity that it
represents, then the logo is supplemental, and its
[`alt`{#a-short-phrase-or-label-with-an-alternative-graphical-representation:-icons,-logos:attr-img-alt-6}](embedded-content.html#attr-img-alt)
attribute must instead be empty.

If the logo is merely used as decorative material (as branding, or, for
example, as a side image in an article that mentions the entity to which
the logo belongs), then the entry below on purely decorative images
applies. If the logo is actually being discussed, then it is being used
as a phrase or paragraph (the description of the logo) with an
alternative graphical representation (the logo itself), and the first
entry above applies.

::: example
In the following snippets, all four of the above cases are present.
First, we see a logo used to represent a company:

``` html
<h1><img src="XYZ.gif" alt="The XYZ company"></h1>
```

Next, we see a paragraph which uses a logo right next to the company
name, and so doesn\'t have any alternative text:

``` html
<article>
 <h2>News</h2>
 <p>We have recently been looking at buying the <img src="alpha.gif"
 alt=""> ΑΒΓ company, a small Greek company
 specializing in our type of product.</p>
```

In this third snippet, we have a logo being used in an aside, as part of
the larger article discussing the acquisition:

``` html
<aside><p><img src="alpha-large.gif" alt=""></p></aside>
 <p>The ΑΒΓ company has had a good quarter, and our
 pie chart studies of their accounts suggest a much bigger blue slice
 than its green and orange slices, which is always a good sign.</p>
</article>
```

Finally, we have an opinion piece talking about a logo, and the logo is
therefore described in detail in the alternative text.

``` html
<p>Consider for a moment their logo:</p>

<p><img src="/images/logo" alt="It consists of a green circle with a
green question mark centered inside it."></p>

<p>How unoriginal can you get? I mean, oooooh, a question mark, how
<em>revolutionary</em>, how utterly <em>ground-breaking</em>, I'm
sure everyone will rush to adopt those specifications now! They could
at least have tried for some sort of, I don't know, sequence of
rounded squares with varying shades of green and bold white outlines,
at least that would look good on the cover of a blue book.</p>
```

This example shows how the alternative text should be written such that
if the image isn\'t *[available](#img-available)*, and the text is used
instead, the text flows seamlessly into the surrounding text, as if the
image had never been there in the first place.
:::

###### [4.8.4.4.5]{.secno} Text that has been rendered to a graphic for typographical effect[](#text-that-has-been-rendered-to-a-graphic-for-typographical-effect){.self-link}

Sometimes, an image just consists of text, and the purpose of the image
is not to highlight the actual typographic effects used to render the
text, but just to convey the text itself.

In such cases, the
[`alt`{#text-that-has-been-rendered-to-a-graphic-for-typographical-effect:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be present but must consist of the same text as written
in the image itself.

::: example
Consider a graphic containing the text \"Earth Day\", but with the
letters all decorated with flowers and plants. If the text is merely
being used as a heading, to spice up the page for graphical users, then
the correct alternative text is just the same text \"Earth Day\", and no
mention need be made of the decorations:

``` html
<h1><img src="earthdayheading.png" alt="Earth Day"></h1>
```
:::

::: example
An illuminated manuscript might use graphics for some of its images. The
alternative text in such a situation is just the character that the
image represents.

``` html
<p><img src="initials/o.svg" alt="O">nce upon a time and a long long time ago, late at
night, when it was dark, over the hills, through the woods, across a great ocean, in a land far
away, in a small house, on a hill, under a full moon...
```
:::

When an image is used to represent a character that cannot otherwise be
represented in Unicode, for example gaiji, itaiji, or new characters
such as novel currency symbols, the alternative text should be a more
conventional way of writing the same thing, e.g. using the phonetic
hiragana or katakana to give the character\'s pronunciation.

::: example
In this example from 1997, a new-fangled currency symbol that looks like
a curly E with two bars in the middle instead of one is represented
using an image. The alternative text gives the character\'s
pronunciation.

``` html
<p>Only <img src="euro.png" alt="euro ">5.99!
```
:::

An image should not be used if characters would serve an identical
purpose. Only when the text cannot be directly represented using text,
e.g., because of decorations or because there is no appropriate
character (as in the case of gaiji), would an image be appropriate.

If an author is tempted to use an image because their default system
font does not support a given character, then web fonts are a better
solution than images.

###### [4.8.4.4.6]{.secno} A graphical representation of some of the surrounding text[](#a-graphical-representation-of-some-of-the-surrounding-text){.self-link}

In many cases, the image is actually just supplementary, and its
presence merely reinforces the surrounding text. In these cases, the
[`alt`{#a-graphical-representation-of-some-of-the-surrounding-text:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be present but its value must be the empty string.

In general, an image falls into this category if removing the image
doesn\'t make the page any less useful, but including the image makes it
a lot easier for users of visual browsers to understand the concept.

::: example
A flowchart that repeats the previous paragraph in graphical form:

``` html
<p>The Network passes data to the Input Stream Preprocessor, which
passes it to the Tokenizer, which passes it to the Tree Construction
stage. From there, data goes to both the DOM and to Script Execution.
Script Execution is linked to the DOM, and, using document.write(),
passes data to the Tokenizer.</p>
<p><img src="images/parsing-model-overview.svg" alt=""></p>
```

In these cases, it would be wrong to include alternative text that
consists of just a caption. If a caption is to be included, then either
the
[`title`{#a-graphical-representation-of-some-of-the-surrounding-text:attr-title}](dom.html#attr-title)
attribute can be used, or the
[`figure`{#a-graphical-representation-of-some-of-the-surrounding-text:the-figure-element}](grouping-content.html#the-figure-element)
and
[`figcaption`{#a-graphical-representation-of-some-of-the-surrounding-text:the-figcaption-element}](grouping-content.html#the-figcaption-element)
elements can be used. In the latter case, the image would in fact be a
phrase or paragraph with an alternative graphical representation, and
would thus require alternative text.

``` html
<!-- Using the title="" attribute -->
<p>The Network passes data to the Input Stream Preprocessor, which
passes it to the Tokenizer, which passes it to the Tree Construction
stage. From there, data goes to both the DOM and to Script Execution.
Script Execution is linked to the DOM, and, using document.write(),
passes data to the Tokenizer.</p>
<p><img src="images/parsing-model-overview.svg" alt=""
        title="Flowchart representation of the parsing model."></p>
```

``` html
<!-- Using <figure> and <figcaption> -->
<p>The Network passes data to the Input Stream Preprocessor, which
passes it to the Tokenizer, which passes it to the Tree Construction
stage. From there, data goes to both the DOM and to Script Execution.
Script Execution is linked to the DOM, and, using document.write(),
passes data to the Tokenizer.</p>
<figure>
 <img src="images/parsing-model-overview.svg" alt="The Network leads to
 the Input Stream Preprocessor, which leads to the Tokenizer, which
 leads to the Tree Construction stage. The Tree Construction stage
 leads to two items. The first is Script Execution, which leads via
 document.write() back to the Tokenizer. The second item from which
 Tree Construction leads is the DOM. The DOM is related to the Script
 Execution.">
 <figcaption>Flowchart representation of the parsing model.</figcaption>
</figure>
```

``` bad
<!-- This is WRONG. Do not do this. Instead, do what the above examples do. -->
<p>The Network passes data to the Input Stream Preprocessor, which
passes it to the Tokenizer, which passes it to the Tree Construction
stage. From there, data goes to both the DOM and to Script Execution.
Script Execution is linked to the DOM, and, using document.write(),
passes data to the Tokenizer.</p>
<p><img src="images/parsing-model-overview.svg"
        alt="Flowchart representation of the parsing model."></p>
<!-- Never put the image's caption in the alt="" attribute! -->
```
:::

::: example
A graph that repeats the previous paragraph in graphical form:

``` html
<p>According to a study covering several billion pages,
about 62% of documents on the web in 2007 triggered the Quirks
rendering mode of web browsers, about 30% triggered the Almost
Standards mode, and about 9% triggered the Standards mode.</p>
<p><img src="rendering-mode-pie-chart.png" alt=""></p>
```
:::

###### [4.8.4.4.7]{.secno} Ancillary images[](#ancillary-images){.self-link}

Sometimes, an image is not critical to the content, but is nonetheless
neither purely decorative nor entirely redundant with the text. In these
cases, the
[`alt`{#ancillary-images:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be present, and its value should either be the empty
string, or a textual representation of the information that the image
conveys. If the image has a caption giving the image\'s title, then the
[`alt`{#ancillary-images:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute\'s value must not be empty (as that would be quite confusing
for non-visual readers).

::: example
Consider a news article about a political figure, in which the
individual\'s face was shown in an image. The image is not purely
decorative, as it is relevant to the story. The image is not entirely
redundant with the story either, as it shows what the politician looks
like. Whether any alternative text need be provided is an authoring
decision, decided by whether the image influences the interpretation of
the prose.

In this first variant, the image is shown without context, and no
alternative text is provided:

``` html
<p><img src="president.jpeg" alt=""> Ahead of today's referendum,
the President wrote an open letter to all registered voters. In it, she admitted that the country was
divided.</p>
```

If the picture is just a face, there might be no value in describing it.
It\'s of no interest to the reader whether the individual has red hair
or blond hair, whether the individual has white skin or black skin,
whether the individual has one eye or two eyes.

However, if the picture is more dynamic, for instance showing the
politician as angry, or particularly happy, or devastated, some
alternative text would be useful in setting the tone of the article, a
tone that might otherwise be missed:

``` html
<p><img src="president.jpeg" alt="The President is sad.">
Ahead of today's referendum, the President wrote an open letter to all
registered voters. In it, she admitted that the country was divided.
</p>
```

``` html
<p><img src="president.jpeg" alt="The President is happy!">
Ahead of today's referendum, the President wrote an open letter to all
registered voters. In it, she admitted that the country was divided.
</p>
```

Whether the individual was \"sad\" or \"happy\" makes a difference to
how the rest of the paragraph is to be interpreted: is she likely saying
that she is unhappy with the country being divided, or is she saying
that the prospect of a divided country is good for her political career?
The interpretation varies based on the image.
:::

::: example
If the image has a caption, then including alternative text avoids
leaving the non-visual user confused as to what the caption refers to.

``` html
<p>Ahead of today's referendum, the President wrote an open letter to
all registered voters. In it, she admitted that the country was divided.</p>
<figure>
 <img src="president.jpeg"
      alt="A high forehead, cheerful disposition, and dark hair round out the President's face.">
 <figcaption> The President of Ruritania. Photo © 2014 PolitiPhoto. </figcaption>
</figure>
```
:::

###### [4.8.4.4.8]{.secno} A purely decorative image that doesn\'t add any information[](#a-purely-decorative-image-that-doesn't-add-any-information){.self-link} {#a-purely-decorative-image-that-doesn't-add-any-information}

If an image is decorative but isn\'t especially page-specific --- for
example an image that forms part of a site-wide design scheme --- the
image should be specified in the site\'s CSS, not in the markup of the
document.

However, a decorative image that isn\'t discussed by the surrounding
text but still has some relevance can be included in a page using the
[`img`{#a-purely-decorative-image-that-doesn't-add-any-information:the-img-element}](embedded-content.html#the-img-element)
element. Such images are decorative, but still form part of the content.
In these cases, the
[`alt`{#a-purely-decorative-image-that-doesn't-add-any-information:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be present but its value must be the empty string.

::: example
Examples where the image is purely decorative despite being relevant
would include things like a photo of the Black Rock City landscape in a
blog post about an event at Burning Man, or an image of a painting
inspired by a poem, on a page reciting that poem. The following snippet
shows an example of the latter case (only the first verse is included in
this snippet):

``` html
<h1>The Lady of Shalott</h1>
<p><img src="shalott.jpeg" alt=""></p>
<p>On either side the river lie<br>
Long fields of barley and of rye,<br>
That clothe the wold and meet the sky;<br>
And through the field the road run by<br>
To many-tower'd Camelot;<br>
And up and down the people go,<br>
Gazing where the lilies blow<br>
Round an island there below,<br>
The island of Shalott.</p>
```
:::

###### [4.8.4.4.9]{.secno} A group of images that form a single larger picture with no links[](#a-group-of-images-that-form-a-single-larger-picture-with-no-links){.self-link}

When a picture has been sliced into smaller image files that are then
displayed together to form the complete picture again, one of the images
must have its
[`alt`{#a-group-of-images-that-form-a-single-larger-picture-with-no-links:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute set as per the relevant rules that would be appropriate for
the picture as a whole, and then all the remaining images must have
their
[`alt`{#a-group-of-images-that-form-a-single-larger-picture-with-no-links:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute set to the empty string.

::: example
In the following example, a picture representing a company logo for XYZ
Corp has been split into two pieces, the first containing the letters
\"XYZ\" and the second with the word \"Corp\". The alternative text
(\"XYZ Corp\") is all in the first image.

``` html
<h1><img src="logo1.png" alt="XYZ Corp"><img src="logo2.png" alt=""></h1>
```
:::

::: example
In the following example, a rating is shown as three filled stars and
two empty stars. While the alternative text could have been \"★★★☆☆\",
the author has instead decided to more helpfully give the rating in the
form \"3 out of 5\". That is the alternative text of the first image,
and the rest have blank alternative text.

``` html
<p>Rating: <meter max=5 value=3><img src="1" alt="3 out of 5"
  ><img src="1" alt=""><img src="1" alt=""><img src="0" alt=""
  ><img src="0" alt=""></meter></p>
```
:::

###### [4.8.4.4.10]{.secno} A group of images that form a single larger picture with links[](#a-group-of-images-that-form-a-single-larger-picture-with-links){.self-link}

Generally, [image
maps](image-maps.html#image-map){#a-group-of-images-that-form-a-single-larger-picture-with-links:image-map}
should be used instead of slicing an image for links.

However, if an image is indeed sliced and any of the components of the
sliced picture are the sole contents of links, then one image per link
must have alternative text in its
[`alt`{#a-group-of-images-that-form-a-single-larger-picture-with-links:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute representing the purpose of the link.

::: example
In the following example, a picture representing the flying spaghetti
monster emblem, with each of the left noodly appendages and the right
noodly appendages in different images, so that the user can pick the
left side or the right side in an adventure.

``` html
<h1>The Church</h1>
<p>You come across a flying spaghetti monster. Which side of His
Noodliness do you wish to reach out for?</p>
<p><a href="?go=left" ><img src="fsm-left.png"  alt="Left side. "></a
  ><img src="fsm-middle.png" alt=""
  ><a href="?go=right"><img src="fsm-right.png" alt="Right side."></a></p>
```
:::

###### [4.8.4.4.11]{.secno} A key part of the content[](#a-key-part-of-the-content){.self-link}

In some cases, the image is a critical part of the content. This could
be the case, for instance, on a page that is part of a photo gallery.
The image is the whole *point* of the page containing it.

How to provide alternative text for an image that is a key part of the
content depends on the image\'s provenance.

The general case

:   When it is possible for detailed alternative text to be provided,
    for example if the image is part of a series of screenshots in a
    magazine review, or part of a comic strip, or is a photograph in a
    blog entry about that photograph, text that can serve as a
    substitute for the image must be given as the contents of the
    [`alt`{#a-key-part-of-the-content:attr-img-alt}](embedded-content.html#attr-img-alt)
    attribute.

    ::: example
    A screenshot in a gallery of screenshots for a new OS, with some
    alternative text:

    ``` html
    <figure>
     <img src="KDE%20Light%20desktop.png"
          alt="The desktop is blue, with icons along the left hand side in
               two columns, reading System, Home, K-Mail, etc. A window is
               open showing that menus wrap to a second line if they
               cannot fit in the window. The window has a list of icons
               along the top, with an address bar below it, a list of
               icons for tabs along the left edge, a status bar on the
               bottom, and two panes in the middle. The desktop has a bar
               at the bottom of the screen with a few buttons, a pager, a
               list of open applications, and a clock.">
     <figcaption>Screenshot of a KDE desktop.</figcaption>
    </figure>
    ```
    :::

    ::: example
    A graph in a financial report:

    ``` html
    <img src="sales.gif"
         title="Sales graph"
         alt="From 1998 to 2005, sales increased by the following percentages
         with each year: 624%, 75%, 138%, 40%, 35%, 9%, 21%">
    ```

    Note that \"sales graph\" would be inadequate alternative text for a
    sales graph. Text that would be a good *caption* is not generally
    suitable as replacement text.
    :::

Images that defy a complete description

:   In certain cases, the nature of the image might be such that
    providing thorough alternative text is impractical. For example, the
    image could be indistinct, or could be a complex fractal, or could
    be a detailed topographical map.

    In these cases, the
    [`alt`{#a-key-part-of-the-content:attr-img-alt-2}](embedded-content.html#attr-img-alt)
    attribute must contain some suitable alternative text, but it may be
    somewhat brief.

    ::: example
    Sometimes there simply is no text that can do justice to an image.
    For example, there is little that can be said to usefully describe a
    Rorschach inkblot test. However, a description, even if brief, is
    still better than nothing:

    ``` html
    <figure>
     <img src="/commons/a/a7/Rorschach1.jpg" alt="A shape with left-right
     symmetry with indistinct edges, with a small gap in the center, two
     larger gaps offset slightly from the center, with two similar gaps
     under them. The outline is wider in the top half than the bottom
     half, with the sides extending upwards higher than the center, and
     the center extending below the sides.">
     <figcaption>A black outline of the first of the ten cards
     in the Rorschach inkblot test.</figcaption>
    </figure>
    ```

    Note that the following would be a very bad use of alternative text:

    ``` bad
    <!-- This example is wrong. Do not copy it. -->
    <figure>
     <img src="/commons/a/a7/Rorschach1.jpg" alt="A black outline
     of the first of the ten cards in the Rorschach inkblot test.">
     <figcaption>A black outline of the first of the ten cards
     in the Rorschach inkblot test.</figcaption>
    </figure>
    ```

    Including the caption in the alternative text like this isn\'t
    useful because it effectively duplicates the caption for users who
    don\'t have images, taunting them twice yet not helping them any
    more than if they had only read or heard the caption once.
    :::

    ::: example
    Another example of an image that defies full description is a
    fractal, which, by definition, is infinite in detail.

    The following example shows one possible way of providing
    alternative text for the full view of an image of the Mandelbrot
    set.

    ``` html
    <img src="ms1.jpeg" alt="The Mandelbrot set appears as a cardioid with
    its cusp on the real axis in the positive direction, with a smaller
    bulb aligned along the same center line, touching it in the negative
    direction, and with these two shapes being surrounded by smaller bulbs
    of various sizes.">
    ```
    :::

    ::: example
    Similarly, a photograph of a person\'s face, for example in a
    biography, can be considered quite relevant and key to the content,
    but it can be hard to fully substitute text for:

    ``` html
    <section class="bio">
     <h1>A Biography of Isaac Asimov</h1>
     <p>Born <b>Isaak Yudovich Ozimov</b> in 1920, Isaac was a prolific author.</p>
     <p><img src="headpics/asimov.jpeg" alt="Isaac Asimov had dark hair, a tall forehead, and wore glasses.
     Later in life, he wore long white sideburns."></p>
     <p>Asimov was born in Russia, and moved to the US when he was three years old.</p>
     <p>...</p>
    </section>
    ```

    In such cases it is unnecessary (and indeed discouraged) to include
    a reference to the presence of the image itself in the alternative
    text, since such text would be redundant with the browser itself
    reporting the presence of the image. For example, if the alternative
    text was \"A photo of Isaac Asimov\", then a conforming user agent
    might read that out as \"(Image) A photo of Isaac Asimov\" rather
    than the more useful \"(Image) Isaac Asimov had dark hair, a tall
    forehead, and wore glasses\...\".
    :::

Images whose contents are not known

:   In some unfortunate cases, there might be no alternative text
    available at all, either because the image is obtained in some
    automated fashion without any associated alternative text (e.g., a
    webcam), or because the page is being generated by a script using
    user-provided images where the user did not provide suitable or
    usable alternative text (e.g. photograph sharing sites), or because
    the author does not themself know what the images represent (e.g. a
    blind photographer sharing an image on their blog).

    In such cases, the
    [`alt`{#a-key-part-of-the-content:attr-img-alt-3}](embedded-content.html#attr-img-alt)
    attribute may be omitted, but one of the following conditions must
    be met as well:

    - [The
      [`img`{#a-key-part-of-the-content:the-img-element}](embedded-content.html#the-img-element)
      element is in a
      [`figure`{#a-key-part-of-the-content:the-figure-element}](grouping-content.html#the-figure-element)
      element that contains a
      [`figcaption`{#a-key-part-of-the-content:the-figcaption-element}](grouping-content.html#the-figcaption-element)
      element that contains content other than [inter-element
      whitespace](dom.html#inter-element-whitespace){#a-key-part-of-the-content:inter-element-whitespace},
      and, ignoring the
      [`figcaption`{#a-key-part-of-the-content:the-figcaption-element-2}](grouping-content.html#the-figcaption-element)
      element and its descendants, the
      [`figure`{#a-key-part-of-the-content:the-figure-element-2}](grouping-content.html#the-figure-element)
      element has no [flow
      content](dom.html#flow-content-2){#a-key-part-of-the-content:flow-content-2}
      descendants other than [inter-element
      whitespace](dom.html#inter-element-whitespace){#a-key-part-of-the-content:inter-element-whitespace-2}
      and the
      [`img`{#a-key-part-of-the-content:the-img-element-2}](embedded-content.html#the-img-element)
      element.]{#figcaption-as-alt-condition}

    - The
      [`title`{#a-key-part-of-the-content:attr-title}](dom.html#attr-title)
      attribute is present and has a non-empty value.

      Relying on the
      [`title`{#a-key-part-of-the-content:attr-title-2}](dom.html#attr-title)
      attribute is currently discouraged as many user agents do not
      expose the attribute in an accessible manner as required by this
      specification (e.g. requiring a pointing device such as a mouse to
      cause a tooltip to appear, which excludes keyboard-only users and
      touch-only users, such as anyone with a modern phone or tablet).

    Such cases are to be kept to an absolute minimum. If there is even
    the slightest possibility of the author having the ability to
    provide real alternative text, then it would not be acceptable to
    omit the
    [`alt`{#a-key-part-of-the-content:attr-img-alt-4}](embedded-content.html#attr-img-alt)
    attribute.

    ::: example
    A photo on a photo-sharing site, if the site received the image with
    no metadata other than the caption, could be marked up as follows:

    ``` html
    <figure>
     <img src="1100670787_6a7c664aef.jpg">
     <figcaption>Bubbles traveled everywhere with us.</figcaption>
    </figure>
    ```

    It would be better, however, if a detailed description of the
    important parts of the image obtained from the user and included on
    the page.
    :::

    ::: example
    A blind user\'s blog in which a photo taken by the user is shown.
    Initially, the user might not have any idea what the photo they took
    shows:

    ``` html
    <article>
     <h1>I took a photo</h1>
     <p>I went out today and took a photo!</p>
     <figure>
      <img src="photo2.jpeg">
      <figcaption>A photograph taken blindly from my front porch.</figcaption>
     </figure>
    </article>
    ```

    Eventually though, the user might obtain a description of the image
    from their friends and could then include alternative text:

    ``` html
    <article>
     <h1>I took a photo</h1>
     <p>I went out today and took a photo!</p>
     <figure>
      <img src="photo2.jpeg" alt="The photograph shows my squirrel
      feeder hanging from the edge of my roof. It is half full, but there
      are no squirrels around. In the background, out-of-focus trees fill the
      shot. The feeder is made of wood with a metal grate, and it contains
      peanuts. The edge of the roof is wooden too, and is painted white
      with light blue streaks.">
      <figcaption>A photograph taken blindly from my front porch.</figcaption>
     </figure>
    </article>
    ```
    :::

    ::: example
    Sometimes the entire point of the image is that a textual
    description is not available, and the user is to provide the
    description. For instance, the point of a CAPTCHA image is to see if
    the user can literally read the graphic. Here is one way to mark up
    a CAPTCHA (note the
    [`title`{#a-key-part-of-the-content:attr-title-3}](dom.html#attr-title)
    attribute):

    ``` html
    <p><label>What does this image say?
    <img src="captcha.cgi?id=8934" title="CAPTCHA">
    <input type=text name=captcha></label>
    (If you cannot see the image, you can use an <a
    href="?audio">audio</a> test instead.)</p>
    ```

    Another example would be software that displays images and asks for
    alternative text precisely for the purpose of then writing a page
    with correct alternative text. Such a page could have a table of
    images, like this:

    ``` html
    <table>
     <thead>
      <tr> <th> Image <th> Description
     <tbody>
      <tr>
       <td> <img src="2421.png" title="Image 640 by 100, filename 'banner.gif'">
       <td> <input name="alt2421">
      <tr>
       <td> <img src="2422.png" title="Image 200 by 480, filename 'ad3.gif'">
       <td> <input name="alt2422">
    </table>
    ```

    Notice that even in this example, as much useful information as
    possible is still included in the
    [`title`{#a-key-part-of-the-content:attr-title-4}](dom.html#attr-title)
    attribute.
    :::

    Since some users cannot use images at all (e.g. because they have a
    very slow connection, or because they are using a text-only browser,
    or because they are listening to the page being read out by a
    hands-free automobile voice web browser, or simply because they are
    blind), the
    [`alt`{#a-key-part-of-the-content:attr-img-alt-5}](embedded-content.html#attr-img-alt)
    attribute is only allowed to be omitted rather than being provided
    with replacement text when no alternative text is available and none
    can be made available, as in the above examples. Lack of effort from
    the part of the author is not an acceptable reason for omitting the
    [`alt`{#a-key-part-of-the-content:attr-img-alt-6}](embedded-content.html#attr-img-alt)
    attribute.

###### [4.8.4.4.12]{.secno} An image not intended for the user[](#an-image-not-intended-for-the-user){.self-link}

Generally authors should avoid using
[`img`{#an-image-not-intended-for-the-user:the-img-element}](embedded-content.html#the-img-element)
elements for purposes other than showing images.

If an
[`img`{#an-image-not-intended-for-the-user:the-img-element-2}](embedded-content.html#the-img-element)
element is being used for purposes other than showing an image, e.g. as
part of a service to count page views, then the
[`alt`{#an-image-not-intended-for-the-user:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute must be the empty string.

In such cases, the
[`width`{#an-image-not-intended-for-the-user:attr-dim-width}](embedded-content-other.html#attr-dim-width)
and
[`height`{#an-image-not-intended-for-the-user:attr-dim-height}](embedded-content-other.html#attr-dim-height)
attributes should both be set to zero.

###### [4.8.4.4.13]{.secno} An image in an email or private document intended for a specific person who is known to be able to view images[](#an-image-in-an-e-mail-or-private-document-intended-for-a-specific-person-who-is-known-to-be-able-to-view-images){.self-link} {#an-image-in-an-e-mail-or-private-document-intended-for-a-specific-person-who-is-known-to-be-able-to-view-images}

*This section does not apply to documents that are publicly accessible,
or whose target audience is not necessarily personally known to the
author, such as documents on a web site, emails sent to public mailing
lists, or software documentation.*

When an image is included in a private communication (such as an HTML
email) aimed at a specific person who is known to be able to view
images, the
[`alt`{#an-image-in-an-e-mail-or-private-document-intended-for-a-specific-person-who-is-known-to-be-able-to-view-images:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute may be omitted. However, even in such cases authors are
strongly urged to include alternative text (as appropriate according to
the kind of image involved, as described in the above entries), so that
the email is still usable should the user use a mail client that does
not support images, or should the document be forwarded on to other
users whose abilities might not include easily seeing images.

###### [4.8.4.4.14]{.secno} Guidance for markup generators[](#guidance-for-markup-generators){.self-link}

Markup generators (such as WYSIWYG authoring tools) should, wherever
possible, obtain alternative text from their users. However, it is
recognized that in many cases, this will not be possible.

For images that are the sole contents of links, markup generators should
examine the link target to determine the title of the target, or the URL
of the target, and use information obtained in this manner as the
alternative text.

For images that have captions, markup generators should use the
[`figure`{#guidance-for-markup-generators:the-figure-element}](grouping-content.html#the-figure-element)
and
[`figcaption`{#guidance-for-markup-generators:the-figcaption-element}](grouping-content.html#the-figcaption-element)
elements, or the
[`title`{#guidance-for-markup-generators:attr-title}](dom.html#attr-title)
attribute, to provide the image\'s caption.

As a last resort, implementers should either set the
[`alt`{#guidance-for-markup-generators:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute to the empty string, under the assumption that the image is a
purely decorative image that doesn\'t add any information but is still
specific to the surrounding content, or omit the
[`alt`{#guidance-for-markup-generators:attr-img-alt-2}](embedded-content.html#attr-img-alt)
attribute altogether, under the assumption that the image is a key part
of the content.

Markup generators may specify a
[`generator-unable-to-provide-required-alt`]{#attr-img-generator-unable-to-provide-required-alt
.dfn} attribute on
[`img`{#guidance-for-markup-generators:the-img-element}](embedded-content.html#the-img-element)
elements for which they have been unable to obtain alternative text and
for which they have therefore omitted the
[`alt`{#guidance-for-markup-generators:attr-img-alt-3}](embedded-content.html#attr-img-alt)
attribute. The value of this attribute must be the empty string.
Documents containing such attributes are not conforming, but conformance
checkers will [silently ignore](#guidance-for-conformance-checkers) this
error.

This is intended to avoid markup generators from being pressured into
replacing the error of omitting the
[`alt`{#guidance-for-markup-generators:attr-img-alt-4}](embedded-content.html#attr-img-alt)
attribute with the even more egregious error of providing phony
alternative text, because state-of-the-art automated conformance
checkers cannot distinguish phony alternative text from correct
alternative text.

Markup generators should generally avoid using the image\'s own filename
as the alternative text. Similarly, markup generators should avoid
generating alternative text from any content that will be equally
available to presentation user agents (e.g., web browsers).

This is because once a page is generated, it will typically not be
updated, whereas the browsers that later read the page can be updated by
the user, therefore the browser is likely to have more up-to-date and
finely-tuned heuristics than the markup generator did when generating
the page.

###### [4.8.4.4.15]{.secno} Guidance for conformance checkers[](#guidance-for-conformance-checkers){.self-link}

A conformance checker must report the lack of an
[`alt`{#guidance-for-conformance-checkers:attr-img-alt}](embedded-content.html#attr-img-alt)
attribute as an error unless one of the conditions listed below applies:

- The
  [`img`{#guidance-for-conformance-checkers:the-img-element}](embedded-content.html#the-img-element)
  element is in a
  [`figure`{#guidance-for-conformance-checkers:the-figure-element}](grouping-content.html#the-figure-element)
  element that satisfies [the conditions described
  above](#figcaption-as-alt-condition).

- The
  [`img`{#guidance-for-conformance-checkers:the-img-element-2}](embedded-content.html#the-img-element)
  element has a
  [`title`{#guidance-for-conformance-checkers:attr-title}](dom.html#attr-title)
  attribute with a value that is not the empty string (also as
  [described above](#unknown-images)).

- The conformance checker has been configured to assume that the
  document is an email or document intended for a specific person who is
  known to be able to view images.

- The
  [`img`{#guidance-for-conformance-checkers:the-img-element-3}](embedded-content.html#the-img-element)
  element has a (non-conforming)
  [`generator-unable-to-provide-required-alt`{#guidance-for-conformance-checkers:attr-img-generator-unable-to-provide-required-alt}](#attr-img-generator-unable-to-provide-required-alt)
  attribute whose value is the empty string. A conformance checker that
  is not reporting the lack of an
  [`alt`{#guidance-for-conformance-checkers:attr-img-alt-2}](embedded-content.html#attr-img-alt)
  attribute as an error must also not report the presence of the empty
  [`generator-unable-to-provide-required-alt`{#guidance-for-conformance-checkers:attr-img-generator-unable-to-provide-required-alt-2}](#attr-img-generator-unable-to-provide-required-alt)
  attribute as an error. (This case does not represent a case where the
  document is conforming, only that the generator could not determine
  appropriate alternative text --- validators are not required to show
  an error in this case, because such an error might encourage markup
  generators to include bogus alternative text purely in an attempt to
  silence validators. Naturally, conformance checkers *may* report the
  lack of an
  [`alt`{#guidance-for-conformance-checkers:attr-img-alt-3}](embedded-content.html#attr-img-alt)
  attribute as an error even in the presence of the
  [`generator-unable-to-provide-required-alt`{#guidance-for-conformance-checkers:attr-img-generator-unable-to-provide-required-alt-3}](#attr-img-generator-unable-to-provide-required-alt)
  attribute; for example, there could be a user option to report *all*
  conformance errors even those that might be the more or less
  inevitable result of using a markup generator.)

[← 4.8 Embedded content](embedded-content.html) --- [Table of
Contents](./) --- [4.8.5 The iframe element →](iframe-embed-object.html)
