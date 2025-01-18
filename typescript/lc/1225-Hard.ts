/*
Given a list of words, list of  single letters (might be repeating) and score of every character.

Return the maximum score of any valid set of words formed by using the given letters (words[i] cannot be used two or more times).

It is not necessary to use all characters in letters and each letter can only be used once. Score of letters 'a', 'b', 'c', ... ,'z' is given by score[0], score[1], ... , score[25] respectively.
*/

function getWordScore(word: string, letters: string[], score: number[]): number {
    let wordScore = 0;
    for (let i = 0; i < word.length; i += 1) {
        let letterIndex = letters.findIndex((l) => l === word[i]);
        if (letterIndex !== -1) {
            letters.splice(letterIndex, 1);
            const alphaIndex = word.charCodeAt(i) - 'a'.charCodeAt(0)
            wordScore += score[alphaIndex];
            continue;
        }
        return 0;
    }
    return wordScore;
}

function maxScoreWords(words: string[], letters: string[], score: number[]): number {
    
};

let words = ["dog", "cat", "dad", "good"],
    letters = ["a", "a", "c", "d", "d", "d", "g", "o", "o"],
    score = [1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

console.log('score:', maxScoreWords(words, letters, score)); //23 w/ "dad" and "good"