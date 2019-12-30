let MySet = Object.create(null);
MySet.id = 0;
if(MySet.id){
    console.log("Id Exist")
}

let MyMap = Object.create(null);
MyMap.name = "Chandler";
let val = MyMap.name;
console.log(val);

MyMap[100] = "Hello";
console.log(MyMap["100"]);

let MSet = new Set();
let object1 = {};
let object2 = {};

MSet.add("Hello");
MSet.add(1);
MSet.add(object1);
MSet.add(object2);
console.log(MSet.size);

let newSet = new Set([1,2,3,4,4,4,5]);
console.log(newSet.size);

let chainSet = new Set().add("Hello").add("World");
console.log(chainSet.size);

console.log(newSet.has(2));
console.log(newSet.has(6));
newSet.delete(4);
console.log(newSet.has(4));
console.log(newSet.size);

//Weak Sets
let MyNewSet = new WeakSet();
let key = {};
MyNewSet.add(key);
console.log(MyNewSet.has(key));

//Maps
let myMap = new Map();

myMap.set("fname","monica");
myMap.set("age",30);

console.log(myMap.get("fname"));

let pob1 = {};
let pob2 = {};

myMap.set(pob1,40);
myMap.set(pob2,30);

console.log(myMap.get(pob1));
myMap.delete("fname");
console.log(myMap.size);
myMap.clear();

console.log(myMap.has("fname"));

//Itirate over maps
let newMap = new Map([
    ["fName" , "Chandler"],
    ["lName", "Bing"]
]);
//???????? for of /////////

//forEach
let num = [2,4,6,8];
function numFunc(element: number, index: number, array: any) {
    console.log("arr["+index+"]= " + element);
}
num.forEach(numFunc);

let aMap = new Map([
    ["fName" , "monica"],
    ["lName","Geller"]
]);
function mapFunc(value: string, key: string, callingMap: Map<any, any>){
    console.log(key+""+value);
    console.log(aMap === callingMap);
}
aMap.forEach(mapFunc);

let aSet = new Set([1,2,3]);
function SetFunc(key: number, value: number, callingSet: Set<number>){
    console.log(key+""+value);
    console.log(aSet === callingSet);
}
aSet.forEach(SetFunc);

//WeakMaps
let tMap = new WeakMap();
let tob1 = {};
tMap.set(tob1,"Hello World");
console.log(tMap.get(tob1));



