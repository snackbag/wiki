# The Assembler

Allows you to assemble text quickly. You can create a new assembler with `NText.assemble()` and use `.build()` to create
a Minecraft `Text` object.

**Normal text**

```java
NText.assemble()
    .text("I'm assembled!")
    .build();
```

Now, don't jump away because it's so "long". That's why we have the basic `NText` QOL methods. With the assembler you
can do many cool things. These cool 'things' are called operations.

Let's apply a color operation to a translation key:

```java
NText.assemble()
    .text("block.minecraft.diamond_block", TextType.TRANSLATION)
    .color(Color.RED)
    .build();
```

As you can see, the operation is applied to the last `.text` element. If none is defined, an error is thrown. This also
allows you to stack multiple texts together. Look at this green text together with some red text:

```java
NText.assemble()
    .text("I'm green").color(Color.GREEN)
    .text(" and I'm red").color(Color.RED)
    .build();
```

## Hover events

In Minecraft, when you hover over text, there is a possibility you see a tooltip. These tooltips are called hover
events. Watch out: a hover event is still an operation, so make sure you created a normal text before.

**Text**

```java
NText.assemble()
    .text("Hover me!") // create normal text
    .hover("I'm a tooltip") // apply hover operation
    .build();
    
// or even styled tooltips
NText.assemble()
    .text("Hover me!")
    .hover(NText.of("Beautiful tooltip", Color.RED)
    .build();
```

**Items**\
Hover events can also show items. This is most often used in Minecraft's death messages.

```java
NText.assemble()
    .text("Hover to see an item")
    .hover(Items.DIAMOND) // or straight-up supply with an ItemStack
    .build();
```

**Entities**
BEGIN NOTE
**Note:** This hover event may result in unexpected behavior. It is recommended to review
the [Minecraft wiki](https://minecraft.wiki/w/Text_component_format#show_entity) first.
END NOTE

```java
NText.assemble()
    .text("Hover to see an entity")
    .hover(<entity>)
    .build();
```

## Click events

Setting up text-click events has been made simpler.

**Run commands**\
Since 1.19.1, Minecraft no longer allows click events make the player send chat messages, therefore putting a `/` in
front of the command is obsolete. This also includes any commands that may make the player say something (such as `/say`
or `/msg`). Neilon simplifies this behaviour and automatically adds a `/` if none is given.

```java
NText.assemble()
    .text("click me")
    .click(ClickType.RUN, "gamemode creative")
    .build();
```

**Copy to clipboard**

```java
NText.assemble()
    .text("click me")
    .click(ClickType.COPY, ":p")
    .build();
```

**Suggestions**\
Minecraft allows text to change the chat bar. That means, when clicking the text, it replaces the content with the
specified content. This is not limited to commands, this can be anything you want.

```java
NText.assemble()
    .text("click me")
    .click(ClickType.SUGGEST, "suggested text")
    .build();
```

**Open URLs**\
Allows the client to open a URL (confirmation window will be shown first). It is not essential to prefix the URL with
the `https://` protocol, as Neilon automatically adds it if none is specified.

```java
NText.assemble()
    .text("click me")
    .click(ClickType.URL, "github.com/snackbag/neilon")
    .build();
```

## Don't forget!

Don't forget that this is the assembler. It allows you to assemble things. Here's an example construction:

```java
NText.assemble()
    .text("[Epic Button]")
    .color(Color.RED)
    .hover("What will this do?")
    .click(ClickType.RUN, "kill @s")
    
    .text(" ")
    
    .text("[Another Button]")
    .color(Color.GREEN)
    .hover(NText.of("I suspect this won't be any better.", Color.GREEN))
    .click(ClickType.COPY, "i like trains")
    
    .build();
```

For reference, this is how it would look like when using vanilla code
```java
Text.literal("[Epic Button]")
    .styled(style -> style
        .withColor(Formatting.RED)
        .withHoverEvent(new HoverEvent(
            HoverEvent.Action.SHOW_TEXT,
            Text.literal("What will this do?")
        ))
        .withClickEvent(new ClickEvent(
            ClickEvent.Action.RUN_COMMAND,
            "kill @s"
        ))
    )
    
    .append(" ")
    
    .append(Text.literal("[Another Button]")
        .styled(style -> style
            .withColor(Formatting.GREEN)
            .withHoverEvent(new HoverEvent(
                HoverEvent.Action.SHOW_TEXT,
                Text.literal("I suspect this won't be any better.")
                    .styled(s -> s.withColor(Formatting.GREEN))
            ))
            .withClickEvent(new ClickEvent(
                ClickEvent.Action.COPY_TO_CLIPBOARD,
                "i like trains"
            ))
        )
    );
```