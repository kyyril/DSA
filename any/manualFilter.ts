const numss: number[] = [1, 2, 3, 4, 5, 6, 7, 8];

let odd: number[] = [];
let even: number[] = [];

for (let num of numss) {
  if (num % 2 === 0) {
    even.push(num);
  } else {
    odd.push(num);
  }
}
console.log(odd);
console.log(even);
