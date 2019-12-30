"use strict";
var MySet = Object.create(null);
MySet.id = 0;
if (MySet.id) {
    console.log("Id Exist");
}
var MyMap = Object.create(null);
MyMap.name = "Chandler";
var val = MyMap.name;
console.log(val);
MyMap[100] = "Hello";
console.log(MyMap["100"]);
var MSet = new Set();
var object1 = {};
var object2 = {};
MSet.add("Hello");
MSet.add(1);
MSet.add(object1);
MSet.add(object2);
console.log(MSet.size);
var newSet = new Set([1, 2, 3, 4, 4, 4, 5]);
console.log(newSet.size);
var chainSet = new Set().add("Hello").add("World");
console.log(chainSet.size);
console.log(newSet.has(2));
console.log(newSet.has(6));
newSet.delete(4);
console.log(newSet.has(4));
console.log(newSet.size);
//Weak Sets
var MyNewSet = new WeakSet();
var key = {};
MyNewSet.add(key);
console.log(MyNewSet.has(key));
//Maps
var myMap = new Map();
myMap.set("fname", "monica");
myMap.set("age", 30);
console.log(myMap.get("fname"));
var pob1 = {};
var pob2 = {};
myMap.set(pob1, 40);
myMap.set(pob2, 30);
console.log(myMap.get(pob1));
myMap.delete("fname");
console.log(myMap.size);
myMap.clear();
console.log(myMap.has("fname"));
//Itirate over maps
var newMap = new Map([
    ["fName", "Chandler"],
    ["lName", "Bing"]
]);
//???????? for of /////////
//forEach
var num = [2, 4, 6, 8];
function numFunc(element, index, array) {
    console.log("arr[" + index + "]= " + element);
}
num.forEach(numFunc);
var aMap = new Map([
    ["fName", "monica"],
    ["lName", "Geller"]
]);
function mapFunc(value, key, callingMap) {
    console.log(key + "" + value);
    console.log(aMap === callingMap);
}
aMap.forEach(mapFunc);
var aSet = new Set([1, 2, 3]);
function SetFunc(key, value, callingSet) {
    console.log(key + "" + value);
    console.log(aSet === callingSet);
}
aSet.forEach(SetFunc);
//WeakMaps
var tMap = new WeakMap();
var tob1 = {};
tMap.set(tob1, "Hello World");
console.log(tMap.get(tob1));
//# sourceMappingURL=../TypeScript/build/SetsAndMaps.js.map