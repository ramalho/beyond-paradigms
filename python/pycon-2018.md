# Is Python OO or functional? Is that the right question?

## Description

Java is object oriented and Haskell is functional. How about Python? Is it really OO with free-standing functions and such weak encapsulation? Python has lambdas and closures, but is it functional? Are these useful questions?

In the last 10 years a new approach to the study of programming languages has emerged: focus on features, not paradigms. This approach offers more direct, practical advice for programmers learning a new language, taking up coding idioms, and choosing suitable design patterns.

This talk will name some key language features, show how they affect the use of design patterns, and conclude with a refactoring guided by this new approach, producing simpler, more efficient code. Theory in practice.

## Audience

This talk does not require much Python knowledge. Knowing more than one programming language will be more helpful than knowing Python in depth. Some features — like iterators and metaclasses — will be mentioned with very brief, high level explanations. The aim is to encourage the audience to get up to speed on fundamental language features, not to dissect them here. My hope is that participants will put this sampling of theory in practice by adopting a pragmatic (ie. practical, not dogmatic) appreciation of clear, idiomatic code, and gaining a critical view of classic design patterns that may be overkill or unnecessary when we can leverage key features that Python offers.

## Outline

This outline is for 25' of presentation + 5' of Q&A. If given a 45' slot I will add 5 minutes to each of the 3 topics marked with *. This will raise the presentation time to 40' + 5' of Q&A. Those topics have tables or code that I can present at a high-level with a small sampling of details or dive deeper explaining more details. 

01:00 Intro

01:00 From assembly to Elixir via Python 

01:00 Programming Language Pragmatics (Scott. 2015) 

01:00 GCD in four languages

02:00 Variations of GCD in Python 

00:30 Scott's classifcation

00:30 Melo & Silva's classification

00:30 Language Categories in The Language List

01:00 Ontology is overrated

01:00 A better approach: teaching languages in the 21st century

00:30 Krishnamurthi's book

00:30 Kamin's book

03:00 6 features in 5 languages: a table*

01:00 Design Patterns are not universal

03:00 Norvig's critique of design pattterns*

03:00 Implementing the Strategy Pattern*

03:00 Refactoring to use functions as objects

00:30 Which is more idiomatic?

00:30 Neither syntax nor paradigms: features are the key 

00:30 Wrapping up: the value of solid foundations

25:00 TOTAL 

## Additional notes


I presented a very similar talk at GopherCon Brasil 2017 and it was well received. The slides (in Portuguese, with examples in Go) are here:

https://speakerdeck.com/ramalho/go-cade-o-paradigma

As it happens, Go and Python are prime examples of pragmatic language design -- Guido and the authors of Go are more interested in solving real problems than sticking to "paradigms".

I will use mostly the same slides, translated to English, except for the code examples which I will port from Go to idiomatic Python. I will also add a short motivating Python example early on, in the topic "Variations of GCD in Python".
