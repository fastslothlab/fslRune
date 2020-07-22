fslRune
=============
## fslRune.IsRuneNoTofu
* Tell you whether a rune/character can be displayed normally by the operating system.
* There is three type of rune/character can not be displayed:
    * Space . like code point 0x20(32). ' '
    * Font not exist. There is a question mark in the square box. (occupy two English letters width) like code point 0x378(888)  "͸"
    * A diamond-shaped black box with a question mark inside (occupy one English letter width) like code point 0xd800(55296) "�"
* Data come from chrome 84 on macOs 10.15.4(2020-03-24), older operation system will see some tofus.