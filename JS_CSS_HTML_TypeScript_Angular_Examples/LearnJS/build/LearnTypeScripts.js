"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
var x = 10;
var y = "hi";
var z = true;
var n1 = 10;
var n2 = true;
var EyeColor;
(function (EyeColor) {
    EyeColor[EyeColor["Brown"] = 5] = "Brown";
    EyeColor[EyeColor["Black"] = 11] = "Black";
    EyeColor[EyeColor["Blue"] = 17] = "Blue";
})(EyeColor || (EyeColor = {}));
var myEyeColor = EyeColor.Black;
console.log(myEyeColor); //11
console.log(EyeColor[myEyeColor]); //Black
var strArray1 = ["hello", "World"];
var anyArray = ["hello", 10, true];
var myTuple = ["Hi", 100, 20];
myTuple[2] = 100;
var PersonJs = /** @class */ (function () {
    function PersonJs(fname, lname) {
        this.fName = fname;
        this.lName = lname;
    }
    return PersonJs;
}());
var PersonTypeScript = /** @class */ (function () {
    function PersonTypeScript(fname, lname) {
        this.fname = fname;
        this.lname = lname;
    }
    return PersonTypeScript;
}());
var Employee1 = {
    FirstName: "Chandler",
    LastName: "Bing",
    age: 30
};
var Employee2 = {
    FirstName: "Ross",
    LastName: "Geller"
};
var AppComponent = /** @class */ (function () {
    function AppComponent() {
    }
    AppComponent = __decorate([
        Component({
            selector: 'my-app',
            template: "<h1>Hello World</h1>"
        })
    ], AppComponent);
    return AppComponent;
}());
exports.AppComponent = AppComponent;
//# sourceMappingURL=../TypeScript/build/LearnTypeScripts.js.map