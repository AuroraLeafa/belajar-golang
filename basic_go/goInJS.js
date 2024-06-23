let arr = ["Senin", "Selasa", "Rabu", "kamis", 'Jumat', "Sabtu", "Minggu"]
let arrSlice = arr.slice(5)
console.log(arrSlice);

arrSlice[0] = "Sabtu baru"
arrSlice[1] = "Minggu baru"

console.log(arr);
console.log(arrSlice);

const factorial = val => {
    if(val === 1) return 1
    return val * factorial(val-1)
}

console.log(factorial(5));

let test = function(){
    let increment = function(){
        return "wtf";
    }

    
}