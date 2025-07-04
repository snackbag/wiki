# Basic Text

Every type of text is supported by Neilon. You can create new text by using the methods from `NText`. Every method will
return a normal Minecraft `Text` object, meaning you can use it anywhere without a problem.

**Literal text**

```java
NText.of("literal text");

// or with quick colors
NText.of("literal text", VColor.MC_RED);
```

**Translations**

```java
NText.translation("block.minecraft.grass_block");

// or with quick colors
NText.translation("block.minecraft.grass_block", VColor.MC_RED);
```

**Keybindings**

```java
NText.keybinding("key.jump");

// you guessed it, with quick colors
NText.keybinding("key.jump", VColor.MC_RED);
```

This is fun, but you also want a little longer text. Maybe something like a message with a keybind within? Or different
colors throughout the text? Then you're ready for the assembler.