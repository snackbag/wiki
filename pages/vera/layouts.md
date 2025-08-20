# Layouts

Let's add a proper click button to our game and center everything. That's where layouts come in. There are two kinds of
layouts:

1. `VVLayout` -> Aligns elements vertically
2. `VHLayout` -> Align elements horizontally

We'll need them both. One time the horizontal layout to center it horizontally and inside the vertical one to vertically
center it.
![Diagram showing an illustration of a white and a blue square, identified by "VHLayout" and "VVLayout"](/static/img/vera-layout-centering-diagram.png)

In code, that is

```java
// DemoApplication
private VHLayout layout;
private VVLayout centerLayout;

@Override
public void init() {
    // ...

    layout = new VHLayout(this, 0, 0);
    layout.alignment = VLayoutAlignmentFlag.CENTER;
    
    centerLayout = new VVLayout(this, 0, 0);
    centerLayout.alignment = VLayoutAlignmentFlag.CENTER;
    
    // ...
}

@Override
public void update() {
    // ...
    
    // set width AND height
    layout.setSize(getWidth(), getHeight());
    
    // only set height, let the widgets scale it automatically
    centerLayout.setHeight(getHeight());
}
```

and instead of using `.alsoAdd()` for our label, we use `.alsoAddTo(centerLayout)` to add it to the center layout.