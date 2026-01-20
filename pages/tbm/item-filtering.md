# Item Filtering
You can open the item filter with the "Edit Item Filter..." button in the configuration GUI. You can find the button right
below the tray.

You will see a GUI like this:

![Item Filter Menu](/static/img/tbm/item_filter_menu.png)

With the item filter, you can filter out any items you do not need. It supports both regular expressions and standard glob:

If you click "Add Filter...", a menu like this will pop up:
![Item Filter Add Menu](/static/img/tbm/item_filter_edit_menu.png)

In the first textbox, you can add the predicate. This is either a glob or a regular expression:

Here's an example filtering out **all** wool blocks:
- Regular Expressions: e.g. ^minecraft:.+_wool$
- Glob: e.g. minecraft:*_wool

Glob is **exclusive**. You can only specify items you do **NOT** want. You cannot only include a item of one type and exclude all others
with one line. The NOT-Operator (!) is NOT supported.

Regular expressions are also mostly exclusive, although you can use a trick involving negative lookaheads to make the
filter exclude out everything else and leave only the thing you specified in the regex.

Here's an example to only leave "minecraft:(something)_wool" and exclude everything else: `^(?!minecraft:.+_wool$).*`

**If you use regular expressions, you _MUST_ tick the "Use Regular Expressions" checkbox**, otherwise it will be interpreted
as a glob and will do nothing.

By clicking on an item group, you can either add or remove it to or from the "Applied Groups" column.

Click save if you are happy, click discard if you want to give up.