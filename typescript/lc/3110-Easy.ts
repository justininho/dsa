/*
You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.

Return the score of s.
*/

function scoreOfString(s: string): number {
    let sum = 0;
    for(let i = 1; i < s.length; i+=1) {
        const diff = s.charCodeAt(i) - s.charCodeAt(i-1);
        sum += Math.abs(diff)
    }
    return sum
};

console.log(scoreOfString("hello"));