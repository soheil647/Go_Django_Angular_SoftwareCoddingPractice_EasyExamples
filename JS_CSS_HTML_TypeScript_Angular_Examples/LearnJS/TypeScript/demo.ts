
// *******************************************************************************
//Hello World
function Greet(name: string) {
    let greet = "";
    if(name === "Chandler"){
         greet = "Hello Chandler"
    }
    else {
         greet = "Hi There"
    }
    console.log(greet)
}
Greet("Chandler");

// *******************************************************************************
//For Loop With Let
for (let i=0;i<=5;i++){
    setTimeout(function (){console.log(i)},2000)
}

// *******************************************************************************
//Let vs Const
let num1;
const num2 = 10;

const obj1 = {
    name: "John"
};
console.log(obj1.name);
obj1.name = "Not John";
console.log(obj1.name);

// *******************************************************************************
//Const usage
const Max_Size = 1000;
const PI = 3.14;

// *******************************************************************************
//Let Usage
let a = 5;
let b = 10;

a = a+b;
b = a-b;
a = a-b;

// *******************************************************************************
//Const Usage
const MyFunc = function Ten() {
    return 10;
};
console.log(MyFunc());
const MyFunc2 = (m: number, bounces: number) => 10*m+bounces;
console.log(MyFunc2(5,10));

// *******************************************************************************
//Arrow Function
const employee = {
    id : 23,
    greet : function () {
        console.log(this.id);
        setTimeout(() => console.log(this.id))
    }
};
console.log(employee.greet());

// *******************************************************************************
//Defualt value With Function
let PersentageBounes = () => 0.1;
let getValue = function (value =10, bounes=10*PersentageBounes(), mines = -2) {
    console.log(value+bounes-mines);
    console.log(arguments.length)
};
getValue();
getValue(20);
getValue(20,30);
getValue(20,30,10);
getValue(undefined,undefined,15);

// *******************************************************************************
//Rest And Seperate Operators
let displayColors = function (message: string, ...colors: string[]) {
    console.log(message);
    console.log(colors);
    console.log(arguments.length);

    for(let i in colors) {
        console.log(colors[i])
    }
};

let message = "List of Colors";
let colorArray =["Orange","Black","Cyan"];
displayColors(message, ...colorArray);

displayColors(message, "Red");
displayColors(message, "Red","Blue");
displayColors(message, "Red","Blue","Green");

// *******************************************************************************
//Return nad Obejct Litration 1
let firstName = "Chandler";
let lastName = "Bing";

let person = {
    firstName,
    lastName
};

console.log(person.firstName);
console.log(person.lastName);

function CreatePerson(firstName: string, lastName: string, age: number) {
    let fullName = firstName + "" + lastName;
    return {
        firstName,
        lastName,
        fullName: fullName,
        IsShinier(){
            return age>60
        }
    };
}
let p =CreatePerson("Ross","Geller",32);
console.log(p.firstName);
console.log(p.lastName);
console.log(p.fullName);
console.log(p.IsShinier());

// *******************************************************************************
//Space and Object litration 2
let ln = "last name";
let body = {
    "first name": "Chandler",
    [ln]: "Bing"
};
console.log(body);
console.log(body[ln]);
console.log(body["first name"]);

// *******************************************************************************
//Destructing Object
//let worker = ["Rachel","Green";
let worker = {
    fName: "Rachel",
    lName: "Green",
    gender: "Female",
};
//let [fName, , gender = "Male"] = worker;
let {fName: f, lName: l, gender: g} = worker;
console.log(f);
console.log(l);
console.log(g);

// *******************************************************************************
//String Templates
let user = "Chandler";
let greet = `Welcome 'single' "double" ${user} to ES2015
            this is second line.
                this is so          on.`;
console.log(greet);

// *******************************************************************************
//For of Loops
let fruits = ['banana', 'carrot', 'apple'];
for(let index in fruits) {
    console.log(fruits[index])
}

for(let key of fruits){
    console.log(key)
}

let letter = "ABC";
for(let Char of letter){
    console.log(Char)
}

// *******************************************************************************
//Classes
class Person {
    name: string;
    constructor(name: string){
        this.name = name;
        console.log(this.name + " Constructor")
    }
    static staticMethod(){
        console.log("Static Method")
    }
    greet(){
        console.log("Hello" + this.name)
    }
    getID(){
        return 10;
    }
}
let gate = new Person("Chandler");
console.log(gate.greet() === Person.prototype.greet());
Person.staticMethod();
gate.greet();

//Class inheritance
class Employee extends Person{
    constructor(name: string){
        super(name);
        console.log(name + " drain constructor")
    }
    getID(): number {
        return super.getID();
    }
}
let e = new Employee("Chandler");
console.log(e.getID());

