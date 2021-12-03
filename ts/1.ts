interface GenericNumber<T> {
    zeroValue: T
    add: (x: T, y: T) => T
}

class Add implements GenericNumber<number> {
    zeroValue!: number
    constructor() { }
    add(x: number, y: number): number {
        return x + y
    }
}

interface Person {
    name: string
    age: number
    location: string
}
type K3 = keyof { [x: string]: Person }


function copyFields<T extends U, U>(target: T, source: U): T {
    for (let id in source) {
        target[id] = (<T>source)[id];
    }

    return target
}

let x = { a: 1, b: 2, c: 3, d: 4 }
copyFields(x, { b: 10, d: 20 })
let a: Add = new Add()
console.log(typeof a)

function returnSomething(): string {
    return ""
}

function* task() {
    let result: ReturnType<typeof returnSomething> = yield returnSomething();
}

type Combine<T> = {
    [P in keyof T]: T[P]
}

type Pagation = {
    Page: Number;
    PageSize: Number;
}

type Query = {
    Status: Number;
}

type Config = Combine<Pagation & Query>

let c: Config = {Status:1,Page:1,PageSize:5}

let d: Record
