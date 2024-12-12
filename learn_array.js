// array destructuring
// const stat = (a, b) => {
//     return [
//         a + b,
//         (a + b) / 2,
//         a - b
//     ]
// }

// // stat = 0

// // console.log(stat)

// [sum, average, diff] = stat(20,15)

// console.log(sum, average, diff)

// //
// // function summ(...number) {
// //     let total = 0
// //     for (const a of number) {
// //         total += a;
// //     }
// //     return total;
// // }

// const summ = (...number) => {
//     let total = 0
//     for (const a of number) {
//         total += a;
//     }
//     return total;
// }

// let result = summ(2,3,4)

// console.log(result);

// E6
const plus = (...data) => {
    console.log(data)
    // return data
    //   .filter(function (e) {
    //     return typeof e === 'number';
    //   })
    //   .reduce(function (prev, curr) {
    //     return prev + curr;
    //   });

    // let filter = data.filter((e) => typeof e === 'number')
    // console.log(filter)

    // reduce = filter.reduce((prev, curr) => prev + curr)

    // return reduce

    return data
      .filter((e) => typeof e === 'number')
      .reduce((prev, curr) => prev + curr)
}

result = plus(10,'Hi',null,undefined,20); 
console.log(result)

