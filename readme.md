gojs
====

The Aim of the gojs module is to provide a simple way to serve shared embedded js files across multiple server apps.

Some examples,

tallorwide.js - A simple test if a browser is tall or wide, changing the css class of an object to either main-tall or main-wide to, so that css can handle both wide and tall pages.


I needed these across multiple server apps, and couldn't work out where to put them, so I created this js library location

To use, run ```go-bindata``` in this library and access with

    
