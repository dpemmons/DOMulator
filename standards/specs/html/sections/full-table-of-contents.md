## Full table of contents {#contents .no-num .no-toc}

1.  [[[1]{.secno}
    Introduction](introduction.html#introduction)]{#toc-introduction}
    1.  [[1.1]{.secno} Where does this specification
        fit?](introduction.html#abstract)
    2.  [[1.2]{.secno} Is this HTML5?](introduction.html#is-this-html5?)
    3.  [[1.3]{.secno} Background](introduction.html#background)
    4.  [[1.4]{.secno} Audience](introduction.html#audience)
    5.  [[1.5]{.secno} Scope](introduction.html#scope)
    6.  [[1.6]{.secno} History](introduction.html#history-2)
    7.  [[1.7]{.secno} Design notes](introduction.html#design-notes)
        1.  [[1.7.1]{.secno} Serializability of script
            execution](introduction.html#serialisability-of-script-execution)
        2.  [[1.7.2]{.secno}
            Extensibility](introduction.html#extensibility)
    8.  [[1.8]{.secno} HTML vs XML
        syntax](introduction.html#html-vs-xhtml)
    9.  [[1.9]{.secno} Structure of this
        specification](introduction.html#structure-of-this-specification)
        1.  [[1.9.1]{.secno} How to read this
            specification](introduction.html#how-to-read-this-specification)
        2.  [[1.9.2]{.secno} Typographic
            conventions](introduction.html#typographic-conventions)
    10. [[1.10]{.secno} A quick introduction to
        HTML](introduction.html#a-quick-introduction-to-html)
        1.  [[1.10.1]{.secno} Writing secure applications with
            HTML](introduction.html#writing-secure-applications-with-html)
        2.  [[1.10.2]{.secno} Common pitfalls to avoid when using the
            scripting
            APIs](introduction.html#common-pitfalls-to-avoid-when-using-the-scripting-apis)
        3.  [[1.10.3]{.secno} How to catch mistakes when writing HTML:
            validators and conformance
            checkers](introduction.html#how-to-catch-mistakes-when-writing-html:-validators-and-conformance-checkers)
    11. [[1.11]{.secno} Conformance requirements for
        authors](introduction.html#conformance-requirements-for-authors)
        1.  [[1.11.1]{.secno} Presentational
            markup](introduction.html#presentational-markup)
        2.  [[1.11.2]{.secno} Syntax
            errors](introduction.html#syntax-errors)
        3.  [[1.11.3]{.secno} Restrictions on content models and on
            attribute
            values](introduction.html#restrictions-on-content-models-and-on-attribute-values)
    12. [[1.12]{.secno} Suggested
        reading](introduction.html#suggested-reading)
2.  [[[2]{.secno} Common
    infrastructure](infrastructure.html#infrastructure)]{#toc-infrastructure}
    1.  [[2.1]{.secno} Terminology](infrastructure.html#terminology)
        1.  [[2.1.1]{.secno}
            Parallelism](infrastructure.html#parallelism)
        2.  [[2.1.2]{.secno} Resources](infrastructure.html#resources)
        3.  [[2.1.3]{.secno} XML compatibility](infrastructure.html#xml)
        4.  [[2.1.4]{.secno} DOM trees](infrastructure.html#dom-trees)
        5.  [[2.1.5]{.secno} Scripting](infrastructure.html#scripting-2)
        6.  [[2.1.6]{.secno} Plugins](infrastructure.html#plugins)
        7.  [[2.1.7]{.secno} Character
            encodings](infrastructure.html#encoding-terminology)
        8.  [[2.1.8]{.secno} Conformance
            classes](infrastructure.html#conformance-classes)
        9.  [[2.1.9]{.secno}
            Dependencies](infrastructure.html#dependencies)
        10. [[2.1.10]{.secno}
            Extensibility](infrastructure.html#extensibility-2)
        11. [[2.1.11]{.secno} Interactions with XPath and
            XSLT](infrastructure.html#interactions-with-xpath-and-xslt)
    2.  [[2.2]{.secno} Policy-controlled
        features](infrastructure.html#policy-controlled-features)
    3.  [[2.3]{.secno} Common
        microsyntaxes](common-microsyntaxes.html#common-microsyntaxes)
        1.  [[2.3.1]{.secno} Common parser
            idioms](common-microsyntaxes.html#common-parser-idioms)
        2.  [[2.3.2]{.secno} Boolean
            attributes](common-microsyntaxes.html#boolean-attributes)
        3.  [[2.3.3]{.secno} Keywords and enumerated
            attributes](common-microsyntaxes.html#keywords-and-enumerated-attributes)
        4.  [[2.3.4]{.secno} Numbers](common-microsyntaxes.html#numbers)
            1.  [[2.3.4.1]{.secno} Signed
                integers](common-microsyntaxes.html#signed-integers)
            2.  [[2.3.4.2]{.secno} Non-negative
                integers](common-microsyntaxes.html#non-negative-integers)
            3.  [[2.3.4.3]{.secno} Floating-point
                numbers](common-microsyntaxes.html#floating-point-numbers)
            4.  [[2.3.4.4]{.secno} Percentages and
                lengths](common-microsyntaxes.html#percentages-and-dimensions)
            5.  [[2.3.4.5]{.secno} Nonzero percentages and
                lengths](common-microsyntaxes.html#nonzero-percentages-and-lengths)
            6.  [[2.3.4.6]{.secno} Lists of floating-point
                numbers](common-microsyntaxes.html#lists-of-floating-point-numbers)
            7.  [[2.3.4.7]{.secno} Lists of
                dimensions](common-microsyntaxes.html#lists-of-dimensions)
        5.  [[2.3.5]{.secno} Dates and
            times](common-microsyntaxes.html#dates-and-times)
            1.  [[2.3.5.1]{.secno}
                Months](common-microsyntaxes.html#months)
            2.  [[2.3.5.2]{.secno}
                Dates](common-microsyntaxes.html#dates)
            3.  [[2.3.5.3]{.secno} Yearless
                dates](common-microsyntaxes.html#yearless-dates)
            4.  [[2.3.5.4]{.secno}
                Times](common-microsyntaxes.html#times)
            5.  [[2.3.5.5]{.secno} Local dates and
                times](common-microsyntaxes.html#local-dates-and-times)
            6.  [[2.3.5.6]{.secno} Time
                zones](common-microsyntaxes.html#time-zones)
            7.  [[2.3.5.7]{.secno} Global dates and
                times](common-microsyntaxes.html#global-dates-and-times)
            8.  [[2.3.5.8]{.secno}
                Weeks](common-microsyntaxes.html#weeks)
            9.  [[2.3.5.9]{.secno}
                Durations](common-microsyntaxes.html#durations)
            10. [[2.3.5.10]{.secno} Vaguer moments in
                time](common-microsyntaxes.html#vaguer-moments-in-time)
        6.  [[2.3.6]{.secno} Legacy
            colors](common-microsyntaxes.html#colours)
        7.  [[2.3.7]{.secno} Space-separated
            tokens](common-microsyntaxes.html#space-separated-tokens)
        8.  [[2.3.8]{.secno} Comma-separated
            tokens](common-microsyntaxes.html#comma-separated-tokens)
        9.  [[2.3.9]{.secno}
            References](common-microsyntaxes.html#syntax-references)
        10. [[2.3.10]{.secno} Media
            queries](common-microsyntaxes.html#mq)
        11. [[2.3.11]{.secno} Unique internal
            values](common-microsyntaxes.html#unique-values)
    4.  [[2.4]{.secno} URLs](urls-and-fetching.html#urls)
        1.  [[2.4.1]{.secno}
            Terminology](urls-and-fetching.html#terminology-2)
        2.  [[2.4.2]{.secno} Parsing
            URLs](urls-and-fetching.html#resolving-urls)
        3.  [[2.4.3]{.secno} Dynamic changes to base
            URLs](urls-and-fetching.html#dynamic-changes-to-base-urls)
    5.  [[2.5]{.secno} Fetching
        resources](urls-and-fetching.html#fetching-resources)
        1.  [[2.5.1]{.secno}
            Terminology](urls-and-fetching.html#terminology-3)
        2.  [[2.5.2]{.secno} Determining the type of a
            resource](urls-and-fetching.html#content-type-sniffing)
        3.  [[2.5.3]{.secno} Extracting character encodings from `meta`
            elements](urls-and-fetching.html#extracting-character-encodings-from-meta-elements)
        4.  [[2.5.4]{.secno} CORS settings
            attributes](urls-and-fetching.html#cors-settings-attributes)
        5.  [[2.5.5]{.secno} Referrer policy
            attributes](urls-and-fetching.html#referrer-policy-attributes)
        6.  [[2.5.6]{.secno} Nonce
            attributes](urls-and-fetching.html#nonce-attributes)
        7.  [[2.5.7]{.secno} Lazy loading
            attributes](urls-and-fetching.html#lazy-loading-attributes)
        8.  [[2.5.8]{.secno} Blocking
            attributes](urls-and-fetching.html#blocking-attributes)
        9.  [[2.5.9]{.secno} Fetch priority
            attributes](urls-and-fetching.html#fetch-priority-attributes)
    6.  [[2.6]{.secno} Common DOM
        interfaces](common-dom-interfaces.html#common-dom-interfaces)
        1.  [[2.6.1]{.secno} Reflecting content attributes in IDL
            attributes](common-dom-interfaces.html#reflecting-content-attributes-in-idl-attributes)
        2.  [[2.6.2]{.secno} Using reflect in
            specifications](common-dom-interfaces.html#using-reflect-in-specifications)
        3.  [[2.6.3]{.secno}
            Collections](common-dom-interfaces.html#collections)
            1.  [[2.6.3.1]{.secno} The `HTMLAllCollection`
                interface](common-dom-interfaces.html#the-htmlallcollection-interface)
                1.  [[2.6.3.1.1]{.secno} \[\[Call\]\] (
                    `thisArgument`{.variable},
                    `argumentsList`{.variable}
                    )](common-dom-interfaces.html#HTMLAllCollection-call)
            2.  [[2.6.3.2]{.secno} The `HTMLFormControlsCollection`
                interface](common-dom-interfaces.html#the-htmlformcontrolscollection-interface)
            3.  [[2.6.3.3]{.secno} The `HTMLOptionsCollection`
                interface](common-dom-interfaces.html#the-htmloptionscollection-interface)
        4.  [[2.6.4]{.secno} The `DOMStringList`
            interface](common-dom-interfaces.html#the-domstringlist-interface)
    7.  [[2.7]{.secno} Safe passing of structured
        data](structured-data.html#safe-passing-of-structured-data)
        1.  [[2.7.1]{.secno} Serializable
            objects](structured-data.html#serializable-objects)
        2.  [[2.7.2]{.secno} Transferable
            objects](structured-data.html#transferable-objects)
        3.  [[2.7.3]{.secno} StructuredSerializeInternal (
            `value`{.variable}, `forStorage`{.variable} \[ ,
            `memory`{.variable} \]
            )](structured-data.html#structuredserializeinternal)
        4.  [[2.7.4]{.secno} StructuredSerialize ( `value`{.variable}
            )](structured-data.html#structuredserialize)
        5.  [[2.7.5]{.secno} StructuredSerializeForStorage (
            `value`{.variable}
            )](structured-data.html#structuredserializeforstorage)
        6.  [[2.7.6]{.secno} StructuredDeserialize (
            `serialized`{.variable}, `targetRealm`{.variable} \[ ,
            `memory`{.variable} \]
            )](structured-data.html#structureddeserialize)
        7.  [[2.7.7]{.secno} StructuredSerializeWithTransfer (
            `value`{.variable}, `transferList`{.variable}
            )](structured-data.html#structuredserializewithtransfer)
        8.  [[2.7.8]{.secno} StructuredDeserializeWithTransfer (
            `serializeWithTransferResult`{.variable},
            `targetRealm`{.variable}
            )](structured-data.html#structureddeserializewithtransfer)
        9.  [[2.7.9]{.secno} Performing serialization and transferring
            from other
            specifications](structured-data.html#performing-structured-clones-from-other-specifications)
        10. [[2.7.10]{.secno} Structured cloning
            API](structured-data.html#structured-cloning)
3.  [[[3]{.secno} Semantics, structure, and APIs of HTML
    documents](dom.html#dom)]{#toc-dom}
    1.  [[3.1]{.secno} Documents](dom.html#documents)
        1.  [[3.1.1]{.secno} The `Document`
            object](dom.html#the-document-object)
        2.  [[3.1.2]{.secno} The `DocumentOrShadowRoot`
            interface](dom.html#the-documentorshadowroot-interface)
        3.  [[3.1.3]{.secno} Resource metadata
            management](dom.html#resource-metadata-management)
        4.  [[3.1.4]{.secno} Reporting document loading
            status](dom.html#reporting-document-loading-status)
        5.  [[3.1.5]{.secno} Render-blocking
            mechanism](dom.html#render-blocking-mechanism)
        6.  [[3.1.6]{.secno} DOM tree
            accessors](dom.html#dom-tree-accessors)
    2.  [[3.2]{.secno} Elements](dom.html#elements)
        1.  [[3.2.1]{.secno} Semantics](dom.html#semantics-2)
        2.  [[3.2.2]{.secno} Elements in the
            DOM](dom.html#elements-in-the-dom)
        3.  [[3.2.3]{.secno} HTML element
            constructors](dom.html#html-element-constructors)
        4.  [[3.2.4]{.secno} Element
            definitions](dom.html#element-definitions)
            1.  [[3.2.4.1]{.secno} Attributes](dom.html#attributes)
        5.  [[3.2.5]{.secno} Content models](dom.html#content-models)
            1.  [[3.2.5.1]{.secno} The \"nothing\" content
                model](dom.html#the-nothing-content-model)
            2.  [[3.2.5.2]{.secno} Kinds of
                content](dom.html#kinds-of-content)
                1.  [[3.2.5.2.1]{.secno} Metadata
                    content](dom.html#metadata-content)
                2.  [[3.2.5.2.2]{.secno} Flow
                    content](dom.html#flow-content)
                3.  [[3.2.5.2.3]{.secno} Sectioning
                    content](dom.html#sectioning-content)
                4.  [[3.2.5.2.4]{.secno} Heading
                    content](dom.html#heading-content)
                5.  [[3.2.5.2.5]{.secno} Phrasing
                    content](dom.html#phrasing-content)
                6.  [[3.2.5.2.6]{.secno} Embedded
                    content](dom.html#embedded-content-2)
                7.  [[3.2.5.2.7]{.secno} Interactive
                    content](dom.html#interactive-content)
                8.  [[3.2.5.2.8]{.secno} Palpable
                    content](dom.html#palpable-content)
                9.  [[3.2.5.2.9]{.secno} Script-supporting
                    elements](dom.html#script-supporting-elements)
            3.  [[3.2.5.3]{.secno} Transparent content
                models](dom.html#transparent-content-models)
            4.  [[3.2.5.4]{.secno} Paragraphs](dom.html#paragraphs)
        6.  [[3.2.6]{.secno} Global
            attributes](dom.html#global-attributes)
            1.  [[3.2.6.1]{.secno} The `title`
                attribute](dom.html#the-title-attribute)
            2.  [[3.2.6.2]{.secno} The `lang` and `xml:lang`
                attributes](dom.html#the-lang-and-xml:lang-attributes)
            3.  [[3.2.6.3]{.secno} The `translate`
                attribute](dom.html#the-translate-attribute)
            4.  [[3.2.6.4]{.secno} The `dir`
                attribute](dom.html#the-dir-attribute)
            5.  [[3.2.6.5]{.secno} The `style`
                attribute](dom.html#the-style-attribute)
            6.  [[3.2.6.6]{.secno} Embedding custom non-visible data
                with the `data-*`
                attributes](dom.html#embedding-custom-non-visible-data-with-the-data-*-attributes)
        7.  [[3.2.7]{.secno} The `innerText` and `outerText`
            properties](dom.html#the-innertext-idl-attribute)
        8.  [[3.2.8]{.secno} Requirements relating to the bidirectional
            algorithm](dom.html#requirements-relating-to-the-bidirectional-algorithm)
            1.  [[3.2.8.1]{.secno} Authoring conformance criteria for
                bidirectional-algorithm formatting
                characters](dom.html#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters)
            2.  [[3.2.8.2]{.secno} User agent conformance
                criteria](dom.html#user-agent-conformance-criteria)
        9.  [[3.2.9]{.secno} Requirements related to ARIA and to
            platform accessibility APIs](dom.html#wai-aria)
4.  [[[4]{.secno} The elements of
    HTML](semantics.html#semantics)]{#toc-semantics}
    1.  [[4.1]{.secno} The document
        element](semantics.html#the-root-element)
        1.  [[4.1.1]{.secno} The `html`
            element](semantics.html#the-html-element)
    2.  [[4.2]{.secno} Document
        metadata](semantics.html#document-metadata)
        1.  [[4.2.1]{.secno} The `head`
            element](semantics.html#the-head-element)
        2.  [[4.2.2]{.secno} The `title`
            element](semantics.html#the-title-element)
        3.  [[4.2.3]{.secno} The `base`
            element](semantics.html#the-base-element)
        4.  [[4.2.4]{.secno} The `link`
            element](semantics.html#the-link-element)
            1.  [[4.2.4.1]{.secno} Processing the `media`
                attribute](semantics.html#processing-the-media-attribute)
            2.  [[4.2.4.2]{.secno} Processing the `type`
                attribute](semantics.html#processing-the-type-attribute)
            3.  [[4.2.4.3]{.secno} Fetching and processing a resource
                from a `link`
                element](semantics.html#fetching-and-processing-a-resource-from-a-link-element)
            4.  [[4.2.4.4]{.secno} Processing \``Link`\`
                headers](semantics.html#processing-link-headers)
            5.  [[4.2.4.5]{.secno} Early
                hints](semantics.html#early-hints)
            6.  [[4.2.4.6]{.secno} Providing users with a means to
                follow hyperlinks created using the `link`
                element](semantics.html#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element)
        5.  [[4.2.5]{.secno} The `meta`
            element](semantics.html#the-meta-element)
            1.  [[4.2.5.1]{.secno} Standard metadata
                names](semantics.html#standard-metadata-names)
            2.  [[4.2.5.2]{.secno} Other metadata
                names](semantics.html#other-metadata-names)
            3.  [[4.2.5.3]{.secno} Pragma
                directives](semantics.html#pragma-directives)
            4.  [[4.2.5.4]{.secno} Specifying the document\'s character
                encoding](semantics.html#charset)
        6.  [[4.2.6]{.secno} The `style`
            element](semantics.html#the-style-element)
        7.  [[4.2.7]{.secno} Interactions of styling and
            scripting](semantics.html#interactions-of-styling-and-scripting)
    3.  [[4.3]{.secno} Sections](sections.html#sections)
        1.  [[4.3.1]{.secno} The `body`
            element](sections.html#the-body-element)
        2.  [[4.3.2]{.secno} The `article`
            element](sections.html#the-article-element)
        3.  [[4.3.3]{.secno} The `section`
            element](sections.html#the-section-element)
        4.  [[4.3.4]{.secno} The `nav`
            element](sections.html#the-nav-element)
        5.  [[4.3.5]{.secno} The `aside`
            element](sections.html#the-aside-element)
        6.  [[4.3.6]{.secno} The `h1`, `h2`, `h3`, `h4`, `h5`, and `h6`
            elements](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
        7.  [[4.3.7]{.secno} The `hgroup`
            element](sections.html#the-hgroup-element)
        8.  [[4.3.8]{.secno} The `header`
            element](sections.html#the-header-element)
        9.  [[4.3.9]{.secno} The `footer`
            element](sections.html#the-footer-element)
        10. [[4.3.10]{.secno} The `address`
            element](sections.html#the-address-element)
        11. [[4.3.11]{.secno} Headings and
            outlines](sections.html#headings-and-outlines-2)
            1.  [[4.3.11.1]{.secno} Sample
                outlines](sections.html#sample-outlines)
            2.  [[4.3.11.2]{.secno} Exposing outlines to
                users](sections.html#exposing-outlines-to-users)
        12. [[4.3.12]{.secno} Usage
            summary](sections.html#usage-summary-2)
            1.  [[4.3.12.1]{.secno} Article or
                section?](sections.html#article-or-section)
    4.  [[4.4]{.secno} Grouping
        content](grouping-content.html#grouping-content)
        1.  [[4.4.1]{.secno} The `p`
            element](grouping-content.html#the-p-element)
        2.  [[4.4.2]{.secno} The `hr`
            element](grouping-content.html#the-hr-element)
        3.  [[4.4.3]{.secno} The `pre`
            element](grouping-content.html#the-pre-element)
        4.  [[4.4.4]{.secno} The `blockquote`
            element](grouping-content.html#the-blockquote-element)
        5.  [[4.4.5]{.secno} The `ol`
            element](grouping-content.html#the-ol-element)
        6.  [[4.4.6]{.secno} The `ul`
            element](grouping-content.html#the-ul-element)
        7.  [[4.4.7]{.secno} The `menu`
            element](grouping-content.html#the-menu-element)
        8.  [[4.4.8]{.secno} The `li`
            element](grouping-content.html#the-li-element)
        9.  [[4.4.9]{.secno} The `dl`
            element](grouping-content.html#the-dl-element)
        10. [[4.4.10]{.secno} The `dt`
            element](grouping-content.html#the-dt-element)
        11. [[4.4.11]{.secno} The `dd`
            element](grouping-content.html#the-dd-element)
        12. [[4.4.12]{.secno} The `figure`
            element](grouping-content.html#the-figure-element)
        13. [[4.4.13]{.secno} The `figcaption`
            element](grouping-content.html#the-figcaption-element)
        14. [[4.4.14]{.secno} The `main`
            element](grouping-content.html#the-main-element)
        15. [[4.4.15]{.secno} The `search`
            element](grouping-content.html#the-search-element)
        16. [[4.4.16]{.secno} The `div`
            element](grouping-content.html#the-div-element)
    5.  [[4.5]{.secno} Text-level
        semantics](text-level-semantics.html#text-level-semantics)
        1.  [[4.5.1]{.secno} The `a`
            element](text-level-semantics.html#the-a-element)
        2.  [[4.5.2]{.secno} The `em`
            element](text-level-semantics.html#the-em-element)
        3.  [[4.5.3]{.secno} The `strong`
            element](text-level-semantics.html#the-strong-element)
        4.  [[4.5.4]{.secno} The `small`
            element](text-level-semantics.html#the-small-element)
        5.  [[4.5.5]{.secno} The `s`
            element](text-level-semantics.html#the-s-element)
        6.  [[4.5.6]{.secno} The `cite`
            element](text-level-semantics.html#the-cite-element)
        7.  [[4.5.7]{.secno} The `q`
            element](text-level-semantics.html#the-q-element)
        8.  [[4.5.8]{.secno} The `dfn`
            element](text-level-semantics.html#the-dfn-element)
        9.  [[4.5.9]{.secno} The `abbr`
            element](text-level-semantics.html#the-abbr-element)
        10. [[4.5.10]{.secno} The `ruby`
            element](text-level-semantics.html#the-ruby-element)
        11. [[4.5.11]{.secno} The `rt`
            element](text-level-semantics.html#the-rt-element)
        12. [[4.5.12]{.secno} The `rp`
            element](text-level-semantics.html#the-rp-element)
        13. [[4.5.13]{.secno} The `data`
            element](text-level-semantics.html#the-data-element)
        14. [[4.5.14]{.secno} The `time`
            element](text-level-semantics.html#the-time-element)
        15. [[4.5.15]{.secno} The `code`
            element](text-level-semantics.html#the-code-element)
        16. [[4.5.16]{.secno} The `var`
            element](text-level-semantics.html#the-var-element)
        17. [[4.5.17]{.secno} The `samp`
            element](text-level-semantics.html#the-samp-element)
        18. [[4.5.18]{.secno} The `kbd`
            element](text-level-semantics.html#the-kbd-element)
        19. [[4.5.19]{.secno} The `sub` and `sup`
            elements](text-level-semantics.html#the-sub-and-sup-elements)
        20. [[4.5.20]{.secno} The `i`
            element](text-level-semantics.html#the-i-element)
        21. [[4.5.21]{.secno} The `b`
            element](text-level-semantics.html#the-b-element)
        22. [[4.5.22]{.secno} The `u`
            element](text-level-semantics.html#the-u-element)
        23. [[4.5.23]{.secno} The `mark`
            element](text-level-semantics.html#the-mark-element)
        24. [[4.5.24]{.secno} The `bdi`
            element](text-level-semantics.html#the-bdi-element)
        25. [[4.5.25]{.secno} The `bdo`
            element](text-level-semantics.html#the-bdo-element)
        26. [[4.5.26]{.secno} The `span`
            element](text-level-semantics.html#the-span-element)
        27. [[4.5.27]{.secno} The `br`
            element](text-level-semantics.html#the-br-element)
        28. [[4.5.28]{.secno} The `wbr`
            element](text-level-semantics.html#the-wbr-element)
        29. [[4.5.29]{.secno} Usage
            summary](text-level-semantics.html#usage-summary)
    6.  [[4.6]{.secno} Links](links.html#links)
        1.  [[4.6.1]{.secno} Introduction](links.html#introduction-2)
        2.  [[4.6.2]{.secno} Links created by `a` and `area`
            elements](links.html#links-created-by-a-and-area-elements)
        3.  [[4.6.3]{.secno} API for `a` and `area`
            elements](links.html#api-for-a-and-area-elements)
        4.  [[4.6.4]{.secno} Following
            hyperlinks](links.html#following-hyperlinks)
        5.  [[4.6.5]{.secno} Downloading
            resources](links.html#downloading-resources)
        6.  [[4.6.6]{.secno} Hyperlink
            auditing](links.html#hyperlink-auditing)
            1.  [[4.6.6.1]{.secno} The \``Ping-From`\` and \``Ping-To`\`
                headers](links.html#the-ping-headers)
        7.  [[4.6.7]{.secno} Link types](links.html#linkTypes)
            1.  [[4.6.7.1]{.secno} Link type
                \"`alternate`\"](links.html#rel-alternate)
            2.  [[4.6.7.2]{.secno} Link type
                \"`author`\"](links.html#link-type-author)
            3.  [[4.6.7.3]{.secno} Link type
                \"`bookmark`\"](links.html#link-type-bookmark)
            4.  [[4.6.7.4]{.secno} Link type
                \"`canonical`\"](links.html#link-type-canonical)
            5.  [[4.6.7.5]{.secno} Link type
                \"`dns-prefetch`\"](links.html#link-type-dns-prefetch)
            6.  [[4.6.7.6]{.secno} Link type
                \"`expect`\"](links.html#link-type-expect)
            7.  [[4.6.7.7]{.secno} Link type
                \"`external`\"](links.html#link-type-external)
            8.  [[4.6.7.8]{.secno} Link type
                \"`help`\"](links.html#link-type-help)
            9.  [[4.6.7.9]{.secno} Link type
                \"`icon`\"](links.html#rel-icon)
            10. [[4.6.7.10]{.secno} Link type
                \"`license`\"](links.html#link-type-license)
            11. [[4.6.7.11]{.secno} Link type
                \"`manifest`\"](links.html#link-type-manifest)
            12. [[4.6.7.12]{.secno} Link type
                \"`modulepreload`\"](links.html#link-type-modulepreload)
            13. [[4.6.7.13]{.secno} Link type
                \"`nofollow`\"](links.html#link-type-nofollow)
            14. [[4.6.7.14]{.secno} Link type
                \"`noopener`\"](links.html#link-type-noopener)
            15. [[4.6.7.15]{.secno} Link type
                \"`noreferrer`\"](links.html#link-type-noreferrer)
            16. [[4.6.7.16]{.secno} Link type
                \"`opener`\"](links.html#link-type-opener)
            17. [[4.6.7.17]{.secno} Link type
                \"`pingback`\"](links.html#link-type-pingback)
            18. [[4.6.7.18]{.secno} Link type
                \"`preconnect`\"](links.html#link-type-preconnect)
            19. [[4.6.7.19]{.secno} Link type
                \"`prefetch`\"](links.html#link-type-prefetch)
            20. [[4.6.7.20]{.secno} Link type
                \"`preload`\"](links.html#link-type-preload)
            21. [[4.6.7.21]{.secno} Link type
                \"`privacy-policy`\"](links.html#link-type-privacy-policy)
            22. [[4.6.7.22]{.secno} Link type
                \"`search`\"](links.html#link-type-search)
            23. [[4.6.7.23]{.secno} Link type
                \"`stylesheet`\"](links.html#link-type-stylesheet)
            24. [[4.6.7.24]{.secno} Link type
                \"`tag`\"](links.html#link-type-tag)
            25. [[4.6.7.25]{.secno} Link Type
                \"`terms-of-service`\"](links.html#link-type-terms-of-service)
            26. [[4.6.7.26]{.secno} Sequential link
                types](links.html#sequential-link-types)
                1.  [[4.6.7.26.1]{.secno} Link type
                    \"`next`\"](links.html#link-type-next)
                2.  [[4.6.7.26.2]{.secno} Link type
                    \"`prev`\"](links.html#link-type-prev)
            27. [[4.6.7.27]{.secno} Other link
                types](links.html#other-link-types)
    7.  [[4.7]{.secno} Edits](edits.html#edits)
        1.  [[4.7.1]{.secno} The `ins`
            element](edits.html#the-ins-element)
        2.  [[4.7.2]{.secno} The `del`
            element](edits.html#the-del-element)
        3.  [[4.7.3]{.secno} Attributes common to `ins` and `del`
            elements](edits.html#attributes-common-to-ins-and-del-elements)
        4.  [[4.7.4]{.secno} Edits and
            paragraphs](edits.html#edits-and-paragraphs)
        5.  [[4.7.5]{.secno} Edits and
            lists](edits.html#edits-and-lists)
        6.  [[4.7.6]{.secno} Edits and
            tables](edits.html#edits-and-tables)
    8.  [[4.8]{.secno} Embedded
        content](embedded-content.html#embedded-content)
        1.  [[4.8.1]{.secno} The `picture`
            element](embedded-content.html#the-picture-element)
        2.  [[4.8.2]{.secno} The `source`
            element](embedded-content.html#the-source-element)
        3.  [[4.8.3]{.secno} The `img`
            element](embedded-content.html#the-img-element)
        4.  [[4.8.4]{.secno} Images](images.html#images)
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
        5.  [[4.8.5]{.secno} The `iframe`
            element](iframe-embed-object.html#the-iframe-element)
        6.  [[4.8.6]{.secno} The `embed`
            element](iframe-embed-object.html#the-embed-element)
        7.  [[4.8.7]{.secno} The `object`
            element](iframe-embed-object.html#the-object-element)
        8.  [[4.8.8]{.secno} The `video`
            element](media.html#the-video-element)
        9.  [[4.8.9]{.secno} The `audio`
            element](media.html#the-audio-element)
        10. [[4.8.10]{.secno} The `track`
            element](media.html#the-track-element)
        11. [[4.8.11]{.secno} Media elements](media.html#media-elements)
            1.  [[4.8.11.1]{.secno} Error codes](media.html#error-codes)
            2.  [[4.8.11.2]{.secno} Location of the media
                resource](media.html#location-of-the-media-resource)
            3.  [[4.8.11.3]{.secno} MIME types](media.html#mime-types)
            4.  [[4.8.11.4]{.secno} Network
                states](media.html#network-states)
            5.  [[4.8.11.5]{.secno} Loading the media
                resource](media.html#loading-the-media-resource)
            6.  [[4.8.11.6]{.secno} Offsets into the media
                resource](media.html#offsets-into-the-media-resource)
            7.  [[4.8.11.7]{.secno} Ready
                states](media.html#ready-states)
            8.  [[4.8.11.8]{.secno} Playing the media
                resource](media.html#playing-the-media-resource)
            9.  [[4.8.11.9]{.secno} Seeking](media.html#seeking)
            10. [[4.8.11.10]{.secno} Media resources with multiple media
                tracks](media.html#media-resources-with-multiple-media-tracks)
                1.  [[4.8.11.10.1]{.secno} `AudioTrackList` and
                    `VideoTrackList`
                    objects](media.html#audiotracklist-and-videotracklist-objects)
                2.  [[4.8.11.10.2]{.secno} Selecting specific audio and
                    video tracks
                    declaratively](media.html#selecting-specific-audio-and-video-tracks-declaratively)
            11. [[4.8.11.11]{.secno} Timed text
                tracks](media.html#timed-text-tracks)
                1.  [[4.8.11.11.1]{.secno} Text track
                    model](media.html#text-track-model)
                2.  [[4.8.11.11.2]{.secno} Sourcing in-band text
                    tracks](media.html#sourcing-in-band-text-tracks)
                3.  [[4.8.11.11.3]{.secno} Sourcing out-of-band text
                    tracks](media.html#sourcing-out-of-band-text-tracks)
                4.  [[4.8.11.11.4]{.secno} Guidelines for exposing cues
                    in various formats as text track
                    cues](media.html#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues)
                5.  [[4.8.11.11.5]{.secno} Text track
                    API](media.html#text-track-api)
                6.  [[4.8.11.11.6]{.secno} Event handlers for objects of
                    the text track APIs](media.html#cue-events)
                7.  [[4.8.11.11.7]{.secno} Best practices for metadata
                    text
                    tracks](media.html#best-practices-for-metadata-text-tracks)
            12. [[4.8.11.12]{.secno} Identifying a track kind through a
                URL](media.html#identifying-a-track-kind-through-a-url)
            13. [[4.8.11.13]{.secno} User
                interface](media.html#user-interface)
            14. [[4.8.11.14]{.secno} Time
                ranges](media.html#time-ranges)
            15. [[4.8.11.15]{.secno} The `TrackEvent`
                interface](media.html#the-trackevent-interface)
            16. [[4.8.11.16]{.secno} Events
                summary](media.html#mediaevents)
            17. [[4.8.11.17]{.secno} Security and privacy
                considerations](media.html#security-and-privacy-considerations)
            18. [[4.8.11.18]{.secno} Best practices for authors using
                media
                elements](media.html#best-practices-for-authors-using-media-elements)
            19. [[4.8.11.19]{.secno} Best practices for implementers of
                media
                elements](media.html#best-practices-for-implementers-of-media-elements)
        12. [[4.8.12]{.secno} The `map`
            element](image-maps.html#the-map-element)
        13. [[4.8.13]{.secno} The `area`
            element](image-maps.html#the-area-element)
        14. [[4.8.14]{.secno} Image maps](image-maps.html#image-maps)
            1.  [[4.8.14.1]{.secno}
                Authoring](image-maps.html#authoring)
            2.  [[4.8.14.2]{.secno} Processing
                model](image-maps.html#image-map-processing-model)
        15. [[4.8.15]{.secno}
            MathML](embedded-content-other.html#mathml)
        16. [[4.8.16]{.secno} SVG](embedded-content-other.html#svg-0)
        17. [[4.8.17]{.secno} Dimension
            attributes](embedded-content-other.html#dimension-attributes)
    9.  [[4.9]{.secno} Tabular data](tables.html#tables)
        1.  [[4.9.1]{.secno} The `table`
            element](tables.html#the-table-element)
            1.  [[4.9.1.1]{.secno} Techniques for describing
                tables](tables.html#table-descriptions-techniques)
            2.  [[4.9.1.2]{.secno} Techniques for table
                design](tables.html#table-layout-techniques)
        2.  [[4.9.2]{.secno} The `caption`
            element](tables.html#the-caption-element)
        3.  [[4.9.3]{.secno} The `colgroup`
            element](tables.html#the-colgroup-element)
        4.  [[4.9.4]{.secno} The `col`
            element](tables.html#the-col-element)
        5.  [[4.9.5]{.secno} The `tbody`
            element](tables.html#the-tbody-element)
        6.  [[4.9.6]{.secno} The `thead`
            element](tables.html#the-thead-element)
        7.  [[4.9.7]{.secno} The `tfoot`
            element](tables.html#the-tfoot-element)
        8.  [[4.9.8]{.secno} The `tr`
            element](tables.html#the-tr-element)
        9.  [[4.9.9]{.secno} The `td`
            element](tables.html#the-td-element)
        10. [[4.9.10]{.secno} The `th`
            element](tables.html#the-th-element)
        11. [[4.9.11]{.secno} Attributes common to `td` and `th`
            elements](tables.html#attributes-common-to-td-and-th-elements)
        12. [[4.9.12]{.secno} Processing
            model](tables.html#table-processing-model)
            1.  [[4.9.12.1]{.secno} Forming a
                table](tables.html#forming-a-table)
            2.  [[4.9.12.2]{.secno} Forming relationships between data
                cells and header
                cells](tables.html#header-and-data-cell-semantics)
        13. [[4.9.13]{.secno} Examples](tables.html#table-examples)
    10. [[4.10]{.secno} Forms](forms.html#forms)
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
        5.  [[4.10.5]{.secno} The `input`
            element](input.html#the-input-element)
            1.  [[4.10.5.1]{.secno} States of the `type`
                attribute](input.html#states-of-the-type-attribute)
                1.  [[4.10.5.1.1]{.secno} Hidden state
                    (`type=hidden`)](input.html#hidden-state-(type=hidden))
                2.  [[4.10.5.1.2]{.secno} Text (`type=text`) state and
                    Search state
                    (`type=search`)](input.html#text-(type=text)-state-and-search-state-(type=search))
                3.  [[4.10.5.1.3]{.secno} Telephone state
                    (`type=tel`)](input.html#telephone-state-(type=tel))
                4.  [[4.10.5.1.4]{.secno} URL state
                    (`type=url`)](input.html#url-state-(type=url))
                5.  [[4.10.5.1.5]{.secno} Email state
                    (`type=email`)](input.html#email-state-(type=email))
                6.  [[4.10.5.1.6]{.secno} Password state
                    (`type=password`)](input.html#password-state-(type=password))
                7.  [[4.10.5.1.7]{.secno} Date state
                    (`type=date`)](input.html#date-state-(type=date))
                8.  [[4.10.5.1.8]{.secno} Month state
                    (`type=month`)](input.html#month-state-(type=month))
                9.  [[4.10.5.1.9]{.secno} Week state
                    (`type=week`)](input.html#week-state-(type=week))
                10. [[4.10.5.1.10]{.secno} Time state
                    (`type=time`)](input.html#time-state-(type=time))
                11. [[4.10.5.1.11]{.secno} Local Date and Time state
                    (`type=datetime-local`)](input.html#local-date-and-time-state-(type=datetime-local))
                12. [[4.10.5.1.12]{.secno} Number state
                    (`type=number`)](input.html#number-state-(type=number))
                13. [[4.10.5.1.13]{.secno} Range state
                    (`type=range`)](input.html#range-state-(type=range))
                14. [[4.10.5.1.14]{.secno} Color state
                    (`type=color`)](input.html#color-state-(type=color))
                15. [[4.10.5.1.15]{.secno} Checkbox state
                    (`type=checkbox`)](input.html#checkbox-state-(type=checkbox))
                16. [[4.10.5.1.16]{.secno} Radio Button state
                    (`type=radio`)](input.html#radio-button-state-(type=radio))
                17. [[4.10.5.1.17]{.secno} File Upload state
                    (`type=file`)](input.html#file-upload-state-(type=file))
                18. [[4.10.5.1.18]{.secno} Submit Button state
                    (`type=submit`)](input.html#submit-button-state-(type=submit))
                19. [[4.10.5.1.19]{.secno} Image Button state
                    (`type=image`)](input.html#image-button-state-(type=image))
                20. [[4.10.5.1.20]{.secno} Reset Button state
                    (`type=reset`)](input.html#reset-button-state-(type=reset))
                21. [[4.10.5.1.21]{.secno} Button state
                    (`type=button`)](input.html#button-state-(type=button))
            2.  [[4.10.5.2]{.secno} Implementation notes regarding
                localization of form
                controls](input.html#input-impl-notes)
            3.  [[4.10.5.3]{.secno} Common `input` element
                attributes](input.html#common-input-element-attributes)
                1.  [[4.10.5.3.1]{.secno} The `maxlength` and
                    `minlength`
                    attributes](input.html#the-maxlength-and-minlength-attributes)
                2.  [[4.10.5.3.2]{.secno} The `size`
                    attribute](input.html#the-size-attribute)
                3.  [[4.10.5.3.3]{.secno} The `readonly`
                    attribute](input.html#the-readonly-attribute)
                4.  [[4.10.5.3.4]{.secno} The `required`
                    attribute](input.html#the-required-attribute)
                5.  [[4.10.5.3.5]{.secno} The `multiple`
                    attribute](input.html#the-multiple-attribute)
                6.  [[4.10.5.3.6]{.secno} The `pattern`
                    attribute](input.html#the-pattern-attribute)
                7.  [[4.10.5.3.7]{.secno} The `min` and `max`
                    attributes](input.html#the-min-and-max-attributes)
                8.  [[4.10.5.3.8]{.secno} The `step`
                    attribute](input.html#the-step-attribute)
                9.  [[4.10.5.3.9]{.secno} The `list`
                    attribute](input.html#the-list-attribute)
                10. [[4.10.5.3.10]{.secno} The `placeholder`
                    attribute](input.html#the-placeholder-attribute)
            4.  [[4.10.5.4]{.secno} Common `input` element
                APIs](input.html#common-input-element-apis)
            5.  [[4.10.5.5]{.secno} Common event
                behaviors](input.html#common-input-element-events)
        6.  [[4.10.6]{.secno} The `button`
            element](form-elements.html#the-button-element)
        7.  [[4.10.7]{.secno} The `select`
            element](form-elements.html#the-select-element)
        8.  [[4.10.8]{.secno} The `datalist`
            element](form-elements.html#the-datalist-element)
        9.  [[4.10.9]{.secno} The `optgroup`
            element](form-elements.html#the-optgroup-element)
        10. [[4.10.10]{.secno} The `option`
            element](form-elements.html#the-option-element)
        11. [[4.10.11]{.secno} The `textarea`
            element](form-elements.html#the-textarea-element)
        12. [[4.10.12]{.secno} The `output`
            element](form-elements.html#the-output-element)
        13. [[4.10.13]{.secno} The `progress`
            element](form-elements.html#the-progress-element)
        14. [[4.10.14]{.secno} The `meter`
            element](form-elements.html#the-meter-element)
        15. [[4.10.15]{.secno} The `fieldset`
            element](form-elements.html#the-fieldset-element)
        16. [[4.10.16]{.secno} The `legend`
            element](form-elements.html#the-legend-element)
        17. [[4.10.17]{.secno} Form control
            infrastructure](form-control-infrastructure.html#form-control-infrastructure)
            1.  [[4.10.17.1]{.secno} A form control\'s
                value](form-control-infrastructure.html#a-form-control's-value)
            2.  [[4.10.17.2]{.secno}
                Mutability](form-control-infrastructure.html#mutability)
            3.  [[4.10.17.3]{.secno} Association of controls and
                forms](form-control-infrastructure.html#association-of-controls-and-forms)
        18. [[4.10.18]{.secno} Attributes common to form
            controls](form-control-infrastructure.html#attributes-common-to-form-controls)
            1.  [[4.10.18.1]{.secno} Naming form controls: the `name`
                attribute](form-control-infrastructure.html#naming-form-controls:-the-name-attribute)
            2.  [[4.10.18.2]{.secno} Submitting element directionality:
                the `dirname`
                attribute](form-control-infrastructure.html#submitting-element-directionality:-the-dirname-attribute)
            3.  [[4.10.18.3]{.secno} Limiting user input length: the
                `maxlength`
                attribute](form-control-infrastructure.html#limiting-user-input-length:-the-maxlength-attribute)
            4.  [[4.10.18.4]{.secno} Setting minimum input length
                requirements: the `minlength`
                attribute](form-control-infrastructure.html#setting-minimum-input-length-requirements:-the-minlength-attribute)
            5.  [[4.10.18.5]{.secno} Enabling and disabling form
                controls: the `disabled`
                attribute](form-control-infrastructure.html#enabling-and-disabling-form-controls:-the-disabled-attribute)
            6.  [[4.10.18.6]{.secno} Form submission
                attributes](form-control-infrastructure.html#form-submission-attributes)
            7.  [[4.10.18.7]{.secno}
                Autofill](form-control-infrastructure.html#autofill)
                1.  [[4.10.18.7.1]{.secno} Autofilling form controls:
                    the `autocomplete`
                    attribute](form-control-infrastructure.html#autofilling-form-controls:-the-autocomplete-attribute)
                2.  [[4.10.18.7.2]{.secno} Processing
                    model](form-control-infrastructure.html#autofill-processing-model)
        19. [[4.10.19]{.secno} APIs for the text control
            selections](form-control-infrastructure.html#textFieldSelection)
        20. [[4.10.20]{.secno}
            Constraints](form-control-infrastructure.html#constraints)
            1.  [[4.10.20.1]{.secno}
                Definitions](form-control-infrastructure.html#definitions)
            2.  [[4.10.20.2]{.secno} Constraint
                validation](form-control-infrastructure.html#constraint-validation)
            3.  [[4.10.20.3]{.secno} The constraint validation
                API](form-control-infrastructure.html#the-constraint-validation-api)
            4.  [[4.10.20.4]{.secno}
                Security](form-control-infrastructure.html#security-forms)
        21. [[4.10.21]{.secno} Form
            submission](form-control-infrastructure.html#form-submission-2)
            1.  [[4.10.21.1]{.secno}
                Introduction](form-control-infrastructure.html#introduction-5)
            2.  [[4.10.21.2]{.secno} Implicit
                submission](form-control-infrastructure.html#implicit-submission)
            3.  [[4.10.21.3]{.secno} Form submission
                algorithm](form-control-infrastructure.html#form-submission-algorithm)
            4.  [[4.10.21.4]{.secno} Constructing the entry
                list](form-control-infrastructure.html#constructing-form-data-set)
            5.  [[4.10.21.5]{.secno} Selecting a form submission
                encoding](form-control-infrastructure.html#selecting-a-form-submission-encoding)
            6.  [[4.10.21.6]{.secno} Converting an entry list to a list
                of name-value
                pairs](form-control-infrastructure.html#converting-an-entry-list-to-a-list-of-name-value-pairs)
            7.  [[4.10.21.7]{.secno} URL-encoded form
                data](form-control-infrastructure.html#url-encoded-form-data)
            8.  [[4.10.21.8]{.secno} Multipart form
                data](form-control-infrastructure.html#multipart-form-data)
            9.  [[4.10.21.9]{.secno} Plain text form
                data](form-control-infrastructure.html#plain-text-form-data)
            10. [[4.10.21.10]{.secno} The `SubmitEvent`
                interface](form-control-infrastructure.html#the-submitevent-interface)
            11. [[4.10.21.11]{.secno} The `FormDataEvent`
                interface](form-control-infrastructure.html#the-formdataevent-interface)
        22. [[4.10.22]{.secno} Resetting a
            form](form-control-infrastructure.html#resetting-a-form)
    11. [[4.11]{.secno} Interactive
        elements](interactive-elements.html#interactive-elements)
        1.  [[4.11.1]{.secno} The `details`
            element](interactive-elements.html#the-details-element)
        2.  [[4.11.2]{.secno} The `summary`
            element](interactive-elements.html#the-summary-element)
        3.  [[4.11.3]{.secno}
            Commands](interactive-elements.html#commands)
            1.  [[4.11.3.1]{.secno}
                Facets](interactive-elements.html#facets-2)
            2.  [[4.11.3.2]{.secno} Using the `a` element to define a
                command](interactive-elements.html#using-the-a-element-to-define-a-command)
            3.  [[4.11.3.3]{.secno} Using the `button` element to define
                a
                command](interactive-elements.html#using-the-button-element-to-define-a-command)
            4.  [[4.11.3.4]{.secno} Using the `input` element to define
                a
                command](interactive-elements.html#using-the-input-element-to-define-a-command)
            5.  [[4.11.3.5]{.secno} Using the `option` element to define
                a
                command](interactive-elements.html#using-the-option-element-to-define-a-command)
            6.  [[4.11.3.6]{.secno} Using the `accesskey` attribute on a
                `legend` element to define a
                command](interactive-elements.html#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command)
            7.  [[4.11.3.7]{.secno} Using the `accesskey` attribute to
                define a command on other
                elements](interactive-elements.html#using-the-accesskey-attribute-to-define-a-command-on-other-elements)
        4.  [[4.11.4]{.secno} The `dialog`
            element](interactive-elements.html#the-dialog-element)
        5.  [[4.11.5]{.secno} Dialog light
            dismiss](interactive-elements.html#dialog-light-dismiss)
    12. [[4.12]{.secno} Scripting](scripting.html#scripting-3)
        1.  [[4.12.1]{.secno} The `script`
            element](scripting.html#the-script-element)
            1.  [[4.12.1.1]{.secno} Processing
                model](scripting.html#script-processing-model)
            2.  [[4.12.1.2]{.secno} Scripting
                languages](scripting.html#scriptingLanguages)
            3.  [[4.12.1.3]{.secno} Restrictions for contents of
                `script`
                elements](scripting.html#restrictions-for-contents-of-script-elements)
            4.  [[4.12.1.4]{.secno} Inline documentation for external
                scripts](scripting.html#inline-documentation-for-external-scripts)
            5.  [[4.12.1.5]{.secno} Interaction of `script` elements and
                XSLT](scripting.html#scriptTagXSLT)
        2.  [[4.12.2]{.secno} The `noscript`
            element](scripting.html#the-noscript-element)
        3.  [[4.12.3]{.secno} The `template`
            element](scripting.html#the-template-element)
            1.  [[4.12.3.1]{.secno} Interaction of `template` elements
                with XSLT and XPath](scripting.html#template-XSLT-XPath)
        4.  [[4.12.4]{.secno} The `slot`
            element](scripting.html#the-slot-element)
        5.  [[4.12.5]{.secno} The `canvas`
            element](canvas.html#the-canvas-element)
            1.  [[4.12.5.1]{.secno} The 2D rendering
                context](canvas.html#2dcontext)
                1.  [[4.12.5.1.1]{.secno} Implementation
                    notes](canvas.html#implementation-notes)
                2.  [[4.12.5.1.2]{.secno} The canvas
                    settings](canvas.html#the-canvas-settings)
                3.  [[4.12.5.1.3]{.secno} The canvas
                    state](canvas.html#the-canvas-state)
                4.  [[4.12.5.1.4]{.secno} Line
                    styles](canvas.html#line-styles)
                5.  [[4.12.5.1.5]{.secno} Text
                    styles](canvas.html#text-styles)
                6.  [[4.12.5.1.6]{.secno} Building
                    paths](canvas.html#building-paths)
                7.  [[4.12.5.1.7]{.secno} `Path2D`
                    objects](canvas.html#path2d-objects)
                8.  [[4.12.5.1.8]{.secno}
                    Transformations](canvas.html#transformations)
                9.  [[4.12.5.1.9]{.secno} Image sources for 2D rendering
                    contexts](canvas.html#image-sources-for-2d-rendering-contexts)
                10. [[4.12.5.1.10]{.secno} Fill and stroke
                    styles](canvas.html#fill-and-stroke-styles)
                11. [[4.12.5.1.11]{.secno} Drawing rectangles to the
                    bitmap](canvas.html#drawing-rectangles-to-the-bitmap)
                12. [[4.12.5.1.12]{.secno} Drawing text to the
                    bitmap](canvas.html#drawing-text-to-the-bitmap)
                13. [[4.12.5.1.13]{.secno} Drawing paths to the
                    canvas](canvas.html#drawing-paths-to-the-canvas)
                14. [[4.12.5.1.14]{.secno} Drawing focus
                    rings](canvas.html#drawing-focus-rings-and-scrolling-paths-into-view)
                15. [[4.12.5.1.15]{.secno} Drawing
                    images](canvas.html#drawing-images)
                16. [[4.12.5.1.16]{.secno} Pixel
                    manipulation](canvas.html#pixel-manipulation)
                17. [[4.12.5.1.17]{.secno}
                    Compositing](canvas.html#compositing)
                18. [[4.12.5.1.18]{.secno} Image
                    smoothing](canvas.html#image-smoothing)
                19. [[4.12.5.1.19]{.secno} Shadows](canvas.html#shadows)
                20. [[4.12.5.1.20]{.secno} Filters](canvas.html#filters)
                21. [[4.12.5.1.21]{.secno} Working with
                    externally-defined SVG
                    filters](canvas.html#working-with-externally-defined-svg-filters)
                22. [[4.12.5.1.22]{.secno} Drawing
                    model](canvas.html#drawing-model)
                23. [[4.12.5.1.23]{.secno} Best
                    practices](canvas.html#best-practices)
                24. [[4.12.5.1.24]{.secno}
                    Examples](canvas.html#examples)
            2.  [[4.12.5.2]{.secno} The `ImageBitmap` rendering
                context](canvas.html#the-imagebitmap-rendering-context)
                1.  [[4.12.5.2.1]{.secno}
                    Introduction](canvas.html#introduction-6)
                2.  [[4.12.5.2.2]{.secno} The
                    `ImageBitmapRenderingContext`
                    interface](canvas.html#the-imagebitmaprenderingcontext-interface)
            3.  [[4.12.5.3]{.secno} The `OffscreenCanvas`
                interface](canvas.html#the-offscreencanvas-interface)
                1.  [[4.12.5.3.1]{.secno} The offscreen 2D rendering
                    context](canvas.html#the-offscreen-2d-rendering-context)
            4.  [[4.12.5.4]{.secno} Color spaces and color space
                conversion](canvas.html#colour-spaces-and-colour-correction)
            5.  [[4.12.5.5]{.secno} Serializing bitmaps to a
                file](canvas.html#serialising-bitmaps-to-a-file)
            6.  [[4.12.5.6]{.secno} Security with `canvas`
                elements](canvas.html#security-with-canvas-elements)
            7.  [[4.12.5.7]{.secno} Premultiplied alpha and the 2D
                rendering
                context](canvas.html#premultiplied-alpha-and-the-2d-rendering-context)
    13. [[4.13]{.secno} Custom
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
    14. [[4.14]{.secno} Common idioms without dedicated
        elements](semantics-other.html#common-idioms)
        1.  [[4.14.1]{.secno} Breadcrumb
            navigation](semantics-other.html#rel-up)
        2.  [[4.14.2]{.secno} Tag
            clouds](semantics-other.html#tag-clouds)
        3.  [[4.14.3]{.secno}
            Conversations](semantics-other.html#conversations)
        4.  [[4.14.4]{.secno} Footnotes](semantics-other.html#footnotes)
    15. [[4.15]{.secno} Disabled
        elements](semantics-other.html#disabled-elements)
    16. [[4.16]{.secno} Matching HTML elements using selectors and
        CSS](semantics-other.html#selectors)
        1.  [[4.16.1]{.secno} Case-sensitivity of the CSS \'attr()\'
            function](semantics-other.html#case-sensitivity-of-the-css-'attr()'-function)
        2.  [[4.16.2]{.secno} Case-sensitivity of
            selectors](semantics-other.html#case-sensitivity-of-selectors)
        3.  [[4.16.3]{.secno}
            Pseudo-classes](semantics-other.html#pseudo-classes)
5.  [[[5]{.secno} Microdata](microdata.html#microdata)]{#toc-microdata}
    1.  [[5.1]{.secno} Introduction](microdata.html#introduction-7)
        1.  [[5.1.1]{.secno} Overview](microdata.html#overview)
        2.  [[5.1.2]{.secno} The basic
            syntax](microdata.html#the-basic-syntax)
        3.  [[5.1.3]{.secno} Typed items](microdata.html#typed-items)
        4.  [[5.1.4]{.secno} Global identifiers for
            items](microdata.html#global-identifiers-for-items)
        5.  [[5.1.5]{.secno} Selecting names when defining
            vocabularies](microdata.html#selecting-names-when-defining-vocabularies)
    2.  [[5.2]{.secno} Encoding
        microdata](microdata.html#encoding-microdata)
        1.  [[5.2.1]{.secno} The microdata
            model](microdata.html#the-microdata-model)
        2.  [[5.2.2]{.secno} Items](microdata.html#items)
        3.  [[5.2.3]{.secno} Names: the `itemprop`
            attribute](microdata.html#names:-the-itemprop-attribute)
        4.  [[5.2.4]{.secno} Values](microdata.html#values)
        5.  [[5.2.5]{.secno} Associating names with
            items](microdata.html#associating-names-with-items)
        6.  [[5.2.6]{.secno} Microdata and other
            namespaces](microdata.html#microdata-and-other-namespaces)
    3.  [[5.3]{.secno} Sample microdata
        vocabularies](microdata.html#mdvocabs)
        1.  [[5.3.1]{.secno} vCard](microdata.html#vcard)
            1.  [[5.3.1.1]{.secno} Conversion to
                vCard](microdata.html#conversion-to-vcard)
            2.  [[5.3.1.2]{.secno} Examples](microdata.html#examples-2)
        2.  [[5.3.2]{.secno} vEvent](microdata.html#vevent)
            1.  [[5.3.2.1]{.secno} Conversion to
                iCalendar](microdata.html#conversion-to-icalendar)
            2.  [[5.3.2.2]{.secno} Examples](microdata.html#examples-3)
        3.  [[5.3.3]{.secno} Licensing
            works](microdata.html#licensing-works)
            1.  [[5.3.3.1]{.secno} Examples](microdata.html#examples-4)
    4.  [[5.4]{.secno} Converting HTML to other
        formats](microdata.html#converting-html-to-other-formats)
        1.  [[5.4.1]{.secno} JSON](microdata.html#json)
6.  [[[6]{.secno} User
    interaction](interaction.html#editing)]{#toc-editing}
    1.  [[6.1]{.secno} The `hidden`
        attribute](interaction.html#the-hidden-attribute)
    2.  [[6.2]{.secno} Page
        visibility](interaction.html#page-visibility)
        1.  [[6.2.1]{.secno} The `VisibilityStateEntry`
            interface](interaction.html#the-visibilitystateentry-interface)
    3.  [[6.3]{.secno} Inert subtrees](interaction.html#inert-subtrees)
        1.  [[6.3.1]{.secno} Modal dialogs and inert
            subtrees](interaction.html#modal-dialogs-and-inert-subtrees)
        2.  [[6.3.2]{.secno} The `inert`
            attribute](interaction.html#the-inert-attribute)
    4.  [[6.4]{.secno} Tracking user
        activation](interaction.html#tracking-user-activation)
        1.  [[6.4.1]{.secno} Data
            model](interaction.html#user-activation-data-model)
        2.  [[6.4.2]{.secno} Processing
            model](interaction.html#user-activation-processing-model)
        3.  [[6.4.3]{.secno} APIs gated by user
            activation](interaction.html#user-activation-gated-apis)
        4.  [[6.4.4]{.secno} The `UserActivation`
            interface](interaction.html#the-useractivation-interface)
        5.  [[6.4.5]{.secno} User agent
            automation](interaction.html#user-activation-user-agent-automation)
    5.  [[6.5]{.secno} Activation behavior of
        elements](interaction.html#activation)
        1.  [[6.5.1]{.secno} The `ToggleEvent`
            interface](interaction.html#the-toggleevent-interface)
        2.  [[6.5.2]{.secno} The `CommandEvent`
            interface](interaction.html#the-commandevent-interface)
    6.  [[6.6]{.secno} Focus](interaction.html#focus)
        1.  [[6.6.1]{.secno}
            Introduction](interaction.html#introduction-8)
        2.  [[6.6.2]{.secno} Data model](interaction.html#data-model)
        3.  [[6.6.3]{.secno} The `tabindex`
            attribute](interaction.html#the-tabindex-attribute)
        4.  [[6.6.4]{.secno} Processing
            model](interaction.html#focus-processing-model)
        5.  [[6.6.5]{.secno} Sequential focus
            navigation](interaction.html#sequential-focus-navigation)
        6.  [[6.6.6]{.secno} Focus management
            APIs](interaction.html#focus-management-apis)
        7.  [[6.6.7]{.secno} The `autofocus`
            attribute](interaction.html#the-autofocus-attribute)
    7.  [[6.7]{.secno} Assigning keyboard
        shortcuts](interaction.html#assigning-keyboard-shortcuts)
        1.  [[6.7.1]{.secno}
            Introduction](interaction.html#introduction-9)
        2.  [[6.7.2]{.secno} The `accesskey`
            attribute](interaction.html#the-accesskey-attribute)
        3.  [[6.7.3]{.secno} Processing
            model](interaction.html#keyboard-shortcuts-processing-model)
    8.  [[6.8]{.secno} Editing](interaction.html#editing-2)
        1.  [[6.8.1]{.secno} Making document regions editable: The
            `contenteditable` content
            attribute](interaction.html#contenteditable)
        2.  [[6.8.2]{.secno} Making entire documents editable: the
            `designMode` getter and
            setter](interaction.html#making-entire-documents-editable:-the-designmode-idl-attribute)
        3.  [[6.8.3]{.secno} Best practices for in-page
            editors](interaction.html#best-practices-for-in-page-editors)
        4.  [[6.8.4]{.secno} Editing
            APIs](interaction.html#editing-apis)
        5.  [[6.8.5]{.secno} Spelling and grammar
            checking](interaction.html#spelling-and-grammar-checking)
        6.  [[6.8.6]{.secno} Writing
            suggestions](interaction.html#writing-suggestions)
        7.  [[6.8.7]{.secno}
            Autocapitalization](interaction.html#autocapitalization)
        8.  [[6.8.8]{.secno}
            Autocorrection](interaction.html#autocorrection)
        9.  [[6.8.9]{.secno} Input modalities: the `inputmode`
            attribute](interaction.html#input-modalities:-the-inputmode-attribute)
        10. [[6.8.10]{.secno} Input modalities: the `enterkeyhint`
            attribute](interaction.html#input-modalities:-the-enterkeyhint-attribute)
    9.  [[6.9]{.secno} Find-in-page](interaction.html#find-in-page)
        1.  [[6.9.1]{.secno}
            Introduction](interaction.html#introduction-10)
        2.  [[6.9.2]{.secno} Interaction with `details` and
            `hidden=until-found`](interaction.html#interaction-with-details-and-hidden=until-found)
        3.  [[6.9.3]{.secno} Interaction with
            selection](interaction.html#interaction-with-selection)
    10. [[6.10]{.secno} Close requests and close
        watchers](interaction.html#close-requests-and-close-watchers)
        1.  [[6.10.1]{.secno} Close
            requests](interaction.html#close-requests)
        2.  [[6.10.2]{.secno} Close watcher
            infrastructure](interaction.html#close-watcher-infrastructure)
        3.  [[6.10.3]{.secno} The `CloseWatcher`
            interface](interaction.html#the-closewatcher-interface)
    11. [[6.11]{.secno} Drag and drop](dnd.html#dnd)
        1.  [[6.11.1]{.secno} Introduction](dnd.html#event-drag)
        2.  [[6.11.2]{.secno} The drag data
            store](dnd.html#the-drag-data-store)
        3.  [[6.11.3]{.secno} The `DataTransfer`
            interface](dnd.html#the-datatransfer-interface)
            1.  [[6.11.3.1]{.secno} The `DataTransferItemList`
                interface](dnd.html#the-datatransferitemlist-interface)
            2.  [[6.11.3.2]{.secno} The `DataTransferItem`
                interface](dnd.html#the-datatransferitem-interface)
        4.  [[6.11.4]{.secno} The `DragEvent`
            interface](dnd.html#the-dragevent-interface)
        5.  [[6.11.5]{.secno} Processing
            model](dnd.html#drag-and-drop-processing-model)
        6.  [[6.11.6]{.secno} Events summary](dnd.html#dndevents)
        7.  [[6.11.7]{.secno} The `draggable`
            attribute](dnd.html#the-draggable-attribute)
        8.  [[6.11.8]{.secno} Security risks in the drag-and-drop
            model](dnd.html#security-risks-in-the-drag-and-drop-model)
    12. [[6.12]{.secno} The `popover`
        attribute](popover.html#the-popover-attribute)
        1.  [[6.12.1]{.secno} The popover target
            attributes](popover.html#the-popover-target-attributes)
        2.  [[6.12.2]{.secno} Popover light
            dismiss](popover.html#popover-light-dismiss)
7.  [[[7]{.secno} Loading web
    pages](browsers.html#browsers)]{#toc-browsers}
    1.  [[7.1]{.secno} Supporting
        concepts](browsers.html#loading-web-pages-supporting-concepts)
        1.  [[7.1.1]{.secno} Origins](browsers.html#origin)
            1.  [[7.1.1.1]{.secno} Sites](browsers.html#sites)
            2.  [[7.1.1.2]{.secno} Relaxing the same-origin
                restriction](browsers.html#relaxing-the-same-origin-restriction)
        2.  [[7.1.2]{.secno} Origin-keyed agent
            clusters](browsers.html#origin-keyed-agent-clusters)
        3.  [[7.1.3]{.secno} Cross-origin opener
            policies](browsers.html#cross-origin-opener-policies)
            1.  [[7.1.3.1]{.secno} The
                headers](browsers.html#the-coop-headers)
            2.  [[7.1.3.2]{.secno} Browsing context group switches due
                to opener
                policy](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy)
            3.  [[7.1.3.3]{.secno}
                Reporting](browsers.html#coop-reporting)
        4.  [[7.1.4]{.secno} Cross-origin embedder
            policies](browsers.html#coep)
            1.  [[7.1.4.1]{.secno} The
                headers](browsers.html#the-coep-headers)
            2.  [[7.1.4.2]{.secno} Embedder policy
                checks](browsers.html#embedder-policy-checks)
        5.  [[7.1.5]{.secno} Sandboxing](browsers.html#sandboxing)
        6.  [[7.1.6]{.secno} Policy
            containers](browsers.html#policy-containers)
    2.  [[7.2]{.secno} APIs related to navigation and session
        history](nav-history-apis.html#nav-traversal-apis)
        1.  [[7.2.1]{.secno} Security infrastructure for `Window`,
            `WindowProxy`, and `Location`
            objects](nav-history-apis.html#cross-origin-objects)
            1.  [[7.2.1.1]{.secno} Integration with
                IDL](nav-history-apis.html#integration-with-idl)
            2.  [[7.2.1.2]{.secno} Shared internal slot:
                \[\[CrossOriginPropertyDescriptorMap\]\]](nav-history-apis.html#shared-internal-slot:-crossoriginpropertydescriptormap)
            3.  [[7.2.1.3]{.secno} Shared abstract
                operations](nav-history-apis.html#shared-abstract-operations)
                1.  [[7.2.1.3.1]{.secno} CrossOriginProperties (
                    `O`{.variable}
                    )](nav-history-apis.html#crossoriginproperties-(-o-))
                2.  [[7.2.1.3.2]{.secno} CrossOriginPropertyFallback (
                    `P`{.variable}
                    )](nav-history-apis.html#crossoriginpropertyfallback-(-p-))
                3.  [[7.2.1.3.3]{.secno} IsPlatformObjectSameOrigin (
                    `O`{.variable}
                    )](nav-history-apis.html#isplatformobjectsameorigin-(-o-))
                4.  [[7.2.1.3.4]{.secno} CrossOriginGetOwnPropertyHelper
                    ( `O`{.variable}, `P`{.variable}
                    )](nav-history-apis.html#crossorigingetownpropertyhelper-(-o,-p-))
                5.  [[7.2.1.3.5]{.secno} CrossOriginGet (
                    `O`{.variable}, `P`{.variable},
                    `Receiver`{.variable}
                    )](nav-history-apis.html#crossoriginget-(-o,-p,-receiver-))
                6.  [[7.2.1.3.6]{.secno} CrossOriginSet (
                    `O`{.variable}, `P`{.variable}, `V`{.variable},
                    `Receiver`{.variable}
                    )](nav-history-apis.html#crossoriginset-(-o,-p,-v,-receiver-))
                7.  [[7.2.1.3.7]{.secno} CrossOriginOwnPropertyKeys (
                    `O`{.variable}
                    )](nav-history-apis.html#crossoriginownpropertykeys-(-o-))
        2.  [[7.2.2]{.secno} The `Window`
            object](nav-history-apis.html#the-window-object)
            1.  [[7.2.2.1]{.secno} Opening and closing
                windows](nav-history-apis.html#apis-for-creating-and-navigating-browsing-contexts-by-name)
            2.  [[7.2.2.2]{.secno} Indexed access on the `Window`
                object](nav-history-apis.html#accessing-other-browsing-contexts)
            3.  [[7.2.2.3]{.secno} Named access on the `Window`
                object](nav-history-apis.html#named-access-on-the-window-object)
            4.  [[7.2.2.4]{.secno} Accessing related
                windows](nav-history-apis.html#navigating-nested-browsing-contexts-in-the-dom)
            5.  [[7.2.2.5]{.secno} Historical browser interface element
                APIs](nav-history-apis.html#browser-interface-elements)
            6.  [[7.2.2.6]{.secno} Script settings for `Window`
                objects](nav-history-apis.html#script-settings-for-window-objects)
        3.  [[7.2.3]{.secno} The `WindowProxy` exotic
            object](nav-history-apis.html#the-windowproxy-exotic-object)
            1.  [[7.2.3.1]{.secno} \[\[GetPrototypeOf\]\] (
                )](nav-history-apis.html#windowproxy-getprototypeof)
            2.  [[7.2.3.2]{.secno} \[\[SetPrototypeOf\]\] (
                `V`{.variable}
                )](nav-history-apis.html#windowproxy-setprototypeof)
            3.  [[7.2.3.3]{.secno} \[\[IsExtensible\]\] (
                )](nav-history-apis.html#windowproxy-isextensible)
            4.  [[7.2.3.4]{.secno} \[\[PreventExtensions\]\] (
                )](nav-history-apis.html#windowproxy-preventextensions)
            5.  [[7.2.3.5]{.secno} \[\[GetOwnProperty\]\] (
                `P`{.variable}
                )](nav-history-apis.html#windowproxy-getownproperty)
            6.  [[7.2.3.6]{.secno} \[\[DefineOwnProperty\]\] (
                `P`{.variable}, `Desc`{.variable}
                )](nav-history-apis.html#windowproxy-defineownproperty)
            7.  [[7.2.3.7]{.secno} \[\[Get\]\] ( `P`{.variable},
                `Receiver`{.variable}
                )](nav-history-apis.html#windowproxy-get)
            8.  [[7.2.3.8]{.secno} \[\[Set\]\] ( `P`{.variable},
                `V`{.variable}, `Receiver`{.variable}
                )](nav-history-apis.html#windowproxy-set)
            9.  [[7.2.3.9]{.secno} \[\[Delete\]\] ( `P`{.variable}
                )](nav-history-apis.html#windowproxy-delete)
            10. [[7.2.3.10]{.secno} \[\[OwnPropertyKeys\]\] (
                )](nav-history-apis.html#windowproxy-ownpropertykeys)
        4.  [[7.2.4]{.secno} The `Location`
            interface](nav-history-apis.html#the-location-interface)
            1.  [[7.2.4.1]{.secno} \[\[GetPrototypeOf\]\] (
                )](nav-history-apis.html#location-getprototypeof)
            2.  [[7.2.4.2]{.secno} \[\[SetPrototypeOf\]\] (
                `V`{.variable}
                )](nav-history-apis.html#location-setprototypeof)
            3.  [[7.2.4.3]{.secno} \[\[IsExtensible\]\] (
                )](nav-history-apis.html#location-isextensible)
            4.  [[7.2.4.4]{.secno} \[\[PreventExtensions\]\] (
                )](nav-history-apis.html#location-preventextensions)
            5.  [[7.2.4.5]{.secno} \[\[GetOwnProperty\]\] (
                `P`{.variable}
                )](nav-history-apis.html#location-getownproperty)
            6.  [[7.2.4.6]{.secno} \[\[DefineOwnProperty\]\] (
                `P`{.variable}, `Desc`{.variable}
                )](nav-history-apis.html#location-defineownproperty)
            7.  [[7.2.4.7]{.secno} \[\[Get\]\] ( `P`{.variable},
                `Receiver`{.variable}
                )](nav-history-apis.html#location-get)
            8.  [[7.2.4.8]{.secno} \[\[Set\]\] ( `P`{.variable},
                `V`{.variable}, `Receiver`{.variable}
                )](nav-history-apis.html#location-set)
            9.  [[7.2.4.9]{.secno} \[\[Delete\]\] ( `P`{.variable}
                )](nav-history-apis.html#location-delete)
            10. [[7.2.4.10]{.secno} \[\[OwnPropertyKeys\]\] (
                )](nav-history-apis.html#location-ownpropertykeys)
        5.  [[7.2.5]{.secno} The `History`
            interface](nav-history-apis.html#the-history-interface)
        6.  [[7.2.6]{.secno} The navigation
            API](nav-history-apis.html#navigation-api)
            1.  [[7.2.6.1]{.secno}
                Introduction](nav-history-apis.html#navigation-api-intro)
            2.  [[7.2.6.2]{.secno} The `Navigation`
                interface](nav-history-apis.html#navigation-interface)
            3.  [[7.2.6.3]{.secno} Core
                infrastructure](nav-history-apis.html#navigation-api-core)
            4.  [[7.2.6.4]{.secno} Initializing and updating the entry
                list](nav-history-apis.html#navigation-api-entry-updates)
            5.  [[7.2.6.5]{.secno} The `NavigationHistoryEntry`
                interface](nav-history-apis.html#the-navigationhistoryentry-interface)
            6.  [[7.2.6.6]{.secno} The history entry
                list](nav-history-apis.html#the-history-entry-list)
            7.  [[7.2.6.7]{.secno} Initiating
                navigations](nav-history-apis.html#navigation-api-initiating-navigations)
            8.  [[7.2.6.8]{.secno} Ongoing navigation
                tracking](nav-history-apis.html#ongoing-navigation-tracking)
            9.  [[7.2.6.9]{.secno} The `NavigationActivation`
                interface](nav-history-apis.html#navigation-activation-interface)
            10. [[7.2.6.10]{.secno} The `navigate`
                event](nav-history-apis.html#the-navigate-event)
                1.  [[7.2.6.10.1]{.secno} The `NavigateEvent`
                    interface](nav-history-apis.html#the-navigateevent-interface)
                2.  [[7.2.6.10.2]{.secno} The `NavigationDestination`
                    interface](nav-history-apis.html#the-navigationdestination-interface)
                3.  [[7.2.6.10.3]{.secno} Firing the
                    event](nav-history-apis.html#navigate-event-firing)
                4.  [[7.2.6.10.4]{.secno} Scroll and focus
                    behavior](nav-history-apis.html#navigate-event-scroll-focus)
        7.  [[7.2.7]{.secno} Event
            interfaces](nav-history-apis.html#nav-traversal-event-interfaces)
            1.  [[7.2.7.1]{.secno} The
                `NavigationCurrentEntryChangeEvent`
                interface](nav-history-apis.html#the-navigationcurrententrychangeevent-interface)
            2.  [[7.2.7.2]{.secno} The `PopStateEvent`
                interface](nav-history-apis.html#the-popstateevent-interface)
            3.  [[7.2.7.3]{.secno} The `HashChangeEvent`
                interface](nav-history-apis.html#the-hashchangeevent-interface)
            4.  [[7.2.7.4]{.secno} The `PageSwapEvent`
                interface](nav-history-apis.html#the-pageswapevent-interface)
            5.  [[7.2.7.5]{.secno} The `PageRevealEvent`
                interface](nav-history-apis.html#the-pagerevealevent-interface)
            6.  [[7.2.7.6]{.secno} The `PageTransitionEvent`
                interface](nav-history-apis.html#the-pagetransitionevent-interface)
            7.  [[7.2.7.7]{.secno} The `BeforeUnloadEvent`
                interface](nav-history-apis.html#the-beforeunloadevent-interface)
        8.  [[7.2.8]{.secno} The `NotRestoredReasons`
            interface](nav-history-apis.html#the-notrestoredreasons-interface)
    3.  [[7.3]{.secno} Infrastructure for sequences of
        documents](document-sequences.html#infrastructure-for-sequences-of-documents)
        1.  [[7.3.1]{.secno}
            Navigables](document-sequences.html#navigables)
            1.  [[7.3.1.1]{.secno} Traversable
                navigables](document-sequences.html#traversable-navigables)
            2.  [[7.3.1.2]{.secno} Top-level
                traversables](document-sequences.html#top-level-traversables)
            3.  [[7.3.1.3]{.secno} Child
                navigables](document-sequences.html#child-navigables)
            4.  [[7.3.1.4]{.secno} Jake
                diagrams](document-sequences.html#jake-diagrams)
            5.  [[7.3.1.5]{.secno} Related navigable
                collections](document-sequences.html#related-navigable-collections)
            6.  [[7.3.1.6]{.secno} Navigable
                destruction](document-sequences.html#garbage-collection-and-browsing-contexts)
            7.  [[7.3.1.7]{.secno} Navigable target
                names](document-sequences.html#navigable-target-names)
        2.  [[7.3.2]{.secno} Browsing
            contexts](document-sequences.html#windows)
            1.  [[7.3.2.1]{.secno} Creating browsing
                contexts](document-sequences.html#creating-browsing-contexts)
            2.  [[7.3.2.2]{.secno} Related browsing
                contexts](document-sequences.html#nested-browsing-contexts)
            3.  [[7.3.2.3]{.secno} Groupings of browsing
                contexts](document-sequences.html#groupings-of-browsing-contexts)
        3.  [[7.3.3]{.secno} Fully active
            documents](document-sequences.html#fully-active-documents)
    4.  [[7.4]{.secno} Navigation and session
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
    5.  [[7.5]{.secno} Document
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
    6.  [[7.6]{.secno} The \``X-Frame-Options`\`
        header](document-lifecycle.html#the-x-frame-options-header)
    7.  [[7.7]{.secno} The \``Refresh`\`
        header](document-lifecycle.html#the-refresh-header)
    8.  [[7.8]{.secno} Browser user interface
        considerations](document-lifecycle.html#nav-traversal-ui)
8.  [[[8]{.secno} Web application
    APIs](webappapis.html#webappapis)]{#toc-webappapis}
    1.  [[8.1]{.secno} Scripting](webappapis.html#scripting)
        1.  [[8.1.1]{.secno}
            Introduction](webappapis.html#introduction-11)
        2.  [[8.1.2]{.secno} Agents and agent
            clusters](webappapis.html#agents-and-agent-clusters)
            1.  [[8.1.2.1]{.secno} Integration with the JavaScript agent
                formalism](webappapis.html#integration-with-the-javascript-agent-formalism)
            2.  [[8.1.2.2]{.secno} Integration with the JavaScript agent
                cluster
                formalism](webappapis.html#integration-with-the-javascript-agent-cluster-formalism)
        3.  [[8.1.3]{.secno} Realms and their
            counterparts](webappapis.html#realms-and-their-counterparts)
            1.  [[8.1.3.1]{.secno}
                Environments](webappapis.html#environments)
            2.  [[8.1.3.2]{.secno} Environment settings
                objects](webappapis.html#environment-settings-objects)
            3.  [[8.1.3.3]{.secno} Realms, settings objects, and global
                objects](webappapis.html#realms-settings-objects-global-objects)
                1.  [[8.1.3.3.1]{.secno} Entry](webappapis.html#entry)
                2.  [[8.1.3.3.2]{.secno}
                    Incumbent](webappapis.html#incumbent)
                3.  [[8.1.3.3.3]{.secno}
                    Current](webappapis.html#current)
                4.  [[8.1.3.3.4]{.secno}
                    Relevant](webappapis.html#relevant)
            4.  [[8.1.3.4]{.secno} Enabling and disabling
                scripting](webappapis.html#enabling-and-disabling-scripting)
            5.  [[8.1.3.5]{.secno} Secure
                contexts](webappapis.html#secure-contexts)
        4.  [[8.1.4]{.secno} Script processing
            model](webappapis.html#scripting-processing-model)
            1.  [[8.1.4.1]{.secno}
                Scripts](webappapis.html#script-structs)
            2.  [[8.1.4.2]{.secno} Fetching
                scripts](webappapis.html#fetching-scripts)
            3.  [[8.1.4.3]{.secno} Creating
                scripts](webappapis.html#creating-scripts)
            4.  [[8.1.4.4]{.secno} Calling
                scripts](webappapis.html#calling-scripts)
            5.  [[8.1.4.5]{.secno} Killing
                scripts](webappapis.html#killing-scripts)
            6.  [[8.1.4.6]{.secno} Runtime script
                errors](webappapis.html#runtime-script-errors)
            7.  [[8.1.4.7]{.secno} Unhandled promise
                rejections](webappapis.html#unhandled-promise-rejections)
            8.  [[8.1.4.8]{.secno} Import map parse
                results](webappapis.html#import-map-parse-results)
        5.  [[8.1.5]{.secno} Module specifier
            resolution](webappapis.html#module-specifier-resolution)
            1.  [[8.1.5.1]{.secno} The resolution
                algorithm](webappapis.html#the-resolution-algorithm)
            2.  [[8.1.5.2]{.secno} Import
                maps](webappapis.html#import-maps)
            3.  [[8.1.5.3]{.secno} Import map processing
                model](webappapis.html#import-map-processing-model)
        6.  [[8.1.6]{.secno} JavaScript specification host
            hooks](webappapis.html#javascript-specification-host-hooks)
            1.  [[8.1.6.1]{.secno}
                HostEnsureCanAddPrivateElement(`O`{.variable})](webappapis.html#the-hostensurecanaddprivateelement-implementation)
            2.  [[8.1.6.2]{.secno}
                HostEnsureCanCompileStrings(`realm`{.variable},
                `parameterStrings`{.variable}, `bodyString`{.variable},
                `codeString`{.variable}, `compilationType`{.variable},
                `parameterArgs`{.variable},
                `bodyArg`{.variable})](webappapis.html#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg))
            3.  [[8.1.6.3]{.secno}
                HostGetCodeForEval(`argument`{.variable})](webappapis.html#hostgetcodeforeval(argument))
            4.  [[8.1.6.4]{.secno}
                HostPromiseRejectionTracker(`promise`{.variable},
                `operation`{.variable})](webappapis.html#the-hostpromiserejectiontracker-implementation)
            5.  [[8.1.6.5]{.secno}
                HostSystemUTCEpochNanoseconds(`global`{.variable})](webappapis.html#hostsystemutcepochnanoseconds)
            6.  [[8.1.6.6]{.secno} Job-related host
                hooks](webappapis.html#integration-with-javascript-jobs)
                1.  [[8.1.6.6.1]{.secno}
                    HostCallJobCallback(`callback`{.variable},
                    `V`{.variable},
                    `argumentsList`{.variable})](webappapis.html#hostcalljobcallback)
                2.  [[8.1.6.6.2]{.secno}
                    HostEnqueueFinalizationRegistryCleanupJob(`finalizationRegistry`{.variable})](webappapis.html#hostenqueuefinalizationregistrycleanupjob)
                3.  [[8.1.6.6.3]{.secno}
                    HostEnqueueGenericJob(`job`{.variable},
                    `realm`{.variable})](webappapis.html#hostenqueuegenericjob)
                4.  [[8.1.6.6.4]{.secno}
                    HostEnqueuePromiseJob(`job`{.variable},
                    `realm`{.variable})](webappapis.html#hostenqueuepromisejob)
                5.  [[8.1.6.6.5]{.secno}
                    HostEnqueueTimeoutJob(`job`{.variable},
                    `realm`{.variable},
                    `milliseconds`{.variable})](webappapis.html#hostenqueuetimeoutjob)
                6.  [[8.1.6.6.6]{.secno}
                    HostMakeJobCallback(`callable`{.variable})](webappapis.html#hostmakejobcallback)
            7.  [[8.1.6.7]{.secno} Module-related host
                hooks](webappapis.html#integration-with-the-javascript-module-system)
                1.  [[8.1.6.7.1]{.secno}
                    HostGetImportMetaProperties(`moduleRecord`{.variable})](webappapis.html#hostgetimportmetaproperties)
                2.  [[8.1.6.7.2]{.secno}
                    HostGetSupportedImportAttributes()](webappapis.html#hostgetsupportedimportattributes)
                3.  [[8.1.6.7.3]{.secno}
                    HostLoadImportedModule(`referrer`{.variable},
                    `moduleRequest`{.variable}, `loadState`{.variable},
                    `payload`{.variable})](webappapis.html#hostloadimportedmodule)
        7.  [[8.1.7]{.secno} Event loops](webappapis.html#event-loops)
            1.  [[8.1.7.1]{.secno}
                Definitions](webappapis.html#definitions-3)
            2.  [[8.1.7.2]{.secno} Queuing
                tasks](webappapis.html#queuing-tasks)
            3.  [[8.1.7.3]{.secno} Processing
                model](webappapis.html#event-loop-processing-model)
            4.  [[8.1.7.4]{.secno} Generic task
                sources](webappapis.html#generic-task-sources)
            5.  [[8.1.7.5]{.secno} Dealing with the event loop from
                other
                specifications](webappapis.html#event-loop-for-spec-authors)
        8.  [[8.1.8]{.secno} Events](webappapis.html#events)
            1.  [[8.1.8.1]{.secno} Event
                handlers](webappapis.html#event-handler-attributes)
            2.  [[8.1.8.2]{.secno} Event handlers on elements,
                `Document` objects, and `Window`
                objects](webappapis.html#event-handlers-on-elements,-document-objects,-and-window-objects)
                1.  [[8.1.8.2.1]{.secno} IDL
                    definitions](webappapis.html#idl-definitions)
            3.  [[8.1.8.3]{.secno} Event
                firing](webappapis.html#event-firing)
    2.  [[8.2]{.secno} The `WindowOrWorkerGlobalScope`
        mixin](webappapis.html#windoworworkerglobalscope-mixin)
    3.  [[8.3]{.secno} Base64 utility methods](webappapis.html#atob)
    4.  [[8.4]{.secno} Dynamic markup
        insertion](dynamic-markup-insertion.html#dynamic-markup-insertion)
        1.  [[8.4.1]{.secno} Opening the input
            stream](dynamic-markup-insertion.html#opening-the-input-stream)
        2.  [[8.4.2]{.secno} Closing the input
            stream](dynamic-markup-insertion.html#closing-the-input-stream)
        3.  [[8.4.3]{.secno}
            `document.write()`](dynamic-markup-insertion.html#document.write())
        4.  [[8.4.4]{.secno}
            `document.writeln()`](dynamic-markup-insertion.html#document.writeln())
    5.  [[8.5]{.secno} DOM parsing and serialization
        APIs](dynamic-markup-insertion.html#dom-parsing-and-serialization)
        1.  [[8.5.1]{.secno} The `DOMParser`
            interface](dynamic-markup-insertion.html#the-domparser-interface)
        2.  [[8.5.2]{.secno} Unsafe HTML parsing
            methods](dynamic-markup-insertion.html#unsafe-html-parsing-methods)
        3.  [[8.5.3]{.secno} HTML serialization
            methods](dynamic-markup-insertion.html#html-serialization-methods)
        4.  [[8.5.4]{.secno} The `innerHTML`
            property](dynamic-markup-insertion.html#the-innerhtml-property)
        5.  [[8.5.5]{.secno} The `outerHTML`
            property](dynamic-markup-insertion.html#the-outerhtml-property)
        6.  [[8.5.6]{.secno} The `insertAdjacentHTML()`
            method](dynamic-markup-insertion.html#the-insertadjacenthtml()-method)
        7.  [[8.5.7]{.secno} The `createContextualFragment()`
            method](dynamic-markup-insertion.html#the-createcontextualfragment()-method)
        8.  [[8.5.8]{.secno} The `XMLSerializer`
            interface](dynamic-markup-insertion.html#the-xmlserializer-interface)
    6.  [[8.6]{.secno} Timers](timers-and-user-prompts.html#timers)
    7.  [[8.7]{.secno} Microtask
        queuing](timers-and-user-prompts.html#microtask-queuing)
    8.  [[8.8]{.secno} User
        prompts](timers-and-user-prompts.html#user-prompts)
        1.  [[8.8.1]{.secno} Simple
            dialogs](timers-and-user-prompts.html#simple-dialogs)
        2.  [[8.8.2]{.secno}
            Printing](timers-and-user-prompts.html#printing)
    9.  [[8.9]{.secno} System state and
        capabilities](system-state.html#system-state-and-capabilities)
        1.  [[8.9.1]{.secno} The `Navigator`
            object](system-state.html#the-navigator-object)
            1.  [[8.9.1.1]{.secno} Client
                identification](system-state.html#client-identification)
            2.  [[8.9.1.2]{.secno} Language
                preferences](system-state.html#language-preferences)
            3.  [[8.9.1.3]{.secno} Browser
                state](system-state.html#navigator.online)
            4.  [[8.9.1.4]{.secno} Custom scheme handlers: the
                `registerProtocolHandler()`
                method](system-state.html#custom-handlers)
                1.  [[8.9.1.4.1]{.secno} Security and
                    privacy](system-state.html#security-and-privacy)
                2.  [[8.9.1.4.2]{.secno} User agent
                    automation](system-state.html#user-agent-automation)
            5.  [[8.9.1.5]{.secno} Cookies](system-state.html#cookies)
            6.  [[8.9.1.6]{.secno} PDF viewing
                support](system-state.html#pdf-viewing-support)
    10. [[8.10]{.secno}
        Images](imagebitmap-and-animations.html#images-2)
        1.  [[8.10.1]{.secno} The `ImageData`
            interface](imagebitmap-and-animations.html#the-imagedata-interface)
        2.  [[8.10.2]{.secno} The `ImageBitmap`
            interface](imagebitmap-and-animations.html#the-imagebitmap-interface)
    11. [[8.11]{.secno} Animation
        frames](imagebitmap-and-animations.html#animation-frames)
9.  [[[9]{.secno} Communication](comms.html#comms)]{#toc-comms}
    1.  [[9.1]{.secno} The `MessageEvent`
        interface](comms.html#the-messageevent-interface)
    2.  [[9.2]{.secno} Server-sent
        events](server-sent-events.html#server-sent-events)
        1.  [[9.2.1]{.secno}
            Introduction](server-sent-events.html#server-sent-events-intro)
        2.  [[9.2.2]{.secno} The `EventSource`
            interface](server-sent-events.html#the-eventsource-interface)
        3.  [[9.2.3]{.secno} Processing
            model](server-sent-events.html#sse-processing-model)
        4.  [[9.2.4]{.secno} The \``Last-Event-ID`\`
            header](server-sent-events.html#the-last-event-id-header)
        5.  [[9.2.5]{.secno} Parsing an event
            stream](server-sent-events.html#parsing-an-event-stream)
        6.  [[9.2.6]{.secno} Interpreting an event
            stream](server-sent-events.html#event-stream-interpretation)
        7.  [[9.2.7]{.secno} Authoring
            notes](server-sent-events.html#authoring-notes)
        8.  [[9.2.8]{.secno} Connectionless push and other
            features](server-sent-events.html#eventsource-push)
        9.  [[9.2.9]{.secno} Garbage
            collection](server-sent-events.html#garbage-collection)
        10. [[9.2.10]{.secno} Implementation
            advice](server-sent-events.html#implementation-advice)
    3.  [[9.3]{.secno} Cross-document
        messaging](web-messaging.html#web-messaging)
        1.  [[9.3.1]{.secno}
            Introduction](web-messaging.html#introduction-12)
        2.  [[9.3.2]{.secno}
            Security](web-messaging.html#security-postmsg)
            1.  [[9.3.2.1]{.secno} Authors](web-messaging.html#authors)
            2.  [[9.3.2.2]{.secno} User
                agents](web-messaging.html#user-agents)
        3.  [[9.3.3]{.secno} Posting
            messages](web-messaging.html#posting-messages)
    4.  [[9.4]{.secno} Channel
        messaging](web-messaging.html#channel-messaging)
        1.  [[9.4.1]{.secno}
            Introduction](web-messaging.html#introduction-13)
            1.  [[9.4.1.1]{.secno}
                Examples](web-messaging.html#examples-5)
            2.  [[9.4.1.2]{.secno} Ports as the basis of an
                object-capability model on the
                web](web-messaging.html#ports-as-the-basis-of-an-object-capability-model-on-the-web)
            3.  [[9.4.1.3]{.secno} Ports as the basis of abstracting out
                service
                implementations](web-messaging.html#ports-as-the-basis-of-abstracting-out-service-implementations)
        2.  [[9.4.2]{.secno} Message
            channels](web-messaging.html#message-channels)
        3.  [[9.4.3]{.secno} The `MessageEventTarget`
            mixin](web-messaging.html#the-messageeventtarget-mixin)
        4.  [[9.4.4]{.secno} Message
            ports](web-messaging.html#message-ports)
        5.  [[9.4.5]{.secno} Ports and garbage
            collection](web-messaging.html#ports-and-garbage-collection)
    5.  [[9.5]{.secno} Broadcasting to other browsing
        contexts](web-messaging.html#broadcasting-to-other-browsing-contexts)
10. [[[10]{.secno} Web workers](workers.html#workers)]{#toc-workers}
    1.  [[10.1]{.secno} Introduction](workers.html#introduction-14)
        1.  [[10.1.1]{.secno} Scope](workers.html#scope-2)
        2.  [[10.1.2]{.secno} Examples](workers.html#examples-6)
            1.  [[10.1.2.1]{.secno} A background number-crunching
                worker](workers.html#a-background-number-crunching-worker)
            2.  [[10.1.2.2]{.secno} Using a JavaScript module as a
                worker](workers.html#module-worker-example)
            3.  [[10.1.2.3]{.secno} Shared workers
                introduction](workers.html#shared-workers-introduction)
            4.  [[10.1.2.4]{.secno} Shared state using a shared
                worker](workers.html#shared-state-using-a-shared-worker)
            5.  [[10.1.2.5]{.secno} Delegation](workers.html#delegation)
            6.  [[10.1.2.6]{.secno} Providing
                libraries](workers.html#providing-libraries)
        3.  [[10.1.3]{.secno} Tutorials](workers.html#tutorials)
            1.  [[10.1.3.1]{.secno} Creating a dedicated
                worker](workers.html#creating-a-dedicated-worker)
            2.  [[10.1.3.2]{.secno} Communicating with a dedicated
                worker](workers.html#communicating-with-a-dedicated-worker)
            3.  [[10.1.3.3]{.secno} Shared
                workers](workers.html#shared-workers)
    2.  [[10.2]{.secno} Infrastructure](workers.html#infrastructure-2)
        1.  [[10.2.1]{.secno} The global
            scope](workers.html#the-global-scope)
            1.  [[10.2.1.1]{.secno} The `WorkerGlobalScope` common
                interface](workers.html#the-workerglobalscope-common-interface)
            2.  [[10.2.1.2]{.secno} Dedicated workers and the
                `DedicatedWorkerGlobalScope`
                interface](workers.html#dedicated-workers-and-the-dedicatedworkerglobalscope-interface)
            3.  [[10.2.1.3]{.secno} Shared workers and the
                `SharedWorkerGlobalScope`
                interface](workers.html#shared-workers-and-the-sharedworkerglobalscope-interface)
        2.  [[10.2.2]{.secno} The event
            loop](workers.html#worker-event-loop)
        3.  [[10.2.3]{.secno} The worker\'s
            lifetime](workers.html#the-worker's-lifetime)
        4.  [[10.2.4]{.secno} Processing
            model](workers.html#worker-processing-model)
        5.  [[10.2.5]{.secno} Runtime script
            errors](workers.html#runtime-script-errors-2)
        6.  [[10.2.6]{.secno} Creating
            workers](workers.html#creating-workers)
            1.  [[10.2.6.1]{.secno} The `AbstractWorker`
                mixin](workers.html#the-abstractworker-mixin)
            2.  [[10.2.6.2]{.secno} Script settings for
                workers](workers.html#script-settings-for-workers)
            3.  [[10.2.6.3]{.secno} Dedicated workers and the `Worker`
                interface](workers.html#dedicated-workers-and-the-worker-interface)
            4.  [[10.2.6.4]{.secno} Shared workers and the
                `SharedWorker`
                interface](workers.html#shared-workers-and-the-sharedworker-interface)
        7.  [[10.2.7]{.secno} Concurrent hardware
            capabilities](workers.html#navigator.hardwareconcurrency)
    3.  [[10.3]{.secno} APIs available to
        workers](workers.html#apis-available-to-workers)
        1.  [[10.3.1]{.secno} Importing scripts and
            libraries](workers.html#importing-scripts-and-libraries)
        2.  [[10.3.2]{.secno} The `WorkerNavigator`
            interface](workers.html#the-workernavigator-object)
        3.  [[10.3.3]{.secno} The `WorkerLocation`
            interface](workers.html#worker-locations)
11. [[[11]{.secno} Worklets](worklets.html#worklets)]{#toc-worklets}
    1.  [[11.1]{.secno} Introduction](worklets.html#worklets-intro)
        1.  [[11.1.1]{.secno}
            Motivations](worklets.html#worklets-motivations)
        2.  [[11.1.2]{.secno} Code
            idempotence](worklets.html#worklets-idempotent)
        3.  [[11.1.3]{.secno} Speculative
            evaluation](worklets.html#worklets-speculative)
    2.  [[11.2]{.secno} Examples](worklets.html#worklets-examples)
        1.  [[11.2.1]{.secno} Loading
            scripts](worklets.html#worklets-examples-loading)
        2.  [[11.2.2]{.secno} Registering a class and invoking its
            methods](worklets.html#worklets-example-registering)
    3.  [[11.3]{.secno}
        Infrastructure](worklets.html#worklets-infrastructure)
        1.  [[11.3.1]{.secno} The global
            scope](worklets.html#worklets-global)
            1.  [[11.3.1.1]{.secno} Agents and event
                loops](worklets.html#worklet-agents-and-event-loops)
            2.  [[11.3.1.2]{.secno} Creation and
                termination](worklets.html#worklets-creation-termination)
            3.  [[11.3.1.3]{.secno} Script settings for
                worklets](worklets.html#script-settings-for-worklets)
        2.  [[11.3.2]{.secno} The `Worklet`
            class](worklets.html#worklets-worklet)
        3.  [[11.3.3]{.secno} The worklet\'s
            lifetime](worklets.html#worklets-lifetime)
12. [[[12]{.secno} Web
    storage](webstorage.html#webstorage)]{#toc-webstorage}
    1.  [[12.1]{.secno} Introduction](webstorage.html#introduction-15)
    2.  [[12.2]{.secno} The API](webstorage.html#storage)
        1.  [[12.2.1]{.secno} The `Storage`
            interface](webstorage.html#the-storage-interface)
        2.  [[12.2.2]{.secno} The `sessionStorage`
            getter](webstorage.html#the-sessionstorage-attribute)
        3.  [[12.2.3]{.secno} The `localStorage`
            getter](webstorage.html#the-localstorage-attribute)
        4.  [[12.2.4]{.secno} The `StorageEvent`
            interface](webstorage.html#the-storageevent-interface)
    3.  [[12.3]{.secno} Privacy](webstorage.html#privacy)
        1.  [[12.3.1]{.secno} User
            tracking](webstorage.html#user-tracking)
        2.  [[12.3.2]{.secno} Sensitivity of
            data](webstorage.html#sensitivity-of-data)
    4.  [[12.4]{.secno} Security](webstorage.html#security-storage)
        1.  [[12.4.1]{.secno} DNS spoofing
            attacks](webstorage.html#dns-spoofing-attacks)
        2.  [[12.4.2]{.secno} Cross-directory
            attacks](webstorage.html#cross-directory-attacks)
        3.  [[12.4.3]{.secno} Implementation
            risks](webstorage.html#implementation-risks)
13. [[[13]{.secno} The HTML syntax](syntax.html#syntax)]{#toc-syntax}
    1.  [[13.1]{.secno} Writing HTML documents](syntax.html#writing)
        1.  [[13.1.1]{.secno} The DOCTYPE](syntax.html#the-doctype)
        2.  [[13.1.2]{.secno} Elements](syntax.html#elements-2)
            1.  [[13.1.2.1]{.secno} Start tags](syntax.html#start-tags)
            2.  [[13.1.2.2]{.secno} End tags](syntax.html#end-tags)
            3.  [[13.1.2.3]{.secno}
                Attributes](syntax.html#attributes-2)
            4.  [[13.1.2.4]{.secno} Optional
                tags](syntax.html#optional-tags)
            5.  [[13.1.2.5]{.secno} Restrictions on content
                models](syntax.html#element-restrictions)
            6.  [[13.1.2.6]{.secno} Restrictions on the contents of raw
                text and escapable raw text
                elements](syntax.html#cdata-rcdata-restrictions)
        3.  [[13.1.3]{.secno} Text](syntax.html#text-2)
            1.  [[13.1.3.1]{.secno} Newlines](syntax.html#newlines)
        4.  [[13.1.4]{.secno} Character
            references](syntax.html#character-references)
        5.  [[13.1.5]{.secno} CDATA
            sections](syntax.html#cdata-sections)
        6.  [[13.1.6]{.secno} Comments](syntax.html#comments)
    2.  [[13.2]{.secno} Parsing HTML documents](parsing.html#parsing)
        1.  [[13.2.1]{.secno} Overview of the parsing
            model](parsing.html#overview-of-the-parsing-model)
        2.  [[13.2.2]{.secno} Parse errors](parsing.html#parse-errors)
        3.  [[13.2.3]{.secno} The input byte
            stream](parsing.html#the-input-byte-stream)
            1.  [[13.2.3.1]{.secno} Parsing with a known character
                encoding](parsing.html#parsing-with-a-known-character-encoding)
            2.  [[13.2.3.2]{.secno} Determining the character
                encoding](parsing.html#determining-the-character-encoding)
            3.  [[13.2.3.3]{.secno} Character
                encodings](parsing.html#character-encodings)
            4.  [[13.2.3.4]{.secno} Changing the encoding while
                parsing](parsing.html#changing-the-encoding-while-parsing)
            5.  [[13.2.3.5]{.secno} Preprocessing the input
                stream](parsing.html#preprocessing-the-input-stream)
        4.  [[13.2.4]{.secno} Parse state](parsing.html#parse-state)
            1.  [[13.2.4.1]{.secno} The insertion
                mode](parsing.html#the-insertion-mode)
            2.  [[13.2.4.2]{.secno} The stack of open
                elements](parsing.html#the-stack-of-open-elements)
            3.  [[13.2.4.3]{.secno} The list of active formatting
                elements](parsing.html#the-list-of-active-formatting-elements)
            4.  [[13.2.4.4]{.secno} The element
                pointers](parsing.html#the-element-pointers)
            5.  [[13.2.4.5]{.secno} Other parsing state
                flags](parsing.html#other-parsing-state-flags)
        5.  [[13.2.5]{.secno} Tokenization](parsing.html#tokenization)
            1.  [[13.2.5.1]{.secno} Data state](parsing.html#data-state)
            2.  [[13.2.5.2]{.secno} RCDATA
                state](parsing.html#rcdata-state)
            3.  [[13.2.5.3]{.secno} RAWTEXT
                state](parsing.html#rawtext-state)
            4.  [[13.2.5.4]{.secno} Script data
                state](parsing.html#script-data-state)
            5.  [[13.2.5.5]{.secno} PLAINTEXT
                state](parsing.html#plaintext-state)
            6.  [[13.2.5.6]{.secno} Tag open
                state](parsing.html#tag-open-state)
            7.  [[13.2.5.7]{.secno} End tag open
                state](parsing.html#end-tag-open-state)
            8.  [[13.2.5.8]{.secno} Tag name
                state](parsing.html#tag-name-state)
            9.  [[13.2.5.9]{.secno} RCDATA less-than sign
                state](parsing.html#rcdata-less-than-sign-state)
            10. [[13.2.5.10]{.secno} RCDATA end tag open
                state](parsing.html#rcdata-end-tag-open-state)
            11. [[13.2.5.11]{.secno} RCDATA end tag name
                state](parsing.html#rcdata-end-tag-name-state)
            12. [[13.2.5.12]{.secno} RAWTEXT less-than sign
                state](parsing.html#rawtext-less-than-sign-state)
            13. [[13.2.5.13]{.secno} RAWTEXT end tag open
                state](parsing.html#rawtext-end-tag-open-state)
            14. [[13.2.5.14]{.secno} RAWTEXT end tag name
                state](parsing.html#rawtext-end-tag-name-state)
            15. [[13.2.5.15]{.secno} Script data less-than sign
                state](parsing.html#script-data-less-than-sign-state)
            16. [[13.2.5.16]{.secno} Script data end tag open
                state](parsing.html#script-data-end-tag-open-state)
            17. [[13.2.5.17]{.secno} Script data end tag name
                state](parsing.html#script-data-end-tag-name-state)
            18. [[13.2.5.18]{.secno} Script data escape start
                state](parsing.html#script-data-escape-start-state)
            19. [[13.2.5.19]{.secno} Script data escape start dash
                state](parsing.html#script-data-escape-start-dash-state)
            20. [[13.2.5.20]{.secno} Script data escaped
                state](parsing.html#script-data-escaped-state)
            21. [[13.2.5.21]{.secno} Script data escaped dash
                state](parsing.html#script-data-escaped-dash-state)
            22. [[13.2.5.22]{.secno} Script data escaped dash dash
                state](parsing.html#script-data-escaped-dash-dash-state)
            23. [[13.2.5.23]{.secno} Script data escaped less-than sign
                state](parsing.html#script-data-escaped-less-than-sign-state)
            24. [[13.2.5.24]{.secno} Script data escaped end tag open
                state](parsing.html#script-data-escaped-end-tag-open-state)
            25. [[13.2.5.25]{.secno} Script data escaped end tag name
                state](parsing.html#script-data-escaped-end-tag-name-state)
            26. [[13.2.5.26]{.secno} Script data double escape start
                state](parsing.html#script-data-double-escape-start-state)
            27. [[13.2.5.27]{.secno} Script data double escaped
                state](parsing.html#script-data-double-escaped-state)
            28. [[13.2.5.28]{.secno} Script data double escaped dash
                state](parsing.html#script-data-double-escaped-dash-state)
            29. [[13.2.5.29]{.secno} Script data double escaped dash
                dash
                state](parsing.html#script-data-double-escaped-dash-dash-state)
            30. [[13.2.5.30]{.secno} Script data double escaped
                less-than sign
                state](parsing.html#script-data-double-escaped-less-than-sign-state)
            31. [[13.2.5.31]{.secno} Script data double escape end
                state](parsing.html#script-data-double-escape-end-state)
            32. [[13.2.5.32]{.secno} Before attribute name
                state](parsing.html#before-attribute-name-state)
            33. [[13.2.5.33]{.secno} Attribute name
                state](parsing.html#attribute-name-state)
            34. [[13.2.5.34]{.secno} After attribute name
                state](parsing.html#after-attribute-name-state)
            35. [[13.2.5.35]{.secno} Before attribute value
                state](parsing.html#before-attribute-value-state)
            36. [[13.2.5.36]{.secno} Attribute value (double-quoted)
                state](parsing.html#attribute-value-(double-quoted)-state)
            37. [[13.2.5.37]{.secno} Attribute value (single-quoted)
                state](parsing.html#attribute-value-(single-quoted)-state)
            38. [[13.2.5.38]{.secno} Attribute value (unquoted)
                state](parsing.html#attribute-value-(unquoted)-state)
            39. [[13.2.5.39]{.secno} After attribute value (quoted)
                state](parsing.html#after-attribute-value-(quoted)-state)
            40. [[13.2.5.40]{.secno} Self-closing start tag
                state](parsing.html#self-closing-start-tag-state)
            41. [[13.2.5.41]{.secno} Bogus comment
                state](parsing.html#bogus-comment-state)
            42. [[13.2.5.42]{.secno} Markup declaration open
                state](parsing.html#markup-declaration-open-state)
            43. [[13.2.5.43]{.secno} Comment start
                state](parsing.html#comment-start-state)
            44. [[13.2.5.44]{.secno} Comment start dash
                state](parsing.html#comment-start-dash-state)
            45. [[13.2.5.45]{.secno} Comment
                state](parsing.html#comment-state)
            46. [[13.2.5.46]{.secno} Comment less-than sign
                state](parsing.html#comment-less-than-sign-state)
            47. [[13.2.5.47]{.secno} Comment less-than sign bang
                state](parsing.html#comment-less-than-sign-bang-state)
            48. [[13.2.5.48]{.secno} Comment less-than sign bang dash
                state](parsing.html#comment-less-than-sign-bang-dash-state)
            49. [[13.2.5.49]{.secno} Comment less-than sign bang dash
                dash
                state](parsing.html#comment-less-than-sign-bang-dash-dash-state)
            50. [[13.2.5.50]{.secno} Comment end dash
                state](parsing.html#comment-end-dash-state)
            51. [[13.2.5.51]{.secno} Comment end
                state](parsing.html#comment-end-state)
            52. [[13.2.5.52]{.secno} Comment end bang
                state](parsing.html#comment-end-bang-state)
            53. [[13.2.5.53]{.secno} DOCTYPE
                state](parsing.html#doctype-state)
            54. [[13.2.5.54]{.secno} Before DOCTYPE name
                state](parsing.html#before-doctype-name-state)
            55. [[13.2.5.55]{.secno} DOCTYPE name
                state](parsing.html#doctype-name-state)
            56. [[13.2.5.56]{.secno} After DOCTYPE name
                state](parsing.html#after-doctype-name-state)
            57. [[13.2.5.57]{.secno} After DOCTYPE public keyword
                state](parsing.html#after-doctype-public-keyword-state)
            58. [[13.2.5.58]{.secno} Before DOCTYPE public identifier
                state](parsing.html#before-doctype-public-identifier-state)
            59. [[13.2.5.59]{.secno} DOCTYPE public identifier
                (double-quoted)
                state](parsing.html#doctype-public-identifier-(double-quoted)-state)
            60. [[13.2.5.60]{.secno} DOCTYPE public identifier
                (single-quoted)
                state](parsing.html#doctype-public-identifier-(single-quoted)-state)
            61. [[13.2.5.61]{.secno} After DOCTYPE public identifier
                state](parsing.html#after-doctype-public-identifier-state)
            62. [[13.2.5.62]{.secno} Between DOCTYPE public and system
                identifiers
                state](parsing.html#between-doctype-public-and-system-identifiers-state)
            63. [[13.2.5.63]{.secno} After DOCTYPE system keyword
                state](parsing.html#after-doctype-system-keyword-state)
            64. [[13.2.5.64]{.secno} Before DOCTYPE system identifier
                state](parsing.html#before-doctype-system-identifier-state)
            65. [[13.2.5.65]{.secno} DOCTYPE system identifier
                (double-quoted)
                state](parsing.html#doctype-system-identifier-(double-quoted)-state)
            66. [[13.2.5.66]{.secno} DOCTYPE system identifier
                (single-quoted)
                state](parsing.html#doctype-system-identifier-(single-quoted)-state)
            67. [[13.2.5.67]{.secno} After DOCTYPE system identifier
                state](parsing.html#after-doctype-system-identifier-state)
            68. [[13.2.5.68]{.secno} Bogus DOCTYPE
                state](parsing.html#bogus-doctype-state)
            69. [[13.2.5.69]{.secno} CDATA section
                state](parsing.html#cdata-section-state)
            70. [[13.2.5.70]{.secno} CDATA section bracket
                state](parsing.html#cdata-section-bracket-state)
            71. [[13.2.5.71]{.secno} CDATA section end
                state](parsing.html#cdata-section-end-state)
            72. [[13.2.5.72]{.secno} Character reference
                state](parsing.html#character-reference-state)
            73. [[13.2.5.73]{.secno} Named character reference
                state](parsing.html#named-character-reference-state)
            74. [[13.2.5.74]{.secno} Ambiguous ampersand
                state](parsing.html#ambiguous-ampersand-state)
            75. [[13.2.5.75]{.secno} Numeric character reference
                state](parsing.html#numeric-character-reference-state)
            76. [[13.2.5.76]{.secno} Hexadecimal character reference
                start
                state](parsing.html#hexadecimal-character-reference-start-state)
            77. [[13.2.5.77]{.secno} Decimal character reference start
                state](parsing.html#decimal-character-reference-start-state)
            78. [[13.2.5.78]{.secno} Hexadecimal character reference
                state](parsing.html#hexadecimal-character-reference-state)
            79. [[13.2.5.79]{.secno} Decimal character reference
                state](parsing.html#decimal-character-reference-state)
            80. [[13.2.5.80]{.secno} Numeric character reference end
                state](parsing.html#numeric-character-reference-end-state)
        6.  [[13.2.6]{.secno} Tree
            construction](parsing.html#tree-construction)
            1.  [[13.2.6.1]{.secno} Creating and inserting
                nodes](parsing.html#creating-and-inserting-nodes)
            2.  [[13.2.6.2]{.secno} Parsing elements that contain only
                text](parsing.html#parsing-elements-that-contain-only-text)
            3.  [[13.2.6.3]{.secno} Closing elements that have implied
                end
                tags](parsing.html#closing-elements-that-have-implied-end-tags)
            4.  [[13.2.6.4]{.secno} The rules for parsing tokens in HTML
                content](parsing.html#parsing-main-inhtml)
                1.  [[13.2.6.4.1]{.secno} The \"initial\" insertion
                    mode](parsing.html#the-initial-insertion-mode)
                2.  [[13.2.6.4.2]{.secno} The \"before html\" insertion
                    mode](parsing.html#the-before-html-insertion-mode)
                3.  [[13.2.6.4.3]{.secno} The \"before head\" insertion
                    mode](parsing.html#the-before-head-insertion-mode)
                4.  [[13.2.6.4.4]{.secno} The \"in head\" insertion
                    mode](parsing.html#parsing-main-inhead)
                5.  [[13.2.6.4.5]{.secno} The \"in head noscript\"
                    insertion
                    mode](parsing.html#parsing-main-inheadnoscript)
                6.  [[13.2.6.4.6]{.secno} The \"after head\" insertion
                    mode](parsing.html#the-after-head-insertion-mode)
                7.  [[13.2.6.4.7]{.secno} The \"in body\" insertion
                    mode](parsing.html#parsing-main-inbody)
                8.  [[13.2.6.4.8]{.secno} The \"text\" insertion
                    mode](parsing.html#parsing-main-incdata)
                9.  [[13.2.6.4.9]{.secno} The \"in table\" insertion
                    mode](parsing.html#parsing-main-intable)
                10. [[13.2.6.4.10]{.secno} The \"in table text\"
                    insertion
                    mode](parsing.html#parsing-main-intabletext)
                11. [[13.2.6.4.11]{.secno} The \"in caption\" insertion
                    mode](parsing.html#parsing-main-incaption)
                12. [[13.2.6.4.12]{.secno} The \"in column group\"
                    insertion
                    mode](parsing.html#parsing-main-incolgroup)
                13. [[13.2.6.4.13]{.secno} The \"in table body\"
                    insertion mode](parsing.html#parsing-main-intbody)
                14. [[13.2.6.4.14]{.secno} The \"in row\" insertion
                    mode](parsing.html#parsing-main-intr)
                15. [[13.2.6.4.15]{.secno} The \"in cell\" insertion
                    mode](parsing.html#parsing-main-intd)
                16. [[13.2.6.4.16]{.secno} The \"in select\" insertion
                    mode](parsing.html#parsing-main-inselect)
                17. [[13.2.6.4.17]{.secno} The \"in select in table\"
                    insertion
                    mode](parsing.html#parsing-main-inselectintable)
                18. [[13.2.6.4.18]{.secno} The \"in template\" insertion
                    mode](parsing.html#parsing-main-intemplate)
                19. [[13.2.6.4.19]{.secno} The \"after body\" insertion
                    mode](parsing.html#parsing-main-afterbody)
                20. [[13.2.6.4.20]{.secno} The \"in frameset\" insertion
                    mode](parsing.html#parsing-main-inframeset)
                21. [[13.2.6.4.21]{.secno} The \"after frameset\"
                    insertion
                    mode](parsing.html#parsing-main-afterframeset)
                22. [[13.2.6.4.22]{.secno} The \"after after body\"
                    insertion
                    mode](parsing.html#the-after-after-body-insertion-mode)
                23. [[13.2.6.4.23]{.secno} The \"after after frameset\"
                    insertion
                    mode](parsing.html#the-after-after-frameset-insertion-mode)
            5.  [[13.2.6.5]{.secno} The rules for parsing tokens in
                foreign content](parsing.html#parsing-main-inforeign)
        7.  [[13.2.7]{.secno} The end](parsing.html#the-end)
        8.  [[13.2.8]{.secno} Speculative HTML
            parsing](parsing.html#speculative-html-parsing)
        9.  [[13.2.9]{.secno} Coercing an HTML DOM into an
            infoset](parsing.html#coercing-an-html-dom-into-an-infoset)
        10. [[13.2.10]{.secno} An introduction to error handling and
            strange cases in the
            parser](parsing.html#an-introduction-to-error-handling-and-strange-cases-in-the-parser)
            1.  [[13.2.10.1]{.secno} Misnested tags:
                \<b\>\<i\>\</b\>\</i\>](parsing.html#misnested-tags:-b-i-/b-/i)
            2.  [[13.2.10.2]{.secno} Misnested tags:
                \<b\>\<p\>\</b\>\</p\>](parsing.html#misnested-tags:-b-p-/b-/p)
            3.  [[13.2.10.3]{.secno} Unexpected markup in
                tables](parsing.html#unexpected-markup-in-tables)
            4.  [[13.2.10.4]{.secno} Scripts that modify the page as it
                is being
                parsed](parsing.html#scripts-that-modify-the-page-as-it-is-being-parsed)
            5.  [[13.2.10.5]{.secno} The execution of scripts that are
                moving across multiple
                documents](parsing.html#the-execution-of-scripts-that-are-moving-across-multiple-documents)
            6.  [[13.2.10.6]{.secno} Unclosed formatting
                elements](parsing.html#unclosed-formatting-elements)
    3.  [[13.3]{.secno} Serializing HTML
        fragments](parsing.html#serialising-html-fragments)
    4.  [[13.4]{.secno} Parsing HTML
        fragments](parsing.html#parsing-html-fragments)
    5.  [[13.5]{.secno} Named character
        references](named-characters.html#named-character-references)
14. [[[14]{.secno} The XML
    syntax](xhtml.html#the-xhtml-syntax)]{#toc-the-xhtml-syntax}
    1.  [[14.1]{.secno} Writing documents in the XML
        syntax](xhtml.html#writing-xhtml-documents)
    2.  [[14.2]{.secno} Parsing XML
        documents](xhtml.html#parsing-xhtml-documents)
    3.  [[14.3]{.secno} Serializing XML
        fragments](xhtml.html#serialising-xhtml-fragments)
    4.  [[14.4]{.secno} Parsing XML
        fragments](xhtml.html#parsing-xhtml-fragments)
15. [[[15]{.secno} Rendering](rendering.html#rendering)]{#toc-rendering}
    1.  [[15.1]{.secno} Introduction](rendering.html#introduction-16)
    2.  [[15.2]{.secno} The CSS user agent style sheet and
        presentational
        hints](rendering.html#the-css-user-agent-style-sheet-and-presentational-hints)
    3.  [[15.3]{.secno} Non-replaced
        elements](rendering.html#non-replaced-elements)
        1.  [[15.3.1]{.secno} Hidden
            elements](rendering.html#hidden-elements)
        2.  [[15.3.2]{.secno} The page](rendering.html#the-page)
        3.  [[15.3.3]{.secno} Flow
            content](rendering.html#flow-content-3)
        4.  [[15.3.4]{.secno} Phrasing
            content](rendering.html#phrasing-content-3)
        5.  [[15.3.5]{.secno} Bidirectional
            text](rendering.html#bidi-rendering)
        6.  [[15.3.6]{.secno} Sections and
            headings](rendering.html#sections-and-headings)
        7.  [[15.3.7]{.secno} Lists](rendering.html#lists)
        8.  [[15.3.8]{.secno} Tables](rendering.html#tables-2)
        9.  [[15.3.9]{.secno} Margin collapsing
            quirks](rendering.html#margin-collapsing-quirks)
        10. [[15.3.10]{.secno} Form
            controls](rendering.html#form-controls)
        11. [[15.3.11]{.secno} The `hr`
            element](rendering.html#the-hr-element-2)
        12. [[15.3.12]{.secno} The `fieldset` and `legend`
            elements](rendering.html#the-fieldset-and-legend-elements)
    4.  [[15.4]{.secno} Replaced
        elements](rendering.html#replaced-elements)
        1.  [[15.4.1]{.secno} Embedded
            content](rendering.html#embedded-content-rendering-rules)
        2.  [[15.4.2]{.secno} Images](rendering.html#images-3)
        3.  [[15.4.3]{.secno} Attributes for embedded content and
            images](rendering.html#attributes-for-embedded-content-and-images)
        4.  [[15.4.4]{.secno} Image maps](rendering.html#image-maps-2)
    5.  [[15.5]{.secno} Widgets](rendering.html#widgets)
        1.  [[15.5.1]{.secno} Native
            appearance](rendering.html#native-appearance-2)
        2.  [[15.5.2]{.secno} Writing mode](rendering.html#writing-mode)
        3.  [[15.5.3]{.secno} Button
            layout](rendering.html#button-layout)
        4.  [[15.5.4]{.secno} The `button`
            element](rendering.html#the-button-element-2)
        5.  [[15.5.5]{.secno} The `details` and `summary`
            elements](rendering.html#the-details-and-summary-elements)
        6.  [[15.5.6]{.secno} The `input` element as a text entry
            widget](rendering.html#the-input-element-as-a-text-entry-widget)
        7.  [[15.5.7]{.secno} The `input` element as domain-specific
            widgets](rendering.html#the-input-element-as-domain-specific-widgets)
        8.  [[15.5.8]{.secno} The `input` element as a range
            control](rendering.html#the-input-element-as-a-range-control)
        9.  [[15.5.9]{.secno} The `input` element as a color
            well](rendering.html#the-input-element-as-a-colour-well)
        10. [[15.5.10]{.secno} The `input` element as a checkbox and
            radio button
            widgets](rendering.html#the-input-element-as-a-checkbox-and-radio-button-widgets)
        11. [[15.5.11]{.secno} The `input` element as a file upload
            control](rendering.html#the-input-element-as-a-file-upload-control)
        12. [[15.5.12]{.secno} The `input` element as a
            button](rendering.html#the-input-element-as-a-button)
        13. [[15.5.13]{.secno} The `marquee`
            element](rendering.html#the-marquee-element-2)
        14. [[15.5.14]{.secno} The `meter`
            element](rendering.html#the-meter-element-2)
        15. [[15.5.15]{.secno} The `progress`
            element](rendering.html#the-progress-element-2)
        16. [[15.5.16]{.secno} The `select`
            element](rendering.html#the-select-element-2)
        17. [[15.5.17]{.secno} The `textarea`
            element](rendering.html#the-textarea-element-2)
    6.  [[15.6]{.secno} Frames and
        framesets](rendering.html#frames-and-framesets)
    7.  [[15.7]{.secno} Interactive
        media](rendering.html#interactive-media)
        1.  [[15.7.1]{.secno} Links, forms, and
            navigation](rendering.html#links,-forms,-and-navigation)
        2.  [[15.7.2]{.secno} The `title`
            attribute](rendering.html#the-title-attribute-2)
        3.  [[15.7.3]{.secno} Editing
            hosts](rendering.html#editing-hosts)
        4.  [[15.7.4]{.secno} Text rendered in native user
            interfaces](rendering.html#text-rendered-in-native-user-interfaces)
    8.  [[15.8]{.secno} Print media](rendering.html#print-media)
    9.  [[15.9]{.secno} Unstyled XML
        documents](rendering.html#unstyled-xml-documents)
16. [[[16]{.secno} Obsolete
    features](obsolete.html#obsolete)]{#toc-obsolete}
    1.  [[16.1]{.secno} Obsolete but conforming
        features](obsolete.html#obsolete-but-conforming-features)
        1.  [[16.1.1]{.secno} Warnings for obsolete but conforming
            features](obsolete.html#warnings-for-obsolete-but-conforming-features)
    2.  [[16.2]{.secno} Non-conforming
        features](obsolete.html#non-conforming-features)
    3.  [[16.3]{.secno} Requirements for
        implementations](obsolete.html#requirements-for-implementations)
        1.  [[16.3.1]{.secno} The `marquee`
            element](obsolete.html#the-marquee-element)
        2.  [[16.3.2]{.secno} Frames](obsolete.html#frames)
        3.  [[16.3.3]{.secno} Other elements, attributes and
            APIs](obsolete.html#other-elements,-attributes-and-apis)
17. [[[17]{.secno} IANA considerations](iana.html#iana)]{#toc-iana}
    1.  [[17.1]{.secno} `text/html`](iana.html#text/html)
    2.  [[17.2]{.secno}
        `multipart/x-mixed-replace`](iana.html#multipart/x-mixed-replace)
    3.  [[17.3]{.secno}
        `application/xhtml+xml`](iana.html#application/xhtml+xml)
    4.  [[17.4]{.secno} `text/ping`](iana.html#text/ping)
    5.  [[17.5]{.secno}
        `application/microdata+json`](iana.html#application/microdata+json)
    6.  [[17.6]{.secno}
        `text/event-stream`](iana.html#text/event-stream)
    7.  [[17.7]{.secno} `web+` scheme
        prefix](iana.html#web+-scheme-prefix)
18. [[Index](indices.html#index)]{#toc-index}
    1.  [Elements](indices.html#elements-3)
    2.  [Element content
        categories](indices.html#element-content-categories)
    3.  [Attributes](indices.html#attributes-3)
    4.  [Element interfaces](indices.html#element-interfaces)
    5.  [All interfaces](indices.html#all-interfaces)
    6.  [Events](indices.html#events-2)
    7.  [HTTP headers](indices.html#http-headers)
    8.  [MIME types](indices.html#mime-types-2)
19. [[References](references.html#references)]{#toc-references}
20. [[Acknowledgments](acknowledgements.html#acknowledgments)]{#toc-acknowledgments}
21. [[Intellectual property
    rights](acknowledgements.html#ipr)]{#toc-ipr}
