var x = 10;
var y = "hi";
var z :boolean = true;

var n1 = 10;
var n2 :any = true;

enum EyeColor {Brown=5,Black=11,Blue=17}
let myEyeColor = EyeColor.Black;
console.log(myEyeColor); //11
console.log(EyeColor[myEyeColor]); //Black

let strArray1 : string[] = ["hello","World"];
let anyArray : any[] = ["hello",10,true];

let myTuple :[string, number, number] = ["Hi",100,20];
myTuple[2] = 100;

class PersonJs {
    public fName: string;
    public lName:string;
    constructor(fname:string,lname:string){
        this.fName = fname;
        this.lName = lname;
    }
}

class PersonTypeScript {
    constructor(public fname:string,public lname:string){

    }
}

interface Human {
    FirstName : string,
    LastName : string,
    age? : number
}
let Employee1 : Human = {
    FirstName : "Chandler",
    LastName : "Bing",
    age : 30
};
let Employee2 : Human = {
    FirstName : "Ross",
    LastName : "Geller"
};

@Component({
    selector : 'my-app',
    template : `<h1>Hello World</h1>`
})
export class AppComponent {
    
}