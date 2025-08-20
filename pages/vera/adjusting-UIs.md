# Adjusting UIs

We've got the basics out of the way. Let's move stuff, add interactions and *colors*.

Moving is very simple. Use the `.move` method on a widget to adjust its position. For QOL there are two move methods.

1. `.move(int x, int y)` -> changes x and y accordingly
2. `.move(int all)` -> sets both x and y to the all argument

So if we want to move our label to x=10 y=10, we can just

```java
// DemoApplication#init

VLabel label = new VLabel("hi there", this).alsoAdd();
label.move(10);
```

Already looks magnitudes better. Since this is a clicking game, we should add the clicking functionality. Thankfully, we
can simply listen to the click event.

```java
private int clicks = 0;

@Override
public void init() {
    // ...

    label.onLeftClick(() -> {
        clicks++;
        label.setText("Clicks: " + clicks);
    });
}
```

Try it! When you left-click the label, it increments the `clicks` variable and adjusts the label's text.
