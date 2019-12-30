//Symbols()
let s = Symbol("First Symbol");
console.log(typeof s);
console.log(s.toString());

let s2 = Symbol("Test");
let s3 = Symbol("Test");
console.log(s2 === s3);

let s4 = Symbol.for('RegSymbol');
let s5 = Symbol.for('RegSymbol');
console.log(s4===s5);
console.log(Symbol.keyFor(s4));

let fName = Symbol("First Name");
let ob1 = {
    [fName] : "Chandler"
};
console.log(Object.getOwnPropertyNames(ob1));

//Symbol Iterators
// for..of ---> itirator method ---> symbol.itrator
let str = "Hello";
let arr = [1,2,3];
let numb = 5;
let obj = {name: "Chandler"};

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
let iterable = [1,2,3,4];

function createIterator(array: any[]) {
    let count = 0;
    return{
        next: function () {
            return count < array.length?{value:array[count++],done:false}:{value:undefined,done:true}
        }
    }
}
let myIterator = createIterator(iterable);
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
