
console.log("Console log module A 1");
import {fName as f,lName as l,obj} from "./moduleB.js";
import {default as FirstName} from "./moduleC.js"
import {greet,GreetClass} from "./moduleD.js";

console.log("Console log module A 2");

obj.name = "Ross";
console.log(obj.name);
console.log(` ${f} ${l} `);
console.log(FirstName);

greet("Hello World");
let g = new GreetClass();
g.greet();
