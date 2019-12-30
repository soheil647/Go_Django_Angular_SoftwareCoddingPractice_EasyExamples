"use strict";
var _a;
//Symbols()
var s = Symbol("First Symbol");
console.log(typeof s);
console.log(s.toString());
var s2 = Symbol("Test");
var s3 = Symbol("Test");
console.log(s2 === s3);
var s4 = Symbol.for('RegSymbol');
var s5 = Symbol.for('RegSymbol');
console.log(s4 === s5);
console.log(Symbol.keyFor(s4));
var fName = Symbol("First Name");
var ob1 = (_a = {},
    _a[fName] = "Chandler",
    _a);
console.log(Object.getOwnPropertyNames(ob1));
//Symbol Iterators
// for..of ---> itirator method ---> symbol.itrator
var str = "Hello";
var arr = [1, 2, 3];
var numb = 5;
var obj = { name: "Chandler" };
//console.log("For String -" + typeof str[Symbol.iterator]);
/*Itranable {
    [Symbol.itrantor]() : Iterator
    }

    Itrator{
        next() : IResultObj
}
IResultObj {
    value : any
    done : bool
}*/
var iterable = [1, 2, 3, 4];
function createIterator(array) {
    var count = 0;
    return {
        next: function () {
            return count < array.length ? { value: array[count++], done: false } : { value: undefined, done: true };
        }
    };
}
var myIterator = createIterator(iterable);
console.log(myIterator.next());
console.log(myIterator.next());
console.log(myIterator.next());
console.log(myIterator.next());
console.log(myIterator.next());
/*//Iterate over Objects
let People = {
    fName:"Monica",
    lName:"Geller"
};

People[Symbol.iterator] = function () {
    let properties = Object.keys(People);
    let count = 0;
    let isDone = false;
    let next = () => {
        if(count < properties.length){
            isDone : true
        }
        return{done:isDone,value: this[properties[count++]]}
    };
    return {next}
};
for (let p in People){
    console.log(p)
}*/
/*//Generators
function *CreateGenerator() {
    console.log("Before yield");
    yield 1;
    console.log("After yield 1");
    yield 2;
    console.log("After yield 2")
}
let MyGen = CreateGenerator();
console.log(MyGen.next());
console.log(MyGen.next());
console.log(MyGen.next());

let People = {
    fName:"Monica",
    lName:"Geller"
};

People[Symbol.iterator] = function* () {
    let properties = Object.keys(People);
    for (let t of properties) {
        yield this[t]
    }
};

for (let p in People){
    console.log(p);
    }
*/
//# sourceMappingURL=../TypeScript/build/morejs.js.map