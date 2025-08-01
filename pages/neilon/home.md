# Neilon

Neilon is a Fabric text library to aid working with Minecraft text.\
View the roadmap and a list of all features on GitHub.

## Setup

Any version of Neilon under 1.1 requires any version of Vera over 1.0, as it uses `VColor` in the backend. This has been
changed.

```groovy
repositories {
    // ...
    maven { url = "https://artifacts.snackbag.net/artifactory/libs-release/" }
}

dependencies {
    // ...
    modImplementation "net.snackbag.neilon:neilon-lib:1.1.1" // latest version
    
    // (for 1.1 or below)
    modImplementation "neilon-lib:neilon-lib:1.1.1"
}
```

You're good to go!

## Why is this necessary?

When modding normally, text is very abstract and convoluted, Neilon aims to make this easier. See this very basic
example:

```java
Text.literal("Hello"); // vanilla
NText.of("Hello"); // Neilon
```

Not much has changed. But once you get into more complex stuff, Neilon will make your life a lot easier. Such as colors
and hover events!

```java
// vanilla
Text.literal("Your Text").styled(style -> style
    .withColor(Formatting.RED)
    .withHoverEvent(new HoverEvent(
        HoverEvent.Action.SHOW_TEXT,
        Text.literal("Hovered Text").styled(s -> s.withColor(Formatting.RED)))
    )
);

// Neilon
NText.assemble()
    .text("Your Text")
    .color(Color.RED)
    .hover(NText.of("Hovered Text", Color.RED))
    .build();
```

now imagine how annoying it would be to add a click event in vanilla. The best part about Neilon is that it immediately
transforms into your expected Fabric `Text` object! No adapters necessary.
