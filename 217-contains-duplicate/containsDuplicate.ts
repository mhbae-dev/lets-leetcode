//Brute Force 
// function containsDuplicate(nums: number[]): boolean {
//     for (let i = 0; i < nums.length; i++) {
//         for (let j = i + 1; j < nums.length; j++) { // Fix: Change 'i' to 'j'
//             if (nums[i] === nums[j]) {
//                 return true;
//             }
//         }
//         return false;
//     }
//     return false;
// }

//HashSet
function containsDuplicate(nums: number[]): boolean {
    const set = new Set();

    for(let number of nums){
        if(set.has(number)) return true
        else set.add(number)
    }

    return false;
}

/*
Notes:
-Brute force solution uses a double loop which is more costly in terms of memory and speed.
-The better approach would be to use a set to check first does the set already have the number? Else add it to the set
and just go through the array once.
*/