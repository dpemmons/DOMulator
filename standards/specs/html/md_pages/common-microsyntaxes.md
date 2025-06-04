HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 2 Common infrastructure](infrastructure.html) --- [Table of
Contents](./) --- [2.4 URLs →](urls-and-fetching.html)

1.  ::: {#toc-infrastructure}
    1.  [[2.3]{.secno} Common
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
    :::

### [2.3]{.secno} Common microsyntaxes[](#common-microsyntaxes){.self-link}

There are various places in HTML that accept particular data types, such
as dates or numbers. This section describes what the conformance
criteria for content in those formats is, and how to parse them.

Implementers are strongly urged to carefully examine any third-party
libraries they might consider using to implement the parsing of syntaxes
described below. For example, date libraries are likely to implement
error handling behavior that differs from what is required in this
specification, since error-handling behavior is often not defined in
specifications that describe date syntaxes similar to those used in this
specification, and thus implementations tend to vary greatly in how they
handle errors.

#### [2.3.1]{.secno} Common parser idioms[](#common-parser-idioms){.self-link}

Some of the micro-parsers described below follow the pattern of having
an `input`{.variable} variable that holds the string being parsed, and
having a `position`{.variable} variable pointing at the next character
to parse in `input`{.variable}.

#### [2.3.2]{.secno} Boolean attributes[](#boolean-attributes){.self-link}

A number of attributes are [boolean attributes]{#boolean-attribute
.dfn}. The presence of a boolean attribute on an element represents the
true value, and the absence of the attribute represents the false value.

If the attribute is present, its value must either be the empty string
or a value that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#boolean-attributes:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the attribute\'s
canonical name, with no leading or trailing whitespace.

The values \"true\" and \"false\" are not allowed on boolean attributes.
To represent a false value, the attribute has to be omitted altogether.

::: example
Here is an example of a checkbox that is checked and disabled. The
[`checked`{#boolean-attributes:attr-input-checked}](input.html#attr-input-checked)
and
[`disabled`{#boolean-attributes:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
attributes are the boolean attributes.

``` html
<label><input type=checkbox checked name=cheese disabled> Cheese</label>
```

This could be equivalently written as this:

``` html
<label><input type=checkbox checked=checked name=cheese disabled=disabled> Cheese</label>
```

You can also mix styles; the following is still equivalent:

``` html
<label><input type='checkbox' checked name=cheese disabled=""> Cheese</label>
```
:::

#### [2.3.3]{.secno} Keywords and enumerated attributes[](#keywords-and-enumerated-attributes){.self-link}

Some attributes, called [enumerated attributes]{#enumerated-attribute
.dfn lt="enumerated attribute" export=""}, take on a finite set of
states. The state for such an attribute is derived by combining the
attribute\'s value, a set of keyword/state mappings given in the
specification of each attribute, and two possible special states that
can also be given in the specification of the attribute. These special
states are the [*invalid value default*]{#invalid-value-default .dfn}
and the [*missing value default*]{#missing-value-default .dfn}.

Multiple keywords can map to the same state.

The empty string can be a valid keyword. Note that the *[missing value
default](#missing-value-default)* applies only when the attribute is
*missing*, not when it is present with an empty string value.

To determine the state of an attribute, use the following steps:

1.  If the attribute is not specified:

    1.  If the attribute has a *[missing value
        default](#missing-value-default)* state defined, then return
        that *[missing value default](#missing-value-default)* state.

    2.  Otherwise, return no state.

2.  If the attribute\'s value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#keywords-and-enumerated-attributes:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for one of the keywords
    defined for the attribute, then return the state represented by that
    keyword.

3.  If the attribute has an *[invalid value
    default](#invalid-value-default)* state defined, then return that
    *[invalid value default](#invalid-value-default)* state.

4.  Return no state.

For authoring conformance purposes, if an enumerated attribute is
specified, the attribute\'s value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#keywords-and-enumerated-attributes:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for one of the conforming
keywords for that attribute, with no leading or trailing whitespace.

For
[reflection](common-dom-interfaces.html#reflect){#keywords-and-enumerated-attributes:reflect}
purposes, states which have any keywords mapping to them are said to
have a [canonical keyword]{#canonical-keyword .dfn}. This is determined
as follows:

- If there is only one keyword mapping to the given state, then it is
  that keyword.

- If there is only one *conforming* keyword mapping to the given state,
  then it is that conforming keyword.

- If there are two conforming keywords mapping to the given state, and
  one is the empty string, then the canonical keyword will be the
  conforming keyword that is not the empty string.

- Otherwise, the canonical keyword for the state will be explicitly
  given in the specification for the attribute.

#### [2.3.4]{.secno} Numbers[](#numbers){.self-link}

##### [2.3.4.1]{.secno} Signed integers[](#signed-integers){.self-link}

A string is a [valid integer]{#valid-integer .dfn} if it consists of one
or more [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#signed-integers:ascii-digits
x-internal="ascii-digits"}, optionally prefixed with a U+002D
HYPHEN-MINUS character (-).

A [valid integer](#valid-integer){#signed-integers:valid-integer}
without a U+002D HYPHEN-MINUS (-) prefix represents the number that is
represented in base ten by that string of digits. A [valid
integer](#valid-integer){#signed-integers:valid-integer-2} *with* a
U+002D HYPHEN-MINUS (-) prefix represents the number represented in base
ten by the string of digits that follows the U+002D HYPHEN-MINUS,
subtracted from zero.

The [rules for parsing integers]{#rules-for-parsing-integers .dfn} are
as given in the following algorithm. When invoked, the steps must be
followed in the order given, aborting at the first step that returns a
value. This algorithm will return either an integer or an error.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Let `sign`{.variable} have the value \"positive\".

4.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#signed-integers:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

5.  If `position`{.variable} is past the end of `input`{.variable},
    return an error.

6.  If the character indicated by `position`{.variable} (the first
    character) is a U+002D HYPHEN-MINUS character (-):

    1.  Let `sign`{.variable} be \"negative\".
    2.  Advance `position`{.variable} to the next character.
    3.  If `position`{.variable} is past the end of `input`{.variable},
        return an error.

    Otherwise, if the character indicated by `position`{.variable} (the
    first character) is a U+002B PLUS SIGN character (+):

    1.  Advance `position`{.variable} to the next character. (The
        \"`+`\" is ignored, but it is not conforming.)
    2.  If `position`{.variable} is past the end of `input`{.variable},
        return an error.

7.  If the character indicated by `position`{.variable} is not an [ASCII
    digit](https://infra.spec.whatwg.org/#ascii-digit){#signed-integers:ascii-digits-2
    x-internal="ascii-digits"}, then return an error.

8.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#signed-integers:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#signed-integers:ascii-digits-3
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}, and interpret the resulting sequence as a
    base-ten integer. Let `value`{.variable} be that integer.

9.  If `sign`{.variable} is \"positive\", return `value`{.variable},
    otherwise return the result of subtracting `value`{.variable} from
    zero.

##### [2.3.4.2]{.secno} Non-negative integers[](#non-negative-integers){.self-link}

A string is a [valid non-negative integer]{#valid-non-negative-integer
.dfn} if it consists of one or more [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#non-negative-integers:ascii-digits
x-internal="ascii-digits"}.

A [valid non-negative
integer](#valid-non-negative-integer){#non-negative-integers:valid-non-negative-integer}
represents the number that is represented in base ten by that string of
digits.

The [rules for parsing non-negative
integers]{#rules-for-parsing-non-negative-integers .dfn} are as given in
the following algorithm. When invoked, the steps must be followed in the
order given, aborting at the first step that returns a value. This
algorithm will return either zero, a positive integer, or an error.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `value`{.variable} be the result of parsing `input`{.variable}
    using the [rules for parsing
    integers](#rules-for-parsing-integers){#non-negative-integers:rules-for-parsing-integers}.

3.  If `value`{.variable} is an error, return an error.

4.  If `value`{.variable} is less than zero, return an error.

5.  Return `value`{.variable}.

##### [2.3.4.3]{.secno} Floating-point numbers[](#floating-point-numbers){.self-link}

A string is a [valid floating-point number]{#valid-floating-point-number
.dfn} if it consists of:

1.  Optionally, a U+002D HYPHEN-MINUS character (-).

2.  One or both of the following, in the given order:

    1.  A series of one or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits
        x-internal="ascii-digits"}.

    2.  Both of the following, in the given order:

        1.  A single U+002E FULL STOP character (.).

        2.  A series of one or more [ASCII
            digits](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-2
            x-internal="ascii-digits"}.

3.  Optionally:

    1.  Either a U+0065 LATIN SMALL LETTER E character (e) or a U+0045
        LATIN CAPITAL LETTER E character (E).

    2.  Optionally, a U+002D HYPHEN-MINUS character (-) or U+002B PLUS
        SIGN character (+).

    3.  A series of one or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-3
        x-internal="ascii-digits"}.

A [valid floating-point
number](#valid-floating-point-number){#floating-point-numbers:valid-floating-point-number}
represents the number obtained by multiplying the significand by ten
raised to the power of the exponent, where the significand is the first
number, interpreted as base ten (including the decimal point and the
number after the decimal point, if any, and interpreting the significand
as a negative number if the whole string starts with a U+002D
HYPHEN-MINUS character (-) and the number is not zero), and where the
exponent is the number after the E, if any (interpreted as a negative
number if there is a U+002D HYPHEN-MINUS character (-) between the E and
the number and the number is not zero, or else ignoring a U+002B PLUS
SIGN character (+) between the E and the number if there is one). If
there is no E, then the exponent is treated as zero.

The Infinity and Not-a-Number (NaN) values are not [valid floating-point
numbers](#valid-floating-point-number){#floating-point-numbers:valid-floating-point-number-2}.

The [valid floating-point
number](#valid-floating-point-number){#floating-point-numbers:valid-floating-point-number-3}
concept is typically only used to restrict what is allowed for authors,
while the user agent requirements use the [rules for parsing
floating-point number
values](#rules-for-parsing-floating-point-number-values){#floating-point-numbers:rules-for-parsing-floating-point-number-values}
below (e.g., the
[`max`{#floating-point-numbers:attr-progress-max}](form-elements.html#attr-progress-max)
attribute of the
[`progress`{#floating-point-numbers:the-progress-element}](form-elements.html#the-progress-element)
element). However, in some cases the user agent requirements include
checking if a string is a [valid floating-point
number](#valid-floating-point-number){#floating-point-numbers:valid-floating-point-number-4}
(e.g., the [value sanitization
algorithm](input.html#value-sanitization-algorithm){#floating-point-numbers:value-sanitization-algorithm}
for the
[Number](input.html#number-state-(type=number)){#floating-point-numbers:number-state-(type=number)}
state of the
[`input`{#floating-point-numbers:the-input-element}](input.html#the-input-element)
element, or the [parse a srcset
attribute](images.html#parse-a-srcset-attribute){#floating-point-numbers:parse-a-srcset-attribute}
algorithm).

The [best representation of the number `n`{.variable} as a
floating-point
number]{#best-representation-of-the-number-as-a-floating-point-number
.dfn} is the string obtained from running
[ToString](https://tc39.es/ecma262/#sec-tostring){#floating-point-numbers:tostring
x-internal="tostring"}(`n`{.variable}). The abstract operation
[ToString](https://tc39.es/ecma262/#sec-tostring){#floating-point-numbers:tostring-2
x-internal="tostring"} is not uniquely determined. When there are
multiple possible strings that could be obtained from
[ToString](https://tc39.es/ecma262/#sec-tostring){#floating-point-numbers:tostring-3
x-internal="tostring"} for a particular value, the user agent must
always return the same string for that value (though it may differ from
the value used by other user agents).

The [rules for parsing floating-point number
values]{#rules-for-parsing-floating-point-number-values .dfn export=""}
are as given in the following algorithm. This algorithm must be aborted
at the first step that returns something. This algorithm will return
either a number or an error.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Let `value`{.variable} have the value 1.

4.  Let `divisor`{.variable} have the value 1.

5.  Let `exponent`{.variable} have the value 1.

6.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#floating-point-numbers:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

7.  If `position`{.variable} is past the end of `input`{.variable},
    return an error.

8.  If the character indicated by `position`{.variable} is a U+002D
    HYPHEN-MINUS character (-):

    1.  Change `value`{.variable} and `divisor`{.variable} to −1.
    2.  Advance `position`{.variable} to the next character.
    3.  If `position`{.variable} is past the end of `input`{.variable},
        return an error.

    Otherwise, if the character indicated by `position`{.variable} (the
    first character) is a U+002B PLUS SIGN character (+):

    1.  Advance `position`{.variable} to the next character. (The
        \"`+`\" is ignored, but it is not conforming.)
    2.  If `position`{.variable} is past the end of `input`{.variable},
        return an error.

9.  If the character indicated by `position`{.variable} is a U+002E FULL
    STOP (.), and that is not the last character in `input`{.variable},
    and the character after the character indicated by
    `position`{.variable} is an [ASCII
    digit](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-4
    x-internal="ascii-digits"}, then set `value`{.variable} to zero and
    jump to the step labeled *fraction*.

10. If the character indicated by `position`{.variable} is not an [ASCII
    digit](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-5
    x-internal="ascii-digits"}, then return an error.

11. [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#floating-point-numbers:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-6
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}, and interpret the resulting sequence as a
    base-ten integer. Multiply `value`{.variable} by that integer.

12. If `position`{.variable} is past the end of `input`{.variable}, jump
    to the step labeled *conversion*.

13. *Fraction*: If the character indicated by `position`{.variable} is a
    U+002E FULL STOP (.), run these substeps:

    1.  Advance `position`{.variable} to the next character.

    2.  If `position`{.variable} is past the end of `input`{.variable},
        or if the character indicated by `position`{.variable} is not an
        [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-7
        x-internal="ascii-digits"}, U+0065 LATIN SMALL LETTER E (e), or
        U+0045 LATIN CAPITAL LETTER E (E), then jump to the step labeled
        *conversion*.

    3.  If the character indicated by `position`{.variable} is a U+0065
        LATIN SMALL LETTER E character (e) or a U+0045 LATIN CAPITAL
        LETTER E character (E), skip the remainder of these substeps.

    4.  *Fraction loop*: Multiply `divisor`{.variable} by ten.

    5.  Add the value of the character indicated by
        `position`{.variable}, interpreted as a base-ten digit (0..9)
        and divided by `divisor`{.variable}, to `value`{.variable}.

    6.  Advance `position`{.variable} to the next character.

    7.  If `position`{.variable} is past the end of `input`{.variable},
        then jump to the step labeled *conversion*.

    8.  If the character indicated by `position`{.variable} is an [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-8
        x-internal="ascii-digits"}, jump back to the step labeled
        *fraction loop* in these substeps.

14. If the character indicated by `position`{.variable} is U+0065 (e) or
    a U+0045 (E), then:

    1.  Advance `position`{.variable} to the next character.

    2.  If `position`{.variable} is past the end of `input`{.variable},
        then jump to the step labeled *conversion*.

    3.  If the character indicated by `position`{.variable} is a U+002D
        HYPHEN-MINUS character (-):

        1.  Change `exponent`{.variable} to −1.

        2.  Advance `position`{.variable} to the next character.

        3.  If `position`{.variable} is past the end of
            `input`{.variable}, then jump to the step labeled
            *conversion*.

        Otherwise, if the character indicated by `position`{.variable}
        is a U+002B PLUS SIGN character (+):

        1.  Advance `position`{.variable} to the next character.

        2.  If `position`{.variable} is past the end of
            `input`{.variable}, then jump to the step labeled
            *conversion*.

    4.  If the character indicated by `position`{.variable} is not an
        [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-9
        x-internal="ascii-digits"}, then jump to the step labeled
        *conversion*.

    5.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#floating-point-numbers:collect-a-sequence-of-code-points-2
        x-internal="collect-a-sequence-of-code-points"} that are [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#floating-point-numbers:ascii-digits-10
        x-internal="ascii-digits"} from `input`{.variable} given
        `position`{.variable}, and interpret the resulting sequence as a
        base-ten integer. Multiply `exponent`{.variable} by that
        integer.

    6.  Multiply `value`{.variable} by ten raised to the
        `exponent`{.variable}th power.

15. *Conversion*: Let `S`{.variable} be the set of finite IEEE 754
    double-precision floating-point values except −0, but with two
    special values added: 2^1024^ and −2^1024^.

16. Let `rounded-value`{.variable} be the number in `S`{.variable} that
    is closest to `value`{.variable}, selecting the number with an even
    significand if there are two equally close values. (The two special
    values 2^1024^ and −2^1024^ are considered to have even significands
    for this purpose.)

17. If `rounded-value`{.variable} is 2^1024^ or −2^1024^, return an
    error.

18. Return `rounded-value`{.variable}.

##### [2.3.4.4]{.secno} Percentages and lengths[](#percentages-and-dimensions){.self-link} {#percentages-and-dimensions}

The [rules for parsing dimension
values]{#rules-for-parsing-dimension-values .dfn} are as given in the
following algorithm. When invoked, the steps must be followed in the
order given, aborting at the first step that returns a value. This
algorithm will return either a number greater than or equal to 0.0, or
failure; if a number is returned, then it is further categorized as
either a percentage or a length.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a [position
    variable](https://infra.spec.whatwg.org/#string-position-variable){#percentages-and-dimensions:position-variable
    x-internal="position-variable"} for `input`{.variable}, initially
    pointing at the start of `input`{.variable}.

3.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#percentages-and-dimensions:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

4.  If `position`{.variable} is past the end of `input`{.variable} or
    the code point at `position`{.variable} within `input`{.variable} is
    not an [ASCII
    digit](https://infra.spec.whatwg.org/#ascii-digit){#percentages-and-dimensions:ascii-digits
    x-internal="ascii-digits"}, then return failure.

5.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#percentages-and-dimensions:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#percentages-and-dimensions:ascii-digits-2
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}, and interpret the resulting sequence as a
    base-ten integer. Let `value`{.variable} be that number.

6.  If `position`{.variable} is past the end of `input`{.variable}, then
    return `value`{.variable} as a length.

7.  If the code point at `position`{.variable} within `input`{.variable}
    is U+002E (.), then:

    1.  Advance `position`{.variable} by 1.

    2.  If `position`{.variable} is past the end of `input`{.variable}
        or the code point at `position`{.variable} within
        `input`{.variable} is not an [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#percentages-and-dimensions:ascii-digits-3
        x-internal="ascii-digits"}, then return the [current dimension
        value](#current-dimension-value){#percentages-and-dimensions:current-dimension-value}
        with `value`{.variable}, `input`{.variable}, and
        `position`{.variable}.

    3.  Let `divisor`{.variable} have the value 1.

    4.  While true:

        1.  Multiply `divisor`{.variable} by ten.

        2.  Add the value of the code point at `position`{.variable}
            within `input`{.variable}, interpreted as a base-ten digit
            (0..9) and divided by `divisor`{.variable}, to
            `value`{.variable}.

        3.  Advance `position`{.variable} by 1.

        4.  If `position`{.variable} is past the end of
            `input`{.variable}, then return `value`{.variable} as a
            length.

        5.  If the code point at `position`{.variable} within
            `input`{.variable} is not an [ASCII
            digit](https://infra.spec.whatwg.org/#ascii-digit){#percentages-and-dimensions:ascii-digits-4
            x-internal="ascii-digits"}, then
            [break](https://infra.spec.whatwg.org/#iteration-break){#percentages-and-dimensions:break
            x-internal="break"}.

8.  Return the [current dimension
    value](#current-dimension-value){#percentages-and-dimensions:current-dimension-value-2}
    with `value`{.variable}, `input`{.variable}, and
    `position`{.variable}.

The [current dimension value]{#current-dimension-value .dfn}, given
`value`{.variable}, `input`{.variable}, and `position`{.variable}, is
determined as follows:

1.  If `position`{.variable} is past the end of `input`{.variable}, then
    return `value`{.variable} as a length.

2.  If the code point at `position`{.variable} within `input`{.variable}
    is U+0025 (%), then return `value`{.variable} as a percentage.

3.  Return `value`{.variable} as a length.

##### [2.3.4.5]{.secno} Nonzero percentages and lengths[](#nonzero-percentages-and-lengths){.self-link}

The [rules for parsing nonzero dimension
values]{#rules-for-parsing-non-zero-dimension-values .dfn} are as given
in the following algorithm. When invoked, the steps must be followed in
the order given, aborting at the first step that returns a value. This
algorithm will return either a number greater than 0.0, or an error; if
a number is returned, then it is further categorized as either a
percentage or a length.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `value`{.variable} be the result of parsing `input`{.variable}
    using the [rules for parsing dimension
    values](#rules-for-parsing-dimension-values){#nonzero-percentages-and-lengths:rules-for-parsing-dimension-values}.

3.  If `value`{.variable} is an error, return an error.

4.  If `value`{.variable} is zero, return an error.

5.  If `value`{.variable} is a percentage, return `value`{.variable} as
    a percentage.

6.  Return `value`{.variable} as a length.

##### [2.3.4.6]{.secno} Lists of floating-point numbers[](#lists-of-floating-point-numbers){.self-link}

A [valid list of floating-point
numbers]{#valid-list-of-floating-point-numbers .dfn} is a number of
[valid floating-point
numbers](#valid-floating-point-number){#lists-of-floating-point-numbers:valid-floating-point-number}
separated by U+002C COMMA characters, with no other characters (e.g. no
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-floating-point-numbers:space-characters
x-internal="space-characters"}). In addition, there might be
restrictions on the number of floating-point numbers that can be given,
or on the range of values allowed.

The [rules for parsing a list of floating-point
numbers]{#rules-for-parsing-a-list-of-floating-point-numbers .dfn} are
as follows:

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Let `numbers`{.variable} be an initially empty list of
    floating-point numbers. This list will be the result of this
    algorithm.

4.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-floating-point-numbers:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-floating-point-numbers:space-characters-2
    x-internal="space-characters"}, U+002C COMMA, or U+003B SEMICOLON
    characters from `input`{.variable} given `position`{.variable}. This
    skips past any leading delimiters.

5.  While `position`{.variable} is not past the end of
    `input`{.variable}:

    1.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-floating-point-numbers:collect-a-sequence-of-code-points-2
        x-internal="collect-a-sequence-of-code-points"} that are not
        [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-floating-point-numbers:space-characters-3
        x-internal="space-characters"}, U+002C COMMA, U+003B SEMICOLON,
        [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#lists-of-floating-point-numbers:ascii-digits
        x-internal="ascii-digits"}, U+002E FULL STOP, or U+002D
        HYPHEN-MINUS characters from `input`{.variable} given
        `position`{.variable}. This skips past leading garbage.

    2.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-floating-point-numbers:collect-a-sequence-of-code-points-3
        x-internal="collect-a-sequence-of-code-points"} that are not
        [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-floating-point-numbers:space-characters-4
        x-internal="space-characters"}, U+002C COMMA, or U+003B
        SEMICOLON characters from `input`{.variable} given
        `position`{.variable}, and let `unparsed number`{.variable} be
        the result.

    3.  Let `number`{.variable} be the result of parsing
        `unparsed number`{.variable} using the [rules for parsing
        floating-point number
        values](#rules-for-parsing-floating-point-number-values){#lists-of-floating-point-numbers:rules-for-parsing-floating-point-number-values}.

    4.  If `number`{.variable} is an error, set `number`{.variable} to
        zero.

    5.  Append `number`{.variable} to `numbers`{.variable}.

    6.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-floating-point-numbers:collect-a-sequence-of-code-points-4
        x-internal="collect-a-sequence-of-code-points"} that are [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-floating-point-numbers:space-characters-5
        x-internal="space-characters"}, U+002C COMMA, or U+003B
        SEMICOLON characters from `input`{.variable} given
        `position`{.variable}. This skips past the delimiter.

6.  Return `numbers`{.variable}.

##### [2.3.4.7]{.secno} Lists of dimensions[](#lists-of-dimensions){.self-link}

The [rules for parsing a list of
dimensions]{#rules-for-parsing-a-list-of-dimensions .dfn} are as
follows. These rules return a list of zero or more pairs consisting of a
number and a unit, the unit being one of *percentage*, *relative*, and
*absolute*.

1.  Let `raw input`{.variable} be the string being parsed.

2.  If the last character in `raw input`{.variable} is a U+002C COMMA
    character (,), then remove that character from
    `raw input`{.variable}.

3.  [Split the string `raw input`{.variable} on
    commas](https://infra.spec.whatwg.org/#split-on-commas){#lists-of-dimensions:split-a-string-on-commas
    x-internal="split-a-string-on-commas"}. Let `raw tokens`{.variable}
    be the resulting list of tokens.

4.  Let `result`{.variable} be an empty list of number/unit pairs.

5.  For each token in `raw tokens`{.variable}, run the following
    substeps:

    1.  Let `input`{.variable} be the token.

    2.  Let `position`{.variable} be a pointer into `input`{.variable},
        initially pointing at the start of the string.

    3.  Let `value`{.variable} be the number 0.

    4.  Let `unit`{.variable} be *absolute*.

    5.  If `position`{.variable} is past the end of `input`{.variable},
        set `unit`{.variable} to *relative* and jump to the last
        substep.

    6.  If the character at `position`{.variable} is an [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#lists-of-dimensions:ascii-digits
        x-internal="ascii-digits"}, [collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-dimensions:collect-a-sequence-of-code-points
        x-internal="collect-a-sequence-of-code-points"} that are [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#lists-of-dimensions:ascii-digits-2
        x-internal="ascii-digits"} from `input`{.variable} given
        `position`{.variable}, interpret the resulting sequence as an
        integer in base ten, and increment `value`{.variable} by that
        integer.

    7.  If the character at `position`{.variable} is U+002E (.), then:

        1.  [Collect a sequence of code
            points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#lists-of-dimensions:collect-a-sequence-of-code-points-2
            x-internal="collect-a-sequence-of-code-points"} consisting
            of [ASCII
            whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-dimensions:space-characters
            x-internal="space-characters"} and [ASCII
            digits](https://infra.spec.whatwg.org/#ascii-digit){#lists-of-dimensions:ascii-digits-3
            x-internal="ascii-digits"} from `input`{.variable} given
            `position`{.variable}. Let `s`{.variable} be the resulting
            sequence.

        2.  Remove all [ASCII
            whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#lists-of-dimensions:space-characters-2
            x-internal="space-characters"} in `s`{.variable}.

        3.  If `s`{.variable} is not the empty string, then:

            1.  Let `length`{.variable} be the number of characters in
                `s`{.variable} (after the spaces were removed).

            2.  Let `fraction`{.variable} be the result of interpreting
                `s`{.variable} as a base-ten integer, and then dividing
                that number by 10^`length`{.variable}^.

            3.  Increment `value`{.variable} by `fraction`{.variable}.

    8.  [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#lists-of-dimensions:skip-ascii-whitespace
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

    9.  If the character at `position`{.variable} is a U+0025 PERCENT
        SIGN character (%), then set `unit`{.variable} to *percentage*.

        Otherwise, if the character at `position`{.variable} is a U+002A
        ASTERISK character (\*), then set `unit`{.variable} to
        *relative*.

    10. Add an entry to `result`{.variable} consisting of the number
        given by `value`{.variable} and the unit given by
        `unit`{.variable}.

6.  Return the list `result`{.variable}.

#### [2.3.5]{.secno} Dates and times[](#dates-and-times){.self-link}

In the algorithms below, the [number of days in month `month`{.variable}
of year `year`{.variable}]{#number-of-days-in-month-month-of-year-year
.dfn} is: *31* if `month`{.variable} is 1, 3, 5, 7, 8, 10, or 12; *30*
if `month`{.variable} is 4, 6, 9, or 11; *29* if `month`{.variable} is 2
and `year`{.variable} is a number divisible by 400, or if
`year`{.variable} is a number divisible by 4 but not by 100; and *28*
otherwise. This takes into account leap years in the Gregorian calendar.
[\[GREGORIAN\]](references.html#refsGREGORIAN)

When [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#dates-and-times:ascii-digits
x-internal="ascii-digits"} are used in the date and time syntaxes
defined in this section, they express numbers in base ten.

While the formats described here are intended to be subsets of the
corresponding ISO8601 formats, this specification defines parsing rules
in much more detail than ISO8601. Implementers are therefore encouraged
to carefully examine any date parsing libraries before using them to
implement the parsing rules described below; ISO8601 libraries might not
parse dates and times in exactly the same manner.
[\[ISO8601\]](references.html#refsISO8601)

Where this specification refers to the [proleptic Gregorian
calendar]{#proleptic-gregorian-calendar .dfn}, it means the modern
Gregorian calendar, extrapolated backwards to year 1. A date in the
[proleptic Gregorian
calendar](#proleptic-gregorian-calendar){#dates-and-times:proleptic-gregorian-calendar},
sometimes explicitly referred to as a [proleptic-Gregorian
date]{#proleptic-gregorian-date .dfn}, is one that is described using
that calendar even if that calendar was not in use at the time (or
place) in question. [\[GREGORIAN\]](references.html#refsGREGORIAN)

The use of the Gregorian calendar as the wire format in this
specification is an arbitrary choice resulting from the cultural biases
of those involved in the decision. See also the section discussing
[date, time, and number formats](forms.html#input-author-notes) in forms
(for authors), [implementation notes regarding localization of form
controls](input.html#input-impl-notes), and the
[`time`{#dates-and-times:the-time-element}](text-level-semantics.html#the-time-element)
element.

##### [2.3.5.1]{.secno} Months[](#months){.self-link}

A [month]{#concept-month .dfn} consists of a specific
[proleptic-Gregorian
date](#proleptic-gregorian-date){#months:proleptic-gregorian-date} with
no time-zone information and no date information beyond a year and a
month. [\[GREGORIAN\]](references.html#refsGREGORIAN)

A string is a [valid month string]{#valid-month-string .dfn}
representing a year `year`{.variable} and month `month`{.variable} if it
consists of the following components in the given order:

1.  Four or more [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#months:ascii-digits
    x-internal="ascii-digits"}, representing `year`{.variable}, where
    `year`{.variable} \> 0
2.  A U+002D HYPHEN-MINUS character (-)
3.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#months:ascii-digits-2
    x-internal="ascii-digits"}, representing the month
    `month`{.variable}, in the range 1 ≤ `month`{.variable} ≤ 12

The rules to [parse a month string]{#parse-a-month-string .dfn} are as
follows. This will return either a year and month, or nothing. If at any
point the algorithm says that it \"fails\", this means that it is
aborted at that point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a month
    component](#parse-a-month-component){#months:parse-a-month-component}
    to obtain `year`{.variable} and `month`{.variable}. If this returns
    nothing, then fail.

4.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

5.  Return `year`{.variable} and `month`{.variable}.

The rules to [parse a month component]{#parse-a-month-component .dfn},
given an `input`{.variable} string and a `position`{.variable}, are as
follows. This will return either a year and a month, or nothing. If at
any point the algorithm says that it \"fails\", this means that it is
aborted at that point and returns nothing.

1.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#months:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#months:ascii-digits-3
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not at least
    four characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `year`{.variable} be that
    number.

2.  If `year`{.variable} is not a number greater than zero, then fail.

3.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+002D
    HYPHEN-MINUS character, then fail. Otherwise, move
    `position`{.variable} forwards one character.

4.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#months:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#months:ascii-digits-4
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `month`{.variable} be that
    number.

5.  If `month`{.variable} is not a number in the range
    1 ≤ `month`{.variable} ≤ 12, then fail.

6.  Return `year`{.variable} and `month`{.variable}.

##### [2.3.5.2]{.secno} Dates[](#dates){.self-link}

A [date]{#concept-date .dfn} consists of a specific [proleptic-Gregorian
date](#proleptic-gregorian-date){#dates:proleptic-gregorian-date} with
no time-zone information, consisting of a year, a month, and a day.
[\[GREGORIAN\]](references.html#refsGREGORIAN)

A string is a [valid date string]{#valid-date-string .dfn} representing
a year `year`{.variable}, month `month`{.variable}, and day
`day`{.variable} if it consists of the following components in the given
order:

1.  A [valid month
    string](#valid-month-string){#dates:valid-month-string},
    representing `year`{.variable} and `month`{.variable}
2.  A U+002D HYPHEN-MINUS character (-)
3.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#dates:ascii-digits
    x-internal="ascii-digits"}, representing `day`{.variable}, in the
    range 1 ≤ `day`{.variable} ≤ `maxday`{.variable} where
    `maxday`{.variable} is the [number of days in the month
    `month`{.variable} and year
    `year`{.variable}](#number-of-days-in-month-month-of-year-year){#dates:number-of-days-in-month-month-of-year-year}

The rules to [parse a date string]{#parse-a-date-string .dfn} are as
follows. This will return either a date, or nothing. If at any point the
algorithm says that it \"fails\", this means that it is aborted at that
point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a date
    component](#parse-a-date-component){#dates:parse-a-date-component}
    to obtain `year`{.variable}, `month`{.variable}, and
    `day`{.variable}. If this returns nothing, then fail.

4.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

5.  Let `date`{.variable} be the date with year `year`{.variable}, month
    `month`{.variable}, and day `day`{.variable}.

6.  Return `date`{.variable}.

The rules to [parse a date component]{#parse-a-date-component .dfn},
given an `input`{.variable} string and a `position`{.variable}, are as
follows. This will return either a year, a month, and a day, or nothing.
If at any point the algorithm says that it \"fails\", this means that it
is aborted at that point and returns nothing.

1.  [Parse a month
    component](#parse-a-month-component){#dates:parse-a-month-component}
    to obtain `year`{.variable} and `month`{.variable}. If this returns
    nothing, then fail.

2.  Let `maxday`{.variable} be the [number of days in month
    `month`{.variable} of year
    `year`{.variable}](#number-of-days-in-month-month-of-year-year){#dates:number-of-days-in-month-month-of-year-year-2}.

3.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+002D
    HYPHEN-MINUS character, then fail. Otherwise, move
    `position`{.variable} forwards one character.

4.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#dates:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#dates:ascii-digits-2
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `day`{.variable} be that number.

5.  If `day`{.variable} is not a number in the range
    1 ≤ `day`{.variable} ≤ `maxday`{.variable}, then fail.

6.  Return `year`{.variable}, `month`{.variable}, and `day`{.variable}.

##### [2.3.5.3]{.secno} Yearless dates[](#yearless-dates){.self-link}

A [yearless date]{#concept-yearless-date .dfn} consists of a Gregorian
month and a day within that month, but with no associated year.
[\[GREGORIAN\]](references.html#refsGREGORIAN)

A string is a [valid yearless date string]{#valid-yearless-date-string
.dfn} representing a month `month`{.variable} and a day `day`{.variable}
if it consists of the following components in the given order:

1.  Optionally, two U+002D HYPHEN-MINUS characters (-)
2.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#yearless-dates:ascii-digits
    x-internal="ascii-digits"}, representing the month
    `month`{.variable}, in the range 1 ≤ `month`{.variable} ≤ 12
3.  A U+002D HYPHEN-MINUS character (-)
4.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#yearless-dates:ascii-digits-2
    x-internal="ascii-digits"}, representing `day`{.variable}, in the
    range 1 ≤ `day`{.variable} ≤ `maxday`{.variable} where
    `maxday`{.variable} is the [number of
    days](#number-of-days-in-month-month-of-year-year){#yearless-dates:number-of-days-in-month-month-of-year-year}
    in the month `month`{.variable} and any arbitrary leap year (e.g. 4
    or 2000)

In other words, if the `month`{.variable} is \"`02`\", meaning February,
then the day can be 29, as if the year was a leap year.

The rules to [parse a yearless date
string]{#parse-a-yearless-date-string .dfn} are as follows. This will
return either a month and a day, or nothing. If at any point the
algorithm says that it \"fails\", this means that it is aborted at that
point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a yearless date
    component](#parse-a-yearless-date-component){#yearless-dates:parse-a-yearless-date-component}
    to obtain `month`{.variable} and `day`{.variable}. If this returns
    nothing, then fail.

4.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

5.  Return `month`{.variable} and `day`{.variable}.

The rules to [parse a yearless date
component]{#parse-a-yearless-date-component .dfn}, given an
`input`{.variable} string and a `position`{.variable}, are as follows.
This will return either a month and a day, or nothing. If at any point
the algorithm says that it \"fails\", this means that it is aborted at
that point and returns nothing.

1.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#yearless-dates:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are U+002D
    HYPHEN-MINUS characters (-) from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly zero
    or two characters long, then fail.

2.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#yearless-dates:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#yearless-dates:ascii-digits-3
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `month`{.variable} be that
    number.

3.  If `month`{.variable} is not a number in the range
    1 ≤ `month`{.variable} ≤ 12, then fail.

4.  Let `maxday`{.variable} be the [number of
    days](#number-of-days-in-month-month-of-year-year){#yearless-dates:number-of-days-in-month-month-of-year-year-2}
    in month `month`{.variable} of any arbitrary leap year (e.g. 4 or
    2000).

5.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+002D
    HYPHEN-MINUS character, then fail. Otherwise, move
    `position`{.variable} forwards one character.

6.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#yearless-dates:collect-a-sequence-of-code-points-3
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#yearless-dates:ascii-digits-4
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `day`{.variable} be that number.

7.  If `day`{.variable} is not a number in the range
    1 ≤ `day`{.variable} ≤ `maxday`{.variable}, then fail.

8.  Return `month`{.variable} and `day`{.variable}.

##### [2.3.5.4]{.secno} Times[](#times){.self-link}

A [time]{#concept-time .dfn} consists of a specific time with no
time-zone information, consisting of an hour, a minute, a second, and a
fraction of a second.

A string is a [valid time string]{#valid-time-string .dfn} representing
an hour `hour`{.variable}, a minute `minute`{.variable}, and a second
`second`{.variable} if it consists of the following components in the
given order:

1.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits
    x-internal="ascii-digits"}, representing `hour`{.variable}, in the
    range 0 ≤ `hour`{.variable} ≤ 23
2.  A U+003A COLON character (:)
3.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-2
    x-internal="ascii-digits"}, representing `minute`{.variable}, in the
    range 0 ≤ `minute`{.variable} ≤ 59
4.  If `second`{.variable} is nonzero, or optionally if
    `second`{.variable} is zero:
    1.  A U+003A COLON character (:)
    2.  Two [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-3
        x-internal="ascii-digits"}, representing the integer part of
        `second`{.variable}, in the range 0 ≤ `s`{.variable} ≤ 59
    3.  If `second`{.variable} is not an integer, or optionally if
        `second`{.variable} is an integer:
        1.  A U+002E FULL STOP character (.)
        2.  One, two, or three [ASCII
            digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-4
            x-internal="ascii-digits"}, representing the fractional part
            of `second`{.variable}

The `second`{.variable} component cannot be 60 or 61; leap seconds
cannot be represented.

The rules to [parse a time string]{#parse-a-time-string .dfn} are as
follows. This will return either a time, or nothing. If at any point the
algorithm says that it \"fails\", this means that it is aborted at that
point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a time
    component](#parse-a-time-component){#times:parse-a-time-component}
    to obtain `hour`{.variable}, `minute`{.variable}, and
    `second`{.variable}. If this returns nothing, then fail.

4.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

5.  Let `time`{.variable} be the time with hour `hour`{.variable},
    minute `minute`{.variable}, and second `second`{.variable}.

6.  Return `time`{.variable}.

The rules to [parse a time component]{#parse-a-time-component .dfn},
given an `input`{.variable} string and a `position`{.variable}, are as
follows. This will return either an hour, a minute, and a second, or
nothing. If at any point the algorithm says that it \"fails\", this
means that it is aborted at that point and returns nothing.

1.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#times:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-5
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `hour`{.variable} be that
    number.

2.  If `hour`{.variable} is not a number in the range
    0 ≤ `hour`{.variable} ≤ 23, then fail.

3.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+003A COLON
    character, then fail. Otherwise, move `position`{.variable} forwards
    one character.

4.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#times:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-6
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `minute`{.variable} be that
    number.

5.  If `minute`{.variable} is not a number in the range
    0 ≤ `minute`{.variable} ≤ 59, then fail.

6.  Let `second`{.variable} be 0.

7.  If `position`{.variable} is not beyond the end of `input`{.variable}
    and the character at `position`{.variable} is U+003A (:), then:

    1.  Advance `position`{.variable} to the next character in
        `input`{.variable}.

    2.  If `position`{.variable} is beyond the end of
        `input`{.variable}, or at the last character in
        `input`{.variable}, or if the next *two* characters in
        `input`{.variable} starting at `position`{.variable} are not
        both [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-7
        x-internal="ascii-digits"}, then fail.

    3.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#times:collect-a-sequence-of-code-points-3
        x-internal="collect-a-sequence-of-code-points"} that are either
        [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#times:ascii-digits-8
        x-internal="ascii-digits"} or U+002E FULL STOP characters from
        `input`{.variable} given `position`{.variable}. If the collected
        sequence is three characters long, or if it is longer than three
        characters long and the third character is not a U+002E FULL
        STOP character, or if it has more than one U+002E FULL STOP
        character, then fail. Otherwise, interpret the resulting
        sequence as a base-ten number (possibly with a fractional part).
        Set `second`{.variable} to that number.

    4.  If `second`{.variable} is not a number in the range
        0 ≤ `second`{.variable} \< 60, then fail.

8.  Return `hour`{.variable}, `minute`{.variable}, and
    `second`{.variable}.

##### [2.3.5.5]{.secno} Local dates and times[](#local-dates-and-times){.self-link}

A [local date and time]{#concept-datetime-local .dfn} consists of a
specific [proleptic-Gregorian
date](#proleptic-gregorian-date){#local-dates-and-times:proleptic-gregorian-date},
consisting of a year, a month, and a day, and a time, consisting of an
hour, a minute, a second, and a fraction of a second, but expressed
without a time zone. [\[GREGORIAN\]](references.html#refsGREGORIAN)

A string is a [valid local date and time
string]{#valid-local-date-and-time-string .dfn} representing a date and
time if it consists of the following components in the given order:

1.  A [valid date
    string](#valid-date-string){#local-dates-and-times:valid-date-string}
    representing the date
2.  A U+0054 LATIN CAPITAL LETTER T character (T) or a U+0020 SPACE
    character
3.  A [valid time
    string](#valid-time-string){#local-dates-and-times:valid-time-string}
    representing the time

A string is a [valid normalized local date and time
string]{#valid-normalised-local-date-and-time-string .dfn} representing
a date and time if it consists of the following components in the given
order:

1.  A [valid date
    string](#valid-date-string){#local-dates-and-times:valid-date-string-2}
    representing the date
2.  A U+0054 LATIN CAPITAL LETTER T character (T)
3.  A [valid time
    string](#valid-time-string){#local-dates-and-times:valid-time-string-2}
    representing the time, expressed as the shortest possible string for
    the given time (e.g. omitting the seconds component entirely if the
    given time is zero seconds past the minute)

The rules to [parse a local date and time
string]{#parse-a-local-date-and-time-string .dfn} are as follows. This
will return either a date and time, or nothing. If at any point the
algorithm says that it \"fails\", this means that it is aborted at that
point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a date
    component](#parse-a-date-component){#local-dates-and-times:parse-a-date-component}
    to obtain `year`{.variable}, `month`{.variable}, and
    `day`{.variable}. If this returns nothing, then fail.

4.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is neither a U+0054 LATIN
    CAPITAL LETTER T character (T) nor a U+0020 SPACE character, then
    fail. Otherwise, move `position`{.variable} forwards one character.

5.  [Parse a time
    component](#parse-a-time-component){#local-dates-and-times:parse-a-time-component}
    to obtain `hour`{.variable}, `minute`{.variable}, and
    `second`{.variable}. If this returns nothing, then fail.

6.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

7.  Let `date`{.variable} be the date with year `year`{.variable}, month
    `month`{.variable}, and day `day`{.variable}.

8.  Let `time`{.variable} be the time with hour `hour`{.variable},
    minute `minute`{.variable}, and second `second`{.variable}.

9.  Return `date`{.variable} and `time`{.variable}.

##### [2.3.5.6]{.secno} Time zones[](#time-zones){.self-link}

A [time-zone offset]{#concept-timezone .dfn} consists of a signed number
of hours and minutes.

A string is a [valid time-zone offset
string]{#valid-time-zone-offset-string .dfn} representing a time-zone
offset if it consists of either:

- A U+005A LATIN CAPITAL LETTER Z character (Z), allowed only if the
  time zone is UTC

- Or, the following components, in the given order:

  1.  Either a U+002B PLUS SIGN character (+) or, if the time-zone
      offset is not zero, a U+002D HYPHEN-MINUS character (-),
      representing the sign of the time-zone offset
  2.  Two [ASCII
      digits](https://infra.spec.whatwg.org/#ascii-digit){#time-zones:ascii-digits
      x-internal="ascii-digits"}, representing the hours component
      `hour`{.variable} of the time-zone offset, in the range
      0 ≤ `hour`{.variable} ≤ 23
  3.  Optionally, a U+003A COLON character (:)
  4.  Two [ASCII
      digits](https://infra.spec.whatwg.org/#ascii-digit){#time-zones:ascii-digits-2
      x-internal="ascii-digits"}, representing the minutes component
      `minute`{.variable} of the time-zone offset, in the range
      0 ≤ `minute`{.variable} ≤ 59

This format allows for time-zone offsets from -23:59 to +23:59. Right
now, in practice, the range of offsets of actual time zones is -12:00 to
+14:00, and the minutes component of offsets of actual time zones is
always either 00, 30, or 45. There is no guarantee that this will remain
so forever, however, since time zones are used as political footballs
and are thus subject to very whimsical policy decisions.

See also the usage notes and examples in the [global date and
time](#concept-datetime){#time-zones:concept-datetime} section below for
details on using time-zone offsets with historical times that predate
the formation of formal time zones.

The rules to [parse a time-zone offset
string]{#parse-a-time-zone-offset-string .dfn} are as follows. This will
return either a time-zone offset, or nothing. If at any point the
algorithm says that it \"fails\", this means that it is aborted at that
point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a time-zone offset
    component](#parse-a-time-zone-offset-component){#time-zones:parse-a-time-zone-offset-component}
    to obtain `timezone`{.variable}~`hours`{.variable}~ and
    `timezone`{.variable}~`minutes`{.variable}~. If this returns
    nothing, then fail.

4.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

5.  Return the time-zone offset that is
    `timezone`{.variable}~`hours`{.variable}~ hours and
    `timezone`{.variable}~`minutes`{.variable}~ minutes from UTC.

The rules to [parse a time-zone offset
component]{#parse-a-time-zone-offset-component .dfn}, given an
`input`{.variable} string and a `position`{.variable}, are as follows.
This will return either time-zone hours and time-zone minutes, or
nothing. If at any point the algorithm says that it \"fails\", this
means that it is aborted at that point and returns nothing.

1.  If the character at `position`{.variable} is a U+005A LATIN CAPITAL
    LETTER Z character (Z), then:

    1.  Let `timezone`{.variable}~`hours`{.variable}~ be 0.

    2.  Let `timezone`{.variable}~`minutes`{.variable}~ be 0.

    3.  Advance `position`{.variable} to the next character in
        `input`{.variable}.

    Otherwise, if the character at `position`{.variable} is either a
    U+002B PLUS SIGN (+) or a U+002D HYPHEN-MINUS (-), then:

    1.  If the character at `position`{.variable} is a U+002B PLUS SIGN
        (+), let `sign`{.variable} be \"positive\". Otherwise, it\'s a
        U+002D HYPHEN-MINUS (-); let `sign`{.variable} be \"negative\".

    2.  Advance `position`{.variable} to the next character in
        `input`{.variable}.

    3.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#time-zones:collect-a-sequence-of-code-points
        x-internal="collect-a-sequence-of-code-points"} that are [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#time-zones:ascii-digits-3
        x-internal="ascii-digits"} from `input`{.variable} given
        `position`{.variable}. Let `s`{.variable} be the collected
        sequence.

    4.  If `s`{.variable} is exactly two characters long, then:

        1.  Interpret `s`{.variable} as a base-ten integer. Let
            `timezone`{.variable}~`hours`{.variable}~ be that number.

        2.  If `position`{.variable} is beyond the end of
            `input`{.variable} or if the character at
            `position`{.variable} is not a U+003A COLON character, then
            fail. Otherwise, move `position`{.variable} forwards one
            character.

        3.  [Collect a sequence of code
            points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#time-zones:collect-a-sequence-of-code-points-2
            x-internal="collect-a-sequence-of-code-points"} that are
            [ASCII
            digits](https://infra.spec.whatwg.org/#ascii-digit){#time-zones:ascii-digits-4
            x-internal="ascii-digits"} from `input`{.variable} given
            `position`{.variable}. If the collected sequence is not
            exactly two characters long, then fail. Otherwise, interpret
            the resulting sequence as a base-ten integer. Let
            `timezone`{.variable}~`minutes`{.variable}~ be that number.

        If `s`{.variable} is exactly four characters long, then:

        1.  Interpret the first two characters of `s`{.variable} as a
            base-ten integer. Let
            `timezone`{.variable}~`hours`{.variable}~ be that number.

        2.  Interpret the last two characters of `s`{.variable} as a
            base-ten integer. Let
            `timezone`{.variable}~`minutes`{.variable}~ be that number.

        Otherwise, fail.

    5.  If `timezone`{.variable}~`hours`{.variable}~ is not a number in
        the range 0 ≤ `timezone`{.variable}~`hours`{.variable}~ ≤ 23,
        then fail.

    6.  If `sign`{.variable} is \"negative\", then negate
        `timezone`{.variable}~`hours`{.variable}~.

    7.  If `timezone`{.variable}~`minutes`{.variable}~ is not a number
        in the range
        0 ≤ `timezone`{.variable}~`minutes`{.variable}~ ≤ 59, then fail.

    8.  If `sign`{.variable} is \"negative\", then negate
        `timezone`{.variable}~`minutes`{.variable}~.

    Otherwise, fail.

2.  Return `timezone`{.variable}~`hours`{.variable}~ and
    `timezone`{.variable}~`minutes`{.variable}~.

##### [2.3.5.7]{.secno} Global dates and times[](#global-dates-and-times){.self-link}

A [global date and time]{#concept-datetime .dfn} consists of a specific
[proleptic-Gregorian
date](#proleptic-gregorian-date){#global-dates-and-times:proleptic-gregorian-date},
consisting of a year, a month, and a day, and a time, consisting of an
hour, a minute, a second, and a fraction of a second, expressed with a
time-zone offset, consisting of a signed number of hours and minutes.
[\[GREGORIAN\]](references.html#refsGREGORIAN)

A string is a [valid global date and time
string]{#valid-global-date-and-time-string .dfn} representing a date,
time, and a time-zone offset if it consists of the following components
in the given order:

1.  A [valid date
    string](#valid-date-string){#global-dates-and-times:valid-date-string}
    representing the date
2.  A U+0054 LATIN CAPITAL LETTER T character (T) or a U+0020 SPACE
    character
3.  A [valid time
    string](#valid-time-string){#global-dates-and-times:valid-time-string}
    representing the time
4.  A [valid time-zone offset
    string](#valid-time-zone-offset-string){#global-dates-and-times:valid-time-zone-offset-string}
    representing the time-zone offset

Times in dates before the formation of UTC in the mid-twentieth century
must be expressed and interpreted in terms of UT1 (contemporary Earth
solar time at the 0° longitude), not UTC (the approximation of UT1 that
ticks in SI seconds). Time before the formation of time zones must be
expressed and interpreted as UT1 times with explicit time zones that
approximate the contemporary difference between the appropriate local
time and the time observed at the location of Greenwich, London.

::: example
The following are some examples of dates written as [valid global date
and time
strings](#valid-global-date-and-time-string){#global-dates-and-times:valid-global-date-and-time-string}.

\"`0037-12-13 00:00Z`\"
:   Midnight in areas using London time on the birthday of Nero (the
    Roman Emperor). See below for further discussion on which date this
    actually corresponds to.

\"`1979-10-14T12:00:00.001-04:00`\"
:   One millisecond after noon on October 14th 1979, in the time zone in
    use on the east coast of the USA during daylight saving time.

\"`8592-01-01T02:09+02:09`\"
:   Midnight UTC on the 1st of January, 8592. The time zone associated
    with that time is two hours and nine minutes ahead of UTC, which is
    not currently a real time zone, but is nonetheless allowed.

Several things are notable about these dates:

- Years with fewer than four digits have to be zero-padded. The date
  \"37-12-13\" would not be a valid date.
- If the \"`T`\" is replaced by a space, it must be a single space
  character. The string \"`2001-12-21  12:00Z`\" (with two spaces
  between the components) would not be parsed successfully.
- To unambiguously identify a moment in time prior to the introduction
  of the Gregorian calendar (insofar as moments in time before the
  formation of UTC can be unambiguously identified), the date has to be
  first converted to the Gregorian calendar from the calendar in use at
  the time (e.g. from the Julian calendar). The date of Nero\'s birth is
  the 15th of December 37, in the Julian Calendar, which is the 13th of
  December 37 in the [proleptic Gregorian
  calendar](#proleptic-gregorian-calendar){#global-dates-and-times:proleptic-gregorian-calendar}.
- The time and time-zone offset components are not optional.
- Dates before the year one can\'t be represented as a datetime in this
  version of HTML.
- Times of specific events in ancient times are, at best,
  approximations, since time was not well coordinated or measured until
  relatively recent decades.
- Time-zone offsets differ based on daylight saving time.
:::

The rules to [parse a global date and time
string]{#parse-a-global-date-and-time-string .dfn} are as follows. This
will return either a time in UTC, with associated time-zone offset
information for round-tripping or display purposes, or nothing. If at
any point the algorithm says that it \"fails\", this means that it is
aborted at that point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Parse a date
    component](#parse-a-date-component){#global-dates-and-times:parse-a-date-component}
    to obtain `year`{.variable}, `month`{.variable}, and
    `day`{.variable}. If this returns nothing, then fail.

4.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is neither a U+0054 LATIN
    CAPITAL LETTER T character (T) nor a U+0020 SPACE character, then
    fail. Otherwise, move `position`{.variable} forwards one character.

5.  [Parse a time
    component](#parse-a-time-component){#global-dates-and-times:parse-a-time-component}
    to obtain `hour`{.variable}, `minute`{.variable}, and
    `second`{.variable}. If this returns nothing, then fail.

6.  If `position`{.variable} is beyond the end of `input`{.variable},
    then fail.

7.  [Parse a time-zone offset
    component](#parse-a-time-zone-offset-component){#global-dates-and-times:parse-a-time-zone-offset-component}
    to obtain `timezone`{.variable}~`hours`{.variable}~ and
    `timezone`{.variable}~`minutes`{.variable}~. If this returns
    nothing, then fail.

8.  If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

9.  Let `time`{.variable} be the moment in time at year
    `year`{.variable}, month `month`{.variable}, day `day`{.variable},
    hours `hour`{.variable}, minute `minute`{.variable}, second
    `second`{.variable}, subtracting
    `timezone`{.variable}~`hours`{.variable}~ hours and
    `timezone`{.variable}~`minutes`{.variable}~ minutes. That moment in
    time is a moment in the UTC time zone.

10. Let `timezone`{.variable} be
    `timezone`{.variable}~`hours`{.variable}~ hours and
    `timezone`{.variable}~`minutes`{.variable}~ minutes from UTC.

11. Return `time`{.variable} and `timezone`{.variable}.

##### [2.3.5.8]{.secno} Weeks[](#weeks){.self-link}

A [week]{#concept-week .dfn} consists of a week-year number and a week
number representing a seven-day period starting on a Monday. Each
week-year in this calendaring system has either 52 or 53 such seven-day
periods, as defined below. The seven-day period starting on the
Gregorian date Monday December 29th 1969 (1969-12-29) is defined as week
number 1 in week-year 1970. Consecutive weeks are numbered sequentially.
The week before the number 1 week in a week-year is the last week in the
previous week-year, and vice versa.
[\[GREGORIAN\]](references.html#refsGREGORIAN)

A week-year with a number `year`{.variable} has 53 weeks if it
corresponds to either a year `year`{.variable} in the [proleptic
Gregorian
calendar](#proleptic-gregorian-calendar){#weeks:proleptic-gregorian-calendar}
that has a Thursday as its first day (January 1st), or a year
`year`{.variable} in the [proleptic Gregorian
calendar](#proleptic-gregorian-calendar){#weeks:proleptic-gregorian-calendar-2}
that has a Wednesday as its first day (January 1st) and where
`year`{.variable} is a number divisible by 400, or a number divisible by
4 but not by 100. All other week-years have 52 weeks.

The [week number of the last day]{#week-number-of-the-last-day .dfn} of
a week-year with 53 weeks is 53; the week number of the last day of a
week-year with 52 weeks is 52.

The week-year number of a particular day can be different than the
number of the year that contains that day in the [proleptic Gregorian
calendar](#proleptic-gregorian-calendar){#weeks:proleptic-gregorian-calendar-3}.
The first week in a week-year `y`{.variable} is the week that contains
the first Thursday of the Gregorian year `y`{.variable}.

For modern purposes, a [week](#concept-week){#weeks:concept-week} as
defined here is equivalent to ISO weeks as defined in ISO 8601.
[\[ISO8601\]](references.html#refsISO8601)

A string is a [valid week string]{#valid-week-string .dfn} representing
a week-year `year`{.variable} and week `week`{.variable} if it consists
of the following components in the given order:

1.  Four or more [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#weeks:ascii-digits
    x-internal="ascii-digits"}, representing `year`{.variable}, where
    `year`{.variable} \> 0
2.  A U+002D HYPHEN-MINUS character (-)
3.  A U+0057 LATIN CAPITAL LETTER W character (W)
4.  Two [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#weeks:ascii-digits-2
    x-internal="ascii-digits"}, representing the week `week`{.variable},
    in the range 1 ≤ `week`{.variable} ≤ `maxweek`{.variable}, where
    `maxweek`{.variable} is the [week number of the last
    day](#week-number-of-the-last-day){#weeks:week-number-of-the-last-day}
    of week-year `year`{.variable}

The rules to [parse a week string]{#parse-a-week-string .dfn} are as
follows. This will return either a week-year number and week number, or
nothing. If at any point the algorithm says that it \"fails\", this
means that it is aborted at that point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#weeks:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#weeks:ascii-digits-3
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not at least
    four characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `year`{.variable} be that
    number.

4.  If `year`{.variable} is not a number greater than zero, then fail.

5.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+002D
    HYPHEN-MINUS character, then fail. Otherwise, move
    `position`{.variable} forwards one character.

6.  If `position`{.variable} is beyond the end of `input`{.variable} or
    if the character at `position`{.variable} is not a U+0057 LATIN
    CAPITAL LETTER W character (W), then fail. Otherwise, move
    `position`{.variable} forwards one character.

7.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#weeks:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#weeks:ascii-digits-4
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}. If the collected sequence is not exactly two
    characters long, then fail. Otherwise, interpret the resulting
    sequence as a base-ten integer. Let `week`{.variable} be that
    number.

8.  Let `maxweek`{.variable} be the [week number of the last
    day](#week-number-of-the-last-day){#weeks:week-number-of-the-last-day-2}
    of year `year`{.variable}.

9.  If `week`{.variable} is not a number in the range
    1 ≤ `week`{.variable} ≤ `maxweek`{.variable}, then fail.

10. If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

11. Return the week-year number `year`{.variable} and the week number
    `week`{.variable}.

##### [2.3.5.9]{.secno} Durations[](#durations){.self-link}

A [duration]{#concept-duration .dfn} consists of a number of seconds.

Since months and seconds are not comparable (a month is not a precise
number of seconds, but is instead a period whose exact length depends on
the precise day from which it is measured) a
[duration](#concept-duration){#durations:concept-duration} as defined in
this specification cannot include months (or years, which are equivalent
to twelve months). Only durations that describe a specific number of
seconds can be described.

A string is a [valid duration string]{#valid-duration-string .dfn}
representing a
[duration](#concept-duration){#durations:concept-duration-2}
`t`{.variable} if it consists of either of the following:

- A literal U+0050 LATIN CAPITAL LETTER P character followed by one or
  more of the following subcomponents, in the order given, where the
  number of days, hours, minutes, and seconds corresponds to the same
  number of seconds as in `t`{.variable}:

  1.  One or more [ASCII
      digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits
      x-internal="ascii-digits"} followed by a U+0044 LATIN CAPITAL
      LETTER D character, representing a number of days.

  2.  A U+0054 LATIN CAPITAL LETTER T character followed by one or more
      of the following subcomponents, in the order given:

      1.  One or more [ASCII
          digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-2
          x-internal="ascii-digits"} followed by a U+0048 LATIN CAPITAL
          LETTER H character, representing a number of hours.

      2.  One or more [ASCII
          digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-3
          x-internal="ascii-digits"} followed by a U+004D LATIN CAPITAL
          LETTER M character, representing a number of minutes.

      3.  The following components:

          1.  One or more [ASCII
              digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-4
              x-internal="ascii-digits"}, representing a number of
              seconds.

          2.  Optionally, a U+002E FULL STOP character (.) followed by
              one, two, or three [ASCII
              digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-5
              x-internal="ascii-digits"}, representing a fraction of a
              second.

          3.  A U+0053 LATIN CAPITAL LETTER S character.

  This, as with a number of other date- and time-related microsyntaxes
  defined in this specification, is based on one of the formats defined
  in ISO 8601. [\[ISO8601\]](references.html#refsISO8601)

- One or more [duration time
  components](#duration-time-component){#durations:duration-time-component},
  each with a different [duration time component
  scale](#duration-time-component-scale){#durations:duration-time-component-scale},
  in any order; the sum of the represented seconds being equal to the
  number of seconds in `t`{.variable}.

  A [duration time component]{#duration-time-component .dfn} is a string
  consisting of the following components:

  1.  Zero or more [ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#durations:space-characters
      x-internal="space-characters"}.

  2.  One or more [ASCII
      digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-6
      x-internal="ascii-digits"}, representing a number of time units,
      scaled by the [duration time component
      scale](#duration-time-component-scale){#durations:duration-time-component-scale-2}
      specified (see below) to represent a number of seconds.

  3.  If the [duration time component
      scale](#duration-time-component-scale){#durations:duration-time-component-scale-3}
      specified is 1 (i.e. the units are seconds), then, optionally, a
      U+002E FULL STOP character (.) followed by one, two, or three
      [ASCII
      digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-7
      x-internal="ascii-digits"}, representing a fraction of a second.

  4.  Zero or more [ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#durations:space-characters-2
      x-internal="space-characters"}.

  5.  One of the following characters, representing the [duration time
      component scale]{#duration-time-component-scale .dfn} of the time
      unit used in the numeric part of the [duration time
      component](#duration-time-component){#durations:duration-time-component-2}:

      U+0057 LATIN CAPITAL LETTER W character\
      U+0077 LATIN SMALL LETTER W character
      :   Weeks. The scale is 604800.

      U+0044 LATIN CAPITAL LETTER D character\
      U+0064 LATIN SMALL LETTER D character
      :   Days. The scale is 86400.

      U+0048 LATIN CAPITAL LETTER H character\
      U+0068 LATIN SMALL LETTER H character
      :   Hours. The scale is 3600.

      U+004D LATIN CAPITAL LETTER M character\
      U+006D LATIN SMALL LETTER M character
      :   Minutes. The scale is 60.

      U+0053 LATIN CAPITAL LETTER S character\
      U+0073 LATIN SMALL LETTER S character
      :   Seconds. The scale is 1.

  6.  Zero or more [ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#durations:space-characters-3
      x-internal="space-characters"}.

  This is not based on any of the formats in ISO 8601. It is intended to
  be a more human-readable alternative to the ISO 8601 duration format.

The rules to [parse a duration string]{#parse-a-duration-string .dfn}
are as follows. This will return either a
[duration](#concept-duration){#durations:concept-duration-3} or nothing.
If at any point the algorithm says that it \"fails\", this means that it
is aborted at that point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Let `months`{.variable}, `seconds`{.variable}, and
    `component count`{.variable} all be zero.

4.  Let `M-disambiguator`{.variable} be *minutes*.

    This flag\'s other value is *months*. It is used to disambiguate the
    \"M\" unit in ISO8601 durations, which use the same unit for months
    and minutes. Months are not allowed, but are parsed for future
    compatibility and to avoid misinterpreting ISO8601 durations that
    would be valid in other contexts.

5.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

6.  If `position`{.variable} is past the end of `input`{.variable}, then
    fail.

7.  If the character in `input`{.variable} pointed to by
    `position`{.variable} is a U+0050 LATIN CAPITAL LETTER P character,
    then advance `position`{.variable} to the next character, set
    `M-disambiguator`{.variable} to *months*, and [skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace-2
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

8.  While true:

    1.  Let `units`{.variable} be undefined. It will be assigned one of
        the following values: *years*, *months*, *weeks*, *days*,
        *hours*, *minutes*, and *seconds*.

    2.  Let `next character`{.variable} be undefined. It is used to
        process characters from the `input`{.variable}.

    3.  If `position`{.variable} is past the end of `input`{.variable},
        then break.

    4.  If the character in `input`{.variable} pointed to by
        `position`{.variable} is a U+0054 LATIN CAPITAL LETTER T
        character, then advance `position`{.variable} to the next
        character, set `M-disambiguator`{.variable} to *minutes*, [skip
        ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace-3
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}, and
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#durations:continue
        x-internal="continue"}.

    5.  Set `next character`{.variable} to the character in
        `input`{.variable} pointed to by `position`{.variable}.

    6.  If `next character`{.variable} is a U+002E FULL STOP character
        (.), then let `N`{.variable} equal zero. (Do not advance
        `position`{.variable}. That is taken care of below.)

        Otherwise, if `next character`{.variable} is an [ASCII
        digit](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-8
        x-internal="ascii-digits"}, then [collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#durations:collect-a-sequence-of-code-points
        x-internal="collect-a-sequence-of-code-points"} that are [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-9
        x-internal="ascii-digits"} from `input`{.variable} given
        `position`{.variable}, interpret the resulting sequence as a
        base-ten integer, and let `N`{.variable} be that number.

        Otherwise, `next character`{.variable} is not part of a number;
        fail.

    7.  If `position`{.variable} is past the end of `input`{.variable},
        then fail.

    8.  Set `next character`{.variable} to the character in
        `input`{.variable} pointed to by `position`{.variable}, and this
        time advance `position`{.variable} to the next character. (If
        `next character`{.variable} was a U+002E FULL STOP character (.)
        before, it will still be that character this time.)

    9.  If `next character`{.variable} is U+002E (.), then:

        1.  [Collect a sequence of code
            points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#durations:collect-a-sequence-of-code-points-2
            x-internal="collect-a-sequence-of-code-points"} that are
            [ASCII
            digits](https://infra.spec.whatwg.org/#ascii-digit){#durations:ascii-digits-10
            x-internal="ascii-digits"} from `input`{.variable} given
            `position`{.variable}. Let `s`{.variable} be the resulting
            sequence.

        2.  If `s`{.variable} is the empty string, then fail.

        3.  Let `length`{.variable} be the number of characters in
            `s`{.variable}.

        4.  Let `fraction`{.variable} be the result of interpreting
            `s`{.variable} as a base-ten integer, and then dividing that
            number by 10^`length`{.variable}^.

        5.  Increment `N`{.variable} by `fraction`{.variable}.

        6.  [Skip ASCII
            whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace-4
            x-internal="skip-ascii-whitespace"} within
            `input`{.variable} given `position`{.variable}.

        7.  If `position`{.variable} is past the end of
            `input`{.variable}, then fail.

        8.  Set `next character`{.variable} to the character in
            `input`{.variable} pointed to by `position`{.variable}, and
            advance `position`{.variable} to the next character.

        9.  If `next character`{.variable} is neither a U+0053 LATIN
            CAPITAL LETTER S character nor a U+0073 LATIN SMALL LETTER S
            character, then fail.

        10. Set `units`{.variable} to *seconds*.

        Otherwise:

        1.  If `next character`{.variable} is [ASCII
            whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#durations:space-characters-4
            x-internal="space-characters"}, then [skip ASCII
            whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace-5
            x-internal="skip-ascii-whitespace"} within
            `input`{.variable} given `position`{.variable}, set
            `next character`{.variable} to the character in
            `input`{.variable} pointed to by `position`{.variable}, and
            advance `position`{.variable} to the next character.

        2.  If `next character`{.variable} is a U+0059 LATIN CAPITAL
            LETTER Y character, or a U+0079 LATIN SMALL LETTER Y
            character, set `units`{.variable} to *years* and set
            `M-disambiguator`{.variable} to *months*.

            If `next character`{.variable} is a U+004D LATIN CAPITAL
            LETTER M character or a U+006D LATIN SMALL LETTER M
            character, and `M-disambiguator`{.variable} is *months*,
            then set `units`{.variable} to *months*.

            If `next character`{.variable} is a U+0057 LATIN CAPITAL
            LETTER W character or a U+0077 LATIN SMALL LETTER W
            character, set `units`{.variable} to *weeks* and set
            `M-disambiguator`{.variable} to *minutes*.

            If `next character`{.variable} is a U+0044 LATIN CAPITAL
            LETTER D character or a U+0064 LATIN SMALL LETTER D
            character, set `units`{.variable} to *days* and set
            `M-disambiguator`{.variable} to *minutes*.

            If `next character`{.variable} is a U+0048 LATIN CAPITAL
            LETTER H character or a U+0068 LATIN SMALL LETTER H
            character, set `units`{.variable} to *hours* and set
            `M-disambiguator`{.variable} to *minutes*.

            If `next character`{.variable} is a U+004D LATIN CAPITAL
            LETTER M character or a U+006D LATIN SMALL LETTER M
            character, and `M-disambiguator`{.variable} is *minutes*,
            then set `units`{.variable} to *minutes*.

            If `next character`{.variable} is a U+0053 LATIN CAPITAL
            LETTER S character or a U+0073 LATIN SMALL LETTER S
            character, set `units`{.variable} to *seconds* and set
            `M-disambiguator`{.variable} to *minutes*.

            Otherwise, if `next character`{.variable} is none of the
            above characters, then fail.

    10. Increment `component count`{.variable}.

    11. Let `multiplier`{.variable} be 1.

    12. If `units`{.variable} is *years*, multiply
        `multiplier`{.variable} by 12 and set `units`{.variable} to
        *months*.

    13. If `units`{.variable} is *months*, add the product of
        `N`{.variable} and `multiplier`{.variable} to
        `months`{.variable}.

        Otherwise:

        1.  If `units`{.variable} is *weeks*, multiply
            `multiplier`{.variable} by 7 and set `units`{.variable} to
            *days*.

        2.  If `units`{.variable} is *days*, multiply
            `multiplier`{.variable} by 24 and set `units`{.variable} to
            *hours*.

        3.  If `units`{.variable} is *hours*, multiply
            `multiplier`{.variable} by 60 and set `units`{.variable} to
            *minutes*.

        4.  If `units`{.variable} is *minutes*, multiply
            `multiplier`{.variable} by 60 and set `units`{.variable} to
            *seconds*.

        5.  Forcibly, `units`{.variable} is now *seconds*. Add the
            product of `N`{.variable} and `multiplier`{.variable} to
            `seconds`{.variable}.

    14. [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#durations:skip-ascii-whitespace-6
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

9.  If `component count`{.variable} is zero, fail.

10. If `months`{.variable} is not zero, fail.

11. Return the
    [duration](#concept-duration){#durations:concept-duration-4}
    consisting of `seconds`{.variable} seconds.

##### [2.3.5.10]{.secno} Vaguer moments in time[](#vaguer-moments-in-time){.self-link}

A string is a [valid date string with optional
time]{#valid-date-string-with-optional-time .dfn} if it is also one of
the following:

- A [valid date
  string](#valid-date-string){#vaguer-moments-in-time:valid-date-string}
- A [valid global date and time
  string](#valid-global-date-and-time-string){#vaguer-moments-in-time:valid-global-date-and-time-string}

------------------------------------------------------------------------

The rules to [parse a date or time string]{#parse-a-date-or-time-string
.dfn} are as follows. The algorithm will return either a
[date](#concept-date){#vaguer-moments-in-time:concept-date}, a
[time](#concept-time){#vaguer-moments-in-time:concept-time}, a [global
date and
time](#concept-datetime){#vaguer-moments-in-time:concept-datetime}, or
nothing. If at any point the algorithm says that it \"fails\", this
means that it is aborted at that point and returns nothing.

1.  Let `input`{.variable} be the string being parsed.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  Set `start position`{.variable} to the same position as
    `position`{.variable}.

4.  Set the `date present`{.variable} and `time present`{.variable}
    flags to true.

5.  [Parse a date
    component](#parse-a-date-component){#vaguer-moments-in-time:parse-a-date-component}
    to obtain `year`{.variable}, `month`{.variable}, and
    `day`{.variable}. If this fails, then set the
    `date present`{.variable} flag to false.

6.  If `date present`{.variable} is true, and `position`{.variable} is
    not beyond the end of `input`{.variable}, and the character at
    `position`{.variable} is either a U+0054 LATIN CAPITAL LETTER T
    character (T) or a U+0020 SPACE character, then advance
    `position`{.variable} to the next character in `input`{.variable}.

    Otherwise, if `date present`{.variable} is true, and either
    `position`{.variable} is beyond the end of `input`{.variable} or the
    character at `position`{.variable} is neither a U+0054 LATIN CAPITAL
    LETTER T character (T) nor a U+0020 SPACE character, then set
    `time present`{.variable} to false.

    Otherwise, if `date present`{.variable} is false, set
    `position`{.variable} back to the same position as
    `start position`{.variable}.

7.  If the `time present`{.variable} flag is true, then [parse a time
    component](#parse-a-time-component){#vaguer-moments-in-time:parse-a-time-component}
    to obtain `hour`{.variable}, `minute`{.variable}, and
    `second`{.variable}. If this returns nothing, then fail.

8.  If the `date present`{.variable} and `time present`{.variable} flags
    are both true, but `position`{.variable} is beyond the end of
    `input`{.variable}, then fail.

9.  If the `date present`{.variable} and `time present`{.variable} flags
    are both true, [parse a time-zone offset
    component](#parse-a-time-zone-offset-component){#vaguer-moments-in-time:parse-a-time-zone-offset-component}
    to obtain `timezone`{.variable}~`hours`{.variable}~ and
    `timezone`{.variable}~`minutes`{.variable}~. If this returns
    nothing, then fail.

10. If `position`{.variable} is *not* beyond the end of
    `input`{.variable}, then fail.

11. If the `date present`{.variable} flag is true and the
    `time present`{.variable} flag is false, then let `date`{.variable}
    be the date with year `year`{.variable}, month `month`{.variable},
    and day `day`{.variable}, and return `date`{.variable}.

    Otherwise, if the `time present`{.variable} flag is true and the
    `date present`{.variable} flag is false, then let `time`{.variable}
    be the time with hour `hour`{.variable}, minute `minute`{.variable},
    and second `second`{.variable}, and return `time`{.variable}.

    Otherwise, let `time`{.variable} be the moment in time at year
    `year`{.variable}, month `month`{.variable}, day `day`{.variable},
    hours `hour`{.variable}, minute `minute`{.variable}, second
    `second`{.variable}, subtracting
    `timezone`{.variable}~`hours`{.variable}~ hours and
    `timezone`{.variable}~`minutes`{.variable}~ minutes, that moment in
    time being a moment in the UTC time zone; let `timezone`{.variable}
    be `timezone`{.variable}~`hours`{.variable}~ hours and
    `timezone`{.variable}~`minutes`{.variable}~ minutes from UTC; and
    return `time`{.variable} and `timezone`{.variable}.

#### [2.3.6]{.secno} Legacy colors[](#colours){.self-link} {#colours}

Some obsolete legacy attributes parse colors using the [rules for
parsing a legacy color value]{#rules-for-parsing-a-legacy-colour-value
.dfn}, given a string `input`{.variable}. They will return either a CSS
color or failure.

1.  If `input`{.variable} is the empty string, then return failure.

2.  [Strip leading and trailing ASCII
    whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#colours:strip-leading-and-trailing-ascii-whitespace
    x-internal="strip-leading-and-trailing-ascii-whitespace"} from
    `input`{.variable}.

3.  If `input`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#colours:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for \"`transparent`\",
    then return failure.

4.  If `input`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#colours:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for one of the [named
    colors](https://drafts.csswg.org/css-color/#named-color){#colours:named-color
    x-internal="named-color"}, then return the CSS color corresponding
    to that keyword. [\[CSSCOLOR\]](references.html#refsCSSCOLOR)

    [CSS2 System Colors](https://www.w3.org/TR/css3-color/#css2-system)
    are not recognized.

5.  If `input`{.variable}\'s [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length
    x-internal="code-point-length"} is four, and the first character in
    `input`{.variable} is U+0023 (#), and the last three characters of
    `input`{.variable} are all [ASCII hex
    digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#colours:ascii-hex-digits
    x-internal="ascii-hex-digits"}, then:

    1.  Let `result`{.variable} be a CSS color.

    2.  Interpret the second character of `input`{.variable} as a
        hexadecimal digit; let the red component of `result`{.variable}
        be the resulting number multiplied by 17.

    3.  Interpret the third character of `input`{.variable} as a
        hexadecimal digit; let the green component of
        `result`{.variable} be the resulting number multiplied by 17.

    4.  Interpret the fourth character of `input`{.variable} as a
        hexadecimal digit; let the blue component of `result`{.variable}
        be the resulting number multiplied by 17.

    5.  Return `result`{.variable}.

6.  Replace any [code
    points](https://infra.spec.whatwg.org/#code-point){#colours:code-point
    x-internal="code-point"} greater than U+FFFF in `input`{.variable}
    (i.e., any characters that are not in the basic multilingual plane)
    with \"`00`\".

7.  If `input`{.variable}\'s [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length-2
    x-internal="code-point-length"} is greater than 128, truncate
    `input`{.variable}, leaving only the first 128 characters.

8.  If the first character in `input`{.variable} is U+0023 (#), then
    remove it.

9.  Replace any character in `input`{.variable} that is not an [ASCII
    hex
    digit](https://infra.spec.whatwg.org/#ascii-hex-digit){#colours:ascii-hex-digits-2
    x-internal="ascii-hex-digits"} with U+0030 (0).

10. While `input`{.variable}\'s [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length-3
    x-internal="code-point-length"} is zero or not a multiple of three,
    append U+0030 (0) to `input`{.variable}.

11. Split `input`{.variable} into three strings of equal [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length-4
    x-internal="code-point-length"}, to obtain three components. Let
    `length`{.variable} be the [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length-5
    x-internal="code-point-length"} that all of those components have
    (one third the [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#colours:code-point-length-6
    x-internal="code-point-length"} of `input`{.variable}).

12. If `length`{.variable} is greater than 8, then remove the leading
    `length`{.variable}-8 characters in each component, and let
    `length`{.variable} be 8.

13. While `length`{.variable} is greater than two and the first
    character in each component is U+0030 (0), remove that character and
    reduce `length`{.variable} by one.

14. If `length`{.variable} is *still* greater than two, truncate each
    component, leaving only the first two characters in each.

15. Let `result`{.variable} be a CSS color.

16. Interpret the first component as a hexadecimal number; let the red
    component of `result`{.variable} be the resulting number.

17. Interpret the second component as a hexadecimal number; let the
    green component of `result`{.variable} be the resulting number.

18. Interpret the third component as a hexadecimal number; let the blue
    component of `result`{.variable} be the resulting number.

19. Return `result`{.variable}.

#### [2.3.7]{.secno} Space-separated tokens[](#space-separated-tokens){.self-link}

A [set of space-separated tokens]{#set-of-space-separated-tokens .dfn
export=""} is a string containing zero or more words (known as tokens)
separated by one or more [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#space-separated-tokens:space-characters
x-internal="space-characters"}, where words consist of any string of one
or more characters, none of which are [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#space-separated-tokens:space-characters-2
x-internal="space-characters"}.

A string containing a [set of space-separated
tokens](#set-of-space-separated-tokens){#space-separated-tokens:set-of-space-separated-tokens}
may have leading or trailing [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#space-separated-tokens:space-characters-3
x-internal="space-characters"}.

An [unordered set of unique space-separated
tokens]{#unordered-set-of-unique-space-separated-tokens .dfn export=""}
is a [set of space-separated
tokens](#set-of-space-separated-tokens){#space-separated-tokens:set-of-space-separated-tokens-2}
where none of the tokens are duplicated.

An [ordered set of unique space-separated
tokens]{#ordered-set-of-unique-space-separated-tokens .dfn} is a [set of
space-separated
tokens](#set-of-space-separated-tokens){#space-separated-tokens:set-of-space-separated-tokens-3}
where none of the tokens are duplicated but where the order of the
tokens is meaningful.

[Sets of space-separated
tokens](#set-of-space-separated-tokens){#space-separated-tokens:set-of-space-separated-tokens-4}
sometimes have a defined set of allowed values. When a set of allowed
values is defined, the tokens must all be from that list of allowed
values; other values are non-conforming. If no such set of allowed
values is provided, then all values are conforming.

How tokens in a [set of space-separated
tokens](#set-of-space-separated-tokens){#space-separated-tokens:set-of-space-separated-tokens-5}
are to be compared (e.g. case-sensitively or not) is defined on a
per-set basis.

#### [2.3.8]{.secno} Comma-separated tokens[](#comma-separated-tokens){.self-link}

A [set of comma-separated tokens]{#set-of-comma-separated-tokens .dfn}
is a string containing zero or more tokens each separated from the next
by a single U+002C COMMA character (,), where tokens consist of any
string of zero or more characters, neither beginning nor ending with
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#comma-separated-tokens:space-characters
x-internal="space-characters"}, nor containing any U+002C COMMA
characters (,), and optionally surrounded by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#comma-separated-tokens:space-characters-2
x-internal="space-characters"}.

For instance, the string \"` a ,b,,d d `\" consists of four tokens:
\"a\", \"b\", the empty string, and \"d d\". Leading and trailing
whitespace around each token doesn\'t count as part of the token, and
the empty string can be a token.

[Sets of comma-separated
tokens](#set-of-comma-separated-tokens){#comma-separated-tokens:set-of-comma-separated-tokens}
sometimes have further restrictions on what consists a valid token. When
such restrictions are defined, the tokens must all fit within those
restrictions; other values are non-conforming. If no such restrictions
are specified, then all values are conforming.

#### [2.3.9]{.secno} References[](#syntax-references){.self-link} {#syntax-references}

A [valid hash-name reference]{#valid-hash-name-reference .dfn} to an
element of type `type`{.variable} is a string consisting of a U+0023
NUMBER SIGN character (#) followed by a string which exactly matches the
value of the `name` attribute of an element with type `type`{.variable}
in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#syntax-references:tree
x-internal="tree"}.

The [rules for parsing a hash-name
reference]{#rules-for-parsing-a-hash-name-reference .dfn} to an element
of type `type`{.variable}, given a context node `scope`{.variable}, are
as follows:

1.  If the string being parsed does not contain a U+0023 NUMBER SIGN
    character, or if the first such character in the string is the last
    character in the string, then return null.

2.  Let `s`{.variable} be the string from the character immediately
    after the first U+0023 NUMBER SIGN character in the string being
    parsed up to the end of that string.

3.  Return the first element of type `type`{.variable} in
    `scope`{.variable}\'s
    [tree](https://dom.spec.whatwg.org/#concept-tree){#syntax-references:tree-2
    x-internal="tree"}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#syntax-references:tree-order
    x-internal="tree-order"}, that has an
    [`id`{#syntax-references:the-id-attribute}](dom.html#the-id-attribute)
    or `name` attribute whose value is `s`{.variable}, or null if there
    is no such element.

    Although
    [`id`{#syntax-references:the-id-attribute-2}](dom.html#the-id-attribute)
    attributes are accounted for when parsing, they are not used in
    determining whether a value is a [*valid* hash-name
    reference](#valid-hash-name-reference){#syntax-references:valid-hash-name-reference}.
    That is, a hash-name reference that refers to an element based on
    [`id`{#syntax-references:the-id-attribute-3}](dom.html#the-id-attribute)
    is a conformance error (unless that element also has a `name`
    attribute with the same value).

#### [2.3.10]{.secno} Media queries[](#mq){.self-link} {#mq}

A string is a [valid media query list]{#valid-media-query-list .dfn} if
it matches the `<media-query-list>` production of Media Queries.
[\[MQ\]](references.html#refsMQ)

A string [matches the environment]{#matches-the-environment .dfn} of the
user if it is the empty string, a string consisting of only [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#mq:space-characters
x-internal="space-characters"}, or is a media query list that matches
the user\'s environment according to the definitions given in Media
Queries. [\[MQ\]](references.html#refsMQ)

#### [2.3.11]{.secno} Unique internal values[](#unique-values){.self-link} {#unique-values}

A [unique internal value]{#unique-internal-value .dfn} is a value that
is serializable, comparable by value, and never exposed to script.

To create a [new unique internal value]{#new-unique-internal-value
.dfn}, return a [unique internal
value](#unique-internal-value){#unique-values:unique-internal-value}
that has never previously been returned by this algorithm.

[← 2 Common infrastructure](infrastructure.html) --- [Table of
Contents](./) --- [2.4 URLs →](urls-and-fetching.html)
