# Styles

The UI looks very bland. Let's add a separate button to click on that increments our counter.

![Minecraft UI with text displaying "Clicks 4" as well as a button named "Click me" with a green border, all centered on the middle of the screen](/static/img/vera-style-click-button-demo.png)

## Full Example
```java
package com.example.demo;

import net.snackbag.vera.Vera;
import net.snackbag.vera.core.VColor;
import net.snackbag.vera.core.VCursorShape;
import net.snackbag.vera.core.VeraApp;
import net.snackbag.vera.event.VShortcut;
import net.snackbag.vera.flag.VLayoutAlignmentFlag;
import net.snackbag.vera.layout.VHLayout;
import net.snackbag.vera.layout.VVLayout;
import net.snackbag.vera.style.StyleState;
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

        centerLayout = new VVLayout(layout);
        centerLayout.alignment = VLayoutAlignmentFlag.START;

        VLabel label = new VLabel("Not clicked yet", this).alsoAddTo(centerLayout);
        label.modifyFontColor().rgb(VColor.white());

        VLabel button = new VLabel("Click me", this).alsoAddTo(centerLayout);

        button.modifyFontColor().rgb(VColor.of(95, 180, 0));
        button.setStyle("background-color", VColor.white());
        button.setStyle("padding", 4);
        button.setStyle("overlay", StyleState.HOVERED, VColor.white().withOpacity(0.5f));
        button.setStyle("border-size", 1);
        button.setStyle("border-color", VColor.of(95, 180, 0));
        button.setStyle("cursor", VCursorShape.POINTING_HAND);

        button.onLeftClick(() -> {
            clicks++;
            label.setText("Clicks: " + clicks);
        });

        new VShortcut(this, "leftctrl+d", () -> {
            System.out.println("Debug hit");
        }).alsoAdd();
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