# Internationalization Instructions

This document describes how to add a new language to the client.

## How to add a new language

1.  Create a new file in the `client/src/locale` directory with the name of the language code (e.g. `en-GB.locale.ts` for British English).
2.  Copy the contents of `client/src/locale/en-US.locale.ts` into the new file.
3.  Translate the strings in the new file.
4.  Add the new language to the `client/src/locale/locale.ts` file. You can add it to the `translations` object, similar to how the other languages are added.
5.  Add the new language key (e.g. `en-GB`) to the `localeFlags` object with the corresponding flag emoji and label.
6.  Add the numeric system used by the new language to the `localeNumericDisplay` object. The key should be the language code (e.g. `en-GB`) and the value should be the numeric system (e.g. `arabic`). All the supported numeric systems are listed in the `client/src/locale/numeric-systems/` directory.

## How to add a new translation key

1.  Add the new key to the `client/src/locale/locale.ts` file, inside the exported `Locale` interface.
2.  Update all the existing language files with the new key.

## How to add a new numeric system

1.  Create a new file in the `client/src/locale/numeric-systems` directory with the name of the numeric system (e.g. `arabic.ts`).
2.  Create and export a function named like the numeric system (e.g. `export const arabic = (value) => { ... }`).
3.  The function should take a number `{ a | a ∈ Q }` as an argument and return a string with the number in the numeric system.
