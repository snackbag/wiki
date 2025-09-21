# Lone Contraptions

Let's make an airship. We first need to know what we want the airship to be. Of course! A contraption. Create
contraptions are extremely customizable and easy to move around. Precisely what we want. The next step is to define what
a contraption actually is:

- it's a type of contraption
- it's a class
- it's an entity
- it's something that needs to be rendered

## Terminology

## A New Contraption

As established, a contraption is a class, an entity and of course a type. Let's create a new class first and make it
extend `TranslatingContraption`. TranslatingContraption handles things that would otherwise be annoying to maintain
yourself, such as colliders and stabilization.

TranslatingContraption requires an `assemble` method as well as a `getType` method.