# Your First UI

Simply extend the `VeraApp` class and you're good to go. Oh, right. What even is an app? In Vera, you either draw on the
screen, or you use elements. In this tutorial, we'll look at the element part of Vera.

An element can be a layout or a widget. Basically any "object" on the screen that has a position and can be rendered. An
application is the holder for elements with a few customization options. Apps do not take over the entire screen and
therefore can be scaled and moved around. All initialization logic is recommended to be written inside the app's `init`
method.

In this tutorial we will create a little clicker game. Before we can do any clicking, we must first learn how to add
widgets to the screen. Vera makes this very easy, like so:

```java
public class DemoApplication extends VeraApp {
    @Override
    public void init() {
        new VLabel("hi there", this).alsoAdd();
    }
}
```

BEGIN NOTE
**Note:** Never forget the `.alsoAdd()` method. It is required due to legacy issues, which would require significant
breaking changes in order to adjust. The method notifies the app about the widget's existence and returns the same
widget back.
END NOTE

Render the UI on screen by simply initializing it and running the `.show()` method. In our case, it's
`new DemoApplication().show()`
![Minecraft game with text overlaying in a black font, displaying "hi there"](/static/img/vera-hi-there.png)

First step's done! But we can't exit.