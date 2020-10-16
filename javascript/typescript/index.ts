import * as _ from 'lodash';

async function test() {
    console.log('qwer')
}

let num: number;

type Style = 'bold' | 'italic' | 123;

let font: Style = 'bold';

interface Person {
    first: string;  // These two have to be included
    last: string;   
    [key: string]: any // This allows for any additional 
}

function pow(x: number, y: number): string {
    return Math.pow(x, y).toString();
}

const arr: number[] = [];

arr.push(1);
// arr.push('2');

type MyList = [number?, string?, boolean?]

const arr2: MyList = [];

class Observable<T> {
    constructor (public value: T) {}
}

let x: Observable<number>;

