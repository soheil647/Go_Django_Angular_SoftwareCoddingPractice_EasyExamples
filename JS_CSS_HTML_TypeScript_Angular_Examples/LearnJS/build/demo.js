"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var _a;
// *******************************************************************************
//Hello World
function Greet(name) {
    var greet = "";
    if (name === "Chandler") {
        greet = "Hello Chandler";
    }
    else {
        greet = "Hi There";
    }
    console.log(greet);
}
Greet("Chandler");
var _loop_1 = function (i) {
    setTimeout(function () { console.log(i); }, 2000);
};
// *******************************************************************************
//For Loop With Let
for (var i = 0; i <= 5; i++) {
    _loop_1(i);
}
// *******************************************************************************
//Let vs Const
var num1;
var num2 = 10;
var obj1 = {
    name: "John"
};
console.log(obj1.name);
obj1.name = "Not John";
console.log(obj1.name);
// *******************************************************************************
//Const usage
var Max_Size = 1000;
var PI = 3.14;
// *******************************************************************************
//Let Usage
var a = 5;
var b = 10;
a = a + b;
b = a - b;
a = a - b;
// *******************************************************************************
//Const Usage
var MyFunc = function Ten() {
    return 10;
};
console.log(MyFunc());
var MyFunc2 = function (m, bounces) { return 10 * m + bounces; };
console.log(MyFunc2(5, 10));
// *******************************************************************************
//Arrow Function
var employee = {
    id: 23,
    greet: function () {
        var _this = this;
        console.log(this.id);
        setTimeout(function () { return console.log(_this.id); });
    }
};
console.log(employee.greet());
// *******************************************************************************
//Defualt value With Function
var PersentageBounes = function () { return 0.1; };
var getValue = function (value, bounes, mines) {
    if (value === void 0) { value = 10; }
    if (bounes === void 0) { bounes = 10 * PersentageBounes(); }
    if (mines === void 0) { mines = -2; }
    console.log(value + bounes - mines);
    console.log(arguments.length);
};
getValue();
getValue(20);
getValue(20, 30);
getValue(20, 30, 10);
getValue(undefined, undefined, 15);
// *******************************************************************************
//Rest And Seperate Operators
var displayColors = function (message) {
    var colors = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        colors[_i - 1] = arguments[_i];
    }
    console.log(message);
    console.log(colors);
    console.log(arguments.length);
    for (var i in colors) {
        console.log(colors[i]);
    }
};
var message = "List of Colors";
var colorArray = ["Orange", "Black", "Cyan"];
displayColors.apply(void 0, [message].concat(colorArray));
displayColors(message, "Red");
displayColors(message, "Red", "Blue");
displayColors(message, "Red", "Blue", "Green");
// *******************************************************************************
//Return nad Obejct Litration 1
var firstName = "Chandler";
var lastName = "Bing";
var person = {
    firstName: firstName,
    lastName: lastName
};
console.log(person.firstName);
console.log(person.lastName);
function CreatePerson(firstName, lastName, age) {
    var fullName = firstName + "" + lastName;
    return {
        firstName: firstName,
        lastName: lastName,
        fullName: fullName,
        IsShinier: function () {
            return age > 60;
        }
    };
}
var p = CreatePerson("Ross", "Geller", 32);
console.log(p.firstName);
console.log(p.lastName);
console.log(p.fullName);
console.log(p.IsShinier());
// *******************************************************************************
//Space and Object litration 2
var ln = "last name";
var body = (_a = {
        "first name": "Chandler"
    },
    _a[ln] = "Bing",
    _a);
console.log(body);
console.log(body[ln]);
console.log(body["first name"]);
// *******************************************************************************
//Destructing Object
//let worker = ["Rachel","Green";
var worker = {
    fName: "Rachel",
    lName: "Green",
    gender: "Female",
};
//let [fName, , gender = "Male"] = worker;
var f = worker.fName, l = worker.lName, g = worker.gender;
console.log(f);
console.log(l);
console.log(g);
// *******************************************************************************
//String Templates
var user = "Chandler";
var greet = "Welcome 'single' \"double\" " + user + " to ES2015\n            this is second line.\n                this is so          on.";
console.log(greet);
// *******************************************************************************
//For of Loops
var fruits = ['banana', 'carrot', 'apple'];
for (var index in fruits) {
    console.log(fruits[index]);
}
for (var _i = 0, fruits_1 = fruits; _i < fruits_1.length; _i++) {
    var key_1 = fruits_1[_i];
    console.log(key_1);
}
var letter = "ABC";
for (var _b = 0, letter_1 = letter; _b < letter_1.length; _b++) {
    var Char = letter_1[_b];
    console.log(Char);
}
// *******************************************************************************
//Classes
var Person = /** @class */ (function () {
    function Person(name) {
        this.name = name;
        console.log(this.name + " Constructor");
    }
    Person.staticMethod = function () {
        console.log("Static Method");
    };
    Person.prototype.greet = function () {
        console.log("Hello" + this.name);
    };
    Person.prototype.getID = function () {
        return 10;
    };
    return Person;
}());
var gate = new Person("Chandler");
console.log(gate.greet() === Person.prototype.greet());
Person.staticMethod();
gate.greet();
//Class inheritance
var Employee = /** @class */ (function (_super) {
    __extends(Employee, _super);
    function Employee(name) {
        var _this = _super.call(this, name) || this;
        console.log(name + " drain constructor");
        return _this;
    }
    Employee.prototype.getID = function () {
        return _super.prototype.getID.call(this);
    };
    return Employee;
}(Person));
var e = new Employee("Chandler");
console.log(e.getID());
//# sourceMappingURL=../TypeScript/build/demo.js.map