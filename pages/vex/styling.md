# Styling

Vex natively comes with a styling library. This is to ease the creation of animations and hover effects (transitions,
aka animations). In this part of the documentation you'll learn how to use the styling system.

## All base styles

**Style keys**

| Name               | Type      | Description                                               |
|--------------------|-----------|-----------------------------------------------------------|
| `color`            | `color`   | Color for any type of text                                |
| `background-color` | `color`   | Color for the background of a widget                      |
| `image-tint`       | `color`   | Tint for images. Default is white                         |
| `font-size`        | `int`     | Size for text in pixels                                   |
| `font-name`        | `string`  | Font for any text                                         |
| `letter-spacing`   | `int`     | Spacing between letters of text                           |
| `border-roundness` | `float32` | Roundness factor for widgets. 1 = circle, 0 = sharp edges |
| `border-width`     | `int`     | Size of border in pixels                                  |
| `border-color`     | `color`   | Color for widget border                                   |
| `border-segments`  | `int`     | Out of how many segments the border is made               |

---

**Style types**

BEGIN NOTE
**Note:** The value of a style key must be one of these values. If it isnt't, the program will panic and crash.
END NOTE

* Color (`color.RGBA{}` or any of the vex color methods)
* String
* Int (may also be `int8`, `int16`, `int32` or `int64`)
* Float32 (may also be `int` or `float64`)
* Float64 (may also be `int` or `float32`)

## Setting styles

You can simply set the style of a widget by using `widget.SetStyle("key", value)`. The value of the style must be one of
the style types. When you set a style, the value type gets reserved for the key. Whenever you try to set that exact key,
regardless of which widget, it must be in the style type you originally set it. If you do not want to set anything but
reserve a style type for a key you can use `Process.StyleSheet.ReserveType("key", extra.SVT<style type>)`

When you set a style (e.g. `float32`) and you give it one of the other possible values (e.g. `int`) it will
automatically convert it to `float32`

## Getting styles

You can simply get the value of a style (of a widget) by using `widget.`