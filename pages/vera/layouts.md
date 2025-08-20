# Layouts

Let's center everything. That's where layouts come in. There are two kinds of layouts:

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

![Darkened Minecraft world with centered white text displaying "hi there"](/static/img/vera-layout-centering-demo.png)

Simple. Ish. Suggestions are always welcome.

## Full Example
```java
package com.example.demo;

import net.snackbag.vera.Vera;
import net.snackbag.vera.core.VColor;
import net.snackbag.vera.core.VeraApp;
import net.snackbag.vera.event.VShortcut;
import net.snackbag.vera.flag.VLayoutAlignmentFlag;
import net.snackbag.vera.layout.VHLayout;
import net.snackbag.vera.layout.VVLayout;
import net.snackbag.vera.widget.VLabel;

public class DemoApplication extends VeraApp {
    private int clicks = 0;
    private VHLayout layout;
    private VVLayout centerLayout;

    @Override
    public void init() {
        new VShortcut(this, "escape", this::hide).alsoAdd();
        setBackgroundColor(VColor.black().withOpacity(0.2f));

        layout = new VHLayout(this, 0, 0);
        layout.alignment = VLayoutAlignmentFlag.CENTER;

        centerLayout = new VVLayout(this, 0, 0).alsoAddTo(layout);
        centerLayout.alignment = VLayoutAlignmentFlag.CENTER;

        VLabel label = new VLabel("hi there", this).alsoAddTo(centerLayout);
        label.modifyFontColor().rgb(VColor.white());
        label.onLeftClick(() -> {
            clicks++;
            label.setText("Clicks: " + clicks);
        });
    }

    @Override
    public void update() {
        super.update();

        setWidth(Vera.provider.getScreenWidth());
        setHeight(Vera.provider.getScreenHeight());

        layout.setSize(getWidth(), getHeight());
        centerLayout.setHeight(getHeight());
    }
}
```